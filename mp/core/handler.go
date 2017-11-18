package core

import (
	"gopkg.in/chanxuehong/wechat.v2/internal/util"
)

const maxHandlerChainSize = 64

type HandlerChain []Handler

type Handler interface {
	ServeMsg(*Context)
}

// HandlerFunc =========================================================================================================

var _ Handler = HandlerFunc(nil)

type HandlerFunc func(*Context)

func (fn HandlerFunc) ServeMsg(ctx *Context) { fn(ctx) }

// ServeMux ============================================================================================================

var _ Handler = (*ServeMux)(nil)

// ServeMux 是一个消息(事件)路由器, 同时也是一个 Handler 的实现.
//  NOTE: ServeMux 非并发安全, 如果需要并发安全的 Handler, 可以参考 ServeMux 实现一个.
type ServeMux struct {
	startedChecker startedChecker

	msgMiddlewares   HandlerChain
	eventMiddlewares HandlerChain

	defaultMsgHandlerChain   HandlerChain
	defaultEventHandlerChain HandlerChain

	msgHandlerChainMap   map[MsgType]HandlerChain
	eventHandlerChainMap map[EventType]HandlerChain
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		msgHandlerChainMap:   make(map[MsgType]HandlerChain),
		eventHandlerChainMap: make(map[EventType]HandlerChain),
	}
}

var successResponseBytes = []byte("success")

// ServeMsg 实现 Handler 接口.
func (mux *ServeMux) ServeMsg(ctx *Context) {
	mux.startedChecker.start()
	if MsgType := ctx.MixedMsg.MsgType; MsgType != "event" {
		handlers := mux.getMsgHandlerChain(MsgType)
		if len(handlers) == 0 {
			ctx.ResponseWriter.Write(successResponseBytes)
			return
		}
		ctx.handlers = handlers
		ctx.Next()
	} else {
		handlers := mux.getEventHandlerChain(ctx.MixedMsg.EventType)
		if len(handlers) == 0 {
			ctx.ResponseWriter.Write(successResponseBytes)
			return
		}
		ctx.handlers = handlers
		ctx.Next()
	}
}

// getMsgHandlerChain 获取 HandlerChain 以处理消息类型为 MsgType 的消息, 如果没有找到返回 nil.
func (mux *ServeMux) getMsgHandlerChain(msgType MsgType) (handlers HandlerChain) {
	if m := mux.msgHandlerChainMap; len(m) > 0 {
		handlers = m[MsgType(util.ToLower(string(msgType)))]
		if len(handlers) == 0 {
			handlers = mux.defaultMsgHandlerChain
		}
	} else {
		handlers = mux.defaultMsgHandlerChain
	}
	return
}

// getEventHandlerChain 获取 HandlerChain 以处理事件类型为 EventType 的事件, 如果没有找到返回 nil.
func (mux *ServeMux) getEventHandlerChain(eventType EventType) (handlers HandlerChain) {
	if m := mux.eventHandlerChainMap; len(m) > 0 {
		handlers = m[EventType(util.ToLower(string(eventType)))]
		if len(handlers) == 0 {
			handlers = mux.defaultEventHandlerChain
		}
	} else {
		handlers = mux.defaultEventHandlerChain
	}
	return
}

// ServeMux: registers HandlerChain ====================================================================================

// Use 注册(新增) middlewares 使其在所有消息(事件)的 Handler 之前处理该处理消息(事件).
func (mux *ServeMux) Use(middlewares ...Handler) {
	mux.startedChecker.check()
	if len(middlewares) == 0 {
		return
	}
	for _, h := range middlewares {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.useForMsg(middlewares)
	mux.useForEvent(middlewares)
}

// UseFunc 注册(新增) middlewares 使其在所有消息(事件)的 Handler 之前处理该处理消息(事件).
func (mux *ServeMux) UseFunc(middlewares ...func(*Context)) {
	mux.startedChecker.check()
	if len(middlewares) == 0 {
		return
	}
	for _, h := range middlewares {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	middlewares2 := make(HandlerChain, len(middlewares))
	for i := 0; i < len(middlewares); i++ {
		middlewares2[i] = HandlerFunc(middlewares[i])
	}
	mux.useForMsg(middlewares2)
	mux.useForEvent(middlewares2)
}

// UseForMsg 注册(新增) middlewares 使其在所有消息的 Handler 之前处理该处理消息.
func (mux *ServeMux) UseForMsg(middlewares ...Handler) {
	mux.startedChecker.check()
	if len(middlewares) == 0 {
		return
	}
	for _, h := range middlewares {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.useForMsg(middlewares)
}

// UseFuncForMsg 注册(新增) middlewares 使其在所有消息的 Handler 之前处理该处理消息.
func (mux *ServeMux) UseFuncForMsg(middlewares ...func(*Context)) {
	mux.startedChecker.check()
	if len(middlewares) == 0 {
		return
	}
	for _, h := range middlewares {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	middlewares2 := make(HandlerChain, len(middlewares))
	for i := 0; i < len(middlewares); i++ {
		middlewares2[i] = HandlerFunc(middlewares[i])
	}
	mux.useForMsg(middlewares2)
}

func (mux *ServeMux) useForMsg(middlewares []Handler) {
	if len(mux.defaultMsgHandlerChain) > 0 || len(mux.msgHandlerChainMap) > 0 {
		panic("please call this method before any other methods those registered handlers for message")
	}
	mux.msgMiddlewares = combineHandlerChain(mux.msgMiddlewares, middlewares)
}

// UseForEvent 注册(新增) middlewares 使其在所有事件的 Handler 之前处理该处理事件.
func (mux *ServeMux) UseForEvent(middlewares ...Handler) {
	mux.startedChecker.check()
	if len(middlewares) == 0 {
		return
	}
	for _, h := range middlewares {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.useForEvent(middlewares)
}

// UseFuncForEvent 注册(新增) middlewares 使其在所有事件的 Handler 之前处理该处理事件.
func (mux *ServeMux) UseFuncForEvent(middlewares ...func(*Context)) {
	mux.startedChecker.check()
	if len(middlewares) == 0 {
		return
	}
	for _, h := range middlewares {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	middlewares2 := make(HandlerChain, len(middlewares))
	for i := 0; i < len(middlewares); i++ {
		middlewares2[i] = HandlerFunc(middlewares[i])
	}
	mux.useForEvent(middlewares2)
}

func (mux *ServeMux) useForEvent(middlewares []Handler) {
	if len(mux.defaultEventHandlerChain) > 0 || len(mux.eventHandlerChainMap) > 0 {
		panic("please call this method before any other methods those registered handlers for event")
	}
	mux.eventMiddlewares = combineHandlerChain(mux.eventMiddlewares, middlewares)
}

// DefaultMsgHandle 设置 handlers 以处理没有匹配到具体类型的 HandlerChain 的消息.
func (mux *ServeMux) DefaultMsgHandle(handlers ...Handler) {
	mux.startedChecker.check()
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.defaultMsgHandlerChain = combineHandlerChain(mux.msgMiddlewares, handlers)
}

// DefaultMsgHandleFunc 设置 handlers 以处理没有匹配到具体类型的 HandlerChain 的消息.
func (mux *ServeMux) DefaultMsgHandleFunc(handlers ...func(*Context)) {
	mux.startedChecker.check()
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	handlers2 := make(HandlerChain, len(handlers))
	for i := 0; i < len(handlers); i++ {
		handlers2[i] = HandlerFunc(handlers[i])
	}
	mux.defaultMsgHandlerChain = combineHandlerChain(mux.msgMiddlewares, handlers2)
}

// DefaultEventHandle 设置 handlers 以处理没有匹配到具体类型的 HandlerChain 的事件.
func (mux *ServeMux) DefaultEventHandle(handlers ...Handler) {
	mux.startedChecker.check()
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.defaultEventHandlerChain = combineHandlerChain(mux.eventMiddlewares, handlers)
}

// DefaultEventHandleFunc 设置 handlers 以处理没有匹配到具体类型的 HandlerChain 的事件.
func (mux *ServeMux) DefaultEventHandleFunc(handlers ...func(*Context)) {
	mux.startedChecker.check()
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	handlers2 := make(HandlerChain, len(handlers))
	for i := 0; i < len(handlers); i++ {
		handlers2[i] = HandlerFunc(handlers[i])
	}
	mux.defaultEventHandlerChain = combineHandlerChain(mux.eventMiddlewares, handlers2)
}

// MsgHandle 设置 handlers 以处理特定类型的消息.
func (mux *ServeMux) MsgHandle(msgType MsgType, handlers ...Handler) {
	mux.startedChecker.check()
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.msgHandlerChainMap[MsgType(util.ToLower(string(msgType)))] = combineHandlerChain(mux.msgMiddlewares, handlers)
}

// MsgHandleFunc 设置 handlers 以处理特定类型的消息.
func (mux *ServeMux) MsgHandleFunc(msgType MsgType, handlers ...func(*Context)) {
	mux.startedChecker.check()
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	handlers2 := make(HandlerChain, len(handlers))
	for i := 0; i < len(handlers); i++ {
		handlers2[i] = HandlerFunc(handlers[i])
	}
	mux.msgHandlerChainMap[MsgType(util.ToLower(string(msgType)))] = combineHandlerChain(mux.msgMiddlewares, handlers2)
}

// EventHandle 设置 handlers 以处理特定类型的事件.
func (mux *ServeMux) EventHandle(eventType EventType, handlers ...Handler) {
	mux.startedChecker.check()
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.eventHandlerChainMap[EventType(util.ToLower(string(eventType)))] = combineHandlerChain(mux.eventMiddlewares, handlers)
}

// EventHandleFunc 设置 handlers 以处理特定类型的事件.
func (mux *ServeMux) EventHandleFunc(eventType EventType, handlers ...func(*Context)) {
	mux.startedChecker.check()
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	handlers2 := make(HandlerChain, len(handlers))
	for i := 0; i < len(handlers); i++ {
		handlers2[i] = HandlerFunc(handlers[i])
	}
	mux.eventHandlerChainMap[EventType(util.ToLower(string(eventType)))] = combineHandlerChain(mux.eventMiddlewares, handlers2)
}

func combineHandlerChain(middlewares, handlers HandlerChain) HandlerChain {
	if len(middlewares)+len(handlers) > maxHandlerChainSize {
		panic("too many handlers")
	}
	return append(middlewares, handlers...)
}
