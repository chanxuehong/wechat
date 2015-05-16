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
type ComponentServerFrontend struct {
	componentServer       ComponentServer
	invalidRequestHandler mp.InvalidRequestHandler
}

func NewComponentServerFrontend(server ComponentServer, handler mp.InvalidRequestHandler) *ComponentServerFrontend {
	if server == nil {
		panic("nil ComponentServer")
	}
	if handler == nil {
		handler = mp.DefaultInvalidRequestHandler
	}

	return &ComponentServerFrontend{
		componentServer:       server,
		invalidRequestHandler: handler,
	}
}

// 实现 http.Handler.
func (frontend *ComponentServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	componentServer := frontend.componentServer
	invalidRequestHandler := frontend.invalidRequestHandler

	urlValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, urlValues, componentServer, invalidRequestHandler)
}
