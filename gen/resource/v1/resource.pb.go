// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/opentelemetry/proto/resource/v1/resource.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/nilebox/opentelemetry-proto-go-experimental/gen/common/v1"
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

// Resource information.
type Resource struct {
	// Set of labels that describe the resource.
	Attributes []*v1.AttributeKeyValue `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty"`
	// dropped_attributes_count is the number of dropped attributes. If the value is 0, then
	// no attributes were dropped.
	DroppedAttributesCount uint32   `protobuf:"varint,2,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"dropped_attributes_count,omitempty"`
	XXX_NoUnkeyedLiteral   struct{} `json:"-"`
	XXX_unrecognized       []byte   `json:"-"`
	XXX_sizecache          int32    `json:"-"`
}

func (m *Resource) Reset()         { *m = Resource{} }
func (m *Resource) String() string { return proto.CompactTextString(m) }
func (*Resource) ProtoMessage()    {}
func (*Resource) Descriptor() ([]byte, []int) {
	return fileDescriptor_c51599a26ae4c3d3, []int{0}
}

func (m *Resource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resource.Unmarshal(m, b)
}
func (m *Resource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resource.Marshal(b, m, deterministic)
}
func (m *Resource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resource.Merge(m, src)
}
func (m *Resource) XXX_Size() int {
	return xxx_messageInfo_Resource.Size(m)
}
func (m *Resource) XXX_DiscardUnknown() {
	xxx_messageInfo_Resource.DiscardUnknown(m)
}

var xxx_messageInfo_Resource proto.InternalMessageInfo

func (m *Resource) GetAttributes() []*v1.AttributeKeyValue {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Resource) GetDroppedAttributesCount() uint32 {
	if m != nil {
		return m.DroppedAttributesCount
	}
	return 0
}

func init() {
	proto.RegisterType((*Resource)(nil), "opentelemetry.proto.resource.v1.Resource")
}

func init() {
	proto.RegisterFile("proto/opentelemetry/proto/resource/v1/resource.proto", fileDescriptor_c51599a26ae4c3d3)
}

var fileDescriptor_c51599a26ae4c3d3 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x32, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x2f, 0x48, 0xcd, 0x2b, 0x49, 0xcd, 0x49, 0xcd, 0x4d, 0x2d, 0x29, 0xaa, 0xd4,
	0x87, 0x88, 0x15, 0xa5, 0x16, 0xe7, 0x97, 0x16, 0x25, 0xa7, 0xea, 0x97, 0x19, 0xc2, 0xd9, 0x7a,
	0x60, 0x29, 0x21, 0x79, 0x14, 0xf5, 0x10, 0x41, 0x3d, 0xb8, 0x9a, 0x32, 0x43, 0x29, 0x2d, 0x6c,
	0x06, 0x26, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0x81, 0x8c, 0x83, 0xb0, 0x20, 0xfa, 0x94, 0xa6, 0x31,
	0x72, 0x71, 0x04, 0x41, 0xf5, 0x0a, 0x05, 0x70, 0x71, 0x25, 0x96, 0x94, 0x14, 0x65, 0x26, 0x95,
	0x96, 0xa4, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x19, 0xe8, 0x61, 0xb3, 0x0e, 0x6a,
	0x46, 0x99, 0xa1, 0x9e, 0x23, 0x4c, 0x83, 0x77, 0x6a, 0x65, 0x58, 0x62, 0x4e, 0x69, 0x6a, 0x10,
	0x92, 0x19, 0x42, 0x16, 0x5c, 0x12, 0x29, 0x45, 0xf9, 0x05, 0x05, 0xa9, 0x29, 0xf1, 0x08, 0xd1,
	0xf8, 0xe4, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xde, 0x20, 0x31, 0xa8, 0x3c,
	0xdc, 0x9c, 0x62, 0x67, 0x90, 0xac, 0x53, 0x2d, 0x97, 0x52, 0x66, 0xbe, 0x1e, 0x01, 0xaf, 0x3a,
	0xf1, 0xc2, 0xdc, 0x1e, 0x00, 0x92, 0x0a, 0x60, 0x8c, 0x72, 0x4b, 0xcf, 0x2c, 0xc9, 0x28, 0x4d,
	0x02, 0x39, 0x50, 0x3f, 0x2f, 0x33, 0x27, 0x35, 0x29, 0xbf, 0x02, 0x35, 0x7c, 0x75, 0xc1, 0x86,
	0xe8, 0xa6, 0xe7, 0xeb, 0xa6, 0x56, 0x14, 0xa4, 0x16, 0x65, 0xe6, 0xa6, 0xe6, 0x95, 0x24, 0xe6,
	0xe8, 0xa7, 0xa7, 0xe6, 0x21, 0x87, 0x79, 0x12, 0x1b, 0x58, 0x99, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0xab, 0x33, 0x08, 0x60, 0xa3, 0x01, 0x00, 0x00,
}
