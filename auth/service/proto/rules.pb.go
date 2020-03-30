// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/service/proto/rules.proto

package auth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Access int32

const (
	Access_UNKNOWN Access = 0
	Access_GRANTED Access = 1
	Access_DENIED  Access = 2
)

var Access_name = map[int32]string{
	0: "UNKNOWN",
	1: "GRANTED",
	2: "DENIED",
}

var Access_value = map[string]int32{
	"UNKNOWN": 0,
	"GRANTED": 1,
	"DENIED":  2,
}

func (x Access) String() string {
	return proto.EnumName(Access_name, int32(x))
}

func (Access) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ce1ef0aa40cdd6dc, []int{0}
}

type Rule struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Role                 string    `protobuf:"bytes,2,opt,name=role,proto3" json:"role,omitempty"`
	Resource             *Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	Access               Access    `protobuf:"varint,4,opt,name=access,proto3,enum=go.micro.auth.Access" json:"access,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Rule) Reset()         { *m = Rule{} }
func (m *Rule) String() string { return proto.CompactTextString(m) }
func (*Rule) ProtoMessage()    {}
func (*Rule) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1ef0aa40cdd6dc, []int{0}
}

func (m *Rule) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rule.Unmarshal(m, b)
}
func (m *Rule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rule.Marshal(b, m, deterministic)
}
func (m *Rule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rule.Merge(m, src)
}
func (m *Rule) XXX_Size() int {
	return xxx_messageInfo_Rule.Size(m)
}
func (m *Rule) XXX_DiscardUnknown() {
	xxx_messageInfo_Rule.DiscardUnknown(m)
}

var xxx_messageInfo_Rule proto.InternalMessageInfo

func (m *Rule) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Rule) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *Rule) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *Rule) GetAccess() Access {
	if m != nil {
		return m.Access
	}
	return Access_UNKNOWN
}

type CreateRequest struct {
	Role                 string    `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Access               Access    `protobuf:"varint,3,opt,name=access,proto3,enum=go.micro.auth.Access" json:"access,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1ef0aa40cdd6dc, []int{1}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *CreateRequest) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *CreateRequest) GetAccess() Access {
	if m != nil {
		return m.Access
	}
	return Access_UNKNOWN
}

type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1ef0aa40cdd6dc, []int{2}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

type DeleteRequest struct {
	Role                 string    `protobuf:"bytes,1,opt,name=role,proto3" json:"role,omitempty"`
	Resource             *Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	Access               Access    `protobuf:"varint,3,opt,name=access,proto3,enum=go.micro.auth.Access" json:"access,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DeleteRequest) Reset()         { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()    {}
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1ef0aa40cdd6dc, []int{3}
}

func (m *DeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteRequest.Unmarshal(m, b)
}
func (m *DeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteRequest.Marshal(b, m, deterministic)
}
func (m *DeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteRequest.Merge(m, src)
}
func (m *DeleteRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteRequest.Size(m)
}
func (m *DeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteRequest proto.InternalMessageInfo

func (m *DeleteRequest) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *DeleteRequest) GetResource() *Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

func (m *DeleteRequest) GetAccess() Access {
	if m != nil {
		return m.Access
	}
	return Access_UNKNOWN
}

type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1ef0aa40cdd6dc, []int{4}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1ef0aa40cdd6dc, []int{5}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

type ListResponse struct {
	Rules                []*Rule  `protobuf:"bytes,1,rep,name=rules,proto3" json:"rules,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce1ef0aa40cdd6dc, []int{6}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRules() []*Rule {
	if m != nil {
		return m.Rules
	}
	return nil
}

func init() {
	proto.RegisterEnum("go.micro.auth.Access", Access_name, Access_value)
	proto.RegisterType((*Rule)(nil), "go.micro.auth.Rule")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.auth.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.auth.CreateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.auth.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.auth.DeleteResponse")
	proto.RegisterType((*ListRequest)(nil), "go.micro.auth.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.auth.ListResponse")
}

func init() { proto.RegisterFile("auth/service/proto/rules.proto", fileDescriptor_ce1ef0aa40cdd6dc) }

var fileDescriptor_ce1ef0aa40cdd6dc = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x93, 0x5f, 0x4b, 0xf3, 0x30,
	0x18, 0xc5, 0x97, 0xae, 0xeb, 0xfb, 0xee, 0xa9, 0x1b, 0x25, 0x22, 0x96, 0xfa, 0x87, 0xb2, 0xab,
	0x2a, 0xd8, 0x41, 0x77, 0x25, 0x5e, 0x4d, 0x3b, 0x86, 0x28, 0x15, 0x82, 0x22, 0x78, 0x37, 0xbb,
	0x07, 0x2d, 0x54, 0x3b, 0x93, 0xd6, 0xaf, 0xe0, 0x9d, 0x9f, 0xd0, 0x0f, 0x23, 0x69, 0xba, 0xe1,
	0x3a, 0x07, 0x7a, 0xe7, 0x4d, 0x48, 0x72, 0x4e, 0x4e, 0x7f, 0x39, 0x6d, 0x61, 0x7f, 0x52, 0xe4,
	0x8f, 0x7d, 0x81, 0xfc, 0x35, 0x89, 0xb1, 0x3f, 0xe3, 0x59, 0x9e, 0xf5, 0x79, 0x91, 0xa2, 0xf0,
	0xcb, 0x39, 0xed, 0x3c, 0x64, 0xfe, 0x53, 0x12, 0xf3, 0xcc, 0x97, 0x46, 0x07, 0xe4, 0xa8, 0xa4,
	0xde, 0x3b, 0x01, 0x9d, 0x15, 0x29, 0xd2, 0x2e, 0x68, 0xc9, 0xd4, 0x26, 0x2e, 0xf1, 0xda, 0x4c,
	0x4b, 0xa6, 0x94, 0x82, 0xce, 0xb3, 0x14, 0x6d, 0xad, 0xdc, 0x29, 0xe7, 0x74, 0x00, 0xff, 0x39,
	0x8a, 0xac, 0xe0, 0x31, 0xda, 0x4d, 0x97, 0x78, 0x66, 0xb0, 0xed, 0x2f, 0x45, 0xfb, 0xac, 0x92,
	0xd9, 0xc2, 0x48, 0x8f, 0xc0, 0x98, 0xc4, 0x31, 0x0a, 0x61, 0xeb, 0x2e, 0xf1, 0xba, 0xc1, 0x56,
	0xed, 0xc8, 0xb0, 0x14, 0x59, 0x65, 0xea, 0xbd, 0x11, 0xe8, 0x9c, 0x71, 0x9c, 0xe4, 0xc8, 0xf0,
	0xa5, 0x40, 0x91, 0x2f, 0x48, 0xc8, 0x1a, 0x12, 0xed, 0xf7, 0x24, 0xcd, 0x9f, 0x90, 0x58, 0xd0,
	0x9d, 0x83, 0x88, 0x59, 0xf6, 0x2c, 0xb0, 0x64, 0x0b, 0x31, 0xc5, 0x3f, 0xc1, 0x36, 0x07, 0xa9,
	0xd8, 0x3a, 0x60, 0x5e, 0x26, 0x22, 0xaf, 0xc0, 0x7a, 0xc7, 0xb0, 0xa1, 0x96, 0x4a, 0xa6, 0x07,
	0xd0, 0x2a, 0xbf, 0x08, 0x9b, 0xb8, 0x4d, 0xcf, 0x0c, 0x36, 0xeb, 0x44, 0x45, 0x8a, 0x4c, 0x39,
	0x0e, 0x7d, 0x30, 0xd4, 0xd3, 0xa8, 0x09, 0xff, 0x6e, 0xa2, 0x8b, 0xe8, 0xea, 0x36, 0xb2, 0x1a,
	0x72, 0x31, 0x66, 0xc3, 0xe8, 0x7a, 0x14, 0x5a, 0x84, 0x02, 0x18, 0xe1, 0x28, 0x3a, 0x1f, 0x85,
	0x96, 0x16, 0x7c, 0x10, 0x68, 0xc9, 0xf3, 0x82, 0x8e, 0xc1, 0x50, 0x8d, 0xd1, 0xdd, 0x5a, 0xfe,
	0xd2, 0x1b, 0x75, 0xf6, 0xd6, 0xa8, 0xd5, 0x55, 0x1a, 0x32, 0x48, 0x5d, 0x6f, 0x25, 0x68, 0xa9,
	0xfe, 0x95, 0xa0, 0x5a, 0x27, 0x0d, 0x3a, 0x04, 0x5d, 0xd6, 0x40, 0x9d, 0x9a, 0xf1, 0x4b, 0x55,
	0xce, 0xce, 0xb7, 0xda, 0x3c, 0xe2, 0xd4, 0xbc, 0x6b, 0xcb, 0xed, 0x13, 0x39, 0xdc, 0x1b, 0xe5,
	0x5f, 0x33, 0xf8, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x63, 0x57, 0x37, 0xef, 0x72, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RulesClient is the client API for Rules service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RulesClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type rulesClient struct {
	cc *grpc.ClientConn
}

func NewRulesClient(cc *grpc.ClientConn) RulesClient {
	return &rulesClient{cc}
}

func (c *rulesClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/go.micro.auth.Rules/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/go.micro.auth.Rules/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rulesClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/go.micro.auth.Rules/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RulesServer is the server API for Rules service.
type RulesServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedRulesServer can be embedded to have forward compatible implementations.
type UnimplementedRulesServer struct {
}

func (*UnimplementedRulesServer) Create(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedRulesServer) Delete(ctx context.Context, req *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (*UnimplementedRulesServer) List(ctx context.Context, req *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}

func RegisterRulesServer(s *grpc.Server, srv RulesServer) {
	s.RegisterService(&_Rules_serviceDesc, srv)
}

func _Rules_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.auth.Rules/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.auth.Rules/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Rules_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RulesServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/go.micro.auth.Rules/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RulesServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rules_serviceDesc = grpc.ServiceDesc{
	ServiceName: "go.micro.auth.Rules",
	HandlerType: (*RulesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Rules_Create_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Rules_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Rules_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/service/proto/rules.proto",
}
