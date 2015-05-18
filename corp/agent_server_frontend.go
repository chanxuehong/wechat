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
}

func NewAgentServerFrontend(server AgentServer, handler InvalidRequestHandler) *AgentServerFrontend {
	if server == nil {
		panic("nil AgentServer")
	}
	if handler == nil {
		handler = DefaultInvalidRequestHandler
	}

	return &AgentServerFrontend{
		agentServer:           server,
		invalidRequestHandler: handler,
	}
}

// 实现 http.Handler.
func (frontend *AgentServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agentServer := frontend.agentServer
	invalidRequestHandler := frontend.invalidRequestHandler

	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, queryValues, agentServer, invalidRequestHandler)
}
