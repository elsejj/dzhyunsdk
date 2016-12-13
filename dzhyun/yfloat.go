package dzhyun

import "fmt"
import "errors"
import "math"

type YFloat int64

func (v YFloat) MarshalJSON() ([]byte, error) {
	number, dp, err := UnmakeValue(int64(v))
	if err != nil {
		return nil, err
	}

	format := fmt.Sprintf("%%.%df", dp)

	text := fmt.Sprintf(format, number)

	return []byte(text), nil

}

const VT_null = int64(2)

var VT_dp = [...]int64{10, 1, 0, 3, 4, 5, 6, 7, 8, 9} // dp: decimal position
const VT_max = int64(11)
const VT_min = int64(12)
const VT_e = int64(13) // e: exponential

var Multiples = [...]float64{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
}

const (
	Null        = int64(2 << 16)
	MaxInfinite = int64(11<<16 + 0)
	MaxInt32    = int64(11<<16 + 1)
	MaxInt64    = int64(11<<16 + 2)
	MaxFloat32  = int64(11<<16 + 3)
	MaxFloat64  = int64(11<<16 + 4)
	MaxUInt32   = int64(11<<16 + 5)
	MaxUInt64   = int64(11<<16 + 6)
	MinInfinite = int64(12<<16 + 0)
	MinInt32    = int64(12<<16 + 1)
	MinInt64    = int64(12<<16 + 2)
	MinFloat32  = int64(12<<16 + 3)
	MinFloat64  = int64(12<<16 + 4)
)

func ValueType(value int64) int64 {
	B := (value >> 16) & 0xFF
	return B & 0x0F
}

func MakeValue(value float64, dp int) int64 {
	L, H := int64(0), int64(0)

	if value < 0 {
		H = 1
		value = -value
	}

	temp := int64(0)

	switch dp {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		L, temp = VT_dp[dp], int64(value*Multiples[dp]+0.5)
	default:
		panic(fmt.Sprintf("MakeValue: invalid decimal places %v", dp))
	}

	return ((temp >> 16) << 24) + (((H << 4) + L) << 16) + (temp & 0xFFFF)
}

func UnmakeValue(value int64) (float64, int, error) {
	if value < 0 || value == Null {
		return 0, 0, errors.New("UnmakeValue: invalid value")
	}

	B := (value >> 16) & 0xFF
	L := B & 0x0F
	H := (B >> 4) & 0x0F

	Bx := ((value >> 24) << 16) + (value & 0xFFFF)
	temp := float64(0)
	dp := 0

	switch L {
	case 10:
		dp, temp = 0, float64(Bx)
	case 1:
		dp, temp = 1, float64(Bx)/10
	case 0:
		dp, temp = 2, float64(Bx)/100
	case 3:
		dp, temp = 3, float64(Bx)/1000
	case 4:
		dp, temp = 4, float64(Bx)/10000
	case 5:
		dp, temp = 5, float64(Bx)/100000
	case 6:
		dp, temp = 6, float64(Bx)/1000000
	case 7:
		dp, temp = 7, float64(Bx)/10000000
	case 8:
		dp, temp = 8, float64(Bx)/100000000
	case 9:
		dp, temp = 9, float64(Bx)/1000000000
	default:
		return 0, 0, errors.New(fmt.Sprintf("UnmakeValue: unknown L %v", L))
	}

	if H != 0 {
		temp = -temp
	}

	return temp, dp, nil
}

func MakeValueE(value float64, expo int) int64 {
	if expo < -128 || expo > 127 {
		panic(fmt.Sprintf("MakeValueE: invalid expo value %v", expo))
	}

	const multiple float64 = 1e12

	L, H := int64(VT_e), int64(0)

	if value < 0 {
		H = 1
		value = -value
	}

	if value > 10 {
		panic(fmt.Sprintf("MakeValueE: invalid main value %v", value))
	}

	temp := int64(value*multiple + 0.5)

	return ((temp >> 8) << 24) + (((H << 4) + L) << 16) + ((temp & 0x0F) << 8) + int64(expo+128)
}

func UnmakeValueE(value int64) (float64, int, error) {
	if value < 0 {
		return 0, 0, errors.New("UnmakeValueE: invalid value")
	}

	const multiple float64 = 1e12

	B := (value >> 16) & 0xFF
	L := B & 0x0F
	H := (B >> 4) & 0x0F

	if L != int64(VT_e) {
		return 0, 0, errors.New(fmt.Sprintf("UnmakeValueE: invalid L %v", L))
	}

	Bx := ((value >> 24) << 8) + ((value >> 8) & 0xFF)
	expo := int(value&0xFF) - 128
	temp := float64(Bx) / multiple

	if H != 0 {
		temp = -temp
	}

	return temp, expo, nil
}

var eNotValue = errors.New("not a value")

func getCompareValue(yv int64) (ov float64, err error) {
	switch ValueType(yv) {
	case VT_null, VT_min, VT_max:
		err = eNotValue
	case VT_e:
		var expo int
		ov, expo, err = UnmakeValueE(yv)
		if err == nil {
			ov = ov * math.Pow10(expo)
		}
	default:
		ov, _, err = UnmakeValue(yv)
	}

	return ov, err
}

// less than
func LT(a int64, b int64) bool {
	var x, y float64
	var e error

	if x, e = getCompareValue(a); e != nil {
		return false
	}

	if y, e = getCompareValue(b); e != nil {
		return false
	}

	return x < y
}

// greater than
func GT(a int64, b int64) bool {
	var x, y float64
	var e error

	if x, e = getCompareValue(a); e != nil {
		return false
	}

	if y, e = getCompareValue(b); e != nil {
		return false
	}

	return x > y
}

func Equal(a int64, b int64) bool {
	var x, y float64
	var e error

	if x, e = getCompareValue(a); e != nil {
		return false
	}

	if y, e = getCompareValue(b); e != nil {
		return false
	}

	return x == y
}

// less than or equal
func LTE(a int64, b int64) bool {
	var x, y float64
	var e error

	if x, e = getCompareValue(a); e != nil {
		return false
	}

	if y, e = getCompareValue(b); e != nil {
		return false
	}

	return x <= y
}

// greater than or equal
func GTE(a int64, b int64) bool {
	var x, y float64
	var e error

	if x, e = getCompareValue(a); e != nil {
		return false
	}

	if y, e = getCompareValue(b); e != nil {
		return false
	}

	return x >= y
}

// compare
// return 0(equal), 1(GT), -1(LT), -2(means not a value)
func Compare(a int64, b int64) int {
	var x, y float64
	var e error

	if x, e = getCompareValue(a); e != nil {
		return -2
	}

	if y, e = getCompareValue(b); e != nil {
		return -2
	}

	if x > y {
		return 1
	}
	if x < y {
		return -1
	}

	return 0
}
