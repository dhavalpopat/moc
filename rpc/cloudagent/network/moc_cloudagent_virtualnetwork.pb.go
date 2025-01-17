// Code generated by protoc-gen-go. DO NOT EDIT.
// source: moc_cloudagent_virtualnetwork.proto

package network

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	common "github.com/microsoft/moc/rpc/common"
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

type VirtualNetworkType int32

const (
	VirtualNetworkType_NAT         VirtualNetworkType = 0
	VirtualNetworkType_Transparent VirtualNetworkType = 1
	VirtualNetworkType_L2Bridge    VirtualNetworkType = 2
	VirtualNetworkType_L2Tunnel    VirtualNetworkType = 3
	VirtualNetworkType_ICS         VirtualNetworkType = 4
	VirtualNetworkType_Private     VirtualNetworkType = 5
	VirtualNetworkType_Overlay     VirtualNetworkType = 6
	VirtualNetworkType_Internal    VirtualNetworkType = 7
	VirtualNetworkType_Mirrored    VirtualNetworkType = 8
)

var VirtualNetworkType_name = map[int32]string{
	0: "NAT",
	1: "Transparent",
	2: "L2Bridge",
	3: "L2Tunnel",
	4: "ICS",
	5: "Private",
	6: "Overlay",
	7: "Internal",
	8: "Mirrored",
}

var VirtualNetworkType_value = map[string]int32{
	"NAT":         0,
	"Transparent": 1,
	"L2Bridge":    2,
	"L2Tunnel":    3,
	"ICS":         4,
	"Private":     5,
	"Overlay":     6,
	"Internal":    7,
	"Mirrored":    8,
}

func (x VirtualNetworkType) String() string {
	return proto.EnumName(VirtualNetworkType_name, int32(x))
}

func (VirtualNetworkType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{0}
}

type VirtualNetworkRequest struct {
	VirtualNetworks      []*VirtualNetwork `protobuf:"bytes,1,rep,name=VirtualNetworks,proto3" json:"VirtualNetworks,omitempty"`
	OperationType        common.Operation  `protobuf:"varint,2,opt,name=OperationType,proto3,enum=moc.Operation" json:"OperationType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *VirtualNetworkRequest) Reset()         { *m = VirtualNetworkRequest{} }
func (m *VirtualNetworkRequest) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkRequest) ProtoMessage()    {}
func (*VirtualNetworkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{0}
}

func (m *VirtualNetworkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkRequest.Unmarshal(m, b)
}
func (m *VirtualNetworkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkRequest.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkRequest.Merge(m, src)
}
func (m *VirtualNetworkRequest) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkRequest.Size(m)
}
func (m *VirtualNetworkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkRequest proto.InternalMessageInfo

func (m *VirtualNetworkRequest) GetVirtualNetworks() []*VirtualNetwork {
	if m != nil {
		return m.VirtualNetworks
	}
	return nil
}

func (m *VirtualNetworkRequest) GetOperationType() common.Operation {
	if m != nil {
		return m.OperationType
	}
	return common.Operation_GET
}

type VirtualNetworkResponse struct {
	VirtualNetworks      []*VirtualNetwork   `protobuf:"bytes,1,rep,name=VirtualNetworks,proto3" json:"VirtualNetworks,omitempty"`
	Result               *wrappers.BoolValue `protobuf:"bytes,2,opt,name=Result,proto3" json:"Result,omitempty"`
	Error                string              `protobuf:"bytes,3,opt,name=Error,proto3" json:"Error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *VirtualNetworkResponse) Reset()         { *m = VirtualNetworkResponse{} }
func (m *VirtualNetworkResponse) String() string { return proto.CompactTextString(m) }
func (*VirtualNetworkResponse) ProtoMessage()    {}
func (*VirtualNetworkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{1}
}

func (m *VirtualNetworkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetworkResponse.Unmarshal(m, b)
}
func (m *VirtualNetworkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetworkResponse.Marshal(b, m, deterministic)
}
func (m *VirtualNetworkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetworkResponse.Merge(m, src)
}
func (m *VirtualNetworkResponse) XXX_Size() int {
	return xxx_messageInfo_VirtualNetworkResponse.Size(m)
}
func (m *VirtualNetworkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetworkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetworkResponse proto.InternalMessageInfo

func (m *VirtualNetworkResponse) GetVirtualNetworks() []*VirtualNetwork {
	if m != nil {
		return m.VirtualNetworks
	}
	return nil
}

func (m *VirtualNetworkResponse) GetResult() *wrappers.BoolValue {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *VirtualNetworkResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type VirtualNetwork struct {
	Name                 string             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Subnets              []*Subnet          `protobuf:"bytes,3,rep,name=subnets,proto3" json:"subnets,omitempty"`
	Dns                  *common.Dns        `protobuf:"bytes,4,opt,name=dns,proto3" json:"dns,omitempty"`
	Type                 VirtualNetworkType `protobuf:"varint,5,opt,name=type,proto3,enum=moc.cloudagent.network.VirtualNetworkType" json:"type,omitempty"`
	Nodefqdn             string             `protobuf:"bytes,6,opt,name=nodefqdn,proto3" json:"nodefqdn,omitempty"`
	GroupName            string             `protobuf:"bytes,7,opt,name=groupName,proto3" json:"groupName,omitempty"`
	Status               *common.Status     `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	LocationName         string             `protobuf:"bytes,9,opt,name=locationName,proto3" json:"locationName,omitempty"`
	MacPoolName          string             `protobuf:"bytes,10,opt,name=macPoolName,proto3" json:"macPoolName,omitempty"`
	Tags                 *common.Tags       `protobuf:"bytes,12,opt,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *VirtualNetwork) Reset()         { *m = VirtualNetwork{} }
func (m *VirtualNetwork) String() string { return proto.CompactTextString(m) }
func (*VirtualNetwork) ProtoMessage()    {}
func (*VirtualNetwork) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{2}
}

func (m *VirtualNetwork) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VirtualNetwork.Unmarshal(m, b)
}
func (m *VirtualNetwork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VirtualNetwork.Marshal(b, m, deterministic)
}
func (m *VirtualNetwork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VirtualNetwork.Merge(m, src)
}
func (m *VirtualNetwork) XXX_Size() int {
	return xxx_messageInfo_VirtualNetwork.Size(m)
}
func (m *VirtualNetwork) XXX_DiscardUnknown() {
	xxx_messageInfo_VirtualNetwork.DiscardUnknown(m)
}

var xxx_messageInfo_VirtualNetwork proto.InternalMessageInfo

func (m *VirtualNetwork) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VirtualNetwork) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VirtualNetwork) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func (m *VirtualNetwork) GetDns() *common.Dns {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *VirtualNetwork) GetType() VirtualNetworkType {
	if m != nil {
		return m.Type
	}
	return VirtualNetworkType_NAT
}

func (m *VirtualNetwork) GetNodefqdn() string {
	if m != nil {
		return m.Nodefqdn
	}
	return ""
}

func (m *VirtualNetwork) GetGroupName() string {
	if m != nil {
		return m.GroupName
	}
	return ""
}

func (m *VirtualNetwork) GetStatus() *common.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *VirtualNetwork) GetLocationName() string {
	if m != nil {
		return m.LocationName
	}
	return ""
}

func (m *VirtualNetwork) GetMacPoolName() string {
	if m != nil {
		return m.MacPoolName
	}
	return ""
}

func (m *VirtualNetwork) GetTags() *common.Tags {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Subnet struct {
	Name                    string                                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                      string                                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Cidr                    string                                `protobuf:"bytes,3,opt,name=cidr,proto3" json:"cidr,omitempty"`
	Routes                  []*common.Route                       `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty"`
	Allocation              common.IPAllocationMethod             `protobuf:"varint,5,opt,name=allocation,proto3,enum=moc.IPAllocationMethod" json:"allocation,omitempty"`
	Vlan                    uint32                                `protobuf:"varint,6,opt,name=vlan,proto3" json:"vlan,omitempty"`
	Ippools                 []*common.IPPool                      `protobuf:"bytes,7,rep,name=ippools,proto3" json:"ippools,omitempty"`
	Networksecuritygroupref *common.NetworkSecurityGroupReference `protobuf:"bytes,9,opt,name=networksecuritygroupref,proto3" json:"networksecuritygroupref,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                              `json:"-"`
	XXX_unrecognized        []byte                                `json:"-"`
	XXX_sizecache           int32                                 `json:"-"`
}

func (m *Subnet) Reset()         { *m = Subnet{} }
func (m *Subnet) String() string { return proto.CompactTextString(m) }
func (*Subnet) ProtoMessage()    {}
func (*Subnet) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{3}
}

func (m *Subnet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Subnet.Unmarshal(m, b)
}
func (m *Subnet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Subnet.Marshal(b, m, deterministic)
}
func (m *Subnet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subnet.Merge(m, src)
}
func (m *Subnet) XXX_Size() int {
	return xxx_messageInfo_Subnet.Size(m)
}
func (m *Subnet) XXX_DiscardUnknown() {
	xxx_messageInfo_Subnet.DiscardUnknown(m)
}

var xxx_messageInfo_Subnet proto.InternalMessageInfo

func (m *Subnet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subnet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Subnet) GetCidr() string {
	if m != nil {
		return m.Cidr
	}
	return ""
}

func (m *Subnet) GetRoutes() []*common.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *Subnet) GetAllocation() common.IPAllocationMethod {
	if m != nil {
		return m.Allocation
	}
	return common.IPAllocationMethod_Invalid
}

func (m *Subnet) GetVlan() uint32 {
	if m != nil {
		return m.Vlan
	}
	return 0
}

func (m *Subnet) GetIppools() []*common.IPPool {
	if m != nil {
		return m.Ippools
	}
	return nil
}

func (m *Subnet) GetNetworksecuritygroupref() *common.NetworkSecurityGroupReference {
	if m != nil {
		return m.Networksecuritygroupref
	}
	return nil
}

type Ipam struct {
	Type                 string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Subnets              []*Subnet `protobuf:"bytes,2,rep,name=subnets,proto3" json:"subnets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Ipam) Reset()         { *m = Ipam{} }
func (m *Ipam) String() string { return proto.CompactTextString(m) }
func (*Ipam) ProtoMessage()    {}
func (*Ipam) Descriptor() ([]byte, []int) {
	return fileDescriptor_65af3e74534f0214, []int{4}
}

func (m *Ipam) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ipam.Unmarshal(m, b)
}
func (m *Ipam) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ipam.Marshal(b, m, deterministic)
}
func (m *Ipam) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ipam.Merge(m, src)
}
func (m *Ipam) XXX_Size() int {
	return xxx_messageInfo_Ipam.Size(m)
}
func (m *Ipam) XXX_DiscardUnknown() {
	xxx_messageInfo_Ipam.DiscardUnknown(m)
}

var xxx_messageInfo_Ipam proto.InternalMessageInfo

func (m *Ipam) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Ipam) GetSubnets() []*Subnet {
	if m != nil {
		return m.Subnets
	}
	return nil
}

func init() {
	proto.RegisterEnum("moc.cloudagent.network.VirtualNetworkType", VirtualNetworkType_name, VirtualNetworkType_value)
	proto.RegisterType((*VirtualNetworkRequest)(nil), "moc.cloudagent.network.VirtualNetworkRequest")
	proto.RegisterType((*VirtualNetworkResponse)(nil), "moc.cloudagent.network.VirtualNetworkResponse")
	proto.RegisterType((*VirtualNetwork)(nil), "moc.cloudagent.network.VirtualNetwork")
	proto.RegisterType((*Subnet)(nil), "moc.cloudagent.network.Subnet")
	proto.RegisterType((*Ipam)(nil), "moc.cloudagent.network.Ipam")
}

func init() {
	proto.RegisterFile("moc_cloudagent_virtualnetwork.proto", fileDescriptor_65af3e74534f0214)
}

var fileDescriptor_65af3e74534f0214 = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xd1, 0x8e, 0xdb, 0x44,
	0x14, 0xad, 0x13, 0xaf, 0xe3, 0x5c, 0x6f, 0xb7, 0xd6, 0x00, 0x5d, 0x2b, 0x82, 0xd5, 0xca, 0x15,
	0xd5, 0xaa, 0x12, 0xb6, 0x08, 0x48, 0xf0, 0x84, 0xb4, 0x0b, 0x08, 0xa5, 0xa2, 0xdb, 0xd5, 0x24,
	0xea, 0x03, 0x42, 0xaa, 0x26, 0xf6, 0xc4, 0xb5, 0x6a, 0xcf, 0xb8, 0x33, 0xe3, 0x54, 0xfb, 0xc6,
	0x23, 0x7c, 0x05, 0x7c, 0x04, 0xbc, 0xf1, 0x71, 0x68, 0xae, 0x9d, 0xdd, 0x64, 0xa1, 0x52, 0x78,
	0xe8, 0x53, 0x3c, 0x73, 0xce, 0x3d, 0xf7, 0xcc, 0xb9, 0x93, 0x81, 0x47, 0xb5, 0xcc, 0x5e, 0x66,
	0x95, 0x6c, 0x73, 0x56, 0x70, 0x61, 0x5e, 0xae, 0x4b, 0x65, 0x5a, 0x56, 0x09, 0x6e, 0xde, 0x4a,
	0xf5, 0x3a, 0x69, 0x94, 0x34, 0x92, 0x3c, 0xac, 0x65, 0x96, 0xdc, 0x92, 0x92, 0x1e, 0x9d, 0x9c,
	0x14, 0x52, 0x16, 0x15, 0x4f, 0x91, 0xb5, 0x6c, 0x57, 0xe9, 0x5b, 0xc5, 0x9a, 0x86, 0x2b, 0xdd,
	0xd5, 0x4d, 0x8e, 0x51, 0x5c, 0xd6, 0xb5, 0x14, 0xfd, 0x4f, 0x0f, 0x9c, 0x6c, 0x01, 0xbd, 0xd8,
	0x36, 0x1e, 0xff, 0xee, 0xc0, 0x47, 0x2f, 0x3a, 0x27, 0x97, 0x1d, 0x4c, 0xf9, 0x9b, 0x96, 0x6b,
	0x43, 0xae, 0xe0, 0xc1, 0x2e, 0xa0, 0x23, 0xe7, 0x74, 0x78, 0x16, 0x4c, 0x1f, 0x27, 0xff, 0x6d,
	0x32, 0xb9, 0xa3, 0x73, 0xb7, 0x9c, 0x7c, 0x09, 0xf7, 0x9f, 0x37, 0x5c, 0x31, 0x53, 0x4a, 0xb1,
	0xb8, 0x6e, 0x78, 0x34, 0x38, 0x75, 0xce, 0x8e, 0xa6, 0x47, 0xa8, 0x77, 0x83, 0xd0, 0x5d, 0x52,
	0xfc, 0xa7, 0x03, 0x0f, 0xef, 0x3a, 0xd4, 0x8d, 0x14, 0x9a, 0xbf, 0x07, 0x8b, 0x53, 0xf0, 0x28,
	0xd7, 0x6d, 0x65, 0xd0, 0x5b, 0x30, 0x9d, 0x24, 0x5d, 0xf0, 0xc9, 0x26, 0xf8, 0xe4, 0x42, 0xca,
	0xea, 0x05, 0xab, 0x5a, 0x4e, 0x7b, 0x26, 0xf9, 0x10, 0x0e, 0xbe, 0x57, 0x4a, 0xaa, 0x68, 0x78,
	0xea, 0x9c, 0x8d, 0x69, 0xb7, 0x88, 0xff, 0x18, 0xc2, 0xd1, 0xae, 0x3a, 0x21, 0xe0, 0x0a, 0x56,
	0xf3, 0xc8, 0x41, 0x1e, 0x7e, 0x93, 0x23, 0x18, 0x94, 0x39, 0x36, 0x1b, 0xd3, 0x41, 0x99, 0x93,
	0xaf, 0x61, 0xa4, 0xdb, 0xa5, 0xe0, 0x46, 0x47, 0x43, 0x3c, 0xca, 0xc9, 0xbb, 0x8e, 0x32, 0x47,
	0x1a, 0xdd, 0xd0, 0xc9, 0x29, 0x0c, 0x73, 0xa1, 0x23, 0x17, 0x7d, 0xfb, 0x58, 0xf5, 0x9d, 0xd0,
	0x17, 0xee, 0xaf, 0x7f, 0x45, 0x0e, 0xb5, 0x10, 0xf9, 0x06, 0x5c, 0x63, 0x63, 0x3f, 0xc0, 0xd8,
	0x9f, 0xec, 0x97, 0x91, 0x9d, 0x01, 0xc5, 0x3a, 0x32, 0x01, 0x5f, 0xc8, 0x9c, 0xaf, 0xde, 0xe4,
	0x22, 0xf2, 0xd0, 0xf1, 0xcd, 0x9a, 0x7c, 0x0c, 0xe3, 0x42, 0xc9, 0xb6, 0xb9, 0xb4, 0x07, 0x1c,
	0x21, 0x78, 0xbb, 0x41, 0x1e, 0x81, 0xa7, 0x0d, 0x33, 0xad, 0x8e, 0x7c, 0xb4, 0x17, 0x60, 0xef,
	0x39, 0x6e, 0xd1, 0x1e, 0x22, 0x31, 0x1c, 0x56, 0x32, 0xc3, 0xc1, 0xa3, 0xca, 0x18, 0x55, 0x76,
	0xf6, 0xc8, 0x63, 0x08, 0x6a, 0x96, 0x5d, 0x49, 0x59, 0x21, 0x05, 0x2c, 0xa5, 0x3f, 0xe2, 0x36,
	0x40, 0x3e, 0x01, 0xd7, 0xb0, 0x42, 0x47, 0x87, 0xd8, 0x6e, 0x8c, 0xed, 0x16, 0xac, 0xd0, 0x14,
	0xb7, 0x9f, 0xba, 0x7e, 0x10, 0x1e, 0xc6, 0x7f, 0x0f, 0xc0, 0xeb, 0x52, 0xdc, 0x6b, 0x34, 0x04,
	0xdc, 0xac, 0xcc, 0x37, 0x63, 0xc6, 0x6f, 0x12, 0x83, 0xa7, 0x64, 0x6b, 0xb8, 0xcd, 0xdd, 0x4e,
	0x0b, 0xb0, 0x13, 0xb5, 0x5b, 0xb4, 0x47, 0xc8, 0x57, 0x00, 0xac, 0xda, 0x9c, 0xa2, 0x0f, 0xff,
	0x18, 0x79, 0xb3, 0xab, 0xf3, 0x1b, 0xe0, 0x19, 0x37, 0xaf, 0x64, 0x4e, 0xb7, 0xa8, 0xb6, 0xe1,
	0xba, 0x62, 0x5d, 0xd6, 0xf7, 0x29, 0x7e, 0x93, 0x4f, 0x61, 0x54, 0x36, 0x8d, 0x94, 0x95, 0x8e,
	0x46, 0xd8, 0x31, 0xe8, 0x95, 0xec, 0xd1, 0xe9, 0x06, 0x23, 0x3f, 0xc3, 0x71, 0x3f, 0x4e, 0xcd,
	0xb3, 0x56, 0x95, 0xe6, 0x1a, 0x87, 0xa1, 0xf8, 0x0a, 0x63, 0x0d, 0xa6, 0x31, 0x96, 0xf5, 0x33,
	0x9e, 0xf7, 0x9c, 0x1f, 0x2c, 0x87, 0xf2, 0x15, 0x57, 0x5c, 0x64, 0x9c, 0xbe, 0x4b, 0xe2, 0xa9,
	0xeb, 0xfb, 0xe1, 0x38, 0x5e, 0x80, 0x3b, 0x6b, 0x58, 0x6d, 0x6d, 0xe2, 0xb5, 0xea, 0xb3, 0xc3,
	0xab, 0xb2, 0x75, 0x8d, 0x07, 0xff, 0xeb, 0x1a, 0x3f, 0xf9, 0xcd, 0x01, 0xf2, 0xef, 0x1b, 0x48,
	0x46, 0x30, 0xbc, 0x3c, 0x5f, 0x84, 0xf7, 0xc8, 0x03, 0x08, 0x16, 0x8a, 0x09, 0xdd, 0x30, 0xc5,
	0x85, 0x09, 0x1d, 0x72, 0x08, 0xfe, 0x8f, 0xd3, 0x0b, 0x55, 0xe6, 0x05, 0x0f, 0x07, 0xdd, 0x6a,
	0xd1, 0x0a, 0xc1, 0xab, 0x70, 0x68, 0xab, 0x66, 0xdf, 0xce, 0x43, 0x97, 0x04, 0x30, 0xba, 0x52,
	0xe5, 0x9a, 0x19, 0x1e, 0x1e, 0xd8, 0xc5, 0xf3, 0x35, 0x57, 0x15, 0xbb, 0x0e, 0x3d, 0x5b, 0x30,
	0x13, 0x86, 0x2b, 0xc1, 0xaa, 0x70, 0x64, 0x57, 0xcf, 0x4a, 0xfb, 0xff, 0xe5, 0x79, 0xe8, 0x4f,
	0x7f, 0x71, 0xe0, 0x83, 0x5d, 0x2f, 0xe7, 0xd6, 0x3b, 0x29, 0xc1, 0x9b, 0x89, 0xb5, 0x7c, 0xcd,
	0xc9, 0x67, 0x7b, 0x3e, 0x34, 0xdd, 0x9b, 0x3a, 0x49, 0xf6, 0xa5, 0x77, 0x0f, 0x5c, 0x7c, 0xef,
	0xe2, 0xf3, 0x9f, 0xd2, 0xa2, 0x34, 0xaf, 0xda, 0x65, 0x92, 0xc9, 0x3a, 0xad, 0xcb, 0x4c, 0x49,
	0x2d, 0x57, 0x26, 0xad, 0x65, 0x96, 0xaa, 0x26, 0x4b, 0x6f, 0xb5, 0xd2, 0x5e, 0x6b, 0xe9, 0xe1,
	0x5b, 0xf5, 0xc5, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x04, 0xe9, 0x22, 0x48, 0x71, 0x06, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VirtualNetworkAgentClient is the client API for VirtualNetworkAgent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VirtualNetworkAgentClient interface {
	Invoke(ctx context.Context, in *VirtualNetworkRequest, opts ...grpc.CallOption) (*VirtualNetworkResponse, error)
}

type virtualNetworkAgentClient struct {
	cc *grpc.ClientConn
}

func NewVirtualNetworkAgentClient(cc *grpc.ClientConn) VirtualNetworkAgentClient {
	return &virtualNetworkAgentClient{cc}
}

func (c *virtualNetworkAgentClient) Invoke(ctx context.Context, in *VirtualNetworkRequest, opts ...grpc.CallOption) (*VirtualNetworkResponse, error) {
	out := new(VirtualNetworkResponse)
	err := c.cc.Invoke(ctx, "/moc.cloudagent.network.VirtualNetworkAgent/Invoke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VirtualNetworkAgentServer is the server API for VirtualNetworkAgent service.
type VirtualNetworkAgentServer interface {
	Invoke(context.Context, *VirtualNetworkRequest) (*VirtualNetworkResponse, error)
}

// UnimplementedVirtualNetworkAgentServer can be embedded to have forward compatible implementations.
type UnimplementedVirtualNetworkAgentServer struct {
}

func (*UnimplementedVirtualNetworkAgentServer) Invoke(ctx context.Context, req *VirtualNetworkRequest) (*VirtualNetworkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Invoke not implemented")
}

func RegisterVirtualNetworkAgentServer(s *grpc.Server, srv VirtualNetworkAgentServer) {
	s.RegisterService(&_VirtualNetworkAgent_serviceDesc, srv)
}

func _VirtualNetworkAgent_Invoke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VirtualNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualNetworkAgentServer).Invoke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/moc.cloudagent.network.VirtualNetworkAgent/Invoke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualNetworkAgentServer).Invoke(ctx, req.(*VirtualNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VirtualNetworkAgent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "moc.cloudagent.network.VirtualNetworkAgent",
	HandlerType: (*VirtualNetworkAgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Invoke",
			Handler:    _VirtualNetworkAgent_Invoke_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "moc_cloudagent_virtualnetwork.proto",
}
