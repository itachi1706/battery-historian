// Code generated by protoc-gen-go.
// source: github.com/google/battery-historian/pb/build_proto/build.proto
// DO NOT EDIT!

/*
Package build is a generated protocol buffer package.

It is generated from these files:

	github.com/google/battery-historian/pb/build_proto/build.proto

It has these top-level messages:

	Build
*/
package build

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type Build struct {
	// Fingerprint, e.g. "google/mysid/toro:4.0.4/IMM06/243892:userdebug/dev-keys"
	Fingerprint *string `protobuf:"bytes,1,opt,name=fingerprint" json:"fingerprint,omitempty"`
	// Carrier, e.g. "google"
	Brand *string `protobuf:"bytes,2,opt,name=brand" json:"brand,omitempty"`
	// Product name, e.g. "mysid"
	Product *string `protobuf:"bytes,3,opt,name=product" json:"product,omitempty"`
	// Product name, e.g. "toro"
	Device *string `protobuf:"bytes,4,opt,name=device" json:"device,omitempty"`
	// Release version, e.g. "4.0.4"
	Release *string `protobuf:"bytes,5,opt,name=release" json:"release,omitempty"`
	// Build id, e.g. "IMM06"
	BuildId *string `protobuf:"bytes,6,opt,name=build_id" json:"build_id,omitempty"`
	// Incremental build id, e.g. "243892"
	Incremental *string `protobuf:"bytes,7,opt,name=incremental" json:"incremental,omitempty"`
	// Type of build, e.g. "userdebug"
	Type *string `protobuf:"bytes,8,opt,name=type" json:"type,omitempty"`
	// Tags, e.g. "dev-keys"
	Tags             []string `protobuf:"bytes,9,rep,name=tags" json:"tags,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Build) Reset()                    { *m = Build{} }
func (m *Build) String() string            { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()               {}
func (*Build) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Build) GetFingerprint() string {
	if m != nil && m.Fingerprint != nil {
		return *m.Fingerprint
	}
	return ""
}

func (m *Build) GetBrand() string {
	if m != nil && m.Brand != nil {
		return *m.Brand
	}
	return ""
}

func (m *Build) GetProduct() string {
	if m != nil && m.Product != nil {
		return *m.Product
	}
	return ""
}

func (m *Build) GetDevice() string {
	if m != nil && m.Device != nil {
		return *m.Device
	}
	return ""
}

func (m *Build) GetRelease() string {
	if m != nil && m.Release != nil {
		return *m.Release
	}
	return ""
}

func (m *Build) GetBuildId() string {
	if m != nil && m.BuildId != nil {
		return *m.BuildId
	}
	return ""
}

func (m *Build) GetIncremental() string {
	if m != nil && m.Incremental != nil {
		return *m.Incremental
	}
	return ""
}

func (m *Build) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return ""
}

func (m *Build) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*Build)(nil), "build.Build")
}

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x3c, 0x8e, 0x4b, 0x4e, 0x03, 0x31,
	0x0c, 0x40, 0x35, 0xb4, 0xd3, 0x76, 0xcc, 0x57, 0xc3, 0xc6, 0xcb, 0x8a, 0x55, 0x37, 0x34, 0x37,
	0x60, 0xc1, 0x45, 0x50, 0x3e, 0x26, 0xb5, 0x94, 0x26, 0x91, 0xc7, 0x83, 0xd4, 0xdb, 0x70, 0x54,
	0x4a, 0x46, 0xb0, 0x7b, 0xef, 0xc9, 0x96, 0x0d, 0x6f, 0x91, 0xf5, 0x34, 0xbb, 0xa3, 0x2f, 0x67,
	0x13, 0x4b, 0x89, 0x89, 0x8c, 0xb3, 0xaa, 0x24, 0x97, 0xd7, 0x13, 0x4f, 0x5a, 0x84, 0x6d, 0x36,
	0xd5, 0x19, 0x37, 0x73, 0x0a, 0x1f, 0x55, 0x8a, 0x96, 0x85, 0x8f, 0x8d, 0xc7, 0xbe, 0xc9, 0xcb,
	0x77, 0x07, 0xfd, 0xfb, 0x2f, 0x8d, 0xcf, 0x70, 0xfb, 0xc9, 0x39, 0x92, 0x54, 0xe1, 0xac, 0xd8,
	0xed, 0xbb, 0xc3, 0x30, 0xde, 0x43, 0xef, 0xc4, 0xe6, 0x80, 0x37, 0x4d, 0x1f, 0x61, 0x7b, 0xdd,
	0x0e, 0xb3, 0x57, 0x5c, 0xb5, 0xf0, 0x00, 0x9b, 0x40, 0x5f, 0xec, 0x09, 0xd7, 0x7f, 0x03, 0x42,
	0x89, 0xec, 0x44, 0xd8, 0xb7, 0xf0, 0x04, 0xbb, 0xe5, 0x03, 0x0e, 0xb8, 0x69, 0xe5, 0x7a, 0x87,
	0xb3, 0x17, 0x3a, 0x53, 0x56, 0x9b, 0x70, 0xdb, 0xe2, 0x1d, 0xac, 0xf5, 0x52, 0x09, 0x77, 0xff,
	0x66, 0xe3, 0x84, 0xc3, 0x7e, 0x75, 0x18, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0x38, 0xcd, 0xd1,
	0xa3, 0xea, 0x00, 0x00, 0x00,
}
