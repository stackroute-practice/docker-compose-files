// Code generated by protoc-gen-go.
// source: adapter.proto
// DO NOT EDIT!

/*
Package voltha is a generated protocol buffer package.

It is generated from these files:
	adapter.proto

It has these top-level messages:
	AdapterConfig
	Adapter
	Adapters
*/
package adapter

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"
import voltha2 "github.com/opencord/voltha/netconf/translator/voltha/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AdapterConfig struct {
	// Common adapter config attributes here
	LogLevel voltha2.LogLevel_LogLevel `protobuf:"varint,1,opt,name=log_level,json=logLevel,enum=voltha.LogLevel_LogLevel" json:"log_level,omitempty"`
	// Custom (vendor-specific) configuration attributes
	AdditionalConfig *google_protobuf.Any `protobuf:"bytes,64,opt,name=additional_config,json=additionalConfig" json:"additional_config,omitempty"`
}

func (m *AdapterConfig) Reset()                    { *m = AdapterConfig{} }
func (m *AdapterConfig) String() string            { return proto.CompactTextString(m) }
func (*AdapterConfig) ProtoMessage()               {}
func (*AdapterConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AdapterConfig) GetLogLevel() voltha2.LogLevel_LogLevel {
	if m != nil {
		return m.LogLevel
	}
	return voltha2.LogLevel_DEBUG
}

func (m *AdapterConfig) GetAdditionalConfig() *google_protobuf.Any {
	if m != nil {
		return m.AdditionalConfig
	}
	return nil
}

// Adapter (software plugin)
type Adapter struct {
	// Unique name of adapter, matching the python packate name under
	// voltha/adapters.
	Id      string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Vendor  string `protobuf:"bytes,2,opt,name=vendor" json:"vendor,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	// Adapter configuration
	Config *AdapterConfig `protobuf:"bytes,16,opt,name=config" json:"config,omitempty"`
	// Custom descriptors and custom configuration
	AdditionalDescription *google_protobuf.Any `protobuf:"bytes,64,opt,name=additional_description,json=additionalDescription" json:"additional_description,omitempty"`
	LogicalDeviceIds      []string             `protobuf:"bytes,4,rep,name=logical_device_ids,json=logicalDeviceIds" json:"logical_device_ids,omitempty"`
}

func (m *Adapter) Reset()                    { *m = Adapter{} }
func (m *Adapter) String() string            { return proto.CompactTextString(m) }
func (*Adapter) ProtoMessage()               {}
func (*Adapter) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Adapter) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Adapter) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *Adapter) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Adapter) GetConfig() *AdapterConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *Adapter) GetAdditionalDescription() *google_protobuf.Any {
	if m != nil {
		return m.AdditionalDescription
	}
	return nil
}

func (m *Adapter) GetLogicalDeviceIds() []string {
	if m != nil {
		return m.LogicalDeviceIds
	}
	return nil
}

type Adapters struct {
	Items []*Adapter `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Adapters) Reset()                    { *m = Adapters{} }
func (m *Adapters) String() string            { return proto.CompactTextString(m) }
func (*Adapters) ProtoMessage()               {}
func (*Adapters) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Adapters) GetItems() []*Adapter {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*AdapterConfig)(nil), "voltha.AdapterConfig")
	proto.RegisterType((*Adapter)(nil), "voltha.Adapter")
	proto.RegisterType((*Adapters)(nil), "voltha.Adapters")
}

func init() { proto.RegisterFile("adapter.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xd3, 0x72, 0x29, 0x70, 0xb8, 0xdc, 0x8b, 0x13, 0x31, 0x85, 0x84, 0xd8, 0x90, 0x98,
	0x74, 0xa1, 0x25, 0x62, 0xe2, 0x5a, 0x94, 0x8d, 0x09, 0xab, 0xbe, 0x00, 0x19, 0x3a, 0x43, 0x9d,
	0x64, 0x3a, 0x87, 0xb4, 0xb5, 0x09, 0xaf, 0xe0, 0xce, 0x07, 0xf3, 0x3d, 0x7c, 0x02, 0xd7, 0x86,
	0x99, 0xa9, 0xa0, 0x0b, 0x77, 0x67, 0xfe, 0xff, 0x9c, 0xf3, 0x7f, 0x67, 0xa0, 0x47, 0x19, 0xdd,
	0x96, 0x3c, 0x8f, 0xb6, 0x39, 0x96, 0x48, 0xbc, 0x0a, 0x65, 0xf9, 0x44, 0x47, 0xc3, 0x14, 0x31,
	0x95, 0x7c, 0xaa, 0xd5, 0xf5, 0xf3, 0x66, 0x4a, 0xd5, 0xce, 0xb4, 0x8c, 0xfe, 0x26, 0x98, 0x65,
	0xa8, 0xec, 0x0b, 0x32, 0x5e, 0x52, 0x53, 0x4f, 0x5e, 0x1c, 0xe8, 0xcd, 0xcd, 0xba, 0x07, 0x54,
	0x1b, 0x91, 0x92, 0x5b, 0xe8, 0x48, 0x4c, 0x57, 0x92, 0x57, 0x5c, 0xfa, 0x4e, 0xe0, 0x84, 0xff,
	0x66, 0xc3, 0xc8, 0x44, 0x44, 0x4b, 0x4c, 0x97, 0x7b, 0xfd, 0xab, 0x88, 0xdb, 0xd2, 0x56, 0x64,
	0x0e, 0x27, 0x94, 0x31, 0x51, 0x0a, 0x54, 0x54, 0xae, 0x12, 0xbd, 0xcc, 0xbf, 0x0b, 0x9c, 0xb0,
	0x3b, 0x3b, 0x8d, 0x0c, 0x5a, 0x54, 0xa3, 0x45, 0x73, 0xb5, 0x8b, 0xfb, 0x87, 0x76, 0x13, 0x3d,
	0x79, 0x75, 0xa1, 0x65, 0x61, 0xc8, 0x00, 0x5c, 0xc1, 0x74, 0x7e, 0xe7, 0xbe, 0xf9, 0xfe, 0xf1,
	0x36, 0x76, 0x62, 0x57, 0x30, 0x32, 0x06, 0xaf, 0xe2, 0x8a, 0x61, 0xee, 0xbb, 0xc7, 0x96, 0x15,
	0xc9, 0x39, 0xb4, 0x2a, 0x9e, 0x17, 0x02, 0x95, 0xdf, 0x38, 0xf6, 0x6b, 0x95, 0x5c, 0x81, 0x67,
	0xd1, 0xfa, 0x1a, 0x6d, 0x50, 0x9f, 0xf6, 0xed, 0x13, 0x62, 0xdb, 0x44, 0x62, 0x38, 0x3b, 0x3a,
	0x8a, 0xf1, 0x22, 0xc9, 0xc5, 0x76, 0xff, 0xfa, 0xed, 0xb2, 0x3a, 0x74, 0x70, 0x18, 0x5d, 0x1c,
	0x26, 0xc9, 0x25, 0x10, 0x89, 0xa9, 0x48, 0xf4, 0xc2, 0x4a, 0x24, 0x7c, 0x25, 0x58, 0xe1, 0xff,
	0x09, 0x1a, 0x61, 0x27, 0xee, 0x5b, 0x67, 0xa1, 0x8d, 0x47, 0x56, 0x4c, 0xae, 0xa1, 0x6d, 0xd1,
	0x0a, 0x72, 0x01, 0x4d, 0x51, 0xf2, 0xac, 0xf0, 0x9d, 0xa0, 0x11, 0x76, 0x67, 0xff, 0x7f, 0xb0,
	0xc7, 0xc6, 0x5d, 0x7b, 0x1a, 0xe6, 0xe6, 0x33, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x5e, 0xc4, 0x9a,
	0x28, 0x02, 0x00, 0x00,
}