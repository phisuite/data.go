// Code generated by protoc-gen-go. DO NOT EDIT.
// source: process.proto

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

type Process struct {
	Id                   string           `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Version              string           `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	Transaction          string           `protobuf:"bytes,4,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Started              string           `protobuf:"bytes,5,opt,name=started,proto3" json:"started,omitempty"`
	Completed            string           `protobuf:"bytes,6,opt,name=completed,proto3" json:"completed,omitempty"`
	In                   *Process_Payload `protobuf:"bytes,10,opt,name=in,proto3" json:"in,omitempty"`
	Out                  *Process_Payload `protobuf:"bytes,11,opt,name=out,proto3" json:"out,omitempty"`
	Error                *Process_Error   `protobuf:"bytes,12,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *Process) Reset()         { *m = Process{} }
func (m *Process) String() string { return proto.CompactTextString(m) }
func (*Process) ProtoMessage()    {}
func (*Process) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0}
}

func (m *Process) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process.Unmarshal(m, b)
}
func (m *Process) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process.Marshal(b, m, deterministic)
}
func (m *Process) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process.Merge(m, src)
}
func (m *Process) XXX_Size() int {
	return xxx_messageInfo_Process.Size(m)
}
func (m *Process) XXX_DiscardUnknown() {
	xxx_messageInfo_Process.DiscardUnknown(m)
}

var xxx_messageInfo_Process proto.InternalMessageInfo

func (m *Process) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Process) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Process) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Process) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

func (m *Process) GetStarted() string {
	if m != nil {
		return m.Started
	}
	return ""
}

func (m *Process) GetCompleted() string {
	if m != nil {
		return m.Completed
	}
	return ""
}

func (m *Process) GetIn() *Process_Payload {
	if m != nil {
		return m.In
	}
	return nil
}

func (m *Process) GetOut() *Process_Payload {
	if m != nil {
		return m.Out
	}
	return nil
}

func (m *Process) GetError() *Process_Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type Process_Payload struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Entity               *Entity  `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Payload) Reset()         { *m = Process_Payload{} }
func (m *Process_Payload) String() string { return proto.CompactTextString(m) }
func (*Process_Payload) ProtoMessage()    {}
func (*Process_Payload) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 0}
}

func (m *Process_Payload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Payload.Unmarshal(m, b)
}
func (m *Process_Payload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Payload.Marshal(b, m, deterministic)
}
func (m *Process_Payload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Payload.Merge(m, src)
}
func (m *Process_Payload) XXX_Size() int {
	return xxx_messageInfo_Process_Payload.Size(m)
}
func (m *Process_Payload) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Payload.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Payload proto.InternalMessageInfo

func (m *Process_Payload) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *Process_Payload) GetEntity() *Entity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type Process_Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Event                *Event   `protobuf:"bytes,3,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Process_Error) Reset()         { *m = Process_Error{} }
func (m *Process_Error) String() string { return proto.CompactTextString(m) }
func (*Process_Error) ProtoMessage()    {}
func (*Process_Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{0, 1}
}

func (m *Process_Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Process_Error.Unmarshal(m, b)
}
func (m *Process_Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Process_Error.Marshal(b, m, deterministic)
}
func (m *Process_Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Process_Error.Merge(m, src)
}
func (m *Process_Error) XXX_Size() int {
	return xxx_messageInfo_Process_Error.Size(m)
}
func (m *Process_Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Process_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Process_Error proto.InternalMessageInfo

func (m *Process_Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Process_Error) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Process_Error) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

type ProcessRequest struct {
	Name                 string              `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string              `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Transaction          string              `protobuf:"bytes,3,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Event                map[string]*any.Any `protobuf:"bytes,10,rep,name=event,proto3" json:"event,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Entity               map[string]*any.Any `protobuf:"bytes,11,rep,name=entity,proto3" json:"entity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProcessRequest) Reset()         { *m = ProcessRequest{} }
func (m *ProcessRequest) String() string { return proto.CompactTextString(m) }
func (*ProcessRequest) ProtoMessage()    {}
func (*ProcessRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{1}
}

func (m *ProcessRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessRequest.Unmarshal(m, b)
}
func (m *ProcessRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessRequest.Marshal(b, m, deterministic)
}
func (m *ProcessRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessRequest.Merge(m, src)
}
func (m *ProcessRequest) XXX_Size() int {
	return xxx_messageInfo_ProcessRequest.Size(m)
}
func (m *ProcessRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessRequest proto.InternalMessageInfo

func (m *ProcessRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ProcessRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ProcessRequest) GetTransaction() string {
	if m != nil {
		return m.Transaction
	}
	return ""
}

func (m *ProcessRequest) GetEvent() map[string]*any.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ProcessRequest) GetEntity() map[string]*any.Any {
	if m != nil {
		return m.Entity
	}
	return nil
}

type ProcessResponse struct {
	Id                   string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Event                map[string]*any.Any `protobuf:"bytes,10,rep,name=event,proto3" json:"event,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Entity               map[string]*any.Any `protobuf:"bytes,11,rep,name=entity,proto3" json:"entity,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProcessResponse) Reset()         { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()    {}
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_54c4d0e8c0aaf5c3, []int{2}
}

func (m *ProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessResponse.Unmarshal(m, b)
}
func (m *ProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessResponse.Marshal(b, m, deterministic)
}
func (m *ProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessResponse.Merge(m, src)
}
func (m *ProcessResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessResponse.Size(m)
}
func (m *ProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessResponse proto.InternalMessageInfo

func (m *ProcessResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ProcessResponse) GetEvent() map[string]*any.Any {
	if m != nil {
		return m.Event
	}
	return nil
}

func (m *ProcessResponse) GetEntity() map[string]*any.Any {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*Process)(nil), "data.Process")
	proto.RegisterType((*Process_Payload)(nil), "data.Process.Payload")
	proto.RegisterType((*Process_Error)(nil), "data.Process.Error")
	proto.RegisterType((*ProcessRequest)(nil), "data.ProcessRequest")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.ProcessRequest.EntityEntry")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.ProcessRequest.EventEntry")
	proto.RegisterType((*ProcessResponse)(nil), "data.ProcessResponse")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.ProcessResponse.EntityEntry")
	proto.RegisterMapType((map[string]*any.Any)(nil), "data.ProcessResponse.EventEntry")
}

func init() {
	proto.RegisterFile("process.proto", fileDescriptor_54c4d0e8c0aaf5c3)
}

var fileDescriptor_54c4d0e8c0aaf5c3 = []byte{
	// 570 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x93, 0xdf, 0x6e, 0xd3, 0x30,
	0x14, 0xc6, 0x95, 0xa4, 0x7f, 0xb4, 0x93, 0x6e, 0x80, 0xd9, 0xa4, 0x10, 0x21, 0xd1, 0x0d, 0x18,
	0x1d, 0x17, 0xd9, 0x54, 0x04, 0x1a, 0xbb, 0xdb, 0xa4, 0x22, 0x21, 0x21, 0x56, 0x85, 0x1b, 0x6e,
	0xbd, 0xe4, 0x30, 0x22, 0x1a, 0x3b, 0xd8, 0x4e, 0x45, 0x1f, 0x8c, 0x97, 0xe0, 0x61, 0xb8, 0xe0,
	0x09, 0x90, 0xff, 0xa4, 0x5d, 0xbb, 0x6e, 0x20, 0x71, 0xc5, 0x5d, 0x7c, 0xbe, 0xef, 0xf3, 0xf1,
	0xf9, 0x39, 0x86, 0xcd, 0x4a, 0xf0, 0x0c, 0xa5, 0x4c, 0x2a, 0xc1, 0x15, 0x27, 0xad, 0x9c, 0x2a,
	0x1a, 0x3f, 0xb8, 0xe4, 0xfc, 0x72, 0x82, 0x87, 0xa6, 0x76, 0x51, 0x7f, 0x3a, 0xa4, 0x6c, 0x66,
	0x0d, 0x71, 0x2f, 0xe3, 0x65, 0xc9, 0x99, 0x5b, 0x85, 0x38, 0x45, 0xa6, 0x1a, 0x09, 0x99, 0x2a,
	0x94, 0x33, 0xee, 0xfd, 0x08, 0xa0, 0x3b, 0xb6, 0x7b, 0x93, 0x2d, 0xf0, 0x8b, 0x3c, 0xf2, 0xfa,
	0xde, 0x60, 0x23, 0xf5, 0x8b, 0x9c, 0x10, 0x68, 0x31, 0x5a, 0x62, 0xe4, 0x9b, 0x8a, 0xf9, 0x26,
	0x11, 0x74, 0xa7, 0x28, 0x64, 0xc1, 0x59, 0x14, 0x98, 0x72, 0xb3, 0x24, 0x7d, 0x08, 0x95, 0xa0,
	0x4c, 0xd2, 0x4c, 0x69, 0xb5, 0x65, 0xd4, 0xab, 0x25, 0x9d, 0x95, 0x8a, 0x0a, 0x85, 0x79, 0xd4,
	0xb6, 0x59, 0xb7, 0x24, 0x0f, 0x61, 0x23, 0xe3, 0x65, 0x35, 0x41, 0xad, 0x75, 0x8c, 0xb6, 0x28,
	0x90, 0xa7, 0xe0, 0x17, 0x2c, 0x82, 0xbe, 0x37, 0x08, 0x87, 0x3b, 0x89, 0x1e, 0x3d, 0x71, 0x47,
	0x4e, 0xc6, 0x74, 0x36, 0xe1, 0x34, 0x4f, 0xfd, 0x82, 0x91, 0x67, 0x10, 0xf0, 0x5a, 0x45, 0xe1,
	0x6d, 0x3e, 0xed, 0x20, 0x07, 0xd0, 0x46, 0x21, 0xb8, 0x88, 0x7a, 0xc6, 0x7a, 0x7f, 0xd9, 0x3a,
	0xd2, 0x52, 0x6a, 0x1d, 0x71, 0x0a, 0x5d, 0x17, 0x25, 0xbb, 0xd0, 0x36, 0x18, 0x0d, 0xa0, 0x70,
	0x18, 0xda, 0xd4, 0x48, 0x97, 0x52, 0xab, 0x90, 0x27, 0xd0, 0xb1, 0x70, 0x0d, 0xb2, 0x70, 0xd8,
	0x73, 0x1e, 0x53, 0x4b, 0x9d, 0x16, 0x7f, 0x84, 0xb6, 0xe9, 0xa1, 0xf9, 0x66, 0x3c, 0x47, 0x47,
	0xdc, 0x7c, 0x6b, 0x46, 0x25, 0x4a, 0x49, 0x2f, 0x1b, 0xec, 0xcd, 0x72, 0xd1, 0x3f, 0xb8, 0xa9,
	0xff, 0xde, 0x2f, 0x1f, 0xb6, 0xdc, 0x18, 0x29, 0x7e, 0xad, 0x51, 0xaa, 0xf9, 0x1d, 0x7a, 0xeb,
	0xef, 0xd0, 0xbf, 0xf5, 0x0e, 0x83, 0xeb, 0x77, 0xf8, 0xb2, 0x39, 0x05, 0xf4, 0x83, 0x41, 0x38,
	0x7c, 0xb4, 0xc4, 0xce, 0x35, 0xb5, 0x87, 0x1a, 0x31, 0x25, 0x66, 0x0d, 0x99, 0xe3, 0x39, 0x99,
	0xd0, 0xe4, 0xfa, 0xeb, 0x73, 0xc6, 0x62, 0x83, 0x0d, 0xad, 0xf7, 0x00, 0x8b, 0xed, 0xc8, 0x5d,
	0x08, 0xbe, 0xe0, 0xcc, 0x4d, 0xa3, 0x3f, 0xc9, 0x73, 0x68, 0x4f, 0xe9, 0xa4, 0x46, 0x87, 0x7c,
	0x3b, 0xb1, 0x8f, 0x22, 0x69, 0x1e, 0x45, 0x72, 0xca, 0x66, 0xa9, 0xb5, 0x9c, 0xf8, 0xc7, 0x5e,
	0x7c, 0x0e, 0xe1, 0x95, 0x36, 0xff, 0xbe, 0xe1, 0xde, 0x77, 0x1f, 0xee, 0xcc, 0xe7, 0x90, 0x15,
	0x67, 0x12, 0xaf, 0xbd, 0xa4, 0x57, 0xcb, 0xd4, 0x56, 0xa7, 0xb7, 0xa9, 0x35, 0xd8, 0x5e, 0xaf,
	0x60, 0xdb, 0xbd, 0x21, 0xf8, 0x1f, 0x72, 0x1b, 0xfe, 0xf4, 0x00, 0xdc, 0x20, 0xa7, 0xe3, 0xb7,
	0xe4, 0x04, 0xba, 0xa3, 0x6f, 0x98, 0xd5, 0x0a, 0xc9, 0xf6, 0xba, 0x9f, 0x23, 0xde, 0x59, 0x3b,
	0xfb, 0xc0, 0x3b, 0xf2, 0xc8, 0x3e, 0x74, 0x52, 0xcc, 0xb8, 0xc8, 0xc9, 0xe6, 0x92, 0x29, 0x5e,
	0x5e, 0x92, 0x7d, 0x68, 0xbd, 0x2b, 0xa4, 0x6a, 0x5c, 0xe7, 0x95, 0xfe, 0xa5, 0x57, 0x5d, 0x47,
	0x1e, 0x39, 0x80, 0xee, 0x87, 0x3a, 0x33, 0x91, 0xbf, 0xb0, 0xbe, 0xa1, 0xc5, 0xa4, 0x16, 0xf8,
	0x27, 0xeb, 0xd9, 0x63, 0xb8, 0x97, 0xf1, 0x32, 0xa9, 0x3e, 0x17, 0xb2, 0x2e, 0x14, 0x1a, 0xf9,
	0xac, 0xe7, 0xf4, 0xb1, 0xa6, 0x34, 0xf6, 0x2e, 0x3a, 0x06, 0xd7, 0x8b, 0xdf, 0x01, 0x00, 0x00,
	0xff, 0xff, 0x8e, 0x37, 0x0d, 0x85, 0xf1, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProcessAPIClient is the client API for ProcessAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessAPIClient interface {
	Execute(ctx context.Context, opts ...grpc.CallOption) (ProcessAPI_ExecuteClient, error)
	Record(ctx context.Context, in *Process, opts ...grpc.CallOption) (*Process, error)
	List(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessAPI_ListClient, error)
	Success(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessAPI_SuccessClient, error)
	Failure(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessAPI_FailureClient, error)
}

type processAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessAPIClient(cc grpc.ClientConnInterface) ProcessAPIClient {
	return &processAPIClient{cc}
}

func (c *processAPIClient) Execute(ctx context.Context, opts ...grpc.CallOption) (ProcessAPI_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessAPI_serviceDesc.Streams[0], "/data.ProcessAPI/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &processAPIExecuteClient{stream}
	return x, nil
}

type ProcessAPI_ExecuteClient interface {
	Send(*ProcessRequest) error
	Recv() (*ProcessResponse, error)
	grpc.ClientStream
}

type processAPIExecuteClient struct {
	grpc.ClientStream
}

func (x *processAPIExecuteClient) Send(m *ProcessRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *processAPIExecuteClient) Recv() (*ProcessResponse, error) {
	m := new(ProcessResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processAPIClient) Record(ctx context.Context, in *Process, opts ...grpc.CallOption) (*Process, error) {
	out := new(Process)
	err := c.cc.Invoke(ctx, "/data.ProcessAPI/Record", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *processAPIClient) List(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessAPI_ListClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessAPI_serviceDesc.Streams[1], "/data.ProcessAPI/List", opts...)
	if err != nil {
		return nil, err
	}
	x := &processAPIListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessAPI_ListClient interface {
	Recv() (*Process, error)
	grpc.ClientStream
}

type processAPIListClient struct {
	grpc.ClientStream
}

func (x *processAPIListClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processAPIClient) Success(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessAPI_SuccessClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessAPI_serviceDesc.Streams[2], "/data.ProcessAPI/Success", opts...)
	if err != nil {
		return nil, err
	}
	x := &processAPISuccessClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessAPI_SuccessClient interface {
	Recv() (*Process, error)
	grpc.ClientStream
}

type processAPISuccessClient struct {
	grpc.ClientStream
}

func (x *processAPISuccessClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *processAPIClient) Failure(ctx context.Context, in *Options, opts ...grpc.CallOption) (ProcessAPI_FailureClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessAPI_serviceDesc.Streams[3], "/data.ProcessAPI/Failure", opts...)
	if err != nil {
		return nil, err
	}
	x := &processAPIFailureClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessAPI_FailureClient interface {
	Recv() (*Process, error)
	grpc.ClientStream
}

type processAPIFailureClient struct {
	grpc.ClientStream
}

func (x *processAPIFailureClient) Recv() (*Process, error) {
	m := new(Process)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessAPIServer is the server API for ProcessAPI service.
type ProcessAPIServer interface {
	Execute(ProcessAPI_ExecuteServer) error
	Record(context.Context, *Process) (*Process, error)
	List(*Options, ProcessAPI_ListServer) error
	Success(*Options, ProcessAPI_SuccessServer) error
	Failure(*Options, ProcessAPI_FailureServer) error
}

// UnimplementedProcessAPIServer can be embedded to have forward compatible implementations.
type UnimplementedProcessAPIServer struct {
}

func (*UnimplementedProcessAPIServer) Execute(srv ProcessAPI_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (*UnimplementedProcessAPIServer) Record(ctx context.Context, req *Process) (*Process, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Record not implemented")
}
func (*UnimplementedProcessAPIServer) List(req *Options, srv ProcessAPI_ListServer) error {
	return status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (*UnimplementedProcessAPIServer) Success(req *Options, srv ProcessAPI_SuccessServer) error {
	return status.Errorf(codes.Unimplemented, "method Success not implemented")
}
func (*UnimplementedProcessAPIServer) Failure(req *Options, srv ProcessAPI_FailureServer) error {
	return status.Errorf(codes.Unimplemented, "method Failure not implemented")
}

func RegisterProcessAPIServer(s *grpc.Server, srv ProcessAPIServer) {
	s.RegisterService(&_ProcessAPI_serviceDesc, srv)
}

func _ProcessAPI_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProcessAPIServer).Execute(&processAPIExecuteServer{stream})
}

type ProcessAPI_ExecuteServer interface {
	Send(*ProcessResponse) error
	Recv() (*ProcessRequest, error)
	grpc.ServerStream
}

type processAPIExecuteServer struct {
	grpc.ServerStream
}

func (x *processAPIExecuteServer) Send(m *ProcessResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *processAPIExecuteServer) Recv() (*ProcessRequest, error) {
	m := new(ProcessRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProcessAPI_Record_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Process)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessAPIServer).Record(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/data.ProcessAPI/Record",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessAPIServer).Record(ctx, req.(*Process))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProcessAPI_List_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessAPIServer).List(m, &processAPIListServer{stream})
}

type ProcessAPI_ListServer interface {
	Send(*Process) error
	grpc.ServerStream
}

type processAPIListServer struct {
	grpc.ServerStream
}

func (x *processAPIListServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessAPI_Success_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessAPIServer).Success(m, &processAPISuccessServer{stream})
}

type ProcessAPI_SuccessServer interface {
	Send(*Process) error
	grpc.ServerStream
}

type processAPISuccessServer struct {
	grpc.ServerStream
}

func (x *processAPISuccessServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

func _ProcessAPI_Failure_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Options)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessAPIServer).Failure(m, &processAPIFailureServer{stream})
}

type ProcessAPI_FailureServer interface {
	Send(*Process) error
	grpc.ServerStream
}

type processAPIFailureServer struct {
	grpc.ServerStream
}

func (x *processAPIFailureServer) Send(m *Process) error {
	return x.ServerStream.SendMsg(m)
}

var _ProcessAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "data.ProcessAPI",
	HandlerType: (*ProcessAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Record",
			Handler:    _ProcessAPI_Record_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _ProcessAPI_Execute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "List",
			Handler:       _ProcessAPI_List_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Success",
			Handler:       _ProcessAPI_Success_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Failure",
			Handler:       _ProcessAPI_Failure_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "process.proto",
}