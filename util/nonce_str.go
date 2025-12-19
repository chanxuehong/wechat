package util

import (
	"crypto/rand"
	"encoding/hex"
	mathrand "math/rand"
)

func NonceStr() string {
	var buf [16]byte
	if _, err := rand.Read(buf[:]); err != nil {
		mathrand.Read(buf[:])
	}
	return hex.EncodeToString(buf[:])
}
