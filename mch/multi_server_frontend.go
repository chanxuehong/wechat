// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mch

import (
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

// 回调 URL 上索引 MessageServer 的 key 的名称.
//  比如下面的回调地址里面就可以根据 server1 来索引对应的 MessageServer.
//  http://www.xxx.com/?msg_server=server1
//
//  索引值一般为 mchid|appid.
const URLQueryMessageServerKeyName = "msg_server"

// 多个 MessageServer 的前端, 负责处理 http 请求, net/http.Handler 的实现
//
//  NOTE:
//  MultiMessageServerFrontend 可以处理多个APP的消息，但是要求在回调 URL 上加上一个
//  查询参数，参考常量 URLQueryMessageServerKeyName，这个参数的值就是 MultiMessageServerFrontend
//  索引 MessageServer 的 key。
//
//  例如回调 URL 为 http://www.xxx.com/notify_url?msg_server=1234567890，那么就可以在后端调用
//
//    MultiMessageServerFrontend.SetMessageServer("1234567890", MessageServer)
//
//  来增加一个 MessageServer 来处理 msg_server=1234567890 的消息。
//
//  MultiMessageServerFrontend 并发安全，可以在运行中动态增加和删除 MessageServer。
type MultiMessageServerFrontend struct {
	rwmutex               sync.RWMutex
	messageServerMap      map[string]MessageServer
	invalidRequestHandler InvalidRequestHandler
}

// 设置 InvalidRequestHandler, 如果 handler == nil 则使用默认的 DefaultInvalidRequestHandler
func (frontend *MultiMessageServerFrontend) SetInvalidRequestHandler(handler InvalidRequestHandler) {
	frontend.rwmutex.Lock()
	defer frontend.rwmutex.Unlock()

	if handler == nil {
		frontend.invalidRequestHandler = DefaultInvalidRequestHandler
	} else {
		frontend.invalidRequestHandler = handler
	}
}

// 设置 serverKey-MessageServer pair.
// 如果 serverKey == "" 或者 server == nil 则不做任何操作
func (frontend *MultiMessageServerFrontend) SetMessageServer(serverKey string, server MessageServer) {
	if serverKey == "" {
		return
	}
	if server == nil {
		return
	}

	frontend.rwmutex.Lock()
	defer frontend.rwmutex.Unlock()

	if frontend.messageServerMap == nil {
		frontend.messageServerMap = make(map[string]MessageServer)
	}
	frontend.messageServerMap[serverKey] = server
}

// 删除 serverKey 对应的 MessageServer
func (frontend *MultiMessageServerFrontend) DeleteMessageServer(serverKey string) {
	frontend.rwmutex.Lock()
	defer frontend.rwmutex.Unlock()

	delete(frontend.messageServerMap, serverKey)
}

// 删除所有的 MessageServer
func (frontend *MultiMessageServerFrontend) DeleteAllMessageServer() {
	frontend.rwmutex.Lock()
	defer frontend.rwmutex.Unlock()

	frontend.messageServerMap = make(map[string]MessageServer)
}

// 实现 http.Handler
func (frontend *MultiMessageServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	urlValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		frontend.rwmutex.RLock()
		invalidRequestHandler := frontend.invalidRequestHandler
		frontend.rwmutex.RUnlock()

		if invalidRequestHandler == nil {
			invalidRequestHandler = DefaultInvalidRequestHandler
		}
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	serverKey := urlValues.Get(URLQueryMessageServerKeyName)
	if serverKey == "" {
		frontend.rwmutex.RLock()
		invalidRequestHandler := frontend.invalidRequestHandler
		frontend.rwmutex.RUnlock()

		if invalidRequestHandler == nil {
			invalidRequestHandler = DefaultInvalidRequestHandler
		}
		err = fmt.Errorf("the url query value with name %s is empty", URLQueryMessageServerKeyName)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	frontend.rwmutex.RLock()
	invalidRequestHandler := frontend.invalidRequestHandler
	messageServer := frontend.messageServerMap[serverKey]
	frontend.rwmutex.RUnlock()

	if invalidRequestHandler == nil {
		invalidRequestHandler = DefaultInvalidRequestHandler
	}
	if messageServer == nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("Not found MessageServer for %s == %s", URLQueryMessageServerKeyName, serverKey))
		return
	}

	ServeHTTP(w, r, nil, messageServer, invalidRequestHandler)
}
