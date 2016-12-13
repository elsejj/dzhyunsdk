// Code generated by protoc-gen-go.
// source: dzhyun.dxspirit.proto
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
type DXSpirit struct {
	ShiJian int64  `protobuf:"varint,1,opt,name=ShiJian" json:"ShiJian,omitempty"`
	Obj     string `protobuf:"bytes,2,opt,name=Obj" json:"Obj,omitempty"`
	TongZhi string `protobuf:"bytes,3,opt,name=TongZhi" json:"TongZhi,omitempty"`
	ShuJu   int64  `protobuf:"varint,4,opt,name=ShuJu" json:"ShuJu,omitempty"`
}

func (m *DXSpirit) Reset()                    { *m = DXSpirit{} }
func (m *DXSpirit) String() string            { return proto.CompactTextString(m) }
func (*DXSpirit) ProtoMessage()               {}
func (*DXSpirit) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *DXSpirit) GetShiJian() int64 {
	if m != nil {
		return m.ShiJian
	}
	return 0
}

func (m *DXSpirit) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *DXSpirit) GetTongZhi() string {
	if m != nil {
		return m.TongZhi
	}
	return ""
}

func (m *DXSpirit) GetShuJu() int64 {
	if m != nil {
		return m.ShuJu
	}
	return 0
}

type DXSpiritStat struct {
	Obj            string `protobuf:"bytes,1,opt,name=Obj" json:"Obj,omitempty"`
	HjfsTotal      int64  `protobuf:"varint,2,opt,name=HjfsTotal" json:"HjfsTotal,omitempty"`
	KsftTotal      int64  `protobuf:"varint,3,opt,name=KsftTotal" json:"KsftTotal,omitempty"`
	GttsTotal      int64  `protobuf:"varint,4,opt,name=GttsTotal" json:"GttsTotal,omitempty"`
	JsxdTotal      int64  `protobuf:"varint,5,opt,name=JsxdTotal" json:"JsxdTotal,omitempty"`
	DbmrTotal      int64  `protobuf:"varint,6,opt,name=DbmrTotal" json:"DbmrTotal,omitempty"`
	DbmrStatistics int64  `protobuf:"varint,7,opt,name=DbmrStatistics" json:"DbmrStatistics,omitempty"`
	DbmcTotal      int64  `protobuf:"varint,8,opt,name=DbmcTotal" json:"DbmcTotal,omitempty"`
	DbmcStatistics int64  `protobuf:"varint,9,opt,name=DbmcStatistics" json:"DbmcStatistics,omitempty"`
	FztbTotal      int64  `protobuf:"varint,10,opt,name=FztbTotal" json:"FztbTotal,omitempty"`
	FdtbTotal      int64  `protobuf:"varint,11,opt,name=FdtbTotal" json:"FdtbTotal,omitempty"`
	DkztTotal      int64  `protobuf:"varint,12,opt,name=DkztTotal" json:"DkztTotal,omitempty"`
	DkdtTotal      int64  `protobuf:"varint,13,opt,name=DkdtTotal" json:"DkdtTotal,omitempty"`
	YdmcPTotal     int64  `protobuf:"varint,14,opt,name=YdmcPTotal" json:"YdmcPTotal,omitempty"`
	YdmrPTotal     int64  `protobuf:"varint,15,opt,name=YdmrPTotal" json:"YdmrPTotal,omitempty"`
	LszsTotal      int64  `protobuf:"varint,16,opt,name=LszsTotal" json:"LszsTotal,omitempty"`
	DyzsTotal      int64  `protobuf:"varint,17,opt,name=DyzsTotal" json:"DyzsTotal,omitempty"`
	JgmrgdTotal    int64  `protobuf:"varint,18,opt,name=JgmrgdTotal" json:"JgmrgdTotal,omitempty"`
	JgmcgdTotal    int64  `protobuf:"varint,19,opt,name=JgmcgdTotal" json:"JgmcgdTotal,omitempty"`
	DcjmrdTotal    int64  `protobuf:"varint,20,opt,name=DcjmrdTotal" json:"DcjmrdTotal,omitempty"`
	DcjmcdTotal    int64  `protobuf:"varint,21,opt,name=DcjmcdTotal" json:"DcjmcdTotal,omitempty"`
	FdmrgdTotal    int64  `protobuf:"varint,22,opt,name=FdmrgdTotal" json:"FdmrgdTotal,omitempty"`
	FdmcgdTotal    int64  `protobuf:"varint,23,opt,name=FdmcgdTotal" json:"FdmcgdTotal,omitempty"`
	MrcdTotal      int64  `protobuf:"varint,24,opt,name=MrcdTotal" json:"MrcdTotal,omitempty"`
	MccdTotal      int64  `protobuf:"varint,25,opt,name=MccdTotal" json:"MccdTotal,omitempty"`
	MrxdTotal      int64  `protobuf:"varint,26,opt,name=MrxdTotal" json:"MrxdTotal,omitempty"`
	McxdTotal      int64  `protobuf:"varint,27,opt,name=McxdTotal" json:"McxdTotal,omitempty"`
}

func (m *DXSpiritStat) Reset()                    { *m = DXSpiritStat{} }
func (m *DXSpiritStat) String() string            { return proto.CompactTextString(m) }
func (*DXSpiritStat) ProtoMessage()               {}
func (*DXSpiritStat) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *DXSpiritStat) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *DXSpiritStat) GetHjfsTotal() int64 {
	if m != nil {
		return m.HjfsTotal
	}
	return 0
}

func (m *DXSpiritStat) GetKsftTotal() int64 {
	if m != nil {
		return m.KsftTotal
	}
	return 0
}

func (m *DXSpiritStat) GetGttsTotal() int64 {
	if m != nil {
		return m.GttsTotal
	}
	return 0
}

func (m *DXSpiritStat) GetJsxdTotal() int64 {
	if m != nil {
		return m.JsxdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetDbmrTotal() int64 {
	if m != nil {
		return m.DbmrTotal
	}
	return 0
}

func (m *DXSpiritStat) GetDbmrStatistics() int64 {
	if m != nil {
		return m.DbmrStatistics
	}
	return 0
}

func (m *DXSpiritStat) GetDbmcTotal() int64 {
	if m != nil {
		return m.DbmcTotal
	}
	return 0
}

func (m *DXSpiritStat) GetDbmcStatistics() int64 {
	if m != nil {
		return m.DbmcStatistics
	}
	return 0
}

func (m *DXSpiritStat) GetFztbTotal() int64 {
	if m != nil {
		return m.FztbTotal
	}
	return 0
}

func (m *DXSpiritStat) GetFdtbTotal() int64 {
	if m != nil {
		return m.FdtbTotal
	}
	return 0
}

func (m *DXSpiritStat) GetDkztTotal() int64 {
	if m != nil {
		return m.DkztTotal
	}
	return 0
}

func (m *DXSpiritStat) GetDkdtTotal() int64 {
	if m != nil {
		return m.DkdtTotal
	}
	return 0
}

func (m *DXSpiritStat) GetYdmcPTotal() int64 {
	if m != nil {
		return m.YdmcPTotal
	}
	return 0
}

func (m *DXSpiritStat) GetYdmrPTotal() int64 {
	if m != nil {
		return m.YdmrPTotal
	}
	return 0
}

func (m *DXSpiritStat) GetLszsTotal() int64 {
	if m != nil {
		return m.LszsTotal
	}
	return 0
}

func (m *DXSpiritStat) GetDyzsTotal() int64 {
	if m != nil {
		return m.DyzsTotal
	}
	return 0
}

func (m *DXSpiritStat) GetJgmrgdTotal() int64 {
	if m != nil {
		return m.JgmrgdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetJgmcgdTotal() int64 {
	if m != nil {
		return m.JgmcgdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetDcjmrdTotal() int64 {
	if m != nil {
		return m.DcjmrdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetDcjmcdTotal() int64 {
	if m != nil {
		return m.DcjmcdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetFdmrgdTotal() int64 {
	if m != nil {
		return m.FdmrgdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetFdmcgdTotal() int64 {
	if m != nil {
		return m.FdmcgdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetMrcdTotal() int64 {
	if m != nil {
		return m.MrcdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetMccdTotal() int64 {
	if m != nil {
		return m.MccdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetMrxdTotal() int64 {
	if m != nil {
		return m.MrxdTotal
	}
	return 0
}

func (m *DXSpiritStat) GetMcxdTotal() int64 {
	if m != nil {
		return m.McxdTotal
	}
	return 0
}

func init() {
	proto.RegisterType((*DXSpirit)(nil), "dzhyun.DXSpirit")
	proto.RegisterType((*DXSpiritStat)(nil), "dzhyun.DXSpiritStat")
}

func init() { proto.RegisterFile("dzhyun.dxspirit.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x5c, 0x93, 0xd1, 0x8e, 0xd2, 0x40,
	0x14, 0x86, 0x53, 0xeb, 0x76, 0xb7, 0xb3, 0x88, 0x38, 0x82, 0x56, 0x25, 0x86, 0x70, 0x61, 0xb8,
	0xe2, 0xc6, 0x37, 0x30, 0x0d, 0x9a, 0xaa, 0xd1, 0x50, 0x2e, 0xd4, 0x3b, 0x3a, 0x13, 0x68, 0xd1,
	0x52, 0x32, 0x1d, 0x12, 0xe8, 0xb3, 0xfa, 0x30, 0x66, 0xe6, 0x9c, 0x33, 0x1d, 0xf7, 0x6a, 0x77,
	0xbe, 0xef, 0x9f, 0xff, 0x1c, 0x0a, 0x65, 0x13, 0xd9, 0x95, 0xd7, 0xf3, 0x71, 0x29, 0x2f, 0xed,
	0xa9, 0x52, 0x95, 0x5e, 0x9e, 0x54, 0xa3, 0x1b, 0x1e, 0x01, 0x9e, 0xef, 0xd8, 0x5d, 0xfa, 0x23,
	0xb7, 0x86, 0x27, 0xec, 0x36, 0x2f, 0xab, 0xac, 0xda, 0x1e, 0x93, 0x60, 0x16, 0x2c, 0xc2, 0x35,
	0x1d, 0xf9, 0x88, 0x85, 0xdf, 0x8a, 0x43, 0xf2, 0x68, 0x16, 0x2c, 0xe2, 0xb5, 0xf9, 0xd7, 0x64,
	0x37, 0xcd, 0x71, 0xff, 0xab, 0xac, 0x92, 0xd0, 0x52, 0x3a, 0xf2, 0x31, 0xbb, 0xc9, 0xcb, 0x73,
	0x76, 0x4e, 0x1e, 0xdb, 0x0e, 0x38, 0xcc, 0xff, 0x46, 0x6c, 0x40, 0x83, 0x72, 0xbd, 0xd5, 0x54,
	0x19, 0xf4, 0x95, 0x53, 0x16, 0x7f, 0x3a, 0xec, 0xda, 0x4d, 0xa3, 0xb7, 0x7f, 0xec, 0xa8, 0x70,
	0xdd, 0x03, 0x63, 0x3f, 0xb7, 0x3b, 0x0d, 0x36, 0x04, 0xeb, 0x80, 0xb1, 0x1f, 0xb5, 0xc6, 0xbb,
	0x30, 0xb8, 0x07, 0xc6, 0x66, 0xed, 0x45, 0x82, 0xbd, 0x01, 0xeb, 0x80, 0xb1, 0x69, 0x51, 0x2b,
	0xb0, 0x11, 0x58, 0x07, 0xf8, 0x3b, 0x36, 0x34, 0x07, 0xb3, 0x73, 0xd5, 0xea, 0x4a, 0xb4, 0xc9,
	0xad, 0x8d, 0x3c, 0xa0, 0xd8, 0x22, 0xa0, 0xe5, 0xce, 0xb5, 0x08, 0xbf, 0x45, 0x78, 0x2d, 0xb1,
	0x6b, 0x11, 0xff, 0xb7, 0xac, 0x3a, 0x5d, 0x40, 0x0b, 0x83, 0x16, 0x07, 0xac, 0x95, 0x64, 0xef,
	0xd1, 0x4a, 0xcf, 0xa6, 0xbf, 0x3b, 0x7c, 0x42, 0x03, 0xdc, 0x80, 0x00, 0x58, 0x89, 0xf6, 0x09,
	0x59, 0x04, 0xfc, 0x2d, 0x63, 0x3f, 0x65, 0x2d, 0xbe, 0x83, 0x1e, 0x5a, 0xed, 0x11, 0xf4, 0x0a,
	0xfd, 0x53, 0xe7, 0x91, 0x98, 0xf6, 0x2f, 0x6d, 0x87, 0xcf, 0x7f, 0x04, 0xed, 0x0e, 0xd8, 0xd9,
	0x57, 0xb2, 0xcf, 0x70, 0x36, 0x01, 0x3e, 0x63, 0xf7, 0xd9, 0xbe, 0x56, 0x7b, 0xfc, 0x7e, 0xb8,
	0xf5, 0x3e, 0xc2, 0x84, 0xa0, 0xc4, 0x73, 0x97, 0x10, 0x5e, 0x22, 0x15, 0x87, 0x5a, 0x61, 0x62,
	0x0c, 0x09, 0x0f, 0x51, 0x42, 0x60, 0x62, 0xd2, 0x27, 0x44, 0x9f, 0x58, 0xc9, 0x7e, 0x8f, 0x17,
	0x90, 0xf0, 0x10, 0x26, 0xdc, 0x1e, 0x2f, 0x5d, 0xc2, 0xed, 0x31, 0x65, 0xf1, 0x57, 0x45, 0x33,
	0x12, 0xf8, 0xa4, 0x0e, 0x58, 0x2b, 0xc8, 0xbe, 0x42, 0x2b, 0x7c, 0xab, 0xe8, 0x57, 0xfa, 0x9a,
	0xee, 0x5e, 0xfc, 0xbb, 0x64, 0xdf, 0xd0, 0x5d, 0x04, 0x1f, 0x38, 0x1b, 0x89, 0xa6, 0x5e, 0xe2,
	0xbb, 0x6e, 0x5f, 0xf1, 0x22, 0xb2, 0x7f, 0xde, 0xff, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x58, 0xae,
	0xed, 0x59, 0x02, 0x04, 0x00, 0x00,
}
