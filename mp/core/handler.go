package core

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
type ServeMux struct {
	msgMiddlewares   HandlerChain
	eventMiddlewares HandlerChain

	defaultMsgHandlerChain   HandlerChain
	defaultEventHandlerChain HandlerChain

	msgHandlerChainMap   map[string]HandlerChain // map[MsgType]HandlerChain
	eventHandlerChainMap map[string]HandlerChain // map[EventType]HandlerChain
}

func NewServeMux() *ServeMux {
	return &ServeMux{
		msgHandlerChainMap:   make(map[string]HandlerChain),
		eventHandlerChainMap: make(map[string]HandlerChain),
	}
}

// ServeMsg 实现 Handler 接口.
func (mux *ServeMux) ServeMsg(ctx *Context) {
	if MsgType := ctx.MixedMsg.MsgType; MsgType != "event" {
		handlers := mux.getMsgHandlerChain(MsgType)
		if len(handlers) == 0 {
			return // 返回空串, 符合微信协议
		}
		ctx.handlers = handlers
		ctx.Next()
	} else {
		handlers := mux.getEventHandlerChain(ctx.MixedMsg.Event)
		if len(handlers) == 0 {
			return // 返回空串, 符合微信协议
		}
		ctx.handlers = handlers
		ctx.Next()
	}
}

// getMsgHandlerChain 获取 HandlerChain 以处理消息类型为 MsgType 的消息, 如果没有找到返回 nil.
func (mux *ServeMux) getMsgHandlerChain(MsgType string) (handlers HandlerChain) {
	if m := mux.msgHandlerChainMap; len(m) > 0 {
		handlers = m[MsgType]
		if len(handlers) == 0 {
			handlers = mux.defaultMsgHandlerChain
		}
	} else {
		handlers = mux.defaultMsgHandlerChain
	}
	return
}

// getEventHandlerChain 获取 HandlerChain 以处理事件类型为 EventType 的事件, 如果没有找到返回 nil.
func (mux *ServeMux) getEventHandlerChain(EventType string) (handlers HandlerChain) {
	if m := mux.eventHandlerChainMap; len(m) > 0 {
		handlers = m[EventType]
		if len(handlers) == 0 {
			handlers = mux.defaultEventHandlerChain
		}
	} else {
		handlers = mux.defaultEventHandlerChain
	}
	return
}

// Use 注册 middlewares 使其在所有消息(事件)的 Handler 之前处理该处理消息(事件).
func (mux *ServeMux) Use(middlewares ...Handler) {
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

// UseForMsg 注册 middlewares 使其在所有消息的 Handler 之前处理该处理消息.
func (mux *ServeMux) UseForMsg(middlewares ...Handler) {
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

func (mux *ServeMux) useForMsg(middlewares []Handler) {
	if len(mux.defaultMsgHandlerChain) > 0 || len(mux.msgHandlerChainMap) > 0 {
		panic("please call this method before any other methods those registered handlers for message")
	}
	mux.msgMiddlewares = combineHandlerChain(mux.msgMiddlewares, middlewares)
}

// UseForEvent 注册 middlewares 使其在所有事件的 Handler 之前处理该处理事件.
func (mux *ServeMux) UseForEvent(middlewares ...Handler) {
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

func (mux *ServeMux) useForEvent(middlewares []Handler) {
	if len(mux.defaultEventHandlerChain) > 0 || len(mux.eventHandlerChainMap) > 0 {
		panic("please call this method before any other methods those registered handlers for event")
	}
	mux.eventMiddlewares = combineHandlerChain(mux.eventMiddlewares, middlewares)
}

// DefaultMsgHandle 注册 handlers 以处理没有匹配到具体类型的 HandlerChain 的消息.
func (mux *ServeMux) DefaultMsgHandle(handlers ...Handler) {
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

// DefaultMsgHandleFunc 注册 handlers 以处理没有匹配到具体类型的 HandlerChain 的消息.
func (mux *ServeMux) DefaultMsgHandleFunc(handlers ...func(*Context)) {
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

// DefaultEventHandle 注册 handlers 以处理没有匹配到具体类型的 HandlerChain 的事件.
func (mux *ServeMux) DefaultEventHandle(handlers ...Handler) {
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

// DefaultEventHandleFunc 注册 handlers 以处理没有匹配到具体类型的 HandlerChain 的事件.
func (mux *ServeMux) DefaultEventHandleFunc(handlers ...func(*Context)) {
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

// MsgHandle 注册 handlers 以处理特定类型的消息.
func (mux *ServeMux) MsgHandle(MsgType string, handlers ...Handler) {
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.msgHandlerChainMap[MsgType] = combineHandlerChain(mux.msgMiddlewares, handlers)
}

// MsgHandleFunc 注册 handlers 以处理特定类型的消息.
func (mux *ServeMux) MsgHandleFunc(MsgType string, handlers ...func(*Context)) {
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
	mux.msgHandlerChainMap[MsgType] = combineHandlerChain(mux.msgMiddlewares, handlers2)
}

// EventHandle 注册 handlers 以处理特定类型的事件.
func (mux *ServeMux) EventHandle(EventType string, handlers ...Handler) {
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	mux.eventHandlerChainMap[EventType] = combineHandlerChain(mux.eventMiddlewares, handlers)
}

// EventHandleFunc 注册 handlers 以处理特定类型的事件.
func (mux *ServeMux) EventHandleFunc(EventType string, handlers ...func(*Context)) {
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
	mux.eventHandlerChainMap[EventType] = combineHandlerChain(mux.eventMiddlewares, handlers2)
}

func combineHandlerChain(middlewares, handlers HandlerChain) HandlerChain {
	size := len(middlewares) + len(handlers)
	if size > maxHandlerChainSize {
		panic("too many handlers")
	}
	combinedHandlerChain := make(HandlerChain, size)
	copy(combinedHandlerChain, middlewares)
	copy(combinedHandlerChain[len(middlewares):], handlers)
	return combinedHandlerChain
}
