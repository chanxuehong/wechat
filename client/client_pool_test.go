// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"testing"
)

var _test_client_pool_client = func() *Client {
	clt := NewClient("appid", "secret", nil)

	// 预热
	buf := clt.getBufferFromPool()
	clt.putBufferToPool(buf)

	return clt
}()

func BenchmarkGetBufferFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			buf := _test_client_pool_client.getBufferFromPool()
			defer _test_client_pool_client.putBufferToPool(buf)
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
