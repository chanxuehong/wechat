// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

// 定义回调 URL 上指定 Agent 的查询参数名
const URLQueryAgentKeyName = "agentkey"

// 多个 Agent 的前端, 负责处理 http 请求, net/http.Handler 的实现
//
//  NOTE:
//  MultiAgentFrontend 可以处理多个公众号的消息（事件），但是要求在回调 URL 上加上一个查询
//  参数，一般为 agentkey（参考常量 URLQueryAgentKeyName），这个参数的值就是 MultiAgentFrontend
//  索引 Agent 的 key。
//  例如回调 URL 为 http://www.xxx.com/weixin?agentkey=1234567890，那么就可以在后端调用
//
//    MultiAgentFrontend.SetAgent("1234567890", agent)
//
//  来增加一个 Agent 来处理 agentkey=1234567890 的消息（事件）。
//
//  MultiAgentFrontend 并发安全，可以在运行中动态增加和删除 Agent。
type MultiAgentFrontend struct {
	rwmutex               sync.RWMutex
	agentMap              map[string]Agent
	invalidRequestHandler InvalidRequestHandler
}

// 设置 InvalidRequestHandler, 如果 handler == nil 则使用默认的 DefaultInvalidRequestHandler
func (this *MultiAgentFrontend) SetInvalidRequestHandler(handler InvalidRequestHandler) {
	this.rwmutex.Lock()
	if handler == nil {
		this.invalidRequestHandler = DefaultInvalidRequestHandler
	} else {
		this.invalidRequestHandler = handler
	}
	this.rwmutex.Unlock()
}

// 添加（设置） agentkey-agent pair, 如果 agent == nil 则不做任何操作
func (this *MultiAgentFrontend) SetAgent(agentkey string, agent Agent) {
	if agentkey == "" {
		return
	}
	if agent == nil {
		return
	}

	this.rwmutex.Lock()
	if this.agentMap == nil {
		this.agentMap = make(map[string]Agent)
	}
	this.agentMap[agentkey] = agent
	this.rwmutex.Unlock()
}

// 删除 agentkey 对应的 Agent
func (this *MultiAgentFrontend) DeleteAgent(agentkey string) {
	this.rwmutex.Lock()
	delete(this.agentMap, agentkey)
	this.rwmutex.Unlock()
}

// 删除所有的 Agent
func (this *MultiAgentFrontend) DeleteAllAgent() {
	this.rwmutex.Lock()
	this.agentMap = nil
	this.rwmutex.Unlock()
}

func (this *MultiAgentFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL == nil {
		this.rwmutex.RLock()
		invalidRequestHandler := this.invalidRequestHandler
		this.rwmutex.RUnlock()
		if invalidRequestHandler == nil {
			invalidRequestHandler = DefaultInvalidRequestHandler
		}

		err := errors.New("input net/http.Request.URL == nil")
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	urlValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		this.rwmutex.RLock()
		invalidRequestHandler := this.invalidRequestHandler
		this.rwmutex.RUnlock()
		if invalidRequestHandler == nil {
			invalidRequestHandler = DefaultInvalidRequestHandler
		}

		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	agentKey := urlValues.Get(URLQueryAgentKeyName)
	if agentKey == "" {
		this.rwmutex.RLock()
		invalidRequestHandler := this.invalidRequestHandler
		this.rwmutex.RUnlock()
		if invalidRequestHandler == nil {
			invalidRequestHandler = DefaultInvalidRequestHandler
		}

		err = fmt.Errorf("the url query value with name %s is empty", URLQueryAgentKeyName)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	this.rwmutex.RLock()
	invalidRequestHandler := this.invalidRequestHandler
	agent := this.agentMap[agentKey]
	this.rwmutex.RUnlock()

	if invalidRequestHandler == nil {
		invalidRequestHandler = DefaultInvalidRequestHandler
	}
	if agent == nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("Not found Agent for %s == %s", URLQueryAgentKeyName, agentKey))
		return
	}

	ServeHTTP(w, r, urlValues, agent, invalidRequestHandler)
}
