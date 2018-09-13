// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto2/enum.proto

package proto2

import proto "github.com/golang/protobuf/proto"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// EnumType1 comment.
type EnumType1 int32

const (
	// EnumType1_ONE comment.
	EnumType1_ONE EnumType1 = 1
	// EnumType1_TWO comment.
	EnumType1_TWO EnumType1 = 2
)

var EnumType1_name = map[int32]string{
	1: "ONE",
	2: "TWO",
}

var EnumType1_value = map[string]int32{
	"ONE": 1,
	"TWO": 2,
}

func (x EnumType1) Enum() *EnumType1 {
	p := new(EnumType1)
	*p = x
	return p
}

func (x EnumType1) String() string {
	return proto.EnumName(EnumType1_name, int32(x))
}

func (x *EnumType1) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumType1_value, data, "EnumType1")
	if err != nil {
		return err
	}
	*x = EnumType1(value)
	return nil
}

func (EnumType1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0}
}

type EnumType2 int32

const (
	EnumType2_duplicate1 EnumType2 = 1
	EnumType2_duplicate2 EnumType2 = 1
)

var EnumType2_name = map[int32]string{
	1: "duplicate1",
	// Duplicate value: 1: "duplicate2",
}

var EnumType2_value = map[string]int32{
	"duplicate1": 1,
	"duplicate2": 1,
}

func (x EnumType2) Enum() *EnumType2 {
	p := new(EnumType2)
	*p = x
	return p
}

func (x EnumType2) String() string {
	return proto.EnumName(EnumType2_name, int32(x))
}

func (x *EnumType2) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumType2_value, data, "EnumType2")
	if err != nil {
		return err
	}
	*x = EnumType2(value)
	return nil
}

func (EnumType2) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{1}
}

// NestedEnumType1A comment.
type EnumContainerMessage1_NestedEnumType1A int32

const (
	// NestedEnumType1A_VALUE comment.
	EnumContainerMessage1_NESTED_1A_VALUE EnumContainerMessage1_NestedEnumType1A = 0
)

var EnumContainerMessage1_NestedEnumType1A_name = map[int32]string{
	0: "NESTED_1A_VALUE",
}

var EnumContainerMessage1_NestedEnumType1A_value = map[string]int32{
	"NESTED_1A_VALUE": 0,
}

func (x EnumContainerMessage1_NestedEnumType1A) Enum() *EnumContainerMessage1_NestedEnumType1A {
	p := new(EnumContainerMessage1_NestedEnumType1A)
	*p = x
	return p
}

func (x EnumContainerMessage1_NestedEnumType1A) String() string {
	return proto.EnumName(EnumContainerMessage1_NestedEnumType1A_name, int32(x))
}

func (x *EnumContainerMessage1_NestedEnumType1A) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumContainerMessage1_NestedEnumType1A_value, data, "EnumContainerMessage1_NestedEnumType1A")
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_NestedEnumType1A(value)
	return nil
}

func (EnumContainerMessage1_NestedEnumType1A) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 0}
}

type EnumContainerMessage1_NestedEnumType1B int32

const (
	EnumContainerMessage1_NESTED_1B_VALUE EnumContainerMessage1_NestedEnumType1B = 0
)

var EnumContainerMessage1_NestedEnumType1B_name = map[int32]string{
	0: "NESTED_1B_VALUE",
}

var EnumContainerMessage1_NestedEnumType1B_value = map[string]int32{
	"NESTED_1B_VALUE": 0,
}

func (x EnumContainerMessage1_NestedEnumType1B) Enum() *EnumContainerMessage1_NestedEnumType1B {
	p := new(EnumContainerMessage1_NestedEnumType1B)
	*p = x
	return p
}

func (x EnumContainerMessage1_NestedEnumType1B) String() string {
	return proto.EnumName(EnumContainerMessage1_NestedEnumType1B_name, int32(x))
}

func (x *EnumContainerMessage1_NestedEnumType1B) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumContainerMessage1_NestedEnumType1B_value, data, "EnumContainerMessage1_NestedEnumType1B")
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_NestedEnumType1B(value)
	return nil
}

func (EnumContainerMessage1_NestedEnumType1B) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 1}
}

// NestedEnumType2A comment.
type EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A int32

const (
	// NestedEnumType2A_VALUE comment.
	EnumContainerMessage1_EnumContainerMessage2_NESTED_2A_VALUE EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A = 0
)

var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_name = map[int32]string{
	0: "NESTED_2A_VALUE",
}

var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_value = map[string]int32{
	"NESTED_2A_VALUE": 0,
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) Enum() *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A {
	p := new(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A)
	*p = x
	return p
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) String() string {
	return proto.EnumName(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_name, int32(x))
}

func (x *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_value, data, "EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A")
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A(value)
	return nil
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 0, 0}
}

type EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B int32

const (
	EnumContainerMessage1_EnumContainerMessage2_NESTED_2B_VALUE EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B = 0
)

var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_name = map[int32]string{
	0: "NESTED_2B_VALUE",
}

var EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_value = map[string]int32{
	"NESTED_2B_VALUE": 0,
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) Enum() *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B {
	p := new(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B)
	*p = x
	return p
}

func (x EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) String() string {
	return proto.EnumName(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_name, int32(x))
}

func (x *EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_value, data, "EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B")
	if err != nil {
		return err
	}
	*x = EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B(value)
	return nil
}

func (EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 0, 1}
}

type EnumContainerMessage1 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumContainerMessage1) Reset()         { *m = EnumContainerMessage1{} }
func (m *EnumContainerMessage1) String() string { return proto.CompactTextString(m) }
func (*EnumContainerMessage1) ProtoMessage()    {}
func (*EnumContainerMessage1) Descriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0}
}
func (m *EnumContainerMessage1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumContainerMessage1.Unmarshal(m, b)
}
func (m *EnumContainerMessage1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumContainerMessage1.Marshal(b, m, deterministic)
}
func (m *EnumContainerMessage1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumContainerMessage1.Merge(m, src)
}
func (m *EnumContainerMessage1) XXX_Size() int {
	return xxx_messageInfo_EnumContainerMessage1.Size(m)
}
func (m *EnumContainerMessage1) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumContainerMessage1.DiscardUnknown(m)
}

var xxx_messageInfo_EnumContainerMessage1 proto.InternalMessageInfo

type EnumContainerMessage1_EnumContainerMessage2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnumContainerMessage1_EnumContainerMessage2) Reset() {
	*m = EnumContainerMessage1_EnumContainerMessage2{}
}
func (m *EnumContainerMessage1_EnumContainerMessage2) String() string {
	return proto.CompactTextString(m)
}
func (*EnumContainerMessage1_EnumContainerMessage2) ProtoMessage() {}
func (*EnumContainerMessage1_EnumContainerMessage2) Descriptor() ([]byte, []int) {
	return fileDescriptor_de9f68860d540858, []int{0, 0}
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.Unmarshal(m, b)
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.Marshal(b, m, deterministic)
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.Merge(m, src)
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_Size() int {
	return xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.Size(m)
}
func (m *EnumContainerMessage1_EnumContainerMessage2) XXX_DiscardUnknown() {
	xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2.DiscardUnknown(m)
}

var xxx_messageInfo_EnumContainerMessage1_EnumContainerMessage2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EnumContainerMessage1)(nil), "goproto.protoc.proto2.EnumContainerMessage1")
	proto.RegisterType((*EnumContainerMessage1_EnumContainerMessage2)(nil), "goproto.protoc.proto2.EnumContainerMessage1.EnumContainerMessage2")
	proto.RegisterEnum("goproto.protoc.proto2.EnumType1", EnumType1_name, EnumType1_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumType2", EnumType2_name, EnumType2_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumContainerMessage1_NestedEnumType1A", EnumContainerMessage1_NestedEnumType1A_name, EnumContainerMessage1_NestedEnumType1A_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumContainerMessage1_NestedEnumType1B", EnumContainerMessage1_NestedEnumType1B_name, EnumContainerMessage1_NestedEnumType1B_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A", EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_name, EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2A_value)
	proto.RegisterEnum("goproto.protoc.proto2.EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B", EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_name, EnumContainerMessage1_EnumContainerMessage2_NestedEnumType2B_value)
}

func init() { proto.RegisterFile("proto2/enum.proto", fileDescriptor_de9f68860d540858) }

var fileDescriptor_de9f68860d540858 = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0xcd, 0x39, 0x88, 0x19, 0x34, 0x56, 0x6e, 0x39, 0x70, 0xb9, 0x45, 0x38, 0xb8, 0x86,
	0x64, 0x13, 0xa7, 0x56, 0xb3, 0x69, 0x6f, 0xb0, 0x2a, 0xb8, 0x1c, 0xa1, 0x7d, 0x3c, 0x0a, 0x6d,
	0x5e, 0x69, 0xd3, 0xc1, 0xff, 0xd3, 0x3f, 0x48, 0xae, 0x81, 0xb3, 0x42, 0x75, 0xca, 0xf7, 0xe5,
	0xfb, 0xf1, 0x1b, 0x1e, 0xbf, 0x6a, 0x3b, 0xf2, 0xa4, 0x25, 0xb8, 0xa1, 0x89, 0xc7, 0x1c, 0x2d,
	0x91, 0xc6, 0x10, 0x6a, 0x11, 0x1e, 0xbd, 0xfe, 0x62, 0x7c, 0x69, 0xdc, 0xd0, 0x3c, 0x90, 0xf3,
	0xb6, 0x72, 0xd0, 0x3d, 0x43, 0xdf, 0x5b, 0x04, 0xb5, 0xaa, 0xe6, 0x07, 0xbd, 0xbe, 0xe5, 0x22,
	0x83, 0xde, 0x43, 0x79, 0x98, 0xf3, 0xcf, 0x16, 0x74, 0x12, 0x5d, 0xf3, 0xcb, 0xcc, 0xbc, 0xe4,
	0xe6, 0x71, 0xaf, 0x93, 0xfd, 0x5b, 0xf2, 0xf4, 0x6a, 0xc4, 0xc9, 0x0c, 0x98, 0x4e, 0xc1, 0xf4,
	0x6f, 0x50, 0x4d, 0x8d, 0xea, 0x1f, 0xa3, 0x9a, 0x1a, 0xd5, 0xd1, 0xb8, 0xb9, 0xe1, 0xe7, 0x47,
	0x24, 0x3a, 0xe3, 0xa7, 0xbb, 0xcc, 0x08, 0x76, 0x08, 0xf9, 0xfb, 0x4e, 0x2c, 0x36, 0xf2, 0x67,
	0xd6, 0xd1, 0x05, 0xe7, 0xe5, 0xd0, 0xd6, 0x55, 0x61, 0x3d, 0x28, 0xc1, 0x7e, 0x75, 0x2d, 0xd8,
	0x6a, 0x21, 0x58, 0x7a, 0xff, 0x71, 0x87, 0x44, 0x58, 0x43, 0x8c, 0x54, 0x5b, 0x87, 0x31, 0x75,
	0x28, 0xc7, 0x13, 0xca, 0xa2, 0x29, 0x43, 0x2a, 0xb6, 0x08, 0x6e, 0x8b, 0x24, 0x3d, 0xf4, 0xbe,
	0xb4, 0xde, 0x86, 0x6f, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xf6, 0x53, 0xf7, 0xde, 0x8e, 0x01,
	0x00, 0x00,
}
