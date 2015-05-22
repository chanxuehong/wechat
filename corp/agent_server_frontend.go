// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package corp

import (
	"net/http"
	"net/url"
)

// 实现了 http.Handler, 处理一个企业号应用的消息(事件)请求.
type AgentServerFrontend struct {
	agentServer           AgentServer
	invalidRequestHandler InvalidRequestHandler
	interceptor           Interceptor
}

// handler, interceptor 均可以为 nil
func NewAgentServerFrontend(srv AgentServer, handler InvalidRequestHandler, interceptor Interceptor) *AgentServerFrontend {
	if srv == nil {
		panic("nil AgentServer")
	}
	if handler == nil {
		handler = DefaultInvalidRequestHandler
	}

	return &AgentServerFrontend{
		agentServer:           srv,
		invalidRequestHandler: handler,
		interceptor:           interceptor,
	}
}

// 实现 http.Handler.
func (frontend *AgentServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		frontend.invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if interceptor := frontend.interceptor; interceptor != nil && !interceptor.Intercept(w, r, queryValues) {
		return
	}

	ServeHTTP(w, r, queryValues, frontend.agentServer, frontend.invalidRequestHandler)
}
