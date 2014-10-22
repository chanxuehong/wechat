// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay2

import (
	"crypto/subtle"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/mp/pay2"
	"io/ioutil"
	"net/http"
)

// 用户在成功完成支付后，微信后台通知（POST）商户服务器（notify_url）支付结果的处理 Handler
type OrderNotifyHandler struct {
	agent                 Agent
	invalidRequestHandler InvalidRequestHandler
}

func NewOrderNotifyHandler(agent Agent, invalidRequestHandler InvalidRequestHandler) *OrderNotifyHandler {
	if agent == nil {
		panic("agent == nil")
	}
	if invalidRequestHandler == nil {
		invalidRequestHandler = InvalidRequestHandlerFunc(defaultInvalidRequestHandlerFunc)
	}

	return &OrderNotifyHandler{
		agent: agent,
		invalidRequestHandler: invalidRequestHandler,
	}
}

// OrderNotifyHandler 实现 http.Handler 接口
func (handler *OrderNotifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agent := handler.agent
	invalidRequestHandler := handler.invalidRequestHandler

	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if r.URL == nil {
		err := errors.New("input net/http.Request.URL == nil")
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	var urlData pay2.OrderNotifyURLData
	if err := urlData.CheckAndInit(r.URL.RawQuery, agent.GetPartnerKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	wantPartnerId := agent.GetPartnerId()
	if len(urlData.PartnerId) != len(wantPartnerId) {
		err := fmt.Errorf("PartnerId mismatch, have: %q, want: %q", urlData.PartnerId, wantPartnerId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(urlData.PartnerId), []byte(wantPartnerId)) != 1 {
		err := fmt.Errorf("PartnerId mismatch, have: %q, want: %q", urlData.PartnerId, wantPartnerId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	rawXMLMsg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	var postData pay2.OrderNotifyPostData
	if err := xml.Unmarshal(rawXMLMsg, &postData); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	wantAppId := agent.GetAppId()
	if len(postData.AppId) != len(wantAppId) {
		err := fmt.Errorf("AppId mismatch, have: %q, want: %q", postData.AppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(postData.AppId), []byte(wantAppId)) != 1 {
		err := fmt.Errorf("AppId mismatch, have: %q, want: %q", postData.AppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if err := postData.CheckSignature(agent.GetAppKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	agent.ServeOrderNotification(w, r, &urlData, &postData, rawXMLMsg)
}
