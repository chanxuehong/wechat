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
	"github.com/chanxuehong/wechat/mp/pay"
	"github.com/chanxuehong/wechat/mp/pay/pay2"
	"io/ioutil"
	"net/http"
)

// native api 请求订单详情的 Handler
type PayPackageRequestHandler struct {
	agent                 Agent
	invalidRequestHandler InvalidRequestHandler
}

func NewPayPackageRequestHandler(agent Agent, invalidRequestHandler InvalidRequestHandler) *PayPackageRequestHandler {
	if agent == nil {
		panic("agent == nil")
	}
	if invalidRequestHandler == nil {
		invalidRequestHandler = InvalidRequestHandlerFunc(defaultInvalidRequestHandlerFunc)
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

	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	rawXMLMsg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	req := make(pay2.PayPackageRequest)
	if err = pay.ParseXMLToMap(bytes.NewReader(rawXMLMsg), req); err != nil {
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

	agent.ServePayPackageRequest(w, r, &req, rawXMLMsg)
}
