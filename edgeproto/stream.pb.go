// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stream.proto

package edgeproto

import (
	context "context"
	"encoding/json"
	fmt "fmt"
	_ "github.com/gogo/googleapis/google/api"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/mobiledgex/edge-cloud/protogen"
	"github.com/mobiledgex/edge-cloud/util"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	reflect "reflect"
	"strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Stream State
//
// Indicates if stream has started/ended or in a bad shape
//
// 0: `STREAM_UNKNOWN`
// 1: `STREAM_START`
// 2: `STREAM_STOP`
// 3: `STREAM_ERROR`
type StreamState int32

const (
	// Stream state is unknown
	StreamState_STREAM_UNKNOWN StreamState = 0
	// Stream has started
	StreamState_STREAM_START StreamState = 1
	// Stream has stopped
	StreamState_STREAM_STOP StreamState = 2
	// Stream is in error state
	StreamState_STREAM_ERROR StreamState = 3
)

var StreamState_name = map[int32]string{
	0: "STREAM_UNKNOWN",
	1: "STREAM_START",
	2: "STREAM_STOP",
	3: "STREAM_ERROR",
}

var StreamState_value = map[string]int32{
	"STREAM_UNKNOWN": 0,
	"STREAM_START":   1,
	"STREAM_STOP":    2,
	"STREAM_ERROR":   3,
}

func (x StreamState) String() string {
	return proto.EnumName(StreamState_name, int32(x))
}

func (StreamState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_bb17ef3f514bfe54, []int{0}
}

func init() {
	proto.RegisterEnum("edgeproto.StreamState", StreamState_name, StreamState_value)
}

func init() { proto.RegisterFile("stream.proto", fileDescriptor_bb17ef3f514bfe54) }

var fileDescriptor_bb17ef3f514bfe54 = []byte{
	// 502 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0x81, 0x90, 0xf0, 0xba, 0x36, 0x35, 0xdb, 0x08, 0xd5, 0x94, 0x43, 0x6e, 0x4c,
	0x25, 0xe6, 0xe5, 0x82, 0xb8, 0x05, 0x98, 0x10, 0x2a, 0x34, 0x23, 0x6d, 0xc7, 0x11, 0xa5, 0xa9,
	0x65, 0x0c, 0x69, 0x6c, 0x25, 0x0e, 0x50, 0x8e, 0x5c, 0x91, 0x10, 0x12, 0x12, 0x7c, 0x05, 0x3e,
	0xca, 0x8e, 0x93, 0xb8, 0x70, 0x1c, 0x2d, 0xa7, 0x9e, 0xcb, 0x1d, 0xd5, 0x4e, 0xba, 0x14, 0x51,
	0xed, 0x12, 0x3d, 0xfe, 0xf9, 0xef, 0x3c, 0x3f, 0xbf, 0xc0, 0x6a, 0x2a, 0x13, 0x12, 0x8c, 0x1c,
	0x91, 0x70, 0xc9, 0xd1, 0x65, 0x32, 0xa4, 0x44, 0x95, 0xcd, 0x6a, 0x42, 0xd2, 0x2c, 0x92, 0x7a,
	0xa2, 0xb9, 0x15, 0x08, 0xc1, 0xe2, 0xb4, 0x18, 0x36, 0xc2, 0x28, 0x4b, 0x25, 0x49, 0x56, 0x11,
	0xcf, 0x86, 0x11, 0x91, 0xaf, 0xc9, 0x38, 0x47, 0xb5, 0x02, 0xe5, 0xe3, 0xbb, 0x94, 0xc9, 0x97,
	0xd9, 0xc0, 0x09, 0xf9, 0x08, 0x8f, 0xf8, 0x80, 0x45, 0x8b, 0x6e, 0xef, 0xf0, 0xe2, 0x7b, 0x43,
	0x45, 0xb1, 0xca, 0x51, 0x12, 0x2f, 0x8b, 0x7c, 0xe5, 0x1e, 0xe5, 0x9c, 0x46, 0x04, 0x07, 0x82,
	0xe1, 0x20, 0x8e, 0xb9, 0x0c, 0x24, 0xe3, 0x71, 0x9a, 0xcf, 0x6e, 0x53, 0x4e, 0xb9, 0x2a, 0xf1,
	0xa2, 0xd2, 0x74, 0xff, 0x08, 0x6e, 0x76, 0xd5, 0xde, 0xba, 0x32, 0x90, 0x04, 0x21, 0x58, 0xeb,
	0xf6, 0xfc, 0x03, 0xf7, 0xe9, 0x8b, 0x7e, 0xa7, 0xdd, 0xf1, 0x9e, 0x77, 0x8c, 0x0a, 0x32, 0x60,
	0x35, 0x67, 0xdd, 0x9e, 0xeb, 0xf7, 0x0c, 0x80, 0xea, 0x70, 0x73, 0x49, 0xbc, 0x43, 0x63, 0xa3,
	0x14, 0x39, 0xf0, 0x7d, 0xcf, 0x37, 0x2e, 0xdc, 0x3e, 0xbd, 0x08, 0xab, 0xfa, 0xc7, 0xde, 0xe0,
	0x95, 0x2b, 0x18, 0xfa, 0x04, 0xe0, 0x96, 0x06, 0xae, 0x10, 0x8f, 0xe3, 0x54, 0xa2, 0x1d, 0x67,
	0x79, 0x8e, 0x4e, 0xce, 0xda, 0x64, 0xdc, 0x6c, 0x94, 0xb0, 0xaf, 0x4e, 0xd7, 0x7e, 0x36, 0x9b,
	0x9b, 0xd8, 0x27, 0x29, 0xcf, 0x92, 0x90, 0xe4, 0xd1, 0xb4, 0xe5, 0x86, 0x8b, 0xed, 0x1d, 0x31,
	0xf2, 0xb6, 0xe5, 0x0a, 0xd1, 0x26, 0x63, 0xc7, 0x4b, 0x68, 0x10, 0xb3, 0xf7, 0x6a, 0xdb, 0xdf,
	0xff, 0x98, 0xe0, 0xc3, 0x8f, 0xdf, 0x5f, 0x36, 0xb6, 0xed, 0x3a, 0xd6, 0xd7, 0x87, 0xf3, 0xeb,
	0xb9, 0x07, 0xf6, 0x6f, 0x02, 0xf4, 0x0d, 0xc0, 0x86, 0x16, 0x7a, 0xa0, 0xaf, 0x49, 0x49, 0x5d,
	0x2b, 0x75, 0x2f, 0xf1, 0x35, 0x62, 0xfe, 0x6c, 0x6e, 0xb6, 0x0a, 0xb1, 0x52, 0x7c, 0x45, 0xee,
	0xbf, 0x56, 0xa6, 0x7d, 0xa5, 0xb0, 0x2a, 0xbd, 0x12, 0x6d, 0xf6, 0x11, 0xc0, 0x5a, 0x61, 0xa6,
	0x9f, 0x06, 0xda, 0x5d, 0xd1, 0xd2, 0x70, 0x8d, 0xd3, 0x93, 0xd9, 0xdc, 0xbc, 0x7e, 0xe6, 0xa4,
	0xb3, 0xe7, 0x0b, 0xed, 0xd8, 0xc6, 0x99, 0x90, 0x5e, 0xa4, 0x6d, 0xbe, 0x02, 0x58, 0xd7, 0x36,
	0x8f, 0x0e, 0xfb, 0x0f, 0x13, 0xf6, 0x86, 0x24, 0xe8, 0x6a, 0xa9, 0xed, 0x92, 0xae, 0xf1, 0xe9,
	0xcf, 0xe6, 0xe6, 0xad, 0x7f, 0x7d, 0xdc, 0x38, 0x88, 0xc6, 0x92, 0x85, 0xe7, 0x7b, 0xed, 0xda,
	0x8d, 0xc2, 0x8b, 0x8a, 0x6c, 0xa8, 0x3a, 0x29, 0xb1, 0xfb, 0x7b, 0xc7, 0xbf, 0xac, 0xca, 0xf1,
	0xc4, 0x02, 0x27, 0x13, 0x0b, 0x9c, 0x4e, 0x2c, 0xf0, 0x79, 0x6a, 0x55, 0x4e, 0xa6, 0x56, 0xe5,
	0xe7, 0xd4, 0xaa, 0x0c, 0x2e, 0x29, 0x89, 0x3b, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x69, 0x5a,
	0xf9, 0x61, 0xbb, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StreamObjApiClient is the client API for StreamObjApi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StreamObjApiClient interface {
	// Stream Application Instance current progress
	StreamAppInst(ctx context.Context, in *AppInstKey, opts ...grpc.CallOption) (StreamObjApi_StreamAppInstClient, error)
	// Stream Cluster Instance current progress
	StreamClusterInst(ctx context.Context, in *ClusterInstKey, opts ...grpc.CallOption) (StreamObjApi_StreamClusterInstClient, error)
	// Stream Cloudlet current progress
	StreamCloudlet(ctx context.Context, in *CloudletKey, opts ...grpc.CallOption) (StreamObjApi_StreamCloudletClient, error)
	// Stream GPU driver current progress
	StreamGPUDriver(ctx context.Context, in *GPUDriverKey, opts ...grpc.CallOption) (StreamObjApi_StreamGPUDriverClient, error)
}

type streamObjApiClient struct {
	cc *grpc.ClientConn
}

func NewStreamObjApiClient(cc *grpc.ClientConn) StreamObjApiClient {
	return &streamObjApiClient{cc}
}

func (c *streamObjApiClient) StreamAppInst(ctx context.Context, in *AppInstKey, opts ...grpc.CallOption) (StreamObjApi_StreamAppInstClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamObjApi_serviceDesc.Streams[0], "/edgeproto.StreamObjApi/StreamAppInst", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamObjApiStreamAppInstClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamObjApi_StreamAppInstClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type streamObjApiStreamAppInstClient struct {
	grpc.ClientStream
}

func (x *streamObjApiStreamAppInstClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamObjApiClient) StreamClusterInst(ctx context.Context, in *ClusterInstKey, opts ...grpc.CallOption) (StreamObjApi_StreamClusterInstClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamObjApi_serviceDesc.Streams[1], "/edgeproto.StreamObjApi/StreamClusterInst", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamObjApiStreamClusterInstClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamObjApi_StreamClusterInstClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type streamObjApiStreamClusterInstClient struct {
	grpc.ClientStream
}

func (x *streamObjApiStreamClusterInstClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamObjApiClient) StreamCloudlet(ctx context.Context, in *CloudletKey, opts ...grpc.CallOption) (StreamObjApi_StreamCloudletClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamObjApi_serviceDesc.Streams[2], "/edgeproto.StreamObjApi/StreamCloudlet", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamObjApiStreamCloudletClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamObjApi_StreamCloudletClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type streamObjApiStreamCloudletClient struct {
	grpc.ClientStream
}

func (x *streamObjApiStreamCloudletClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamObjApiClient) StreamGPUDriver(ctx context.Context, in *GPUDriverKey, opts ...grpc.CallOption) (StreamObjApi_StreamGPUDriverClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StreamObjApi_serviceDesc.Streams[3], "/edgeproto.StreamObjApi/StreamGPUDriver", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamObjApiStreamGPUDriverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamObjApi_StreamGPUDriverClient interface {
	Recv() (*Result, error)
	grpc.ClientStream
}

type streamObjApiStreamGPUDriverClient struct {
	grpc.ClientStream
}

func (x *streamObjApiStreamGPUDriverClient) Recv() (*Result, error) {
	m := new(Result)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamObjApiServer is the server API for StreamObjApi service.
type StreamObjApiServer interface {
	// Stream Application Instance current progress
	StreamAppInst(*AppInstKey, StreamObjApi_StreamAppInstServer) error
	// Stream Cluster Instance current progress
	StreamClusterInst(*ClusterInstKey, StreamObjApi_StreamClusterInstServer) error
	// Stream Cloudlet current progress
	StreamCloudlet(*CloudletKey, StreamObjApi_StreamCloudletServer) error
	// Stream GPU driver current progress
	StreamGPUDriver(*GPUDriverKey, StreamObjApi_StreamGPUDriverServer) error
}

// UnimplementedStreamObjApiServer can be embedded to have forward compatible implementations.
type UnimplementedStreamObjApiServer struct {
}

func (*UnimplementedStreamObjApiServer) StreamAppInst(req *AppInstKey, srv StreamObjApi_StreamAppInstServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAppInst not implemented")
}
func (*UnimplementedStreamObjApiServer) StreamClusterInst(req *ClusterInstKey, srv StreamObjApi_StreamClusterInstServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamClusterInst not implemented")
}
func (*UnimplementedStreamObjApiServer) StreamCloudlet(req *CloudletKey, srv StreamObjApi_StreamCloudletServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamCloudlet not implemented")
}
func (*UnimplementedStreamObjApiServer) StreamGPUDriver(req *GPUDriverKey, srv StreamObjApi_StreamGPUDriverServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamGPUDriver not implemented")
}

func RegisterStreamObjApiServer(s *grpc.Server, srv StreamObjApiServer) {
	s.RegisterService(&_StreamObjApi_serviceDesc, srv)
}

func _StreamObjApi_StreamAppInst_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AppInstKey)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamObjApiServer).StreamAppInst(m, &streamObjApiStreamAppInstServer{stream})
}

type StreamObjApi_StreamAppInstServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type streamObjApiStreamAppInstServer struct {
	grpc.ServerStream
}

func (x *streamObjApiStreamAppInstServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamObjApi_StreamClusterInst_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ClusterInstKey)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamObjApiServer).StreamClusterInst(m, &streamObjApiStreamClusterInstServer{stream})
}

type StreamObjApi_StreamClusterInstServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type streamObjApiStreamClusterInstServer struct {
	grpc.ServerStream
}

func (x *streamObjApiStreamClusterInstServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamObjApi_StreamCloudlet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CloudletKey)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamObjApiServer).StreamCloudlet(m, &streamObjApiStreamCloudletServer{stream})
}

type StreamObjApi_StreamCloudletServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type streamObjApiStreamCloudletServer struct {
	grpc.ServerStream
}

func (x *streamObjApiStreamCloudletServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

func _StreamObjApi_StreamGPUDriver_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GPUDriverKey)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamObjApiServer).StreamGPUDriver(m, &streamObjApiStreamGPUDriverServer{stream})
}

type StreamObjApi_StreamGPUDriverServer interface {
	Send(*Result) error
	grpc.ServerStream
}

type streamObjApiStreamGPUDriverServer struct {
	grpc.ServerStream
}

func (x *streamObjApiStreamGPUDriverServer) Send(m *Result) error {
	return x.ServerStream.SendMsg(m)
}

var _StreamObjApi_serviceDesc = grpc.ServiceDesc{
	ServiceName: "edgeproto.StreamObjApi",
	HandlerType: (*StreamObjApiServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAppInst",
			Handler:       _StreamObjApi_StreamAppInst_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamClusterInst",
			Handler:       _StreamObjApi_StreamClusterInst_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamCloudlet",
			Handler:       _StreamObjApi_StreamCloudlet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "StreamGPUDriver",
			Handler:       _StreamObjApi_StreamGPUDriver_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stream.proto",
}

var StreamStateStrings = []string{
	"STREAM_UNKNOWN",
	"STREAM_START",
	"STREAM_STOP",
	"STREAM_ERROR",
}

const (
	StreamStateSTREAM_UNKNOWN uint64 = 1 << 0
	StreamStateSTREAM_START   uint64 = 1 << 1
	StreamStateSTREAM_STOP    uint64 = 1 << 2
	StreamStateSTREAM_ERROR   uint64 = 1 << 3
)

var StreamState_CamelName = map[int32]string{
	// STREAM_UNKNOWN -> StreamUnknown
	0: "StreamUnknown",
	// STREAM_START -> StreamStart
	1: "StreamStart",
	// STREAM_STOP -> StreamStop
	2: "StreamStop",
	// STREAM_ERROR -> StreamError
	3: "StreamError",
}
var StreamState_CamelValue = map[string]int32{
	"StreamUnknown": 0,
	"StreamStart":   1,
	"StreamStop":    2,
	"StreamError":   3,
}

func ParseStreamState(data interface{}) (StreamState, error) {
	if val, ok := data.(StreamState); ok {
		return val, nil
	} else if str, ok := data.(string); ok {
		val, ok := StreamState_CamelValue[util.CamelCase(str)]
		if !ok {
			// may have omitted common prefix
			val, ok = StreamState_CamelValue["Stream"+util.CamelCase(str)]
		}
		if !ok {
			// may be int value instead of enum name
			ival, err := strconv.Atoi(str)
			val = int32(ival)
			if err == nil {
				_, ok = StreamState_CamelName[val]
			}
		}
		if !ok {
			return StreamState(0), fmt.Errorf("Invalid StreamState value %q", str)
		}
		return StreamState(val), nil
	} else if ival, ok := data.(int32); ok {
		if _, ok := StreamState_CamelName[ival]; ok {
			return StreamState(ival), nil
		} else {
			return StreamState(0), fmt.Errorf("Invalid StreamState value %d", ival)
		}
	}
	return StreamState(0), fmt.Errorf("Invalid StreamState value %v", data)
}

func (e *StreamState) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var str string
	err := unmarshal(&str)
	if err != nil {
		return err
	}
	val, err := ParseStreamState(str)
	if err != nil {
		return err
	}
	*e = val
	return nil
}

func (e StreamState) MarshalYAML() (interface{}, error) {
	str := proto.EnumName(StreamState_CamelName, int32(e))
	str = strings.TrimPrefix(str, "Stream")
	return str, nil
}

// custom JSON encoding/decoding
func (e *StreamState) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err == nil {
		val, err := ParseStreamState(str)
		if err != nil {
			return &json.UnmarshalTypeError{
				Value: "string " + str,
				Type:  reflect.TypeOf(StreamState(0)),
			}
		}
		*e = StreamState(val)
		return nil
	}
	var ival int32
	err = json.Unmarshal(b, &ival)
	if err == nil {
		val, err := ParseStreamState(ival)
		if err == nil {
			*e = val
			return nil
		}
	}
	return &json.UnmarshalTypeError{
		Value: "value " + string(b),
		Type:  reflect.TypeOf(StreamState(0)),
	}
}

func (e StreamState) MarshalJSON() ([]byte, error) {
	str := proto.EnumName(StreamState_CamelName, int32(e))
	str = strings.TrimPrefix(str, "Stream")
	return json.Marshal(str)
}

var StreamStateCommonPrefix = "Stream"

func (m *AppInstKey) IsValidArgsForStreamAppInst() error {
	return nil
}

func (m *ClusterInstKey) IsValidArgsForStreamClusterInst() error {
	return nil
}

func (m *CloudletKey) IsValidArgsForStreamCloudlet() error {
	return nil
}

func (m *GPUDriverKey) IsValidArgsForStreamGPUDriver() error {
	return nil
}
