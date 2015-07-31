// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package suite

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"github.com/chanxuehong/wechat/corp"
)

// 多个 CorpServer 的前端, http.Handler 的实现,对普通corp.AgentServer的实现.
//
//  MultiCorpServerFrontend 可以处理多个企业号应用的消息(事件), 但是要求在回调 URL 上加上一个
//  查询参数(参数名与 urlCorpServerQueryName 一致), 通过这个参数的值来索引对应的 AgentServer.
//
//  例如回调 URL 为(urlCorpServerQueryName == "corp_server"):
//    http://www.xxx.com/weixin?corp_server=1234567890
//  那么就可以在后端调用
//   MultiCorpServerFrontend.SetCorpServer("1234567890", corpServer)
//  来增加一个 CorpServer 来处理 corp_server=1234567890 的消息(事件).
//
//  MultiCorpServerFrontend 并发安全, 可以在运行中动态增加和删除 CorpServer.
type MultiCorpServerFrontend struct {
	urlCorpServerQueryName string

	errHandler  corp.ErrorHandler
	interceptor corp.Interceptor

	rwmutex        sync.RWMutex
	corpServerMap map[string]corp.AgentServer
}

// NewMultiCorpServerFrontend 创建一个新的 MultiCorpServerFrontend.
//  urlCorpServerQueryName: 回调 URL 上参数名, 这个参数的值就是索引 CorpServer 的 key
//  errHandler:              错误处理 handler, 可以为 nil
//  interceptor:             拦截器, 可以为 nil
func NewMultiCorpServerFrontend(urlCorpServerQueryName string, errHandler corp.ErrorHandler, interceptor corp.Interceptor) *MultiCorpServerFrontend {
	if urlCorpServerQueryName == "" {
		urlCorpServerQueryName = "corp_server"
	}
	if errHandler == nil {
		errHandler = corp.DefaultErrorHandler
	}

	return &MultiCorpServerFrontend{
		urlCorpServerQueryName: urlCorpServerQueryName,
		errHandler:              errHandler,
		interceptor:             interceptor,
		corpServerMap:          make(map[string]corp.AgentServer),
	}
}

func (frontend *MultiCorpServerFrontend) SetCorpServer(serverKey string, server corp.AgentServer) (err error) {
	if serverKey == "" {
		return errors.New("empty serverKey")
	}
	if server == nil {
		return errors.New("nil CorpServer")
	}

	frontend.rwmutex.Lock()
	frontend.corpServerMap[serverKey] = server
	frontend.rwmutex.Unlock()
	return
}

func (frontend *MultiCorpServerFrontend) DeleteCorpServer(serverKey string) {
	frontend.rwmutex.Lock()
	delete(frontend.corpServerMap, serverKey)
	frontend.rwmutex.Unlock()
}

func (frontend *MultiCorpServerFrontend) DeleteAllCorpServer() {
	frontend.rwmutex.Lock()
	frontend.corpServerMap = make(map[string]corp.AgentServer)
	frontend.rwmutex.Unlock()
}

func (frontend *MultiCorpServerFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	queryValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		frontend.errHandler.ServeError(w, r, err)
		return
	}

	if interceptor := frontend.interceptor; interceptor != nil && !interceptor.Intercept(w, r, queryValues) {
		return
	}

	serverKey := queryValues.Get(frontend.urlCorpServerQueryName)
	if serverKey == "" {
		err := fmt.Errorf("the url query value with name %s is empty", frontend.urlCorpServerQueryName)
		frontend.errHandler.ServeError(w, r, err)
		return
	}

	frontend.rwmutex.RLock()
	agentServer := frontend.corpServerMap[serverKey]
	frontend.rwmutex.RUnlock()

	if agentServer == nil {
		err := fmt.Errorf("Not found CorpServer for %s == %s", frontend.urlCorpServerQueryName, serverKey)
		frontend.errHandler.ServeError(w, r, err)
		return
	}

	CorpServeHTTP(w, r, queryValues, agentServer, frontend.errHandler)
}
