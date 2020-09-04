// Code generated by protoc-gen-go. DO NOT EDIT.
// source: virtualharddisk.proto

package storage

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

type VirtualHardDiskType int32

const (
	VirtualHardDiskType_OS_VIRTUALHARDDISK       VirtualHardDiskType = 0
	VirtualHardDiskType_DATADISK_VIRTUALHARDDISK VirtualHardDiskType = 1
)

var VirtualHardDiskType_name = map[int32]string{
	0: "OS_VIRTUALHARDDISK",
	1: "DATADISK_VIRTUALHARDDISK",
}

var VirtualHardDiskType_value = map[string]int32{
	"OS_VIRTUALHARDDISK":       0,
	"DATADISK_VIRTUALHARDDISK": 1,
}

func (x VirtualHardDiskType) String() string {
	return proto.EnumName(VirtualHardDiskType_name, int32(x))
}

func (VirtualHardDiskType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{0}
}

type VirtualHardDiskRequest struct {
	VirtualHardDiskSystems []*VirtualHardDisk `protobuf:"bytes,1,rep,name=VirtualHardDiskSystems,proto3" json:"VirtualHardDiskSystems,omitempty"`
	OperationType          common.Operation   `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}           `json:"-"`
	XXX_unrecognized       []byte             `json:"-"`
	XXX_sizecache          int32              `json:"-"`
}

func (m *VirtualHardDiskRequest) Reset()         { *m = VirtualHardDiskRequest{} }
func (m *VirtualHardDiskRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskRequest) ProtoMessage()    {}
func (*VirtualHardDiskRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{0}
}

func (m *VirtualHardDiskRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskRequest.Unmarshal(m, b)
}
func (m *VirtualHardDiskRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskRequest.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskRequest.Merge(m, src)
}
func (m *VirtualHardDiskRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskRequest.Size(m)
}
func (m *VirtualHardDiskRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskRequest proto.InternalMessageInfo

func (m *VirtualHardDiskRequest) GetVirtualHardDiskSystems() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDiskSystems
	}
	return nil
}

func (m *VirtualHardDiskRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualHardDiskResponse struct {
	VirtualHardDiskSystems []*VirtualHardDisk  `protobuf:"bytes,1,rep,name=VirtualHardDiskSystems,proto3" json:"VirtualHardDiskSystems,omitempty"`
	Result                 *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                  string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}            `json:"-"`
	XXX_unrecognized       []byte              `json:"-"`
	XXX_sizecache          int32               `json:"-"`
}

func (m *VirtualHardDiskResponse) Reset()         { *m = VirtualHardDiskResponse{} }
func (m *VirtualHardDiskResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDiskResponse) ProtoMessage()    {}
func (*VirtualHardDiskResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{1}
}

func (m *VirtualHardDiskResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDiskResponse.Unmarshal(m, b)
}
func (m *VirtualHardDiskResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDiskResponse.Marshal(b, m, deterministic)
}
func (m *VirtualHardDiskResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDiskResponse.Merge(m, src)
}
func (m *VirtualHardDiskResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDiskResponse.Size(m)
}
func (m *VirtualHardDiskResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDiskResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDiskResponse proto.InternalMessageInfo

func (m *VirtualHardDiskResponse) GetVirtualHardDiskSystems() []*VirtualHardDisk {
	if m != nil {
		return m.VirtualHardDiskSystems
	}
	return nil
}

func (m *VirtualHardDiskResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualHardDiskResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VirtualHardDisk struct {
	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id     string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Source string `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Path   string `protobuf:"bytes,4,opt,name=path,proto3" json:"path,omitempty"`
	// Storage container name to hold this vhd
	ContainerName        string              `protobuf:"bytes,5,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Status               *common.Status      `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Size                 int64               `protobuf:"varint,7,opt,name=size,proto3" json:"size,omitempty"`
	Dynamic              bool                `protobuf:"varint,8,opt,name=dynamic,proto3" json:"dynamic,omitempty"`
	Blocksizebytes       int32               `protobuf:"varint,9,opt,name=blocksizebytes,proto3" json:"blocksizebytes,omitempty"`
	Logicalsectorbytes   int32               `protobuf:"varint,10,opt,name=logicalsectorbytes,proto3" json:"logicalsectorbytes,omitempty"`
	Physicalsectorbytes  int32               `protobuf:"varint,11,opt,name=physicalsectorbytes,proto3" json:"physicalsectorbytes,omitempty"`
	Controllernumber     int64               `protobuf:"varint,12,opt,name=controllernumber,proto3" json:"controllernumber,omitempty"`
	Controllerlocation   int64               `protobuf:"varint,13,opt,name=controllerlocation,proto3" json:"controllerlocation,omitempty"`
	Disknumber           int64               `protobuf:"varint,14,opt,name=disknumber,proto3" json:"disknumber,omitempty"`
	VirtualmachineName   string              `protobuf:"bytes,15,opt,name=virtualmachineName,proto3" json:"virtualmachineName,omitempty"`
	Scsipath             string              `protobuf:"bytes,16,opt,name=scsipath,proto3" json:"scsipath,omitempty"`
	Virtualharddisktype  VirtualHardDiskType `protobuf:"varint,17,opt,name=virtualharddisktype,proto3,enum=moc.nodeagent.storage.VirtualHardDiskType" json:"virtualharddisktype,omitempty"`
	Entity               *common.Entity      `protobuf:"bytes,18,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VirtualHardDisk) Reset()         { *m = VirtualHardDisk{} }
func (m *VirtualHardDisk) String() string { return proto.CompactTextString(m) }
func (*VirtualHardDisk) ProtoMessage()    {}
func (*VirtualHardDisk) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4b382f86170a6e5, []int{2}
}

func (m *VirtualHardDisk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualHardDisk.Unmarshal(m, b)
}
func (m *VirtualHardDisk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualHardDisk.Marshal(b, m, deterministic)
}
func (m *VirtualHardDisk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualHardDisk.Merge(m, src)
}
func (m *VirtualHardDisk) XXX_Size() int {
	return xxx_messageInfo_VirtualHardDisk.Size(m)
}
func (m *VirtualHardDisk) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualHardDisk.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualHardDisk proto.InternalMessageInfo

func (m *VirtualHardDisk) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualHardDisk) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualHardDisk) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *VirtualHardDisk) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *VirtualHardDisk) GetContainerName() string {
	if m != nil {
		return m.ContainerName
	}
	return ""
}

func (m *VirtualHardDisk) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualHardDisk) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *VirtualHardDisk) GetDynamic() bool {
	if m != nil {
		return m.Dynamic
	}
	return false
}

func (m *VirtualHardDisk) GetBlocksizebytes() int32 {
	if m != nil {
		return m.Blocksizebytes
	}
	return 0
}

func (m *VirtualHardDisk) GetLogicalsectorbytes() int32 {
	if m != nil {
		return m.Logicalsectorbytes
	}
	return 0
}

func (m *VirtualHardDisk) GetPhysicalsectorbytes() int32 {
	if m != nil {
		return m.Physicalsectorbytes
	}
	return 0
}

func (m *VirtualHardDisk) GetControllernumber() int64 {
	if m != nil {
		return m.Controllernumber
	}
	return 0
}

func (m *VirtualHardDisk) GetControllerlocation() int64 {
	if m != nil {
		return m.Controllerlocation
	}
	return 0
}

func (m *VirtualHardDisk) GetDisknumber() int64 {
	if m != nil {
		return m.Disknumber
	}
	return 0
}

func (m *VirtualHardDisk) GetVirtualmachineName() string {
	if m != nil {
		return m.VirtualmachineName
	}
	return ""
}

func (m *VirtualHardDisk) GetScsipath() string {
	if m != nil {
		return m.Scsipath
	}
	return ""
}

func (m *VirtualHardDisk) GetVirtualharddisktype() VirtualHardDiskType {
	if m != nil {
		return m.Virtualharddisktype
	}
	return VirtualHardDiskType_OS_VIRTUALHARDDISK
}

func (m *VirtualHardDisk) GetEntity() *common.Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.nodeagent.storage.VirtualHardDiskType", VirtualHardDiskType_name, VirtualHardDiskType_value)
	proto.RegisterType((*VirtualHardDiskRequest)(nil), "moc.nodeagent.storage.VirtualHardDiskRequest")
	proto.RegisterType((*VirtualHardDiskResponse)(nil), "moc.nodeagent.storage.VirtualHardDiskResponse")
	proto.RegisterType((*VirtualHardDisk)(nil), "moc.nodeagent.storage.VirtualHardDisk")
}

func init() { proto.RegisterFile("virtualharddisk.proto", fileDescriptor_f4b382f86170a6e5) }

var fileDescriptor_f4b382f86170a6e5 = []byte{
	// 683 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0x5d, 0x4f, 0x13, 0x41,
	0x14, 0x65, 0x81, 0x16, 0xb8, 0x85, 0x02, 0x17, 0xa8, 0x6b, 0x35, 0xa4, 0xa9, 0x86, 0x34, 0x24,
	0x6c, 0x4d, 0x25, 0xbe, 0x17, 0x4b, 0x02, 0xc1, 0x40, 0xb2, 0x45, 0x1e, 0x8c, 0xd1, 0x4c, 0xa7,
	0x43, 0x3b, 0xe9, 0xee, 0xce, 0x3a, 0x33, 0x8b, 0xa9, 0x3f, 0x4a, 0x7f, 0x84, 0x8f, 0xfe, 0x29,
	0x33, 0xb3, 0x4b, 0x29, 0xed, 0x9a, 0xf0, 0xe2, 0xdb, 0xdc, 0x7b, 0xce, 0x3d, 0x73, 0xe6, 0xe3,
	0x5e, 0xd8, 0xbb, 0xe3, 0x52, 0x27, 0x24, 0x18, 0x12, 0xd9, 0xef, 0x73, 0x35, 0xf2, 0x62, 0x29,
	0xb4, 0xc0, 0xbd, 0x50, 0x50, 0x2f, 0x12, 0x7d, 0x46, 0x06, 0x2c, 0xd2, 0x9e, 0xd2, 0x42, 0x92,
	0x01, 0xab, 0xee, 0x0f, 0x84, 0x18, 0x04, 0xac, 0x69, 0x49, 0xbd, 0xe4, 0xb6, 0xf9, 0x5d, 0x92,
	0x38, 0x66, 0x52, 0xa5, 0x65, 0xd5, 0x75, 0x2a, 0xc2, 0x50, 0x44, 0x59, 0x84, 0x91, 0xd0, 0xfc,
	0x96, 0x53, 0xa2, 0xf9, 0x24, 0xf7, 0x62, 0x56, 0x81, 0x85, 0xb1, 0x1e, 0xa7, 0x60, 0xfd, 0xa7,
	0x03, 0x95, 0x9b, 0xd4, 0xcf, 0x19, 0x91, 0xfd, 0x0e, 0x57, 0x23, 0x9f, 0x7d, 0x4b, 0x98, 0xd2,
	0xf8, 0x65, 0x0e, 0xe9, 0x8e, 0x95, 0x66, 0xa1, 0x72, 0x9d, 0xda, 0x52, 0xa3, 0xd4, 0x3a, 0xf0,
	0x72, 0x1d, 0x7b, 0xb3, 0x72, 0xff, 0x50, 0xc1, 0x63, 0xd8, 0xb8, 0x8a, 0x99, 0xb4, 0x56, 0xaf,
	0xc7, 0x31, 0x73, 0x17, 0x6b, 0x4e, 0xa3, 0xdc, 0x2a, 0x5b, 0xd9, 0x09, 0xe2, 0x3f, 0x26, 0xd5,
	0x7f, 0x3b, 0xf0, 0x6c, 0xce, 0xb0, 0x8a, 0x45, 0xa4, 0xd8, 0x7f, 0x77, 0xdc, 0x82, 0xa2, 0xcf,
	0x54, 0x12, 0x68, 0x6b, 0xb5, 0xd4, 0xaa, 0x7a, 0xe9, 0xd5, 0x7a, 0xf7, 0x57, 0xeb, 0x9d, 0x08,
	0x11, 0xdc, 0x90, 0x20, 0x61, 0x7e, 0xc6, 0xc4, 0x5d, 0x28, 0x9c, 0x4a, 0x29, 0xa4, 0xbb, 0x54,
	0x73, 0x1a, 0x6b, 0x7e, 0x1a, 0xd4, 0x7f, 0x15, 0x60, 0x73, 0x66, 0x13, 0x44, 0x58, 0x8e, 0x48,
	0xc8, 0x5c, 0xc7, 0x12, 0xed, 0x1a, 0xcb, 0xb0, 0xc8, 0xfb, 0x76, 0xb7, 0x35, 0x7f, 0x91, 0xf7,
	0xb1, 0x02, 0x45, 0x25, 0x12, 0x49, 0x59, 0x26, 0x97, 0x45, 0xa6, 0x36, 0x26, 0x7a, 0xe8, 0x2e,
	0xa7, 0xb5, 0x66, 0x8d, 0xaf, 0x61, 0x83, 0x8a, 0x48, 0x13, 0x1e, 0x31, 0x79, 0x69, 0x84, 0x0b,
	0x16, 0x7c, 0x9c, 0xc4, 0x57, 0x50, 0x54, 0x9a, 0xe8, 0x44, 0xb9, 0x45, 0x7b, 0xa6, 0x92, 0xbd,
	0xa3, 0xae, 0x4d, 0xf9, 0x19, 0x64, 0xe4, 0x15, 0xff, 0xc1, 0xdc, 0x95, 0x9a, 0xd3, 0x58, 0xf2,
	0xed, 0x1a, 0x5d, 0x58, 0xe9, 0x8f, 0x23, 0x12, 0x72, 0xea, 0xae, 0xd6, 0x9c, 0xc6, 0xaa, 0x7f,
	0x1f, 0xe2, 0x01, 0x94, 0x7b, 0x81, 0xa0, 0x23, 0x43, 0xeb, 0x8d, 0x35, 0x53, 0xee, 0x5a, 0xcd,
	0x69, 0x14, 0xfc, 0x99, 0x2c, 0x7a, 0x80, 0x81, 0x18, 0x70, 0x4a, 0x02, 0xc5, 0xa8, 0x16, 0x32,
	0xe5, 0x82, 0xe5, 0xe6, 0x20, 0xf8, 0x06, 0x76, 0xe2, 0xe1, 0x58, 0xcd, 0x16, 0x94, 0x6c, 0x41,
	0x1e, 0x84, 0x87, 0xb0, 0x65, 0x4e, 0x2b, 0x45, 0x10, 0x30, 0x19, 0x25, 0x61, 0x8f, 0x49, 0x77,
	0xdd, 0x9e, 0x61, 0x2e, 0x6f, 0xdc, 0x3c, 0xe4, 0x02, 0x91, 0xb6, 0x90, 0xbb, 0x61, 0xd9, 0x39,
	0x08, 0xee, 0x03, 0x98, 0xee, 0xcd, 0x54, 0xcb, 0x96, 0x37, 0x95, 0x31, 0x7a, 0x59, 0xa3, 0x87,
	0x84, 0x0e, 0x79, 0xc4, 0xec, 0x1b, 0x6c, 0xda, 0x37, 0xc8, 0x41, 0xb0, 0x0a, 0xab, 0x8a, 0x2a,
	0x6e, 0x9f, 0x71, 0xcb, 0xb2, 0x26, 0x31, 0x7e, 0x86, 0x9d, 0x99, 0xa1, 0xa1, 0x4d, 0xc3, 0x6c,
	0xdb, 0x86, 0x39, 0x7c, 0xda, 0xaf, 0x36, 0xdd, 0xe3, 0xe7, 0xc9, 0x98, 0x2f, 0xc0, 0x22, 0xcd,
	0xf5, 0xd8, 0xc5, 0xa9, 0x2f, 0x70, 0x6a, 0x53, 0x7e, 0x06, 0x1d, 0x5e, 0xc0, 0x4e, 0x8e, 0x20,
	0x56, 0x00, 0xaf, 0xba, 0x5f, 0x6f, 0xce, 0xfd, 0xeb, 0x8f, 0xed, 0x0f, 0x67, 0x6d, 0xbf, 0xd3,
	0x39, 0xef, 0x5e, 0x6c, 0x2d, 0xe0, 0x4b, 0x70, 0x3b, 0xed, 0xeb, 0xb6, 0x89, 0xe6, 0x50, 0xa7,
	0xf5, 0xc7, 0x81, 0xdd, 0x19, 0xb5, 0xb6, 0xf1, 0x8e, 0x1c, 0x8a, 0xe7, 0xd1, 0x9d, 0x18, 0x31,
	0x3c, 0x7a, 0x62, 0xaf, 0xa6, 0xc3, 0xaa, 0xea, 0x3d, 0x95, 0x9e, 0x8e, 0x8a, 0xfa, 0x02, 0x9e,
	0xc1, 0xf6, 0xfb, 0x21, 0xa3, 0xa3, 0xcb, 0xa9, 0x89, 0x89, 0x95, 0xb9, 0x8e, 0x3e, 0x35, 0xc3,
	0xb2, 0xfa, 0xdc, 0xca, 0x4f, 0x53, 0x1f, 0x94, 0x4e, 0xde, 0x7d, 0x3a, 0x1e, 0x70, 0x3d, 0x4c,
	0x7a, 0x1e, 0x15, 0x61, 0x33, 0xe4, 0x54, 0x0a, 0x25, 0x6e, 0x75, 0x33, 0x14, 0xf4, 0xc8, 0xea,
	0x34, 0x65, 0x4c, 0x9b, 0x13, 0x6f, 0xcd, 0xcc, 0x5b, 0xaf, 0x68, 0xc1, 0xb7, 0x7f, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x64, 0xed, 0x11, 0x3c, 0x11, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualHardDiskAgentClient is the client API for VirtualHardDiskAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualHardDiskAgentClient interface {
	Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error)
	CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error)
}

type virtualHardDiskAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualHardDiskAgentClient(cc *grpc.ClientConn) VirtualHardDiskAgentClient {
	return &virtualHardDiskAgentClient{cc}
}

func (c *virtualHardDiskAgentClient) Invoke(ctx context.Context, in *VirtualHardDiskRequest, opts ...grpc.CallOption) (*VirtualHardDiskResponse, error) {
	out := new(VirtualHardDiskResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.VirtualHardDiskAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualHardDiskAgentClient) CheckNotification(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.NotificationResponse, error) {
	out := new(common.NotificationResponse)
	err := c.cc.Invoke(ctx, "/moc.nodeagent.storage.VirtualHardDiskAgent/CheckNotification", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualHardDiskAgentServer is the server API for VirtualHardDiskAgent service.
type VirtualHardDiskAgentServer interface {
	Invoke(context.Context, *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error)
	CheckNotification(context.Context, *empty.Empty) (*common.NotificationResponse, error)
}

// UnimplementedVirtualHardDiskAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualHardDiskAgentServer struct {
}

func (*UnimplementedVirtualHardDiskAgentServer) Invoke(ctx context.Context, req *VirtualHardDiskRequest) (*VirtualHardDiskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}
func (*UnimplementedVirtualHardDiskAgentServer) CheckNotification(ctx context.Context, req *empty.Empty) (*common.NotificationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckNotification not implemented")
}

func RegisterVirtualHardDiskAgentServer(s *grpc.Server, srv VirtualHardDiskAgentServer) {
	s.RegisterService(&_VirtualHardDiskAgent_serviceDesc, srv)
}

func _VirtualHardDiskAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualHardDiskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.VirtualHardDiskAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).Invoke(ctx, req.(*VirtualHardDiskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualHardDiskAgent_CheckNotification_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualHardDiskAgentServer).CheckNotification(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.nodeagent.storage.VirtualHardDiskAgent/CheckNotification",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualHardDiskAgentServer).CheckNotification(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualHardDiskAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.nodeagent.storage.VirtualHardDiskAgent",
	HandlerType: (*VirtualHardDiskAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualHardDiskAgent_Invoke_Handler,
		},
		{
			MethodName: "CheckNotification",
			Handler:    _VirtualHardDiskAgent_CheckNotification_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "virtualharddisk.proto",
}
