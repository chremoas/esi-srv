// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

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

type SearchRequest struct {
	SearchString string `protobuf:"bytes,1,opt,name=SearchString" json:"SearchString,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *SearchRequest) GetSearchString() string {
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
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

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
	proto.RegisterType((*SearchRequest)(nil), "chremoas.esi.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "chremoas.esi.SearchResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for SearchQuery service

type SearchQueryClient interface {
	GetSearch(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
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

func (c *searchQueryClient) GetSearch(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
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
	GetSearch(context.Context, *SearchRequest, *SearchResponse) error
}

func RegisterSearchQueryHandler(s server.Server, hdlr SearchQueryHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&SearchQuery{hdlr}, opts...))
}

type SearchQuery struct {
	SearchQueryHandler
}

func (h *SearchQuery) GetSearch(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.SearchQueryHandler.GetSearch(ctx, in, out)
}

func init() { proto.RegisterFile("search.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4d, 0x6b, 0xf2, 0x40,
	0x10, 0xc7, 0x1f, 0xf5, 0xf1, 0x25, 0xa3, 0xf6, 0xb0, 0x94, 0xb2, 0x58, 0x0f, 0x12, 0x7a, 0xf0,
	0x94, 0x43, 0xfd, 0x04, 0x22, 0xf4, 0xe5, 0xd8, 0xe4, 0xe0, 0x79, 0x1b, 0x06, 0x13, 0x88, 0x3b,
	0xe9, 0xec, 0x58, 0xc8, 0x37, 0xed, 0xc7, 0x29, 0xd9, 0xd5, 0xd6, 0x40, 0x8f, 0xff, 0xdf, 0x6f,
	0x96, 0xff, 0x32, 0x03, 0x33, 0x87, 0x86, 0xf3, 0x22, 0xa9, 0x99, 0x84, 0xd4, 0x2c, 0x2f, 0x18,
	0x8f, 0x64, 0x5c, 0x82, 0xae, 0x8c, 0x37, 0x30, 0xcf, 0xbc, 0x4d, 0xf1, 0xe3, 0x84, 0x4e, 0x54,
	0x0c, 0xb3, 0x00, 0x32, 0xe1, 0xd2, 0x1e, 0x74, 0x6f, 0xd5, 0x5b, 0x47, 0x69, 0x87, 0xc5, 0x5f,
	0x7d, 0xb8, 0xb9, 0xbc, 0x72, 0x35, 0x59, 0x87, 0xea, 0x16, 0x86, 0xdb, 0x03, 0x5a, 0xd1, 0xbd,
	0xd5, 0x60, 0x3d, 0x4c, 0x43, 0x50, 0x0b, 0x98, 0x6c, 0xab, 0xaa, 0x34, 0x36, 0x47, 0xdd, 0xf7,
	0xe2, 0x27, 0xab, 0x25, 0x44, 0xbb, 0xc2, 0xb0, 0xc9, 0x05, 0x59, 0x0f, 0xbc, 0xfc, 0x05, 0xea,
	0x01, 0xe6, 0x3b, 0xb2, 0x4e, 0xb0, 0xaa, 0x8c, 0x94, 0x64, 0xf5, 0x7f, 0x3f, 0xd1, 0x85, 0x6a,
	0x05, 0xd3, 0x1d, 0x71, 0x4d, 0x1c, 0x66, 0x86, 0x7e, 0xe6, 0x1a, 0x29, 0x0d, 0xe3, 0x27, 0x93,
	0x7b, 0x3b, 0xf2, 0xf6, 0x12, 0xdb, 0x86, 0x57, 0xfb, 0x89, 0x56, 0x88, 0x1b, 0x69, 0x6a, 0xd4,
	0xe3, 0xd0, 0xd0, 0x81, 0xea, 0x0e, 0x46, 0x29, 0x1e, 0xda, 0xe7, 0x13, 0xaf, 0xcf, 0xa9, 0x6d,
	0xce, 0xa8, 0x32, 0xec, 0x1a, 0x27, 0x78, 0xd4, 0x51, 0x68, 0xbe, 0x42, 0x6d, 0x73, 0x26, 0xe1,
	0x5f, 0x10, 0x9a, 0xcf, 0xb1, 0xdd, 0xca, 0x9e, 0xf8, 0x58, 0x50, 0x85, 0x7a, 0x1a, 0xb6, 0x72,
	0xc9, 0x8f, 0x7b, 0x98, 0x86, 0xcd, 0xbe, 0x9d, 0x90, 0x1b, 0xf5, 0x02, 0xd1, 0x33, 0x4a, 0x20,
	0xea, 0x3e, 0xb9, 0x3e, 0x5d, 0xd2, 0xb9, 0xdb, 0x62, 0xf9, 0xb7, 0x0c, 0xe7, 0x89, 0xff, 0xbd,
	0x8f, 0xfc, 0xf5, 0x37, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x07, 0xe7, 0x73, 0xc3, 0x0d, 0x02,
	0x00, 0x00,
}
