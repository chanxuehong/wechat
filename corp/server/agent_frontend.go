// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"errors"
	"net/http"
	"net/url"
)

// Agent 的前端, 负责处理 http 请求, net/http.Handler 的实现
//  NOTE: 只能处理一个企业号应用的消息
type AgentFrontend struct {
	agent                 Agent
	invalidRequestHandler InvalidRequestHandler
}

// 创建一个新的 AgentFrontend.
//  agent 不能为 nil, 如果 invalidRequestHandler == nil 则使用 DefaultInvalidRequestHandler
func NewAgentFrontend(agent Agent, invalidRequestHandler InvalidRequestHandler) *AgentFrontend {
	if agent == nil {
		panic("agent == nil")
	}
	if invalidRequestHandler == nil {
		invalidRequestHandler = DefaultInvalidRequestHandler
	}

	return &AgentFrontend{
		agent: agent,
		invalidRequestHandler: invalidRequestHandler,
	}
}

func (this *AgentFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agent := this.agent
	invalidRequestHandler := this.invalidRequestHandler

	if r.URL == nil {
		err := errors.New("input net/http.Request.URL == nil")
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	urlValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, urlValues, agent, invalidRequestHandler)
}
