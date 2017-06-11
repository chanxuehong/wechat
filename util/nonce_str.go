package util

import "github.com/chanxuehong/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
