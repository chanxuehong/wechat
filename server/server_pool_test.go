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

var _test_server_pool_server = func() *Server {
	srv := NewServer(&ServerSetting{})
	// 先分配一个对象
	unit := srv.getBufferUnitFromPool()
	srv.putBufferUnitToPool(unit)

	return srv
}()

func BenchmarkGetBufferUnitFromPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			unit := _test_server_pool_server.getBufferUnitFromPool()
			defer _test_server_pool_server.putBufferUnitToPool(unit)
		}()
	}
}

func BenchmarkGetBufferUnitFromNew(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			_ = bytes.NewBuffer(make([]byte, 512)) // server 平均收到消息的估计大小
			_ = request.Request{}
			_ = make([]byte, 128) // checkSignature 内申请切片平均估计大小
		}()
	}
}
