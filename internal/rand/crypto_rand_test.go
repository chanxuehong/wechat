package rand

import (
	"bytes"
	"testing"
)

func TestCryptoRand_Read(t *testing.T) {
	r := newCryptoRand()

	buf1 := make([]byte, 32)
	{
		n, err := r.Read(buf1)
		if err != nil {
			t.Error(err)
			return
		}
		if n != len(buf1) {
			t.Error("want equal")
			return
		}
	}

	buf2 := make([]byte, 32)
	{
		n, err := r.Read(buf2)
		if err != nil {
			t.Error(err)
			return
		}
		if n != len(buf2) {
			t.Error("want equal")
			return
		}
	}

	if bytes.Equal(buf1, buf2) {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Float32(t *testing.T) {
	r := newCryptoRand()

	f1, err := r.Float32()
	if err != nil {
		t.Error(err)
		return
	}
	if f1 < 0 || f1 >= 1 {
		t.Error("not in range [0.0,1.0)")
		return
	}

	f2, err := r.Float32()
	if err != nil {
		t.Error(err)
		return
	}
	if f2 < 0 || f2 >= 1 {
		t.Error("not in range [0.0,1.0)")
		return
	}

	f3, err := r.Float32()
	if err != nil {
		t.Error(err)
		return
	}
	if f3 < 0 || f3 >= 1 {
		t.Error("not in range [0.0,1.0)")
		return
	}

	if f1 == f2 && f1 == f3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Float64(t *testing.T) {
	r := newCryptoRand()

	f1, err := r.Float64()
	if err != nil {
		t.Error(err)
		return
	}
	if f1 < 0 || f1 >= 1 {
		t.Error("not in range [0.0,1.0)")
		return
	}

	f2, err := r.Float64()
	if err != nil {
		t.Error(err)
		return
	}
	if f2 < 0 || f2 >= 1 {
		t.Error("not in range [0.0,1.0)")
		return
	}

	f3, err := r.Float64()
	if err != nil {
		t.Error(err)
		return
	}
	if f3 < 0 || f3 >= 1 {
		t.Error("not in range [0.0,1.0)")
		return
	}

	if f1 == f2 && f1 == f3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Int(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Int()
	if err != nil {
		t.Error(err)
		return
	}
	if n1 < 0 {
		t.Error("should be >=0")
		return
	}

	n2, err := r.Int()
	if err != nil {
		t.Error(err)
		return
	}
	if n2 < 0 {
		t.Error("should be >=0")
		return
	}

	n3, err := r.Int()
	if err != nil {
		t.Error(err)
		return
	}
	if n3 < 0 {
		t.Error("should be >=0")
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Intn(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Intn(64)
	if err != nil {
		t.Error(err)
		return
	}
	if n1 < 0 || n1 >= 64 {
		t.Error("not in range")
		return
	}

	n2, err := r.Intn(maxInt32 - 100)
	if err != nil {
		t.Error(err)
		return
	}
	if n2 < 0 || n2 >= maxInt32-100 {
		t.Error("not in range")
		return
	}

	n3, err := r.Intn(maxInt - 100)
	if err != nil {
		t.Error(err)
		return
	}
	if n3 < 0 || n3 >= maxInt-100 {
		t.Error("not in range")
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Int31(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Int31()
	if err != nil {
		t.Error(err)
		return
	}
	if n1 < 0 {
		t.Error("should be >=0")
		return
	}

	n2, err := r.Int31()
	if err != nil {
		t.Error(err)
		return
	}
	if n2 < 0 {
		t.Error("should be >=0")
		return
	}

	n3, err := r.Int31()
	if err != nil {
		t.Error(err)
		return
	}
	if n3 < 0 {
		t.Error("should be >=0")
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Int31n(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Int31n(64)
	if err != nil {
		t.Error(err)
		return
	}
	if n1 < 0 || n1 >= 64 {
		t.Error("not in range")
		return
	}

	n2, err := r.Int31n(64)
	if err != nil {
		t.Error(err)
		return
	}
	if n2 < 0 || n2 >= 64 {
		t.Error("not in range")
		return
	}

	n3, err := r.Int31n(63)
	if err != nil {
		t.Error(err)
		return
	}
	if n3 < 0 || n3 >= 63 {
		t.Error("not in range")
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Int63(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Int63()
	if err != nil {
		t.Error(err)
		return
	}
	if n1 < 0 {
		t.Error("should be >=0")
		return
	}

	n2, err := r.Int63()
	if err != nil {
		t.Error(err)
		return
	}
	if n2 < 0 {
		t.Error("should be >=0")
		return
	}

	n3, err := r.Int63()
	if err != nil {
		t.Error(err)
		return
	}
	if n3 < 0 {
		t.Error("should be >=0")
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Int63n(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Int63n(1 << 48)
	if err != nil {
		t.Error(err)
		return
	}
	if n1 < 0 || n1 >= 1<<48 {
		t.Error("not in range")
		return
	}

	n2, err := r.Int63n(1 << 48)
	if err != nil {
		t.Error(err)
		return
	}
	if n2 < 0 || n2 >= 1<<48 {
		t.Error("not in range")
		return
	}

	n3, err := r.Int63n(1<<48 + 1)
	if err != nil {
		t.Error(err)
		return
	}
	if n3 < 0 || n3 >= 1<<48+1 {
		t.Error("not in range")
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Uint(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Uint()
	if err != nil {
		t.Error(err)
		return
	}

	n2, err := r.Uint()
	if err != nil {
		t.Error(err)
		return
	}

	n3, err := r.Uint()
	if err != nil {
		t.Error(err)
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Uint32(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Uint32()
	if err != nil {
		t.Error(err)
		return
	}

	n2, err := r.Uint32()
	if err != nil {
		t.Error(err)
		return
	}

	n3, err := r.Uint32()
	if err != nil {
		t.Error(err)
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}

func TestCryptoRand_Uint64(t *testing.T) {
	r := newCryptoRand()

	n1, err := r.Uint64()
	if err != nil {
		t.Error(err)
		return
	}

	n2, err := r.Uint64()
	if err != nil {
		t.Error(err)
		return
	}

	n3, err := r.Uint64()
	if err != nil {
		t.Error(err)
		return
	}

	if n1 == n2 && n1 == n3 {
		t.Error("expect not equal")
		return
	}
}
