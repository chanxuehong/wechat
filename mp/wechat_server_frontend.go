// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"net/http"
	"net/url"
)

// 实现了 http.Handler, 处理一个公众号的消息(事件)请求.
type WechatServerFrontend struct {
	wechatServer          WechatServer
	invalidRequestHandler InvalidRequestHandler
}

func NewWechatServerFrontend(server WechatServer, handler InvalidRequestHandler) *WechatServerFrontend {
	if server == nil {
		panic("nil WechatServer")
	}
	if handler == nil {
		handler = DefaultInvalidRequestHandler
	}

	return &WechatServerFrontend{
		wechatServer:          server,
		invalidRequestHandler: handler,
	}
}

// 实现 http.Handler.
func (frontend *WechatServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wechatServer := frontend.wechatServer
	invalidRequestHandler := frontend.invalidRequestHandler

	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, queryValues, wechatServer, invalidRequestHandler)
}
