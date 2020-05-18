// Code generated by protoc-gen-go. DO NOT EDIT.
// source: components/automate-gateway/api/iam/v2/request/rules.proto

package request

import (
	fmt "fmt"
	common "github.com/chef/automate/components/automate-gateway/api/iam/v2/common"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
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

type CreateRuleReq struct {
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique ID of the project this rule belongs to. Cannot be changed.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Name for the project rule.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the rule affects nodes (`NODE`) or events (`EVENT`).
	Type common.RuleType `protobuf:"varint,4,opt,name=type,proto3,enum=chef.automate.api.iam.v2.RuleType" json:"type,omitempty"`
	// Conditions that ingested resources must match to belong to the project.
	// Will contain one or more.
	Conditions           []*common.Condition `protobuf:"bytes,5,rep,name=conditions,proto3" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CreateRuleReq) Reset()         { *m = CreateRuleReq{} }
func (m *CreateRuleReq) String() string { return proto.CompactTextString(m) }
func (*CreateRuleReq) ProtoMessage()    {}
func (*CreateRuleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{0}
}

func (m *CreateRuleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRuleReq.Unmarshal(m, b)
}
func (m *CreateRuleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRuleReq.Marshal(b, m, deterministic)
}
func (m *CreateRuleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRuleReq.Merge(m, src)
}
func (m *CreateRuleReq) XXX_Size() int {
	return xxx_messageInfo_CreateRuleReq.Size(m)
}
func (m *CreateRuleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRuleReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRuleReq proto.InternalMessageInfo

func (m *CreateRuleReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CreateRuleReq) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *CreateRuleReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateRuleReq) GetType() common.RuleType {
	if m != nil {
		return m.Type
	}
	return common.RuleType_RULE_TYPE_UNSET
}

func (m *CreateRuleReq) GetConditions() []*common.Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type UpdateRuleReq struct {
	// Unique ID. Cannot be changed.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Unique ID of the project this rule belongs to. Cannot be changed.
	ProjectId string `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// Name for the project rule.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// Whether the rule applies to ingested `NODE` or `EVENT` resources.
	// Cannot be changed.
	Type common.RuleType `protobuf:"varint,4,opt,name=type,proto3,enum=chef.automate.api.iam.v2.RuleType" json:"type,omitempty"`
	// Conditions that ingested resources must match to belong to the project.
	// Will contain one or more.
	Conditions           []*common.Condition `protobuf:"bytes,5,rep,name=conditions,proto3" json:"conditions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *UpdateRuleReq) Reset()         { *m = UpdateRuleReq{} }
func (m *UpdateRuleReq) String() string { return proto.CompactTextString(m) }
func (*UpdateRuleReq) ProtoMessage()    {}
func (*UpdateRuleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{1}
}

func (m *UpdateRuleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRuleReq.Unmarshal(m, b)
}
func (m *UpdateRuleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRuleReq.Marshal(b, m, deterministic)
}
func (m *UpdateRuleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRuleReq.Merge(m, src)
}
func (m *UpdateRuleReq) XXX_Size() int {
	return xxx_messageInfo_UpdateRuleReq.Size(m)
}
func (m *UpdateRuleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRuleReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRuleReq proto.InternalMessageInfo

func (m *UpdateRuleReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateRuleReq) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

func (m *UpdateRuleReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateRuleReq) GetType() common.RuleType {
	if m != nil {
		return m.Type
	}
	return common.RuleType_RULE_TYPE_UNSET
}

func (m *UpdateRuleReq) GetConditions() []*common.Condition {
	if m != nil {
		return m.Conditions
	}
	return nil
}

type GetRuleReq struct {
	// ID of the project rule.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the project the rule belongs to.
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetRuleReq) Reset()         { *m = GetRuleReq{} }
func (m *GetRuleReq) String() string { return proto.CompactTextString(m) }
func (*GetRuleReq) ProtoMessage()    {}
func (*GetRuleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{2}
}

func (m *GetRuleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetRuleReq.Unmarshal(m, b)
}
func (m *GetRuleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetRuleReq.Marshal(b, m, deterministic)
}
func (m *GetRuleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetRuleReq.Merge(m, src)
}
func (m *GetRuleReq) XXX_Size() int {
	return xxx_messageInfo_GetRuleReq.Size(m)
}
func (m *GetRuleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_GetRuleReq.DiscardUnknown(m)
}

var xxx_messageInfo_GetRuleReq proto.InternalMessageInfo

func (m *GetRuleReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *GetRuleReq) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

type ListRulesReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRulesReq) Reset()         { *m = ListRulesReq{} }
func (m *ListRulesReq) String() string { return proto.CompactTextString(m) }
func (*ListRulesReq) ProtoMessage()    {}
func (*ListRulesReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{3}
}

func (m *ListRulesReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRulesReq.Unmarshal(m, b)
}
func (m *ListRulesReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRulesReq.Marshal(b, m, deterministic)
}
func (m *ListRulesReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRulesReq.Merge(m, src)
}
func (m *ListRulesReq) XXX_Size() int {
	return xxx_messageInfo_ListRulesReq.Size(m)
}
func (m *ListRulesReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRulesReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListRulesReq proto.InternalMessageInfo

type ListRulesForProjectReq struct {
	// ID of the project.
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRulesForProjectReq) Reset()         { *m = ListRulesForProjectReq{} }
func (m *ListRulesForProjectReq) String() string { return proto.CompactTextString(m) }
func (*ListRulesForProjectReq) ProtoMessage()    {}
func (*ListRulesForProjectReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{4}
}

func (m *ListRulesForProjectReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRulesForProjectReq.Unmarshal(m, b)
}
func (m *ListRulesForProjectReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRulesForProjectReq.Marshal(b, m, deterministic)
}
func (m *ListRulesForProjectReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRulesForProjectReq.Merge(m, src)
}
func (m *ListRulesForProjectReq) XXX_Size() int {
	return xxx_messageInfo_ListRulesForProjectReq.Size(m)
}
func (m *ListRulesForProjectReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRulesForProjectReq.DiscardUnknown(m)
}

var xxx_messageInfo_ListRulesForProjectReq proto.InternalMessageInfo

func (m *ListRulesForProjectReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteRuleReq struct {
	// ID of the project rule.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// ID of the project the rule belongs to.
	ProjectId            string   `protobuf:"bytes,2,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteRuleReq) Reset()         { *m = DeleteRuleReq{} }
func (m *DeleteRuleReq) String() string { return proto.CompactTextString(m) }
func (*DeleteRuleReq) ProtoMessage()    {}
func (*DeleteRuleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{5}
}

func (m *DeleteRuleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRuleReq.Unmarshal(m, b)
}
func (m *DeleteRuleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRuleReq.Marshal(b, m, deterministic)
}
func (m *DeleteRuleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRuleReq.Merge(m, src)
}
func (m *DeleteRuleReq) XXX_Size() int {
	return xxx_messageInfo_DeleteRuleReq.Size(m)
}
func (m *DeleteRuleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRuleReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRuleReq proto.InternalMessageInfo

func (m *DeleteRuleReq) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DeleteRuleReq) GetProjectId() string {
	if m != nil {
		return m.ProjectId
	}
	return ""
}

type ApplyRulesStartReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRulesStartReq) Reset()         { *m = ApplyRulesStartReq{} }
func (m *ApplyRulesStartReq) String() string { return proto.CompactTextString(m) }
func (*ApplyRulesStartReq) ProtoMessage()    {}
func (*ApplyRulesStartReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{6}
}

func (m *ApplyRulesStartReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRulesStartReq.Unmarshal(m, b)
}
func (m *ApplyRulesStartReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRulesStartReq.Marshal(b, m, deterministic)
}
func (m *ApplyRulesStartReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRulesStartReq.Merge(m, src)
}
func (m *ApplyRulesStartReq) XXX_Size() int {
	return xxx_messageInfo_ApplyRulesStartReq.Size(m)
}
func (m *ApplyRulesStartReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRulesStartReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRulesStartReq proto.InternalMessageInfo

type ApplyRulesCancelReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRulesCancelReq) Reset()         { *m = ApplyRulesCancelReq{} }
func (m *ApplyRulesCancelReq) String() string { return proto.CompactTextString(m) }
func (*ApplyRulesCancelReq) ProtoMessage()    {}
func (*ApplyRulesCancelReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{7}
}

func (m *ApplyRulesCancelReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRulesCancelReq.Unmarshal(m, b)
}
func (m *ApplyRulesCancelReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRulesCancelReq.Marshal(b, m, deterministic)
}
func (m *ApplyRulesCancelReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRulesCancelReq.Merge(m, src)
}
func (m *ApplyRulesCancelReq) XXX_Size() int {
	return xxx_messageInfo_ApplyRulesCancelReq.Size(m)
}
func (m *ApplyRulesCancelReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRulesCancelReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRulesCancelReq proto.InternalMessageInfo

type ApplyRulesStatusReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ApplyRulesStatusReq) Reset()         { *m = ApplyRulesStatusReq{} }
func (m *ApplyRulesStatusReq) String() string { return proto.CompactTextString(m) }
func (*ApplyRulesStatusReq) ProtoMessage()    {}
func (*ApplyRulesStatusReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d81370c9c0e6e2e, []int{8}
}

func (m *ApplyRulesStatusReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplyRulesStatusReq.Unmarshal(m, b)
}
func (m *ApplyRulesStatusReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplyRulesStatusReq.Marshal(b, m, deterministic)
}
func (m *ApplyRulesStatusReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplyRulesStatusReq.Merge(m, src)
}
func (m *ApplyRulesStatusReq) XXX_Size() int {
	return xxx_messageInfo_ApplyRulesStatusReq.Size(m)
}
func (m *ApplyRulesStatusReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplyRulesStatusReq.DiscardUnknown(m)
}

var xxx_messageInfo_ApplyRulesStatusReq proto.InternalMessageInfo

func init() {
	proto.RegisterType((*CreateRuleReq)(nil), "chef.automate.api.iam.v2.CreateRuleReq")
	proto.RegisterType((*UpdateRuleReq)(nil), "chef.automate.api.iam.v2.UpdateRuleReq")
	proto.RegisterType((*GetRuleReq)(nil), "chef.automate.api.iam.v2.GetRuleReq")
	proto.RegisterType((*ListRulesReq)(nil), "chef.automate.api.iam.v2.ListRulesReq")
	proto.RegisterType((*ListRulesForProjectReq)(nil), "chef.automate.api.iam.v2.ListRulesForProjectReq")
	proto.RegisterType((*DeleteRuleReq)(nil), "chef.automate.api.iam.v2.DeleteRuleReq")
	proto.RegisterType((*ApplyRulesStartReq)(nil), "chef.automate.api.iam.v2.ApplyRulesStartReq")
	proto.RegisterType((*ApplyRulesCancelReq)(nil), "chef.automate.api.iam.v2.ApplyRulesCancelReq")
	proto.RegisterType((*ApplyRulesStatusReq)(nil), "chef.automate.api.iam.v2.ApplyRulesStatusReq")
}

func init() {
	proto.RegisterFile("components/automate-gateway/api/iam/v2/request/rules.proto", fileDescriptor_2d81370c9c0e6e2e)
}

var fileDescriptor_2d81370c9c0e6e2e = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0x41, 0x6b, 0xd4, 0x40,
	0x18, 0x65, 0xb2, 0xad, 0xd2, 0xd1, 0xf6, 0x10, 0xab, 0x86, 0x82, 0x10, 0xc6, 0x4b, 0x0f, 0x4d,
	0x02, 0xab, 0x08, 0xe6, 0x96, 0x6e, 0xd3, 0x5a, 0x68, 0xad, 0x66, 0x5b, 0x05, 0xa5, 0x2c, 0xb3,
	0xc9, 0xe7, 0x76, 0x24, 0xc9, 0xcc, 0x26, 0x93, 0xad, 0x4b, 0xe9, 0xd5, 0x83, 0xc7, 0xfd, 0x45,
	0x82, 0x47, 0x4f, 0x2e, 0x1e, 0xfd, 0x1b, 0x1e, 0xbc, 0xc9, 0xcc, 0xee, 0xb6, 0x2b, 0x75, 0xc1,
	0x96, 0x9e, 0x3c, 0x25, 0xf3, 0xbd, 0xf7, 0xbd, 0xef, 0xcd, 0xbc, 0x61, 0xb0, 0x1f, 0xf3, 0x4c,
	0xf0, 0x1c, 0x72, 0x59, 0x7a, 0xb4, 0x92, 0x3c, 0xa3, 0x12, 0x9c, 0x0e, 0x95, 0x70, 0x4c, 0xfb,
	0x1e, 0x15, 0xcc, 0x63, 0x34, 0xf3, 0x7a, 0x75, 0xaf, 0x80, 0x6e, 0x05, 0xa5, 0xf4, 0x8a, 0x2a,
	0x85, 0xd2, 0x15, 0x05, 0x97, 0xdc, 0xb4, 0xe2, 0x23, 0x78, 0xe7, 0x4e, 0xba, 0x5c, 0x2a, 0x98,
	0xcb, 0x68, 0xe6, 0xf6, 0xea, 0x2b, 0x4f, 0xff, 0x51, 0x35, 0xe6, 0x59, 0xc6, 0xf3, 0x69, 0xd1,
	0x95, 0x35, 0xfd, 0x89, 0x9d, 0x0e, 0xe4, 0x4e, 0x79, 0x4c, 0x3b, 0x1d, 0x28, 0x3c, 0x2e, 0x24,
	0xe3, 0x79, 0xe9, 0xd1, 0x3c, 0xe7, 0x92, 0xea, 0xff, 0x11, 0x9b, 0xfc, 0xa8, 0xe1, 0xc5, 0x46,
	0x01, 0x54, 0x42, 0x54, 0xa5, 0x10, 0x41, 0xd7, 0x5c, 0xc2, 0x06, 0x4b, 0x2c, 0x64, 0xa3, 0xd5,
	0x85, 0xc8, 0x60, 0x89, 0xf9, 0x00, 0x63, 0x51, 0xf0, 0xf7, 0x10, 0xcb, 0x16, 0x4b, 0x2c, 0x43,
	0xd7, 0x17, 0xc6, 0x95, 0xed, 0xc4, 0x34, 0xf1, 0x5c, 0x4e, 0x33, 0xb0, 0x6a, 0x1a, 0xd0, 0xff,
	0xe6, 0x13, 0x3c, 0x27, 0xfb, 0x02, 0xac, 0x39, 0x1b, 0xad, 0x2e, 0xd5, 0x89, 0x3b, 0x6b, 0x9b,
	0xae, 0x9a, 0xb9, 0xdf, 0x17, 0x10, 0x69, 0xbe, 0xd9, 0xc0, 0x38, 0xe6, 0x79, 0xc2, 0xb4, 0x41,
	0x6b, 0xde, 0xae, 0xad, 0xde, 0xaa, 0x3f, 0x9c, 0xdd, 0xdd, 0x98, 0x70, 0xa3, 0xa9, 0x36, 0xff,
	0x93, 0x31, 0x08, 0x3e, 0x1a, 0xd8, 0x19, 0x22, 0x83, 0x25, 0x43, 0x34, 0x65, 0x7e, 0x88, 0xb4,
	0xbd, 0x21, 0xd2, 0xd3, 0x86, 0x68, 0xaa, 0xab, 0xfe, 0x1d, 0x99, 0xdf, 0xd0, 0x09, 0x61, 0x09,
	0xf1, 0x6d, 0x02, 0x1f, 0x68, 0x26, 0x52, 0x70, 0xd4, 0xb1, 0x92, 0x35, 0x9b, 0x8c, 0x25, 0x9c,
	0x3f, 0xd1, 0x71, 0x55, 0x11, 0x94, 0xac, 0x82, 0x76, 0xfb, 0x76, 0x38, 0x42, 0xed, 0x68, 0xdc,
	0xab, 0x46, 0x29, 0xe8, 0xf9, 0xde, 0x46, 0xa8, 0xd6, 0xe7, 0x33, 0x89, 0x6f, 0xbf, 0x3d, 0x21,
	0x54, 0xca, 0x82, 0xb5, 0x2b, 0xa9, 0x49, 0x8d, 0x67, 0xe1, 0x66, 0xab, 0x19, 0x46, 0xaf, 0xc2,
	0x48, 0x71, 0xb9, 0x80, 0x82, 0x4a, 0x5e, 0x68, 0xe9, 0x70, 0x77, 0x3d, 0x8c, 0x5a, 0x7b, 0x9b,
	0x0a, 0xe8, 0xd1, 0xb4, 0x02, 0x2d, 0x30, 0x71, 0xe3, 0xc6, 0x5c, 0x01, 0x93, 0x15, 0xe3, 0xe4,
	0xf0, 0xf4, 0xf0, 0x94, 0x7c, 0xad, 0xe1, 0xc5, 0x03, 0x91, 0xfc, 0x47, 0xf1, 0xfe, 0x42, 0x83,
	0xe0, 0x27, 0xba, 0x6c, 0xbc, 0x5f, 0x90, 0xf9, 0xf9, 0x5a, 0xe2, 0x1d, 0x1d, 0x66, 0x72, 0xed,
	0xf1, 0x86, 0x2f, 0x0f, 0x82, 0x9d, 0xe6, 0xec, 0x6c, 0x75, 0x9a, 0xfb, 0x18, 0x6f, 0x81, 0xbc,
	0x5a, 0x92, 0xfe, 0xfd, 0x41, 0xb0, 0x8c, 0xcd, 0x8b, 0xc7, 0x46, 0x96, 0xf0, 0xed, 0x1d, 0x56,
	0x6a, 0xd9, 0x32, 0x82, 0x2e, 0x79, 0x8c, 0xef, 0x9d, 0xad, 0x37, 0x79, 0xf1, 0x62, 0xc4, 0xfc,
	0xcb, 0x44, 0x1f, 0x0f, 0x82, 0x9b, 0x78, 0x5e, 0x4b, 0x92, 0xd7, 0x78, 0x71, 0x03, 0x52, 0xb8,
	0xea, 0x45, 0x9b, 0x6d, 0x6f, 0x19, 0x9b, 0x81, 0x10, 0x69, 0x5f, 0xfb, 0x69, 0x4a, 0x5a, 0x28,
	0x2b, 0xe4, 0x2e, 0xbe, 0x73, 0x5e, 0x6d, 0xd0, 0x3c, 0x86, 0xf4, 0x42, 0xb9, 0x29, 0xa9, 0xac,
	0xd4, 0x96, 0xd6, 0xb7, 0xdf, 0x6c, 0x75, 0x98, 0x3c, 0xaa, 0xda, 0x6e, 0xcc, 0x33, 0x4f, 0xdd,
	0xb8, 0xb3, 0x57, 0xd5, 0xbb, 0xdc, 0xfb, 0xdd, 0xbe, 0xa1, 0xdf, 0xcd, 0x47, 0xbf, 0x03, 0x00,
	0x00, 0xff, 0xff, 0x57, 0x4d, 0x85, 0x84, 0xf8, 0x05, 0x00, 0x00,
}
