// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: refs.proto

package testutil

import (
	"context"
	fmt "fmt"
	"github.com/edgexr/edge-cloud/edgectl/wrapper"
	"github.com/edgexr/edge-cloud/edgeproto"
	"github.com/edgexr/edge-cloud/log"
	_ "github.com/edgexr/edge-cloud/protogen"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"io"
	math "math"
	"testing"
	"time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Auto-generated code: DO NOT EDIT

type ShowCloudletRefs struct {
	Data map[string]edgeproto.CloudletRefs
	grpc.ServerStream
	Ctx context.Context
}

func (x *ShowCloudletRefs) Init() {
	x.Data = make(map[string]edgeproto.CloudletRefs)
}

func (x *ShowCloudletRefs) Send(m *edgeproto.CloudletRefs) error {
	x.Data[m.GetKey().GetKeyString()] = *m
	return nil
}

func (x *ShowCloudletRefs) Context() context.Context {
	return x.Ctx
}

var CloudletRefsShowExtraCount = 0

func (x *ShowCloudletRefs) ReadStream(stream edgeproto.CloudletRefsApi_ShowCloudletRefsClient, err error) {
	x.Data = make(map[string]edgeproto.CloudletRefs)
	if err != nil {
		return
	}
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		x.Data[obj.GetKey().GetKeyString()] = *obj
	}
}

func (x *ShowCloudletRefs) CheckFound(obj *edgeproto.CloudletRefs) bool {
	_, found := x.Data[obj.GetKey().GetKeyString()]
	return found
}

func (x *ShowCloudletRefs) AssertFound(t *testing.T, obj *edgeproto.CloudletRefs) {
	check, found := x.Data[obj.GetKey().GetKeyString()]
	require.True(t, found, "find CloudletRefs %s", obj.GetKey().GetKeyString())
	if found && !check.Matches(obj, edgeproto.MatchIgnoreBackend(), edgeproto.MatchSortArrayedKeys()) {
		require.Equal(t, *obj, check, "CloudletRefs are equal")
	}
	if found {
		// remove in case there are dups in the list, so the
		// same object cannot be used again
		delete(x.Data, obj.GetKey().GetKeyString())
	}
}

func (x *ShowCloudletRefs) AssertNotFound(t *testing.T, obj *edgeproto.CloudletRefs) {
	_, found := x.Data[obj.GetKey().GetKeyString()]
	require.False(t, found, "do not find CloudletRefs %s", obj.GetKey().GetKeyString())
}

func WaitAssertFoundCloudletRefs(t *testing.T, api edgeproto.CloudletRefsApiClient, obj *edgeproto.CloudletRefs, count int, retry time.Duration) {
	show := ShowCloudletRefs{}
	for ii := 0; ii < count; ii++ {
		ctx, cancel := context.WithTimeout(context.Background(), retry)
		stream, err := api.ShowCloudletRefs(ctx, obj)
		show.ReadStream(stream, err)
		cancel()
		if show.CheckFound(obj) {
			break
		}
		time.Sleep(retry)
	}
	show.AssertFound(t, obj)
}

func WaitAssertNotFoundCloudletRefs(t *testing.T, api edgeproto.CloudletRefsApiClient, obj *edgeproto.CloudletRefs, count int, retry time.Duration) {
	show := ShowCloudletRefs{}
	filterNone := edgeproto.CloudletRefs{}
	for ii := 0; ii < count; ii++ {
		ctx, cancel := context.WithTimeout(context.Background(), retry)
		stream, err := api.ShowCloudletRefs(ctx, &filterNone)
		show.ReadStream(stream, err)
		cancel()
		if !show.CheckFound(obj) {
			break
		}
		time.Sleep(retry)
	}
	show.AssertNotFound(t, obj)
}

// Wrap the api with a common interface
type CloudletRefsCommonApi struct {
	internal_api edgeproto.CloudletRefsApiServer
	client_api   edgeproto.CloudletRefsApiClient
}

func (x *CloudletRefsCommonApi) ShowCloudletRefs(ctx context.Context, filter *edgeproto.CloudletRefs, showData *ShowCloudletRefs) error {
	if x.internal_api != nil {
		showData.Ctx = ctx
		return x.internal_api.ShowCloudletRefs(filter, showData)
	} else {
		stream, err := x.client_api.ShowCloudletRefs(ctx, filter)
		showData.ReadStream(stream, err)
		return err
	}
}

func NewInternalCloudletRefsApi(api edgeproto.CloudletRefsApiServer) *CloudletRefsCommonApi {
	apiWrap := CloudletRefsCommonApi{}
	apiWrap.internal_api = api
	return &apiWrap
}

func NewClientCloudletRefsApi(api edgeproto.CloudletRefsApiClient) *CloudletRefsCommonApi {
	apiWrap := CloudletRefsCommonApi{}
	apiWrap.client_api = api
	return &apiWrap
}

type CloudletRefsTestOptions struct {
	createdData []edgeproto.CloudletRefs
}

type CloudletRefsTestOp func(opts *CloudletRefsTestOptions)

func WithCreatedCloudletRefsTestData(createdData []edgeproto.CloudletRefs) CloudletRefsTestOp {
	return func(opts *CloudletRefsTestOptions) { opts.createdData = createdData }
}

func InternalCloudletRefsTest(t *testing.T, test string, api edgeproto.CloudletRefsApiServer, testData []edgeproto.CloudletRefs, ops ...CloudletRefsTestOp) {
	span := log.StartSpan(log.DebugLevelApi, "InternalCloudletRefsTest")
	defer span.Finish()
	ctx := log.ContextWithSpan(context.Background(), span)

	switch test {
	case "show":
		basicCloudletRefsShowTest(t, ctx, NewInternalCloudletRefsApi(api), testData)
	}
}

func ClientCloudletRefsTest(t *testing.T, test string, api edgeproto.CloudletRefsApiClient, testData []edgeproto.CloudletRefs, ops ...CloudletRefsTestOp) {
	span := log.StartSpan(log.DebugLevelApi, "ClientCloudletRefsTest")
	defer span.Finish()
	ctx := log.ContextWithSpan(context.Background(), span)

	switch test {
	case "show":
		basicCloudletRefsShowTest(t, ctx, NewClientCloudletRefsApi(api), testData)
	}
}

func basicCloudletRefsShowTest(t *testing.T, ctx context.Context, api *CloudletRefsCommonApi, testData []edgeproto.CloudletRefs) {
	var err error

	show := ShowCloudletRefs{}
	show.Init()
	filterNone := edgeproto.CloudletRefs{}
	err = api.ShowCloudletRefs(ctx, &filterNone, &show)
	require.Nil(t, err, "show data")
	require.Equal(t, len(testData)+CloudletRefsShowExtraCount, len(show.Data), "Show count")
	for _, obj := range testData {
		show.AssertFound(t, &obj)
	}
}

func GetCloudletRefs(t *testing.T, ctx context.Context, api *CloudletRefsCommonApi, key *edgeproto.CloudletKey, out *edgeproto.CloudletRefs) bool {
	var err error

	show := ShowCloudletRefs{}
	show.Init()
	filter := edgeproto.CloudletRefs{}
	filter.SetKey(key)
	err = api.ShowCloudletRefs(ctx, &filter, &show)
	require.Nil(t, err, "show data")
	obj, found := show.Data[key.GetKeyString()]
	if found {
		*out = obj
	}
	return found
}

func FindCloudletRefsData(key *edgeproto.CloudletKey, testData []edgeproto.CloudletRefs) (*edgeproto.CloudletRefs, bool) {
	for ii, _ := range testData {
		if testData[ii].GetKey().Matches(key) {
			return &testData[ii], true
		}
	}
	return nil, false
}

type ShowClusterRefs struct {
	Data map[string]edgeproto.ClusterRefs
	grpc.ServerStream
	Ctx context.Context
}

func (x *ShowClusterRefs) Init() {
	x.Data = make(map[string]edgeproto.ClusterRefs)
}

func (x *ShowClusterRefs) Send(m *edgeproto.ClusterRefs) error {
	x.Data[m.GetKey().GetKeyString()] = *m
	return nil
}

func (x *ShowClusterRefs) Context() context.Context {
	return x.Ctx
}

var ClusterRefsShowExtraCount = 0

func (x *ShowClusterRefs) ReadStream(stream edgeproto.ClusterRefsApi_ShowClusterRefsClient, err error) {
	x.Data = make(map[string]edgeproto.ClusterRefs)
	if err != nil {
		return
	}
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		x.Data[obj.GetKey().GetKeyString()] = *obj
	}
}

func (x *ShowClusterRefs) CheckFound(obj *edgeproto.ClusterRefs) bool {
	_, found := x.Data[obj.GetKey().GetKeyString()]
	return found
}

func (x *ShowClusterRefs) AssertFound(t *testing.T, obj *edgeproto.ClusterRefs) {
	check, found := x.Data[obj.GetKey().GetKeyString()]
	require.True(t, found, "find ClusterRefs %s", obj.GetKey().GetKeyString())
	if found && !check.Matches(obj, edgeproto.MatchIgnoreBackend(), edgeproto.MatchSortArrayedKeys()) {
		require.Equal(t, *obj, check, "ClusterRefs are equal")
	}
	if found {
		// remove in case there are dups in the list, so the
		// same object cannot be used again
		delete(x.Data, obj.GetKey().GetKeyString())
	}
}

func (x *ShowClusterRefs) AssertNotFound(t *testing.T, obj *edgeproto.ClusterRefs) {
	_, found := x.Data[obj.GetKey().GetKeyString()]
	require.False(t, found, "do not find ClusterRefs %s", obj.GetKey().GetKeyString())
}

func WaitAssertFoundClusterRefs(t *testing.T, api edgeproto.ClusterRefsApiClient, obj *edgeproto.ClusterRefs, count int, retry time.Duration) {
	show := ShowClusterRefs{}
	for ii := 0; ii < count; ii++ {
		ctx, cancel := context.WithTimeout(context.Background(), retry)
		stream, err := api.ShowClusterRefs(ctx, obj)
		show.ReadStream(stream, err)
		cancel()
		if show.CheckFound(obj) {
			break
		}
		time.Sleep(retry)
	}
	show.AssertFound(t, obj)
}

func WaitAssertNotFoundClusterRefs(t *testing.T, api edgeproto.ClusterRefsApiClient, obj *edgeproto.ClusterRefs, count int, retry time.Duration) {
	show := ShowClusterRefs{}
	filterNone := edgeproto.ClusterRefs{}
	for ii := 0; ii < count; ii++ {
		ctx, cancel := context.WithTimeout(context.Background(), retry)
		stream, err := api.ShowClusterRefs(ctx, &filterNone)
		show.ReadStream(stream, err)
		cancel()
		if !show.CheckFound(obj) {
			break
		}
		time.Sleep(retry)
	}
	show.AssertNotFound(t, obj)
}

// Wrap the api with a common interface
type ClusterRefsCommonApi struct {
	internal_api edgeproto.ClusterRefsApiServer
	client_api   edgeproto.ClusterRefsApiClient
}

func (x *ClusterRefsCommonApi) ShowClusterRefs(ctx context.Context, filter *edgeproto.ClusterRefs, showData *ShowClusterRefs) error {
	if x.internal_api != nil {
		showData.Ctx = ctx
		return x.internal_api.ShowClusterRefs(filter, showData)
	} else {
		stream, err := x.client_api.ShowClusterRefs(ctx, filter)
		showData.ReadStream(stream, err)
		return err
	}
}

func NewInternalClusterRefsApi(api edgeproto.ClusterRefsApiServer) *ClusterRefsCommonApi {
	apiWrap := ClusterRefsCommonApi{}
	apiWrap.internal_api = api
	return &apiWrap
}

func NewClientClusterRefsApi(api edgeproto.ClusterRefsApiClient) *ClusterRefsCommonApi {
	apiWrap := ClusterRefsCommonApi{}
	apiWrap.client_api = api
	return &apiWrap
}

type ClusterRefsTestOptions struct {
	createdData []edgeproto.ClusterRefs
}

type ClusterRefsTestOp func(opts *ClusterRefsTestOptions)

func WithCreatedClusterRefsTestData(createdData []edgeproto.ClusterRefs) ClusterRefsTestOp {
	return func(opts *ClusterRefsTestOptions) { opts.createdData = createdData }
}

func InternalClusterRefsTest(t *testing.T, test string, api edgeproto.ClusterRefsApiServer, testData []edgeproto.ClusterRefs, ops ...ClusterRefsTestOp) {
	span := log.StartSpan(log.DebugLevelApi, "InternalClusterRefsTest")
	defer span.Finish()
	ctx := log.ContextWithSpan(context.Background(), span)

	switch test {
	case "show":
		basicClusterRefsShowTest(t, ctx, NewInternalClusterRefsApi(api), testData)
	}
}

func ClientClusterRefsTest(t *testing.T, test string, api edgeproto.ClusterRefsApiClient, testData []edgeproto.ClusterRefs, ops ...ClusterRefsTestOp) {
	span := log.StartSpan(log.DebugLevelApi, "ClientClusterRefsTest")
	defer span.Finish()
	ctx := log.ContextWithSpan(context.Background(), span)

	switch test {
	case "show":
		basicClusterRefsShowTest(t, ctx, NewClientClusterRefsApi(api), testData)
	}
}

func basicClusterRefsShowTest(t *testing.T, ctx context.Context, api *ClusterRefsCommonApi, testData []edgeproto.ClusterRefs) {
	var err error

	show := ShowClusterRefs{}
	show.Init()
	filterNone := edgeproto.ClusterRefs{}
	err = api.ShowClusterRefs(ctx, &filterNone, &show)
	require.Nil(t, err, "show data")
	require.Equal(t, len(testData)+ClusterRefsShowExtraCount, len(show.Data), "Show count")
	for _, obj := range testData {
		show.AssertFound(t, &obj)
	}
}

func GetClusterRefs(t *testing.T, ctx context.Context, api *ClusterRefsCommonApi, key *edgeproto.ClusterInstKey, out *edgeproto.ClusterRefs) bool {
	var err error

	show := ShowClusterRefs{}
	show.Init()
	filter := edgeproto.ClusterRefs{}
	filter.SetKey(key)
	err = api.ShowClusterRefs(ctx, &filter, &show)
	require.Nil(t, err, "show data")
	obj, found := show.Data[key.GetKeyString()]
	if found {
		*out = obj
	}
	return found
}

func FindClusterRefsData(key *edgeproto.ClusterInstKey, testData []edgeproto.ClusterRefs) (*edgeproto.ClusterRefs, bool) {
	for ii, _ := range testData {
		if testData[ii].GetKey().Matches(key) {
			return &testData[ii], true
		}
	}
	return nil, false
}

type ShowAppInstRefs struct {
	Data map[string]edgeproto.AppInstRefs
	grpc.ServerStream
	Ctx context.Context
}

func (x *ShowAppInstRefs) Init() {
	x.Data = make(map[string]edgeproto.AppInstRefs)
}

func (x *ShowAppInstRefs) Send(m *edgeproto.AppInstRefs) error {
	x.Data[m.GetKey().GetKeyString()] = *m
	return nil
}

func (x *ShowAppInstRefs) Context() context.Context {
	return x.Ctx
}

var AppInstRefsShowExtraCount = 0

func (x *ShowAppInstRefs) ReadStream(stream edgeproto.AppInstRefsApi_ShowAppInstRefsClient, err error) {
	x.Data = make(map[string]edgeproto.AppInstRefs)
	if err != nil {
		return
	}
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		x.Data[obj.GetKey().GetKeyString()] = *obj
	}
}

func (x *ShowAppInstRefs) CheckFound(obj *edgeproto.AppInstRefs) bool {
	_, found := x.Data[obj.GetKey().GetKeyString()]
	return found
}

func (x *ShowAppInstRefs) AssertFound(t *testing.T, obj *edgeproto.AppInstRefs) {
	check, found := x.Data[obj.GetKey().GetKeyString()]
	require.True(t, found, "find AppInstRefs %s", obj.GetKey().GetKeyString())
	if found && !check.Matches(obj, edgeproto.MatchIgnoreBackend(), edgeproto.MatchSortArrayedKeys()) {
		require.Equal(t, *obj, check, "AppInstRefs are equal")
	}
	if found {
		// remove in case there are dups in the list, so the
		// same object cannot be used again
		delete(x.Data, obj.GetKey().GetKeyString())
	}
}

func (x *ShowAppInstRefs) AssertNotFound(t *testing.T, obj *edgeproto.AppInstRefs) {
	_, found := x.Data[obj.GetKey().GetKeyString()]
	require.False(t, found, "do not find AppInstRefs %s", obj.GetKey().GetKeyString())
}

func WaitAssertFoundAppInstRefs(t *testing.T, api edgeproto.AppInstRefsApiClient, obj *edgeproto.AppInstRefs, count int, retry time.Duration) {
	show := ShowAppInstRefs{}
	for ii := 0; ii < count; ii++ {
		ctx, cancel := context.WithTimeout(context.Background(), retry)
		stream, err := api.ShowAppInstRefs(ctx, obj)
		show.ReadStream(stream, err)
		cancel()
		if show.CheckFound(obj) {
			break
		}
		time.Sleep(retry)
	}
	show.AssertFound(t, obj)
}

func WaitAssertNotFoundAppInstRefs(t *testing.T, api edgeproto.AppInstRefsApiClient, obj *edgeproto.AppInstRefs, count int, retry time.Duration) {
	show := ShowAppInstRefs{}
	filterNone := edgeproto.AppInstRefs{}
	for ii := 0; ii < count; ii++ {
		ctx, cancel := context.WithTimeout(context.Background(), retry)
		stream, err := api.ShowAppInstRefs(ctx, &filterNone)
		show.ReadStream(stream, err)
		cancel()
		if !show.CheckFound(obj) {
			break
		}
		time.Sleep(retry)
	}
	show.AssertNotFound(t, obj)
}

// Wrap the api with a common interface
type AppInstRefsCommonApi struct {
	internal_api edgeproto.AppInstRefsApiServer
	client_api   edgeproto.AppInstRefsApiClient
}

func (x *AppInstRefsCommonApi) ShowAppInstRefs(ctx context.Context, filter *edgeproto.AppInstRefs, showData *ShowAppInstRefs) error {
	if x.internal_api != nil {
		showData.Ctx = ctx
		return x.internal_api.ShowAppInstRefs(filter, showData)
	} else {
		stream, err := x.client_api.ShowAppInstRefs(ctx, filter)
		showData.ReadStream(stream, err)
		return err
	}
}

func NewInternalAppInstRefsApi(api edgeproto.AppInstRefsApiServer) *AppInstRefsCommonApi {
	apiWrap := AppInstRefsCommonApi{}
	apiWrap.internal_api = api
	return &apiWrap
}

func NewClientAppInstRefsApi(api edgeproto.AppInstRefsApiClient) *AppInstRefsCommonApi {
	apiWrap := AppInstRefsCommonApi{}
	apiWrap.client_api = api
	return &apiWrap
}

type AppInstRefsTestOptions struct {
	createdData []edgeproto.AppInstRefs
}

type AppInstRefsTestOp func(opts *AppInstRefsTestOptions)

func WithCreatedAppInstRefsTestData(createdData []edgeproto.AppInstRefs) AppInstRefsTestOp {
	return func(opts *AppInstRefsTestOptions) { opts.createdData = createdData }
}

func InternalAppInstRefsTest(t *testing.T, test string, api edgeproto.AppInstRefsApiServer, testData []edgeproto.AppInstRefs, ops ...AppInstRefsTestOp) {
	span := log.StartSpan(log.DebugLevelApi, "InternalAppInstRefsTest")
	defer span.Finish()
	ctx := log.ContextWithSpan(context.Background(), span)

	switch test {
	case "show":
		basicAppInstRefsShowTest(t, ctx, NewInternalAppInstRefsApi(api), testData)
	}
}

func ClientAppInstRefsTest(t *testing.T, test string, api edgeproto.AppInstRefsApiClient, testData []edgeproto.AppInstRefs, ops ...AppInstRefsTestOp) {
	span := log.StartSpan(log.DebugLevelApi, "ClientAppInstRefsTest")
	defer span.Finish()
	ctx := log.ContextWithSpan(context.Background(), span)

	switch test {
	case "show":
		basicAppInstRefsShowTest(t, ctx, NewClientAppInstRefsApi(api), testData)
	}
}

func basicAppInstRefsShowTest(t *testing.T, ctx context.Context, api *AppInstRefsCommonApi, testData []edgeproto.AppInstRefs) {
	var err error

	show := ShowAppInstRefs{}
	show.Init()
	filterNone := edgeproto.AppInstRefs{}
	err = api.ShowAppInstRefs(ctx, &filterNone, &show)
	require.Nil(t, err, "show data")
	require.Equal(t, len(testData)+AppInstRefsShowExtraCount, len(show.Data), "Show count")
	for _, obj := range testData {
		show.AssertFound(t, &obj)
	}
}

func GetAppInstRefs(t *testing.T, ctx context.Context, api *AppInstRefsCommonApi, key *edgeproto.AppKey, out *edgeproto.AppInstRefs) bool {
	var err error

	show := ShowAppInstRefs{}
	show.Init()
	filter := edgeproto.AppInstRefs{}
	filter.SetKey(key)
	err = api.ShowAppInstRefs(ctx, &filter, &show)
	require.Nil(t, err, "show data")
	obj, found := show.Data[key.GetKeyString()]
	if found {
		*out = obj
	}
	return found
}

func FindAppInstRefsData(key *edgeproto.AppKey, testData []edgeproto.AppInstRefs) (*edgeproto.AppInstRefs, bool) {
	for ii, _ := range testData {
		if testData[ii].GetKey().Matches(key) {
			return &testData[ii], true
		}
	}
	return nil, false
}

func (r *Run) CloudletRefsApi(data *[]edgeproto.CloudletRefs, dataMap interface{}, dataOut interface{}) {
	log.DebugLog(log.DebugLevelApi, "API for CloudletRefs", "mode", r.Mode)
	if r.Mode == "show" {
		obj := &edgeproto.CloudletRefs{}
		out, err := r.client.ShowCloudletRefs(r.ctx, obj)
		if err != nil {
			r.logErr("CloudletRefsApi", err)
		} else {
			outp, ok := dataOut.(*[]edgeproto.CloudletRefs)
			if !ok {
				panic(fmt.Sprintf("RunCloudletRefsApi expected dataOut type *[]edgeproto.CloudletRefs, but was %T", dataOut))
			}
			*outp = append(*outp, out...)
		}
		return
	}
	for ii, objD := range *data {
		obj := &objD
		switch r.Mode {
		case "showfiltered":
			out, err := r.client.ShowCloudletRefs(r.ctx, obj)
			if err != nil {
				r.logErr(fmt.Sprintf("CloudletRefsApi[%d]", ii), err)
			} else {
				outp, ok := dataOut.(*[]edgeproto.CloudletRefs)
				if !ok {
					panic(fmt.Sprintf("RunCloudletRefsApi expected dataOut type *[]edgeproto.CloudletRefs, but was %T", dataOut))
				}
				*outp = append(*outp, out...)
			}
		}
	}
}

func (s *DummyServer) ShowCloudletRefs(in *edgeproto.CloudletRefs, server edgeproto.CloudletRefsApi_ShowCloudletRefsServer) error {
	var err error
	obj := &edgeproto.CloudletRefs{}
	if obj.Matches(in, edgeproto.MatchFilter()) {
		for ii := 0; ii < s.ShowDummyCount; ii++ {
			server.Send(&edgeproto.CloudletRefs{})
		}
		if ch, ok := s.MidstreamFailChs["ShowCloudletRefs"]; ok {
			// Wait until client receives the SendMsg, since they
			// are buffered and dropped once we return err here.
			select {
			case <-ch:
			case <-time.After(5 * time.Second):
			}
			return fmt.Errorf("midstream failure!")
		}
	}
	err = s.CloudletRefsCache.Show(in, func(obj *edgeproto.CloudletRefs) error {
		err := server.Send(obj)
		return err
	})
	return err
}

func (r *Run) ClusterRefsApi(data *[]edgeproto.ClusterRefs, dataMap interface{}, dataOut interface{}) {
	log.DebugLog(log.DebugLevelApi, "API for ClusterRefs", "mode", r.Mode)
	if r.Mode == "show" {
		obj := &edgeproto.ClusterRefs{}
		out, err := r.client.ShowClusterRefs(r.ctx, obj)
		if err != nil {
			r.logErr("ClusterRefsApi", err)
		} else {
			outp, ok := dataOut.(*[]edgeproto.ClusterRefs)
			if !ok {
				panic(fmt.Sprintf("RunClusterRefsApi expected dataOut type *[]edgeproto.ClusterRefs, but was %T", dataOut))
			}
			*outp = append(*outp, out...)
		}
		return
	}
	for ii, objD := range *data {
		obj := &objD
		switch r.Mode {
		case "showfiltered":
			out, err := r.client.ShowClusterRefs(r.ctx, obj)
			if err != nil {
				r.logErr(fmt.Sprintf("ClusterRefsApi[%d]", ii), err)
			} else {
				outp, ok := dataOut.(*[]edgeproto.ClusterRefs)
				if !ok {
					panic(fmt.Sprintf("RunClusterRefsApi expected dataOut type *[]edgeproto.ClusterRefs, but was %T", dataOut))
				}
				*outp = append(*outp, out...)
			}
		}
	}
}

func (s *DummyServer) ShowClusterRefs(in *edgeproto.ClusterRefs, server edgeproto.ClusterRefsApi_ShowClusterRefsServer) error {
	var err error
	obj := &edgeproto.ClusterRefs{}
	if obj.Matches(in, edgeproto.MatchFilter()) {
		for ii := 0; ii < s.ShowDummyCount; ii++ {
			server.Send(&edgeproto.ClusterRefs{})
		}
		if ch, ok := s.MidstreamFailChs["ShowClusterRefs"]; ok {
			// Wait until client receives the SendMsg, since they
			// are buffered and dropped once we return err here.
			select {
			case <-ch:
			case <-time.After(5 * time.Second):
			}
			return fmt.Errorf("midstream failure!")
		}
	}
	err = s.ClusterRefsCache.Show(in, func(obj *edgeproto.ClusterRefs) error {
		err := server.Send(obj)
		return err
	})
	return err
}

func (r *Run) AppInstRefsApi(data *[]edgeproto.AppInstRefs, dataMap interface{}, dataOut interface{}) {
	log.DebugLog(log.DebugLevelApi, "API for AppInstRefs", "mode", r.Mode)
	if r.Mode == "show" {
		obj := &edgeproto.AppInstRefs{}
		out, err := r.client.ShowAppInstRefs(r.ctx, obj)
		if err != nil {
			r.logErr("AppInstRefsApi", err)
		} else {
			outp, ok := dataOut.(*[]edgeproto.AppInstRefs)
			if !ok {
				panic(fmt.Sprintf("RunAppInstRefsApi expected dataOut type *[]edgeproto.AppInstRefs, but was %T", dataOut))
			}
			*outp = append(*outp, out...)
		}
		return
	}
	for ii, objD := range *data {
		obj := &objD
		switch r.Mode {
		case "showfiltered":
			out, err := r.client.ShowAppInstRefs(r.ctx, obj)
			if err != nil {
				r.logErr(fmt.Sprintf("AppInstRefsApi[%d]", ii), err)
			} else {
				outp, ok := dataOut.(*[]edgeproto.AppInstRefs)
				if !ok {
					panic(fmt.Sprintf("RunAppInstRefsApi expected dataOut type *[]edgeproto.AppInstRefs, but was %T", dataOut))
				}
				*outp = append(*outp, out...)
			}
		}
	}
}

func (s *DummyServer) ShowAppInstRefs(in *edgeproto.AppInstRefs, server edgeproto.AppInstRefsApi_ShowAppInstRefsServer) error {
	var err error
	obj := &edgeproto.AppInstRefs{}
	if obj.Matches(in, edgeproto.MatchFilter()) {
		for ii := 0; ii < s.ShowDummyCount; ii++ {
			server.Send(&edgeproto.AppInstRefs{})
		}
		if ch, ok := s.MidstreamFailChs["ShowAppInstRefs"]; ok {
			// Wait until client receives the SendMsg, since they
			// are buffered and dropped once we return err here.
			select {
			case <-ch:
			case <-time.After(5 * time.Second):
			}
			return fmt.Errorf("midstream failure!")
		}
	}
	err = s.AppInstRefsCache.Show(in, func(obj *edgeproto.AppInstRefs) error {
		err := server.Send(obj)
		return err
	})
	return err
}

type CloudletRefsStream interface {
	Recv() (*edgeproto.CloudletRefs, error)
}

func CloudletRefsReadStream(stream CloudletRefsStream) ([]edgeproto.CloudletRefs, error) {
	output := []edgeproto.CloudletRefs{}
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return output, fmt.Errorf("read CloudletRefs stream failed, %v", err)
		}
		output = append(output, *obj)
	}
	return output, nil
}

func (s *ApiClient) ShowCloudletRefs(ctx context.Context, in *edgeproto.CloudletRefs) ([]edgeproto.CloudletRefs, error) {
	api := edgeproto.NewCloudletRefsApiClient(s.Conn)
	stream, err := api.ShowCloudletRefs(ctx, in)
	if err != nil {
		return nil, err
	}
	return CloudletRefsReadStream(stream)
}

func (s *CliClient) ShowCloudletRefs(ctx context.Context, in *edgeproto.CloudletRefs) ([]edgeproto.CloudletRefs, error) {
	output := []edgeproto.CloudletRefs{}
	args := append(s.BaseArgs, "controller", "ShowCloudletRefs")
	err := wrapper.RunEdgectlObjs(args, in, &output, s.RunOps...)
	return output, err
}

type CloudletRefsApiClient interface {
	ShowCloudletRefs(ctx context.Context, in *edgeproto.CloudletRefs) ([]edgeproto.CloudletRefs, error)
}

type ClusterRefsStream interface {
	Recv() (*edgeproto.ClusterRefs, error)
}

func ClusterRefsReadStream(stream ClusterRefsStream) ([]edgeproto.ClusterRefs, error) {
	output := []edgeproto.ClusterRefs{}
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return output, fmt.Errorf("read ClusterRefs stream failed, %v", err)
		}
		output = append(output, *obj)
	}
	return output, nil
}

func (s *ApiClient) ShowClusterRefs(ctx context.Context, in *edgeproto.ClusterRefs) ([]edgeproto.ClusterRefs, error) {
	api := edgeproto.NewClusterRefsApiClient(s.Conn)
	stream, err := api.ShowClusterRefs(ctx, in)
	if err != nil {
		return nil, err
	}
	return ClusterRefsReadStream(stream)
}

func (s *CliClient) ShowClusterRefs(ctx context.Context, in *edgeproto.ClusterRefs) ([]edgeproto.ClusterRefs, error) {
	output := []edgeproto.ClusterRefs{}
	args := append(s.BaseArgs, "controller", "ShowClusterRefs")
	err := wrapper.RunEdgectlObjs(args, in, &output, s.RunOps...)
	return output, err
}

type ClusterRefsApiClient interface {
	ShowClusterRefs(ctx context.Context, in *edgeproto.ClusterRefs) ([]edgeproto.ClusterRefs, error)
}

type AppInstRefsStream interface {
	Recv() (*edgeproto.AppInstRefs, error)
}

func AppInstRefsReadStream(stream AppInstRefsStream) ([]edgeproto.AppInstRefs, error) {
	output := []edgeproto.AppInstRefs{}
	for {
		obj, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return output, fmt.Errorf("read AppInstRefs stream failed, %v", err)
		}
		output = append(output, *obj)
	}
	return output, nil
}

func (s *ApiClient) ShowAppInstRefs(ctx context.Context, in *edgeproto.AppInstRefs) ([]edgeproto.AppInstRefs, error) {
	api := edgeproto.NewAppInstRefsApiClient(s.Conn)
	stream, err := api.ShowAppInstRefs(ctx, in)
	if err != nil {
		return nil, err
	}
	return AppInstRefsReadStream(stream)
}

func (s *CliClient) ShowAppInstRefs(ctx context.Context, in *edgeproto.AppInstRefs) ([]edgeproto.AppInstRefs, error) {
	output := []edgeproto.AppInstRefs{}
	args := append(s.BaseArgs, "controller", "ShowAppInstRefs")
	err := wrapper.RunEdgectlObjs(args, in, &output, s.RunOps...)
	return output, err
}

type AppInstRefsApiClient interface {
	ShowAppInstRefs(ctx context.Context, in *edgeproto.AppInstRefs) ([]edgeproto.AppInstRefs, error)
}
