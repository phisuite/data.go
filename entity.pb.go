// Code generated by protoc-gen-go. DO NOT EDIT.
// source: entity.proto

package data

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Entity struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string              `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Transaction          string              `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Data                 map[string]*any.Any `protobuf:"bytes,10,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Entity) Reset()         { *m = Entity{} }
func (m *Entity) String() string { return proto.CompactTextString(m) }
func (*Entity) ProtoMessage()    {}
func (*Entity) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf50d946d740d100, []int{0}
}

func (m *Entity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entity.Unmarshal(m, b)
}
func (m *Entity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entity.Marshal(b, m, deterministic)
}
func (m *Entity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entity.Merge(m, src)
}
func (m *Entity) XXX_Size() int {
	return xxx_messageInfo_Entity.Size(m)
}
func (m *Entity) XXX_DiscardUnknown() {
	xxx_messageInfo_Entity.DiscardUnknown(m)
}

var xxx_messageInfo_Entity proto.InternalMessageInfo

func (m *Entity) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Entity) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Entity) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Entity) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

func (m *Entity) GetData() map[string]*any.Any {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Entity)(nil), "data.Entity")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.Entity.DataEntry")
}

func init() {
	proto.RegisterFile("entity.proto", fileDescriptor_cf50d946d740d100)
}

var fileDescriptor_cf50d946d740d100 = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xdf, 0x4a, 0xc3, 0x30,
	0x14, 0xc6, 0x49, 0x5b, 0x27, 0x3b, 0x9d, 0xa2, 0x07, 0x91, 0xba, 0xab, 0x52, 0x14, 0xc6, 0x2e,
	0x32, 0xa9, 0x37, 0xe2, 0xdd, 0xa6, 0x43, 0x04, 0xc5, 0x52, 0xf0, 0x01, 0xb2, 0x36, 0xce, 0xe0,
	0x9a, 0x94, 0x36, 0x1d, 0xf4, 0x75, 0x7c, 0x35, 0x5f, 0x44, 0x9a, 0x6c, 0x32, 0x11, 0xf4, 0x2e,
	0xdf, 0x9f, 0x70, 0x7e, 0x1f, 0x0c, 0xb8, 0xd4, 0x42, 0xb7, 0xb4, 0xac, 0x94, 0x56, 0xe8, 0xe5,
	0x4c, 0xb3, 0xe1, 0xd9, 0x52, 0xa9, 0xe5, 0x8a, 0x4f, 0x8c, 0xb7, 0x68, 0x5e, 0x27, 0x4c, 0x6e,
	0x0a, 0xc3, 0x41, 0xa6, 0x8a, 0x42, 0x49, 0xab, 0xa2, 0x4f, 0x02, 0xbd, 0xb9, 0xf9, 0x8f, 0x87,
	0xe0, 0x88, 0x3c, 0x20, 0x21, 0x19, 0xf5, 0x53, 0x47, 0xe4, 0x88, 0xe0, 0x49, 0x56, 0xf0, 0xc0,
	0x31, 0x8e, 0x79, 0x63, 0x00, 0xfb, 0x6b, 0x5e, 0xd5, 0x42, 0xc9, 0xc0, 0x35, 0xf6, 0x56, 0x62,
	0x08, 0xbe, 0xae, 0x98, 0xac, 0x59, 0xa6, 0xbb, 0xd4, 0x33, 0xe9, 0xae, 0x85, 0x63, 0x30, 0x6c,
	0x01, 0x84, 0xee, 0xc8, 0x8f, 0x4f, 0x69, 0x27, 0xa8, 0xbd, 0x4d, 0xef, 0x98, 0x66, 0x73, 0xa9,
	0xab, 0x36, 0xb5, 0xfc, 0x4f, 0xd0, 0xff, 0xb6, 0xf0, 0x08, 0xdc, 0x77, 0xde, 0x6e, 0xc8, 0xba,
	0x27, 0x8e, 0x61, 0x6f, 0xcd, 0x56, 0x8d, 0x65, 0xf3, 0xe3, 0x13, 0x6a, 0xe7, 0xd2, 0xed, 0x5c,
	0x3a, 0x95, 0x6d, 0x6a, 0x2b, 0x37, 0xce, 0x35, 0x89, 0x3f, 0x08, 0xf4, 0xed, 0xa5, 0x69, 0xf2,
	0x80, 0x17, 0xe0, 0x3d, 0x8a, 0x5a, 0xe3, 0x81, 0x45, 0x78, 0x2e, 0x3b, 0xbc, 0x7a, 0x38, 0xd8,
	0x25, 0xba, 0x24, 0x18, 0x81, 0x7b, 0xcf, 0xff, 0x6e, 0xe1, 0x39, 0xf4, 0x6e, 0x2b, 0xce, 0x34,
	0xc7, 0x1f, 0xfe, 0xef, 0xd6, 0x4b, 0x99, 0xff, 0xd3, 0x9a, 0x45, 0x70, 0x9c, 0xa9, 0x82, 0x96,
	0x6f, 0xa2, 0x6e, 0x84, 0xe6, 0x26, 0x9b, 0xf9, 0x36, 0x4c, 0xba, 0x61, 0x09, 0x59, 0xf4, 0xcc,
	0xc2, 0xab, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe5, 0x98, 0xdc, 0xaa, 0xf4, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EntityAPIClient is the client API for EntityAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EntityAPIClient interface {
	List(ctx context.Context, in *Options, opts ...grpc.CallOption) (EntityAPI_ListClient, error)
	Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Entity, error)
	Create(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error)
	Update(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error)
}

type entityAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewEntityAPIClient(cc grpc.ClientConnInterface) EntityAPIClient {
	return &entityAPIClient{cc}
}

func (c *entityAPIClient) List(ctx context.Context, in *Options, opts ...grpc.CallOption) (EntityAPI_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_EntityAPI_serviceDesc.Streams[0], "/data.EntityAPI/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &entityAPIListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EntityAPI_ListClient interface {
	Recv() (*Entity, error)
	grpc.ClientStream
}

type entityAPIListClient struct {
	grpc.ClientStream
}

func (x *entityAPIListClient) Recv() (*Entity, error) {
	m := new(Entity)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *entityAPIClient) Get(ctx context.Context, in *Options, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/data.EntityAPI/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityAPIClient) Create(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/data.EntityAPI/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *entityAPIClient) Update(ctx context.Context, in *Entity, opts ...grpc.CallOption) (*Entity, error) {
	out := new(Entity)
	err := c.cc.Invoke(ctx, "/data.EntityAPI/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EntityAPIServer is the server API for EntityAPI service.
type EntityAPIServer interface {
	List(*Options, EntityAPI_ListServer) error
	Get(context.Context, *Options) (*Entity, error)
	Create(context.Context, *Entity) (*Entity, error)
	Update(context.Context, *Entity) (*Entity, error)
}

// UnimplementedEntityAPIServer can be embedded to have forward compatible implementations.
type UnimplementedEntityAPIServer struct {
}

func (*UnimplementedEntityAPIServer) List(req *Options, srv EntityAPI_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedEntityAPIServer) Get(ctx context.Context, req *Options) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedEntityAPIServer) Create(ctx context.Context, req *Entity) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedEntityAPIServer) Update(ctx context.Context, req *Entity) (*Entity, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterEntityAPIServer(s *grpc.Server, srv EntityAPIServer) {
	s.RegisterService(&_EntityAPI_serviceDesc, srv)
}

func _EntityAPI_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EntityAPIServer).List(m, &entityAPIListServer{stream})
}

type EntityAPI_ListServer interface {
	Send(*Entity) error
	grpc.ServerStream
}

type entityAPIListServer struct {
	grpc.ServerStream
}

func (x *entityAPIListServer) Send(m *Entity) error {
	return x.ServerStream.SendMsg(m)
}

func _EntityAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Options)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.EntityAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityAPIServer).Get(ctx, req.(*Options))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityAPI_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityAPIServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.EntityAPI/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityAPIServer).Create(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

func _EntityAPI_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Entity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EntityAPIServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.EntityAPI/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EntityAPIServer).Update(ctx, req.(*Entity))
	}
	return interceptor(ctx, in, info, handler)
}

var _EntityAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.EntityAPI",
	HandlerType: (*EntityAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _EntityAPI_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _EntityAPI_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _EntityAPI_Update_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "List",
			Handler:       _EntityAPI_List_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "entity.proto",
}
