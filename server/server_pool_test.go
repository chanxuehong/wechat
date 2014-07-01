// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"bytes"
	"github.com/chanxuehong/wechat/message/request"
	"testing"
)

func BenchmarkGetBufferUnitFromPool(b *testing.B) {
	srv := NewServer(&ServerSetting{})
	// 先分配一个对象
	unit := srv.getBufferUnitFromPool()
	srv.putBufferUnitToPool(unit)

	for i := 0; i < b.N; i++ {
		unit := srv.getBufferUnitFromPool()
		srv.putBufferUnitToPool(unit)
	}
}

func BenchmarkGetBufferUnitFromNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bytes.NewBuffer(make([]byte, 1024))
		_ = new(request.Request)
		_ = make([]byte, 256)
	}
}
