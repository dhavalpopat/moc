// Code generated by protoc-gen-go. DO NOT EDIT.
// source: certificate.proto

package security

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc-proto/rpc/common"
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

type CertificateType int32

const (
	CertificateType_Client CertificateType = 0
	CertificateType_Server CertificateType = 1
)

var CertificateType_name = map[int32]string{
	0: "Client",
	1: "Server",
}

var CertificateType_value = map[string]int32{
	"Client": 0,
	"Server": 1,
}

func (x CertificateType) String() string {
	return proto.EnumName(CertificateType_name, int32(x))
}

func (CertificateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{0}
}

type CertificateRequest struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CertificateRequest) Reset()         { *m = CertificateRequest{} }
func (m *CertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateRequest) ProtoMessage()    {}
func (*CertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{0}
}

func (m *CertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateRequest.Unmarshal(m, b)
}
func (m *CertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateRequest.Marshal(b, m, deterministic)
}
func (m *CertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateRequest.Merge(m, src)
}
func (m *CertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateRequest.Size(m)
}
func (m *CertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateRequest proto.InternalMessageInfo

func (m *CertificateRequest) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type CertificateResponse struct {
	Certificates         []*Certificate      `protobuf:"bytes,1,rep,name=Certificates,proto3" json:"Certificates,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CertificateResponse) Reset()         { *m = CertificateResponse{} }
func (m *CertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CertificateResponse) ProtoMessage()    {}
func (*CertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{1}
}

func (m *CertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateResponse.Unmarshal(m, b)
}
func (m *CertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateResponse.Marshal(b, m, deterministic)
}
func (m *CertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateResponse.Merge(m, src)
}
func (m *CertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CertificateResponse.Size(m)
}
func (m *CertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateResponse proto.InternalMessageInfo

func (m *CertificateResponse) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

func (m *CertificateResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *CertificateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Certificate struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string          `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	NotBefore            int64           `protobuf:"varint,3,opt,name=notBefore,proto3" json:"notBefore,omitempty"`
	NotAfter             int64           `protobuf:"varint,4,opt,name=notAfter,proto3" json:"notAfter,omitempty"`
	Certificate          []byte          `protobuf:"bytes,5,opt,name=certificate,proto3" json:"certificate,omitempty"`
	Status               *common.Status  `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Type                 CertificateType `protobuf:"varint,7,opt,name=type,proto3,enum=moc.nodeagent.security.CertificateType" json:"type,omitempty"`
	Entity               *common.Entity  `protobuf:"bytes,8,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_c0d34c34dd33be4b, []int{2}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Certificate) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Certificate) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *Certificate) GetNotAfter() int64 {
	if m != nil {
		return m.NotAfter
	}
	return 0
}

func (m *Certificate) GetCertificate() []byte {
	if m != nil {
		return m.Certificate
	}
	return nil
}

func (m *Certificate) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Certificate) GetType() CertificateType {
	if m != nil {
		return m.Type
	}
	return CertificateType_Client
}

func (m *Certificate) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.nodeagent.security.CertificateType", CertificateType_name, CertificateType_value)
	proto.RegisterType((*CertificateRequest)(nil), "moc.nodeagent.security.CertificateRequest")
	proto.RegisterType((*CertificateResponse)(nil), "moc.nodeagent.security.CertificateResponse")
	proto.RegisterType((*Certificate)(nil), "moc.nodeagent.security.Certificate")
}

func init() { proto.RegisterFile("certificate.proto", fileDescriptor_c0d34c34dd33be4b) }

var fileDescriptor_c0d34c34dd33be4b = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x93, 0xd4, 0x34, 0x93, 0x28, 0x84, 0x05, 0x21, 0xcb, 0x42, 0xc8, 0x4a, 0x0f, 0x98,
	0x22, 0x6c, 0xc9, 0x08, 0x71, 0xe0, 0xd4, 0x84, 0xaa, 0x47, 0xa4, 0x2d, 0x70, 0x40, 0x42, 0xc2,
	0xd9, 0x8c, 0x83, 0x85, 0xed, 0x31, 0xbb, 0x63, 0x50, 0xfe, 0x83, 0x8f, 0xe0, 0x4b, 0xf8, 0x2e,
	0x94, 0x75, 0xda, 0xba, 0x88, 0x43, 0x0e, 0xf4, 0xb6, 0xf3, 0xe6, 0xcd, 0x7b, 0xfb, 0x76, 0x07,
	0xee, 0x29, 0xd4, 0x9c, 0x67, 0xb9, 0x4a, 0x19, 0xa3, 0x5a, 0x13, 0x93, 0x78, 0x58, 0x92, 0x8a,
	0x2a, 0x5a, 0x61, 0xba, 0xc6, 0x8a, 0x23, 0x83, 0xaa, 0xd1, 0x39, 0x6f, 0xfc, 0xc7, 0x6b, 0xa2,
	0x75, 0x81, 0xb1, 0x65, 0x2d, 0x9b, 0x2c, 0xfe, 0xa1, 0xd3, 0xba, 0x46, 0x6d, 0xda, 0x39, 0x7f,
	0xac, 0xa8, 0x2c, 0xa9, 0x6a, 0xab, 0xd9, 0x27, 0x10, 0x8b, 0x6b, 0x69, 0x89, 0xdf, 0x1a, 0x34,
	0x2c, 0xce, 0x61, 0xdc, 0x41, 0x8d, 0xe7, 0x04, 0xfd, 0x70, 0x94, 0x1c, 0x47, 0xff, 0xb6, 0x8c,
	0xba, 0x0a, 0x37, 0x06, 0x67, 0xbf, 0x1c, 0xb8, 0x7f, 0x43, 0xdf, 0xd4, 0x54, 0x19, 0xfc, 0x6f,
	0x06, 0x22, 0x01, 0x57, 0xa2, 0x69, 0x0a, 0xf6, 0x7a, 0x81, 0x13, 0x8e, 0x12, 0x3f, 0x6a, 0xe3,
	0x47, 0x97, 0xf1, 0xa3, 0x39, 0x51, 0xf1, 0x21, 0x2d, 0x1a, 0x94, 0x3b, 0xa6, 0x78, 0x00, 0x87,
	0x67, 0x5a, 0x93, 0xf6, 0xfa, 0x81, 0x13, 0x0e, 0x65, 0x5b, 0xcc, 0x7e, 0xf6, 0x60, 0xd4, 0x91,
	0x16, 0x02, 0x06, 0x55, 0x5a, 0xa2, 0xe7, 0x58, 0x92, 0x3d, 0x8b, 0x09, 0xf4, 0xf2, 0x95, 0x75,
	0x1a, 0xca, 0x5e, 0xbe, 0x12, 0x8f, 0x60, 0x58, 0x11, 0xcf, 0x31, 0x23, 0x8d, 0x56, 0xad, 0x2f,
	0xaf, 0x01, 0xe1, 0xc3, 0x51, 0x45, 0x7c, 0x9a, 0x31, 0x6a, 0x6f, 0x60, 0x9b, 0x57, 0xb5, 0x08,
	0x60, 0xd4, 0xf9, 0x52, 0xef, 0x30, 0x70, 0xc2, 0xb1, 0xec, 0x42, 0xe2, 0x18, 0x5c, 0xc3, 0x29,
	0x37, 0xc6, 0x73, 0x6d, 0xb2, 0x91, 0x7d, 0x9c, 0x0b, 0x0b, 0xc9, 0x5d, 0x4b, 0xbc, 0x86, 0x01,
	0x6f, 0x6a, 0xf4, 0xee, 0x04, 0x4e, 0x38, 0x49, 0x9e, 0xec, 0xf1, 0x7e, 0xef, 0x36, 0x35, 0x4a,
	0x3b, 0xb4, 0x75, 0xc0, 0x8a, 0x73, 0xde, 0x78, 0x47, 0x1d, 0x87, 0x33, 0x0b, 0xc9, 0x5d, 0xeb,
	0xe4, 0x29, 0xdc, 0xfd, 0x6b, 0x5a, 0x00, 0xb8, 0x8b, 0x22, 0xc7, 0x8a, 0xa7, 0x07, 0xdb, 0xf3,
	0x05, 0xea, 0xef, 0xa8, 0xa7, 0x4e, 0xf2, 0xbb, 0x07, 0xd3, 0x0e, 0xf7, 0x74, 0x7b, 0x07, 0xf1,
	0x15, 0x26, 0x0b, 0x8d, 0x29, 0xe3, 0x5b, 0xfd, 0xbe, 0x5e, 0x6d, 0x83, 0x9d, 0xec, 0xf3, 0xcb,
	0xed, 0x22, 0xfa, 0xcf, 0xf6, 0xe2, 0xb6, 0x4b, 0x35, 0x3b, 0x10, 0x9f, 0xa1, 0x7f, 0x8e, 0x7c,
	0x9b, 0x0e, 0x0a, 0xdc, 0x37, 0x58, 0xe0, 0xad, 0xc6, 0x98, 0xbf, 0xfa, 0xf8, 0x72, 0x9d, 0xf3,
	0x97, 0x66, 0x19, 0x29, 0x2a, 0xe3, 0x32, 0x57, 0x9a, 0x0c, 0x65, 0x1c, 0x97, 0xa4, 0x9e, 0xdb,
	0xbd, 0x8e, 0x75, 0xad, 0xe2, 0x2b, 0xb9, 0xf8, 0x52, 0x6e, 0xe9, 0xda, 0xee, 0x8b, 0x3f, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x02, 0xc3, 0xb5, 0x3d, 0x2f, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertificateAgentClient is the client API for CertificateAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAgentClient interface {
	CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
	Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error)
}

type certificateAgentClient struct {
	cc *grpc.ClientConn
}

func NewCertificateAgentClient(cc *grpc.ClientConn) CertificateAgentClient {
	return &certificateAgentClient{cc}
}

func (c *certificateAgentClient) CreateOrUpdate(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/CreateOrUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Get(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAgentClient) Delete(ctx context.Context, in *CertificateRequest, opts ...grpc.CallOption) (*CertificateResponse, error) {
	out := new(CertificateResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.security.CertificateAgent/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAgentServer is the server API for CertificateAgent service.
type CertificateAgentServer interface {
	CreateOrUpdate(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Get(context.Context, *CertificateRequest) (*CertificateResponse, error)
	Delete(context.Context, *CertificateRequest) (*CertificateResponse, error)
}

// UnimplementedCertificateAgentServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAgentServer struct {
}

func (*UnimplementedCertificateAgentServer) CreateOrUpdate(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdate not implemented")
}
func (*UnimplementedCertificateAgentServer) Get(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedCertificateAgentServer) Delete(ctx context.Context, req *CertificateRequest) (*CertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

func RegisterCertificateAgentServer(s *grpc.Server, srv CertificateAgentServer) {
	s.RegisterService(&_CertificateAgent_serviceDesc, srv)
}

func _CertificateAgent_CreateOrUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.CertificateAgent/CreateOrUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).CreateOrUpdate(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.CertificateAgent/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Get(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAgent_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAgentServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.security.CertificateAgent/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAgentServer).Delete(ctx, req.(*CertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.security.CertificateAgent",
	HandlerType: (*CertificateAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdate",
			Handler:    _CertificateAgent_CreateOrUpdate_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _CertificateAgent_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _CertificateAgent_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "certificate.proto",
}
