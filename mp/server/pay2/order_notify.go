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
		invalidRequestHandler = DefaultInvalidRequestHandler
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

	if r.URL == nil {
		err := errors.New("input net/http.Request.URL == nil")
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	urlValues, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	ServeOrderNotifyHTTP(w, r, urlValues, agent, invalidRequestHandler)
}

// ServeOrderNotifyHTTP 处理 http 消息请求
//  NOTE: 确保所有参数合法, r.Body 能正确读取数据
func ServeOrderNotifyHTTP(w http.ResponseWriter, r *http.Request,
	urlValues url.Values, agent Agent, invalidRequestHandler InvalidRequestHandler) {

	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	// 处理 post 部分
	postRawXMLMsg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	postData := make(pay2.OrderNotifyPostData)
	if err = pay.ParseXMLToMap(bytes.NewReader(postRawXMLMsg), postData); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	haveAppId := postData.AppId()
	wantAppId := agent.GetAppId()
	if len(haveAppId) != len(wantAppId) {
		err := fmt.Errorf("AppId mismatch, have: %q, want: %q", haveAppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(haveAppId), []byte(wantAppId)) != 1 {
		err := fmt.Errorf("AppId mismatch, have: %q, want: %q", haveAppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if err := postData.CheckSignature(agent.GetAppKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	// 处理 url 部分
	urlData := pay2.OrderNotifyURLData(urlValues)

	havePartnerId := urlData.PartnerId()
	wantPartnerId := agent.GetPartnerId()
	if len(havePartnerId) != len(wantPartnerId) {
		err := fmt.Errorf("PartnerId mismatch, have: %q, want: %q", havePartnerId, wantPartnerId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(havePartnerId), []byte(wantPartnerId)) != 1 {
		err := fmt.Errorf("PartnerId mismatch, have: %q, want: %q", havePartnerId, wantPartnerId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if err := urlData.CheckSignature(agent.GetPartnerKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	agent.ServeOrderNotification(w, r, urlData, postData, postRawXMLMsg)
}
