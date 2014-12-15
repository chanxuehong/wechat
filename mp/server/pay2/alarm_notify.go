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

// 微信后台向商户推送告警通知的 Handler
type AlarmNotifyHandler struct {
	agent                 Agent
	invalidRequestHandler InvalidRequestHandler
}

// 创建一个新的 AlarmNotifyHandler.
//  agent 不能为 nil, 如果 invalidRequestHandler == nil 则使用 DefaultInvalidRequestHandler
func NewAlarmNotifyHandler(agent Agent, invalidRequestHandler InvalidRequestHandler) *AlarmNotifyHandler {
	if agent == nil {
		panic("agent == nil")
	}
	if invalidRequestHandler == nil {
		invalidRequestHandler = DefaultInvalidRequestHandler
	}

	return &AlarmNotifyHandler{
		agent: agent,
		invalidRequestHandler: invalidRequestHandler,
	}
}

// AlarmNotifyHandler 实现 http.Handler 接口
func (handler *AlarmNotifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agent := handler.agent
	invalidRequestHandler := handler.invalidRequestHandler

	ServeAlarmNotifyHTTP(w, r, nil, agent, invalidRequestHandler)
}

// ServeAlarmNotifyHTTP 处理 http 消息请求
//  NOTE: 确保所有参数合法, r.Body 能正确读取数据
func ServeAlarmNotifyHTTP(w http.ResponseWriter, r *http.Request,
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

	alarmData := make(pay2.AlarmNotifyPostData)
	if err := pay.ParseXMLToMap(bytes.NewReader(postRawXMLMsg), alarmData); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	haveAppId := alarmData.AppId()
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

	if err := alarmData.CheckSignature(agent.GetAppKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	para := RequestParameters{
		HTTPResponseWriter: w,
		HTTPRequest:        r,
		PostRawXMLMsg:      postRawXMLMsg,
	}
	agent.ServeAlarmNotification(alarmData, &para)
}
