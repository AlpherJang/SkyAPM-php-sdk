// Code generated by protoc-gen-go. DO NOT EDIT.
// source: language-agent/ApplicationRegisterService.proto

package agent

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

type Application struct {
	ApplicationCode      string   `protobuf:"bytes,1,opt,name=applicationCode,proto3" json:"applicationCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1c2b07c1f73c603, []int{0}
}

func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (m *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(m, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetApplicationCode() string {
	if m != nil {
		return m.ApplicationCode
	}
	return ""
}

type ApplicationMapping struct {
	Application          *KeyWithIntegerValue `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *ApplicationMapping) Reset()         { *m = ApplicationMapping{} }
func (m *ApplicationMapping) String() string { return proto.CompactTextString(m) }
func (*ApplicationMapping) ProtoMessage()    {}
func (*ApplicationMapping) Descriptor() ([]byte, []int) {
	return fileDescriptor_e1c2b07c1f73c603, []int{1}
}

func (m *ApplicationMapping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMapping.Unmarshal(m, b)
}
func (m *ApplicationMapping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMapping.Marshal(b, m, deterministic)
}
func (m *ApplicationMapping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMapping.Merge(m, src)
}
func (m *ApplicationMapping) XXX_Size() int {
	return xxx_messageInfo_ApplicationMapping.Size(m)
}
func (m *ApplicationMapping) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMapping.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMapping proto.InternalMessageInfo

func (m *ApplicationMapping) GetApplication() *KeyWithIntegerValue {
	if m != nil {
		return m.Application
	}
	return nil
}

func init() {
	proto.RegisterType((*Application)(nil), "Application")
	proto.RegisterType((*ApplicationMapping)(nil), "ApplicationMapping")
}

func init() {
	proto.RegisterFile("language-agent/ApplicationRegisterService.proto", fileDescriptor_e1c2b07c1f73c603)
}

var fileDescriptor_e1c2b07c1f73c603 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0x87, 0x5d, 0x0f, 0x82, 0x59, 0x41, 0x88, 0x8a, 0xb2, 0x27, 0xd9, 0xd3, 0x5e, 0x9c, 0x4a,
	0x85, 0x7a, 0x13, 0xd4, 0x93, 0xf8, 0x87, 0xb2, 0x05, 0x0b, 0xe2, 0x65, 0xba, 0x0e, 0xd9, 0xb0,
	0x31, 0x09, 0x69, 0x6a, 0xd9, 0x57, 0xf2, 0x29, 0xa5, 0x9b, 0x16, 0xc3, 0x6a, 0x2f, 0x21, 0x19,
	0x7e, 0xdf, 0x4c, 0xe6, 0x63, 0x03, 0x85, 0x5a, 0x2c, 0x50, 0xd0, 0x05, 0x0a, 0xd2, 0x7e, 0x70,
	0x6b, 0xad, 0x92, 0x15, 0x7a, 0x69, 0x74, 0x49, 0x42, 0xce, 0x3d, 0xb9, 0x09, 0xb9, 0x2f, 0x59,
	0x11, 0x58, 0x67, 0xbc, 0xc9, 0x8a, 0x1e, 0xf0, 0x48, 0xed, 0x54, 0xfa, 0xfa, 0x41, 0x7b, 0x12,
	0xe4, 0x5e, 0x51, 0x2d, 0xd6, 0xc9, 0xfc, 0x9a, 0xa5, 0x51, 0x37, 0x5e, 0xb0, 0x43, 0xfc, 0x7d,
	0xde, 0x9b, 0x0f, 0x3a, 0x4b, 0xce, 0x93, 0x62, 0xbf, 0xec, 0x97, 0xf3, 0x27, 0xc6, 0x23, 0xf0,
	0x19, 0xad, 0x95, 0x5a, 0xf0, 0x11, 0x4b, 0xa3, 0x60, 0xc7, 0xa6, 0xc3, 0x63, 0xf8, 0x67, 0x7e,
	0x19, 0x07, 0x87, 0xef, 0x2c, 0xdb, 0xbe, 0x14, 0xbf, 0x61, 0xa7, 0xbd, 0xf1, 0x9b, 0x04, 0x3f,
	0x80, 0x88, 0xcb, 0x8e, 0xe0, 0xef, 0x9f, 0xf2, 0x9d, 0xbb, 0x9a, 0x5d, 0x1a, 0x27, 0x00, 0x2d,
	0x56, 0x35, 0xc1, 0xbc, 0x69, 0x97, 0xa8, 0x1a, 0xa9, 0x57, 0x95, 0x4f, 0xd0, 0xe4, 0x97, 0xc6,
	0x35, 0xb0, 0x51, 0x06, 0x9d, 0xb2, 0x71, 0xf2, 0x76, 0x12, 0xdc, 0x85, 0xd3, 0xce, 0x46, 0xe1,
	0xf6, 0xbd, 0x9b, 0x4d, 0x9a, 0x76, 0xba, 0x6e, 0xf0, 0x12, 0xe0, 0xf1, 0xca, 0x65, 0x65, 0xd4,
	0x6c, 0xaf, 0xb3, 0x7a, 0xf5, 0x13, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x61, 0x65, 0xfd, 0xb2, 0x01,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationRegisterServiceClient is the client API for ApplicationRegisterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationRegisterServiceClient interface {
	ApplicationCodeRegister(ctx context.Context, in *Application, opts ...grpc.CallOption) (*ApplicationMapping, error)
}

type applicationRegisterServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationRegisterServiceClient(cc *grpc.ClientConn) ApplicationRegisterServiceClient {
	return &applicationRegisterServiceClient{cc}
}

func (c *applicationRegisterServiceClient) ApplicationCodeRegister(ctx context.Context, in *Application, opts ...grpc.CallOption) (*ApplicationMapping, error) {
	out := new(ApplicationMapping)
	err := c.cc.Invoke(ctx, "/ApplicationRegisterService/applicationCodeRegister", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationRegisterServiceServer is the server API for ApplicationRegisterService service.
type ApplicationRegisterServiceServer interface {
	ApplicationCodeRegister(context.Context, *Application) (*ApplicationMapping, error)
}

// UnimplementedApplicationRegisterServiceServer can be embedded to have forward compatible implementations.
type UnimplementedApplicationRegisterServiceServer struct {
}

func (*UnimplementedApplicationRegisterServiceServer) ApplicationCodeRegister(ctx context.Context, req *Application) (*ApplicationMapping, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplicationCodeRegister not implemented")
}

func RegisterApplicationRegisterServiceServer(s *grpc.Server, srv ApplicationRegisterServiceServer) {
	s.RegisterService(&_ApplicationRegisterService_serviceDesc, srv)
}

func _ApplicationRegisterService_ApplicationCodeRegister_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Application)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationRegisterServiceServer).ApplicationCodeRegister(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ApplicationRegisterService/ApplicationCodeRegister",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationRegisterServiceServer).ApplicationCodeRegister(ctx, req.(*Application))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationRegisterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ApplicationRegisterService",
	HandlerType: (*ApplicationRegisterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "applicationCodeRegister",
			Handler:    _ApplicationRegisterService_ApplicationCodeRegister_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "language-agent/ApplicationRegisterService.proto",
}