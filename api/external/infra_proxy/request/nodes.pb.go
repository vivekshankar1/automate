// Code generated by protoc-gen-go. DO NOT EDIT.
// source: external/infra_proxy/request/nodes.proto

package request

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Chef object type (e.g. 'cookbooks', 'roles', 'chef_environment').
	ChefType string `protobuf:"bytes,3,opt,name=chef_type,json=chefType,proto3" json:"chef_type,omitempty"`
	// Chef object name.
	Name string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	// Chef object version.
	Version              string   `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AffectedNodes) Reset()         { *m = AffectedNodes{} }
func (m *AffectedNodes) String() string { return proto.CompactTextString(m) }
func (*AffectedNodes) ProtoMessage()    {}
func (*AffectedNodes) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dd0f3d187bd3788, []int{0}
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

func (m *AffectedNodes) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *AffectedNodes) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *AffectedNodes) GetChefType() string {
	if m != nil {
		return m.ChefType
	}
	return ""
}

func (m *AffectedNodes) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AffectedNodes) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type DeleteNode struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Node name.
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteNode) Reset()         { *m = DeleteNode{} }
func (m *DeleteNode) String() string { return proto.CompactTextString(m) }
func (*DeleteNode) ProtoMessage()    {}
func (*DeleteNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dd0f3d187bd3788, []int{1}
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

func (m *DeleteNode) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *DeleteNode) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *DeleteNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type UpdateNode struct {
	// Chef organization ID.
	OrgId string `protobuf:"bytes,1,opt,name=org_id,json=orgId,proto3" json:"org_id,omitempty"`
	// Chef Infra Server ID.
	ServerId string `protobuf:"bytes,2,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
	// Node name.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Node environment.
	Environment string `protobuf:"bytes,4,opt,name=environment,proto3" json:"environment,omitempty"`
	// Node policy name.
	PolicyName string `protobuf:"bytes,5,opt,name=policy_name,json=policyName,proto3" json:"policy_name,omitempty"`
	// Node policy group.
	PolicyGroup string `protobuf:"bytes,6,opt,name=policy_group,json=policyGroup,proto3" json:"policy_group,omitempty"`
	// Node run-list.
	RunList []string `protobuf:"bytes,7,rep,name=run_list,json=runList,proto3" json:"run_list,omitempty"`
	// Node automatic attributes JSON.
	AutomaticAttributes *_struct.Struct `protobuf:"bytes,8,opt,name=automatic_attributes,json=automaticAttributes,proto3" json:"automatic_attributes,omitempty"`
	// Node default attributes JSON.
	DefaultAttributes *_struct.Struct `protobuf:"bytes,9,opt,name=default_attributes,json=defaultAttributes,proto3" json:"default_attributes,omitempty"`
	// Node normal attributes JSON.
	NormalAttributes *_struct.Struct `protobuf:"bytes,10,opt,name=normal_attributes,json=normalAttributes,proto3" json:"normal_attributes,omitempty"`
	// Node override attributes JSON.
	OverrideAttributes   *_struct.Struct `protobuf:"bytes,11,opt,name=override_attributes,json=overrideAttributes,proto3" json:"override_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateNode) Reset()         { *m = UpdateNode{} }
func (m *UpdateNode) String() string { return proto.CompactTextString(m) }
func (*UpdateNode) ProtoMessage()    {}
func (*UpdateNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dd0f3d187bd3788, []int{2}
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

func (m *UpdateNode) GetOrgId() string {
	if m != nil {
		return m.OrgId
	}
	return ""
}

func (m *UpdateNode) GetServerId() string {
	if m != nil {
		return m.ServerId
	}
	return ""
}

func (m *UpdateNode) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateNode) GetEnvironment() string {
	if m != nil {
		return m.Environment
	}
	return ""
}

func (m *UpdateNode) GetPolicyName() string {
	if m != nil {
		return m.PolicyName
	}
	return ""
}

func (m *UpdateNode) GetPolicyGroup() string {
	if m != nil {
		return m.PolicyGroup
	}
	return ""
}

func (m *UpdateNode) GetRunList() []string {
	if m != nil {
		return m.RunList
	}
	return nil
}

func (m *UpdateNode) GetAutomaticAttributes() *_struct.Struct {
	if m != nil {
		return m.AutomaticAttributes
	}
	return nil
}

func (m *UpdateNode) GetDefaultAttributes() *_struct.Struct {
	if m != nil {
		return m.DefaultAttributes
	}
	return nil
}

func (m *UpdateNode) GetNormalAttributes() *_struct.Struct {
	if m != nil {
		return m.NormalAttributes
	}
	return nil
}

func (m *UpdateNode) GetOverrideAttributes() *_struct.Struct {
	if m != nil {
		return m.OverrideAttributes
	}
	return nil
}

func init() {
	proto.RegisterType((*AffectedNodes)(nil), "chef.automate.api.infra_proxy.request.AffectedNodes")
	proto.RegisterType((*DeleteNode)(nil), "chef.automate.api.infra_proxy.request.DeleteNode")
	proto.RegisterType((*UpdateNode)(nil), "chef.automate.api.infra_proxy.request.UpdateNode")
}

func init() {
	proto.RegisterFile("external/infra_proxy/request/nodes.proto", fileDescriptor_4dd0f3d187bd3788)
}

var fileDescriptor_4dd0f3d187bd3788 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x4b, 0x6f, 0xd4, 0x30,
	0x10, 0xc7, 0xb5, 0x6c, 0x1f, 0xbb, 0xb3, 0x20, 0x51, 0x17, 0x44, 0x78, 0x48, 0x2c, 0x2b, 0x21,
	0xed, 0x29, 0x96, 0xe0, 0x84, 0x38, 0x15, 0x55, 0xbc, 0x84, 0x7a, 0x28, 0xe5, 0xc2, 0x25, 0xf2,
	0x26, 0x93, 0xd4, 0x52, 0x62, 0x9b, 0xc9, 0x78, 0xd5, 0xfd, 0x12, 0x7c, 0x5f, 0x6e, 0xc8, 0x71,
	0xd2, 0xe6, 0xb4, 0x07, 0xc4, 0x2d, 0xf9, 0x4f, 0x7e, 0x3f, 0x8f, 0x33, 0x36, 0xac, 0xf1, 0x86,
	0x91, 0x8c, 0xaa, 0xa5, 0x36, 0x25, 0xa9, 0xcc, 0x91, 0xbd, 0xd9, 0x49, 0xc2, 0x5f, 0x1e, 0x5b,
	0x96, 0xc6, 0x16, 0xd8, 0xa6, 0x8e, 0x2c, 0x5b, 0xf1, 0x3a, 0xbf, 0xc6, 0x32, 0x55, 0x9e, 0x6d,
	0xa3, 0x18, 0x53, 0xe5, 0x74, 0x3a, 0x42, 0xd2, 0x1e, 0x79, 0xf6, 0xa2, 0xb2, 0xb6, 0xaa, 0x51,
	0x76, 0xd0, 0xc6, 0x97, 0xb2, 0x65, 0xf2, 0x39, 0x47, 0xc9, 0xea, 0xf7, 0x04, 0x1e, 0x9c, 0x95,
	0x25, 0xe6, 0x8c, 0xc5, 0x45, 0x90, 0x8b, 0xc7, 0x70, 0x64, 0xa9, 0xca, 0x74, 0x91, 0x4c, 0x96,
	0x93, 0xf5, 0xfc, 0xf2, 0xd0, 0x52, 0xf5, 0xa5, 0x10, 0xcf, 0x61, 0xde, 0x22, 0x6d, 0x91, 0x42,
	0xe5, 0x5e, 0x57, 0x99, 0xc5, 0x20, 0x16, 0x43, 0x33, 0x19, 0xef, 0x1c, 0x26, 0xd3, 0x58, 0x0c,
	0xc1, 0xd5, 0xce, 0xa1, 0x10, 0x70, 0x60, 0x54, 0x83, 0xc9, 0x41, 0x97, 0x77, 0xcf, 0x22, 0x81,
	0xe3, 0x2d, 0x52, 0xab, 0xad, 0x49, 0x0e, 0xbb, 0x78, 0x78, 0x5d, 0x5d, 0x01, 0x9c, 0x63, 0x8d,
	0x8c, 0xa1, 0x9b, 0x7f, 0x6a, 0x66, 0x58, 0x6f, 0x7a, 0xb7, 0xde, 0xea, 0xcf, 0x14, 0xe0, 0x87,
	0x2b, 0xd4, 0xff, 0xd5, 0x8a, 0x25, 0x2c, 0xd0, 0x6c, 0x35, 0x59, 0xd3, 0xa0, 0xe1, 0x7e, 0x87,
	0xe3, 0x48, 0xbc, 0x84, 0x85, 0xb3, 0xb5, 0xce, 0x77, 0x59, 0x07, 0xc7, 0xcd, 0x42, 0x8c, 0x2e,
	0x82, 0xe2, 0x15, 0xdc, 0xef, 0x3f, 0xa8, 0xc8, 0x7a, 0x97, 0x1c, 0x45, 0x47, 0xcc, 0x3e, 0x85,
	0x48, 0x3c, 0x85, 0x19, 0x79, 0x93, 0xd5, 0xba, 0xe5, 0xe4, 0x78, 0x39, 0x0d, 0x7f, 0x8b, 0xbc,
	0xf9, 0xa6, 0x5b, 0x16, 0x5f, 0xe1, 0x51, 0x7f, 0x00, 0x74, 0x9e, 0x29, 0x66, 0xd2, 0x1b, 0xcf,
	0xd8, 0x26, 0xb3, 0xe5, 0x64, 0xbd, 0x78, 0xf3, 0x24, 0x8d, 0xb3, 0x4f, 0x87, 0xd9, 0xa7, 0xdf,
	0xbb, 0xd9, 0x5f, 0x9e, 0xde, 0x42, 0x67, 0xb7, 0x8c, 0xf8, 0x08, 0xa2, 0xc0, 0x52, 0xf9, 0x9a,
	0xc7, 0xa6, 0xf9, 0x7e, 0xd3, 0x49, 0x8f, 0x8c, 0x3c, 0xe7, 0x70, 0x62, 0x2c, 0x35, 0xaa, 0x1e,
	0x6b, 0x60, 0xbf, 0xe6, 0x61, 0x24, 0x46, 0x96, 0xcf, 0x70, 0x6a, 0xb7, 0x48, 0xa4, 0x0b, 0x1c,
	0x7b, 0x16, 0xfb, 0x3d, 0x62, 0x60, 0xee, 0x4c, 0x1f, 0xde, 0xff, 0x7c, 0x57, 0x69, 0xbe, 0xf6,
	0x9b, 0x34, 0xb7, 0x8d, 0x0c, 0xc7, 0x52, 0x0e, 0x97, 0x46, 0x2a, 0xa7, 0xe5, 0xbe, 0x0b, 0xb7,
	0x39, 0xea, 0x56, 0x78, 0xfb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x53, 0x04, 0xa3, 0x97, 0x03,
	0x00, 0x00,
}
