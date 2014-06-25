package server

import (
	"github.com/chanxuehong/wechat/message/request"
)

// 对于 server, 处理 request 消息每次都会用到一些中间变量, 可以缓存起来
type bufferUnit struct {
	buf        []byte          // 主要用于数字签名的验证, len(buf) == 256
	msgRequest request.Request // 用于 xml Decode
}

func newBufferUnit() interface{} {
	return &bufferUnit{
		buf: make([]byte, 256),
	}
}

func (s *Server) getBufferUnitFromPool() *bufferUnit {
	unit := s.bufferUnitPool.Get().(*bufferUnit)
	unit.msgRequest.Zero() // important!
	return unit
}

func (s *Server) putBufferUnitToPool(bufferUnit *bufferUnit) {
	s.bufferUnitPool.Put(bufferUnit)
}
