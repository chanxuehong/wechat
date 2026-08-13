package rand

import (
	"crypto/rand"
)

func newCryptoRand() cryptoRand {
	return cryptoRand{}
}

type cryptoRand struct{}

func (cryptoRand) Read(p []byte) (n int, err error) {
	return rand.Read(p)
}

func (r cryptoRand) Float32() (float32, error) {
	n, err := r.Uint32()
	if err != nil {
		return 0, err
	}
	const significandBits = 24
	const nonSignificandBits = 32 - significandBits
	return float32((n<<nonSignificandBits)>>nonSignificandBits) / (1 << significandBits), nil
}

func (r cryptoRand) Float64() (float64, error) {
	n, err := r.Uint64()
	if err != nil {
		return 0, err
	}
	const significandBits = 53
	const nonSignificandBits = 64 - significandBits
	return float64((n<<nonSignificandBits)>>nonSignificandBits) / (1 << significandBits), nil
}

func (r cryptoRand) Int() (int, error) {
	if intSize == 32 {
		n, err := r.Int31()
		return int(n), err
	}
	n, err := r.Int63()
	return int(n), err
}

func (r cryptoRand) Intn(n int) (int, error) {
	if n <= 0 {
		panic("invalid argument to Intn")
	}
	if n <= maxInt32 {
		x, err := r.Int31n(int32(n))
		return int(x), err
	}
	x, err := r.Int63n(int64(n))
	return int(x), err
}

func (r cryptoRand) Int31() (int32, error) {
	n, err := r.Uint32()
	if err != nil {
		return 0, err
	}
	return int32((n << 1) >> 1), nil
}

func (r cryptoRand) Int31n(n int32) (int32, error) {
	if n <= 0 {
		panic("invalid argument to Int31n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		x, err := r.Int31()
		return x & (n - 1), err
	}
	max := int32(maxInt32 - (1<<31)%uint32(n)) // (max-0+1)%n == 0
	x, err := r.Int31()
	if err != nil {
		return 0, err
	}
	for x > max {
		x, err = r.Int31()
		if err != nil {
			return 0, err
		}
	}
	return x % n, nil
}

func (r cryptoRand) Int63() (int64, error) {
	n, err := r.Uint64()
	if err != nil {
		return 0, err
	}
	return int64((n << 1) >> 1), nil
}

func (r cryptoRand) Int63n(n int64) (int64, error) {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	if n&(n-1) == 0 { // n is power of two, can mask
		x, err := r.Int63()
		return x & (n - 1), err
	}
	max := int64(maxInt64 - (1<<63)%uint64(n)) // (max-0+1)%n == 0
	x, err := r.Int63()
	if err != nil {
		return 0, err
	}
	for x > max {
		x, err = r.Int63()
		if err != nil {
			return 0, err
		}
	}
	return x % n, err
}

func (r cryptoRand) Uint() (uint, error) {
	if intSize == 32 {
		n, err := r.Uint32()
		return uint(n), err
	}
	n, err := r.Uint64()
	return uint(n), err
}

func (r cryptoRand) Uint32() (uint32, error) {
	var buf [4]byte
	if _, err := r.Read(buf[:]); err != nil {
		return 0, err
	}
	return uint32(buf[3]) | uint32(buf[2])<<8 | uint32(buf[1])<<16 | uint32(buf[0])<<24, nil
}

func (r cryptoRand) Uint64() (uint64, error) {
	var buf [8]byte
	if _, err := r.Read(buf[:]); err != nil {
		return 0, err
	}
	return uint64(buf[7]) | uint64(buf[6])<<8 | uint64(buf[5])<<16 | uint64(buf[4])<<24 |
		uint64(buf[3])<<32 | uint64(buf[2])<<40 | uint64(buf[1])<<48 | uint64(buf[0])<<56, nil
}
