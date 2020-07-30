// Code generated by protoc-gen-go. DO NOT EDIT.
// source: interservice/infra_proxy/response/nodes.proto

package response

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type AffectedNodes struct {
	// List of the nodes which are affected by the chef object.
	Nodes                []*NodeAttribute `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty" toml:"nodes,omitempty" mapstructure:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte           `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32            `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *AffectedNodes) Reset()         { *m = AffectedNodes{} }
func (m *AffectedNodes) String() string { return proto.CompactTextString(m) }
func (*AffectedNodes) ProtoMessage()    {}
func (*AffectedNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_3abe3cd8dc5dfbb0, []int{0}
}

func (m *AffectedNodes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AffectedNodes.Unmarshal(m, b)
}
func (m *AffectedNodes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AffectedNodes.Marshal(b, m, deterministic)
}
func (m *AffectedNodes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AffectedNodes.Merge(m, src)
}
func (m *AffectedNodes) XXX_Size() int {
	return xxx_messageInfo_AffectedNodes.Size(m)
}
func (m *AffectedNodes) XXX_DiscardUnknown() {
	xxx_messageInfo_AffectedNodes.DiscardUnknown(m)
}

var xxx_messageInfo_AffectedNodes proto.InternalMessageInfo

func (m *AffectedNodes) GetNodes() []*NodeAttribute {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type NodeAttribute struct {
	// Node ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" toml:"id,omitempty" mapstructure:"id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	// Node last checkin.
	CheckIn string `protobuf:"bytes,3,opt,name=check_in,json=checkIn,proto3" json:"check_in,omitempty" toml:"check_in,omitempty" mapstructure:"check_in,omitempty"`
	// Node uptime.
	Uptime string `protobuf:"bytes,4,opt,name=uptime,proto3" json:"uptime,omitempty" toml:"uptime,omitempty" mapstructure:"uptime,omitempty"`
	// Node platform.
	Platform string `protobuf:"bytes,5,opt,name=platform,proto3" json:"platform,omitempty" toml:"platform,omitempty" mapstructure:"platform,omitempty"`
	// Node environment name.
	Environment string `protobuf:"bytes,6,opt,name=environment,proto3" json:"environment,omitempty" toml:"environment,omitempty" mapstructure:"environment,omitempty"`
	// Node policy group.
	PolicyGroup          string   `protobuf:"bytes,7,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty" toml:"policy_group,omitempty" mapstructure:"policy_group,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *NodeAttribute) Reset()         { *m = NodeAttribute{} }
func (m *NodeAttribute) String() string { return proto.CompactTextString(m) }
func (*NodeAttribute) ProtoMessage()    {}
func (*NodeAttribute) Descriptor() ([]byte, []int) {
	return fileDescriptor_3abe3cd8dc5dfbb0, []int{1}
}

func (m *NodeAttribute) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeAttribute.Unmarshal(m, b)
}
func (m *NodeAttribute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeAttribute.Marshal(b, m, deterministic)
}
func (m *NodeAttribute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeAttribute.Merge(m, src)
}
func (m *NodeAttribute) XXX_Size() int {
	return xxx_messageInfo_NodeAttribute.Size(m)
}
func (m *NodeAttribute) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeAttribute.DiscardUnknown(m)
}

var xxx_messageInfo_NodeAttribute proto.InternalMessageInfo

func (m *NodeAttribute) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *NodeAttribute) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NodeAttribute) GetCheckIn() string {
	if m != nil {
		return m.CheckIn
	}
	return ""
}

func (m *NodeAttribute) GetUptime() string {
	if m != nil {
		return m.Uptime
	}
	return ""
}

func (m *NodeAttribute) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *NodeAttribute) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *NodeAttribute) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

type DeleteNode struct {
	// Node name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *DeleteNode) Reset()         { *m = DeleteNode{} }
func (m *DeleteNode) String() string { return proto.CompactTextString(m) }
func (*DeleteNode) ProtoMessage()    {}
func (*DeleteNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_3abe3cd8dc5dfbb0, []int{2}
}

func (m *DeleteNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteNode.Unmarshal(m, b)
}
func (m *DeleteNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteNode.Marshal(b, m, deterministic)
}
func (m *DeleteNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteNode.Merge(m, src)
}
func (m *DeleteNode) XXX_Size() int {
	return xxx_messageInfo_DeleteNode.Size(m)
}
func (m *DeleteNode) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteNode.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteNode proto.InternalMessageInfo

func (m *DeleteNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateNode struct {
	// Node name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" toml:"name,omitempty" mapstructure:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_unrecognized     []byte   `json:"-" toml:"-" mapstructure:"-,omitempty"`
	XXX_sizecache        int32    `json:"-" toml:"-" mapstructure:"-,omitempty"`
}

func (m *UpdateNode) Reset()         { *m = UpdateNode{} }
func (m *UpdateNode) String() string { return proto.CompactTextString(m) }
func (*UpdateNode) ProtoMessage()    {}
func (*UpdateNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_3abe3cd8dc5dfbb0, []int{3}
}

func (m *UpdateNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateNode.Unmarshal(m, b)
}
func (m *UpdateNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateNode.Marshal(b, m, deterministic)
}
func (m *UpdateNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateNode.Merge(m, src)
}
func (m *UpdateNode) XXX_Size() int {
	return xxx_messageInfo_UpdateNode.Size(m)
}
func (m *UpdateNode) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateNode.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateNode proto.InternalMessageInfo

func (m *UpdateNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*AffectedNodes)(nil), "chef.automate.domain.infra_proxy.response.AffectedNodes")
	proto.RegisterType((*NodeAttribute)(nil), "chef.automate.domain.infra_proxy.response.NodeAttribute")
	proto.RegisterType((*DeleteNode)(nil), "chef.automate.domain.infra_proxy.response.DeleteNode")
	proto.RegisterType((*UpdateNode)(nil), "chef.automate.domain.infra_proxy.response.UpdateNode")
}

func init() {
	proto.RegisterFile("interservice/infra_proxy/response/nodes.proto", fileDescriptor_3abe3cd8dc5dfbb0)
}

var fileDescriptor_3abe3cd8dc5dfbb0 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x95, 0xfe, 0xc7, 0xa5, 0x0c, 0x1e, 0x90, 0x61, 0x0a, 0x9d, 0xca, 0x80, 0x2d, 0xc1,
	0xc2, 0x84, 0x28, 0x42, 0x42, 0x2c, 0x1d, 0x2a, 0xb1, 0xb0, 0x44, 0xae, 0x73, 0x69, 0x2d, 0xea,
	0x3f, 0x72, 0x2e, 0x15, 0xfd, 0x84, 0x7c, 0x2d, 0x14, 0x87, 0x96, 0x32, 0x20, 0xd8, 0x7c, 0xef,
	0xbd, 0xfb, 0xe9, 0x7c, 0x47, 0xae, 0xb4, 0x45, 0x08, 0x25, 0x84, 0x8d, 0x56, 0x20, 0xb4, 0x2d,
	0x82, 0xcc, 0x7c, 0x70, 0xef, 0x5b, 0x11, 0xa0, 0xf4, 0xce, 0x96, 0x20, 0xac, 0xcb, 0xa1, 0xe4,
	0x3e, 0x38, 0x74, 0xf4, 0x52, 0xad, 0xa0, 0xe0, 0xb2, 0x42, 0x67, 0x24, 0x02, 0xcf, 0x9d, 0x91,
	0xda, 0xf2, 0x83, 0x36, 0xbe, 0x6b, 0x1b, 0x67, 0x64, 0x34, 0x2d, 0x0a, 0x50, 0x08, 0xf9, 0xac,
	0x26, 0xd0, 0x19, 0xe9, 0x46, 0x14, 0x4b, 0xd2, 0xf6, 0x64, 0x78, 0x7d, 0xcb, 0xff, 0xcd, 0xe2,
	0x35, 0x60, 0x8a, 0x18, 0xf4, 0xa2, 0x42, 0x98, 0x37, 0x98, 0xf1, 0x47, 0x42, 0x46, 0x3f, 0x0c,
	0x7a, 0x42, 0x5a, 0x3a, 0x67, 0x49, 0x9a, 0x4c, 0x8e, 0xe6, 0x2d, 0x9d, 0x53, 0x4a, 0x3a, 0x56,
	0x1a, 0x60, 0xad, 0xa8, 0xc4, 0x37, 0x3d, 0x23, 0x03, 0xb5, 0x02, 0xf5, 0x96, 0x69, 0xcb, 0xda,
	0x51, 0xef, 0xc7, 0xfa, 0xd9, 0xd2, 0x53, 0xd2, 0xab, 0x3c, 0x6a, 0x03, 0xac, 0x13, 0x8d, 0xaf,
	0x8a, 0x9e, 0x93, 0x81, 0x5f, 0x4b, 0x2c, 0x5c, 0x30, 0xac, 0x1b, 0x9d, 0x7d, 0x4d, 0x53, 0x32,
	0x04, 0xbb, 0xd1, 0xc1, 0x59, 0x03, 0x16, 0x59, 0x2f, 0xda, 0x87, 0x12, 0xbd, 0x20, 0xc7, 0xde,
	0xad, 0xb5, 0xda, 0x66, 0xcb, 0xe0, 0x2a, 0xcf, 0xfa, 0x4d, 0xa4, 0xd1, 0x9e, 0x6a, 0x69, 0x9c,
	0x12, 0xf2, 0x08, 0x6b, 0x40, 0xa8, 0xbf, 0xb3, 0x9f, 0x3a, 0xf9, 0x9e, 0xba, 0x4e, 0xbc, 0xf8,
	0x5c, 0xfe, 0x9e, 0x78, 0xb8, 0x7f, 0xbd, 0x5b, 0x6a, 0x5c, 0x55, 0x0b, 0xae, 0x9c, 0x11, 0xf5,
	0x6a, 0xc5, 0x6e, 0xb5, 0x42, 0x7a, 0x2d, 0xfe, 0xbc, 0xf3, 0xa2, 0x17, 0x4f, 0x7c, 0xf3, 0x19,
	0x00, 0x00, 0xff, 0xff, 0x90, 0xf1, 0x73, 0x17, 0x13, 0x02, 0x00, 0x00,
}
