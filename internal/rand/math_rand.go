package rand

import (
	"math/rand"
	"sync"
	"time"
)

func newMathRand() *mathRand {
	return &mathRand{
		r: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

type mathRand struct {
	mu sync.Mutex
	r  *rand.Rand
}

func (r *mathRand) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	n, err = r.r.Read(p)
	r.mu.Unlock()
	return
}

func (r *mathRand) Float32() float32 {
	r.mu.Lock()
	x := r.r.Float32()
	r.mu.Unlock()
	return x
}

func (r *mathRand) Float64() float64 {
	r.mu.Lock()
	x := r.r.Float64()
	r.mu.Unlock()
	return x
}

func (r *mathRand) Int() int {
	r.mu.Lock()
	x := r.r.Int()
	r.mu.Unlock()
	return x
}

func (r *mathRand) Intn(n int) int {
	r.mu.Lock()
	x := r.r.Intn(n)
	r.mu.Unlock()
	return x
}

func (r *mathRand) Int31() int32 {
	r.mu.Lock()
	x := r.r.Int31()
	r.mu.Unlock()
	return x
}

func (r *mathRand) Int31n(n int32) int32 {
	r.mu.Lock()
	x := r.r.Int31n(n)
	r.mu.Unlock()
	return x
}

func (r *mathRand) Int63() int64 {
	r.mu.Lock()
	x := r.r.Int63()
	r.mu.Unlock()
	return x
}

func (r *mathRand) Int63n(n int64) int64 {
	r.mu.Lock()
	x := r.r.Int63n(n)
	r.mu.Unlock()
	return x
}

func (r *mathRand) Uint() uint {
	if intSize == 32 {
		return uint(r.Uint32())
	}
	return uint(r.Uint64())
}

func (r *mathRand) Uint32() uint32 {
	r.mu.Lock()
	x := r.r.Uint32()
	r.mu.Unlock()
	return x
}

func (r *mathRand) Uint64() uint64 {
	r.mu.Lock()
	x := r.r.Uint64()
	r.mu.Unlock()
	return x
}

func (r *mathRand) ExpFloat64() float64 {
	r.mu.Lock()
	x := r.r.ExpFloat64()
	r.mu.Unlock()
	return x
}

func (r *mathRand) NormFloat64() float64 {
	r.mu.Lock()
	x := r.r.NormFloat64()
	r.mu.Unlock()
	return x
}

func (r *mathRand) Seed(seed int64) {
	r.mu.Lock()
	r.r.Seed(seed)
	r.mu.Unlock()
}

func (r *mathRand) Perm(n int) []int {
	r.mu.Lock()
	x := r.r.Perm(n)
	r.mu.Unlock()
	return x
}

func (r *mathRand) Shuffle(n int, swap func(i, j int)) {
	r.mu.Lock()
	r.r.Shuffle(n, swap)
	r.mu.Unlock()
}
