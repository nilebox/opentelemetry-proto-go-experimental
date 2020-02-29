// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/opentelemetry/proto/common/v1/common.proto

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

// ValueType is the enumeration of possible types that value can have.
type AttributeKeyValue_ValueType int32

const (
	AttributeKeyValue_STRING AttributeKeyValue_ValueType = 0
	AttributeKeyValue_INT    AttributeKeyValue_ValueType = 1
	AttributeKeyValue_DOUBLE AttributeKeyValue_ValueType = 2
	AttributeKeyValue_BOOL   AttributeKeyValue_ValueType = 3
)

var AttributeKeyValue_ValueType_name = map[int32]string{
	0: "STRING",
	1: "INT",
	2: "DOUBLE",
	3: "BOOL",
}

var AttributeKeyValue_ValueType_value = map[string]int32{
	"STRING": 0,
	"INT":    1,
	"DOUBLE": 2,
	"BOOL":   3,
}

func (x AttributeKeyValue_ValueType) String() string {
	return proto.EnumName(AttributeKeyValue_ValueType_name, int32(x))
}

func (AttributeKeyValue_ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_28046834ca3959bb, []int{0, 0}
}

// AttributeKeyValue is a key-value pair that is used to store Span attributes, Link
// attributes, etc.
type AttributeKeyValue struct {
	// key part of the key-value pair.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// type of the value.
	Type                 AttributeKeyValue_ValueType `protobuf:"varint,2,opt,name=type,proto3,enum=opentelemetry.proto.common.v1.AttributeKeyValue_ValueType" json:"type,omitempty"`
	StringValue          string                      `protobuf:"bytes,3,opt,name=string_value,json=stringValue,proto3" json:"string_value,omitempty"`
	IntValue             int64                       `protobuf:"varint,4,opt,name=int_value,json=intValue,proto3" json:"int_value,omitempty"`
	DoubleValue          float64                     `protobuf:"fixed64,5,opt,name=double_value,json=doubleValue,proto3" json:"double_value,omitempty"`
	BoolValue            bool                        `protobuf:"varint,6,opt,name=bool_value,json=boolValue,proto3" json:"bool_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *AttributeKeyValue) Reset()         { *m = AttributeKeyValue{} }
func (m *AttributeKeyValue) String() string { return proto.CompactTextString(m) }
func (*AttributeKeyValue) ProtoMessage()    {}
func (*AttributeKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_28046834ca3959bb, []int{0}
}

func (m *AttributeKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttributeKeyValue.Unmarshal(m, b)
}
func (m *AttributeKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttributeKeyValue.Marshal(b, m, deterministic)
}
func (m *AttributeKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttributeKeyValue.Merge(m, src)
}
func (m *AttributeKeyValue) XXX_Size() int {
	return xxx_messageInfo_AttributeKeyValue.Size(m)
}
func (m *AttributeKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttributeKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttributeKeyValue proto.InternalMessageInfo

func (m *AttributeKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AttributeKeyValue) GetType() AttributeKeyValue_ValueType {
	if m != nil {
		return m.Type
	}
	return AttributeKeyValue_STRING
}

func (m *AttributeKeyValue) GetStringValue() string {
	if m != nil {
		return m.StringValue
	}
	return ""
}

func (m *AttributeKeyValue) GetIntValue() int64 {
	if m != nil {
		return m.IntValue
	}
	return 0
}

func (m *AttributeKeyValue) GetDoubleValue() float64 {
	if m != nil {
		return m.DoubleValue
	}
	return 0
}

func (m *AttributeKeyValue) GetBoolValue() bool {
	if m != nil {
		return m.BoolValue
	}
	return false
}

// StringKeyValue is a pair of key/value strings. This is the simpler (and faster) version
// of AttributeKeyValue that only supports string values.
type StringKeyValue struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringKeyValue) Reset()         { *m = StringKeyValue{} }
func (m *StringKeyValue) String() string { return proto.CompactTextString(m) }
func (*StringKeyValue) ProtoMessage()    {}
func (*StringKeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_28046834ca3959bb, []int{1}
}

func (m *StringKeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringKeyValue.Unmarshal(m, b)
}
func (m *StringKeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringKeyValue.Marshal(b, m, deterministic)
}
func (m *StringKeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringKeyValue.Merge(m, src)
}
func (m *StringKeyValue) XXX_Size() int {
	return xxx_messageInfo_StringKeyValue.Size(m)
}
func (m *StringKeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringKeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringKeyValue proto.InternalMessageInfo

func (m *StringKeyValue) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StringKeyValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterEnum("opentelemetry.proto.common.v1.AttributeKeyValue_ValueType", AttributeKeyValue_ValueType_name, AttributeKeyValue_ValueType_value)
	proto.RegisterType((*AttributeKeyValue)(nil), "opentelemetry.proto.common.v1.AttributeKeyValue")
	proto.RegisterType((*StringKeyValue)(nil), "opentelemetry.proto.common.v1.StringKeyValue")
}

func init() {
	proto.RegisterFile("proto/opentelemetry/proto/common/v1/common.proto", fileDescriptor_28046834ca3959bb)
}

var fileDescriptor_28046834ca3959bb = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6b, 0xe2, 0x40,
	0x14, 0xc6, 0x77, 0x12, 0x75, 0xcd, 0x53, 0x24, 0x3b, 0xec, 0x41, 0x58, 0x84, 0xe8, 0x29, 0x17,
	0x93, 0xb5, 0x85, 0x52, 0x7a, 0x6b, 0x6a, 0x29, 0x52, 0x51, 0x89, 0xb6, 0x87, 0x5e, 0x8a, 0x69,
	0x1f, 0xe9, 0xd0, 0x64, 0x26, 0xa4, 0x13, 0x6b, 0xfe, 0xaa, 0xfe, 0x8b, 0x25, 0x33, 0xa1, 0x45,
	0x0a, 0x5e, 0x86, 0x37, 0xbf, 0xef, 0xe3, 0x7b, 0x8f, 0x37, 0x03, 0xff, 0xb3, 0x5c, 0x48, 0xe1,
	0x8b, 0x0c, 0xb9, 0xc4, 0x04, 0x53, 0x94, 0x79, 0xe9, 0x6b, 0xf6, 0x24, 0xd2, 0x54, 0x70, 0x7f,
	0x37, 0xa9, 0x2b, 0x4f, 0x61, 0x3a, 0x38, 0xf0, 0x6a, 0xe8, 0xd5, 0x8e, 0xdd, 0x64, 0xf4, 0x61,
	0xc0, 0x9f, 0x4b, 0x29, 0x73, 0x16, 0x15, 0x12, 0x6f, 0xb1, 0xbc, 0xdf, 0x26, 0x05, 0x52, 0x1b,
	0xcc, 0x57, 0x2c, 0xfb, 0xc4, 0x21, 0xae, 0x15, 0x56, 0x25, 0x5d, 0x40, 0x43, 0x96, 0x19, 0xf6,
	0x0d, 0x87, 0xb8, 0xbd, 0x93, 0x0b, 0xef, 0x68, 0xaa, 0xf7, 0x23, 0xd1, 0x53, 0xe7, 0xa6, 0xcc,
	0x30, 0x54, 0x39, 0x74, 0x08, 0xdd, 0x37, 0x99, 0x33, 0x1e, 0x3f, 0xee, 0x2a, 0xa5, 0x6f, 0xaa,
	0x56, 0x1d, 0xcd, 0xf4, 0x10, 0xff, 0xc0, 0x62, 0x5c, 0xd6, 0x7a, 0xc3, 0x21, 0xae, 0x19, 0xb6,
	0x19, 0x97, 0x5a, 0x1c, 0x42, 0xf7, 0x59, 0x14, 0x51, 0x82, 0xb5, 0xde, 0x74, 0x88, 0x4b, 0xc2,
	0x8e, 0x66, 0xda, 0x32, 0x00, 0x88, 0x84, 0x48, 0x6a, 0x43, 0xcb, 0x21, 0x6e, 0x3b, 0xb4, 0x2a,
	0xa2, 0xe4, 0xd1, 0x19, 0x58, 0x5f, 0x43, 0x51, 0x80, 0xd6, 0x7a, 0x13, 0xce, 0x16, 0x37, 0xf6,
	0x2f, 0xfa, 0x1b, 0xcc, 0xd9, 0x62, 0x63, 0x93, 0x0a, 0x4e, 0x97, 0x77, 0xc1, 0xfc, 0xda, 0x36,
	0x68, 0x1b, 0x1a, 0xc1, 0x72, 0x39, 0xb7, 0xcd, 0xd1, 0x39, 0xf4, 0xd6, 0x6a, 0xca, 0x23, 0xdb,
	0xfa, 0x0b, 0x4d, 0xdd, 0xd5, 0x50, 0x4c, 0x5f, 0x82, 0x77, 0x70, 0x98, 0x38, 0xbe, 0xb9, 0xa0,
	0x73, 0xa5, 0xca, 0x55, 0x85, 0x57, 0xe4, 0x61, 0x1a, 0x33, 0xf9, 0x52, 0x44, 0x95, 0xc1, 0xe7,
	0x2c, 0xc1, 0x48, 0xec, 0x0f, 0x1f, 0x7f, 0xac, 0x02, 0xc6, 0xb1, 0x18, 0xe3, 0x3e, 0xc3, 0x9c,
	0xa5, 0xc8, 0xe5, 0x36, 0xf1, 0x63, 0xe4, 0xdf, 0x1f, 0x22, 0x6a, 0x29, 0xd3, 0xe9, 0x67, 0x00,
	0x00, 0x00, 0xff, 0xff, 0xac, 0xdf, 0x76, 0x17, 0x3e, 0x02, 0x00, 0x00,
}
