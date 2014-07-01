// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"testing"
)

func BenchmarkGetBufferFromPool(b *testing.B) {
	clt := NewClient("appid", "secret", nil)
	// 先分配一个对象
	buf := clt.getBufferFromPool()
	clt.putBufferToPool(buf)

	for i := 0; i < b.N; i++ {
		buf := clt.getBufferFromPool()
		clt.putBufferToPool(buf)
	}
}

func BenchmarkGetBufferFromNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bytes.NewBuffer(make([]byte, 2<<20)) // 默认 2MB
	}
}
