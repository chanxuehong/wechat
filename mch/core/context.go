package core

import (
	"net/http"

	"github.com/chanxuehong/util"
)

const (
	initHandlerIndex  = -1
	abortHandlerIndex = maxHandlerChainSize
)

// Context 是 Handler 处理消息(事件)的上下文环境. 非并发安全!
type Context struct {
	Server *Server

	ResponseWriter http.ResponseWriter
	Request        *http.Request

	RequestBody []byte            // 回调请求的 http-body, 就是消息体的原始内容, 记录log可能需要这个信息
	Msg         map[string]string // 请求消息, return_code == "SUCCESS" && result_code == "SUCCESS"

	handlers     HandlerChain
	handlerIndex int

	kvs map[string]interface{}
}

// IsAborted 返回 true 如果 Context.Abort() 被调用了, 否则返回 false.
func (ctx *Context) IsAborted() bool {
	return ctx.handlerIndex >= abortHandlerIndex
}

// Abort 阻止系统调用当前 handler 后续的 handlers, 即当前的 handler 处理完毕就返回, 一般在 middleware 中调用.
func (ctx *Context) Abort() {
	ctx.handlerIndex = abortHandlerIndex
}

// Next 中断当前 handler 程序逻辑执行其后续的 handlers, 一般在 middleware 中调用.
func (ctx *Context) Next() {
	for {
		ctx.handlerIndex++
		if ctx.handlerIndex >= len(ctx.handlers) {
			ctx.handlerIndex--
			break
		}
		handler := ctx.handlers[ctx.handlerIndex]
		if handler != nil {
			handler.ServeMsg(ctx)
		}
	}
}

// SetHandlers 设置 handlers 给 Context.Next() 调用, 务必在 Context.Next() 调用之前设置, 否则会 panic.
//  NOTE: 此方法一般用不到, 除非你自己实现一个 Handler 给 Server 使用, 参考 HandlerChain.
func (ctx *Context) SetHandlers(handlers HandlerChain) {
	if len(handlers) > maxHandlerChainSize {
		panic("too many handlers")
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	if ctx.handlerIndex != initHandlerIndex {
		panic("can't set handlers after Context.Next() called")
	}
	ctx.handlers = handlers
}

// Response 回复消息给微信服务器
func (ctx *Context) Response(msg map[string]string) (err error) {
	return util.EncodeXMLFromMap(ctx.ResponseWriter, msg, "xml")
}

// Set 存储 key-value pair 到 Context 中.
func (ctx *Context) Set(key string, value interface{}) {
	if ctx.kvs == nil {
		ctx.kvs = make(map[string]interface{})
	}
	ctx.kvs[key] = value
}

// Get 返回 Context 中 key 对应的 value, 如果 key 存在的返回 (value, true), 否则返回 (nil, false).
func (ctx *Context) Get(key string) (value interface{}, exists bool) {
	value, exists = ctx.kvs[key]
	return
}

// MustGet 返回 Context 中 key 对应的 value, 如果 key 不存在则会 panic.
func (ctx *Context) MustGet(key string) interface{} {
	if value, exists := ctx.Get(key); exists {
		return value
	}
	panic(`[kvs] key "` + key + `" does not exist`)
}
