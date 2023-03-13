package util

import (
	"encoding/hex"
	"math/rand"
	"time"
)

var randSeed = rand.New(rand.NewSource(time.Now().UnixNano()))

func randomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := randSeed.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func NonceStr() string {
	str, err := randomHex(32)
	if err != nil {
		return ""
	}
	return str
}
