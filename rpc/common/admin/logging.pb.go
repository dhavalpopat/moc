// Code generated by protoc-gen-go. DO NOT EDIT.
// source: admin/logging/logging.proto

package admin

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type LogRotateRequest struct {
	LogRotation          *LogRotation `protobuf:"bytes,1,opt,name=logRotation,proto3" json:"logRotation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *LogRotateRequest) Reset()         { *m = LogRotateRequest{} }
func (m *LogRotateRequest) String() string { return proto.CompactTextString(m) }
func (*LogRotateRequest) ProtoMessage()    {}
func (*LogRotateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_343ba182a88f29f3, []int{0}
}

func (m *LogRotateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRotateRequest.Unmarshal(m, b)
}
func (m *LogRotateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRotateRequest.Marshal(b, m, deterministic)
}
func (m *LogRotateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRotateRequest.Merge(m, src)
}
func (m *LogRotateRequest) XXX_Size() int {
	return xxx_messageInfo_LogRotateRequest.Size(m)
}
func (m *LogRotateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRotateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRotateRequest proto.InternalMessageInfo

func (m *LogRotateRequest) GetLogRotation() *LogRotation {
	if m != nil {
		return m.LogRotation
	}
	return nil
}

type LogRotateResponse struct {
	Result               bool     `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRotateResponse) Reset()         { *m = LogRotateResponse{} }
func (m *LogRotateResponse) String() string { return proto.CompactTextString(m) }
func (*LogRotateResponse) ProtoMessage()    {}
func (*LogRotateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_343ba182a88f29f3, []int{1}
}

func (m *LogRotateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRotateResponse.Unmarshal(m, b)
}
func (m *LogRotateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRotateResponse.Marshal(b, m, deterministic)
}
func (m *LogRotateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRotateResponse.Merge(m, src)
}
func (m *LogRotateResponse) XXX_Size() int {
	return xxx_messageInfo_LogRotateResponse.Size(m)
}
func (m *LogRotateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRotateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogRotateResponse proto.InternalMessageInfo

func (m *LogRotateResponse) GetResult() bool {
	if m != nil {
		return m.Result
	}
	return false
}

func (m *LogRotateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type LogRotation struct {
	Minutes              int32    `protobuf:"varint,1,opt,name=minutes,proto3" json:"minutes,omitempty"`
	Size                 int32    `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	EnableTime           bool     `protobuf:"varint,3,opt,name=enableTime,proto3" json:"enableTime,omitempty"`
	EnableSize           bool     `protobuf:"varint,4,opt,name=enableSize,proto3" json:"enableSize,omitempty"`
	DisableTime          bool     `protobuf:"varint,5,opt,name=disableTime,proto3" json:"disableTime,omitempty"`
	DisableSize          bool     `protobuf:"varint,6,opt,name=disableSize,proto3" json:"disableSize,omitempty"`
	Limit                int32    `protobuf:"varint,7,opt,name=limit,proto3" json:"limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRotation) Reset()         { *m = LogRotation{} }
func (m *LogRotation) String() string { return proto.CompactTextString(m) }
func (*LogRotation) ProtoMessage()    {}
func (*LogRotation) Descriptor() ([]byte, []int) {
	return fileDescriptor_343ba182a88f29f3, []int{2}
}

func (m *LogRotation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRotation.Unmarshal(m, b)
}
func (m *LogRotation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRotation.Marshal(b, m, deterministic)
}
func (m *LogRotation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRotation.Merge(m, src)
}
func (m *LogRotation) XXX_Size() int {
	return xxx_messageInfo_LogRotation.Size(m)
}
func (m *LogRotation) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRotation.DiscardUnknown(m)
}

var xxx_messageInfo_LogRotation proto.InternalMessageInfo

func (m *LogRotation) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *LogRotation) GetSize() int32 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *LogRotation) GetEnableTime() bool {
	if m != nil {
		return m.EnableTime
	}
	return false
}

func (m *LogRotation) GetEnableSize() bool {
	if m != nil {
		return m.EnableSize
	}
	return false
}

func (m *LogRotation) GetDisableTime() bool {
	if m != nil {
		return m.DisableTime
	}
	return false
}

func (m *LogRotation) GetDisableSize() bool {
	if m != nil {
		return m.DisableSize
	}
	return false
}

func (m *LogRotation) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type LogRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_343ba182a88f29f3, []int{3}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

type LogFileResponse struct {
	File                 []byte              `protobuf:"bytes,1,opt,name=file,proto3" json:"file,omitempty"`
	Done                 *wrappers.BoolValue `protobuf:"bytes,2,opt,name=done,proto3" json:"done,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *LogFileResponse) Reset()         { *m = LogFileResponse{} }
func (m *LogFileResponse) String() string { return proto.CompactTextString(m) }
func (*LogFileResponse) ProtoMessage()    {}
func (*LogFileResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_343ba182a88f29f3, []int{4}
}

func (m *LogFileResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogFileResponse.Unmarshal(m, b)
}
func (m *LogFileResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogFileResponse.Marshal(b, m, deterministic)
}
func (m *LogFileResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogFileResponse.Merge(m, src)
}
func (m *LogFileResponse) XXX_Size() int {
	return xxx_messageInfo_LogFileResponse.Size(m)
}
func (m *LogFileResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogFileResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogFileResponse proto.InternalMessageInfo

func (m *LogFileResponse) GetFile() []byte {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *LogFileResponse) GetDone() *wrappers.BoolValue {
	if m != nil {
		return m.Done
	}
	return nil
}

func (m *LogFileResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*LogRotateRequest)(nil), "moc.common.admin.LogRotateRequest")
	proto.RegisterType((*LogRotateResponse)(nil), "moc.common.admin.LogRotateResponse")
	proto.RegisterType((*LogRotation)(nil), "moc.common.admin.LogRotation")
	proto.RegisterType((*LogRequest)(nil), "moc.common.admin.LogRequest")
	proto.RegisterType((*LogFileResponse)(nil), "moc.common.admin.LogFileResponse")
}

func init() { proto.RegisterFile("admin/logging/logging.proto", fileDescriptor_343ba182a88f29f3) }

var fileDescriptor_343ba182a88f29f3 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcd, 0x6e, 0xd4, 0x30,
	0x10, 0xc7, 0x09, 0xdd, 0x6c, 0xcb, 0xa4, 0x12, 0xc5, 0xaa, 0x50, 0xb4, 0x40, 0xb5, 0x84, 0x4b,
	0x2f, 0xd8, 0xb0, 0x3c, 0x00, 0x6a, 0x25, 0x40, 0x42, 0x7b, 0xf2, 0x22, 0x0e, 0xdc, 0xb2, 0x59,
	0xaf, 0xb1, 0xb0, 0x3d, 0x21, 0x76, 0x84, 0xc4, 0x2b, 0xf1, 0x3c, 0xbc, 0x0f, 0xca, 0xb8, 0x61,
	0x03, 0x2a, 0x9c, 0xec, 0x99, 0xf9, 0xcf, 0xcf, 0xf3, 0x61, 0x78, 0x54, 0xef, 0x9c, 0xf1, 0xc2,
	0xa2, 0xd6, 0xc6, 0xeb, 0xf1, 0xe4, 0x6d, 0x87, 0x11, 0xd9, 0x99, 0xc3, 0x86, 0x37, 0xe8, 0x1c,
	0x7a, 0x4e, 0xba, 0xc5, 0x85, 0x46, 0xd4, 0x56, 0x09, 0x8a, 0x6f, 0xfb, 0xbd, 0xf8, 0xd6, 0xd5,
	0x6d, 0xab, 0xba, 0x90, 0x32, 0xaa, 0x0d, 0x9c, 0xad, 0x51, 0x4b, 0x8c, 0x75, 0x54, 0x52, 0x7d,
	0xed, 0x55, 0x88, 0xec, 0x35, 0x14, 0xf6, 0xc6, 0x67, 0xd0, 0x97, 0xd9, 0x32, 0xbb, 0x2c, 0x56,
	0x4f, 0xf8, 0xdf, 0x6c, 0xbe, 0x3e, 0x88, 0xe4, 0x34, 0xa3, 0xba, 0x82, 0x07, 0x13, 0x68, 0x68,
	0xd1, 0x07, 0xc5, 0x1e, 0xc2, 0x5c, 0xaa, 0xd0, 0xdb, 0x48, 0xc0, 0x13, 0x79, 0x63, 0xb1, 0x73,
	0xc8, 0xdf, 0x74, 0x1d, 0x76, 0xe5, 0xdd, 0x65, 0x76, 0x79, 0x4f, 0x26, 0xa3, 0xfa, 0x99, 0x41,
	0x31, 0xe1, 0xb3, 0x12, 0x8e, 0x9d, 0xf1, 0x7d, 0x54, 0x81, 0xd2, 0x73, 0x39, 0x9a, 0x8c, 0xc1,
	0x2c, 0x98, 0xef, 0x8a, 0xd2, 0x73, 0x49, 0x77, 0x76, 0x01, 0xa0, 0x7c, 0xbd, 0xb5, 0xea, 0x83,
	0x71, 0xaa, 0x3c, 0xa2, 0xf7, 0x26, 0x9e, 0x43, 0x7c, 0x33, 0x64, 0xce, 0xa6, 0xf1, 0xc1, 0xc3,
	0x96, 0x50, 0xec, 0x4c, 0xf8, 0x0d, 0xc8, 0x49, 0x30, 0x75, 0x4d, 0x14, 0x84, 0x98, 0xff, 0xa1,
	0x20, 0xc6, 0x39, 0xe4, 0xd6, 0x38, 0x13, 0xcb, 0x63, 0x2a, 0x2c, 0x19, 0xd5, 0x29, 0xc0, 0xd0,
	0x56, 0x9a, 0x74, 0xf5, 0x05, 0xee, 0xaf, 0x51, 0xbf, 0x35, 0xf6, 0x30, 0x26, 0x06, 0xb3, 0xbd,
	0xb1, 0x8a, 0xba, 0x3c, 0x95, 0x74, 0x67, 0x1c, 0x66, 0x3b, 0xf4, 0xa9, 0xc5, 0x62, 0xb5, 0xe0,
	0x69, 0xa7, 0x7c, 0xdc, 0x29, 0xbf, 0x46, 0xb4, 0x1f, 0x6b, 0xdb, 0x2b, 0x49, 0xba, 0xe1, 0x69,
	0x45, 0x23, 0x3d, 0x4a, 0x23, 0x25, 0x63, 0xf5, 0x23, 0x83, 0x93, 0x35, 0xea, 0x2b, 0xad, 0x7c,
	0x64, 0xef, 0xe1, 0xe8, 0x9d, 0x8a, 0xec, 0xf1, 0xed, 0x5b, 0x4d, 0xe5, 0x2d, 0x9e, 0xde, 0x1a,
	0x9d, 0x96, 0x5b, 0xdd, 0x79, 0x91, 0xb1, 0x0d, 0xcc, 0xd3, 0xae, 0x59, 0xf5, 0xef, 0x4f, 0x32,
	0xfe, 0xae, 0xc5, 0xb3, 0xff, 0x6a, 0x46, 0xec, 0xf5, 0xcb, 0x4f, 0x42, 0x9b, 0xf8, 0xb9, 0xdf,
	0x0e, 0x52, 0xe1, 0x4c, 0xd3, 0x61, 0xc0, 0x7d, 0x14, 0x0e, 0x9b, 0xe7, 0xd4, 0xb8, 0xe8, 0xda,
	0x46, 0x24, 0x8c, 0x20, 0xcc, 0x76, 0x4e, 0xfe, 0x57, 0xbf, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0e,
	0xc0, 0xbc, 0x7b, 0x23, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LogAgentClient is the client API for LogAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LogAgentClient interface {
	Get(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (LogAgent_GetClient, error)
	Rotate(ctx context.Context, in *LogRotateRequest, opts ...grpc.CallOption) (*LogRotateResponse, error)
}

type logAgentClient struct {
	cc *grpc.ClientConn
}

func NewLogAgentClient(cc *grpc.ClientConn) LogAgentClient {
	return &logAgentClient{cc}
}

func (c *logAgentClient) Get(ctx context.Context, in *LogRequest, opts ...grpc.CallOption) (LogAgent_GetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LogAgent_serviceDesc.Streams[0], "/moc.common.admin.LogAgent/Get", opts...)
	if err != nil {
		return nil, err
	}
	x := &logAgentGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LogAgent_GetClient interface {
	Recv() (*LogFileResponse, error)
	grpc.ClientStream
}

type logAgentGetClient struct {
	grpc.ClientStream
}

func (x *logAgentGetClient) Recv() (*LogFileResponse, error) {
	m := new(LogFileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *logAgentClient) Rotate(ctx context.Context, in *LogRotateRequest, opts ...grpc.CallOption) (*LogRotateResponse, error) {
	out := new(LogRotateResponse)
	err := c.cc.Invoke(ctx, "/moc.common.admin.LogAgent/Rotate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogAgentServer is the server API for LogAgent service.
type LogAgentServer interface {
	Get(*LogRequest, LogAgent_GetServer) error
	Rotate(context.Context, *LogRotateRequest) (*LogRotateResponse, error)
}

// UnimplementedLogAgentServer can be embedded to have forward compatible implementations.
type UnimplementedLogAgentServer struct {
}

func (*UnimplementedLogAgentServer) Get(req *LogRequest, srv LogAgent_GetServer) error {
	return status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedLogAgentServer) Rotate(ctx context.Context, req *LogRotateRequest) (*LogRotateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Rotate not implemented")
}

func RegisterLogAgentServer(s *grpc.Server, srv LogAgentServer) {
	s.RegisterService(&_LogAgent_serviceDesc, srv)
}

func _LogAgent_Get_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LogRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LogAgentServer).Get(m, &logAgentGetServer{stream})
}

type LogAgent_GetServer interface {
	Send(*LogFileResponse) error
	grpc.ServerStream
}

type logAgentGetServer struct {
	grpc.ServerStream
}

func (x *logAgentGetServer) Send(m *LogFileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _LogAgent_Rotate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LogRotateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogAgentServer).Rotate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.common.admin.LogAgent/Rotate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogAgentServer).Rotate(ctx, req.(*LogRotateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LogAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.common.admin.LogAgent",
	HandlerType: (*LogAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Rotate",
			Handler:    _LogAgent_Rotate_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Get",
			Handler:       _LogAgent_Get_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "admin/logging/logging.proto",
}
