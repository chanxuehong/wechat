// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"testing"
)

func TestCheckSignature(t *testing.T) {
	const (
		signature = "1279a21cfde35abf7b408274163323b5bbcda731"
		timestamp = "1403965161"
		nonce     = "1556813842"
		token     = "your_weixin_token"
	)

	if !checkSignature(signature, timestamp, nonce, token, nil) {
		t.Error("checkSignature failed without buffer")
		return
	}
	if !checkSignature(signature, timestamp, nonce, token, make([]byte, 256)) {
		t.Error("checkSignature failed with buffer")
		return
	}
}

func BenchmarkCheckSignatureWithBuffer(b *testing.B) {
	const (
		signature = "1279a21cfde35abf7b408274163323b5bbcda731"
		timestamp = "1403965161"
		nonce     = "1556813842"
		token     = "your_weixin_token"
	)

	var buffer = make([]byte, 256)

	for i := 0; i < b.N; i++ {
		checkSignature(signature, timestamp, nonce, token, buffer)
	}
}

func BenchmarkCheckSignatureWithoutBuffer(b *testing.B) {
	const (
		signature = "1279a21cfde35abf7b408274163323b5bbcda731"
		timestamp = "1403965161"
		nonce     = "1556813842"
		token     = "your_weixin_token"
	)

	for i := 0; i < b.N; i++ {
		checkSignature(signature, timestamp, nonce, token, nil)
	}
}
