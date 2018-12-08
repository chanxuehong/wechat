package core

import (
	"bufio"
	"encoding/base64"
	"encoding/xml"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chanxuehong/wechat/internal/debug/callback"
	"github.com/chanxuehong/wechat/internal/util"
)

const (
	initHandlerIndex  = -1
	abortHandlerIndex = maxHandlerChainSize
)

// Context 是 Handler 处理消息(事件)的上下文环境. 非并发安全!
type Context struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request

	QueryParams  url.Values // 回调请求 URL 的查询参数集合
	EncryptType  string     // 回调请求 URL 的加密方式参数: encrypt_type
	MsgSignature string     // 回调请求 URL 的消息体签名参数: msg_signature
	Signature    string     // 回调请求 URL 的签名参数: signature
	Timestamp    int64      // 回调请求 URL 的时间戳参数: timestamp
	Nonce        string     // 回调请求 URL 的随机数参数: nonce

	MsgCiphertext []byte    // 消息的密文文本
	MsgPlaintext  []byte    // 消息的明文文本, xml格式
	MixedMsg      *MixedMsg // 消息

	Token  string // 当前消息所属公众号的 Token
	AESKey []byte // 当前消息加密所用的 aes-key, read-only!!!
	Random []byte // 当前消息加密所用的 random, 16-bytes
	AppId  string // 当前消息加密所用的 AppId

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
		ctx.handlers[ctx.handlerIndex].ServeMsg(ctx)
	}
}

// SetHandlers 设置 handlers 给 Context.Next() 调用, 务必在 Context.Next() 调用之前设置, 否则会 panic.
//  NOTE: 此方法一般用不到, 除非你自己实现一个 Handler 给 Server 使用, 参考 ServeMux.
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

// Context:kvs =========================================================================================================

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

// Context:response ====================================================================================================

// NoneResponse 表示没有消息回复给微信服务器.
func (ctx *Context) NoneResponse() (err error) {
	_, err = ctx.ResponseWriter.Write(successResponseBytes)
	return
}

// RawResponse 回复明文消息给微信服务器.
//  msg: 经过 encoding/xml.Marshal 得到的结果符合微信消息格式的任何数据结构
func (ctx *Context) RawResponse(msg interface{}) (err error) {
	return callback.XmlEncodeResponseMessage(ctx.ResponseWriter, msg)
}

// stringWriter is the interface that wraps the WriteString method.
type stringWriter interface {
	WriteString(s string) (n int, err error)
}

// AESResponse 回复aes加密的消息给微信服务器.
//  msg:       经过 encoding/xml.Marshal 得到的结果符合微信消息格式的任何数据结构
//  timestamp: 时间戳, 如果为 0 则默认使用 Context.Timestamp
//  nonce:     随机数, 如果为 "" 则默认使用 Context.Nonce
//  random:    16字节的随机字符串, 如果为 nil 则默认使用 Context.Random
func (ctx *Context) AESResponse(msg interface{}, timestamp int64, nonce string, random []byte) (err error) {
	if timestamp == 0 {
		timestamp = ctx.Timestamp
	}
	if nonce == "" {
		nonce = ctx.Nonce
	}
	if len(random) == 0 {
		random = ctx.Random
	}

	msgPlaintext, err := callback.XmlMarshalResponseMessage(msg)
	if err != nil {
		return
	}

	encryptedMsg := util.AESEncryptMsg(random, msgPlaintext, ctx.AppId, ctx.AESKey)
	base64EncryptedMsg := base64.StdEncoding.EncodeToString(encryptedMsg)
	timestampString := strconv.FormatInt(timestamp, 10)
	msgSignature := util.MsgSign(ctx.Token, timestampString, nonce, base64EncryptedMsg)

	if sw, ok := ctx.ResponseWriter.(stringWriter); ok {
		if _, err = sw.WriteString("<xml><Encrypt>"); err != nil {
			return
		}
		if _, err = sw.WriteString(base64EncryptedMsg); err != nil {
			return
		}
		if _, err = sw.WriteString("</Encrypt><MsgSignature>"); err != nil {
			return
		}
		if _, err = sw.WriteString(msgSignature); err != nil {
			return
		}
		if _, err = sw.WriteString("</MsgSignature><TimeStamp>"); err != nil {
			return
		}
		if _, err = sw.WriteString(timestampString); err != nil {
			return
		}
		if _, err = sw.WriteString("</TimeStamp><Nonce>"); err != nil {
			return
		}
		if err = xml.EscapeText(ctx.ResponseWriter, []byte(nonce)); err != nil {
			return
		}
		_, err = sw.WriteString("</Nonce></xml>")
		return
	} else {
		bufw := bufio.NewWriterSize(ctx.ResponseWriter, 256)
		if _, err = bufw.WriteString("<xml><Encrypt>"); err != nil {
			return
		}
		if _, err = bufw.WriteString(base64EncryptedMsg); err != nil {
			return
		}
		if _, err = bufw.WriteString("</Encrypt><MsgSignature>"); err != nil {
			return
		}
		if _, err = bufw.WriteString(msgSignature); err != nil {
			return
		}
		if _, err = bufw.WriteString("</MsgSignature><TimeStamp>"); err != nil {
			return
		}
		if _, err = bufw.WriteString(timestampString); err != nil {
			return
		}
		if _, err = bufw.WriteString("</TimeStamp><Nonce>"); err != nil {
			return
		}
		if err = xml.EscapeText(bufw, []byte(nonce)); err != nil {
			return
		}
		if _, err = bufw.WriteString("</Nonce></xml>"); err != nil {
			return
		}
		return bufw.Flush()
	}
}
