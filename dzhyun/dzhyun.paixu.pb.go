// Code generated by protoc-gen-go.
// source: dzhyun.paixu.proto
// DO NOT EDIT!

package dzhyun

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 排序存储结构
type PaiXu struct {
	Obj   string `protobuf:"bytes,1,opt,name=Obj" json:"Obj,omitempty"`
	Value YFloat  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
	Text  string `protobuf:"bytes,3,opt,name=Text" json:"Text,omitempty"`
}

func (m *PaiXu) Reset()                    { *m = PaiXu{} }
func (m *PaiXu) String() string            { return proto.CompactTextString(m) }
func (*PaiXu) ProtoMessage()               {}
func (*PaiXu) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{0} }

func (m *PaiXu) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *PaiXu) GetValue() YFloat {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PaiXu) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

// 排名存储结构
type PaiMing struct {
	Obj    string `protobuf:"bytes,1,opt,name=Obj" json:"Obj,omitempty"`
	Value  YFloat  `protobuf:"varint,2,opt,name=Value" json:"Value,omitempty"`
	Text   string `protobuf:"bytes,3,opt,name=Text" json:"Text,omitempty"`
	MingCi YFloat  `protobuf:"varint,4,opt,name=MingCi" json:"MingCi,omitempty"`
}

func (m *PaiMing) Reset()                    { *m = PaiMing{} }
func (m *PaiMing) String() string            { return proto.CompactTextString(m) }
func (*PaiMing) ProtoMessage()               {}
func (*PaiMing) Descriptor() ([]byte, []int) { return fileDescriptor10, []int{1} }

func (m *PaiMing) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *PaiMing) GetValue() YFloat {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *PaiMing) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *PaiMing) GetMingCi() YFloat {
	if m != nil {
		return m.MingCi
	}
	return 0
}

func init() {
	proto.RegisterType((*PaiXu)(nil), "dzhyun.PaiXu")
	proto.RegisterType((*PaiMing)(nil), "dzhyun.PaiMing")
}

func init() { proto.RegisterFile("dzhyun.paixu.proto", fileDescriptor10) }

var fileDescriptor10 = []byte{
	// 149 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xa9, 0xca, 0xa8,
	0x2c, 0xcd, 0xd3, 0x2b, 0x48, 0xcc, 0xac, 0x28, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x83, 0x88, 0x29, 0x39, 0x73, 0xb1, 0x06, 0x24, 0x66, 0x46, 0x94, 0x0a, 0x09, 0x70, 0x31, 0xfb,
	0x27, 0x65, 0x49, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0x22, 0x5c, 0xac, 0x61,
	0x89, 0x39, 0xa5, 0xa9, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x10, 0x8e, 0x90, 0x10, 0x17,
	0x4b, 0x48, 0x6a, 0x45, 0x89, 0x04, 0x33, 0x58, 0x21, 0x98, 0xad, 0x14, 0xcb, 0xc5, 0x1e, 0x90,
	0x98, 0xe9, 0x9b, 0x99, 0x97, 0x4e, 0x89, 0x31, 0x42, 0x62, 0x5c, 0x6c, 0x20, 0x33, 0x9c, 0x33,
	0x25, 0x58, 0xc0, 0x4a, 0xa1, 0x3c, 0x27, 0x21, 0x2e, 0x81, 0xe4, 0xfc, 0x5c, 0x3d, 0x98, 0x2f,
	0x40, 0xee, 0x4f, 0x62, 0x03, 0x53, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xeb, 0x63, 0xa8,
	0x27, 0xdc, 0x00, 0x00, 0x00,
}
