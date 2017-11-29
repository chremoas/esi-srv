// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alliance.proto

/*
Package chremoas_esi is a generated protocol buffer package.

It is generated from these files:
	alliance.proto

It has these top-level messages:
	AllianceRequest
	AllianceResponse
	Alliance
*/
package chremoas_esi

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AllianceRequest struct {
	EntityId int32 `protobuf:"varint,1,opt,name=EntityId" json:"EntityId,omitempty"`
}

func (m *AllianceRequest) Reset()                    { *m = AllianceRequest{} }
func (m *AllianceRequest) String() string            { return proto.CompactTextString(m) }
func (*AllianceRequest) ProtoMessage()               {}
func (*AllianceRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AllianceRequest) GetEntityId() int32 {
	if m != nil {
		return m.EntityId
	}
	return 0
}

type AllianceResponse struct {
	Result *Alliance `protobuf:"bytes,1,opt,name=Result" json:"Result,omitempty"`
}

func (m *AllianceResponse) Reset()                    { *m = AllianceResponse{} }
func (m *AllianceResponse) String() string            { return proto.CompactTextString(m) }
func (*AllianceResponse) ProtoMessage()               {}
func (*AllianceResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AllianceResponse) GetResult() *Alliance {
	if m != nil {
		return m.Result
	}
	return nil
}

type Alliance struct {
	Id           int32  `protobuf:"varint,1,opt,name=Id" json:"Id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Ticker       string `protobuf:"bytes,3,opt,name=Ticker" json:"Ticker,omitempty"`
	ExecutorCorp int32  `protobuf:"varint,4,opt,name=ExecutorCorp" json:"ExecutorCorp,omitempty"`
	DateFounded  int64  `protobuf:"varint,5,opt,name=DateFounded" json:"DateFounded,omitempty"`
}

func (m *Alliance) Reset()                    { *m = Alliance{} }
func (m *Alliance) String() string            { return proto.CompactTextString(m) }
func (*Alliance) ProtoMessage()               {}
func (*Alliance) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Alliance) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Alliance) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Alliance) GetTicker() string {
	if m != nil {
		return m.Ticker
	}
	return ""
}

func (m *Alliance) GetExecutorCorp() int32 {
	if m != nil {
		return m.ExecutorCorp
	}
	return 0
}

func (m *Alliance) GetDateFounded() int64 {
	if m != nil {
		return m.DateFounded
	}
	return 0
}

func init() {
	proto.RegisterType((*AllianceRequest)(nil), "chremoas.esi.AllianceRequest")
	proto.RegisterType((*AllianceResponse)(nil), "chremoas.esi.AllianceResponse")
	proto.RegisterType((*Alliance)(nil), "chremoas.esi.Alliance")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AllianceQuery service

type AllianceQueryClient interface {
	GetAlliance(ctx context.Context, in *AllianceRequest, opts ...client.CallOption) (*AllianceResponse, error)
}

type allianceQueryClient struct {
	c           client.Client
	serviceName string
}

func NewAllianceQueryClient(serviceName string, c client.Client) AllianceQueryClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "chremoas.esi"
	}
	return &allianceQueryClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *allianceQueryClient) GetAlliance(ctx context.Context, in *AllianceRequest, opts ...client.CallOption) (*AllianceResponse, error) {
	req := c.c.NewRequest(c.serviceName, "AllianceQuery.GetAlliance", in)
	out := new(AllianceResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AllianceQuery service

type AllianceQueryHandler interface {
	GetAlliance(context.Context, *AllianceRequest, *AllianceResponse) error
}

func RegisterAllianceQueryHandler(s server.Server, hdlr AllianceQueryHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&AllianceQuery{hdlr}, opts...))
}

type AllianceQuery struct {
	AllianceQueryHandler
}

func (h *AllianceQuery) GetAlliance(ctx context.Context, in *AllianceRequest, out *AllianceResponse) error {
	return h.AllianceQueryHandler.GetAlliance(ctx, in, out)
}

func init() { proto.RegisterFile("alliance.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0x4d, 0xda, 0x86, 0x3a, 0xa9, 0x55, 0xe6, 0x50, 0x96, 0x82, 0x12, 0xf6, 0x94, 0x8b,
	0x7b, 0xa8, 0x4f, 0xe0, 0x47, 0x95, 0x5e, 0x0a, 0x2e, 0xde, 0x65, 0x4d, 0x06, 0x0c, 0xa6, 0xd9,
	0xb8, 0x1f, 0x60, 0xdf, 0xc1, 0x87, 0x16, 0xb6, 0x69, 0x6c, 0x0f, 0xbd, 0xed, 0xfc, 0xe6, 0xc7,
	0xce, 0xcc, 0x1f, 0xa6, 0xaa, 0xae, 0x2b, 0xd5, 0x14, 0x24, 0x5a, 0xa3, 0x9d, 0xc6, 0x49, 0xf1,
	0x69, 0x68, 0xa3, 0x95, 0x15, 0x64, 0x2b, 0x7e, 0x0b, 0x97, 0xf7, 0x5d, 0x5f, 0xd2, 0xb7, 0x27,
	0xeb, 0x70, 0x0e, 0xe3, 0x65, 0xe3, 0x2a, 0xb7, 0x5d, 0x95, 0x2c, 0xca, 0xa2, 0x7c, 0x24, 0xfb,
	0x9a, 0x3f, 0xc0, 0xd5, 0xbf, 0x6e, 0x5b, 0xdd, 0x58, 0x42, 0x01, 0x89, 0x24, 0xeb, 0x6b, 0x17,
	0xec, 0x74, 0x31, 0x13, 0x87, 0x13, 0x44, 0xef, 0x77, 0x16, 0xff, 0x8d, 0x60, 0xbc, 0x87, 0x38,
	0x85, 0xb8, 0x1f, 0x13, 0xaf, 0x4a, 0x44, 0x18, 0xae, 0xd5, 0x86, 0x58, 0x9c, 0x45, 0xf9, 0xb9,
	0x0c, 0x6f, 0x9c, 0x41, 0xf2, 0x56, 0x15, 0x5f, 0x64, 0xd8, 0x20, 0xd0, 0xae, 0x42, 0x0e, 0x93,
	0xe5, 0x0f, 0x15, 0xde, 0x69, 0xf3, 0xa8, 0x4d, 0xcb, 0x86, 0xe1, 0x97, 0x23, 0x86, 0x19, 0xa4,
	0x4f, 0xca, 0xd1, 0xb3, 0xf6, 0x4d, 0x49, 0x25, 0x1b, 0x65, 0x51, 0x3e, 0x90, 0x87, 0x68, 0xf1,
	0x0e, 0x17, 0xfb, 0x6d, 0x5e, 0x3d, 0x99, 0x2d, 0xae, 0x21, 0x7d, 0x21, 0xd7, 0x6f, 0x78, 0x7d,
	0xe2, 0x9c, 0x5d, 0x5a, 0xf3, 0x9b, 0x53, 0xed, 0x5d, 0x3a, 0xfc, 0xec, 0x23, 0x09, 0xb9, 0xdf,
	0xfd, 0x05, 0x00, 0x00, 0xff, 0xff, 0x24, 0xd3, 0x03, 0xf3, 0x89, 0x01, 0x00, 0x00,
}
