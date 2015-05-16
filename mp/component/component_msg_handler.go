// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"io"
	"net/http"
	"net/url"
)

// 微信服务器推送过来的消息(事件)处理接口
type ComponentMessageHandler interface {
	ServeComponentMessage(w http.ResponseWriter, r *Request)
}

type ComponentMessageHandlerFunc func(http.ResponseWriter, *Request)

func (fn ComponentMessageHandlerFunc) ServeComponentMessage(w http.ResponseWriter, r *Request) {
	fn(w, r)
}

type httpResponseWriter struct {
	io.Writer
}

func (httpResponseWriter) Header() http.Header {
	return make(map[string][]string)
}
func (httpResponseWriter) WriteHeader(int) {}

// 将 io.Writer 从语义上实现 http.ResponseWriter.
//  某些 http 框架可能没有提供 http.ResponseWriter, 而只是提供了 io.Writer.
func HttpResponseWriter(w io.Writer) http.ResponseWriter {
	if rw, ok := w.(http.ResponseWriter); ok {
		return rw
	}
	return httpResponseWriter{Writer: w}
}

// 消息(事件)请求信息
type Request struct {
	HttpRequest *http.Request // 可以为 nil, 因为某些 http 框架没有提供此参数

	// 下面的字段必须提供

	QueryValues  url.Values // 回调请求 URL 中的查询参数集合
	MsgSignature string     // 请求 URL 中的消息体签名: msg_signature
	EncryptType  string     // 请求 URL 中的加密方式: encrypt_type
	TimeStamp    int64      // 回调请求 URL 中的时间戳: timestamp
	Nonce        string     // 回调请求 URL 中的随机数: nonce

	RawMsgXML []byte                 // 消息的"明文"XML 文本
	MixedMsg  *MixedComponentMessage // RawMsgXML 解析后的消息

	AESKey [32]byte // 当前消息 AES 加密的 key
	Random []byte   // 当前消息加密时所用的 random, 16 bytes

	ComponentAppId string // 请求消息所属第三方平台的 ID
	ComponentToken string // 请求消息所属第三方平台的 Token
}

// 微信服务器推送过来的消息(事件)的合集.
type MixedComponentMessage struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ComponentAppId string `xml:"AppId"      json:"AppId"`
	CreateTime     int64  `xml:"CreateTime" json:"CreateTime"`
	InfoType       string `xml:"InfoType"   json:"InfoType"`

	ComponentVerifyTicket string `xml:"ComponentVerifyTicket" json:"ComponentVerifyTicket"`
	AuthorizerAppid       string `xml:"AuthorizerAppid"       json:"AuthorizerAppid"`
}
