// Copyright 2022 MobiledgeX, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/coreos/etcd/clientv3/concurrency"
	"github.com/jarcoal/httpmock"
	"github.com/edgexr/edge-cloud/cloudcommon"
	"github.com/edgexr/edge-cloud/cloudcommon/node"
	dme "github.com/edgexr/edge-cloud/d-match-engine/dme-proto"
	"github.com/edgexr/edge-cloud/edgeproto"
	"github.com/edgexr/edge-cloud/integration/process"
	"github.com/edgexr/edge-cloud/log"
	"github.com/edgexr/edge-cloud/notify"
	"github.com/edgexr/edge-cloud/objstore"
	"github.com/edgexr/edge-cloud/testutil"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
)

type stateTransition struct {
	triggerState   dme.CloudletState
	triggerVersion string
	expectedState  edgeproto.TrackedState
	ignoreState    bool
}

const (
	crm_v1 = "2001-01-31"
	crm_v2 = "2002-01-31"
)

var eMock *EventMock

type EventMock struct {
	names map[string][]node.EventTag
	addr  string
	mux   sync.Mutex
}

func NewEventMock(addr string) *EventMock {
	event := EventMock{}
	event.addr = addr
	event.names = make(map[string][]node.EventTag)
	return &event
}

func (e *EventMock) registerResponders(t *testing.T) {
	// register mock responders
	api := fmt.Sprintf("%s/_template/events-log", e.addr)
	httpmock.RegisterResponder("PUT", api,
		func(req *http.Request) (*http.Response, error) {
			return httpmock.NewStringResponse(200, "Success"), nil
		},
	)
	recordEvent := func(data []byte) {
		eData := node.EventData{}
		err := json.Unmarshal(data, &eData)
		require.Nil(t, err, "json unmarshal event data")
		require.NotEmpty(t, eData.Name, "event name exists")
		e.mux.Lock()
		e.names[strings.ToLower(eData.Name)] = eData.Tags
		e.mux.Unlock()
	}

	api = fmt.Sprintf("=~%s/events-log-.*/_doc", e.addr)
	httpmock.RegisterResponder("POST", api,
		func(req *http.Request) (*http.Response, error) {
			data, _ := ioutil.ReadAll(req.Body)
			recordEvent(data)
			return httpmock.NewStringResponse(200, "Success"), nil
		},
	)
	api = fmt.Sprintf("=~%s/.*/_bulk", e.addr)
	httpmock.RegisterResponder("POST", api,
		func(req *http.Request) (*http.Response, error) {
			data, _ := ioutil.ReadAll(req.Body)
			lines := strings.Split(string(data), "\n")
			// each record is 2 lines, first line is metadata,
			// second line is data. Final line is blank.
			for ii := 0; ii < len(lines)-1; ii += 2 {
				recordEvent([]byte(lines[ii+1]))
			}
			return httpmock.NewStringResponse(200, "Success"), nil
		},
	)
}

func (e *EventMock) verifyEvent(t *testing.T, name string, tags []node.EventTag) {
	// Events are written in a separate thread so we need to poll
	// to check when they're registered.
	var eTags []node.EventTag
	var ok bool
	for ii := 0; ii < 20; ii++ {
		e.mux.Lock()
		eTags, ok = e.names[strings.ToLower(name)]
		e.mux.Unlock()
		if ok {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	require.True(t, ok, "event exists")
	require.NotEqual(t, len(eTags), 0, "there should be more than 0 tags")
	require.NotEqual(t, len(tags), 0, "there should be more than 0 tags")
	eTagsMap := make(map[string]string)
	for _, eTag := range eTags {
		eTagsMap[eTag.Key] = eTag.Value
	}
	for _, tag := range tags {
		val, ok := eTagsMap[tag.Key]
		require.True(t, ok, "tag key exists")
		require.Equal(t, val, tag.Value, "tag value matches")
	}
}

func TestCloudletApi(t *testing.T) {
	log.SetDebugLevel(log.DebugLevelEtcd | log.DebugLevelApi | log.DebugLevelNotify | log.DebugLevelEvents)
	log.InitTracer(nil)
	defer log.FinishTracer()
	ctx := log.StartTestSpan(context.Background())

	testSvcs := testinit(ctx, t)
	defer testfinish(testSvcs)

	dummy := dummyEtcd{}
	dummy.Start()
	defer dummy.Stop()

	sync := InitSync(&dummy)
	apis := NewAllApis(sync)
	sync.Start()
	defer sync.Done()

	// mock http to redirect requests
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	// any requests that don't have a registered URL will be fetched normally
	httpmock.RegisterNoResponder(httpmock.InitialTransport.RoundTrip)

	esURL := "http://dummy-es"
	eMock = NewEventMock(esURL)
	eMock.registerResponders(t)

	// setup nodeMgr for events
	nodeMgr = node.NodeMgr{}
	ctx, _, err := nodeMgr.Init(node.NodeTypeController, "", node.WithRegion("unit-test"),
		node.WithESUrls(esURL))
	require.Nil(t, err)
	require.NotNil(t, nodeMgr.ESClient)
	defer nodeMgr.Finish()

	// create flavors
	cloudletData := testutil.CloudletData()
	testutil.InternalFlavorCreate(t, apis.flavorApi, testutil.FlavorData)
	testutil.InternalGPUDriverTest(t, "cud", apis.gpuDriverApi, testutil.GPUDriverData)
	testutil.InternalResTagTableCreate(t, apis.resTagTableApi, testutil.ResTagTableData)
	testutil.InternalCloudletTest(t, "cud", apis.cloudletApi, cloudletData)

	// test invalid location values
	clbad := cloudletData[0]
	clbad.Key.Name = "bad loc"
	testBadLat(t, ctx, &clbad, []float64{90.1, -90.1, -1323213, 1232334}, "create", apis)
	testBadLong(t, ctx, &clbad, []float64{180.1, -180.1, -1323213, 1232334}, "create", apis)

	clbad = cloudletData[0]
	clbad.Key.Name = "test num dyn ips"
	err = apis.cloudletApi.CreateCloudlet(&clbad, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err)
	clbad.NumDynamicIps = 0
	clbad.Fields = []string{edgeproto.CloudletFieldNumDynamicIps}
	err = apis.cloudletApi.UpdateCloudlet(&clbad, testutil.NewCudStreamoutCloudlet(ctx))
	require.NotNil(t, err)

	cl := cloudletData[1]
	cl.Key.Name = "test invalid lat-long"
	err = apis.cloudletApi.CreateCloudlet(&cl, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err)
	testBadLat(t, ctx, &cl, []float64{90.1, -90.1, -1323213, 1232334}, "update", apis)
	testBadLong(t, ctx, &cl, []float64{180.1, -180.1, -1323213, 1232334}, "update", apis)

	testCloudletDnsLabel(t, ctx, apis)

	// Resource Mapping tests
	testResMapKeysApi(t, ctx, &cl, apis)
	testGpuResourceMapping(t, ctx, &cl, apis)

	// Cloudlet state tests
	testCloudletStates(t, ctx, apis)
	testManualBringup(t, ctx, apis)

	testShowFlavorsForCloudlet(t, ctx, apis)
	testAllianceOrgs(t, ctx, apis)
}

func testBadLat(t *testing.T, ctx context.Context, clbad *edgeproto.Cloudlet, lats []float64, action string, apis *AllApis) {
	for _, lat := range lats {
		clbad.Location.Latitude = lat
		clbad.Fields = []string{edgeproto.CloudletFieldLocationLatitude}
		switch action {
		case "create":
			err := apis.cloudletApi.CreateCloudlet(clbad, testutil.NewCudStreamoutCloudlet(ctx))
			require.NotNil(t, err, "create cloudlet bad latitude")
		case "update":
			err := apis.cloudletApi.UpdateCloudlet(clbad, testutil.NewCudStreamoutCloudlet(ctx))
			require.NotNil(t, err, "update cloudlet bad latitude")
		}
	}
}

func testBadLong(t *testing.T, ctx context.Context, clbad *edgeproto.Cloudlet, longs []float64, action string, apis *AllApis) {
	for _, long := range longs {
		clbad.Location.Longitude = long
		clbad.Fields = []string{edgeproto.CloudletFieldLocationLongitude}
		switch action {
		case "create":
			err := apis.cloudletApi.CreateCloudlet(clbad, testutil.NewCudStreamoutCloudlet(ctx))
			require.NotNil(t, err, "create cloudlet bad longitude")
		case "update":
			err := apis.cloudletApi.CreateCloudlet(clbad, testutil.NewCudStreamoutCloudlet(ctx))
			require.NotNil(t, err, "update cloudlet bad longitude")
		}
	}
}

func waitForState(key *edgeproto.CloudletKey, state edgeproto.TrackedState, apis *AllApis) error {
	lastState := edgeproto.TrackedState_TRACKED_STATE_UNKNOWN
	for i := 0; i < 10; i++ {
		cloudlet := edgeproto.Cloudlet{}
		if apis.cloudletApi.cache.Get(key, &cloudlet) {
			if cloudlet.State == state {
				return nil
			}
			lastState = cloudlet.State
		}
		time.Sleep(10 * time.Millisecond)
	}

	return fmt.Errorf("Unable to get desired cloudlet state, actual state %s, desired state %s", lastState, state)
}

func forceCloudletInfoState(ctx context.Context, key *edgeproto.CloudletKey, state dme.CloudletState, taskName, version string, apis *AllApis) {
	info := edgeproto.CloudletInfo{}
	info.Key = *key
	info.State = state
	info.ContainerVersion = version
	info.Status.SetTask(taskName)
	apis.cloudletInfoApi.Update(ctx, &info, 0)
}

func forceCloudletInfoMaintenanceState(ctx context.Context, key *edgeproto.CloudletKey, state dme.MaintenanceState, apis *AllApis) {
	info := edgeproto.CloudletInfo{}
	if !apis.cloudletInfoApi.cache.Get(key, &info) {
		info.Key = *key
	}
	info.MaintenanceState = state
	apis.cloudletInfoApi.Update(ctx, &info, 0)
}

func deleteCloudletInfo(ctx context.Context, key *edgeproto.CloudletKey, apis *AllApis) {
	info := edgeproto.CloudletInfo{}
	info.Key = *key
	apis.cloudletInfoApi.Delete(ctx, &info, 0)
}

func testNotifyId(t *testing.T, ctrlHandler *notify.DummyHandler, key *edgeproto.CloudletKey, nodeCount, notifyId int, crmVersion string) {
	require.Equal(t, nodeCount, len(ctrlHandler.NodeCache.Objs), "node count matches")
	nodeVersion, nodeNotifyId, err := ctrlHandler.GetCloudletDetails(key)
	require.Nil(t, err, "get cloudlet version & notifyId from node cache")
	require.Equal(t, crmVersion, nodeVersion, "node version matches")
	require.Equal(t, int64(notifyId), nodeNotifyId, "node notifyId matches")
}

func testCloudletStates(t *testing.T, ctx context.Context, apis *AllApis) {
	ctrlHandler := notify.NewDummyHandler()
	ctrlMgr := notify.ServerMgr{}
	ctrlHandler.RegisterServer(&ctrlMgr)
	ctrlMgr.Start("ctrl", "127.0.0.1:50001", nil)
	defer ctrlMgr.Stop()

	getPublicCertApi := &cloudcommon.TestPublicCertApi{}
	publicCertManager, err := node.NewPublicCertManager("localhost", getPublicCertApi, "", "")
	require.Nil(t, err)
	tlsConfig, err := publicCertManager.GetServerTlsConfig(ctx)
	require.Nil(t, err)
	err = services.accessKeyGrpcServer.Start(*accessApiAddr, apis.cloudletApi.accessKeyServer, tlsConfig, func(accessServer *grpc.Server) {
		edgeproto.RegisterCloudletAccessApiServer(accessServer, apis.cloudletApi)
		edgeproto.RegisterCloudletAccessKeyApiServer(accessServer, apis.cloudletApi)
	})
	require.Nil(t, err, "start access server")
	defer services.accessKeyGrpcServer.Stop()

	crm_notifyaddr := "127.0.0.1:0"
	cloudlet := testutil.CloudletData()[2]
	cloudlet.ContainerVersion = crm_v1
	cloudlet.Key.Name = "testcloudletstates"
	cloudlet.NotifySrvAddr = crm_notifyaddr
	cloudlet.CrmOverride = edgeproto.CRMOverride_NO_OVERRIDE
	pfConfig, err := apis.cloudletApi.getPlatformConfig(ctx, &cloudlet)
	require.Nil(t, err, "get platform config")
	pfConfig.EnvVar["E2ETEST_TLS"] = "true"

	err = apis.cloudletApi.CreateCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err, "create cloudlet")
	defer apis.cloudletApi.DeleteCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	res, err := apis.cloudletApi.GenerateAccessKey(ctx, &cloudlet.Key)
	require.Nil(t, err, "generate access key")
	pfConfig.CrmAccessPrivateKey = res.Message
	pfConfig.AccessApiAddr = services.accessKeyGrpcServer.ApiAddr()

	streamCloudlet := NewStreamoutMsg(ctx)
	go func() {
		// copy objects required for WatchKey on cloudletInfo
		apis.cloudletInfoApi.cache.Objs = ctrlHandler.CloudletInfoCache.Objs
		apis.cloudletInfoApi.cache.KeyWatchers = ctrlHandler.CloudletInfoCache.KeyWatchers
		// setup cloudlet stream
		err = apis.streamObjApi.StreamCloudlet(&cloudlet.Key, streamCloudlet)
		require.Nil(t, err, "stream cloudlet")
	}()

	err = cloudcommon.StartCRMService(ctx, &cloudlet, pfConfig, process.HARolePrimary, nil)
	require.Nil(t, err, "start cloudlet")
	defer func() {
		// Delete CRM
		err = cloudcommon.StopCRMService(ctx, &cloudlet, process.HARolePrimary)
		require.Nil(t, err, "stop cloudlet")
	}()

	err = ctrlHandler.WaitForCloudletState(&cloudlet.Key, dme.CloudletState_CLOUDLET_STATE_INIT)
	require.Nil(t, err, "cloudlet state transition")

	cloudlet.State = edgeproto.TrackedState_CRM_INITOK
	ctrlHandler.CloudletCache.Update(ctx, &cloudlet, 0)

	err = ctrlHandler.WaitForCloudletState(&cloudlet.Key, dme.CloudletState_CLOUDLET_STATE_READY)
	require.Nil(t, err, "cloudlet state transition")

	cloudlet.State = edgeproto.TrackedState_READY
	ctrlHandler.CloudletCache.Update(ctx, &cloudlet, 0)

	require.Equal(t, len(streamCloudlet.Msgs), 5, "progress messages")
	cloudletMsgs := []string{"Setting up cloudlet", "Initializing platform", "Done initializing fake platform", "Gathering Cloudlet Info", "Cloudlet setup successfully"}
	for ii, msg := range cloudletMsgs {
		require.Equal(t, streamCloudlet.Msgs[ii].Message, msg, "message matches")
	}

	cloudlet.State = edgeproto.TrackedState_UPDATE_REQUESTED
	ctrlHandler.CloudletCache.Update(ctx, &cloudlet, 0)

	err = ctrlHandler.WaitForCloudletState(&cloudlet.Key, dme.CloudletState_CLOUDLET_STATE_UPGRADE)
	require.Nil(t, err, "cloudlet state transition")

	cloudlet.State = edgeproto.TrackedState_UPDATING
	ctrlHandler.CloudletCache.Update(ctx, &cloudlet, 0)

	err = ctrlHandler.WaitForCloudletState(&cloudlet.Key, dme.CloudletState_CLOUDLET_STATE_READY)
	require.Nil(t, err, "cloudlet state transition")

	cloudletInfo := edgeproto.CloudletInfo{}
	found := ctrlHandler.CloudletInfoCache.Get(&cloudlet.Key, &cloudletInfo)
	require.True(t, found, "cloudlet info exists")
	require.Equal(t, len(cloudletInfo.ResourcesSnapshot.Info), 4, "cloudlet resources info exists")
	for _, resInfo := range cloudletInfo.ResourcesSnapshot.Info {
		switch resInfo.Name {
		case cloudcommon.ResourceRamMb:
			require.Equal(t, resInfo.Value, uint64(8192), "cloudlet resources info exists")
		case cloudcommon.ResourceVcpus:
			require.Equal(t, resInfo.Value, uint64(4), "cloudlet resources info exists")
		case cloudcommon.ResourceExternalIPs:
			require.Equal(t, resInfo.Value, uint64(1), "cloudlet resources info exists")
		case cloudcommon.ResourceInstances:
			require.Equal(t, resInfo.Value, uint64(2), "cloudlet resources info exists")
		default:
			require.True(t, false, fmt.Sprintf("invalid resinfo name: %s", resInfo.Name))
		}
	}
}

func testManualBringup(t *testing.T, ctx context.Context, apis *AllApis) {
	var err error
	cloudlet := testutil.CloudletData()[2]
	cloudlet.Key.Name = "crmmanualbringup"
	cloudlet.ContainerVersion = crm_v1
	err = apis.cloudletApi.CreateCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err)

	err = waitForState(&cloudlet.Key, edgeproto.TrackedState_READY, apis)
	require.Nil(t, err, "cloudlet obj created")

	forceCloudletInfoState(ctx, &cloudlet.Key, dme.CloudletState_CLOUDLET_STATE_INIT, "sending init", crm_v2, apis)
	err = waitForState(&cloudlet.Key, edgeproto.TrackedState_CRM_INITOK, apis)
	require.Nil(t, err, fmt.Sprintf("cloudlet state transtions"))
	eMock.verifyEvent(t, "upgrading cloudlet", []node.EventTag{
		node.EventTag{
			Key:   "from-version",
			Value: crm_v1,
		},
		node.EventTag{
			Key:   "to-version",
			Value: crm_v2,
		},
	})

	forceCloudletInfoState(ctx, &cloudlet.Key, dme.CloudletState_CLOUDLET_STATE_READY, "sending ready", crm_v2, apis)
	err = waitForState(&cloudlet.Key, edgeproto.TrackedState_READY, apis)
	require.Nil(t, err, fmt.Sprintf("cloudlet state transtions"))
	eMock.verifyEvent(t, "cloudlet online", []node.EventTag{
		node.EventTag{
			Key:   "state",
			Value: "CLOUDLET_STATE_READY",
		},
		node.EventTag{
			Key:   "version",
			Value: crm_v2,
		},
	})

	stateTransitions := map[dme.MaintenanceState]dme.MaintenanceState{
		dme.MaintenanceState_FAILOVER_REQUESTED:    dme.MaintenanceState_FAILOVER_DONE,
		dme.MaintenanceState_CRM_REQUESTED:         dme.MaintenanceState_CRM_UNDER_MAINTENANCE,
		dme.MaintenanceState_NORMAL_OPERATION_INIT: dme.MaintenanceState_NORMAL_OPERATION,
	}

	cancel := apis.cloudletApi.cache.WatchKey(&cloudlet.Key, func(ctx context.Context) {
		cl := edgeproto.Cloudlet{}
		if !apis.cloudletApi.cache.Get(&cloudlet.Key, &cl) {
			return
		}
		switch cl.MaintenanceState {
		case dme.MaintenanceState_FAILOVER_REQUESTED:
			info := edgeproto.AutoProvInfo{}
			if !apis.autoProvInfoApi.cache.Get(&cloudlet.Key, &info) {
				info.Key = cloudlet.Key
			}
			info.MaintenanceState = stateTransitions[cl.MaintenanceState]
			apis.autoProvInfoApi.cache.Update(ctx, &info, 0)
		case dme.MaintenanceState_CRM_REQUESTED:
			fallthrough
		case dme.MaintenanceState_NORMAL_OPERATION_INIT:
			info := edgeproto.CloudletInfo{}
			if !apis.cloudletInfoApi.cache.Get(&cloudlet.Key, &info) {
				info.Key = cloudlet.Key
			}
			info.MaintenanceState = stateTransitions[cl.MaintenanceState]
			apis.cloudletInfoApi.cache.Update(ctx, &info, 0)
		}
	})

	defer cancel()

	cloudlet.MaintenanceState = dme.MaintenanceState_MAINTENANCE_START
	cloudlet.Fields = append(cloudlet.Fields, edgeproto.CloudletFieldMaintenanceState)
	err = apis.cloudletApi.UpdateCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err, fmt.Sprintf("update cloudlet maintenance state"))

	eMock.verifyEvent(t, "cloudlet maintenance start", []node.EventTag{
		node.EventTag{
			Key:   "maintenance-state",
			Value: "UNDER_MAINTENANCE",
		},
	})

	cloudlet.MaintenanceState = dme.MaintenanceState_NORMAL_OPERATION
	err = apis.cloudletApi.UpdateCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err, fmt.Sprintf("update cloudlet maintenance state"))
	eMock.verifyEvent(t, "cloudlet maintenance done", []node.EventTag{
		node.EventTag{
			Key:   "maintenance-state",
			Value: "NORMAL_OPERATION",
		},
	})

	deleteCloudletInfo(ctx, &cloudlet.Key, apis)
	eMock.verifyEvent(t, "cloudlet offline", []node.EventTag{
		node.EventTag{
			Key:   "reason",
			Value: "notify disconnect",
		},
	})

	// Cloudlet state is INITOK but from old CRM (crm_v1)
	forceCloudletInfoState(ctx, &cloudlet.Key, dme.CloudletState_CLOUDLET_STATE_INIT, "sending init", crm_v1, apis)
	err = waitForState(&cloudlet.Key, edgeproto.TrackedState_CRM_INITOK, apis)
	require.Nil(t, err, fmt.Sprintf("cloudlet state transtions"))

	// Cloudlet should still be ready, ignoring the above stale entry
	forceCloudletInfoState(ctx, &cloudlet.Key, dme.CloudletState_CLOUDLET_STATE_READY, "sending ready", crm_v2, apis)
	err = waitForState(&cloudlet.Key, edgeproto.TrackedState_READY, apis)
	require.Nil(t, err, fmt.Sprintf("cloudlet state transtions"))

	found := apis.autoProvInfoApi.cache.Get(&cloudlet.Key, &edgeproto.AutoProvInfo{})
	require.True(t, found, "autoprovinfo for cloudlet exists")

	err = apis.cloudletApi.DeleteCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err)

	found = apis.autoProvInfoApi.cache.Get(&cloudlet.Key, &edgeproto.AutoProvInfo{})
	require.False(t, found, "autoprovinfo for cloudlet should be cleaned up")
}

func testResMapKeysApi(t *testing.T, ctx context.Context, cl *edgeproto.Cloudlet, apis *AllApis) {
	// We can add/remove edgeproto.ResTagTableKey values to the cl.ResTagMap map
	// which then can be used in the GetVMSpec call when matching our meta-resource specificer
	// to a deployments actual resources/flavrs.
	resmap := edgeproto.CloudletResMap{}
	resmap.Key = cl.Key
	// test_data contains sample resource tag maps, add them to the cloudlet
	// verify, and remove them. ClI should follow suit.
	if cl.ResTagMap == nil {
		cl.ResTagMap = make(map[string]*edgeproto.ResTagTableKey)
	}
	if resmap.Mapping == nil {
		resmap.Mapping = make(map[string]string)
	}

	// use the OptResNames as clould.ResTagMap[key] = tblkey in test
	// gpu, nas and nic are the current set of Resource Names.
	// setup the test map using the test_data objects
	// The AddCloudResMapKey is setup to accept multiple res tbl keys at once
	// but we're doing it one by one.

	resmap.Mapping[strings.ToLower(edgeproto.OptResNames_name[0])] = testutil.Restblkeys[0].Name
	_, err := apis.cloudletApi.AddCloudletResMapping(ctx, &resmap)
	require.Nil(t, err, "AddCloudletResMapKey")

	resmap.Mapping[strings.ToLower(edgeproto.OptResNames_name[1])] = testutil.Restblkeys[1].Name
	_, err = apis.cloudletApi.AddCloudletResMapping(ctx, &resmap)
	require.Nil(t, err, "AddCloudletResMapKey")

	resmap.Mapping[strings.ToLower(edgeproto.OptResNames_name[2])] = testutil.Restblkeys[2].Name
	_, err = apis.cloudletApi.AddCloudletResMapping(ctx, &resmap)
	require.Nil(t, err, "AddCloudletResMapKey")

	testcl := &edgeproto.Cloudlet{}
	// now it's all stored, fetch a copy of the cloudlet and verify
	err = apis.cloudletApi.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		if !apis.cloudletApi.store.STMGet(stm, &cl.Key, testcl) {
			return cl.Key.NotFoundError()
		}
		return err
	})

	// what's in our testcl? Check the resource map
	tkey := testcl.ResTagMap[strings.ToLower(edgeproto.OptResNames_name[0])]
	require.Equal(t, testutil.Restblkeys[0].Name, tkey.Name, "AddCloudletResMapKey")
	tkey = testcl.ResTagMap[strings.ToLower(edgeproto.OptResNames_name[1])]
	require.Equal(t, testutil.Restblkeys[1].Name, tkey.Name, "AddCloudletResMapKey")
	tkey = testcl.ResTagMap[strings.ToLower(edgeproto.OptResNames_name[2])]
	require.Equal(t, testutil.Restblkeys[2].Name, tkey.Name, "AddCloudletResMapKey")

	// and the actual keys should match as well
	require.Equal(t, testutil.Restblkeys[0], *testcl.ResTagMap[testutil.Restblkeys[0].Name], "AddCloudletResMapKey")
	require.Equal(t, testutil.Restblkeys[1], *testcl.ResTagMap[testutil.Restblkeys[1].Name], "AddCloudletResMapKey")
	require.Equal(t, testutil.Restblkeys[2], *testcl.ResTagMap[testutil.Restblkeys[2].Name], "AddCloudletResMapKey")

	resmap1 := edgeproto.CloudletResMap{}
	resmap1.Mapping = make(map[string]string)
	resmap1.Mapping[strings.ToLower(edgeproto.OptResNames_name[2])] = testutil.Restblkeys[2].Name
	resmap1.Mapping[strings.ToLower(edgeproto.OptResNames_name[1])] = testutil.Restblkeys[1].Name
	resmap1.Key = cl.Key

	_, err = apis.cloudletApi.RemoveCloudletResMapping(ctx, &resmap1)
	require.Nil(t, err, "RemoveCloudletResMapKey")

	rmcl := &edgeproto.Cloudlet{}
	if rmcl.ResTagMap == nil {
		rmcl.ResTagMap = make(map[string]*edgeproto.ResTagTableKey)
	}
	rmcl.Key = resmap1.Key

	err = apis.cloudletApi.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {
		if !apis.cloudletApi.store.STMGet(stm, &cl.Key, rmcl) {
			return cl.Key.NotFoundError()
		}
		return err
	})

	require.Nil(t, err, "STMGet failure")
	// and check the maps len = 1
	require.Equal(t, 1, len(rmcl.ResTagMap), "RemoveCloudletResMapKey")
	// and might as well check the key "gpu" exists
	_, ok := rmcl.ResTagMap[testutil.Restblkeys[0].Name]
	require.Equal(t, true, ok, "RemoveCloudletResMapKey")
}

func testGpuResourceMapping(t *testing.T, ctx context.Context, cl *edgeproto.Cloudlet, apis *AllApis) {
	// Cloudlet has a map key'ed by resource name/type whose value is a res tag tbl key.
	// We init this map, and create a resource table, and place its key into this map
	// and pass this map to the matcher routine, this allows the matcher to have access
	// to all optional resource tag maps present in the cloudlet. A meta-flavor has a
	// similar map to request generic resources that need to be mapped to specific
	// platform resources. We create such a edgeproto.Flavor and set it's request
	// map to ask for a gpu and a nas storage volume. The game for the matcher/mapper
	// is to take our meta-flavor resourse request object, and return, for this
	// operator/cloudlet the closest matching available flavor to use in the eventual
	// launch of a suitable image.
	var cli edgeproto.CloudletInfo = testutil.CloudletInfoData[0]

	if cl.ResTagMap == nil {
		cl.ResTagMap = make(map[string]*edgeproto.ResTagTableKey)
	}
	var gputab = edgeproto.ResTagTable{
		Key: edgeproto.ResTagTableKey{
			Name: "gpumap",
		},
		Tags: map[string]string{"vgpu": "nvidia-63:1", "pci": "t4:1", "gpu": "T4:1", "vmware": "vgpu=1", "resources": "VGPU=1", "pci_": "alias=t4gpu:1"},
	}

	var nastab = edgeproto.ResTagTable{
		Key: edgeproto.ResTagTableKey{
			Name: "nasmap",
		},
		Tags: map[string]string{"nas": "ceph-20:1"},
	}
	_, err := apis.resTagTableApi.CreateResTagTable(ctx, &gputab)
	require.Nil(t, nil, err, "CreateResTagTable")

	// Our clouldets resource map, maps from resource type names, to ResTagTableKeys.
	// The ResTagTableKey is a resource name, and the owning operator key.
	cl.ResTagMap["gpu"] = &gputab.Key

	// We also  need a list of edgeproto.FlavorInfo structs
	// which it so happens we have in the testutils.CloudletInfoData.Flavors array
	tbl1, err := apis.resTagTableApi.GetResTagTable(ctx, &gputab.Key)
	require.Nil(t, err, "GetResTagTable")
	require.Equal(t, 6, len(tbl1.Tags), "tag count mismatch")

	// specify a pci pass_throuh, don't care what kind
	// should match flavor.large-pci
	var flavorPciMatch = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "x1.large-pci-mex",
		},
		Ram:   8192,
		Vcpus: 10,
		Disk:  40,
		// This requests a passthru
		OptResMap: map[string]string{"gpu": "pci:1"},
	}

	// map to a generic nvidia vgpu type, should match flavor.large-nvidia
	var flavorVgpuMatch = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "x1.large-vgpu-mex",
		},
		Ram:   8192,
		Vcpus: 10,
		Disk:  40,
		// This requests 1 vgpu instances, (not supported by nvidia yet)
		OptResMap: map[string]string{"gpu": "vgpu:1"},
	}
	// don't care what kind of gpu resource

	// don't care what kind of gpu resource
	var testflavor = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "x1.large-mex",
		},
		Ram:   8192,
		Vcpus: 8,
		Disk:  40,
		// This says I want one gpu, don't care if it's vgpu or passthrough
		OptResMap: map[string]string{"gpu": "gpu:1"},
	}
	// request two optional resources
	var testflavor2 = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "x1.large-2-Resources",
		},
		Ram:   8192,
		Vcpus: 8,
		Disk:  40,
		// This says I want one gpu, don't care if it's vgpu or passthrough
		OptResMap: map[string]string{"gpu": "gpu:1", "nas": "nas:ceph-20:1"},
	}
	// request nas optional resource only
	var testflavorNas = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "x1.large-2-Resources",
		},
		Ram:       8192,
		Vcpus:     8,
		Disk:      40,
		OptResMap: map[string]string{"nas": "nas:ceph-20:1"},
	}

	// test request for a specific type of pci  ( one T4 )
	// should match flavor.large from testutils.
	var testPciT4flavor = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "mex.large-pci-T4",
		},
		Ram:   8192,
		Vcpus: 8,
		Disk:  40,
		// This says I want one gpu of kind pci:t4
		OptResMap: map[string]string{"gpu": "pci:t4:1"},
	}

	var flavorVgpuNvidiaMatch = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "mex.large-vgpu-nvidia-63",
		},
		Ram:   8192,
		Vcpus: 10,
		Disk:  40,
		// This requests 1 vgpu instance of spec nvidia-63
		OptResMap: map[string]string{"gpu": "vgpu:nvidia-63:1"},
	}

	// should match flavor.large-generic-gpu
	var flavorVIOMatch = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "x1.large-vmware-vgpu",
		},
		Ram:   8192,
		Vcpus: 10,
		Disk:  80,
		// This requests a 1 vgpu instance of any kind
		OptResMap: map[string]string{"gpu": "vmware=vgpu=1"},
	}

	// Two mex flavors differing only in GPU vs VGPU
	var flavorT4VGPUMatch = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "mex.large-vgpuT48Q",
		},
		Ram:   4096,
		Vcpus: 12,
		Disk:  20,
		// This requests a vgpu
		OptResMap: map[string]string{"gpu": "resources:VGPU:1"},
	}

	var flavorT4GPUMatch = edgeproto.Flavor{
		Key: edgeproto.FlavorKey{
			Name: "mex.large-gpuT48Q",
		},
		Ram:   4096,
		Vcpus: 12,
		Disk:  20,
		// This requests a vgpu
		OptResMap: map[string]string{"gpu": "pci_:alias=t4gpu:1"},
	}

	taz := edgeproto.OSAZone{Name: "AZ1_GPU", Status: "available"}
	timg := edgeproto.OSImage{Name: "gpu_image"}
	cli.AvailabilityZones = append(cli.AvailabilityZones, &taz)
	cli.OsImages = append(cli.OsImages, &timg)

	// testflavor wants some generic GPU resource, it should match
	// the first flavor offering some type of gpu reosurce.
	// We can direct a generic request to a given flavor though,
	// which is the case here.

	err = apis.cloudletApi.sync.ApplySTMWait(ctx, func(stm concurrency.STM) error {

		spec, vmerr := apis.resTagTableApi.GetVMSpec(ctx, stm, testflavor, *cl, cli)
		require.Nil(t, vmerr, "GetVmSpec")
		require.Equal(t, "flavor.large", spec.FlavorName)
		require.Equal(t, "AZ1_GPU", spec.AvailabilityZone)
		require.Equal(t, "gpu_image", spec.ImageName)

		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, flavorVgpuMatch, *cl, cli)
		require.Nil(t, vmerr, "GetVmSpec vgpu request")
		require.Equal(t, "flavor.large-nvidia", spec.FlavorName)

		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, flavorPciMatch, *cl, cli)
		require.Nil(t, vmerr, "GetVMSpec")
		require.Equal(t, "flavor.large", spec.FlavorName)

		// non-nominal, ask for more resources than the would-be match supports.
		// change testflavor to request 10 gpus of any kind.
		testflavor.OptResMap["gpu"] = "gpu:10"
		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, testflavor, *cl, cli)
		require.Equal(t, "no suitable platform flavor found for x1.large-mex, please try a smaller flavor", vmerr.Error(), "nil table")

		// specific pci passthrough
		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, testPciT4flavor, *cl, cli)
		require.Nil(t, vmerr, "GetVmSpec")
		require.Equal(t, "flavor.large", spec.FlavorName)

		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, flavorVgpuNvidiaMatch, *cl, cli)
		require.Nil(t, vmerr, "GetVmSpec")
		require.Equal(t, "flavor.large-nvidia", spec.FlavorName)
		uses := apis.resTagTableApi.UsesGpu(ctx, stm, *spec.FlavorInfo, *cl)
		require.Equal(t, true, uses)

		// vmware vio syntax
		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, flavorVIOMatch, *cl, cli)
		require.Nil(t, vmerr, "GetVmSpec")
		require.Equal(t, "flavor.large-generic-gpu", spec.FlavorName)

		// Now try 2 optional resources requested by one flavor, first non-nominal, no res tag table for nas tags
		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, testflavor2, *cl, cli)
		if vmerr != nil {
			require.Equal(t, "no suitable platform flavor found for x1.large-2-Resources, please try a smaller flavor", vmerr.Error())
		}

		// now, add cloudlet mapping for nas to the cloudlet, making the above test nominal...
		cl.ResTagMap["nas"] = &nastab.Key

		// ...and actually create the new nas res tag table
		_, err := apis.resTagTableApi.CreateResTagTable(ctx, &nastab)
		require.Nil(t, err)

		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, testflavor2, *cl, cli)
		require.Nil(t, vmerr, "GetVMSpec")
		require.Equal(t, "flavor.large2", spec.FlavorName)

		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, flavorT4VGPUMatch, *cl, cli)
		require.Nil(t, vmerr, "GetVMSpec")
		require.Equal(t, "flavor.m4.large-vgpu", spec.FlavorName)

		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, flavorT4GPUMatch, *cl, cli)
		require.Nil(t, vmerr, "GetVMSpec")
		require.Equal(t, "flavor.m4.large-gpu", spec.FlavorName)

		// Non-nominal: ask for nas only, should reject testflavor2 as there are no
		// os flavors with only a nas resource
		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, testflavorNas, *cl, cli)
		require.Equal(t, "no suitable platform flavor found for x1.large-2-Resources, please try a smaller flavor", vmerr.Error())
		// Non-nominal: flavor requests optional resource, while cloudlet's OptResMap is nil (cloudlet supports none)
		cl.ResTagMap = nil
		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, testflavor, *cl, cli)
		require.Equal(t, "Cloudlet San Jose Site doesn't support GPU", vmerr.Error())

		nulCL := edgeproto.Cloudlet{}
		// and finally, Non-nominal, request a resource, and cloudlet has none to give (nil cloudlet/cloudlet.ResTagMap)
		spec, vmerr = apis.resTagTableApi.GetVMSpec(ctx, stm, testflavor, nulCL, cli)
		require.Equal(t, "Cloudlet San Jose Site doesn't support GPU", vmerr.Error(), "nil table")
		return nil
	})
}

func testShowFlavorsForCloudlet(t *testing.T, ctx context.Context, apis *AllApis) {
	insertCloudletInfo(ctx, apis, testutil.CloudletInfoData)
	// Use a clouldet with no ResourceTagMap
	cCldApi := testutil.NewInternalCloudletApi(apis.cloudletApi)
	cld := testutil.CloudletData()[1]

	show := testutil.ShowFlavorsForCloudlet{}
	show.Init()

	err := cCldApi.ShowFlavorsForCloudlet(ctx, &cld.Key, &show)
	require.Nil(t, err)
	require.Equal(t, 2, len(show.Data))

	// Show flavors for a chosen operator.
	show.Init()
	cld.Key.Name = ""

	err = cCldApi.ShowFlavorsForCloudlet(ctx, &cld.Key, &show)
	require.Nil(t, err)
	require.Equal(t, 5, len(show.Data))

	// Show flavors for a chosen cloudlet name.
	show.Init()
	cld = testutil.CloudletData()[1]
	cld.Key.Organization = ""

	err = cCldApi.ShowFlavorsForCloudlet(ctx, &cld.Key, &show)
	require.Nil(t, err)
	require.Equal(t, 2, len(show.Data))
}

func testAllianceOrgs(t *testing.T, ctx context.Context, apis *AllApis) {
	data := testutil.CloudletData()
	cloudlet := data[0]

	// negative tests
	selfOrgErr := `Cannot add cloudlet's own org "UFGT Inc." as alliance org`
	dupOrgErr := `Duplicate alliance org "foo" specified`

	// update cloudlet checks
	cloudlet.AllianceOrgs = []string{cloudlet.Key.Organization}
	err := apis.cloudletApi.UpdateCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.NotNil(t, err)
	require.Equal(t, selfOrgErr, err.Error())
	cloudlet.AllianceOrgs = []string{"foo", "bar", "foo"}
	err = apis.cloudletApi.UpdateCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.NotNil(t, err)
	require.Equal(t, dupOrgErr, err.Error())

	// create cloudlet checks
	cloudlet.Key.Name += "allianceorgtest"
	cloudlet.AllianceOrgs = []string{cloudlet.Key.Organization}
	err = apis.cloudletApi.CreateCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))

	require.NotNil(t, err)
	require.Equal(t, selfOrgErr, err.Error())
	cloudlet.AllianceOrgs = []string{"foo", "bar", "foo"}
	err = apis.cloudletApi.CreateCloudlet(&cloudlet, testutil.NewCudStreamoutCloudlet(ctx))
	require.NotNil(t, err)
	require.Equal(t, dupOrgErr, err.Error())

	// add alliance org checks
	cao := edgeproto.CloudletAllianceOrg{
		Key:          data[0].Key,
		Organization: data[0].Key.Organization,
	}
	_, err = apis.cloudletApi.AddCloudletAllianceOrg(ctx, &cao)
	require.NotNil(t, err)
	require.Equal(t, selfOrgErr, err.Error())
	cao.Organization = "foo"
	_, err = apis.cloudletApi.AddCloudletAllianceOrg(ctx, &cao)
	require.Nil(t, err)
	_, err = apis.cloudletApi.AddCloudletAllianceOrg(ctx, &cao)
	require.NotNil(t, err)
	require.Equal(t, dupOrgErr, err.Error())
	_, err = apis.cloudletApi.RemoveCloudletAllianceOrg(ctx, &cao)
	require.Nil(t, err)
	_, err = apis.cloudletApi.RemoveCloudletAllianceOrg(ctx, &cao)
	require.Nil(t, err)
	// verify removed
	check := edgeproto.Cloudlet{}
	found := apis.cloudletApi.cache.Get(&data[0].Key, &check)
	require.True(t, found)
	require.Equal(t, 0, len(check.AllianceOrgs))
}

func TestShowCloudletsAppDeploy(t *testing.T) {
	log.SetDebugLevel(log.DebugLevelEtcd | log.DebugLevelApi)
	log.InitTracer(nil)
	defer log.FinishTracer()
	ctx := log.StartTestSpan(context.Background())

	testSvcs := testinit(ctx, t)
	defer testfinish(testSvcs)

	dummy := dummyEtcd{}
	dummy.Start()

	sync := InitSync(&dummy)
	apis := NewAllApis(sync)
	sync.Start()
	defer sync.Done()

	cAppApi := testutil.NewInternalAppApi(apis.appApi)

	show := testutil.ShowCloudletsForAppDeployment{}
	show.Init()

	app := testutil.AppData[2]
	request := edgeproto.DeploymentCloudletRequest{
		App:          &app,
		DryRunDeploy: false,
	}
	app.DefaultFlavor = testutil.FlavorData[0].Key // x1.tiny
	app.Deployment = cloudcommon.DeploymentTypeVM
	filter := request

	// test data
	testutil.InternalFlavorCreate(t, apis.flavorApi, testutil.FlavorData)
	testutil.InternalGPUDriverCreate(t, apis.gpuDriverApi, testutil.GPUDriverData)
	testutil.InternalResTagTableCreate(t, apis.resTagTableApi, testutil.ResTagTableData)
	testutil.InternalCloudletCreate(t, apis.cloudletApi, testutil.CloudletData())
	insertCloudletInfo(ctx, apis, testutil.CloudletInfoData)

	// without a responder, clusterInst create waits forever
	dummyResponder := DummyInfoResponder{
		AppInstCache:        &apis.appInstApi.cache,
		ClusterInstCache:    &apis.clusterInstApi.cache,
		RecvAppInstInfo:     apis.appInstInfoApi,
		RecvClusterInstInfo: apis.clusterInstInfoApi,
	}
	dummyResponder.InitDummyInfoResponder()

	reduceInfoTimeouts(t, ctx, apis)

	// either create the policy expected by one of all cloudlets, or remove that bit of config, or
	// just don't create that specific cloudlet. #1 create the policy.
	testutil.InternalAutoProvPolicyCreate(t, apis.autoProvPolicyApi, testutil.AutoProvPolicyData)
	testutil.InternalAutoScalePolicyCreate(t, apis.autoScalePolicyApi, testutil.AutoScalePolicyData)

	for _, obj := range testutil.ClusterInstData {
		err := apis.clusterInstApi.CreateClusterInst(&obj, testutil.NewCudStreamoutClusterInst(ctx))
		require.Nil(t, err, "Create ClusterInst")
	}

	err := cAppApi.ShowCloudletsForAppDeployment(ctx, &filter, &show)
	require.Nil(t, err, "ShowCloudletsForAppDeployment")
	require.Equal(t, 4, len(show.Data), "SHowCloudletsForAppDeployment")

	for k, v := range show.Data {
		fmt.Printf("\t next k: %s v: %+v flavor %s \n", k, v, filter.App.DefaultFlavor)
	}
	show.Init()
	// increase the flavor size, and expect fewer cloudlet matches
	// TODO: create sets of OS flavors to attach to our CloudletInfo objs  as substitues for whats there in test_data.go
	// for more complex matching.
	app.DefaultFlavor = testutil.FlavorData[2].Key // 3 = x1.large 4 = x1.tiny.gpu 2 = x1.medium
	err = cAppApi.ShowCloudletsForAppDeployment(ctx, &filter, &show)
	require.Nil(t, err, "ShowCloudletsForAppDeployment")
	require.Equal(t, 3, len(show.Data), "SHowCloudletsForAppDeployment")

	show.Init()
	app.DefaultFlavor = testutil.FlavorData[3].Key // 3 = x1.large 4 = x1.tiny.gpu 2 = x1.medium
	err = cAppApi.ShowCloudletsForAppDeployment(ctx, &filter, &show)
	require.Nil(t, err, "ShowCloudletsForAppDeployment")
	require.Equal(t, 1, len(show.Data), "SHowCloudletsForAppDeployment")
	show.Init()

	filter.DryRunDeploy = true

	err = cAppApi.ShowCloudletsForAppDeployment(ctx, &filter, &show)
	require.Nil(t, err, "ShowCloudletsForAppDeployment")
	require.Equal(t, 1, len(show.Data), "ShowCloudletsForAppDeployment DryRun=True")
	// TODO: Increase cloudlets refs such that San Jose can no longer support the App deployment
	dummy.Stop()
}

func testCloudletDnsLabel(t *testing.T, ctx context.Context, apis *AllApis) {
	var err error

	data := testutil.CloudletData()
	// Check that dns segment ids are unique for cloudlets.
	cl0 := data[0]
	cl0.Key.Name = "abc"
	cl0.Key.Organization = "def"
	cl0.ResTagMap = nil
	cl0.GpuConfig = edgeproto.GPUConfig{}

	cl1 := data[1]
	cl1.Key.Name = "ab,c"
	cl1.Key.Organization = "d,ef"
	cl1.ResTagMap = nil
	cl1.GpuConfig = edgeproto.GPUConfig{}

	dnsLabel0 := "abc-def"
	dnsLabel1 := "abc-def1"

	err = apis.cloudletApi.CreateCloudlet(&cl0, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err)
	err = apis.cloudletApi.CreateCloudlet(&cl1, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err)

	check0 := edgeproto.Cloudlet{}
	require.True(t, apis.cloudletApi.cache.Get(&cl0.Key, &check0))
	require.Equal(t, dnsLabel0, check0.DnsLabel)

	check1 := edgeproto.Cloudlet{}
	require.True(t, apis.cloudletApi.cache.Get(&cl1.Key, &check1))
	require.Equal(t, dnsLabel1, check1.DnsLabel)

	require.NotEqual(t, dnsLabel0, dnsLabel1)
	// check that ids are present in database
	require.True(t, testHasCloudletDnsLabel(apis.cloudletApi.sync.store, dnsLabel0))
	require.True(t, testHasCloudletDnsLabel(apis.cloudletApi.sync.store, dnsLabel1))

	// clean up
	err = apis.cloudletApi.DeleteCloudlet(&cl0, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err)
	err = apis.cloudletApi.DeleteCloudlet(&cl1, testutil.NewCudStreamoutCloudlet(ctx))
	require.Nil(t, err)
	// check that ids are removed from database
	require.False(t, testHasCloudletDnsLabel(apis.cloudletApi.sync.store, dnsLabel0))
	require.False(t, testHasCloudletDnsLabel(apis.cloudletApi.sync.store, dnsLabel1))
}

func testHasCloudletDnsLabel(kvstore objstore.KVStore, id string) bool {
	return testKVStoreHasKey(kvstore, edgeproto.CloudletDnsLabelDbKey(id))
}
