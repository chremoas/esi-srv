// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common.proto

/*
Package chremoas_esi is a generated protocol buffer package.

It is generated from these files:
	common.proto

It has these top-level messages:
*/
package chremoas_esi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EntityType int32

const (
	EntityType_ALLIANCE    EntityType = 0
	EntityType_CORPORATION EntityType = 1
	EntityType_CHARACTER   EntityType = 2
)

var EntityType_name = map[int32]string{
	0: "ALLIANCE",
	1: "CORPORATION",
	2: "CHARACTER",
}
var EntityType_value = map[string]int32{
	"ALLIANCE":    0,
	"CORPORATION": 1,
	"CHARACTER":   2,
}

func (x EntityType) String() string {
	return proto.EnumName(EntityType_name, int32(x))
}
func (EntityType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterEnum("chremoas.esi.EntityType", EntityType_name, EntityType_value)
}

func init() { proto.RegisterFile("common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xce, 0xcf, 0xcd,
	0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x49, 0xce, 0x28, 0x4a, 0xcd, 0xcd,
	0x4f, 0x2c, 0xd6, 0x4b, 0x2d, 0xce, 0xd4, 0xb2, 0xe2, 0xe2, 0x72, 0xcd, 0x2b, 0xc9, 0x2c, 0xa9,
	0x0c, 0xa9, 0x2c, 0x48, 0x15, 0xe2, 0xe1, 0xe2, 0x70, 0xf4, 0xf1, 0xf1, 0x74, 0xf4, 0x73, 0x76,
	0x15, 0x60, 0x10, 0xe2, 0xe7, 0xe2, 0x76, 0xf6, 0x0f, 0x0a, 0xf0, 0x0f, 0x72, 0x0c, 0xf1, 0xf4,
	0xf7, 0x13, 0x60, 0x14, 0xe2, 0xe5, 0xe2, 0x74, 0xf6, 0x70, 0x0c, 0x72, 0x74, 0x0e, 0x71, 0x0d,
	0x12, 0x60, 0x4a, 0x62, 0x03, 0x1b, 0x68, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xed, 0xe9, 0xa4,
	0xb6, 0x60, 0x00, 0x00, 0x00,
}
