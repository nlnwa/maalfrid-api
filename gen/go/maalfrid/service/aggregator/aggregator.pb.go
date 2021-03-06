// Code generated by protoc-gen-go. DO NOT EDIT.
// source: maalfrid/service/aggregator/aggregator.proto

package aggregator

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import empty "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Label struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_aggregator_325aacbd761c3a1c, []int{0}
}
func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (dst *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(dst, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SyncEntitiesRequest struct {
	// if set then only entities with a matching name will be synchronized
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// if set then only entities labeled with any of the labels will be synchronized
	Labels               []*Label `protobuf:"bytes,2,rep,name=labels" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncEntitiesRequest) Reset()         { *m = SyncEntitiesRequest{} }
func (m *SyncEntitiesRequest) String() string { return proto.CompactTextString(m) }
func (*SyncEntitiesRequest) ProtoMessage()    {}
func (*SyncEntitiesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aggregator_325aacbd761c3a1c, []int{1}
}
func (m *SyncEntitiesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncEntitiesRequest.Unmarshal(m, b)
}
func (m *SyncEntitiesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncEntitiesRequest.Marshal(b, m, deterministic)
}
func (dst *SyncEntitiesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncEntitiesRequest.Merge(dst, src)
}
func (m *SyncEntitiesRequest) XXX_Size() int {
	return xxx_messageInfo_SyncEntitiesRequest.Size(m)
}
func (m *SyncEntitiesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncEntitiesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncEntitiesRequest proto.InternalMessageInfo

func (m *SyncEntitiesRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SyncEntitiesRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type RunAggregationRequest struct {
	JobExecutionId       string   `protobuf:"bytes,1,opt,name=job_execution_id,json=jobExecutionId" json:"job_execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunAggregationRequest) Reset()         { *m = RunAggregationRequest{} }
func (m *RunAggregationRequest) String() string { return proto.CompactTextString(m) }
func (*RunAggregationRequest) ProtoMessage()    {}
func (*RunAggregationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aggregator_325aacbd761c3a1c, []int{2}
}
func (m *RunAggregationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunAggregationRequest.Unmarshal(m, b)
}
func (m *RunAggregationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunAggregationRequest.Marshal(b, m, deterministic)
}
func (dst *RunAggregationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunAggregationRequest.Merge(dst, src)
}
func (m *RunAggregationRequest) XXX_Size() int {
	return xxx_messageInfo_RunAggregationRequest.Size(m)
}
func (m *RunAggregationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunAggregationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunAggregationRequest proto.InternalMessageInfo

func (m *RunAggregationRequest) GetJobExecutionId() string {
	if m != nil {
		return m.JobExecutionId
	}
	return ""
}

type FilterAggregateRequest struct {
	JobExecutionId       string   `protobuf:"bytes,1,opt,name=job_execution_id,json=jobExecutionId" json:"job_execution_id,omitempty"`
	SeedId               string   `protobuf:"bytes,2,opt,name=seed_id,json=seedId" json:"seed_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FilterAggregateRequest) Reset()         { *m = FilterAggregateRequest{} }
func (m *FilterAggregateRequest) String() string { return proto.CompactTextString(m) }
func (*FilterAggregateRequest) ProtoMessage()    {}
func (*FilterAggregateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aggregator_325aacbd761c3a1c, []int{3}
}
func (m *FilterAggregateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FilterAggregateRequest.Unmarshal(m, b)
}
func (m *FilterAggregateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FilterAggregateRequest.Marshal(b, m, deterministic)
}
func (dst *FilterAggregateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FilterAggregateRequest.Merge(dst, src)
}
func (m *FilterAggregateRequest) XXX_Size() int {
	return xxx_messageInfo_FilterAggregateRequest.Size(m)
}
func (m *FilterAggregateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FilterAggregateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FilterAggregateRequest proto.InternalMessageInfo

func (m *FilterAggregateRequest) GetJobExecutionId() string {
	if m != nil {
		return m.JobExecutionId
	}
	return ""
}

func (m *FilterAggregateRequest) GetSeedId() string {
	if m != nil {
		return m.SeedId
	}
	return ""
}

type RunLanguageDetectionRequest struct {
	// If language detection should process extracted texts already processed
	DetectAll            bool     `protobuf:"varint,1,opt,name=detect_all,json=detectAll" json:"detect_all,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunLanguageDetectionRequest) Reset()         { *m = RunLanguageDetectionRequest{} }
func (m *RunLanguageDetectionRequest) String() string { return proto.CompactTextString(m) }
func (*RunLanguageDetectionRequest) ProtoMessage()    {}
func (*RunLanguageDetectionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aggregator_325aacbd761c3a1c, []int{4}
}
func (m *RunLanguageDetectionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunLanguageDetectionRequest.Unmarshal(m, b)
}
func (m *RunLanguageDetectionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunLanguageDetectionRequest.Marshal(b, m, deterministic)
}
func (dst *RunLanguageDetectionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunLanguageDetectionRequest.Merge(dst, src)
}
func (m *RunLanguageDetectionRequest) XXX_Size() int {
	return xxx_messageInfo_RunLanguageDetectionRequest.Size(m)
}
func (m *RunLanguageDetectionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunLanguageDetectionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunLanguageDetectionRequest proto.InternalMessageInfo

func (m *RunLanguageDetectionRequest) GetDetectAll() bool {
	if m != nil {
		return m.DetectAll
	}
	return false
}

func init() {
	proto.RegisterType((*Label)(nil), "maalfrid.service.aggregator.Label")
	proto.RegisterType((*SyncEntitiesRequest)(nil), "maalfrid.service.aggregator.SyncEntitiesRequest")
	proto.RegisterType((*RunAggregationRequest)(nil), "maalfrid.service.aggregator.RunAggregationRequest")
	proto.RegisterType((*FilterAggregateRequest)(nil), "maalfrid.service.aggregator.FilterAggregateRequest")
	proto.RegisterType((*RunLanguageDetectionRequest)(nil), "maalfrid.service.aggregator.RunLanguageDetectionRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Aggregator service

type AggregatorClient interface {
	// Detect language of extracted texts
	RunLanguageDetection(ctx context.Context, in *RunLanguageDetectionRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Aggregate data from veidemann to maalfrid
	RunAggregation(ctx context.Context, in *RunAggregationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Synchronize entities from veidemann to maalfrid
	SyncEntities(ctx context.Context, in *SyncEntitiesRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	FilterAggregate(ctx context.Context, in *FilterAggregateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type aggregatorClient struct {
	cc *grpc.ClientConn
}

func NewAggregatorClient(cc *grpc.ClientConn) AggregatorClient {
	return &aggregatorClient{cc}
}

func (c *aggregatorClient) RunLanguageDetection(ctx context.Context, in *RunLanguageDetectionRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/maalfrid.service.aggregator.Aggregator/RunLanguageDetection", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) RunAggregation(ctx context.Context, in *RunAggregationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/maalfrid.service.aggregator.Aggregator/RunAggregation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) SyncEntities(ctx context.Context, in *SyncEntitiesRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/maalfrid.service.aggregator.Aggregator/SyncEntities", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aggregatorClient) FilterAggregate(ctx context.Context, in *FilterAggregateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := grpc.Invoke(ctx, "/maalfrid.service.aggregator.Aggregator/FilterAggregate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Aggregator service

type AggregatorServer interface {
	// Detect language of extracted texts
	RunLanguageDetection(context.Context, *RunLanguageDetectionRequest) (*empty.Empty, error)
	// Aggregate data from veidemann to maalfrid
	RunAggregation(context.Context, *RunAggregationRequest) (*empty.Empty, error)
	// Synchronize entities from veidemann to maalfrid
	SyncEntities(context.Context, *SyncEntitiesRequest) (*empty.Empty, error)
	FilterAggregate(context.Context, *FilterAggregateRequest) (*empty.Empty, error)
}

func RegisterAggregatorServer(s *grpc.Server, srv AggregatorServer) {
	s.RegisterService(&_Aggregator_serviceDesc, srv)
}

func _Aggregator_RunLanguageDetection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunLanguageDetectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).RunLanguageDetection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/maalfrid.service.aggregator.Aggregator/RunLanguageDetection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).RunLanguageDetection(ctx, req.(*RunLanguageDetectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_RunAggregation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunAggregationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).RunAggregation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/maalfrid.service.aggregator.Aggregator/RunAggregation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).RunAggregation(ctx, req.(*RunAggregationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_SyncEntities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SyncEntitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).SyncEntities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/maalfrid.service.aggregator.Aggregator/SyncEntities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).SyncEntities(ctx, req.(*SyncEntitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Aggregator_FilterAggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterAggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AggregatorServer).FilterAggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/maalfrid.service.aggregator.Aggregator/FilterAggregate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AggregatorServer).FilterAggregate(ctx, req.(*FilterAggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Aggregator_serviceDesc = grpc.ServiceDesc{
	ServiceName: "maalfrid.service.aggregator.Aggregator",
	HandlerType: (*AggregatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunLanguageDetection",
			Handler:    _Aggregator_RunLanguageDetection_Handler,
		},
		{
			MethodName: "RunAggregation",
			Handler:    _Aggregator_RunAggregation_Handler,
		},
		{
			MethodName: "SyncEntities",
			Handler:    _Aggregator_SyncEntities_Handler,
		},
		{
			MethodName: "FilterAggregate",
			Handler:    _Aggregator_FilterAggregate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "maalfrid/service/aggregator/aggregator.proto",
}

func init() {
	proto.RegisterFile("maalfrid/service/aggregator/aggregator.proto", fileDescriptor_aggregator_325aacbd761c3a1c)
}

var fileDescriptor_aggregator_325aacbd761c3a1c = []byte{
	// 382 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4f, 0xea, 0x40,
	0x14, 0x7d, 0xc0, 0x83, 0xf7, 0xb8, 0x12, 0x24, 0x23, 0x22, 0x81, 0x98, 0x90, 0xae, 0x58, 0x98,
	0xa9, 0x81, 0x8d, 0x31, 0x6e, 0x6a, 0xc4, 0x84, 0x84, 0x55, 0xdd, 0x61, 0x22, 0x4e, 0xdb, 0x4b,
	0x53, 0x1c, 0x3a, 0xd8, 0x4e, 0x89, 0xfc, 0x11, 0x7f, 0xaf, 0xe9, 0x97, 0xa0, 0xe2, 0x18, 0x77,
	0x77, 0xce, 0x9c, 0x7b, 0xce, 0x9d, 0x7b, 0x06, 0xce, 0x96, 0x8c, 0xf1, 0x79, 0xe0, 0x39, 0x7a,
	0x88, 0xc1, 0xda, 0xb3, 0x51, 0x67, 0xae, 0x1b, 0xa0, 0xcb, 0xa4, 0x08, 0x76, 0x4a, 0xba, 0x0a,
	0x84, 0x14, 0xa4, 0x9b, 0xb3, 0x69, 0xc6, 0xa6, 0x5b, 0x4a, 0xa7, 0xeb, 0x0a, 0xe1, 0x72, 0xd4,
	0x13, 0xaa, 0x15, 0xcd, 0x75, 0x5c, 0xae, 0xe4, 0x26, 0xed, 0xd4, 0x74, 0x28, 0x4f, 0x98, 0x85,
	0x9c, 0x34, 0xa0, 0xf4, 0x84, 0x9b, 0x76, 0xa1, 0x57, 0xe8, 0x57, 0xcd, 0xb8, 0x24, 0x4d, 0x28,
	0xaf, 0x19, 0x8f, 0xb0, 0x5d, 0x4c, 0xb0, 0xf4, 0xa0, 0x21, 0x1c, 0xdd, 0x6d, 0x7c, 0x7b, 0xe4,
	0x4b, 0x4f, 0x7a, 0x18, 0x9a, 0xf8, 0x1c, 0x61, 0x28, 0x09, 0x81, 0xbf, 0x3e, 0x5b, 0x62, 0xd6,
	0x9f, 0xd4, 0xe4, 0x12, 0x2a, 0x3c, 0xd6, 0x0e, 0xdb, 0xc5, 0x5e, 0xa9, 0x7f, 0x30, 0xd0, 0xa8,
	0x62, 0x4c, 0x9a, 0x8c, 0x61, 0x66, 0x1d, 0x9a, 0x01, 0xc7, 0x66, 0xe4, 0x1b, 0xd9, 0xb5, 0x27,
	0xfc, 0xdc, 0xa8, 0x0f, 0x8d, 0x85, 0xb0, 0x66, 0xf8, 0x82, 0x76, 0x14, 0xe3, 0x33, 0xcf, 0xc9,
	0x4c, 0xeb, 0x0b, 0x61, 0x8d, 0x72, 0x78, 0xec, 0x68, 0xf7, 0xd0, 0xba, 0xf5, 0xb8, 0xc4, 0x20,
	0x57, 0xc1, 0x5f, 0x6b, 0x90, 0x13, 0xf8, 0x17, 0x22, 0x3a, 0x31, 0x21, 0xdd, 0x42, 0x25, 0x3e,
	0x8e, 0x1d, 0xed, 0x0a, 0xba, 0x66, 0xe4, 0x4f, 0x98, 0xef, 0x46, 0xcc, 0xc5, 0x1b, 0x94, 0x68,
	0xef, 0x4e, 0x79, 0x0a, 0xe0, 0x24, 0xd8, 0x8c, 0x71, 0x9e, 0x68, 0xff, 0x37, 0xab, 0x29, 0x62,
	0x70, 0x3e, 0x78, 0x2d, 0x01, 0x18, 0xef, 0x4f, 0x27, 0x0b, 0x68, 0xee, 0x13, 0x23, 0x17, 0xca,
	0x85, 0x29, 0xfc, 0x3b, 0x2d, 0x9a, 0x86, 0x4e, 0xf3, 0xd0, 0xe9, 0x28, 0x0e, 0x5d, 0xfb, 0x43,
	0x1e, 0xa0, 0xfe, 0x71, 0xb1, 0x64, 0xf0, 0x93, 0xcb, 0xd7, 0x14, 0x14, 0xfa, 0x53, 0xa8, 0xed,
	0xfe, 0x0f, 0x72, 0xae, 0x54, 0xdf, 0xf3, 0x95, 0x14, 0xda, 0x8f, 0x70, 0xf8, 0x29, 0x51, 0x32,
	0x54, 0xca, 0xef, 0xcf, 0xff, 0x7b, 0x87, 0xeb, 0xda, 0x14, 0xb6, 0xed, 0x56, 0x25, 0xb9, 0x1f,
	0xbe, 0x05, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x4a, 0xcf, 0x8b, 0x8d, 0x03, 0x00, 0x00,
}
