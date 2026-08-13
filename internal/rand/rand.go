package rand

const intSize = 32 << ((^uint(0) >> 32) & 1)

const (
	maxInt   = 1<<(intSize-1) - 1
	maxInt32 = 1<<31 - 1
	maxInt64 = 1<<63 - 1
)

// Read generates len(p) random bytes and writes them into p. It always returns len(p) and a nil error.
func Read(p []byte) (n int, err error) { return globalRand.Read(p) }

// Float32 returns, as a float32, a random number in the half-open interval [0.0,1.0).
func Float32() float32 { return globalRand.Float32() }

// Float64 returns, as a float64, a random number in the half-open interval [0.0,1.0).
func Float64() float64 { return globalRand.Float64() }

// Int returns a non-negative random int.
func Int() int { return globalRand.Int() }

// Intn returns, as an int, a non-negative random number in the half-open interval [0,n). It panics if n <= 0.
func Intn(n int) int { return globalRand.Intn(n) }

// Int31 returns a non-negative random 31-bit integer as an int32.
func Int31() int32 { return globalRand.Int31() }

// Int31n returns, as an int32, a non-negative random number in the half-open interval [0,n). It panics if n <= 0.
func Int31n(n int32) int32 { return globalRand.Int31n(n) }

// Int63 returns a non-negative random 63-bit integer as an int64.
func Int63() int64 { return globalRand.Int63() }

// Int63n returns, as an int64, a non-negative random number in the half-open interval [0,n). It panics if n <= 0.
func Int63n(n int64) int64 { return globalRand.Int63n(n) }

// Uint returns a random uint.
func Uint() uint { return globalRand.Uint() }

// Uint32 returns a random 32-bit value as a uint32.
func Uint32() uint32 { return globalRand.Uint32() }

// Uint64 returns a random 64-bit value as a uint64.
func Uint64() uint64 { return globalRand.Uint64() }

var globalRand = NewRand()

// NewRand returns a new Rand.
func NewRand() *Rand {
	return &Rand{
		cryptoRand: newCryptoRand(),
		mathRand:   newMathRand(),
	}
}

// Rand is similar to "math/rand.Rand", but it prefers "crypto/rand" to generate random bytes.
//
// unlike "math/rand.Rand", all methods of *Rand are safe for concurrent use by multiple goroutines.
type Rand struct {
	cryptoRand cryptoRand
	mathRand   *mathRand
}

// Read generates len(p) random bytes and writes them into p. It always returns len(p) and a nil error.
func (r *Rand) Read(p []byte) (n int, err error) {
	if n, err = r.cryptoRand.Read(p); err == nil {
		return
	}
	n, err = r.mathRand.Read(p)
	return
}

// Float32 returns, as a float32, a random number in the half-open interval [0.0,1.0).
func (r *Rand) Float32() float32 {
	if x, err := r.cryptoRand.Float32(); err == nil {
		return x
	}
	return r.mathRand.Float32()
}

// Float64 returns, as a float64, a random number in the half-open interval [0.0,1.0).
func (r *Rand) Float64() float64 {
	if x, err := r.cryptoRand.Float64(); err == nil {
		return x
	}
	return r.mathRand.Float64()
}

// Int returns a non-negative random int.
func (r *Rand) Int() int {
	if x, err := r.cryptoRand.Int(); err == nil {
		return x
	}
	return r.mathRand.Int()
}

// Intn returns, as an int, a non-negative random number in the half-open interval [0,n). It panics if n <= 0.
func (r *Rand) Intn(n int) int {
	if x, err := r.cryptoRand.Intn(n); err == nil {
		return x
	}
	return r.mathRand.Intn(n)
}

// Int31 returns a non-negative random 31-bit integer as an int32.
func (r *Rand) Int31() int32 {
	if x, err := r.cryptoRand.Int31(); err == nil {
		return x
	}
	return r.mathRand.Int31()
}

// Int31n returns, as an int32, a non-negative random number in the half-open interval [0,n). It panics if n <= 0.
func (r *Rand) Int31n(n int32) int32 {
	if x, err := r.cryptoRand.Int31n(n); err == nil {
		return x
	}
	return r.mathRand.Int31n(n)
}

// Int63 returns a non-negative random 63-bit integer as an int64.
func (r *Rand) Int63() int64 {
	if x, err := r.cryptoRand.Int63(); err == nil {
		return x
	}
	return r.mathRand.Int63()
}

// Int63n returns, as an int64, a non-negative random number in the half-open interval [0,n). It panics if n <= 0.
func (r *Rand) Int63n(n int64) int64 {
	if x, err := r.cryptoRand.Int63n(n); err == nil {
		return x
	}
	return r.mathRand.Int63n(n)
}

// Uint returns a random uint.
func (r *Rand) Uint() uint {
	if x, err := r.cryptoRand.Uint(); err == nil {
		return x
	}
	return r.mathRand.Uint()
}

// Uint32 returns a random 32-bit value as a uint32.
func (r *Rand) Uint32() uint32 {
	if x, err := r.cryptoRand.Uint32(); err == nil {
		return x
	}
	return r.mathRand.Uint32()
}

// Uint64 returns a random 64-bit value as a uint64.
func (r *Rand) Uint64() uint64 {
	if x, err := r.cryptoRand.Uint64(); err == nil {
		return x
	}
	return r.mathRand.Uint64()
}
