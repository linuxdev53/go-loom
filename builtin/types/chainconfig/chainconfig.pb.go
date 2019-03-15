// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/chainconfig/chainconfig.proto

package chainconfig

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

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
	// Only the owner will be allowed to add known features
	Owner                *types.Address `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ChainConfigInitRequest) Reset()         { *m = ChainConfigInitRequest{} }
func (m *ChainConfigInitRequest) String() string { return proto.CompactTextString(m) }
func (*ChainConfigInitRequest) ProtoMessage()    {}
func (*ChainConfigInitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{0}
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

func (m *ChainConfigInitRequest) GetOwner() *types.Address {
	if m != nil {
		return m.Owner
	}
	return nil
}

type GetConfigRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConfigRequest) Reset()         { *m = GetConfigRequest{} }
func (m *GetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*GetConfigRequest) ProtoMessage()    {}
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{1}
}
func (m *GetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConfigRequest.Unmarshal(m, b)
}
func (m *GetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConfigRequest.Marshal(b, m, deterministic)
}
func (dst *GetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConfigRequest.Merge(dst, src)
}
func (m *GetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_GetConfigRequest.Size(m)
}
func (m *GetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConfigRequest proto.InternalMessageInfo

func (m *GetConfigRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetFeatureRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeatureRequest) Reset()         { *m = GetFeatureRequest{} }
func (m *GetFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*GetFeatureRequest) ProtoMessage()    {}
func (*GetFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{2}
}
func (m *GetFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeatureRequest.Unmarshal(m, b)
}
func (m *GetFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *GetFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeatureRequest.Merge(dst, src)
}
func (m *GetFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_GetFeatureRequest.Size(m)
}
func (m *GetFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeatureRequest proto.InternalMessageInfo

func (m *GetFeatureRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SetFeatureRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetFeatureRequest) Reset()         { *m = SetFeatureRequest{} }
func (m *SetFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*SetFeatureRequest) ProtoMessage()    {}
func (*SetFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{3}
}
func (m *SetFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFeatureRequest.Unmarshal(m, b)
}
func (m *SetFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *SetFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFeatureRequest.Merge(dst, src)
}
func (m *SetFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_SetFeatureRequest.Size(m)
}
func (m *SetFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetFeatureRequest proto.InternalMessageInfo

func (m *SetFeatureRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type SetFeatureResponse struct {
	Feature              *GetFeatureResponse `protobuf:"bytes,1,opt,name=feature" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SetFeatureResponse) Reset()         { *m = SetFeatureResponse{} }
func (m *SetFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*SetFeatureResponse) ProtoMessage()    {}
func (*SetFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{4}
}
func (m *SetFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetFeatureResponse.Unmarshal(m, b)
}
func (m *SetFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetFeatureResponse.Marshal(b, m, deterministic)
}
func (dst *SetFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetFeatureResponse.Merge(dst, src)
}
func (m *SetFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_SetFeatureResponse.Size(m)
}
func (m *SetFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetFeatureResponse proto.InternalMessageInfo

func (m *SetFeatureResponse) GetFeature() *GetFeatureResponse {
	if m != nil {
		return m.Feature
	}
	return nil
}

type GetFeatureResponse struct {
	Feature              *Feature `protobuf:"bytes,2,opt,name=feature" json:"feature,omitempty"`
	Percentage           uint64   `protobuf:"varint,3,opt,name=percentage,proto3" json:"percentage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetFeatureResponse) Reset()         { *m = GetFeatureResponse{} }
func (m *GetFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*GetFeatureResponse) ProtoMessage()    {}
func (*GetFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{5}
}
func (m *GetFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetFeatureResponse.Unmarshal(m, b)
}
func (m *GetFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetFeatureResponse.Marshal(b, m, deterministic)
}
func (dst *GetFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetFeatureResponse.Merge(dst, src)
}
func (m *GetFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_GetFeatureResponse.Size(m)
}
func (m *GetFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetFeatureResponse proto.InternalMessageInfo

func (m *GetFeatureResponse) GetFeature() *Feature {
	if m != nil {
		return m.Feature
	}
	return nil
}

func (m *GetFeatureResponse) GetPercentage() uint64 {
	if m != nil {
		return m.Percentage
	}
	return 0
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
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{6}
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
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Enabled              string           `protobuf:"bytes,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	Status               string           `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Validators           []*types.Address `protobuf:"bytes,4,rep,name=validators" json:"validators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}
func (*Feature) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{7}
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

func (m *Feature) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *Feature) GetValidators() []*types.Address {
	if m != nil {
		return m.Validators
	}
	return nil
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
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{8}
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

type ListFeaturesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListFeaturesRequest) Reset()         { *m = ListFeaturesRequest{} }
func (m *ListFeaturesRequest) String() string { return proto.CompactTextString(m) }
func (*ListFeaturesRequest) ProtoMessage()    {}
func (*ListFeaturesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{9}
}
func (m *ListFeaturesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListFeaturesRequest.Unmarshal(m, b)
}
func (m *ListFeaturesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListFeaturesRequest.Marshal(b, m, deterministic)
}
func (dst *ListFeaturesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListFeaturesRequest.Merge(dst, src)
}
func (m *ListFeaturesRequest) XXX_Size() int {
	return xxx_messageInfo_ListFeaturesRequest.Size(m)
}
func (m *ListFeaturesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListFeaturesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListFeaturesRequest proto.InternalMessageInfo

type ListFeaturesResponse struct {
	Features             []*GetFeatureResponse `protobuf:"bytes,1,rep,name=features" json:"features,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ListFeaturesResponse) Reset()         { *m = ListFeaturesResponse{} }
func (m *ListFeaturesResponse) String() string { return proto.CompactTextString(m) }
func (*ListFeaturesResponse) ProtoMessage()    {}
func (*ListFeaturesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{10}
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

func (m *ListFeaturesResponse) GetFeatures() []*GetFeatureResponse {
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
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{11}
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

type EnableFeatureRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnableFeatureRequest) Reset()         { *m = EnableFeatureRequest{} }
func (m *EnableFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*EnableFeatureRequest) ProtoMessage()    {}
func (*EnableFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{12}
}
func (m *EnableFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnableFeatureRequest.Unmarshal(m, b)
}
func (m *EnableFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnableFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *EnableFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnableFeatureRequest.Merge(dst, src)
}
func (m *EnableFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_EnableFeatureRequest.Size(m)
}
func (m *EnableFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnableFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnableFeatureRequest proto.InternalMessageInfo

func (m *EnableFeatureRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type EnableFeatureResponse struct {
	Feature              *GetFeatureResponse `protobuf:"bytes,1,opt,name=feature" json:"feature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *EnableFeatureResponse) Reset()         { *m = EnableFeatureResponse{} }
func (m *EnableFeatureResponse) String() string { return proto.CompactTextString(m) }
func (*EnableFeatureResponse) ProtoMessage()    {}
func (*EnableFeatureResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{13}
}
func (m *EnableFeatureResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnableFeatureResponse.Unmarshal(m, b)
}
func (m *EnableFeatureResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnableFeatureResponse.Marshal(b, m, deterministic)
}
func (dst *EnableFeatureResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnableFeatureResponse.Merge(dst, src)
}
func (m *EnableFeatureResponse) XXX_Size() int {
	return xxx_messageInfo_EnableFeatureResponse.Size(m)
}
func (m *EnableFeatureResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EnableFeatureResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EnableFeatureResponse proto.InternalMessageInfo

func (m *EnableFeatureResponse) GetFeature() *GetFeatureResponse {
	if m != nil {
		return m.Feature
	}
	return nil
}

type UpdateFeatureRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Enabled              bool     `protobuf:"varint,2,opt,name=enabled,proto3" json:"enabled,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateFeatureRequest) Reset()         { *m = UpdateFeatureRequest{} }
func (m *UpdateFeatureRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateFeatureRequest) ProtoMessage()    {}
func (*UpdateFeatureRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_chainconfig_e0b5057d16b78105, []int{14}
}
func (m *UpdateFeatureRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateFeatureRequest.Unmarshal(m, b)
}
func (m *UpdateFeatureRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateFeatureRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateFeatureRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateFeatureRequest.Merge(dst, src)
}
func (m *UpdateFeatureRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateFeatureRequest.Size(m)
}
func (m *UpdateFeatureRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateFeatureRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateFeatureRequest proto.InternalMessageInfo

func (m *UpdateFeatureRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *UpdateFeatureRequest) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func init() {
	proto.RegisterType((*ChainConfigInitRequest)(nil), "ChainConfigInitRequest")
	proto.RegisterType((*GetConfigRequest)(nil), "GetConfigRequest")
	proto.RegisterType((*GetFeatureRequest)(nil), "GetFeatureRequest")
	proto.RegisterType((*SetFeatureRequest)(nil), "SetFeatureRequest")
	proto.RegisterType((*SetFeatureResponse)(nil), "SetFeatureResponse")
	proto.RegisterType((*GetFeatureResponse)(nil), "GetFeatureResponse")
	proto.RegisterType((*Config)(nil), "Config")
	proto.RegisterType((*Feature)(nil), "Feature")
	proto.RegisterType((*ListConfig)(nil), "ListConfig")
	proto.RegisterType((*ListFeaturesRequest)(nil), "ListFeaturesRequest")
	proto.RegisterType((*ListFeaturesResponse)(nil), "ListFeaturesResponse")
	proto.RegisterType((*ListConfigResponse)(nil), "ListConfigResponse")
	proto.RegisterType((*EnableFeatureRequest)(nil), "EnableFeatureRequest")
	proto.RegisterType((*EnableFeatureResponse)(nil), "EnableFeatureResponse")
	proto.RegisterType((*UpdateFeatureRequest)(nil), "UpdateFeatureRequest")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/chainconfig/chainconfig.proto", fileDescriptor_chainconfig_e0b5057d16b78105)
}

var fileDescriptor_chainconfig_e0b5057d16b78105 = []byte{
	// 425 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdf, 0x6b, 0xd4, 0x40,
	0x10, 0x26, 0xde, 0x35, 0xb9, 0x1b, 0x7d, 0xa8, 0xdb, 0xb4, 0x04, 0x1f, 0xea, 0xb1, 0x0f, 0x7a,
	0x08, 0x4d, 0xa4, 0x22, 0xf8, 0xaa, 0x87, 0x0d, 0x8a, 0x4f, 0x11, 0xc1, 0xd7, 0x4d, 0x32, 0x4d,
	0x97, 0xe6, 0x76, 0x63, 0x76, 0xd3, 0x72, 0xff, 0xbd, 0x64, 0x77, 0x73, 0x4d, 0xaf, 0x07, 0x2d,
	0x7d, 0x09, 0xf3, 0xe3, 0xfb, 0xbe, 0x99, 0x9d, 0x99, 0xc0, 0xcf, 0x8a, 0xeb, 0xab, 0x2e, 0x8f,
	0x0b, 0xb9, 0x4e, 0x6a, 0x29, 0xd7, 0x02, 0xf5, 0xad, 0x6c, 0xaf, 0x93, 0x4a, 0x9e, 0xf5, 0x6e,
	0x92, 0x77, 0xbc, 0xd6, 0x5c, 0x24, 0x7a, 0xd3, 0xa0, 0x4a, 0x8a, 0x2b, 0xc6, 0x45, 0x21, 0xc5,
	0x25, 0xaf, 0xc6, 0x76, 0xdc, 0xb4, 0x52, 0xcb, 0x37, 0x1f, 0x1f, 0xd1, 0xb2, 0x1a, 0xe6, 0x6b,
	0x19, 0xf4, 0x0b, 0x9c, 0xac, 0x7a, 0x99, 0x95, 0x91, 0xf9, 0x21, 0xb8, 0xce, 0xf0, 0x5f, 0x87,
	0x4a, 0x93, 0x53, 0x38, 0x90, 0xb7, 0x02, 0xdb, 0xc8, 0x5b, 0x78, 0xcb, 0x97, 0xe7, 0xb3, 0xf8,
	0x6b, 0x59, 0xb6, 0xa8, 0x54, 0x66, 0xc3, 0xf4, 0x1d, 0x1c, 0xa6, 0xa8, 0x2d, 0x6f, 0xe0, 0x10,
	0x98, 0x0a, 0xb6, 0x46, 0x43, 0x99, 0x67, 0xc6, 0xa6, 0xef, 0xe1, 0x75, 0x8a, 0xfa, 0x02, 0x99,
	0xee, 0x5a, 0x7c, 0x04, 0xf8, 0xfb, 0x49, 0xc0, 0x15, 0x90, 0x31, 0x50, 0x35, 0x52, 0x28, 0x24,
	0x67, 0x10, 0x5c, 0xda, 0x90, 0xeb, 0xf8, 0x28, 0x4e, 0x1f, 0xa0, 0xb2, 0x01, 0x43, 0xff, 0x02,
	0x79, 0x98, 0x26, 0xf4, 0x4e, 0xe4, 0x85, 0x7b, 0xf6, 0x00, 0x19, 0x12, 0xe4, 0x14, 0xa0, 0xc1,
	0xb6, 0x40, 0xa1, 0x59, 0x85, 0xd1, 0x64, 0xe1, 0x2d, 0xa7, 0xd9, 0x28, 0x42, 0xcf, 0xc1, 0xb7,
	0x53, 0xd9, 0xd7, 0x3c, 0x09, 0xe1, 0xe0, 0x86, 0xd5, 0x9d, 0xd5, 0x9f, 0x67, 0xd6, 0xa1, 0x1b,
	0x08, 0x5c, 0x9d, 0xbd, 0xa4, 0x08, 0x02, 0x14, 0x2c, 0xaf, 0xb1, 0x74, 0xb4, 0xc1, 0x25, 0x27,
	0xe0, 0x2b, 0xcd, 0x74, 0xa7, 0x4c, 0x23, 0xf3, 0xcc, 0x79, 0x64, 0x09, 0x70, 0xc3, 0x6a, 0x5e,
	0x32, 0x2d, 0x5b, 0x15, 0x4d, 0x17, 0x93, 0x7b, 0x2b, 0x1c, 0xe5, 0xe8, 0x2b, 0x80, 0x5f, 0x5c,
	0xb9, 0x45, 0xd2, 0x63, 0x38, 0xea, 0x3d, 0xd7, 0x8c, 0x72, 0x6b, 0xa0, 0x29, 0x84, 0xf7, 0xc3,
	0x6e, 0x5e, 0x09, 0xcc, 0xdc, 0x58, 0x54, 0xe4, 0x99, 0x22, 0x7b, 0xa7, 0xbe, 0x05, 0xd1, 0xcf,
	0x40, 0xee, 0xaa, 0x6d, 0x65, 0xde, 0x82, 0x6f, 0xef, 0xd8, 0x89, 0x04, 0xb1, 0x03, 0xb8, 0x30,
	0xfd, 0x00, 0xe1, 0x77, 0xf3, 0xe2, 0x27, 0x9c, 0xc7, 0x05, 0x1c, 0xef, 0x60, 0x9f, 0x77, 0x21,
	0xdf, 0x20, 0xfc, 0xd3, 0x94, 0x4c, 0xef, 0xd6, 0x3c, 0x84, 0xc9, 0x35, 0x6e, 0x5c, 0xc9, 0xde,
	0xdc, 0x5d, 0xcf, 0x6c, 0xbb, 0x9e, 0xdc, 0x37, 0x7f, 0xd9, 0xa7, 0xff, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x1f, 0x97, 0x70, 0x1a, 0xe5, 0x03, 0x00, 0x00,
}
