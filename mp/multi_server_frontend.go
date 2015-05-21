// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

// 回调 URL 上索引 Server 的 key 的名称.
//  比如下面的回调地址里面就可以根据 wechat1 来索引对应的 Server.
//  http://www.xxx.com/?wechat_server=wechat1&signature=XXX&timestamp=123456789&nonce=12345678
const URLQueryServerKeyName = "wechat_server"

// 多个 Server 的前端, 负责处理 http 请求, net/http.Handler 的实现
//
//  NOTE:
//  MultiServerFrontend 可以处理多个公众号的消息（事件），但是要求在回调 URL 上加上一个
//  查询参数，参考常量 URLQueryServerKeyName，这个参数的值就是 MultiServerFrontend
//  索引 Server 的 key。
//
//  例如回调 URL 为 http://www.xxx.com/weixin?wechat_server=1234567890，那么就可以在后端调用
//
//    MultiServerFrontend.SetServer("1234567890", Server)
//
//  来增加一个 Server 来处理 wechat_server=1234567890 的消息（事件）。
//
//  MultiServerFrontend 并发安全，可以在运行中动态增加和删除 Server。
type MultiServerFrontend struct {
	rwmutex               sync.RWMutex
	serverMap             map[string]Server
	invalidRequestHandler InvalidRequestHandler
}

// 设置 InvalidRequestHandler, 如果 handler == nil 则使用默认的 DefaultInvalidRequestHandler
func (frontend *MultiServerFrontend) SetInvalidRequestHandler(handler InvalidRequestHandler) {
	frontend.rwmutex.Lock()
	if handler == nil {
		frontend.invalidRequestHandler = DefaultInvalidRequestHandler
	} else {
		frontend.invalidRequestHandler = handler
	}
	frontend.rwmutex.Unlock()
}

// 设置 serverKey-Server pair.
// 如果 serverKey == "" 或者 server == nil 则不做任何操作
func (frontend *MultiServerFrontend) SetServer(serverKey string, server Server) {
	if serverKey == "" {
		return
	}
	if server == nil {
		return
	}

	frontend.rwmutex.Lock()
	if frontend.serverMap == nil {
		frontend.serverMap = make(map[string]Server)
	}
	frontend.serverMap[serverKey] = server
	frontend.rwmutex.Unlock()
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

func (frontend *MultiServerFrontend) getInvalidRequestHandler() (h InvalidRequestHandler) {
	frontend.rwmutex.RLock()
	h = frontend.invalidRequestHandler
	if h == nil {
		h = DefaultInvalidRequestHandler
	}
	frontend.rwmutex.RUnlock()
	return
}

// 实现 http.Handler
func (frontend *MultiServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		frontend.getInvalidRequestHandler().ServeInvalidRequest(w, r, err)
		return
	}

	serverKey := queryValues.Get(URLQueryServerKeyName)
	if serverKey == "" {
		err = fmt.Errorf("the url query value with name %s is empty", URLQueryServerKeyName)
		frontend.getInvalidRequestHandler().ServeInvalidRequest(w, r, err)
		return
	}

	frontend.rwmutex.RLock()
	invalidRequestHandler := frontend.invalidRequestHandler
	server := frontend.serverMap[serverKey]
	frontend.rwmutex.RUnlock()

	if invalidRequestHandler == nil {
		invalidRequestHandler = DefaultInvalidRequestHandler
	}
	if server == nil {
		err = fmt.Errorf("Not found Server for %s == %s", URLQueryServerKeyName, serverKey)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, queryValues, server, invalidRequestHandler)
}
