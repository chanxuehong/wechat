package core

import (
	"github.com/bububa/wechat/mp/core"
	"net/http"
	"net/url"
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

	MsgCiphertext []byte         // 消息的密文文本
	MsgPlaintext  []byte         // 消息的明文文本, xml格式
	MixedMsg      *core.MixedMsg // 消息

	Token  string // 当前消息所属公众号的 Token
	AESKey []byte // 当前消息加密所用的 aes-key, read-only!!!
	Random []byte // 当前消息加密所用的 random, 16-bytes
	AppId  string // 当前消息加密所用的 AppId
}
