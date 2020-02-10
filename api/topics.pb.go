// Code generated by protoc-gen-go. DO NOT EDIT.
// source: topics.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Topic struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Tag                  string   `protobuf:"bytes,2,opt,name=tag,proto3" json:"tag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Topic) Reset()         { *m = Topic{} }
func (m *Topic) String() string { return proto.CompactTextString(m) }
func (*Topic) ProtoMessage()    {}
func (*Topic) Descriptor() ([]byte, []int) {
	return fileDescriptor_72af2b0eeeefc419, []int{0}
}

func (m *Topic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Topic.Unmarshal(m, b)
}
func (m *Topic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Topic.Marshal(b, m, deterministic)
}
func (m *Topic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Topic.Merge(m, src)
}
func (m *Topic) XXX_Size() int {
	return xxx_messageInfo_Topic.Size(m)
}
func (m *Topic) XXX_DiscardUnknown() {
	xxx_messageInfo_Topic.DiscardUnknown(m)
}

var xxx_messageInfo_Topic proto.InternalMessageInfo

func (m *Topic) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Topic) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

type TopicReq struct {
	List                 []*Topic `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicReq) Reset()         { *m = TopicReq{} }
func (m *TopicReq) String() string { return proto.CompactTextString(m) }
func (*TopicReq) ProtoMessage()    {}
func (*TopicReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_72af2b0eeeefc419, []int{1}
}

func (m *TopicReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicReq.Unmarshal(m, b)
}
func (m *TopicReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicReq.Marshal(b, m, deterministic)
}
func (m *TopicReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicReq.Merge(m, src)
}
func (m *TopicReq) XXX_Size() int {
	return xxx_messageInfo_TopicReq.Size(m)
}
func (m *TopicReq) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicReq.DiscardUnknown(m)
}

var xxx_messageInfo_TopicReq proto.InternalMessageInfo

func (m *TopicReq) GetList() []*Topic {
	if m != nil {
		return m.List
	}
	return nil
}

type TopicResp struct {
	List                 []*Topic `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TopicResp) Reset()         { *m = TopicResp{} }
func (m *TopicResp) String() string { return proto.CompactTextString(m) }
func (*TopicResp) ProtoMessage()    {}
func (*TopicResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_72af2b0eeeefc419, []int{2}
}

func (m *TopicResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicResp.Unmarshal(m, b)
}
func (m *TopicResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicResp.Marshal(b, m, deterministic)
}
func (m *TopicResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicResp.Merge(m, src)
}
func (m *TopicResp) XXX_Size() int {
	return xxx_messageInfo_TopicResp.Size(m)
}
func (m *TopicResp) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicResp.DiscardUnknown(m)
}

var xxx_messageInfo_TopicResp proto.InternalMessageInfo

func (m *TopicResp) GetList() []*Topic {
	if m != nil {
		return m.List
	}
	return nil
}

func init() {
	proto.RegisterType((*Topic)(nil), "api.Topic")
	proto.RegisterType((*TopicReq)(nil), "api.TopicReq")
	proto.RegisterType((*TopicResp)(nil), "api.TopicResp")
}

func init() { proto.RegisterFile("topics.proto", fileDescriptor_72af2b0eeeefc419) }

var fileDescriptor_72af2b0eeeefc419 = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xc9, 0x2f, 0xc8,
	0x4c, 0x2e, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0xe2, 0x2b,
	0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x49, 0x2d, 0x82, 0x08, 0x2a, 0x69, 0x72, 0xb1, 0x86, 0x80, 0x14,
	0x09, 0xf1, 0x71, 0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6,
	0x08, 0x09, 0x70, 0x31, 0x97, 0x24, 0xa6, 0x4b, 0x30, 0x81, 0x05, 0x40, 0x4c, 0x25, 0x2d, 0x2e,
	0x0e, 0xb0, 0xd2, 0xa0, 0xd4, 0x42, 0x21, 0x39, 0x2e, 0x96, 0x9c, 0xcc, 0xe2, 0x12, 0x09, 0x46,
	0x05, 0x66, 0x0d, 0x6e, 0x23, 0x2e, 0xbd, 0xc4, 0x82, 0x4c, 0x3d, 0x88, 0x24, 0x58, 0x5c, 0x49,
	0x9b, 0x8b, 0x13, 0xaa, 0xb6, 0xb8, 0x80, 0x90, 0x62, 0xa3, 0x46, 0x46, 0x2e, 0xde, 0xe0, 0xd2,
	0xa4, 0xe2, 0xe4, 0xa2, 0xcc, 0x82, 0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x21, 0x15, 0x2e, 0xe6, 0xe2,
	0xd2, 0x24, 0x21, 0x5e, 0x24, 0xa5, 0xa9, 0x85, 0x52, 0x7c, 0xc8, 0xdc, 0xe2, 0x02, 0x21, 0x35,
	0x2e, 0xd6, 0xd2, 0x3c, 0x22, 0xd4, 0x29, 0x41, 0xec, 0x17, 0x82, 0xd8, 0xec, 0x9a, 0x5b, 0x50,
	0x52, 0x89, 0xae, 0x26, 0x89, 0x0d, 0x1c, 0x1c, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xda,
	0x41, 0x35, 0xc4, 0x33, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SubscriptionsClient is the client API for Subscriptions service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SubscriptionsClient interface {
	Sub(ctx context.Context, in *TopicReq, opts ...grpc.CallOption) (*TopicResp, error)
	Unsub(ctx context.Context, in *TopicReq, opts ...grpc.CallOption) (*TopicResp, error)
	List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicResp, error)
}

type subscriptionsClient struct {
	cc *grpc.ClientConn
}

func NewSubscriptionsClient(cc *grpc.ClientConn) SubscriptionsClient {
	return &subscriptionsClient{cc}
}

func (c *subscriptionsClient) Sub(ctx context.Context, in *TopicReq, opts ...grpc.CallOption) (*TopicResp, error) {
	out := new(TopicResp)
	err := c.cc.Invoke(ctx, "/api.Subscriptions/sub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsClient) Unsub(ctx context.Context, in *TopicReq, opts ...grpc.CallOption) (*TopicResp, error) {
	out := new(TopicResp)
	err := c.cc.Invoke(ctx, "/api.Subscriptions/unsub", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *subscriptionsClient) List(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*TopicResp, error) {
	out := new(TopicResp)
	err := c.cc.Invoke(ctx, "/api.Subscriptions/list", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionsServer is the server API for Subscriptions service.
type SubscriptionsServer interface {
	Sub(context.Context, *TopicReq) (*TopicResp, error)
	Unsub(context.Context, *TopicReq) (*TopicResp, error)
	List(context.Context, *Empty) (*TopicResp, error)
}

// UnimplementedSubscriptionsServer can be embedded to have forward compatible implementations.
type UnimplementedSubscriptionsServer struct {
}

func (*UnimplementedSubscriptionsServer) Sub(ctx context.Context, req *TopicReq) (*TopicResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sub not implemented")
}
func (*UnimplementedSubscriptionsServer) Unsub(ctx context.Context, req *TopicReq) (*TopicResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unsub not implemented")
}
func (*UnimplementedSubscriptionsServer) List(ctx context.Context, req *Empty) (*TopicResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterSubscriptionsServer(s *grpc.Server, srv SubscriptionsServer) {
	s.RegisterService(&_Subscriptions_serviceDesc, srv)
}

func _Subscriptions_Sub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).Sub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Subscriptions/Sub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).Sub(ctx, req.(*TopicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriptions_Unsub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TopicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).Unsub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Subscriptions/Unsub",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).Unsub(ctx, req.(*TopicReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Subscriptions_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionsServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Subscriptions/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionsServer).List(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Subscriptions_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Subscriptions",
	HandlerType: (*SubscriptionsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "sub",
			Handler:    _Subscriptions_Sub_Handler,
		},
		{
			MethodName: "unsub",
			Handler:    _Subscriptions_Unsub_Handler,
		},
		{
			MethodName: "list",
			Handler:    _Subscriptions_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "topics.proto",
}
