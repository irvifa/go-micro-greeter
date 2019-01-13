// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/saying/saying.proto

package go_micro_srv_greeter

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

type Request struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_saying_d9a1fc0f38f46db0, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_saying_d9a1fc0f38f46db0, []int{1}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.greeter.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.greeter.Response")
}

func init() { proto.RegisterFile("proto/saying/saying.proto", fileDescriptor_saying_d9a1fc0f38f46db0) }

var fileDescriptor_saying_d9a1fc0f38f46db0 = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x4e, 0xac, 0xcc, 0xcc, 0x4b, 0x87, 0x52, 0x7a, 0x60, 0x31, 0x21, 0x91, 0xf4,
	0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xe2, 0xa2, 0x32, 0xbd, 0xf4, 0xa2, 0xd4, 0xd4,
	0x92, 0xd4, 0x22, 0x25, 0x59, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21,
	0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49,
	0x86, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x80, 0x8b, 0x39, 0xb7,
	0x38, 0x1d, 0x2a, 0x0d, 0x62, 0x1a, 0x05, 0x71, 0xb1, 0x05, 0x83, 0xad, 0x10, 0xf2, 0xe0, 0x62,
	0xf5, 0x48, 0xcd, 0xc9, 0xc9, 0x17, 0x92, 0xd5, 0xc3, 0x66, 0x8d, 0x1e, 0xd4, 0x0e, 0x29, 0x39,
	0x5c, 0xd2, 0x10, 0x3b, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0xae, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x1d, 0x5c, 0x14, 0x4d, 0xca, 0x00, 0x00, 0x00,
}
