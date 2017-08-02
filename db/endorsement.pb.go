// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db/endorsement.proto

package db

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

type Endorsement struct {
	Emitter   string `protobuf:"bytes,1,opt,name=emitter" json:"emitter,omitempty"`
	Uuid      string `protobuf:"bytes,2,opt,name=uuid" json:"uuid,omitempty"`
	Signature []byte `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Endorsement) Reset()                    { *m = Endorsement{} }
func (m *Endorsement) String() string            { return proto.CompactTextString(m) }
func (*Endorsement) ProtoMessage()               {}
func (*Endorsement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Endorsement) GetEmitter() string {
	if m != nil {
		return m.Emitter
	}
	return ""
}

func (m *Endorsement) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Endorsement) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func init() {
	proto.RegisterType((*Endorsement)(nil), "db.Endorsement")
}

func init() { proto.RegisterFile("db/endorsement.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x49, 0xd2, 0x4f,
	0xcd, 0x4b, 0xc9, 0x2f, 0x2a, 0x4e, 0xcd, 0x4d, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x4a, 0x49, 0x52, 0x8a, 0xe4, 0xe2, 0x76, 0x45, 0x48, 0x08, 0x49, 0x70, 0xb1, 0xa7,
	0xe6, 0x66, 0x96, 0x94, 0xa4, 0x16, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42,
	0x42, 0x5c, 0x2c, 0xa5, 0xa5, 0x99, 0x29, 0x12, 0x4c, 0x60, 0x61, 0x30, 0x5b, 0x48, 0x86, 0x8b,
	0xb3, 0x38, 0x33, 0x3d, 0x2f, 0xb1, 0xa4, 0xb4, 0x28, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0x27,
	0x08, 0x21, 0x90, 0xc4, 0x06, 0xb6, 0xc5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x19, 0x64, 0x32,
	0x49, 0x7d, 0x00, 0x00, 0x00,
}
