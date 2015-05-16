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

// 回调 URL 上索引 WechatServer 的 key 的名称.
//  比如下面的回调地址里面就可以根据 wechat1 来索引对应的 WechatServer.
//  http://www.xxx.com/?wechat_server=wechat1&signature=XXX&timestamp=123456789&nonce=12345678
const URLQueryWechatServerKeyName = "wechat_server"

// 多个 WechatServer 的前端, 负责处理 http 请求, net/http.Handler 的实现
//
//  NOTE:
//  MultiWechatServerFrontend 可以处理多个公众号的消息（事件），但是要求在回调 URL 上加上一个
//  查询参数，参考常量 URLQueryWechatServerKeyName，这个参数的值就是 MultiWechatServerFrontend
//  索引 WechatServer 的 key。
//
//  例如回调 URL 为 http://www.xxx.com/weixin?wechat_server=1234567890，那么就可以在后端调用
//
//    MultiWechatServerFrontend.SetWechatServer("1234567890", WechatServer)
//
//  来增加一个 WechatServer 来处理 wechat_server=1234567890 的消息（事件）。
//
//  MultiWechatServerFrontend 并发安全，可以在运行中动态增加和删除 WechatServer。
type MultiWechatServerFrontend struct {
	rwmutex               sync.RWMutex
	wechatServerMap       map[string]WechatServer
	invalidRequestHandler InvalidRequestHandler
}

// 设置 InvalidRequestHandler, 如果 handler == nil 则使用默认的 DefaultInvalidRequestHandler
func (frontend *MultiWechatServerFrontend) SetInvalidRequestHandler(handler InvalidRequestHandler) {
	frontend.rwmutex.Lock()
	defer frontend.rwmutex.Unlock()

	if handler == nil {
		frontend.invalidRequestHandler = DefaultInvalidRequestHandler
	} else {
		frontend.invalidRequestHandler = handler
	}
}

// 设置 serverKey-WechatServer pair.
// 如果 serverKey == "" 或者 server == nil 则不做任何操作
func (frontend *MultiWechatServerFrontend) SetWechatServer(serverKey string, server WechatServer) {
	if serverKey == "" {
		return
	}
	if server == nil {
		return
	}

	frontend.rwmutex.Lock()
	defer frontend.rwmutex.Unlock()

	if frontend.wechatServerMap == nil {
		frontend.wechatServerMap = make(map[string]WechatServer)
	}
	frontend.wechatServerMap[serverKey] = server
}

// 删除 serverKey 对应的 WechatServer
func (frontend *MultiWechatServerFrontend) DeleteWechatServer(serverKey string) {
	frontend.rwmutex.Lock()
	defer frontend.rwmutex.Unlock()

	delete(frontend.wechatServerMap, serverKey)
}

// 删除所有的 WechatServer
func (frontend *MultiWechatServerFrontend) DeleteAllWechatServer() {
	frontend.rwmutex.Lock()
	defer frontend.rwmutex.Unlock()

	frontend.wechatServerMap = make(map[string]WechatServer)
}

func (frontend *MultiWechatServerFrontend) getInvalidRequestHandler() (h InvalidRequestHandler) {
	frontend.rwmutex.RLock()

	h = frontend.invalidRequestHandler
	if h == nil {
		h = DefaultInvalidRequestHandler
	}

	frontend.rwmutex.RUnlock()
	return
}

// 实现 http.Handler
func (frontend *MultiWechatServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		frontend.getInvalidRequestHandler().ServeInvalidRequest(w, r, err)
		return
	}

	serverKey := queryValues.Get(URLQueryWechatServerKeyName)
	if serverKey == "" {
		err = fmt.Errorf("the url query value with name %s is empty", URLQueryWechatServerKeyName)
		frontend.getInvalidRequestHandler().ServeInvalidRequest(w, r, err)
		return
	}

	frontend.rwmutex.RLock()
	invalidRequestHandler := frontend.invalidRequestHandler
	wechatServer := frontend.wechatServerMap[serverKey]
	frontend.rwmutex.RUnlock()

	if invalidRequestHandler == nil {
		invalidRequestHandler = DefaultInvalidRequestHandler
	}
	if wechatServer == nil {
		err = fmt.Errorf("Not found WechatServer for %s == %s", URLQueryWechatServerKeyName, serverKey)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeHTTP(w, r, queryValues, wechatServer, invalidRequestHandler)
}
