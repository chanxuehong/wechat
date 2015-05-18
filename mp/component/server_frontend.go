// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

import (
	"net/http"
	"net/url"

	"github.com/chanxuehong/wechat/mp"
)

// 实现了 http.Handler.
type ServerFrontend struct {
	server                Server
	invalidRequestHandler mp.InvalidRequestHandler
}

func NewServerFrontend(server Server, handler mp.InvalidRequestHandler) *ServerFrontend {
	if server == nil {
		panic("nil Server")
	}
	if handler == nil {
		handler = mp.DefaultInvalidRequestHandler
	}

	return &ServerFrontend{
		server:                server,
		invalidRequestHandler: handler,
	}
}

// 实现 http.Handler.
func (frontend *ServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	Server := frontend.server
	invalidRequestHandler := frontend.invalidRequestHandler

	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, queryValues, Server, invalidRequestHandler)
}
