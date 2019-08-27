// Code generated by protoc-gen-go. DO NOT EDIT.
// source: divide.proto

package divide

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

type Numbers struct {
	A                    int32    `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    int32    `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Numbers) Reset()         { *m = Numbers{} }
func (m *Numbers) String() string { return proto.CompactTextString(m) }
func (*Numbers) ProtoMessage()    {}
func (*Numbers) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e0e467ac0d16a7, []int{0}
}

func (m *Numbers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Numbers.Unmarshal(m, b)
}
func (m *Numbers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Numbers.Marshal(b, m, deterministic)
}
func (m *Numbers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Numbers.Merge(m, src)
}
func (m *Numbers) XXX_Size() int {
	return xxx_messageInfo_Numbers.Size(m)
}
func (m *Numbers) XXX_DiscardUnknown() {
	xxx_messageInfo_Numbers.DiscardUnknown(m)
}

var xxx_messageInfo_Numbers proto.InternalMessageInfo

func (m *Numbers) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Numbers) GetB() int32 {
	if m != nil {
		return m.B
	}
	return 0
}

type Resp struct {
	Q                    int32    `protobuf:"varint,1,opt,name=q,proto3" json:"q,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_68e0e467ac0d16a7, []int{1}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetQ() int32 {
	if m != nil {
		return m.Q
	}
	return 0
}

func init() {
	proto.RegisterType((*Numbers)(nil), "divide.Numbers")
	proto.RegisterType((*Resp)(nil), "divide.Resp")
}

func init() { proto.RegisterFile("divide.proto", fileDescriptor_68e0e467ac0d16a7) }

var fileDescriptor_68e0e467ac0d16a7 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xc9, 0x2c, 0xcb,
	0x4c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x94, 0x54, 0xb9, 0xd8,
	0xfd, 0x4a, 0x73, 0x93, 0x52, 0x8b, 0x8a, 0x85, 0x78, 0xb8, 0x18, 0x13, 0x25, 0x18, 0x15, 0x18,
	0x35, 0x58, 0x83, 0x18, 0x13, 0x41, 0xbc, 0x24, 0x09, 0x26, 0x08, 0x2f, 0x49, 0x49, 0x84, 0x8b,
	0x25, 0x28, 0xb5, 0xb8, 0x00, 0x24, 0x5a, 0x08, 0x53, 0x53, 0x68, 0x64, 0xca, 0xc5, 0xe6, 0x02,
	0x36, 0x46, 0x48, 0x9b, 0x8b, 0xc3, 0x25, 0x1f, 0xca, 0xe6, 0xd7, 0x83, 0xda, 0x04, 0x35, 0x58,
	0x8a, 0x07, 0x26, 0x00, 0x32, 0x42, 0x89, 0xc1, 0x49, 0x83, 0x4b, 0x34, 0x39, 0x3f, 0x57, 0x2f,
	0x3d, 0x3f, 0x2b, 0x35, 0x5b, 0xaf, 0x24, 0xb5, 0xb8, 0x04, 0x2a, 0xef, 0xc4, 0x0b, 0x31, 0x21,
	0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x35, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x58, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x92, 0x24, 0xab, 0x2e, 0xbc, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// DivideClient is the client API for Divide service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DivideClient interface {
	DoDivide(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Resp, error)
}

type divideClient struct {
	cc *grpc.ClientConn
}

func NewDivideClient(cc *grpc.ClientConn) DivideClient {
	return &divideClient{cc}
}

func (c *divideClient) DoDivide(ctx context.Context, in *Numbers, opts ...grpc.CallOption) (*Resp, error) {
	out := new(Resp)
	err := c.cc.Invoke(ctx, "/divide.Divide/DoDivide", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DivideServer is the server API for Divide service.
type DivideServer interface {
	DoDivide(context.Context, *Numbers) (*Resp, error)
}

// UnimplementedDivideServer can be embedded to have forward compatible implementations.
type UnimplementedDivideServer struct {
}

func (*UnimplementedDivideServer) DoDivide(ctx context.Context, req *Numbers) (*Resp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoDivide not implemented")
}

func RegisterDivideServer(s *grpc.Server, srv DivideServer) {
	s.RegisterService(&_Divide_serviceDesc, srv)
}

func _Divide_DoDivide_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Numbers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DivideServer).DoDivide(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/divide.Divide/DoDivide",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DivideServer).DoDivide(ctx, req.(*Numbers))
	}
	return interceptor(ctx, in, info, handler)
}

var _Divide_serviceDesc = grpc.ServiceDesc{
	ServiceName: "divide.Divide",
	HandlerType: (*DivideServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoDivide",
			Handler:    _Divide_DoDivide_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "divide.proto",
}
