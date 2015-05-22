// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mch

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

// 回调 URL 上索引 Server 的 key 的名称.
//  比如下面的回调地址里面就可以根据 server1 来索引对应的 Server.
//  http://www.xxx.com/?msg_server=server1
//
//  索引值一般为 mchid|appid.
const URLQueryServerKeyName = "msg_server"

// 多个 Server 的前端, 负责处理 http 请求, net/http.Handler 的实现
//
//  NOTE:
//  MultiServerFrontend 可以处理多个APP的消息，但是要求在回调 URL 上加上一个
//  查询参数，参考常量 URLQueryServerKeyName，这个参数的值就是 MultiServerFrontend
//  索引 Server 的 key。
//
//  例如回调 URL 为 http://www.xxx.com/notify_url?msg_server=1234567890，那么就可以在后端调用
//
//    MultiServerFrontend.SetServer("1234567890", Server)
//
//  来增加一个 Server 来处理 msg_server=1234567890 的消息。
//
//  MultiServerFrontend 并发安全，可以在运行中动态增加和删除 Server。
type MultiServerFrontend struct {
	invalidRequestHandler InvalidRequestHandler
	interceptor           Interceptor

	rwmutex   sync.RWMutex
	serverMap map[string]Server
}

// handler, interceptor 均可以为 nil
func NewMultiServerFrontend(handler InvalidRequestHandler, interceptor Interceptor) *MultiServerFrontend {
	if handler == nil {
		handler = DefaultInvalidRequestHandler
	}

	return &MultiServerFrontend{
		invalidRequestHandler: handler,
		interceptor:           interceptor,
		serverMap:             make(map[string]Server),
	}
}

// 设置 serverKey-Server pair.
func (frontend *MultiServerFrontend) SetServer(serverKey string, server Server) (err error) {
	if serverKey == "" {
		return errors.New("empty serverKey")
	}
	if server == nil {
		return errors.New("nil Server")
	}

	frontend.rwmutex.Lock()
	frontend.serverMap[serverKey] = server
	frontend.rwmutex.Unlock()
	return
}

// 删除 serverKey 对应的 Server
func (frontend *MultiServerFrontend) DeleteServer(serverKey string) {
	frontend.rwmutex.Lock()
	delete(frontend.serverMap, serverKey)
	frontend.rwmutex.Unlock()
}

// 删除所有的 Server
func (frontend *MultiServerFrontend) DeleteAllServer() {
	frontend.rwmutex.Lock()
	frontend.serverMap = make(map[string]Server)
	frontend.rwmutex.Unlock()
}

// 实现 http.Handler
func (frontend *MultiServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		frontend.invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if interceptor := frontend.interceptor; interceptor != nil && !interceptor.Intercept(w, r, queryValues) {
		return
	}

	serverKey := queryValues.Get(URLQueryServerKeyName)
	if serverKey == "" {
		err = fmt.Errorf("the url query value with name %s is empty", URLQueryServerKeyName)
		frontend.invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	frontend.rwmutex.RLock()
	server := frontend.serverMap[serverKey]
	frontend.rwmutex.RUnlock()

	if server == nil {
		err = fmt.Errorf("Not found Server for %s == %s", URLQueryServerKeyName, serverKey)
		frontend.invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, queryValues, server, frontend.invalidRequestHandler)
}
