// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"io"
	"net/http"
	"sync"
)

var _ ComponentMessageHandler = (*ComponentMessageServeMux)(nil)

// ComponentMessageServeMux 实现了一个简单的消息路由器, 同时也是一个 ComponentMessageHandler.
type ComponentMessageServeMux struct {
	rwmutex                        sync.RWMutex
	componentMessageHandlers       map[string]ComponentMessageHandler
	defaultComponentMessageHandler ComponentMessageHandler
}

func NewComponentMessageServeMux() *ComponentMessageServeMux {
	return &ComponentMessageServeMux{
		componentMessageHandlers: make(map[string]ComponentMessageHandler),
	}
}

// 注册 ComponentMessageHandler, 处理特定类型的消息.
func (mux *ComponentMessageServeMux) MessageHandle(msgType string, handler ComponentMessageHandler) {
	if msgType == "" {
		panic("empty msgType")
	}
	if handler == nil {
		panic("nil ComponentMessageHandler")
	}

	mux.rwmutex.Lock()
	if mux.componentMessageHandlers == nil {
		mux.componentMessageHandlers = make(map[string]ComponentMessageHandler)
	}
	mux.componentMessageHandlers[msgType] = handler
	mux.rwmutex.Unlock()
}

// 注册 ComponentMessageHandlerFunc, 处理特定类型的消息.
func (mux *ComponentMessageServeMux) MessageHandleFunc(msgType string, handler func(http.ResponseWriter, *Request)) {
	mux.MessageHandle(msgType, ComponentMessageHandlerFunc(handler))
}

// 注册 ComponentMessageHandler, 处理未知类型的消息.
func (mux *ComponentMessageServeMux) DefaultMessageHandle(handler ComponentMessageHandler) {
	if handler == nil {
		panic("nil ComponentMessageHandler")
	}

	mux.rwmutex.Lock()
	mux.defaultComponentMessageHandler = handler
	mux.rwmutex.Unlock()
}

// 注册 ComponentMessageHandlerFunc, 处理未知类型的消息.
func (mux *ComponentMessageServeMux) DefaultMessageHandleFunc(handler func(http.ResponseWriter, *Request)) {
	mux.DefaultMessageHandle(ComponentMessageHandlerFunc(handler))
}

// 获取 msgType 对应的 ComponentMessageHandler, 如果没有找到 nil.
func (mux *ComponentMessageServeMux) componentMessageHandler(msgType string) (handler ComponentMessageHandler) {
	if msgType == "" {
		return nil
	}

	mux.rwmutex.RLock()
	handler = mux.componentMessageHandlers[msgType]
	if handler == nil {
		handler = mux.defaultComponentMessageHandler
	}
	mux.rwmutex.RUnlock()
	return
}

// ComponentMessageServeMux 实现了 ComponentMessageHandler 接口.
func (mux *ComponentMessageServeMux) ServeComponentMessage(w http.ResponseWriter, r *Request) {
	handler := mux.componentMessageHandler(r.MixedMsg.InfoType)
	if handler == nil {
		io.WriteString(w, "success")
		return
	}
	handler.ServeComponentMessage(w, r)
}
