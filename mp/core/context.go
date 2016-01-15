package core

import (
	"net/http"
	"net/url"
)

const (
	initHandlerIndex  = -1
	abortHandlerIndex = maxHandlerChainSize
)

type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	QueryParams  url.Values // 回调请求 URL 的查询参数集合
	EncryptType  string     // 回调请求 URL 的加密方式参数: encrypt_type
	MsgSignature string     // 回调请求 URL 的消息体签名参数: msg_signature
	Signature    string     // 回调请求 URL 的签名参数: signature
	Timestamp    int64      // 回调请求 URL 的时间戳参数: timestamp
	Nonce        string     // 回调请求 URL 的随机数参数: nonce

	MsgPlaintext []byte    // 消息的xml明文文本
	MixedMsg     *MixedMsg // 消息

	Token  string // 当前消息所属公众号的 Token
	AESKey []byte // 当前消息加密所用的 aes-key, read-only!!!
	Random []byte // 当前消息加密所用的 random, 16-bytes
	AppId  string // 当前消息加密所用的 AppId

	handlerIndex int
	handlers     HandlerChain

	kvs map[string]interface{}
}

func (ctx *Context) IsAborted() bool {
	return ctx.handlerIndex >= abortHandlerIndex
}

func (ctx *Context) Abort() {
	ctx.handlerIndex = abortHandlerIndex
}

func (ctx *Context) SetHandlers(handlers HandlerChain) {
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

func (ctx *Context) Next() {
	for ctx.handlerIndex++; ctx.handlerIndex < len(ctx.handlers); ctx.handlerIndex++ {
		ctx.handlers[ctx.handlerIndex].ServeMsg(ctx)
	}
}

// ================================== kvs ======================================

// Set is used to store a new key/value pair exclusivelly for this context.
func (ctx *Context) Set(key string, value interface{}) {
	if ctx.kvs == nil {
		ctx.kvs = make(map[string]interface{})
	}
	ctx.kvs[key] = value
}

// Get returns the value for the given key, ie: (value, true).
// If the value does not exists it returns (nil, false)
func (ctx *Context) Get(key string) (value interface{}, exists bool) {
	value, exists = ctx.kvs[key]
	return
}

// MustGet Returns the value for the given key if it exists, otherwise it panics.
func (ctx *Context) MustGet(key string) interface{} {
	if value, exists := ctx.Get(key); exists {
		return value
	}
	panic(`[kvs] key "` + key + `" does not exist`)
}
