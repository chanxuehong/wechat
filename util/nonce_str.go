package util

import "gopkg.in/chanxuehong/wechat.v2/internal/rand"

func NonceStr() string {
	return string(rand.NewHex())
}
