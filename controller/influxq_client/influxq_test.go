package influxq_test

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strconv"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/mobiledgex/edge-cloud/cloudcommon"
	influxq "github.com/mobiledgex/edge-cloud/controller/influxq_client"
	"github.com/mobiledgex/edge-cloud/edgeproto"
	"github.com/mobiledgex/edge-cloud/integration/process"
	"github.com/mobiledgex/edge-cloud/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestInfluxQ(t *testing.T) {
	log.SetDebugLevel(log.DebugLevelMetrics)
	log.InitTracer("")
	defer log.FinishTracer()

	addr := "http://127.0.0.1:8086"
	// lower the interval so we don't have to wait so long
	influxq.InfluxQPushInterval = 10 * time.Millisecond
	influxq.InfluxQReconnectDelay = 10 * time.Millisecond
	ctx := log.StartTestSpan(context.Background())

	// start influxd if not already running
	_, err := exec.Command("sh", "-c", "pgrep -x influxd").Output()
	if err != nil {
		p := process.Influx{}
		p.Common.Name = "influx-test"
		p.HttpAddr = addr
		p.DataDir = "/var/tmp/.influxdb"
		// start influx
		err = p.StartLocal("/var/tmp/influxdb.log",
			process.WithCleanStartup())
		require.Nil(t, err, "start InfluxDB server")
		defer p.StopLocal()
	}

	q := influxq.NewInfluxQ("metrics", "", "")
	err = q.Start(addr)
	require.Nil(t, err, "new influx q")
	defer q.Stop()

	connected := q.WaitConnected()
	assert.True(t, connected, "connected")

	// clear test metrics
	_, err = q.QueryDB(`DROP SERIES FROM "test-metric"`)
	require.Nil(t, err, "clear test metrics")

	count := 0
	iilimit := 10

	for tt := 0; tt < 10; tt++ {
		ts, _ := types.TimestampProto(time.Now())
		for ii := 0; ii < iilimit; ii++ {
			metric := edgeproto.Metric{}
			metric.Name = "test-metric"
			metric.Timestamp = *ts
			metric.Tags = make([]*edgeproto.MetricTag, 0)
			metric.Tags = append(metric.Tags, &edgeproto.MetricTag{
				Name: "tag1",
				Val:  "someval" + strconv.Itoa(ii),
			})
			metric.Tags = append(metric.Tags, &edgeproto.MetricTag{
				Name: "tag2",
				Val:  "someval",
			})
			metric.Vals = make([]*edgeproto.MetricVal, 0)
			metric.Vals = append(metric.Vals, &edgeproto.MetricVal{
				Name: "val1",
				Value: &edgeproto.MetricVal_Ival{
					Ival: uint64(ii + tt*iilimit),
				},
			})
			metric.Vals = append(metric.Vals, &edgeproto.MetricVal{
				Name: "val2",
				Value: &edgeproto.MetricVal_Dval{
					Dval: float64(ii+tt*iilimit) / 2.0,
				},
			})
			q.AddMetric(&metric)
			time.Sleep(time.Microsecond)
			count++
		}
	}

	// wait for metrics to get pushed to db
	time.Sleep(2 * influxq.InfluxQPushInterval)
	assert.Equal(t, uint64(0), q.ErrBatch, "batch errors")
	assert.Equal(t, uint64(0), q.ErrPoint, "point errors")
	assert.Equal(t, uint64(0), q.Qfull, "Qfulls")

	// wait for records to get updated in database and become queryable.
	num := 0
	for ii := 0; ii < 10; ii++ {
		res, err := q.QueryDB("select count(val1) from \"test-metric\"")
		if err == nil && len(res) > 0 && len(res[0].Series) > 0 && len(res[0].Series[0].Values) > 0 {
			jnum, ok := res[0].Series[0].Values[0][1].(json.Number)
			if ok {
				val, err := jnum.Int64()
				if err == nil && int(val) == count {
					num = count
					break
				}
			}

		}
		time.Sleep(10 * time.Millisecond)
	}
	assert.Equal(t, count, num, "num unique values")

	// query all test-metrics
	res, err := q.QueryDB("select * from \"test-metric\"")
	assert.Nil(t, err, "select *")
	assert.Equal(t, 1, len(res), "num results")
	if len(res) > 0 {
		assert.Equal(t, 1, len(res[0].Series), "num series")
		if len(res[0].Series) > 0 {
			assert.Equal(t, count, len(res[0].Series[0].Values), "num values")
			// prints results if needed
			if false {
				for ii, _ := range res[0].Series[0].Values {
					fmt.Printf("%d: %v\n", ii, res[0].Series[0].Values[ii])
				}
			}
		}
	}
	testAutoProvCounts(t, ctx, q)
}

// Test pushing auto prov counts to influxdb and reading back.
func testAutoProvCounts(t *testing.T, ctx context.Context, q *influxq.InfluxQ) {
	_, err := q.QueryDB(fmt.Sprintf(`DROP SERIES FROM "%s"`, cloudcommon.AutoProvMeasurement))
	require.Nil(t, err, "clear test metrics")

	ap := edgeproto.AutoProvCount{}
	ap.AppKey.DeveloperKey.Name = "dev1"
	ap.AppKey.Name = "app1"
	ap.AppKey.Version = "1.0.0"
	ap.CloudletKey.OperatorKey.Name = "oper1"
	ap.CloudletKey.Name = "cloudlet1"
	ap.Count = 42

	msg := edgeproto.AutoProvCounts{}
	msg.DmeNodeName = "dmeid"
	tsp, err := types.TimestampProto(time.Now())
	require.Nil(t, err, "timestamp proto")
	msg.Timestamp = *tsp
	msg.Counts = []*edgeproto.AutoProvCount{&ap}

	ts, err := types.TimestampFromProto(&msg.Timestamp)
	require.Nil(t, err, "timestamp from proto")

	err = q.PushAutoProvCounts(ctx, &msg)
	require.Nil(t, err, "push auto prov counts")

	res, err := q.QueryDB(fmt.Sprintf(`select * from "%s"`, cloudcommon.AutoProvMeasurement))
	require.Nil(t, err, "select %s", cloudcommon.AutoProvMeasurement)
	require.Equal(t, 1, len(res), "num results")
	require.Equal(t, 1, len(res[0].Series))
	row := res[0].Series[0]
	require.Equal(t, 1, len(row.Values))
	apCheck, dmeid, tsCheck, err := influxq.ParseAutoProvCount(row.Columns, row.Values[0])
	require.Nil(t, err, "parse auto prov counts")
	require.Equal(t, msg.DmeNodeName, dmeid, "check dmeid")
	require.Equal(t, ts, tsCheck, "check timestmap")
	require.Equal(t, ap, *apCheck, "check auto prov stats")
}