// Code generated by protoc-gen-go. DO NOT EDIT.
// source: db/spore.proto

package db

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"
import version "gitlab.com/SporeDB/sporedb/db/version"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Operation_Op int32

const (
	// Operations on every values
	Operation_SET    Operation_Op = 0
	Operation_CONCAT Operation_Op = 1
	// Operations on numeric values
	Operation_ADD Operation_Op = 10
	Operation_MUL Operation_Op = 11
	// Operations on set values
	Operation_SADD Operation_Op = 20
	Operation_SREM Operation_Op = 21
)

var Operation_Op_name = map[int32]string{
	0:  "SET",
	1:  "CONCAT",
	10: "ADD",
	11: "MUL",
	20: "SADD",
	21: "SREM",
}
var Operation_Op_value = map[string]int32{
	"SET":    0,
	"CONCAT": 1,
	"ADD":    10,
	"MUL":    11,
	"SADD":   20,
	"SREM":   21,
}

func (x Operation_Op) String() string {
	return proto.EnumName(Operation_Op_name, int32(x))
}
func (Operation_Op) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

type Spore struct {
	Uuid         string                     `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Policy       string                     `protobuf:"bytes,2,opt,name=policy" json:"policy,omitempty"`
	Emitter      string                     `protobuf:"bytes,3,opt,name=emitter" json:"emitter,omitempty"`
	Deadline     *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=deadline" json:"deadline,omitempty"`
	Requirements map[string]*version.V      `protobuf:"bytes,5,rep,name=requirements" json:"requirements,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Operations   []*Operation               `protobuf:"bytes,6,rep,name=operations" json:"operations,omitempty"`
	// The signature of the spore is computed on the spore with an empty signature.
	// signature = signature by emitter ( hash ( marshal ( spore without signature ) ) )
	//
	// It is used to check spore's integrity.
	Signature []byte `protobuf:"bytes,10,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Spore) Reset()                    { *m = Spore{} }
func (m *Spore) String() string            { return proto.CompactTextString(m) }
func (*Spore) ProtoMessage()               {}
func (*Spore) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Spore) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *Spore) GetPolicy() string {
	if m != nil {
		return m.Policy
	}
	return ""
}

func (m *Spore) GetEmitter() string {
	if m != nil {
		return m.Emitter
	}
	return ""
}

func (m *Spore) GetDeadline() *google_protobuf.Timestamp {
	if m != nil {
		return m.Deadline
	}
	return nil
}

func (m *Spore) GetRequirements() map[string]*version.V {
	if m != nil {
		return m.Requirements
	}
	return nil
}

func (m *Spore) GetOperations() []*Operation {
	if m != nil {
		return m.Operations
	}
	return nil
}

func (m *Spore) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Operation struct {
	Key      string       `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Op       Operation_Op `protobuf:"varint,2,opt,name=op,enum=db.Operation_Op" json:"op,omitempty"`
	Data     []byte       `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Metadata []byte       `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (m *Operation) Reset()                    { *m = Operation{} }
func (m *Operation) String() string            { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()               {}
func (*Operation) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Operation) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Operation) GetOp() Operation_Op {
	if m != nil {
		return m.Op
	}
	return Operation_SET
}

func (m *Operation) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Operation) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type RecoverRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *RecoverRequest) Reset()                    { *m = RecoverRequest{} }
func (m *RecoverRequest) String() string            { return proto.CompactTextString(m) }
func (*RecoverRequest) ProtoMessage()               {}
func (*RecoverRequest) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{2} }

func (m *RecoverRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func init() {
	proto.RegisterType((*Spore)(nil), "db.Spore")
	proto.RegisterType((*Operation)(nil), "db.Operation")
	proto.RegisterType((*RecoverRequest)(nil), "db.RecoverRequest")
	proto.RegisterEnum("db.Operation_Op", Operation_Op_name, Operation_Op_value)
}

func init() { proto.RegisterFile("db/spore.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 408 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x8e, 0xd3, 0x30,
	0x14, 0x85, 0x49, 0xfa, 0x33, 0xed, 0x6d, 0xa9, 0x82, 0x05, 0xc8, 0x0a, 0x48, 0x44, 0x59, 0x75,
	0x83, 0x2b, 0x15, 0x09, 0x21, 0x36, 0xa3, 0xd1, 0x4c, 0x57, 0x30, 0x54, 0x72, 0x0b, 0x7b, 0x67,
	0x7c, 0xa9, 0x2c, 0x9a, 0xd8, 0x38, 0x4e, 0xa5, 0xbe, 0x15, 0x8f, 0xc4, 0xa3, 0x20, 0x3b, 0x93,
	0x32, 0xa3, 0xce, 0x2a, 0xe7, 0x7e, 0xf7, 0x27, 0x47, 0xc7, 0x30, 0x93, 0xc5, 0xa2, 0x36, 0xda,
	0x22, 0x33, 0x56, 0x3b, 0x4d, 0x62, 0x59, 0xa4, 0xef, 0x76, 0x5a, 0xef, 0xf6, 0xb8, 0x08, 0xa4,
	0x68, 0x7e, 0x2e, 0x9c, 0x2a, 0xb1, 0x76, 0xa2, 0x34, 0xed, 0x50, 0x4a, 0x65, 0xb1, 0x38, 0xa0,
	0xad, 0x95, 0xae, 0xba, 0x6f, 0xdb, 0xc9, 0xff, 0xc6, 0x30, 0xd8, 0xf8, 0x73, 0x84, 0x40, 0xbf,
	0x69, 0x94, 0xa4, 0x51, 0x16, 0xcd, 0xc7, 0x3c, 0x68, 0xf2, 0x1a, 0x86, 0x46, 0xef, 0xd5, 0xdd,
	0x91, 0xc6, 0x81, 0xde, 0x57, 0x84, 0xc2, 0x05, 0x96, 0xca, 0x39, 0xb4, 0xb4, 0x17, 0x1a, 0x5d,
	0x49, 0x3e, 0xc2, 0x48, 0xa2, 0x90, 0x7b, 0x55, 0x21, 0xed, 0x67, 0xd1, 0x7c, 0xb2, 0x4c, 0x59,
	0xeb, 0x8e, 0x75, 0xee, 0xd8, 0xb6, 0x73, 0xc7, 0x4f, 0xb3, 0xe4, 0x12, 0xa6, 0x16, 0x7f, 0x37,
	0xca, 0x62, 0x89, 0x95, 0xab, 0xe9, 0x20, 0xeb, 0xcd, 0x27, 0xcb, 0x37, 0x4c, 0x16, 0x2c, 0xd8,
	0x63, 0xfc, 0x41, 0x77, 0x55, 0x39, 0x7b, 0xe4, 0x8f, 0x16, 0xc8, 0x7b, 0x00, 0x6d, 0xd0, 0x0a,
	0xa7, 0x74, 0x55, 0xd3, 0x61, 0x58, 0x7f, 0xee, 0xd7, 0xd7, 0x1d, 0xe5, 0x0f, 0x06, 0xc8, 0x5b,
	0x18, 0xd7, 0x6a, 0x57, 0x09, 0xd7, 0x58, 0xa4, 0x90, 0x45, 0xf3, 0x29, 0xff, 0x0f, 0xd2, 0x2f,
	0xf0, 0xe2, 0xec, 0x7f, 0x24, 0x81, 0xde, 0x2f, 0x3c, 0xde, 0xe7, 0xe3, 0x25, 0xc9, 0x60, 0x70,
	0x10, 0xfb, 0x06, 0x43, 0x3a, 0x93, 0x25, 0xb0, 0x2e, 0xdb, 0x1f, 0xbc, 0x6d, 0x7c, 0x8e, 0x3f,
	0x45, 0xf9, 0x9f, 0x08, 0xc6, 0x27, 0x13, 0x4f, 0x5e, 0x89, 0xb5, 0x09, 0x27, 0x66, 0xcb, 0xe4,
	0x91, 0x63, 0xb6, 0x36, 0x3c, 0xd6, 0xc6, 0x3f, 0x8d, 0x14, 0x4e, 0x84, 0xac, 0xa7, 0x3c, 0x68,
	0x92, 0xc2, 0xa8, 0x44, 0x27, 0x02, 0xef, 0x07, 0x7e, 0xaa, 0xf3, 0x4b, 0x88, 0xd7, 0x86, 0x5c,
	0x40, 0x6f, 0xb3, 0xda, 0x26, 0xcf, 0x08, 0xc0, 0xf0, 0x7a, 0xfd, 0xed, 0xfa, 0x6a, 0x9b, 0x44,
	0x1e, 0x5e, 0xdd, 0xdc, 0x24, 0xe0, 0xc5, 0xed, 0xf7, 0xaf, 0xc9, 0x84, 0x8c, 0xa0, 0xbf, 0xf1,
	0xe8, 0x65, 0x50, 0x7c, 0x75, 0x9b, 0xbc, 0xca, 0x73, 0x98, 0x71, 0xbc, 0xd3, 0x07, 0xb4, 0x3e,
	0x06, 0xac, 0xdd, 0xb9, 0xed, 0x62, 0x18, 0xde, 0xf3, 0xc3, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x87, 0xc4, 0x90, 0xfa, 0x91, 0x02, 0x00, 0x00,
}
