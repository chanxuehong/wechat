package wechat

import (
	"github.com/chanxuehong/wechat/message"
)

func newMessageRequest() interface{} {
	return new(message.Request)
}

func (s *Server) getRequestFromPool() *message.Request {
	msg := s.messageRequestPool.Get().(*message.Request)
	return msg.Zero() // important!
}

func (s *Server) putRequestToPool(msg *message.Request) {
	s.messageRequestPool.Put(msg)
}
