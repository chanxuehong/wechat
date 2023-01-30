package util

import (
	"strconv"
	"strings"
)

// Uint64 support string quoted number in json
type Uint64 uint64

// UnmarshalJSON implement json Unmarshal interface
func (u64 *Uint64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseUint(string(b), 10, 64)
	*u64 = Uint64(i)
	return
}

func (u64 Uint64) Uint64() uint64 {
	return uint64(u64)
}

// Int64 support string quoted number in json
type Int64 int64

// UnmarshalJSON implement json Unmarshal interface
func (i64 *Int64) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	i, _ := strconv.ParseInt(string(b), 10, 64)
	*i64 = Int64(i)
	return
}

func (i64 Int64) Int64() int64 {
	return int64(i64)
}

// MoneyFloat support string quoted number in json
type MoneyFloat float64

// UnmarshalJSON implement json Unmarshal interface
func (f64 *MoneyFloat) UnmarshalJSON(b []byte) (err error) {
	if b[0] == '"' && b[len(b)-1] == '"' {
		b = b[1 : len(b)-1]
	}
	str := strings.TrimPrefix(string(b), "￥")
	i, _ := strconv.ParseFloat(str, 64)
	*f64 = MoneyFloat(i)
	return
}

func (f64 MoneyFloat) Float64() float64 {
	return float64(f64)
}
