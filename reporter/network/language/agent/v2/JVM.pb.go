// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/JVM.proto

package v2

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	v2 "github.com/SkyAPM/SkyAPM-php-sdk/reporter/network/common/v2"
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

type PoolTypeV2 int32

const (
	PoolTypeV2_CODE_CACHE_USAGE PoolTypeV2 = 0
	PoolTypeV2_NEWGEN_USAGE     PoolTypeV2 = 1
	PoolTypeV2_OLDGEN_USAGE     PoolTypeV2 = 2
	PoolTypeV2_SURVIVOR_USAGE   PoolTypeV2 = 3
	PoolTypeV2_PERMGEN_USAGE    PoolTypeV2 = 4
	PoolTypeV2_METASPACE_USAGE  PoolTypeV2 = 5
)

var PoolTypeV2_name = map[int32]string{
	0: "CODE_CACHE_USAGE",
	1: "NEWGEN_USAGE",
	2: "OLDGEN_USAGE",
	3: "SURVIVOR_USAGE",
	4: "PERMGEN_USAGE",
	5: "METASPACE_USAGE",
}

var PoolTypeV2_value = map[string]int32{
	"CODE_CACHE_USAGE": 0,
	"NEWGEN_USAGE":     1,
	"OLDGEN_USAGE":     2,
	"SURVIVOR_USAGE":   3,
	"PERMGEN_USAGE":    4,
	"METASPACE_USAGE":  5,
}

func (x PoolTypeV2) String() string {
	return proto.EnumName(PoolTypeV2_name, int32(x))
}

func (PoolTypeV2) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42f5f61b58cf3158, []int{0}
}

type GCPhraseV2 int32

const (
	GCPhraseV2_NEW GCPhraseV2 = 0
	GCPhraseV2_OLD GCPhraseV2 = 1
)

var GCPhraseV2_name = map[int32]string{
	0: "NEW",
	1: "OLD",
}

var GCPhraseV2_value = map[string]int32{
	"NEW": 0,
	"OLD": 1,
}

func (x GCPhraseV2) String() string {
	return proto.EnumName(GCPhraseV2_name, int32(x))
}

func (GCPhraseV2) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_42f5f61b58cf3158, []int{1}
}

type JVMMetricV2 struct {
	Time                 int64           `protobuf:"varint,1,opt,name=time,proto3" json:"time,omitempty"`
	Cpu                  *v2.CPUV2       `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memory               []*MemoryV2     `protobuf:"bytes,3,rep,name=memory,proto3" json:"memory,omitempty"`
	MemoryPool           []*MemoryPoolV2 `protobuf:"bytes,4,rep,name=memoryPool,proto3" json:"memoryPool,omitempty"`
	Gc                   []*GCV2         `protobuf:"bytes,5,rep,name=gc,proto3" json:"gc,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *JVMMetricV2) Reset()         { *m = JVMMetricV2{} }
func (m *JVMMetricV2) String() string { return proto.CompactTextString(m) }
func (*JVMMetricV2) ProtoMessage()    {}
func (*JVMMetricV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_42f5f61b58cf3158, []int{0}
}

func (m *JVMMetricV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JVMMetricV2.Unmarshal(m, b)
}
func (m *JVMMetricV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JVMMetricV2.Marshal(b, m, deterministic)
}
func (m *JVMMetricV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JVMMetricV2.Merge(m, src)
}
func (m *JVMMetricV2) XXX_Size() int {
	return xxx_messageInfo_JVMMetricV2.Size(m)
}
func (m *JVMMetricV2) XXX_DiscardUnknown() {
	xxx_messageInfo_JVMMetricV2.DiscardUnknown(m)
}

var xxx_messageInfo_JVMMetricV2 proto.InternalMessageInfo

func (m *JVMMetricV2) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *JVMMetricV2) GetCpu() *v2.CPUV2 {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *JVMMetricV2) GetMemory() []*MemoryV2 {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *JVMMetricV2) GetMemoryPool() []*MemoryPoolV2 {
	if m != nil {
		return m.MemoryPool
	}
	return nil
}

func (m *JVMMetricV2) GetGc() []*GCV2 {
	if m != nil {
		return m.Gc
	}
	return nil
}

type MemoryV2 struct {
	IsHeap               bool     `protobuf:"varint,1,opt,name=isHeap,proto3" json:"isHeap,omitempty"`
	Init                 int64    `protobuf:"varint,2,opt,name=init,proto3" json:"init,omitempty"`
	Max                  int64    `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Used                 int64    `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
	Committed            int64    `protobuf:"varint,5,opt,name=committed,proto3" json:"committed,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemoryV2) Reset()         { *m = MemoryV2{} }
func (m *MemoryV2) String() string { return proto.CompactTextString(m) }
func (*MemoryV2) ProtoMessage()    {}
func (*MemoryV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_42f5f61b58cf3158, []int{1}
}

func (m *MemoryV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryV2.Unmarshal(m, b)
}
func (m *MemoryV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryV2.Marshal(b, m, deterministic)
}
func (m *MemoryV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryV2.Merge(m, src)
}
func (m *MemoryV2) XXX_Size() int {
	return xxx_messageInfo_MemoryV2.Size(m)
}
func (m *MemoryV2) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryV2.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryV2 proto.InternalMessageInfo

func (m *MemoryV2) GetIsHeap() bool {
	if m != nil {
		return m.IsHeap
	}
	return false
}

func (m *MemoryV2) GetInit() int64 {
	if m != nil {
		return m.Init
	}
	return 0
}

func (m *MemoryV2) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *MemoryV2) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *MemoryV2) GetCommitted() int64 {
	if m != nil {
		return m.Committed
	}
	return 0
}

type MemoryPoolV2 struct {
	Type                 PoolTypeV2 `protobuf:"varint,1,opt,name=type,proto3,enum=PoolTypeV2" json:"type,omitempty"`
	Init                 int64      `protobuf:"varint,2,opt,name=init,proto3" json:"init,omitempty"`
	Max                  int64      `protobuf:"varint,3,opt,name=max,proto3" json:"max,omitempty"`
	Used                 int64      `protobuf:"varint,4,opt,name=used,proto3" json:"used,omitempty"`
	Commited             int64      `protobuf:"varint,5,opt,name=commited,proto3" json:"commited,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *MemoryPoolV2) Reset()         { *m = MemoryPoolV2{} }
func (m *MemoryPoolV2) String() string { return proto.CompactTextString(m) }
func (*MemoryPoolV2) ProtoMessage()    {}
func (*MemoryPoolV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_42f5f61b58cf3158, []int{2}
}

func (m *MemoryPoolV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemoryPoolV2.Unmarshal(m, b)
}
func (m *MemoryPoolV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemoryPoolV2.Marshal(b, m, deterministic)
}
func (m *MemoryPoolV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemoryPoolV2.Merge(m, src)
}
func (m *MemoryPoolV2) XXX_Size() int {
	return xxx_messageInfo_MemoryPoolV2.Size(m)
}
func (m *MemoryPoolV2) XXX_DiscardUnknown() {
	xxx_messageInfo_MemoryPoolV2.DiscardUnknown(m)
}

var xxx_messageInfo_MemoryPoolV2 proto.InternalMessageInfo

func (m *MemoryPoolV2) GetType() PoolTypeV2 {
	if m != nil {
		return m.Type
	}
	return PoolTypeV2_CODE_CACHE_USAGE
}

func (m *MemoryPoolV2) GetInit() int64 {
	if m != nil {
		return m.Init
	}
	return 0
}

func (m *MemoryPoolV2) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *MemoryPoolV2) GetUsed() int64 {
	if m != nil {
		return m.Used
	}
	return 0
}

func (m *MemoryPoolV2) GetCommited() int64 {
	if m != nil {
		return m.Commited
	}
	return 0
}

type GCV2 struct {
	Phrase               GCPhraseV2 `protobuf:"varint,1,opt,name=phrase,proto3,enum=GCPhraseV2" json:"phrase,omitempty"`
	Count                int64      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Time                 int64      `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GCV2) Reset()         { *m = GCV2{} }
func (m *GCV2) String() string { return proto.CompactTextString(m) }
func (*GCV2) ProtoMessage()    {}
func (*GCV2) Descriptor() ([]byte, []int) {
	return fileDescriptor_42f5f61b58cf3158, []int{3}
}

func (m *GCV2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GCV2.Unmarshal(m, b)
}
func (m *GCV2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GCV2.Marshal(b, m, deterministic)
}
func (m *GCV2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GCV2.Merge(m, src)
}
func (m *GCV2) XXX_Size() int {
	return xxx_messageInfo_GCV2.Size(m)
}
func (m *GCV2) XXX_DiscardUnknown() {
	xxx_messageInfo_GCV2.DiscardUnknown(m)
}

var xxx_messageInfo_GCV2 proto.InternalMessageInfo

func (m *GCV2) GetPhrase() GCPhraseV2 {
	if m != nil {
		return m.Phrase
	}
	return GCPhraseV2_NEW
}

func (m *GCV2) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GCV2) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterEnum("PoolTypeV2", PoolTypeV2_name, PoolTypeV2_value)
	proto.RegisterEnum("GCPhraseV2", GCPhraseV2_name, GCPhraseV2_value)
	proto.RegisterType((*JVMMetricV2)(nil), "JVMMetricV2")
	proto.RegisterType((*MemoryV2)(nil), "MemoryV2")
	proto.RegisterType((*MemoryPoolV2)(nil), "MemoryPoolV2")
	proto.RegisterType((*GCV2)(nil), "GCV2")
}

func init() { proto.RegisterFile("common/JVM.proto", fileDescriptor_42f5f61b58cf3158) }

var fileDescriptor_42f5f61b58cf3158 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x97, 0x26, 0x2d, 0xdd, 0xe9, 0x36, 0x82, 0x37, 0x50, 0x54, 0x21, 0x28, 0x85, 0x8b,
	0x6a, 0x12, 0x2e, 0x32, 0x4f, 0x50, 0xd2, 0xa8, 0x63, 0x5a, 0xda, 0xc8, 0x5d, 0x5d, 0x89, 0x9b,
	0x29, 0x64, 0x56, 0x16, 0xb5, 0x89, 0xa3, 0x34, 0xa5, 0x84, 0x4b, 0x24, 0xde, 0x84, 0x2b, 0x9e,
	0x12, 0xd9, 0x49, 0x9b, 0xdd, 0xef, 0x2a, 0xff, 0xf9, 0xce, 0xaf, 0x93, 0xdf, 0xc7, 0x32, 0x98,
	0x81, 0x88, 0x63, 0x91, 0x0c, 0xaf, 0x99, 0x8b, 0xd3, 0x4c, 0xe4, 0xa2, 0x7b, 0x5e, 0x91, 0xf2,
	0x53, 0xc2, 0xfe, 0x5f, 0x0d, 0x3a, 0xd7, 0xcc, 0x75, 0x79, 0x9e, 0x45, 0x01, 0x23, 0x08, 0x81,
	0x91, 0x47, 0x31, 0xb7, 0xb4, 0x9e, 0x36, 0xd0, 0xa9, 0xd2, 0xc8, 0x02, 0x3d, 0x48, 0xb7, 0x56,
	0xa3, 0xa7, 0x0d, 0x3a, 0xa4, 0x85, 0x6d, 0x6f, 0xc1, 0x08, 0x95, 0x08, 0xbd, 0x83, 0x56, 0xcc,
	0x63, 0x91, 0x15, 0x96, 0xde, 0xd3, 0x07, 0x1d, 0x72, 0x8c, 0x5d, 0x55, 0x32, 0x42, 0xab, 0x06,
	0xfa, 0x08, 0x50, 0x2a, 0x4f, 0x88, 0xb5, 0x65, 0x28, 0xdb, 0x69, 0x65, 0x93, 0x88, 0x11, 0xfa,
	0xc8, 0x80, 0x5e, 0x42, 0x23, 0x0c, 0xac, 0xa6, 0xb2, 0x35, 0xf1, 0xc4, 0x66, 0x84, 0x36, 0xc2,
	0xa0, 0xff, 0x0b, 0xda, 0xfb, 0xc9, 0xe8, 0x15, 0xb4, 0xa2, 0xcd, 0x15, 0xf7, 0x53, 0x15, 0xb2,
	0x4d, 0xab, 0x4a, 0x46, 0x8f, 0x92, 0x28, 0x57, 0x39, 0x75, 0xaa, 0x34, 0x32, 0x41, 0x8f, 0xfd,
	0x9f, 0x96, 0xae, 0x90, 0x94, 0xd2, 0xb5, 0xdd, 0xf0, 0x7b, 0xcb, 0x28, 0x5d, 0x52, 0xa3, 0xd7,
	0x70, 0x2c, 0x97, 0x12, 0xe5, 0x39, 0xbf, 0xb7, 0x9a, 0xaa, 0x51, 0x83, 0xfe, 0x1f, 0x0d, 0x4e,
	0x1e, 0xe7, 0x45, 0x6f, 0xc1, 0xc8, 0x8b, 0xb4, 0xdc, 0xd1, 0x19, 0xe9, 0x60, 0x89, 0x6f, 0x8b,
	0x94, 0x33, 0x42, 0x55, 0xe3, 0x09, 0x49, 0xba, 0xd0, 0x2e, 0x7f, 0x7c, 0x08, 0x72, 0xa8, 0xfb,
	0x0b, 0x30, 0xe4, 0x3e, 0xd0, 0x7b, 0x68, 0xa5, 0x0f, 0x99, 0xbf, 0xa9, 0x03, 0x4c, 0x6c, 0x4f,
	0x01, 0xb9, 0xf6, 0xb2, 0x85, 0x2e, 0xa0, 0x19, 0x88, 0x6d, 0xb2, 0xcf, 0x50, 0x16, 0x87, 0xdb,
	0xd5, 0xeb, 0xdb, 0xbd, 0xfc, 0xad, 0x01, 0xd4, 0x27, 0x40, 0x17, 0x60, 0xda, 0xb3, 0xb1, 0x73,
	0x67, 0x8f, 0xec, 0x2b, 0xe7, 0x6e, 0x31, 0x1f, 0x4d, 0x1c, 0xf3, 0x08, 0x99, 0x70, 0x32, 0x75,
	0x96, 0x13, 0x67, 0x5a, 0x11, 0x4d, 0x92, 0xd9, 0xcd, 0xb8, 0x26, 0x0d, 0x84, 0xe0, 0x6c, 0xbe,
	0xa0, 0xec, 0x2b, 0x9b, 0xd1, 0x8a, 0xe9, 0xe8, 0x05, 0x9c, 0x7a, 0x0e, 0x75, 0x6b, 0x9b, 0x81,
	0xce, 0xe1, 0xb9, 0xeb, 0xdc, 0x8e, 0xe6, 0xde, 0xc8, 0xde, 0xcf, 0x6f, 0x5e, 0xbe, 0x01, 0xa8,
	0x0f, 0x81, 0x9e, 0x81, 0x3e, 0x75, 0x96, 0xe6, 0x91, 0x14, 0xb3, 0x9b, 0xb1, 0xa9, 0x7d, 0xd9,
	0xc1, 0x27, 0x91, 0x85, 0xd8, 0x4f, 0xfd, 0xe0, 0x81, 0xe3, 0xcd, 0xaa, 0xd8, 0xf9, 0xeb, 0x55,
	0x94, 0x48, 0x12, 0xe3, 0x84, 0xe7, 0x3b, 0x91, 0xad, 0xf0, 0xda, 0x4f, 0xc2, 0xad, 0x1f, 0x72,
	0xec, 0x87, 0x3c, 0xc9, 0x3d, 0xed, 0xdb, 0x87, 0xda, 0x38, 0xac, 0x4c, 0xc3, 0xbd, 0x69, 0xa8,
	0x4c, 0xc3, 0x1f, 0xe4, 0x5f, 0xa3, 0x3b, 0x5f, 0x15, 0xcb, 0x6a, 0xde, 0xb4, 0xb4, 0x79, 0xf2,
	0x71, 0x04, 0x62, 0xfd, 0xbd, 0xa5, 0x9e, 0xc9, 0xe7, 0xff, 0x01, 0x00, 0x00, 0xff, 0xff, 0x86,
	0x6c, 0x8e, 0xdb, 0x4f, 0x03, 0x00, 0x00,
}
