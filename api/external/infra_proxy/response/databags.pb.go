// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/external/infra_proxy/response/databags.proto

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

type DataBags struct {
	// Data bags item list.
	DataBags             []*DataBagListItem `protobuf:"bytes,2,rep,name=data_bags,json=dataBags,proto3" json:"data_bags,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DataBags) Reset()         { *m = DataBags{} }
func (m *DataBags) String() string { return proto.CompactTextString(m) }
func (*DataBags) ProtoMessage()    {}
func (*DataBags) Descriptor() ([]byte, []int) {
	return fileDescriptor_e518f553399bb7b1, []int{0}
}

func (m *DataBags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBags.Unmarshal(m, b)
}
func (m *DataBags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBags.Marshal(b, m, deterministic)
}
func (m *DataBags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBags.Merge(m, src)
}
func (m *DataBags) XXX_Size() int {
	return xxx_messageInfo_DataBags.Size(m)
}
func (m *DataBags) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBags.DiscardUnknown(m)
}

var xxx_messageInfo_DataBags proto.InternalMessageInfo

func (m *DataBags) GetDataBags() []*DataBagListItem {
	if m != nil {
		return m.DataBags
	}
	return nil
}

type DataBagListItem struct {
	// Data bag item name.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataBagListItem) Reset()         { *m = DataBagListItem{} }
func (m *DataBagListItem) String() string { return proto.CompactTextString(m) }
func (*DataBagListItem) ProtoMessage()    {}
func (*DataBagListItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_e518f553399bb7b1, []int{1}
}

func (m *DataBagListItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBagListItem.Unmarshal(m, b)
}
func (m *DataBagListItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBagListItem.Marshal(b, m, deterministic)
}
func (m *DataBagListItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBagListItem.Merge(m, src)
}
func (m *DataBagListItem) XXX_Size() int {
	return xxx_messageInfo_DataBagListItem.Size(m)
}
func (m *DataBagListItem) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBagListItem.DiscardUnknown(m)
}

var xxx_messageInfo_DataBagListItem proto.InternalMessageInfo

func (m *DataBagListItem) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DataBag struct {
	// Data bag name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Data bag item ID.
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Stringified json of data bag item.
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DataBag) Reset()         { *m = DataBag{} }
func (m *DataBag) String() string { return proto.CompactTextString(m) }
func (*DataBag) ProtoMessage()    {}
func (*DataBag) Descriptor() ([]byte, []int) {
	return fileDescriptor_e518f553399bb7b1, []int{2}
}

func (m *DataBag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataBag.Unmarshal(m, b)
}
func (m *DataBag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataBag.Marshal(b, m, deterministic)
}
func (m *DataBag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataBag.Merge(m, src)
}
func (m *DataBag) XXX_Size() int {
	return xxx_messageInfo_DataBag.Size(m)
}
func (m *DataBag) XXX_DiscardUnknown() {
	xxx_messageInfo_DataBag.DiscardUnknown(m)
}

var xxx_messageInfo_DataBag proto.InternalMessageInfo

func (m *DataBag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataBag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *DataBag) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*DataBags)(nil), "chef.automate.api.infra_proxy.response.DataBags")
	proto.RegisterType((*DataBagListItem)(nil), "chef.automate.api.infra_proxy.response.DataBagListItem")
	proto.RegisterType((*DataBag)(nil), "chef.automate.api.infra_proxy.response.DataBag")
}

func init() {
	proto.RegisterFile("api/external/infra_proxy/response/databags.proto", fileDescriptor_e518f553399bb7b1)
}

var fileDescriptor_e518f553399bb7b1 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0x3f, 0x4b, 0x43, 0x31,
	0x14, 0xc5, 0xe9, 0xab, 0x68, 0x1b, 0x41, 0x21, 0x53, 0xc6, 0xf2, 0x40, 0xe9, 0x74, 0x23, 0x3a,
	0x08, 0xe2, 0xa2, 0xb8, 0x08, 0x4e, 0xc5, 0xc9, 0xa5, 0xde, 0xd7, 0xdc, 0xbe, 0x06, 0xcc, 0x1f,
	0x92, 0x5b, 0xa8, 0xdf, 0x5e, 0x12, 0x5f, 0x41, 0x44, 0xb0, 0x5b, 0xc8, 0x3d, 0xe7, 0x77, 0x0e,
	0x47, 0x5c, 0x61, 0xb4, 0x9a, 0x76, 0x4c, 0xc9, 0xe3, 0x87, 0xb6, 0x7e, 0x9d, 0x70, 0x19, 0x53,
	0xd8, 0x7d, 0xea, 0x44, 0x39, 0x06, 0x9f, 0x49, 0x1b, 0x64, 0xec, 0xb0, 0xcf, 0x10, 0x53, 0xe0,
	0x20, 0x2f, 0x57, 0x1b, 0x5a, 0x03, 0x6e, 0x39, 0x38, 0x64, 0x02, 0x8c, 0x16, 0x7e, 0xd8, 0x60,
	0x6f, 0x6b, 0xdf, 0xc5, 0xe4, 0x09, 0x19, 0x1f, 0xb1, 0xcf, 0xf2, 0x55, 0x4c, 0x0b, 0x65, 0x59,
	0x30, 0xaa, 0x99, 0x8d, 0xe7, 0xa7, 0xd7, 0xb7, 0x70, 0x18, 0x07, 0x06, 0xc8, 0x8b, 0xcd, 0xfc,
	0xcc, 0xe4, 0x16, 0x13, 0x33, 0x50, 0xdb, 0x0b, 0x71, 0xfe, 0xeb, 0x28, 0xa5, 0x38, 0xf2, 0xe8,
	0x48, 0x8d, 0x66, 0xa3, 0xf9, 0x74, 0x51, 0xdf, 0xed, 0x83, 0x38, 0x19, 0x64, 0x7f, 0x9d, 0xe5,
	0x99, 0x68, 0xac, 0x51, 0x4d, 0xfd, 0x69, 0xac, 0x29, 0x9a, 0x92, 0xa0, 0xc6, 0xdf, 0x9a, 0x9a,
	0x76, 0xff, 0x76, 0xd7, 0x5b, 0xde, 0x6c, 0x3b, 0x58, 0x05, 0xa7, 0x4b, 0x71, 0xbd, 0x2f, 0xae,
	0xff, 0x1d, 0xb0, 0x3b, 0xae, 0xc3, 0xdd, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xb5, 0xec,
	0xf8, 0x6c, 0x01, 0x00, 0x00,
}
