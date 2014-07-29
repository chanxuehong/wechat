// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"testing"
)

var testClient = func() *Client {
	// 填入正确的 appid, appsecret
	clt := NewClient("appid", "appsecret", nil)

	// 预热
	buf := clt.getBufferFromPool()
	clt.putBufferToPool(buf)

	return clt
}()

// NOTE: 这个测试的结果如果很小则说明 pool 也是正常工作
func BenchmarkGetBufferFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			buf := testClient.getBufferFromPool()
			defer testClient.putBufferToPool(buf)
		}()
	}
}

func BenchmarkGetBufferFromNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			_ = bytes.NewBuffer(make([]byte, 2<<20)) // 默认 2MB
		}()
	}
}
