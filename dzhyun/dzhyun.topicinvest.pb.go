// Code generated by protoc-gen-go.
// source: dzhyun.topicinvest.proto
// DO NOT EDIT!

package dzhyun

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// 领涨股
type LingZhangGuShuJu struct {
	Obj               string `protobuf:"bytes,1,opt,name=Obj" json:"Obj,omitempty"`
	ZhongWenJianCheng string `protobuf:"bytes,2,opt,name=ZhongWenJianCheng" json:"ZhongWenJianCheng,omitempty"`
	ZuiXinJia         int64  `protobuf:"varint,3,opt,name=ZuiXinJia" json:"ZuiXinJia,omitempty"`
	ZhangFu           int64  `protobuf:"varint,4,opt,name=ZhangFu" json:"ZhangFu,omitempty"`
}

func (m *LingZhangGuShuJu) Reset()                    { *m = LingZhangGuShuJu{} }
func (m *LingZhangGuShuJu) String() string            { return proto.CompactTextString(m) }
func (*LingZhangGuShuJu) ProtoMessage()               {}
func (*LingZhangGuShuJu) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{0} }

func (m *LingZhangGuShuJu) GetObj() string {
	if m != nil {
		return m.Obj
	}
	return ""
}

func (m *LingZhangGuShuJu) GetZhongWenJianCheng() string {
	if m != nil {
		return m.ZhongWenJianCheng
	}
	return ""
}

func (m *LingZhangGuShuJu) GetZuiXinJia() int64 {
	if m != nil {
		return m.ZuiXinJia
	}
	return 0
}

func (m *LingZhangGuShuJu) GetZhangFu() int64 {
	if m != nil {
		return m.ZhangFu
	}
	return 0
}

// 主题排名
type TopicInvestRank struct {
	BianHao                   int64 `protobuf:"varint,1,opt,name=BianHao" json:"BianHao,omitempty"`
	ShiFouReMenZhuTi          int64 `protobuf:"varint,2,opt,name=ShiFouReMenZhuTi" json:"ShiFouReMenZhuTi,omitempty"`
	RiPingJunZhangFuPaiMing14 int64 `protobuf:"varint,3,opt,name=RiPingJunZhangFuPaiMing14" json:"RiPingJunZhangFuPaiMing14,omitempty"`
	RiPingJunZhangFuPaiMing30 int64 `protobuf:"varint,4,opt,name=RiPingJunZhangFuPaiMing30" json:"RiPingJunZhangFuPaiMing30,omitempty"`
	RiReDuZhi14               int64 `protobuf:"varint,5,opt,name=RiReDuZhi14" json:"RiReDuZhi14,omitempty"`
	RiReDuZhi30               int64 `protobuf:"varint,6,opt,name=RiReDuZhi30" json:"RiReDuZhi30,omitempty"`
	RiReDuZhi10               int64 `protobuf:"varint,7,opt,name=RiReDuZhi10" json:"RiReDuZhi10,omitempty"`
}

func (m *TopicInvestRank) Reset()                    { *m = TopicInvestRank{} }
func (m *TopicInvestRank) String() string            { return proto.CompactTextString(m) }
func (*TopicInvestRank) ProtoMessage()               {}
func (*TopicInvestRank) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{1} }

func (m *TopicInvestRank) GetBianHao() int64 {
	if m != nil {
		return m.BianHao
	}
	return 0
}

func (m *TopicInvestRank) GetShiFouReMenZhuTi() int64 {
	if m != nil {
		return m.ShiFouReMenZhuTi
	}
	return 0
}

func (m *TopicInvestRank) GetRiPingJunZhangFuPaiMing14() int64 {
	if m != nil {
		return m.RiPingJunZhangFuPaiMing14
	}
	return 0
}

func (m *TopicInvestRank) GetRiPingJunZhangFuPaiMing30() int64 {
	if m != nil {
		return m.RiPingJunZhangFuPaiMing30
	}
	return 0
}

func (m *TopicInvestRank) GetRiReDuZhi14() int64 {
	if m != nil {
		return m.RiReDuZhi14
	}
	return 0
}

func (m *TopicInvestRank) GetRiReDuZhi30() int64 {
	if m != nil {
		return m.RiReDuZhi30
	}
	return 0
}

func (m *TopicInvestRank) GetRiReDuZhi10() int64 {
	if m != nil {
		return m.RiReDuZhi10
	}
	return 0
}

type TopicInvest struct {
	BianHao                   int64             `protobuf:"varint,1,opt,name=BianHao" json:"BianHao,omitempty"`
	MingCheng                 string            `protobuf:"bytes,2,opt,name=MingCheng" json:"MingCheng,omitempty"`
	ShiJian                   int64             `protobuf:"varint,3,opt,name=ShiJian" json:"ShiJian,omitempty"`
	ZhangFu                   int64             `protobuf:"varint,4,opt,name=ZhangFu" json:"ZhangFu,omitempty"`
	LiangBi                   int64             `protobuf:"varint,5,opt,name=LiangBi" json:"LiangBi,omitempty"`
	HuanShou                  int64             `protobuf:"varint,6,opt,name=HuanShou" json:"HuanShou,omitempty"`
	ShangZhangJiaShu          int64             `protobuf:"varint,7,opt,name=ShangZhangJiaShu" json:"ShangZhangJiaShu,omitempty"`
	PingPanJiaShu             int64             `protobuf:"varint,8,opt,name=PingPanJiaShu" json:"PingPanJiaShu,omitempty"`
	XiaDieJiaShu              int64             `protobuf:"varint,9,opt,name=XiaDieJiaShu" json:"XiaDieJiaShu,omitempty"`
	LingZhangGu               *LingZhangGuShuJu `protobuf:"bytes,10,opt,name=LingZhangGu" json:"LingZhangGu,omitempty"`
	ShiFouReMenZhuTi          int64             `protobuf:"varint,11,opt,name=ShiFouReMenZhuTi" json:"ShiFouReMenZhuTi,omitempty"`
	RiPingJunZhangFuPaiMing14 int64             `protobuf:"varint,12,opt,name=RiPingJunZhangFuPaiMing14" json:"RiPingJunZhangFuPaiMing14,omitempty"`
	RiPingJunZhangFuPaiMing30 int64             `protobuf:"varint,13,opt,name=RiPingJunZhangFuPaiMing30" json:"RiPingJunZhangFuPaiMing30,omitempty"`
	RiReDuZhi14               int64             `protobuf:"varint,14,opt,name=RiReDuZhi14" json:"RiReDuZhi14,omitempty"`
	RiReDuZhi30               int64             `protobuf:"varint,15,opt,name=RiReDuZhi30" json:"RiReDuZhi30,omitempty"`
	RiReDuZhi10               int64             `protobuf:"varint,16,opt,name=RiReDuZhi10" json:"RiReDuZhi10,omitempty"`
	GeGuShu                   int64             `protobuf:"varint,17,opt,name=GeGuShu" json:"GeGuShu,omitempty"`
	DangYueZhangFu            int64             `protobuf:"varint,18,opt,name=DangYueZhangFu" json:"DangYueZhangFu,omitempty"`
}

func (m *TopicInvest) Reset()                    { *m = TopicInvest{} }
func (m *TopicInvest) String() string            { return proto.CompactTextString(m) }
func (*TopicInvest) ProtoMessage()               {}
func (*TopicInvest) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{2} }

func (m *TopicInvest) GetBianHao() int64 {
	if m != nil {
		return m.BianHao
	}
	return 0
}

func (m *TopicInvest) GetMingCheng() string {
	if m != nil {
		return m.MingCheng
	}
	return ""
}

func (m *TopicInvest) GetShiJian() int64 {
	if m != nil {
		return m.ShiJian
	}
	return 0
}

func (m *TopicInvest) GetZhangFu() int64 {
	if m != nil {
		return m.ZhangFu
	}
	return 0
}

func (m *TopicInvest) GetLiangBi() int64 {
	if m != nil {
		return m.LiangBi
	}
	return 0
}

func (m *TopicInvest) GetHuanShou() int64 {
	if m != nil {
		return m.HuanShou
	}
	return 0
}

func (m *TopicInvest) GetShangZhangJiaShu() int64 {
	if m != nil {
		return m.ShangZhangJiaShu
	}
	return 0
}

func (m *TopicInvest) GetPingPanJiaShu() int64 {
	if m != nil {
		return m.PingPanJiaShu
	}
	return 0
}

func (m *TopicInvest) GetXiaDieJiaShu() int64 {
	if m != nil {
		return m.XiaDieJiaShu
	}
	return 0
}

func (m *TopicInvest) GetLingZhangGu() *LingZhangGuShuJu {
	if m != nil {
		return m.LingZhangGu
	}
	return nil
}

func (m *TopicInvest) GetShiFouReMenZhuTi() int64 {
	if m != nil {
		return m.ShiFouReMenZhuTi
	}
	return 0
}

func (m *TopicInvest) GetRiPingJunZhangFuPaiMing14() int64 {
	if m != nil {
		return m.RiPingJunZhangFuPaiMing14
	}
	return 0
}

func (m *TopicInvest) GetRiPingJunZhangFuPaiMing30() int64 {
	if m != nil {
		return m.RiPingJunZhangFuPaiMing30
	}
	return 0
}

func (m *TopicInvest) GetRiReDuZhi14() int64 {
	if m != nil {
		return m.RiReDuZhi14
	}
	return 0
}

func (m *TopicInvest) GetRiReDuZhi30() int64 {
	if m != nil {
		return m.RiReDuZhi30
	}
	return 0
}

func (m *TopicInvest) GetRiReDuZhi10() int64 {
	if m != nil {
		return m.RiReDuZhi10
	}
	return 0
}

func (m *TopicInvest) GetGeGuShu() int64 {
	if m != nil {
		return m.GeGuShu
	}
	return 0
}

func (m *TopicInvest) GetDangYueZhangFu() int64 {
	if m != nil {
		return m.DangYueZhangFu
	}
	return 0
}

// 历史行情
type LiShiHangQing struct {
	ShiJian int64 `protobuf:"varint,1,opt,name=ShiJian" json:"ShiJian,omitempty"`
	ZhangFu int64 `protobuf:"varint,2,opt,name=ZhangFu" json:"ZhangFu,omitempty"`
}

func (m *LiShiHangQing) Reset()                    { *m = LiShiHangQing{} }
func (m *LiShiHangQing) String() string            { return proto.CompactTextString(m) }
func (*LiShiHangQing) ProtoMessage()               {}
func (*LiShiHangQing) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{3} }

func (m *LiShiHangQing) GetShiJian() int64 {
	if m != nil {
		return m.ShiJian
	}
	return 0
}

func (m *LiShiHangQing) GetZhangFu() int64 {
	if m != nil {
		return m.ZhangFu
	}
	return 0
}

// 历史走势
type LiShiZouShi struct {
	HangQing []*LiShiHangQing `protobuf:"bytes,1,rep,name=HangQing" json:"HangQing,omitempty"`
}

func (m *LiShiZouShi) Reset()                    { *m = LiShiZouShi{} }
func (m *LiShiZouShi) String() string            { return proto.CompactTextString(m) }
func (*LiShiZouShi) ProtoMessage()               {}
func (*LiShiZouShi) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{4} }

func (m *LiShiZouShi) GetHangQing() []*LiShiHangQing {
	if m != nil {
		return m.HangQing
	}
	return nil
}

type TopicInvestHistory struct {
	BianHao   int64        `protobuf:"varint,1,opt,name=BianHao" json:"BianHao,omitempty"`
	MingCheng string       `protobuf:"bytes,2,opt,name=MingCheng" json:"MingCheng,omitempty"`
	ShiJian   int64        `protobuf:"varint,3,opt,name=ShiJian" json:"ShiJian,omitempty"`
	LiShi     *LiShiZouShi `protobuf:"bytes,4,opt,name=LiShi" json:"LiShi,omitempty"`
}

func (m *TopicInvestHistory) Reset()                    { *m = TopicInvestHistory{} }
func (m *TopicInvestHistory) String() string            { return proto.CompactTextString(m) }
func (*TopicInvestHistory) ProtoMessage()               {}
func (*TopicInvestHistory) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{5} }

func (m *TopicInvestHistory) GetBianHao() int64 {
	if m != nil {
		return m.BianHao
	}
	return 0
}

func (m *TopicInvestHistory) GetMingCheng() string {
	if m != nil {
		return m.MingCheng
	}
	return ""
}

func (m *TopicInvestHistory) GetShiJian() int64 {
	if m != nil {
		return m.ShiJian
	}
	return 0
}

func (m *TopicInvestHistory) GetLiShi() *LiShiZouShi {
	if m != nil {
		return m.LiShi
	}
	return nil
}

// 主题投资成份股数据
type TopicInvestInfo struct {
	BianHao    int64    `protobuf:"varint,1,opt,name=BianHao" json:"BianHao,omitempty"`
	MingCheng  string   `protobuf:"bytes,2,opt,name=MingCheng" json:"MingCheng,omitempty"`
	ChengFenGu []string `protobuf:"bytes,3,rep,name=ChengFenGu" json:"ChengFenGu,omitempty"`
	LeiBie     int64    `protobuf:"varint,4,opt,name=LeiBie" json:"LeiBie,omitempty"`
}

func (m *TopicInvestInfo) Reset()                    { *m = TopicInvestInfo{} }
func (m *TopicInvestInfo) String() string            { return proto.CompactTextString(m) }
func (*TopicInvestInfo) ProtoMessage()               {}
func (*TopicInvestInfo) Descriptor() ([]byte, []int) { return fileDescriptor12, []int{6} }

func (m *TopicInvestInfo) GetBianHao() int64 {
	if m != nil {
		return m.BianHao
	}
	return 0
}

func (m *TopicInvestInfo) GetMingCheng() string {
	if m != nil {
		return m.MingCheng
	}
	return ""
}

func (m *TopicInvestInfo) GetChengFenGu() []string {
	if m != nil {
		return m.ChengFenGu
	}
	return nil
}

func (m *TopicInvestInfo) GetLeiBie() int64 {
	if m != nil {
		return m.LeiBie
	}
	return 0
}

func init() {
	proto.RegisterType((*LingZhangGuShuJu)(nil), "dzhyun.LingZhangGuShuJu")
	proto.RegisterType((*TopicInvestRank)(nil), "dzhyun.TopicInvestRank")
	proto.RegisterType((*TopicInvest)(nil), "dzhyun.TopicInvest")
	proto.RegisterType((*LiShiHangQing)(nil), "dzhyun.LiShiHangQing")
	proto.RegisterType((*LiShiZouShi)(nil), "dzhyun.LiShiZouShi")
	proto.RegisterType((*TopicInvestHistory)(nil), "dzhyun.TopicInvestHistory")
	proto.RegisterType((*TopicInvestInfo)(nil), "dzhyun.TopicInvestInfo")
}

func init() { proto.RegisterFile("dzhyun.topicinvest.proto", fileDescriptor12) }

var fileDescriptor12 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x55, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x6b, 0x9a, 0x26, 0xe3, 0xa6, 0x4d, 0x17, 0x81, 0x16, 0x54, 0x21, 0xcb, 0x42, 0x28,
	0x20, 0x14, 0xe5, 0xa3, 0x27, 0xc4, 0x01, 0xa5, 0x51, 0x92, 0x46, 0xa9, 0x08, 0xeb, 0x4a, 0x14,
	0xdf, 0xb6, 0x65, 0xf1, 0x0e, 0x88, 0x75, 0xd5, 0x66, 0x91, 0xca, 0x8d, 0x2b, 0x37, 0x7e, 0x0a,
	0x47, 0xfe, 0x1d, 0xf2, 0x66, 0xdd, 0xd8, 0x09, 0x69, 0x45, 0x11, 0xa7, 0x78, 0xde, 0x3c, 0xef,
	0xbe, 0x99, 0x79, 0xe3, 0x00, 0x7d, 0xff, 0x55, 0x5e, 0x6a, 0xd5, 0x98, 0x26, 0x67, 0x78, 0x8a,
	0xea, 0x8b, 0xb8, 0x98, 0x36, 0xce, 0xce, 0x93, 0x69, 0x42, 0x4a, 0xb3, 0x4c, 0xf0, 0xdd, 0x81,
	0xda, 0x18, 0x55, 0x1c, 0x49, 0xae, 0xe2, 0x81, 0x0e, 0xa5, 0x1e, 0x69, 0x52, 0x03, 0xf7, 0xf5,
	0xc9, 0x47, 0xea, 0xf8, 0x4e, 0xbd, 0xc2, 0xd2, 0x47, 0xf2, 0x1c, 0x76, 0x22, 0x99, 0xa8, 0xf8,
	0xad, 0x50, 0x23, 0xe4, 0x6a, 0x5f, 0x0a, 0x15, 0xd3, 0x35, 0x93, 0x5f, 0x4e, 0x90, 0x5d, 0xa8,
	0x44, 0x1a, 0x8f, 0x31, 0x85, 0xa8, 0xeb, 0x3b, 0x75, 0x97, 0xcd, 0x01, 0x42, 0x61, 0xc3, 0xdc,
	0xd6, 0xd7, 0xf4, 0x8e, 0xc9, 0x65, 0x61, 0xf0, 0x6b, 0x0d, 0xb6, 0x8f, 0x52, 0xa9, 0x07, 0x46,
	0x2a, 0xe3, 0xea, 0x53, 0xca, 0xee, 0x22, 0x57, 0x43, 0x9e, 0x18, 0x3d, 0x2e, 0xcb, 0x42, 0xf2,
	0x0c, 0x6a, 0xa1, 0xc4, 0x7e, 0xa2, 0x99, 0x38, 0x14, 0x2a, 0x92, 0xfa, 0x08, 0x8d, 0x24, 0x97,
	0x2d, 0xe1, 0xe4, 0x25, 0x3c, 0x60, 0x38, 0x41, 0x15, 0x8f, 0xb4, 0xb2, 0xb7, 0x4d, 0x38, 0x1e,
	0xa2, 0x8a, 0x5b, 0x7b, 0x56, 0xe1, 0x6a, 0xc2, 0x35, 0x6f, 0x77, 0x9a, 0xb6, 0x86, 0xd5, 0x04,
	0xe2, 0x83, 0xc7, 0x90, 0x89, 0x9e, 0x8e, 0x24, 0xb6, 0xf6, 0xe8, 0xba, 0xe1, 0xe7, 0xa1, 0x02,
	0xa3, 0xd3, 0xa4, 0xa5, 0x05, 0xc6, 0xe2, 0x19, 0x4d, 0xba, 0xb1, 0x78, 0x46, 0x33, 0xf8, 0xb9,
	0x0e, 0x5e, 0xae, 0x77, 0xd7, 0xf4, 0x6d, 0x17, 0x2a, 0xa9, 0xb2, 0xfc, 0x0c, 0xe7, 0x40, 0xfa,
	0x5e, 0x28, 0x31, 0x9d, 0xa5, 0xed, 0x4b, 0x16, 0xae, 0x9e, 0x5b, 0x9a, 0x19, 0x23, 0x57, 0x71,
	0x17, 0x6d, 0x75, 0x59, 0x48, 0x1e, 0x42, 0x79, 0xa8, 0xb9, 0x0a, 0x65, 0xa2, 0x6d, 0x59, 0x57,
	0xf1, 0x6c, 0x7e, 0xdc, 0x5a, 0x6f, 0x84, 0x3c, 0x94, 0xda, 0x16, 0xb6, 0x84, 0x93, 0xc7, 0x50,
	0x4d, 0xdb, 0x3b, 0xe1, 0xca, 0x12, 0xcb, 0x86, 0x58, 0x04, 0x49, 0x00, 0x9b, 0xc7, 0xc8, 0x7b,
	0x28, 0x2c, 0xa9, 0x62, 0x48, 0x05, 0x8c, 0xbc, 0x00, 0x2f, 0xe7, 0x77, 0x0a, 0xbe, 0x53, 0xf7,
	0xda, 0xb4, 0x61, 0x17, 0x65, 0x71, 0x15, 0x58, 0x9e, 0xfc, 0x47, 0xc7, 0x79, 0xb7, 0x71, 0xdc,
	0xe6, 0x3f, 0x39, 0xae, 0xfa, 0x97, 0x8e, 0xdb, 0xba, 0xd1, 0x71, 0xdb, 0x37, 0x3a, 0xae, 0xb6,
	0xe4, 0xb8, 0x74, 0xea, 0x03, 0x61, 0xfa, 0x44, 0x77, 0x66, 0x53, 0xb7, 0x21, 0x79, 0x02, 0x5b,
	0x3d, 0xae, 0xe2, 0x77, 0x5a, 0x64, 0x86, 0x21, 0x86, 0xb0, 0x80, 0x06, 0xfb, 0x50, 0x1d, 0x63,
	0x28, 0x71, 0xc8, 0x55, 0xfc, 0x06, 0x8b, 0xe6, 0x73, 0x56, 0x9a, 0x6f, 0xad, 0xf8, 0xd1, 0x78,
	0x95, 0x0e, 0x34, 0x94, 0x18, 0x25, 0x3a, 0x94, 0x48, 0x5a, 0x50, 0xce, 0x8e, 0xa3, 0x8e, 0xef,
	0xd6, 0xbd, 0xf6, 0xbd, 0xf9, 0x70, 0x73, 0x77, 0xb1, 0x2b, 0x5a, 0xf0, 0xc3, 0x01, 0x92, 0x5b,
	0x9d, 0x21, 0x5e, 0x4c, 0x93, 0xf3, 0xcb, 0xff, 0xb0, 0x41, 0x4f, 0x61, 0xdd, 0x68, 0x30, 0xfb,
	0xe3, 0xb5, 0xef, 0x16, 0x84, 0xcd, 0xf4, 0xb3, 0x19, 0x23, 0xf8, 0xe6, 0x14, 0x3e, 0x85, 0x07,
	0xea, 0x43, 0x72, 0x6b, 0x41, 0x8f, 0x00, 0xcc, 0x43, 0x5f, 0xa8, 0x81, 0xa6, 0xae, 0xef, 0xd6,
	0x2b, 0x2c, 0x87, 0x90, 0xfb, 0x50, 0x1a, 0x0b, 0xec, 0xa2, 0xb0, 0x7b, 0x6d, 0xa3, 0x2e, 0x81,
	0xda, 0x69, 0xf2, 0x39, 0x13, 0x69, 0xfe, 0x37, 0x4e, 0x4a, 0xe6, 0xa7, 0xf3, 0x3b, 0x00, 0x00,
	0xff, 0xff, 0x77, 0x06, 0xac, 0xaf, 0x5a, 0x06, 0x00, 0x00,
}