// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secret.proto

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

type SecretRequest struct {
	Secrets              []*Secret        `protobuf:"bytes,1,rep,name=Secrets,proto3" json:"Secrets,omitempty"`
	OperationType        common.Operation `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SecretRequest) Reset()         { *m = SecretRequest{} }
func (m *SecretRequest) String() string { return proto.CompactTextString(m) }
func (*SecretRequest) ProtoMessage()    {}
func (*SecretRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acf428160d7a216, []int{0}
}

func (m *SecretRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretRequest.Unmarshal(m, b)
}
func (m *SecretRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretRequest.Marshal(b, m, deterministic)
}
func (m *SecretRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretRequest.Merge(m, src)
}
func (m *SecretRequest) XXX_Size() int {
	return xxx_messageInfo_SecretRequest.Size(m)
}
func (m *SecretRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SecretRequest proto.InternalMessageInfo

func (m *SecretRequest) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *SecretRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type SecretResponse struct {
	Secrets              []*Secret           `protobuf:"bytes,1,rep,name=Secrets,proto3" json:"Secrets,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SecretResponse) Reset()         { *m = SecretResponse{} }
func (m *SecretResponse) String() string { return proto.CompactTextString(m) }
func (*SecretResponse) ProtoMessage()    {}
func (*SecretResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acf428160d7a216, []int{1}
}

func (m *SecretResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SecretResponse.Unmarshal(m, b)
}
func (m *SecretResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SecretResponse.Marshal(b, m, deterministic)
}
func (m *SecretResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SecretResponse.Merge(m, src)
}
func (m *SecretResponse) XXX_Size() int {
	return xxx_messageInfo_SecretResponse.Size(m)
}
func (m *SecretResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SecretResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SecretResponse proto.InternalMessageInfo

func (m *SecretResponse) GetSecrets() []*Secret {
	if m != nil {
		return m.Secrets
	}
	return nil
}

func (m *SecretResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *SecretResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type Secret struct {
	Name                 string         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string         `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Filename             string         `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	Value                []byte         `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
	VaultId              string         `protobuf:"bytes,5,opt,name=vaultId,proto3" json:"vaultId,omitempty"`
	VaultName            string         `protobuf:"bytes,6,opt,name=vaultName,proto3" json:"vaultName,omitempty"`
	GroupName            string         `protobuf:"bytes,7,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Status               *common.Status `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	LocationName         string         `protobuf:"bytes,10,opt,name=locationName,proto3" json:"locationName,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_6acf428160d7a216, []int{2}
}

func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Secret) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Secret) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Secret) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Secret) GetVaultId() string {
	if m != nil {
		return m.VaultId
	}
	return ""
}

func (m *Secret) GetVaultName() string {
	if m != nil {
		return m.VaultName
	}
	return ""
}

func (m *Secret) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *Secret) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *Secret) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func init() {
	proto.RegisterType((*SecretRequest)(nil), "moc.cloudagent.security.SecretRequest")
	proto.RegisterType((*SecretResponse)(nil), "moc.cloudagent.security.SecretResponse")
	proto.RegisterType((*Secret)(nil), "moc.cloudagent.security.Secret")
}

func init() { proto.RegisterFile("secret.proto", fileDescriptor_6acf428160d7a216) }

var fileDescriptor_6acf428160d7a216 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x3d, 0x8f, 0xd3, 0x40,
	0x10, 0xc5, 0xb9, 0x3b, 0xe7, 0x32, 0xc9, 0xa5, 0x58, 0x21, 0x61, 0x59, 0x08, 0x22, 0x23, 0x41,
	0x1a, 0xd6, 0x92, 0x41, 0x08, 0x4a, 0x4e, 0xa2, 0xb8, 0x06, 0x24, 0x1f, 0xa2, 0x80, 0xca, 0x59,
	0x4f, 0x8c, 0xc1, 0xf6, 0x98, 0xfd, 0x08, 0xba, 0x8e, 0x3f, 0xc1, 0xcf, 0x45, 0x42, 0x9e, 0x8d,
	0x73, 0xba, 0x02, 0xa5, 0xa0, 0xdb, 0x79, 0xef, 0xed, 0x7b, 0xeb, 0x37, 0x86, 0x85, 0x41, 0xa5,
	0xd1, 0xca, 0x5e, 0x93, 0x25, 0xf1, 0xa0, 0x25, 0x25, 0x55, 0x43, 0xae, 0x2c, 0x2a, 0xec, 0xac,
	0x34, 0xa8, 0x9c, 0xae, 0xed, 0x4d, 0xfc, 0xa8, 0x22, 0xaa, 0x1a, 0x4c, 0x59, 0xb6, 0x71, 0xdb,
	0xf4, 0xa7, 0x2e, 0xfa, 0x1e, 0xb5, 0xf1, 0x17, 0xe3, 0x85, 0xa2, 0xb6, 0xa5, 0xce, 0x4f, 0xc9,
	0xaf, 0x00, 0x2e, 0xae, 0xd9, 0x37, 0xc7, 0x1f, 0x0e, 0x8d, 0x15, 0x6f, 0x60, 0xea, 0x01, 0x13,
	0x05, 0xab, 0x93, 0xf5, 0x3c, 0x7b, 0x2c, 0xff, 0x11, 0x25, 0xf7, 0x17, 0x47, 0xbd, 0x78, 0x09,
	0x17, 0x1f, 0x7a, 0xd4, 0x85, 0xad, 0xa9, 0xfb, 0x78, 0xd3, 0x63, 0x34, 0x59, 0x05, 0xeb, 0x65,
	0xb6, 0x64, 0x83, 0x03, 0x93, 0xdf, 0x15, 0x25, 0xbf, 0x03, 0x58, 0x8e, 0x4f, 0x30, 0x3d, 0x75,
	0x06, 0xff, 0xe7, 0x0d, 0x19, 0x84, 0x39, 0x1a, 0xd7, 0x58, 0x0e, 0x9f, 0x67, 0xb1, 0xf4, 0x7d,
	0xc8, 0xb1, 0x0f, 0x79, 0x49, 0xd4, 0x7c, 0x2a, 0x1a, 0x87, 0xf9, 0x5e, 0x29, 0xee, 0xc3, 0xd9,
	0x3b, 0xad, 0x49, 0x47, 0x27, 0xab, 0x60, 0x3d, 0xcb, 0xfd, 0x90, 0xfc, 0x09, 0x20, 0xf4, 0xae,
	0x42, 0xc0, 0x69, 0x57, 0xb4, 0x18, 0x05, 0xcc, 0xf3, 0x59, 0x2c, 0x61, 0x52, 0x97, 0x1c, 0x32,
	0xcb, 0x27, 0x75, 0x29, 0x62, 0x38, 0xdf, 0xd6, 0x0d, 0xb2, 0xce, 0xfb, 0x1c, 0xe6, 0x21, 0x60,
	0x37, 0x24, 0x46, 0xa7, 0xab, 0x60, 0xbd, 0xc8, 0xfd, 0x20, 0x22, 0x98, 0xee, 0x0a, 0xd7, 0xd8,
	0xab, 0x32, 0x3a, 0xe3, 0x0b, 0xe3, 0x28, 0x1e, 0xc2, 0x8c, 0x8f, 0xef, 0x07, 0xb3, 0x90, 0xb9,
	0x5b, 0x60, 0x60, 0x2b, 0x4d, 0xae, 0x67, 0x76, 0xea, 0xd9, 0x03, 0x20, 0x9e, 0x40, 0x68, 0x6c,
	0x61, 0x9d, 0x89, 0xce, 0xb9, 0x80, 0x39, 0x57, 0x77, 0xcd, 0x50, 0xbe, 0xa7, 0x44, 0x02, 0x8b,
	0x86, 0x14, 0xef, 0x80, 0x5d, 0x80, 0x5d, 0xee, 0x60, 0xd9, 0x37, 0x98, 0xfb, 0xcf, 0x7f, 0x3b,
	0x34, 0x2e, 0xbe, 0x40, 0x78, 0xd5, 0xed, 0xe8, 0x3b, 0x8a, 0xa7, 0xc7, 0x96, 0xe1, 0xff, 0xa4,
	0xf8, 0xd9, 0x51, 0x9d, 0x5f, 0x77, 0x72, 0xef, 0xf2, 0xf5, 0xe7, 0x57, 0x55, 0x6d, 0xbf, 0xba,
	0x8d, 0x54, 0xd4, 0xa6, 0x6d, 0xad, 0x34, 0x19, 0xda, 0xda, 0xb4, 0x25, 0xf5, 0x9c, 0x17, 0x97,
	0xea, 0x5e, 0xa5, 0xb7, 0x56, 0xe9, 0x68, 0xb5, 0x09, 0x99, 0x7e, 0xf1, 0x37, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0xcf, 0x6d, 0xf4, 0x1e, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SecretAgentClient is the client API for SecretAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SecretAgentClient interface {
	Invoke(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*SecretResponse, error)
}

type secretAgentClient struct {
	cc *grpc.ClientConn
}

func NewSecretAgentClient(cc *grpc.ClientConn) SecretAgentClient {
	return &secretAgentClient{cc}
}

func (c *secretAgentClient) Invoke(ctx context.Context, in *SecretRequest, opts ...grpc.CallOption) (*SecretResponse, error) {
	out := new(SecretResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.security.SecretAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SecretAgentServer is the server API for SecretAgent service.
type SecretAgentServer interface {
	Invoke(context.Context, *SecretRequest) (*SecretResponse, error)
}

// UnimplementedSecretAgentServer can be embedded to have forward compatible implementations.
type UnimplementedSecretAgentServer struct {
}

func (*UnimplementedSecretAgentServer) Invoke(ctx context.Context, req *SecretRequest) (*SecretResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterSecretAgentServer(s *grpc.Server, srv SecretAgentServer) {
	s.RegisterService(&_SecretAgent_serviceDesc, srv)
}

func _SecretAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SecretRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.security.SecretAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretAgentServer).Invoke(ctx, req.(*SecretRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.security.SecretAgent",
	HandlerType: (*SecretAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _SecretAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret.proto",
}
