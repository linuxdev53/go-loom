// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/chainconfig/chainconfig.proto

package chainconfig

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ChainConfigInitRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainConfigInitRequest) Reset()         { *m = ChainConfigInitRequest{} }
func (m *ChainConfigInitRequest) String() string { return proto.CompactTextString(m) }
func (*ChainConfigInitRequest) ProtoMessage()    {}
func (*ChainConfigInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{0}
}
func (m *ChainConfigInitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainConfigInitRequest.Unmarshal(m, b)
}
func (m *ChainConfigInitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainConfigInitRequest.Marshal(b, m, deterministic)
}
func (dst *ChainConfigInitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainConfigInitRequest.Merge(dst, src)
}
func (m *ChainConfigInitRequest) XXX_Size() int {
	return xxx_messageInfo_ChainConfigInitRequest.Size(m)
}
func (m *ChainConfigInitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainConfigInitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChainConfigInitRequest proto.InternalMessageInfo

type GetConfig struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfig) Reset()         { *m = GetConfig{} }
func (m *GetConfig) String() string { return proto.CompactTextString(m) }
func (*GetConfig) ProtoMessage()    {}
func (*GetConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{1}
}
func (m *GetConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfig.Unmarshal(m, b)
}
func (m *GetConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfig.Marshal(b, m, deterministic)
}
func (dst *GetConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfig.Merge(dst, src)
}
func (m *GetConfig) XXX_Size() int {
	return xxx_messageInfo_GetConfig.Size(m)
}
func (m *GetConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfig proto.InternalMessageInfo

func (m *GetConfig) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetFeature struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeature) Reset()         { *m = GetFeature{} }
func (m *GetFeature) String() string { return proto.CompactTextString(m) }
func (*GetFeature) ProtoMessage()    {}
func (*GetFeature) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{2}
}
func (m *GetFeature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeature.Unmarshal(m, b)
}
func (m *GetFeature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeature.Marshal(b, m, deterministic)
}
func (dst *GetFeature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeature.Merge(dst, src)
}
func (m *GetFeature) XXX_Size() int {
	return xxx_messageInfo_GetFeature.Size(m)
}
func (m *GetFeature) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeature.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeature proto.InternalMessageInfo

func (m *GetFeature) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type Config struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Config) Reset()         { *m = Config{} }
func (m *Config) String() string { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()    {}
func (*Config) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{3}
}
func (m *Config) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Config.Unmarshal(m, b)
}
func (m *Config) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Config.Marshal(b, m, deterministic)
}
func (dst *Config) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Config.Merge(dst, src)
}
func (m *Config) XXX_Size() int {
	return xxx_messageInfo_Config.Size(m)
}
func (m *Config) XXX_DiscardUnknown() {
	xxx_messageInfo_Config.DiscardUnknown(m)
}

var xxx_messageInfo_Config proto.InternalMessageInfo

func (m *Config) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Config) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Feature struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled              string   `protobuf:"bytes,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{4}
}
func (m *Feature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Feature.Unmarshal(m, b)
}
func (m *Feature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Feature.Marshal(b, m, deterministic)
}
func (dst *Feature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Feature.Merge(dst, src)
}
func (m *Feature) XXX_Size() int {
	return xxx_messageInfo_Feature.Size(m)
}
func (m *Feature) XXX_DiscardUnknown() {
	xxx_messageInfo_Feature.DiscardUnknown(m)
}

var xxx_messageInfo_Feature proto.InternalMessageInfo

func (m *Feature) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Feature) GetEnabled() string {
	if m != nil {
		return m.Enabled
	}
	return ""
}

type ListConfig struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListConfig) Reset()         { *m = ListConfig{} }
func (m *ListConfig) String() string { return proto.CompactTextString(m) }
func (*ListConfig) ProtoMessage()    {}
func (*ListConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{5}
}
func (m *ListConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfig.Unmarshal(m, b)
}
func (m *ListConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfig.Marshal(b, m, deterministic)
}
func (dst *ListConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfig.Merge(dst, src)
}
func (m *ListConfig) XXX_Size() int {
	return xxx_messageInfo_ListConfig.Size(m)
}
func (m *ListConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfig.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfig proto.InternalMessageInfo

type ListFeatures struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFeatures) Reset()         { *m = ListFeatures{} }
func (m *ListFeatures) String() string { return proto.CompactTextString(m) }
func (*ListFeatures) ProtoMessage()    {}
func (*ListFeatures) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{6}
}
func (m *ListFeatures) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFeatures.Unmarshal(m, b)
}
func (m *ListFeatures) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFeatures.Marshal(b, m, deterministic)
}
func (dst *ListFeatures) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFeatures.Merge(dst, src)
}
func (m *ListFeatures) XXX_Size() int {
	return xxx_messageInfo_ListFeatures.Size(m)
}
func (m *ListFeatures) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFeatures.DiscardUnknown(m)
}

var xxx_messageInfo_ListFeatures proto.InternalMessageInfo

type ListFeaturesResponse struct {
	Features             []*Feature `protobuf:"bytes,1,rep,name=features" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *ListFeaturesResponse) Reset()         { *m = ListFeaturesResponse{} }
func (m *ListFeaturesResponse) String() string { return proto.CompactTextString(m) }
func (*ListFeaturesResponse) ProtoMessage()    {}
func (*ListFeaturesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{7}
}
func (m *ListFeaturesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFeaturesResponse.Unmarshal(m, b)
}
func (m *ListFeaturesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFeaturesResponse.Marshal(b, m, deterministic)
}
func (dst *ListFeaturesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFeaturesResponse.Merge(dst, src)
}
func (m *ListFeaturesResponse) XXX_Size() int {
	return xxx_messageInfo_ListFeaturesResponse.Size(m)
}
func (m *ListFeaturesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFeaturesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListFeaturesResponse proto.InternalMessageInfo

func (m *ListFeaturesResponse) GetFeatures() []*Feature {
	if m != nil {
		return m.Features
	}
	return nil
}

type ListConfigResponse struct {
	Config               []*Config `protobuf:"bytes,1,rep,name=config" json:"config,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *ListConfigResponse) Reset()         { *m = ListConfigResponse{} }
func (m *ListConfigResponse) String() string { return proto.CompactTextString(m) }
func (*ListConfigResponse) ProtoMessage()    {}
func (*ListConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{8}
}
func (m *ListConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListConfigResponse.Unmarshal(m, b)
}
func (m *ListConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListConfigResponse.Marshal(b, m, deterministic)
}
func (dst *ListConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListConfigResponse.Merge(dst, src)
}
func (m *ListConfigResponse) XXX_Size() int {
	return xxx_messageInfo_ListConfigResponse.Size(m)
}
func (m *ListConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListConfigResponse proto.InternalMessageInfo

func (m *ListConfigResponse) GetConfig() []*Config {
	if m != nil {
		return m.Config
	}
	return nil
}

type UpdateFeature struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFeature) Reset()         { *m = UpdateFeature{} }
func (m *UpdateFeature) String() string { return proto.CompactTextString(m) }
func (*UpdateFeature) ProtoMessage()    {}
func (*UpdateFeature) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_48aa456c62a61fb8, []int{9}
}
func (m *UpdateFeature) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFeature.Unmarshal(m, b)
}
func (m *UpdateFeature) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFeature.Marshal(b, m, deterministic)
}
func (dst *UpdateFeature) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeature.Merge(dst, src)
}
func (m *UpdateFeature) XXX_Size() int {
	return xxx_messageInfo_UpdateFeature.Size(m)
}
func (m *UpdateFeature) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeature.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeature proto.InternalMessageInfo

func (m *UpdateFeature) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpdateFeature) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*ChainConfigInitRequest)(nil), "ChainConfigInitRequest")
	proto.RegisterType((*GetConfig)(nil), "GetConfig")
	proto.RegisterType((*GetFeature)(nil), "GetFeature")
	proto.RegisterType((*Config)(nil), "Config")
	proto.RegisterType((*Feature)(nil), "Feature")
	proto.RegisterType((*ListConfig)(nil), "ListConfig")
	proto.RegisterType((*ListFeatures)(nil), "ListFeatures")
	proto.RegisterType((*ListFeaturesResponse)(nil), "ListFeaturesResponse")
	proto.RegisterType((*ListConfigResponse)(nil), "ListConfigResponse")
	proto.RegisterType((*UpdateFeature)(nil), "UpdateFeature")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/chainconfig/chainconfig.proto", fileDescriptor_chainconfig_48aa456c62a61fb8)
}

var fileDescriptor_chainconfig_48aa456c62a61fb8 = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x89, 0xd5, 0x24, 0x1d, 0xab, 0x48, 0x28, 0x12, 0x04, 0xb5, 0x2c, 0x1e, 0x7a, 0x31,
	0x91, 0x8a, 0x78, 0xd0, 0x5b, 0xc1, 0xa2, 0x78, 0x0a, 0xf8, 0x03, 0x36, 0xe9, 0x34, 0x5d, 0x9a,
	0xec, 0xc6, 0xec, 0x44, 0xe9, 0xbf, 0x97, 0x24, 0xdb, 0x58, 0xa4, 0xd2, 0x4b, 0x98, 0xf7, 0x66,
	0xde, 0x47, 0x66, 0x16, 0xde, 0x52, 0x41, 0xcb, 0x2a, 0x0e, 0x12, 0x95, 0x87, 0x99, 0x52, 0xb9,
	0x44, 0xfa, 0x56, 0xe5, 0x2a, 0x4c, 0xd5, 0x6d, 0x2d, 0xc3, 0xb8, 0x12, 0x19, 0x09, 0x19, 0xd2,
	0xba, 0x40, 0x1d, 0x26, 0x4b, 0x2e, 0x64, 0xa2, 0xe4, 0x42, 0xa4, 0xdb, 0x75, 0x50, 0x94, 0x8a,
	0xd4, 0xc5, 0xdd, 0x1e, 0x56, 0xcb, 0x68, 0xbe, 0x6d, 0x82, 0xf9, 0x70, 0x3e, 0xad, 0x31, 0xd3,
	0x06, 0xf3, 0x2a, 0x05, 0x45, 0xf8, 0x59, 0xa1, 0x26, 0x76, 0x09, 0xfd, 0x19, 0x52, 0xeb, 0x7b,
	0x67, 0xd0, 0x5b, 0xe1, 0xda, 0xb7, 0x46, 0xd6, 0xb8, 0x1f, 0xd5, 0x25, 0xbb, 0x02, 0x98, 0x21,
	0xbd, 0x20, 0xa7, 0xaa, 0xc4, 0x1d, 0xfd, 0x09, 0xd8, 0x26, 0xeb, 0xc1, 0xa1, 0xe4, 0x39, 0x9a,
	0x66, 0x53, 0x7b, 0x43, 0x38, 0xfa, 0xe2, 0x59, 0x85, 0xfe, 0x41, 0x63, 0xb6, 0x82, 0x3d, 0x82,
	0xb3, 0x01, 0xee, 0x0a, 0xf9, 0xe0, 0xa0, 0xe4, 0x71, 0x86, 0x73, 0x13, 0xdb, 0x48, 0x36, 0x00,
	0x78, 0x17, 0xda, 0xfc, 0x2c, 0x3b, 0x85, 0x41, 0xad, 0x0c, 0x4a, 0xb3, 0x67, 0x18, 0x6e, 0xeb,
	0x08, 0x75, 0xa1, 0xa4, 0x46, 0xef, 0x06, 0xdc, 0x85, 0xf1, 0x7c, 0x6b, 0xd4, 0x1b, 0x1f, 0x4f,
	0xdc, 0xc0, 0x0c, 0x45, 0x5d, 0x87, 0x3d, 0x80, 0xf7, 0xcb, 0xee, 0xb2, 0xd7, 0x60, 0xb7, 0x97,
	0x37, 0x49, 0x27, 0x30, 0x03, 0xc6, 0x66, 0x4f, 0x70, 0xf2, 0x51, 0xcc, 0x39, 0xe1, 0xbf, 0x27,
	0xfa, 0xbb, 0x8f, 0xdb, 0xed, 0x13, 0xdb, 0xcd, 0xe3, 0xdc, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff,
	0x73, 0x5f, 0xc4, 0xce, 0x1c, 0x02, 0x00, 0x00,
}
