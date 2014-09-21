// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"bytes"
	"github.com/chanxuehong/wechat/message/passive/request"
	"testing"
)

// 下面两个性能基准测试是看看使用缓存是否能提高性能

// 使用缓存
func BenchmarkGetBufferUnitFromPool(b *testing.B) {
	var buf1 *bytes.Buffer
	var buf2 []byte
	var x int

	for i := 0; i < b.N; i++ {
		func() {
			unit := testHandler.getBufferUnitFromPool()
			defer testHandler.putBufferUnitToPool(unit)

			buf1 = unit.msgBuf
			buf2 = unit.signatureBuf[:]
			x = unit.msgRequest.ErrorCount
		}()
	}
}

// 不使用缓存
func BenchmarkGetBufferUnitFromNew(b *testing.B) {
	var buf1 *bytes.Buffer
	var buf2 []byte
	var x int

	for i := 0; i < b.N; i++ {
		func() {
			var req request.Request

			buf1 = bytes.NewBuffer(make([]byte, 512)) // Handler 平均收到消息的估计大小
			buf2 = make([]byte, 128)                  // checkSignature 内申请切片平均估计大小
			x = req.ErrorCount
		}()
	}
}
