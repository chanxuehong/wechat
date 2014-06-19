package wechat

import (
	"github.com/chanxuehong/wechat/message"
)

func serverNewMessageRequest() interface{} {
	return new(message.Request)
}

func (s *Server) getMessageRequestFromPool() *message.Request {
	msg := s.messageRequestPool.Get().(*message.Request)
	return msg.Zero() // important!
}

func (s *Server) putMessageRequestToPool(msg *message.Request) {
	s.messageRequestPool.Put(msg)
}
