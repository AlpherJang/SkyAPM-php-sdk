// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/common.proto

package v1

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

// In most cases, detect point should be `server` or `client`.
// Even in service mesh, this means `server`/`client` side sidecar
// `proxy` is reserved only.
type DetectPointV1 int32

const (
	DetectPointV1_client DetectPointV1 = 0
	DetectPointV1_server DetectPointV1 = 1
	DetectPointV1_proxy  DetectPointV1 = 2
)

var DetectPointV1_name = map[int32]string{
	0: "client",
	1: "server",
	2: "proxy",
}

var DetectPointV1_value = map[string]int32{
	"client": 0,
	"server": 1,
	"proxy":  2,
}

func (x DetectPointV1) String() string {
	return proto.EnumName(DetectPointV1_name, int32(x))
}

func (DetectPointV1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

type KeyStringValuePairV1 struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyStringValuePairV1) Reset()         { *m = KeyStringValuePairV1{} }
func (m *KeyStringValuePairV1) String() string { return proto.CompactTextString(m) }
func (*KeyStringValuePairV1) ProtoMessage()    {}
func (*KeyStringValuePairV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{0}
}

func (m *KeyStringValuePairV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyStringValuePairV1.Unmarshal(m, b)
}
func (m *KeyStringValuePairV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyStringValuePairV1.Marshal(b, m, deterministic)
}
func (m *KeyStringValuePairV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyStringValuePairV1.Merge(m, src)
}
func (m *KeyStringValuePairV1) XXX_Size() int {
	return xxx_messageInfo_KeyStringValuePairV1.Size(m)
}
func (m *KeyStringValuePairV1) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyStringValuePairV1.DiscardUnknown(m)
}

var xxx_messageInfo_KeyStringValuePairV1 proto.InternalMessageInfo

func (m *KeyStringValuePairV1) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyStringValuePairV1) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KeyIntValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyIntValuePair) Reset()         { *m = KeyIntValuePair{} }
func (m *KeyIntValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyIntValuePair) ProtoMessage()    {}
func (*KeyIntValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{1}
}

func (m *KeyIntValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyIntValuePair.Unmarshal(m, b)
}
func (m *KeyIntValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyIntValuePair.Marshal(b, m, deterministic)
}
func (m *KeyIntValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyIntValuePair.Merge(m, src)
}
func (m *KeyIntValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyIntValuePair.Size(m)
}
func (m *KeyIntValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyIntValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyIntValuePair proto.InternalMessageInfo

func (m *KeyIntValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyIntValuePair) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

type CPUV1 struct {
	UsagePercent         float64  `protobuf:"fixed64,2,opt,name=usagePercent,proto3" json:"usagePercent,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPUV1) Reset()         { *m = CPUV1{} }
func (m *CPUV1) String() string { return proto.CompactTextString(m) }
func (*CPUV1) ProtoMessage()    {}
func (*CPUV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{2}
}

func (m *CPUV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPUV1.Unmarshal(m, b)
}
func (m *CPUV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPUV1.Marshal(b, m, deterministic)
}
func (m *CPUV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPUV1.Merge(m, src)
}
func (m *CPUV1) XXX_Size() int {
	return xxx_messageInfo_CPUV1.Size(m)
}
func (m *CPUV1) XXX_DiscardUnknown() {
	xxx_messageInfo_CPUV1.DiscardUnknown(m)
}

var xxx_messageInfo_CPUV1 proto.InternalMessageInfo

func (m *CPUV1) GetUsagePercent() float64 {
	if m != nil {
		return m.UsagePercent
	}
	return 0
}

type CommandsV1 struct {
	Commands             []*CommandV1 `protobuf:"bytes,1,rep,name=commands,proto3" json:"commands,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CommandsV1) Reset()         { *m = CommandsV1{} }
func (m *CommandsV1) String() string { return proto.CompactTextString(m) }
func (*CommandsV1) ProtoMessage()    {}
func (*CommandsV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{3}
}

func (m *CommandsV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandsV1.Unmarshal(m, b)
}
func (m *CommandsV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandsV1.Marshal(b, m, deterministic)
}
func (m *CommandsV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandsV1.Merge(m, src)
}
func (m *CommandsV1) XXX_Size() int {
	return xxx_messageInfo_CommandsV1.Size(m)
}
func (m *CommandsV1) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandsV1.DiscardUnknown(m)
}

var xxx_messageInfo_CommandsV1 proto.InternalMessageInfo

func (m *CommandsV1) GetCommands() []*CommandV1 {
	if m != nil {
		return m.Commands
	}
	return nil
}

type CommandV1 struct {
	Command              string                  `protobuf:"bytes,1,opt,name=command,proto3" json:"command,omitempty"`
	Args                 []*KeyStringValuePairV1 `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CommandV1) Reset()         { *m = CommandV1{} }
func (m *CommandV1) String() string { return proto.CompactTextString(m) }
func (*CommandV1) ProtoMessage()    {}
func (*CommandV1) Descriptor() ([]byte, []int) {
	return fileDescriptor_8f954d82c0b891f6, []int{4}
}

func (m *CommandV1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommandV1.Unmarshal(m, b)
}
func (m *CommandV1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommandV1.Marshal(b, m, deterministic)
}
func (m *CommandV1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommandV1.Merge(m, src)
}
func (m *CommandV1) XXX_Size() int {
	return xxx_messageInfo_CommandV1.Size(m)
}
func (m *CommandV1) XXX_DiscardUnknown() {
	xxx_messageInfo_CommandV1.DiscardUnknown(m)
}

var xxx_messageInfo_CommandV1 proto.InternalMessageInfo

func (m *CommandV1) GetCommand() string {
	if m != nil {
		return m.Command
	}
	return ""
}

func (m *CommandV1) GetArgs() []*KeyStringValuePairV1 {
	if m != nil {
		return m.Args
	}
	return nil
}

func init() {
	proto.RegisterEnum("DetectPointV1", DetectPointV1_name, DetectPointV1_value)
	proto.RegisterType((*KeyStringValuePairV1)(nil), "KeyStringValuePairV1")
	proto.RegisterType((*KeyIntValuePair)(nil), "KeyIntValuePair")
	proto.RegisterType((*CPUV1)(nil), "CPUV1")
	proto.RegisterType((*CommandsV1)(nil), "CommandsV1")
	proto.RegisterType((*CommandV1)(nil), "CommandV1")
}

func init() { proto.RegisterFile("common/common.proto", fileDescriptor_8f954d82c0b891f6) }

var fileDescriptor_8f954d82c0b891f6 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xeb, 0x40,
	0x14, 0x85, 0x5f, 0xd2, 0x97, 0xbe, 0xd7, 0xfb, 0x9e, 0x18, 0xc6, 0x0a, 0x41, 0x5c, 0x94, 0x2c,
	0xa4, 0x2a, 0x4c, 0x49, 0x75, 0xe3, 0xc6, 0x85, 0x75, 0x23, 0x05, 0x19, 0x52, 0x8c, 0xe0, 0x6e,
	0x1c, 0x2f, 0x31, 0x24, 0x99, 0x29, 0x93, 0x69, 0x6b, 0xfe, 0x92, 0xbf, 0x52, 0x92, 0x4c, 0x2b,
	0x82, 0xe0, 0x6a, 0xee, 0x3d, 0xdf, 0x39, 0xc3, 0x30, 0x07, 0x0e, 0x84, 0x2a, 0x4b, 0x25, 0x27,
	0xdd, 0x41, 0x97, 0x5a, 0x19, 0x15, 0x5e, 0xc3, 0x70, 0x8e, 0xf5, 0xc2, 0xe8, 0x4c, 0xa6, 0x09,
	0x2f, 0x56, 0xc8, 0x78, 0xa6, 0x93, 0x88, 0xf8, 0xd0, 0xcb, 0xb1, 0x0e, 0x9c, 0x91, 0x33, 0x1e,
	0xc4, 0xcd, 0x48, 0x86, 0xe0, 0xad, 0x1b, 0x43, 0xe0, 0xb6, 0x5a, 0xb7, 0x84, 0x57, 0xb0, 0x3f,
	0xc7, 0xfa, 0x4e, 0x9a, 0x5d, 0xf8, 0xa7, 0xa8, 0xb7, 0x8d, 0x9e, 0x83, 0x37, 0x63, 0x0f, 0x49,
	0x44, 0x42, 0xf8, 0xbf, 0xaa, 0x78, 0x8a, 0x0c, 0xb5, 0x40, 0x69, 0x5a, 0x97, 0x13, 0x7f, 0xd1,
	0xc2, 0x4b, 0x80, 0x99, 0x2a, 0x4b, 0x2e, 0x5f, 0xaa, 0x24, 0x22, 0x27, 0xf0, 0x57, 0xd8, 0x2d,
	0x70, 0x46, 0xbd, 0xf1, 0xbf, 0x29, 0x50, 0x8b, 0x93, 0x28, 0xde, 0xb1, 0x90, 0xc1, 0x60, 0x27,
	0x93, 0x00, 0xfe, 0x58, 0x60, 0xdf, 0xb6, 0x5d, 0xc9, 0x29, 0xfc, 0xe6, 0x3a, 0xad, 0x02, 0xb7,
	0xbd, 0xea, 0x90, 0x7e, 0xf7, 0x23, 0x71, 0x6b, 0x39, 0x9b, 0xc2, 0xde, 0x2d, 0x1a, 0x14, 0x86,
	0xa9, 0x4c, 0x9a, 0x24, 0x22, 0x00, 0x7d, 0x51, 0x64, 0x28, 0x8d, 0xff, 0xab, 0x99, 0x2b, 0xd4,
	0x6b, 0xd4, 0xbe, 0x43, 0x06, 0xe0, 0x2d, 0xb5, 0x7a, 0xab, 0x7d, 0xf7, 0x26, 0x85, 0xb1, 0xd2,
	0x29, 0xe5, 0x4b, 0x2e, 0x5e, 0x91, 0x56, 0x79, 0xbd, 0xe1, 0x45, 0x9e, 0xc9, 0x46, 0x29, 0xa9,
	0x44, 0xb3, 0x51, 0x3a, 0xa7, 0x5d, 0x2b, 0xcc, 0x79, 0x3a, 0xfe, 0x34, 0x4c, 0x2c, 0xb4, 0x95,
	0x4d, 0xd6, 0xd1, 0xbb, 0x7b, 0xb4, 0xc8, 0xeb, 0x47, 0x9b, 0xbf, 0xef, 0x30, 0x6b, 0x9a, 0x14,
	0xaa, 0x78, 0xee, 0xb7, 0x9d, 0x5e, 0x7c, 0x04, 0x00, 0x00, 0xff, 0xff, 0x1d, 0xd7, 0xb0, 0x33,
	0xea, 0x01, 0x00, 0x00,
}
