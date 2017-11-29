// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

/*
Package chremoas_esi is a generated protocol buffer package.

It is generated from these files:
	search.proto

It has these top-level messages:
	EntitySearchRequest
	SearchResponse
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

type EntitySearchRequest struct {
	SearchString string `protobuf:"bytes,1,opt,name=SearchString" json:"SearchString,omitempty"`
}

func (m *EntitySearchRequest) Reset()                    { *m = EntitySearchRequest{} }
func (m *EntitySearchRequest) String() string            { return proto.CompactTextString(m) }
func (*EntitySearchRequest) ProtoMessage()               {}
func (*EntitySearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *EntitySearchRequest) GetSearchString() string {
	if m != nil {
		return m.SearchString
	}
	return ""
}

type SearchResponse struct {
	Agent         []int32 `protobuf:"varint,1,rep,packed,name=Agent" json:"Agent,omitempty"`
	Alliance      []int32 `protobuf:"varint,2,rep,packed,name=Alliance" json:"Alliance,omitempty"`
	Character     []int32 `protobuf:"varint,3,rep,packed,name=Character" json:"Character,omitempty"`
	Constellation []int32 `protobuf:"varint,4,rep,packed,name=Constellation" json:"Constellation,omitempty"`
	Corporation   []int32 `protobuf:"varint,5,rep,packed,name=Corporation" json:"Corporation,omitempty"`
	Faction       []int32 `protobuf:"varint,6,rep,packed,name=Faction" json:"Faction,omitempty"`
	Inventorytype []int32 `protobuf:"varint,7,rep,packed,name=Inventorytype" json:"Inventorytype,omitempty"`
	Region        []int32 `protobuf:"varint,8,rep,packed,name=Region" json:"Region,omitempty"`
	Solarsystem   []int32 `protobuf:"varint,9,rep,packed,name=Solarsystem" json:"Solarsystem,omitempty"`
	Station       []int32 `protobuf:"varint,10,rep,packed,name=Station" json:"Station,omitempty"`
	Wormhole      []int32 `protobuf:"varint,11,rep,packed,name=Wormhole" json:"Wormhole,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SearchResponse) GetAgent() []int32 {
	if m != nil {
		return m.Agent
	}
	return nil
}

func (m *SearchResponse) GetAlliance() []int32 {
	if m != nil {
		return m.Alliance
	}
	return nil
}

func (m *SearchResponse) GetCharacter() []int32 {
	if m != nil {
		return m.Character
	}
	return nil
}

func (m *SearchResponse) GetConstellation() []int32 {
	if m != nil {
		return m.Constellation
	}
	return nil
}

func (m *SearchResponse) GetCorporation() []int32 {
	if m != nil {
		return m.Corporation
	}
	return nil
}

func (m *SearchResponse) GetFaction() []int32 {
	if m != nil {
		return m.Faction
	}
	return nil
}

func (m *SearchResponse) GetInventorytype() []int32 {
	if m != nil {
		return m.Inventorytype
	}
	return nil
}

func (m *SearchResponse) GetRegion() []int32 {
	if m != nil {
		return m.Region
	}
	return nil
}

func (m *SearchResponse) GetSolarsystem() []int32 {
	if m != nil {
		return m.Solarsystem
	}
	return nil
}

func (m *SearchResponse) GetStation() []int32 {
	if m != nil {
		return m.Station
	}
	return nil
}

func (m *SearchResponse) GetWormhole() []int32 {
	if m != nil {
		return m.Wormhole
	}
	return nil
}

func init() {
	proto.RegisterType((*EntitySearchRequest)(nil), "chremoas.esi.EntitySearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "chremoas.esi.SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SearchQuery service

type SearchQueryClient interface {
	GetSearch(ctx context.Context, in *EntitySearchRequest, opts ...client.CallOption) (*SearchResponse, error)
}

type searchQueryClient struct {
	c           client.Client
	serviceName string
}

func NewSearchQueryClient(serviceName string, c client.Client) SearchQueryClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "chremoas.esi"
	}
	return &searchQueryClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *searchQueryClient) GetSearch(ctx context.Context, in *EntitySearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "SearchQuery.GetSearch", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SearchQuery service

type SearchQueryHandler interface {
	GetSearch(context.Context, *EntitySearchRequest, *SearchResponse) error
}

func RegisterSearchQueryHandler(s server.Server, hdlr SearchQueryHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&SearchQuery{hdlr}, opts...))
}

type SearchQuery struct {
	SearchQueryHandler
}

func (h *SearchQuery) GetSearch(ctx context.Context, in *EntitySearchRequest, out *SearchResponse) error {
	return h.SearchQueryHandler.GetSearch(ctx, in, out)
}

func init() { proto.RegisterFile("search.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0x69, 0x4b, 0xff, 0xe4, 0x5a, 0x18, 0x0c, 0x42, 0x56, 0xd5, 0xa1, 0x44, 0x0c, 0x9d,
	0x32, 0xc0, 0xc4, 0x58, 0x55, 0x80, 0x58, 0x90, 0x48, 0x06, 0x26, 0x06, 0x13, 0x9d, 0x9a, 0x48,
	0xae, 0x1d, 0xce, 0x57, 0xa4, 0x7c, 0x53, 0x3e, 0x0e, 0x8a, 0xdd, 0x42, 0x22, 0x31, 0xbe, 0xdf,
	0x7b, 0xa7, 0x67, 0xdd, 0x19, 0x66, 0x0e, 0x15, 0xe5, 0x45, 0x52, 0x91, 0x65, 0x2b, 0x66, 0x79,
	0x41, 0xb8, 0xb3, 0xca, 0x25, 0xe8, 0xca, 0xf8, 0x1e, 0x2e, 0x1e, 0x0c, 0x97, 0x5c, 0x67, 0x3e,
	0x93, 0xe2, 0xe7, 0x1e, 0x1d, 0x8b, 0x18, 0x66, 0x01, 0x64, 0x4c, 0xa5, 0xd9, 0xca, 0xde, 0xb2,
	0xb7, 0x8a, 0xd2, 0x0e, 0x8b, 0xbf, 0xfb, 0x70, 0x7e, 0x9c, 0x72, 0x95, 0x35, 0x0e, 0xc5, 0x25,
	0x0c, 0xd7, 0x5b, 0x34, 0x2c, 0x7b, 0xcb, 0xc1, 0x6a, 0x98, 0x06, 0x21, 0xe6, 0x30, 0x59, 0x6b,
	0x5d, 0x2a, 0x93, 0xa3, 0xec, 0x7b, 0xe3, 0x57, 0x8b, 0x05, 0x44, 0x9b, 0x42, 0x91, 0xca, 0x19,
	0x49, 0x0e, 0xbc, 0xf9, 0x07, 0xc4, 0x0d, 0x9c, 0x6d, 0xac, 0x71, 0x8c, 0x5a, 0x2b, 0x2e, 0xad,
	0x91, 0xa7, 0x3e, 0xd1, 0x85, 0x62, 0x09, 0xd3, 0x8d, 0xa5, 0xca, 0x52, 0xc8, 0x0c, 0x7d, 0xa6,
	0x8d, 0x84, 0x84, 0xf1, 0xa3, 0xca, 0xbd, 0x3b, 0xf2, 0xee, 0x51, 0x36, 0x0d, 0xcf, 0xe6, 0x0b,
	0x0d, 0x5b, 0xaa, 0xb9, 0xae, 0x50, 0x8e, 0x43, 0x43, 0x07, 0x8a, 0x2b, 0x18, 0xa5, 0xb8, 0x6d,
	0xc6, 0x27, 0xde, 0x3e, 0xa8, 0xa6, 0x39, 0xb3, 0x5a, 0x91, 0xab, 0x1d, 0xe3, 0x4e, 0x46, 0xa1,
	0xb9, 0x85, 0x9a, 0xe6, 0x8c, 0xc3, 0xbb, 0x20, 0x34, 0x1f, 0x64, 0xb3, 0x95, 0x37, 0x4b, 0xbb,
	0xc2, 0x6a, 0x94, 0xd3, 0xb0, 0x95, 0xa3, 0xbe, 0x7d, 0x87, 0x69, 0xd8, 0xec, 0xeb, 0x1e, 0xa9,
	0x16, 0x2f, 0x10, 0x3d, 0x21, 0x07, 0x22, 0xae, 0x93, 0xf6, 0x01, 0x93, 0x7f, 0xae, 0x37, 0x5f,
	0x74, 0x23, 0xdd, 0x23, 0xc5, 0x27, 0x1f, 0x23, 0xff, 0x13, 0xee, 0x7e, 0x02, 0x00, 0x00, 0xff,
	0xff, 0xfd, 0x4a, 0x1e, 0xe6, 0x19, 0x02, 0x00, 0x00,
}
