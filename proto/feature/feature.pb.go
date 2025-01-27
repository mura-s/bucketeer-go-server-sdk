// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/feature/feature.proto

package feature

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

type Feature_VariationType int32

const (
	Feature_STRING  Feature_VariationType = 0
	Feature_BOOLEAN Feature_VariationType = 1
	Feature_NUMBER  Feature_VariationType = 2
	Feature_JSON    Feature_VariationType = 3
)

var Feature_VariationType_name = map[int32]string{
	0: "STRING",
	1: "BOOLEAN",
	2: "NUMBER",
	3: "JSON",
}

var Feature_VariationType_value = map[string]int32{
	"STRING":  0,
	"BOOLEAN": 1,
	"NUMBER":  2,
	"JSON":    3,
}

func (x Feature_VariationType) String() string {
	return proto.EnumName(Feature_VariationType_name, int32(x))
}

func (Feature_VariationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_01855a1c92e13124, []int{0, 0}
}

type Feature struct {
	Id                    string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                  string                `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description           string                `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Enabled               bool                  `protobuf:"varint,4,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Deleted               bool                  `protobuf:"varint,5,opt,name=deleted,proto3" json:"deleted,omitempty"`
	EvaluationUndelayable bool                  `protobuf:"varint,6,opt,name=evaluation_undelayable,json=evaluationUndelayable,proto3" json:"evaluation_undelayable,omitempty"` // Deprecated: Do not use.
	Ttl                   int32                 `protobuf:"varint,7,opt,name=ttl,proto3" json:"ttl,omitempty"`
	Version               int32                 `protobuf:"varint,8,opt,name=version,proto3" json:"version,omitempty"`
	CreatedAt             int64                 `protobuf:"varint,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt             int64                 `protobuf:"varint,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Variations            []*Variation          `protobuf:"bytes,11,rep,name=variations,proto3" json:"variations,omitempty"`
	Targets               []*Target             `protobuf:"bytes,12,rep,name=targets,proto3" json:"targets,omitempty"`
	Rules                 []*Rule               `protobuf:"bytes,13,rep,name=rules,proto3" json:"rules,omitempty"`
	DefaultStrategy       *Strategy             `protobuf:"bytes,14,opt,name=default_strategy,json=defaultStrategy,proto3" json:"default_strategy,omitempty"`
	OffVariation          string                `protobuf:"bytes,15,opt,name=off_variation,json=offVariation,proto3" json:"off_variation,omitempty"`
	Tags                  []string              `protobuf:"bytes,16,rep,name=tags,proto3" json:"tags,omitempty"`
	LastUsedInfo          *FeatureLastUsedInfo  `protobuf:"bytes,17,opt,name=last_used_info,json=lastUsedInfo,proto3" json:"last_used_info,omitempty"`
	Maintainer            string                `protobuf:"bytes,18,opt,name=maintainer,proto3" json:"maintainer,omitempty"`
	VariationType         Feature_VariationType `protobuf:"varint,19,opt,name=variation_type,json=variationType,proto3,enum=bucketeer.feature.Feature_VariationType" json:"variation_type,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}              `json:"-"`
	XXX_unrecognized      []byte                `json:"-"`
	XXX_sizecache         int32                 `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_01855a1c92e13124, []int{0}
}

func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (m *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(m, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Feature) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *Feature) GetDeleted() bool {
	if m != nil {
		return m.Deleted
	}
	return false
}

// Deprecated: Do not use.
func (m *Feature) GetEvaluationUndelayable() bool {
	if m != nil {
		return m.EvaluationUndelayable
	}
	return false
}

func (m *Feature) GetTtl() int32 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Feature) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Feature) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Feature) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Feature) GetVariations() []*Variation {
	if m != nil {
		return m.Variations
	}
	return nil
}

func (m *Feature) GetTargets() []*Target {
	if m != nil {
		return m.Targets
	}
	return nil
}

func (m *Feature) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func (m *Feature) GetDefaultStrategy() *Strategy {
	if m != nil {
		return m.DefaultStrategy
	}
	return nil
}

func (m *Feature) GetOffVariation() string {
	if m != nil {
		return m.OffVariation
	}
	return ""
}

func (m *Feature) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *Feature) GetLastUsedInfo() *FeatureLastUsedInfo {
	if m != nil {
		return m.LastUsedInfo
	}
	return nil
}

func (m *Feature) GetMaintainer() string {
	if m != nil {
		return m.Maintainer
	}
	return ""
}

func (m *Feature) GetVariationType() Feature_VariationType {
	if m != nil {
		return m.VariationType
	}
	return Feature_STRING
}

type TagFeatures struct {
	Tag                  string     `protobuf:"bytes,1,opt,name=tag,proto3" json:"tag,omitempty"`
	Features             []*Feature `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *TagFeatures) Reset()         { *m = TagFeatures{} }
func (m *TagFeatures) String() string { return proto.CompactTextString(m) }
func (*TagFeatures) ProtoMessage()    {}
func (*TagFeatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_01855a1c92e13124, []int{1}
}

func (m *TagFeatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagFeatures.Unmarshal(m, b)
}
func (m *TagFeatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagFeatures.Marshal(b, m, deterministic)
}
func (m *TagFeatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFeatures.Merge(m, src)
}
func (m *TagFeatures) XXX_Size() int {
	return xxx_messageInfo_TagFeatures.Size(m)
}
func (m *TagFeatures) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFeatures.DiscardUnknown(m)
}

var xxx_messageInfo_TagFeatures proto.InternalMessageInfo

func (m *TagFeatures) GetTag() string {
	if m != nil {
		return m.Tag
	}
	return ""
}

func (m *TagFeatures) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

type Tag struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            int64    `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_01855a1c92e13124, []int{2}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Tag) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Tag) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterEnum("bucketeer.feature.Feature_VariationType", Feature_VariationType_name, Feature_VariationType_value)
	proto.RegisterType((*Feature)(nil), "bucketeer.feature.Feature")
	proto.RegisterType((*TagFeatures)(nil), "bucketeer.feature.TagFeatures")
	proto.RegisterType((*Tag)(nil), "bucketeer.feature.Tag")
}

func init() { proto.RegisterFile("proto/feature/feature.proto", fileDescriptor_01855a1c92e13124) }

var fileDescriptor_01855a1c92e13124 = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x71, 0x6b, 0xda, 0x5e,
	0x14, 0xfd, 0x25, 0x69, 0xd5, 0x5e, 0xab, 0x4d, 0xdf, 0x8f, 0x6d, 0x6f, 0xb6, 0x1d, 0xc1, 0xc1,
	0x08, 0x85, 0x2a, 0xb4, 0x30, 0x18, 0x8c, 0x41, 0x85, 0x76, 0x74, 0x74, 0x0a, 0x4f, 0xbb, 0xc1,
	0xfe, 0x09, 0x4f, 0x73, 0x93, 0x85, 0xa5, 0x89, 0x24, 0x2f, 0x82, 0x5f, 0x71, 0x9f, 0x6a, 0xbc,
	0xe7, 0x8b, 0x9a, 0xd5, 0xee, 0x2f, 0xef, 0x3d, 0xe7, 0x5c, 0xee, 0x09, 0xf7, 0x3c, 0xe1, 0x64,
	0x9e, 0xa5, 0x22, 0xed, 0x07, 0xc8, 0x45, 0x91, 0x61, 0xf9, 0xdb, 0x53, 0x28, 0x39, 0x9e, 0x16,
	0xb3, 0x5f, 0x28, 0x10, 0xb3, 0x9e, 0x26, 0x3a, 0xb4, 0xaa, 0xcf, 0x8a, 0x58, 0x8b, 0x3b, 0x9d,
	0x2a, 0x23, 0x78, 0x16, 0xa2, 0xd0, 0xdc, 0x59, 0x95, 0x5b, 0xf0, 0x2c, 0xe2, 0x22, 0x4a, 0x13,
	0x4d, 0x9f, 0x56, 0xe9, 0x5c, 0x64, 0x5c, 0x60, 0xb8, 0xd4, 0xec, 0xf9, 0x4e, 0x8b, 0x5e, 0xcc,
	0x73, 0xe1, 0x15, 0x39, 0xfa, 0x5e, 0x94, 0x04, 0xe9, 0x4a, 0xdb, 0xfd, 0x5d, 0x83, 0xfa, 0xed,
	0x4a, 0x40, 0xda, 0x60, 0x46, 0x3e, 0x35, 0x1c, 0xc3, 0x3d, 0x60, 0x66, 0xe4, 0x13, 0x02, 0x7b,
	0x09, 0x7f, 0x44, 0x6a, 0x2a, 0x44, 0xd5, 0xc4, 0x81, 0xa6, 0x8f, 0xf9, 0x2c, 0x8b, 0xe6, 0xd2,
	0x0e, 0xb5, 0x14, 0xb5, 0x0d, 0x11, 0x0a, 0x75, 0x4c, 0xf8, 0x34, 0x46, 0x9f, 0xee, 0x39, 0x86,
	0xdb, 0x60, 0x65, 0x2b, 0x19, 0x1f, 0x63, 0x14, 0xe8, 0xd3, 0xfd, 0x15, 0xa3, 0x5b, 0xf2, 0x01,
	0x5e, 0xe2, 0x82, 0xc7, 0x85, 0xfa, 0x46, 0xaf, 0x48, 0x7c, 0x8c, 0xf9, 0x52, 0x0e, 0xd1, 0x9a,
	0x14, 0x0e, 0x4c, 0x6a, 0xb0, 0x17, 0x1b, 0xc5, 0xc3, 0x46, 0x40, 0x6c, 0xb0, 0x84, 0x88, 0x69,
	0xdd, 0x31, 0xdc, 0x7d, 0x26, 0x4b, 0xb9, 0x66, 0x81, 0x59, 0x2e, 0xed, 0x35, 0x14, 0x5a, 0xb6,
	0xe4, 0x0c, 0x60, 0x96, 0x21, 0x17, 0xe8, 0x7b, 0x5c, 0xd0, 0x03, 0xc7, 0x70, 0x2d, 0x76, 0xa0,
	0x91, 0x6b, 0x21, 0xe9, 0x62, 0xee, 0x97, 0x34, 0xac, 0x68, 0x8d, 0x5c, 0x0b, 0xf2, 0x11, 0x60,
	0x7d, 0x87, 0x9c, 0x36, 0x1d, 0xcb, 0x6d, 0x5e, 0x9e, 0xf6, 0x9e, 0x5c, 0xbc, 0xf7, 0xad, 0x14,
	0xb1, 0x2d, 0x3d, 0xb9, 0x82, 0xfa, 0xea, 0xc2, 0x39, 0x3d, 0x54, 0xa3, 0xaf, 0x77, 0x8c, 0x4e,
	0x94, 0x82, 0x95, 0x4a, 0x72, 0x01, 0xfb, 0x32, 0x30, 0x39, 0x6d, 0xa9, 0x91, 0x57, 0x3b, 0x46,
	0x58, 0x11, 0x23, 0x5b, 0xa9, 0xc8, 0x2d, 0xd8, 0x3e, 0x06, 0xbc, 0x88, 0x85, 0x57, 0x46, 0x82,
	0xb6, 0x1d, 0xc3, 0x6d, 0x5e, 0x9e, 0xec, 0x98, 0x1c, 0x6b, 0x09, 0x3b, 0xd2, 0x43, 0x25, 0x40,
	0xde, 0x42, 0x2b, 0x0d, 0x02, 0x6f, 0xed, 0x9e, 0x1e, 0xa9, 0x33, 0x1f, 0xa6, 0x41, 0xb0, 0xfe,
	0x38, 0x99, 0x0e, 0xc1, 0xc3, 0x9c, 0xda, 0x8e, 0x25, 0xd3, 0x21, 0x6b, 0x72, 0x0f, 0xed, 0x6a,
	0xca, 0xe8, 0xb1, 0x5a, 0xff, 0x6e, 0xc7, 0x7a, 0x9d, 0xba, 0x7b, 0x9e, 0x8b, 0x87, 0x1c, 0xfd,
	0xbb, 0x24, 0x48, 0xd9, 0x61, 0xbc, 0xd5, 0x91, 0x37, 0x00, 0x8f, 0x3c, 0x4a, 0x04, 0x8f, 0x12,
	0xcc, 0x28, 0x51, 0x1e, 0xb6, 0x10, 0x32, 0x82, 0xf6, 0xda, 0xa2, 0x27, 0x96, 0x73, 0xa4, 0xff,
	0x3b, 0x86, 0xdb, 0xbe, 0x74, 0x9f, 0xdf, 0xb6, 0x39, 0xce, 0x64, 0x39, 0x47, 0xd6, 0x5a, 0x6c,
	0xb7, 0xdd, 0x4f, 0xd0, 0xaa, 0xf0, 0x04, 0xa0, 0x36, 0x9e, 0xb0, 0xbb, 0xe1, 0x67, 0xfb, 0x3f,
	0xd2, 0x84, 0xfa, 0x60, 0x34, 0xba, 0xbf, 0xb9, 0x1e, 0xda, 0x86, 0x24, 0x86, 0x0f, 0x5f, 0x07,
	0x37, 0xcc, 0x36, 0x49, 0x03, 0xf6, 0xbe, 0x8c, 0x47, 0x43, 0xdb, 0xea, 0x7e, 0x87, 0xe6, 0x84,
	0x87, 0x7a, 0x55, 0xae, 0xa2, 0xc9, 0x43, 0xfd, 0xa0, 0x64, 0x49, 0xde, 0x43, 0x43, 0x1b, 0xca,
	0xa9, 0xa9, 0x4e, 0xda, 0x79, 0xde, 0x2b, 0x5b, 0x6b, 0xbb, 0x63, 0xb0, 0x26, 0x3c, 0x7c, 0xf2,
	0x40, 0xab, 0x79, 0x36, 0xff, 0x9d, 0x67, 0xeb, 0xaf, 0x3c, 0x0f, 0xce, 0x7f, 0xb8, 0x61, 0x24,
	0x7e, 0x16, 0xd3, 0xde, 0x2c, 0x7d, 0xec, 0xcf, 0xf8, 0x85, 0x3f, 0xef, 0xaf, 0xcd, 0xf4, 0x2b,
	0xff, 0x21, 0xd3, 0x9a, 0x6a, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x48, 0x3a, 0x83, 0x55,
	0xfe, 0x04, 0x00, 0x00,
}
