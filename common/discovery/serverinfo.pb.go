// Code generated by protoc-gen-go. DO NOT EDIT.
// source: serverinfo.proto

/*
Package discovery is a generated protocol buffer package.

It is generated from these files:
	serverinfo.proto

It has these top-level messages:
	ServerInfo
*/
package discovery

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

type ServerInfo struct {
	Type       uint32 `protobuf:"varint,1,opt,name=type" json:"type,omitempty"`
	Id         string `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	ExternalIp string `protobuf:"bytes,3,opt,name=external_ip,json=externalIp" json:"external_ip,omitempty"`
	IntranetIp string `protobuf:"bytes,4,opt,name=intranet_ip,json=intranetIp" json:"intranet_ip,omitempty"`
	Ordered    uint32 `protobuf:"varint,5,opt,name=ordered" json:"ordered,omitempty"`
}

func (m *ServerInfo) Reset()                    { *m = ServerInfo{} }
func (m *ServerInfo) String() string            { return proto.CompactTextString(m) }
func (*ServerInfo) ProtoMessage()               {}
func (*ServerInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ServerInfo) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *ServerInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ServerInfo) GetExternalIp() string {
	if m != nil {
		return m.ExternalIp
	}
	return ""
}

func (m *ServerInfo) GetIntranetIp() string {
	if m != nil {
		return m.IntranetIp
	}
	return ""
}

func (m *ServerInfo) GetOrdered() uint32 {
	if m != nil {
		return m.Ordered
	}
	return 0
}

func init() {
	proto.RegisterType((*ServerInfo)(nil), "discovery.ServerInfo")
}

func init() { proto.RegisterFile("serverinfo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8e, 0xc1, 0xca, 0xc2, 0x30,
	0x0c, 0x80, 0xe9, 0xfe, 0xfd, 0xca, 0x22, 0x8a, 0xe4, 0xd4, 0x9b, 0xc3, 0xd3, 0x4e, 0x5e, 0x7c,
	0x8a, 0x5e, 0xe7, 0x03, 0xc8, 0xb4, 0x19, 0x04, 0xa4, 0x29, 0x59, 0x19, 0xee, 0x1d, 0x7c, 0x68,
	0x59, 0xa5, 0xb7, 0xe4, 0xfb, 0x3e, 0x48, 0xe0, 0x38, 0x91, 0xce, 0xa4, 0x1c, 0x46, 0xb9, 0x44,
	0x95, 0x24, 0xd8, 0x78, 0x9e, 0x9e, 0x32, 0x93, 0x2e, 0xe7, 0x8f, 0x01, 0xb8, 0x65, 0xef, 0xc2,
	0x28, 0x88, 0x50, 0xa7, 0x25, 0x92, 0x35, 0xad, 0xe9, 0xf6, 0x7d, 0x9e, 0xf1, 0x00, 0x15, 0x7b,
	0x5b, 0xb5, 0xa6, 0x6b, 0xfa, 0x8a, 0x3d, 0x9e, 0x60, 0x47, 0xef, 0x44, 0x1a, 0x86, 0xd7, 0x9d,
	0xa3, 0xfd, 0xcb, 0x02, 0x0a, 0x72, 0x71, 0x0d, 0x38, 0x24, 0x1d, 0x02, 0xa5, 0x35, 0xa8, 0x7f,
	0x41, 0x41, 0x2e, 0xa2, 0x85, 0xad, 0xa8, 0x27, 0x25, 0x6f, 0xff, 0xf3, 0xa1, 0xb2, 0x3e, 0x36,
	0xf9, 0xc1, 0xeb, 0x37, 0x00, 0x00, 0xff, 0xff, 0xef, 0xb2, 0x43, 0x2b, 0xb4, 0x00, 0x00, 0x00,
}