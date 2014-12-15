// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"bytes"
	"crypto/subtle"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/chanxuehong/wechat/mp/pay"
	"github.com/chanxuehong/wechat/mp/pay/pay2"
)

// native api 请求订单详情的 Handler
type PayPackageRequestHandler struct {
	agent                 Agent
	invalidRequestHandler InvalidRequestHandler
}

// 创建一个新的 PayPackageRequestHandler.
//  agent 不能为 nil, 如果 invalidRequestHandler == nil 则使用 DefaultInvalidRequestHandler
func NewPayPackageRequestHandler(agent Agent, invalidRequestHandler InvalidRequestHandler) *PayPackageRequestHandler {
	if agent == nil {
		panic("agent == nil")
	}
	if invalidRequestHandler == nil {
		invalidRequestHandler = DefaultInvalidRequestHandler
	}

	return &PayPackageRequestHandler{
		agent: agent,
		invalidRequestHandler: invalidRequestHandler,
	}
}

// PayPackageRequestHandler 实现 http.Handler 接口
func (handler *PayPackageRequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agent := handler.agent
	invalidRequestHandler := handler.invalidRequestHandler

	ServePayPackageRequestHTTP(w, r, nil, agent, invalidRequestHandler)
}

// ServePayPackageRequestHTTP 处理 http 消息请求
//  NOTE: 确保所有参数合法, r.Body 能正确读取数据
func ServePayPackageRequestHTTP(w http.ResponseWriter, r *http.Request,
	urlValues url.Values, agent Agent, invalidRequestHandler InvalidRequestHandler) {

	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	postRawXMLMsg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	req := make(pay2.PayPackageRequest)
	if err = pay.ParseXMLToMap(bytes.NewReader(postRawXMLMsg), req); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	haveAppId := req.AppId()
	wantAppId := agent.GetAppId()
	if len(haveAppId) != len(wantAppId) {
		err = fmt.Errorf("AppId mismatch, have: %q, want: %q", haveAppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(haveAppId), []byte(wantAppId)) != 1 {
		err = fmt.Errorf("AppId mismatch, have: %q, want: %q", haveAppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if err := req.CheckSignature(agent.GetAppKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	para := RequestParameters{
		HTTPResponseWriter: w,
		HTTPRequest:        r,
		PostRawXMLMsg:      postRawXMLMsg,
	}
	agent.ServePayPackageRequest(req, &para)
}
