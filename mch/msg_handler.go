// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mch

import (
	"io"
	"net/http"
)

// 微信服务器推送过来的消息(事件)处理接口
type MessageHandler interface {
	ServeMessage(w http.ResponseWriter, r *Request)
}

type MessageHandlerFunc func(http.ResponseWriter, *Request)

func (fn MessageHandlerFunc) ServeMessage(w http.ResponseWriter, r *Request) {
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

	RawMsgXML []byte            // 消息的 XML 文本
	Msg       map[string]string // 解析后的消息
}
