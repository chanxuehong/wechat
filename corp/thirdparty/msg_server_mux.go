// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package thirdparty

import (
	"io"
	"net/http"
	"sync"
)

var _ SuiteMessageHandler = (*SuiteMessageServeMux)(nil)

// SuiteMessageServeMux 实现了一个简单的消息路由器, 同时也是一个 SuiteMessageHandler.
type SuiteMessageServeMux struct {
	rwmutex                    sync.RWMutex
	messageHandlers            map[string]SuiteMessageHandler
	defaultSuiteMessageHandler SuiteMessageHandler
}

func NewSuiteMessageServeMux() *SuiteMessageServeMux {
	return &SuiteMessageServeMux{
		messageHandlers: make(map[string]SuiteMessageHandler),
	}
}

// 注册 SuiteMessageHandler, 处理特定类型的消息.
func (mux *SuiteMessageServeMux) MessageHandle(msgType string, handler SuiteMessageHandler) {
	if msgType == "" {
		panic("empty msgType")
	}
	if handler == nil {
		panic("nil SuiteMessageHandler")
	}

	mux.rwmutex.Lock()
	if mux.messageHandlers == nil {
		mux.messageHandlers = make(map[string]SuiteMessageHandler)
	}
	mux.messageHandlers[msgType] = handler
	mux.rwmutex.Unlock()
}

// 注册 SuiteMessageHandlerFunc, 处理特定类型的消息.
func (mux *SuiteMessageServeMux) MessageHandleFunc(msgType string, handler func(http.ResponseWriter, *Request)) {
	mux.MessageHandle(msgType, SuiteMessageHandlerFunc(handler))
}

// 注册 SuiteMessageHandler, 处理未知类型的消息.
func (mux *SuiteMessageServeMux) DefaultMessageHandle(handler SuiteMessageHandler) {
	if handler == nil {
		panic("nil SuiteMessageHandler")
	}

	mux.rwmutex.Lock()
	mux.defaultSuiteMessageHandler = handler
	mux.rwmutex.Unlock()
}

// 注册 SuiteMessageHandlerFunc, 处理未知类型的消息.
func (mux *SuiteMessageServeMux) DefaultMessageHandleFunc(handler func(http.ResponseWriter, *Request)) {
	mux.DefaultMessageHandle(SuiteMessageHandlerFunc(handler))
}

// 获取 msgType 对应的 SuiteMessageHandler, 如果没有找到 nil.
func (mux *SuiteMessageServeMux) messageHandler(msgType string) (handler SuiteMessageHandler) {
	if msgType == "" {
		return nil
	}

	mux.rwmutex.RLock()
	handler = mux.messageHandlers[msgType]
	if handler == nil {
		handler = mux.defaultSuiteMessageHandler
	}
	mux.rwmutex.RUnlock()
	return
}

// SuiteMessageServeMux 实现了 SuiteMessageHandler 接口.
func (mux *SuiteMessageServeMux) ServeMessage(w http.ResponseWriter, r *Request) {
	handler := mux.messageHandler(r.MixedMsg.InfoType)
	if handler == nil {
		io.WriteString(w, "success")
		return
	}
	handler.ServeMessage(w, r)
}
