// Code generated by protoc-gen-go. DO NOT EDIT.
// source: type.proto

package mpt

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

type PersistNode struct {
	// Types that are valid to be assigned to Content:
	//	*PersistNode_Full
	//	*PersistNode_Short
	//	*PersistNode_Value
	Content              isPersistNode_Content `protobuf_oneof:"Content"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PersistNode) Reset()         { *m = PersistNode{} }
func (m *PersistNode) String() string { return proto.CompactTextString(m) }
func (*PersistNode) ProtoMessage()    {}
func (*PersistNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eaed4801c3a9059, []int{0}
}

func (m *PersistNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistNode.Unmarshal(m, b)
}
func (m *PersistNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistNode.Marshal(b, m, deterministic)
}
func (m *PersistNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistNode.Merge(m, src)
}
func (m *PersistNode) XXX_Size() int {
	return xxx_messageInfo_PersistNode.Size(m)
}
func (m *PersistNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistNode.DiscardUnknown(m)
}

var xxx_messageInfo_PersistNode proto.InternalMessageInfo

type isPersistNode_Content interface {
	isPersistNode_Content()
}

type PersistNode_Full struct {
	Full *PersistFullNode `protobuf:"bytes,1,opt,name=full,proto3,oneof"`
}

type PersistNode_Short struct {
	Short *PersistShortNode `protobuf:"bytes,2,opt,name=short,proto3,oneof"`
}

type PersistNode_Value struct {
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3,oneof"`
}

func (*PersistNode_Full) isPersistNode_Content() {}

func (*PersistNode_Short) isPersistNode_Content() {}

func (*PersistNode_Value) isPersistNode_Content() {}

func (m *PersistNode) GetContent() isPersistNode_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *PersistNode) GetFull() *PersistFullNode {
	if x, ok := m.GetContent().(*PersistNode_Full); ok {
		return x.Full
	}
	return nil
}

func (m *PersistNode) GetShort() *PersistShortNode {
	if x, ok := m.GetContent().(*PersistNode_Short); ok {
		return x.Short
	}
	return nil
}

func (m *PersistNode) GetValue() []byte {
	if x, ok := m.GetContent().(*PersistNode_Value); ok {
		return x.Value
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*PersistNode) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*PersistNode_Full)(nil),
		(*PersistNode_Short)(nil),
		(*PersistNode_Value)(nil),
	}
}

type PersistFullNode struct {
	Children             [][]byte `protobuf:"bytes,1,rep,name=Children,proto3" json:"Children,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersistFullNode) Reset()         { *m = PersistFullNode{} }
func (m *PersistFullNode) String() string { return proto.CompactTextString(m) }
func (*PersistFullNode) ProtoMessage()    {}
func (*PersistFullNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eaed4801c3a9059, []int{1}
}

func (m *PersistFullNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistFullNode.Unmarshal(m, b)
}
func (m *PersistFullNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistFullNode.Marshal(b, m, deterministic)
}
func (m *PersistFullNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistFullNode.Merge(m, src)
}
func (m *PersistFullNode) XXX_Size() int {
	return xxx_messageInfo_PersistFullNode.Size(m)
}
func (m *PersistFullNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistFullNode.DiscardUnknown(m)
}

var xxx_messageInfo_PersistFullNode proto.InternalMessageInfo

func (m *PersistFullNode) GetChildren() [][]byte {
	if m != nil {
		return m.Children
	}
	return nil
}

type PersistShortNode struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersistShortNode) Reset()         { *m = PersistShortNode{} }
func (m *PersistShortNode) String() string { return proto.CompactTextString(m) }
func (*PersistShortNode) ProtoMessage()    {}
func (*PersistShortNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eaed4801c3a9059, []int{2}
}

func (m *PersistShortNode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistShortNode.Unmarshal(m, b)
}
func (m *PersistShortNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistShortNode.Marshal(b, m, deterministic)
}
func (m *PersistShortNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistShortNode.Merge(m, src)
}
func (m *PersistShortNode) XXX_Size() int {
	return xxx_messageInfo_PersistShortNode.Size(m)
}
func (m *PersistShortNode) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistShortNode.DiscardUnknown(m)
}

var xxx_messageInfo_PersistShortNode proto.InternalMessageInfo

func (m *PersistShortNode) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PersistShortNode) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type PersistTrie struct {
	Pairs                []*PersistKV `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PersistTrie) Reset()         { *m = PersistTrie{} }
func (m *PersistTrie) String() string { return proto.CompactTextString(m) }
func (*PersistTrie) ProtoMessage()    {}
func (*PersistTrie) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eaed4801c3a9059, []int{3}
}

func (m *PersistTrie) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistTrie.Unmarshal(m, b)
}
func (m *PersistTrie) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistTrie.Marshal(b, m, deterministic)
}
func (m *PersistTrie) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistTrie.Merge(m, src)
}
func (m *PersistTrie) XXX_Size() int {
	return xxx_messageInfo_PersistTrie.Size(m)
}
func (m *PersistTrie) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistTrie.DiscardUnknown(m)
}

var xxx_messageInfo_PersistTrie proto.InternalMessageInfo

func (m *PersistTrie) GetPairs() []*PersistKV {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type PersistKV struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PersistKV) Reset()         { *m = PersistKV{} }
func (m *PersistKV) String() string { return proto.CompactTextString(m) }
func (*PersistKV) ProtoMessage()    {}
func (*PersistKV) Descriptor() ([]byte, []int) {
	return fileDescriptor_8eaed4801c3a9059, []int{4}
}

func (m *PersistKV) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PersistKV.Unmarshal(m, b)
}
func (m *PersistKV) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PersistKV.Marshal(b, m, deterministic)
}
func (m *PersistKV) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PersistKV.Merge(m, src)
}
func (m *PersistKV) XXX_Size() int {
	return xxx_messageInfo_PersistKV.Size(m)
}
func (m *PersistKV) XXX_DiscardUnknown() {
	xxx_messageInfo_PersistKV.DiscardUnknown(m)
}

var xxx_messageInfo_PersistKV proto.InternalMessageInfo

func (m *PersistKV) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *PersistKV) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func init() {
	proto.RegisterType((*PersistNode)(nil), "mpt.PersistNode")
	proto.RegisterType((*PersistFullNode)(nil), "mpt.PersistFullNode")
	proto.RegisterType((*PersistShortNode)(nil), "mpt.PersistShortNode")
	proto.RegisterType((*PersistTrie)(nil), "mpt.PersistTrie")
	proto.RegisterType((*PersistKV)(nil), "mpt.PersistKV")
}

func init() { proto.RegisterFile("type.proto", fileDescriptor_8eaed4801c3a9059) }

var fileDescriptor_8eaed4801c3a9059 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xdd, 0xc6, 0xa8, 0x9d, 0x04, 0x2d, 0x43, 0x95, 0xe0, 0x29, 0x2c, 0x1e, 0x82, 0xd0,
	0x1c, 0x9a, 0x9b, 0x47, 0x0b, 0x22, 0x14, 0x44, 0xa2, 0x78, 0xaf, 0x74, 0xa4, 0xa1, 0xdb, 0xec,
	0xb2, 0x99, 0x08, 0xfd, 0x0d, 0xfe, 0x69, 0xd9, 0x4d, 0x8d, 0xa5, 0xd0, 0x5b, 0xde, 0x7b, 0x1f,
	0x7c, 0x99, 0x05, 0xe0, 0xad, 0xa1, 0xdc, 0x58, 0xcd, 0x1a, 0x83, 0x8d, 0x61, 0xf9, 0x23, 0x20,
	0x7a, 0x25, 0xdb, 0x54, 0x0d, 0xbf, 0xe8, 0x25, 0xe1, 0x3d, 0x9c, 0x7e, 0xb5, 0x4a, 0x25, 0x22,
	0x15, 0x59, 0x34, 0x1d, 0xe7, 0x1b, 0xc3, 0xf9, 0x6e, 0x7f, 0x6a, 0x95, 0x72, 0xcc, 0xf3, 0x49,
	0xe9, 0x19, 0x9c, 0x40, 0xd8, 0xac, 0xb4, 0xe5, 0x64, 0xe0, 0xe1, 0xeb, 0x7d, 0xf8, 0xcd, 0x0d,
	0x3b, 0xba, 0xa3, 0xf0, 0x06, 0xc2, 0xef, 0x85, 0x6a, 0x29, 0x09, 0x52, 0x91, 0xc5, 0xae, 0xf7,
	0xf1, 0x71, 0x08, 0xe7, 0x33, 0x5d, 0x33, 0xd5, 0x2c, 0x27, 0x70, 0x75, 0x20, 0xc3, 0x5b, 0xb8,
	0x98, 0xad, 0x2a, 0xb5, 0xb4, 0x54, 0x27, 0x22, 0x0d, 0xb2, 0xb8, 0xec, 0xb3, 0x7c, 0x80, 0xd1,
	0xa1, 0x0e, 0x47, 0x10, 0xcc, 0x69, 0xeb, 0xff, 0x3f, 0x2e, 0xdd, 0x27, 0x8e, 0xff, 0xbc, 0x03,
	0xdf, 0x75, 0x41, 0x16, 0xfd, 0xdd, 0xef, 0xb6, 0x22, 0xbc, 0x83, 0xd0, 0x2c, 0x2a, 0xdb, 0x78,
	0x47, 0x34, 0xbd, 0xdc, 0xbf, 0x65, 0xfe, 0x51, 0x76, 0xa3, 0x2c, 0x60, 0xd8, 0x77, 0xce, 0xb4,
	0xfe, 0x37, 0xad, 0x8f, 0x99, 0x3e, 0xcf, 0xfc, 0x73, 0x17, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xe4, 0x00, 0x0f, 0x60, 0x7c, 0x01, 0x00, 0x00,
}
