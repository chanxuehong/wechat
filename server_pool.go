package wechat

import (
	"github.com/chanxuehong/wechat/message"
)

// 对于 server, 处理 request 消息每次都会用到一些中间变量, 可以缓存起来
type serverBufferUnit struct {
	buf        []byte          // 主要用于数字签名的验证, len(buf) == 256
	msgRequest message.Request // 用于 xml Decode
}

func newServerBufferUnit() interface{} {
	return &serverBufferUnit{
		buf: make([]byte, 256),
	}
}

func (s *Server) getBufferUnitFromPool() *serverBufferUnit {
	unit := s.bufferUnitPool.Get().(*serverBufferUnit)
	unit.msgRequest.Zero() // important!
	return unit
}

func (s *Server) putBufferUnitToPool(bufferUnit *serverBufferUnit) {
	s.bufferUnitPool.Put(bufferUnit)
}
