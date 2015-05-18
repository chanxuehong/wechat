// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package thirdparty

import (
	"net/http"
	"net/url"

	"github.com/chanxuehong/wechat/corp"
)

// 实现了 http.Handler.
type SuiteServerFrontend struct {
	suiteServer           SuiteServer
	invalidRequestHandler corp.InvalidRequestHandler
}

func NewSuiteServerFrontend(server SuiteServer, handler corp.InvalidRequestHandler) *SuiteServerFrontend {
	if server == nil {
		panic("nil SuiteServer")
	}
	if handler == nil {
		handler = corp.DefaultInvalidRequestHandler
	}

	return &SuiteServerFrontend{
		suiteServer:           server,
		invalidRequestHandler: handler,
	}
}

// 实现 http.Handler.
func (frontend *SuiteServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	suiteServer := frontend.suiteServer
	invalidRequestHandler := frontend.invalidRequestHandler

	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, queryValues, suiteServer, invalidRequestHandler)
}
