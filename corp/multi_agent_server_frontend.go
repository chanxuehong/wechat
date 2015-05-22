// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package corp

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

// 回调 URL 上索引 AgentServer 的 key 的名称.
//  比如下面的回调地址里面就可以根据 agent1 来索引对应的 AgentServer.
//  http://www.xxx.com/?agent_server=agent1&msg_signature=XXX&timestamp=123456789&nonce=12345678
const URLQueryAgentServerKeyName = "agent_server"

// 多个 AgentServer 的前端, 负责处理 http 请求, net/http.Handler 的实现
//
//  NOTE:
//  MultiAgentServerFrontend 可以处理多个企业号应用的消息（事件），但是要求在回调 URL 上加上一个
//  查询参数，参考常量 URLQueryAgentServerKeyName，这个参数的值就是 MultiAgentServerFrontend
//  索引 AgentServer 的 key。
//
//  例如回调 URL 为 http://www.xxx.com/weixin?agent_server=1234567890，那么就可以在后端调用
//
//    MultiAgentServerFrontend.SetAgentServer("1234567890", AgentServer)
//
//  来增加一个 AgentServer 来处理 agent_server=1234567890 的消息（事件）。
//
//  MultiAgentServerFrontend 并发安全，可以在运行中动态增加和删除 AgentServer。
type MultiAgentServerFrontend struct {
	invalidRequestHandler InvalidRequestHandler
	interceptor           Interceptor

	rwmutex        sync.RWMutex
	agentServerMap map[string]AgentServer
}

// handler, interceptor 均可以为 nil
func NewMultiAgentServerFrontend(handler InvalidRequestHandler, interceptor Interceptor) *MultiAgentServerFrontend {
	if handler == nil {
		handler = DefaultInvalidRequestHandler
	}

	return &MultiAgentServerFrontend{
		invalidRequestHandler: handler,
		interceptor:           interceptor,
		agentServerMap:        make(map[string]AgentServer),
	}
}

// 设置 serverKey-AgentServer pair.
func (frontend *MultiAgentServerFrontend) SetAgentServer(serverKey string, server AgentServer) (err error) {
	if serverKey == "" {
		return errors.New("empty serverKey")
	}
	if server == nil {
		return errors.New("nil AgentServer")
	}

	frontend.rwmutex.Lock()
	frontend.agentServerMap[serverKey] = server
	frontend.rwmutex.Unlock()
	return
}

// 删除 serverKey 对应的 AgentServer
func (frontend *MultiAgentServerFrontend) DeleteAgentServer(serverKey string) {
	frontend.rwmutex.Lock()
	delete(frontend.agentServerMap, serverKey)
	frontend.rwmutex.Unlock()
}

// 删除所有的 AgentServer
func (frontend *MultiAgentServerFrontend) DeleteAllAgentServer() {
	frontend.rwmutex.Lock()
	frontend.agentServerMap = make(map[string]AgentServer)
	frontend.rwmutex.Unlock()
}

// 实现 http.Handler
func (frontend *MultiAgentServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		frontend.invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if interceptor := frontend.interceptor; interceptor != nil && !interceptor.Intercept(w, r, queryValues) {
		return
	}

	serverKey := queryValues.Get(URLQueryAgentServerKeyName)
	if serverKey == "" {
		err = fmt.Errorf("the url query value with name %s is empty", URLQueryAgentServerKeyName)
		frontend.invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	frontend.rwmutex.RLock()
	agentServer := frontend.agentServerMap[serverKey]
	frontend.rwmutex.RUnlock()

	if agentServer == nil {
		err = fmt.Errorf("Not found AgentServer for %s == %s", URLQueryAgentServerKeyName, serverKey)
		frontend.invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, queryValues, agentServer, frontend.invalidRequestHandler)
}
