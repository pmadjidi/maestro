// Code generated by protoc-gen-go. DO NOT EDIT.
// source: register.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type RegisterReq struct {
	UserName             string               `protobuf:"bytes,1,opt,name=userName,proto3" json:"userName,omitempty"`
	PassWord             []byte               `protobuf:"bytes,2,opt,name=passWord,proto3" json:"passWord,omitempty"`
	FirstName            string               `protobuf:"bytes,3,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName             string               `protobuf:"bytes,4,opt,name=lastName,proto3" json:"lastName,omitempty"`
	Email                string               `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Phone                string               `protobuf:"bytes,6,opt,name=phone,proto3" json:"phone,omitempty"`
	Address              *RegisterReq_Address `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	Device               string               `protobuf:"bytes,8,opt,name=device,proto3" json:"device,omitempty"`
	TimeName             *timestamp.Timestamp `protobuf:"bytes,9,opt,name=time_name,json=timeName,proto3" json:"time_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *RegisterReq) Reset()         { *m = RegisterReq{} }
func (m *RegisterReq) String() string { return proto.CompactTextString(m) }
func (*RegisterReq) ProtoMessage()    {}
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0}
}

func (m *RegisterReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq.Unmarshal(m, b)
}
func (m *RegisterReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq.Marshal(b, m, deterministic)
}
func (m *RegisterReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq.Merge(m, src)
}
func (m *RegisterReq) XXX_Size() int {
	return xxx_messageInfo_RegisterReq.Size(m)
}
func (m *RegisterReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq proto.InternalMessageInfo

func (m *RegisterReq) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *RegisterReq) GetPassWord() []byte {
	if m != nil {
		return m.PassWord
	}
	return nil
}

func (m *RegisterReq) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *RegisterReq) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *RegisterReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterReq) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *RegisterReq) GetAddress() *RegisterReq_Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *RegisterReq) GetDevice() string {
	if m != nil {
		return m.Device
	}
	return ""
}

func (m *RegisterReq) GetTimeName() *timestamp.Timestamp {
	if m != nil {
		return m.TimeName
	}
	return nil
}

type RegisterReq_Address struct {
	Street               string   `protobuf:"bytes,1,opt,name=street,proto3" json:"street,omitempty"`
	City                 string   `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	State                string   `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`
	Zip                  string   `protobuf:"bytes,4,opt,name=zip,proto3" json:"zip,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterReq_Address) Reset()         { *m = RegisterReq_Address{} }
func (m *RegisterReq_Address) String() string { return proto.CompactTextString(m) }
func (*RegisterReq_Address) ProtoMessage()    {}
func (*RegisterReq_Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{0, 0}
}

func (m *RegisterReq_Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterReq_Address.Unmarshal(m, b)
}
func (m *RegisterReq_Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterReq_Address.Marshal(b, m, deterministic)
}
func (m *RegisterReq_Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterReq_Address.Merge(m, src)
}
func (m *RegisterReq_Address) XXX_Size() int {
	return xxx_messageInfo_RegisterReq_Address.Size(m)
}
func (m *RegisterReq_Address) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterReq_Address.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterReq_Address proto.InternalMessageInfo

func (m *RegisterReq_Address) GetStreet() string {
	if m != nil {
		return m.Street
	}
	return ""
}

func (m *RegisterReq_Address) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *RegisterReq_Address) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *RegisterReq_Address) GetZip() string {
	if m != nil {
		return m.Zip
	}
	return ""
}

type RegisterResp struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status               Status   `protobuf:"varint,2,opt,name=status,proto3,enum=api.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegisterResp) Reset()         { *m = RegisterResp{} }
func (m *RegisterResp) String() string { return proto.CompactTextString(m) }
func (*RegisterResp) ProtoMessage()    {}
func (*RegisterResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_1303fe8288f4efb6, []int{1}
}

func (m *RegisterResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegisterResp.Unmarshal(m, b)
}
func (m *RegisterResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegisterResp.Marshal(b, m, deterministic)
}
func (m *RegisterResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisterResp.Merge(m, src)
}
func (m *RegisterResp) XXX_Size() int {
	return xxx_messageInfo_RegisterResp.Size(m)
}
func (m *RegisterResp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisterResp.DiscardUnknown(m)
}

var xxx_messageInfo_RegisterResp proto.InternalMessageInfo

func (m *RegisterResp) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegisterResp) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_SUCCESS
}

func init() {
	proto.RegisterType((*RegisterReq)(nil), "api.RegisterReq")
	proto.RegisterType((*RegisterReq_Address)(nil), "api.RegisterReq.Address")
	proto.RegisterType((*RegisterResp)(nil), "api.RegisterResp")
}

func init() { proto.RegisterFile("register.proto", fileDescriptor_1303fe8288f4efb6) }

var fileDescriptor_1303fe8288f4efb6 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0x4d, 0x6b, 0xfa, 0x40,
	0x10, 0xc6, 0x89, 0xf1, 0x25, 0x19, 0x45, 0xfc, 0x2f, 0x7f, 0xca, 0x12, 0x0a, 0x15, 0x7b, 0xf1,
	0xb4, 0x42, 0x7a, 0xe8, 0xa1, 0x27, 0xe9, 0xbd, 0x87, 0x6d, 0xa1, 0xa7, 0x52, 0x56, 0x33, 0xda,
	0x05, 0x63, 0xb6, 0xbb, 0x6b, 0xa1, 0xfd, 0x12, 0xfd, 0xca, 0x65, 0x5f, 0xa2, 0xe2, 0x6d, 0x9e,
	0x67, 0x66, 0x9e, 0x64, 0x7e, 0x0b, 0x63, 0x8d, 0x5b, 0x69, 0x2c, 0x6a, 0xa6, 0x74, 0x63, 0x1b,
	0x92, 0x0a, 0x25, 0x8b, 0x91, 0xb1, 0xc2, 0x1e, 0x4c, 0xb0, 0x8a, 0x9b, 0x6d, 0xd3, 0x6c, 0x77,
	0xb8, 0xf0, 0x6a, 0x75, 0xd8, 0x2c, 0xac, 0xac, 0xd1, 0x58, 0x51, 0xab, 0x30, 0x30, 0xfb, 0x4d,
	0x61, 0xc8, 0x63, 0x0c, 0xc7, 0x4f, 0x52, 0x40, 0x76, 0x30, 0xa8, 0x9f, 0x44, 0x8d, 0x34, 0x99,
	0x26, 0xf3, 0x9c, 0x1f, 0xb5, 0xeb, 0x29, 0x61, 0xcc, 0x6b, 0xa3, 0x2b, 0xda, 0x99, 0x26, 0xf3,
	0x11, 0x3f, 0x6a, 0x72, 0x0d, 0xf9, 0x46, 0x6a, 0x63, 0xfd, 0x62, 0xea, 0x17, 0x4f, 0x86, 0xdb,
	0xdc, 0x89, 0xd8, 0xec, 0x86, 0xd4, 0x56, 0x93, 0xff, 0xd0, 0xc3, 0x5a, 0xc8, 0x1d, 0xed, 0xf9,
	0x46, 0x10, 0xce, 0x55, 0x1f, 0xcd, 0x1e, 0x69, 0x3f, 0xb8, 0x5e, 0x90, 0x12, 0x06, 0xa2, 0xaa,
	0x34, 0x1a, 0x43, 0x07, 0xd3, 0x64, 0x3e, 0x2c, 0x29, 0x13, 0x4a, 0xb2, 0xb3, 0x03, 0xd8, 0x32,
	0xf4, 0x79, 0x3b, 0x48, 0xae, 0xa0, 0x5f, 0xe1, 0x97, 0x5c, 0x23, 0xcd, 0x7c, 0x54, 0x54, 0xe4,
	0x1e, 0x72, 0x07, 0xe3, 0x7d, 0xef, 0x7e, 0x2a, 0xf7, 0x69, 0x05, 0x0b, 0xb8, 0x58, 0x8b, 0x8b,
	0xbd, 0xb4, 0xb8, 0x78, 0xe6, 0x86, 0xdd, 0x0f, 0x17, 0x6f, 0x30, 0x58, 0x9e, 0xb2, 0x8d, 0xd5,
	0x88, 0x36, 0xb2, 0x8a, 0x8a, 0x10, 0xe8, 0xae, 0xa5, 0xfd, 0xf6, 0x94, 0x72, 0xee, 0x6b, 0x77,
	0x91, 0x7b, 0x9a, 0x96, 0x4e, 0x10, 0x64, 0x02, 0xe9, 0x8f, 0x54, 0x11, 0x8a, 0x2b, 0x67, 0x8f,
	0x30, 0x3a, 0xdd, 0x63, 0x14, 0x19, 0x43, 0x47, 0x56, 0x31, 0xbf, 0x23, 0x2b, 0x72, 0xeb, 0xbe,
	0xe9, 0x9e, 0xd8, 0xa7, 0x8f, 0xcb, 0xa1, 0x47, 0xf0, 0xec, 0x2d, 0x1e, 0x5b, 0xe5, 0x03, 0x64,
	0x6d, 0x08, 0x59, 0x9c, 0xd5, 0x93, 0x4b, 0x5e, 0xc5, 0xbf, 0x0b, 0xc7, 0xa8, 0x55, 0xdf, 0x9f,
	0x7f, 0xf7, 0x17, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x89, 0xa0, 0xb6, 0x60, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RegisterClient is the client API for Register service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RegisterClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
}

type registerClient struct {
	cc *grpc.ClientConn
}

func NewRegisterClient(cc *grpc.ClientConn) RegisterClient {
	return &registerClient{cc}
}

func (c *registerClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, "/api.Register/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegisterServer is the server API for Register service.
type RegisterServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
}

// UnimplementedRegisterServer can be embedded to have forward compatible implementations.
type UnimplementedRegisterServer struct {
}

func (*UnimplementedRegisterServer) Register(ctx context.Context, req *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterRegisterServer(s *grpc.Server, srv RegisterServer) {
	s.RegisterService(&_Register_serviceDesc, srv)
}

func _Register_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegisterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Register/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegisterServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Register_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Register",
	HandlerType: (*RegisterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Register_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "register.proto",
}
