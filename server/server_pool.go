// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"bytes"
	"github.com/chanxuehong/wechat/message/request"
)

// 对于 server, 处理 request 消息每次都会用到一些中间变量, 可以缓存起来
type bufferUnit struct {
	signatureBuf [256]byte       // 主要用于数字签名 checkSignature 的参数
	msgBuf       *bytes.Buffer   // 缓存读取的消息体
	msgRequest   request.Request // 用于 xml Decode
}

func newBufferUnit() interface{} {
	return &bufferUnit{
		msgBuf: bytes.NewBuffer(make([]byte, 1024)),
	}
}

func (s *Server) getBufferUnitFromPool() (unit *bufferUnit) {
	unit = s.bufferUnitPool.Get().(*bufferUnit)
	unit.msgBuf.Reset()    // important!
	unit.msgRequest.Zero() // important!
	return
}

func (s *Server) putBufferUnitToPool(unit *bufferUnit) {
	s.bufferUnitPool.Put(unit)
}
