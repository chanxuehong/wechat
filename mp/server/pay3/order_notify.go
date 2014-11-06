// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/mp/pay"
	"github.com/chanxuehong/wechat/mp/pay/pay3"
	"io/ioutil"
	"net/http"
	"net/url"
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

	ServeOrderNotifyHTTP(w, r, nil, agent, invalidRequestHandler)
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

	postRawXMLMsg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	data := make(map[string]string)
	if err = pay.ParseXMLToMap(bytes.NewReader(postRawXMLMsg), data); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	haveAppId := data["appid"]
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

	haveMchId := data["mch_id"]
	wantMchId := agent.GetMchId()
	if len(haveMchId) != len(wantMchId) {
		err = fmt.Errorf("MchId mismatch, have: %q, want: %q", haveMchId, wantMchId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(haveMchId), []byte(wantMchId)) != 1 {
		err = fmt.Errorf("MchId mismatch, have: %q, want: %q", haveMchId, wantMchId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if err := pay3.CheckMD5Signature(data, agent.GetKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	agent.ServeOrderNotification(w, r, data, postRawXMLMsg)
}
