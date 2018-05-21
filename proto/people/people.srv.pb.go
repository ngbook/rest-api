// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ngbook/proto/people/people.srv.proto

/*
Package people_srv is a generated protocol buffer package.

It is generated from these files:
	ngbook/proto/people/people.srv.proto

It has these top-level messages:
	People
	GetListReq
	GetListRsp
*/
package people_srv

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

type People struct {
	Id       int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username" json:"username,omitempty"`
	Nickname string `protobuf:"bytes,3,opt,name=nickname" json:"nickname,omitempty"`
	Age      int64  `protobuf:"varint,4,opt,name=age" json:"age,omitempty"`
	Tel      string `protobuf:"bytes,5,opt,name=tel" json:"tel,omitempty"`
	Avatar   string `protobuf:"bytes,6,opt,name=avatar" json:"avatar,omitempty"`
	Friends  string `protobuf:"bytes,7,opt,name=friends" json:"friends,omitempty"`
	Face     string `protobuf:"bytes,8,opt,name=face" json:"face,omitempty"`
}

func (m *People) Reset()                    { *m = People{} }
func (m *People) String() string            { return proto.CompactTextString(m) }
func (*People) ProtoMessage()               {}
func (*People) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *People) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *People) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *People) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *People) GetAge() int64 {
	if m != nil {
		return m.Age
	}
	return 0
}

func (m *People) GetTel() string {
	if m != nil {
		return m.Tel
	}
	return ""
}

func (m *People) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *People) GetFriends() string {
	if m != nil {
		return m.Friends
	}
	return ""
}

func (m *People) GetFace() string {
	if m != nil {
		return m.Face
	}
	return ""
}

type GetListReq struct {
	Start    int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize" json:"pageSize,omitempty"`
}

func (m *GetListReq) Reset()                    { *m = GetListReq{} }
func (m *GetListReq) String() string            { return proto.CompactTextString(m) }
func (*GetListReq) ProtoMessage()               {}
func (*GetListReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetListReq) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *GetListReq) GetPageSize() int64 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type GetListRsp struct {
	People []*People `protobuf:"bytes,1,rep,name=people" json:"people,omitempty"`
}

func (m *GetListRsp) Reset()                    { *m = GetListRsp{} }
func (m *GetListRsp) String() string            { return proto.CompactTextString(m) }
func (*GetListRsp) ProtoMessage()               {}
func (*GetListRsp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetListRsp) GetPeople() []*People {
	if m != nil {
		return m.People
	}
	return nil
}

func init() {
	proto.RegisterType((*People)(nil), "People")
	proto.RegisterType((*GetListReq)(nil), "GetListReq")
	proto.RegisterType((*GetListRsp)(nil), "GetListRsp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PeopleService service

type PeopleServiceClient interface {
	GetList(ctx context.Context, in *GetListReq, opts ...client.CallOption) (*GetListRsp, error)
}

type peopleServiceClient struct {
	c           client.Client
	serviceName string
}

func NewPeopleServiceClient(serviceName string, c client.Client) PeopleServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "peopleservice"
	}
	return &peopleServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *peopleServiceClient) GetList(ctx context.Context, in *GetListReq, opts ...client.CallOption) (*GetListRsp, error) {
	req := c.c.NewRequest(c.serviceName, "PeopleService.GetList", in)
	out := new(GetListRsp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PeopleService service

type PeopleServiceHandler interface {
	GetList(context.Context, *GetListReq, *GetListRsp) error
}

func RegisterPeopleServiceHandler(s server.Server, hdlr PeopleServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&PeopleService{hdlr}, opts...))
}

type PeopleService struct {
	PeopleServiceHandler
}

func (h *PeopleService) GetList(ctx context.Context, in *GetListReq, out *GetListRsp) error {
	return h.PeopleServiceHandler.GetList(ctx, in, out)
}

func init() { proto.RegisterFile("ngbook/proto/people/people.srv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x4e, 0x83, 0x40,
	0x10, 0x86, 0x05, 0x5a, 0xa8, 0xd3, 0x68, 0xcc, 0xc4, 0x98, 0x4d, 0x2f, 0x12, 0xa2, 0x09, 0x17,
	0x69, 0x52, 0x13, 0x8f, 0x5e, 0xbd, 0x78, 0x30, 0xf4, 0x09, 0xb6, 0x30, 0x25, 0x9b, 0x56, 0x58,
	0x77, 0x57, 0x0e, 0x3e, 0x9a, 0x4f, 0xd7, 0x30, 0x4b, 0xcb, 0x89, 0xf9, 0xbe, 0x9f, 0xd9, 0xcc,
	0x0c, 0x3c, 0xb5, 0xcd, 0xae, 0xeb, 0x0e, 0x6b, 0x6d, 0x3a, 0xd7, 0xad, 0x35, 0x75, 0xfa, 0x48,
	0xe3, 0xa7, 0xb0, 0xa6, 0x2f, 0xd8, 0x67, 0xff, 0x01, 0xc4, 0x5f, 0x2c, 0xf1, 0x16, 0x42, 0x55,
	0x8b, 0x20, 0x0d, 0xf2, 0xa8, 0x0c, 0x55, 0x8d, 0x2b, 0x58, 0xfc, 0x5a, 0x32, 0xad, 0xfc, 0x26,
	0x11, 0xa6, 0x41, 0x7e, 0x5d, 0x5e, 0x78, 0xc8, 0x5a, 0x55, 0x1d, 0x38, 0x8b, 0x7c, 0x76, 0x66,
	0xbc, 0x83, 0x48, 0x36, 0x24, 0x66, 0xfc, 0xd0, 0x50, 0x0e, 0xc6, 0xd1, 0x51, 0xcc, 0xf9, 0xc7,
	0xa1, 0xc4, 0x07, 0x88, 0x65, 0x2f, 0x9d, 0x34, 0x22, 0x66, 0x39, 0x12, 0x0a, 0x48, 0xf6, 0x46,
	0x51, 0x5b, 0x5b, 0x91, 0x70, 0x70, 0x46, 0x44, 0x98, 0xed, 0x65, 0x45, 0x62, 0xc1, 0x9a, 0xeb,
	0xec, 0x1d, 0xe0, 0x83, 0xdc, 0xa7, 0xb2, 0xae, 0xa4, 0x1f, 0xbc, 0x87, 0xb9, 0x75, 0xd2, 0xb8,
	0x71, 0x05, 0x0f, 0xc3, 0xa4, 0x5a, 0x36, 0xb4, 0x55, 0x7f, 0x7e, 0x8b, 0xa8, 0xbc, 0x70, 0xf6,
	0x32, 0xf5, 0x5b, 0x8d, 0x8f, 0x10, 0xfb, 0xf3, 0x88, 0x20, 0x8d, 0xf2, 0xe5, 0x26, 0x29, 0xfc,
	0x61, 0xca, 0x51, 0x6f, 0xde, 0xe0, 0xc6, 0x9b, 0x2d, 0x99, 0x5e, 0x55, 0x84, 0xcf, 0x90, 0x8c,
	0xfd, 0xb8, 0x2c, 0xa6, 0x49, 0x56, 0x13, 0x58, 0x9d, 0x5d, 0xed, 0x62, 0x3e, 0xf5, 0xeb, 0x29,
	0x00, 0x00, 0xff, 0xff, 0x66, 0xaa, 0x1e, 0xa3, 0x92, 0x01, 0x00, 0x00,
}
