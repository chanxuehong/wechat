// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package chat

import (
	"errors"
	"net/http"
	"net/url"
	"sync"
)

// 多个 AgentServer 的前端, http.Handler 的实现.
//
//  MultiAgentServerFrontend 可以处理多个企业号应用的消息(事件), 但是要求在回调 URL 上加上一个
//  查询参数(参数名与 urlAgentServerQueryName 一致), 通过这个参数的值来索引对应的 AgentServer.
//
//  例如回调 URL 为(urlAgentServerQueryName == "agent_server"):
//    http://www.xxx.com/weixin?agent_server=1234567890
//  那么就可以在后端调用
//   MultiAgentServerFrontend.SetAgentServer("1234567890", AgentServer)
//  来增加一个 AgentServer 来处理 agent_server=1234567890 的消息(事件).
//
//  MultiAgentServerFrontend 并发安全, 可以在运行中动态增加和删除 AgentServer.
type MultiChatServerFrontend struct {
	errHandler     ErrorHandler
	interceptor    Interceptor

	rwmutex        sync.RWMutex
	chatServerMap map[string]ChatServer
}

// NewMultiAgentServerFrontend 创建一个新的 MultiAgentServerFrontend.
//  urlAgentServerQueryName: 回调 URL 上参数名, 这个参数的值就是索引 AgentServer 的 key
//  errHandler:              错误处理 handler, 可以为 nil
//  interceptor:             拦截器, 可以为 nil
func NewMultiChatServerFrontend( errHandler ErrorHandler, interceptor Interceptor) *MultiChatServerFrontend {

	if errHandler == nil {
		errHandler = DefaultErrorHandler
	}

	return &MultiChatServerFrontend{
		errHandler:              errHandler,
		interceptor:             interceptor,
		chatServerMap:          make(map[string]ChatServer),
	}
}

func (frontend *MultiChatServerFrontend) SetChatServer(serverKey string, server ChatServer) (err error) {
	if serverKey == "" {
		return errors.New("empty serverKey")
	}
	if server == nil {
		return errors.New("nil ChatServer")
	}

	frontend.rwmutex.Lock()
	frontend.chatServerMap[serverKey] = server
	frontend.rwmutex.Unlock()
	return
}

func (frontend *MultiChatServerFrontend) DeleteChatServer(serverKey string) {
	frontend.rwmutex.Lock()
	delete(frontend.chatServerMap, serverKey)
	frontend.rwmutex.Unlock()
}

func (frontend *MultiChatServerFrontend) DeleteAllChatServer() {
	frontend.rwmutex.Lock()
	frontend.chatServerMap = make(map[string]ChatServer)
	frontend.rwmutex.Unlock()
}

func (frontend *MultiChatServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		frontend.errHandler.ServeError(w, r, err)
		return
	}

	if interceptor := frontend.interceptor; interceptor != nil && !interceptor.Intercept(w, r, queryValues) {
		return
	}


	ServeHTTP(w, r, queryValues, frontend)
}
