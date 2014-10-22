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

// 微信后台向商户推送告警通知的 Handler
type AlarmNotifyHandler struct {
	agent                 Agent
	invalidRequestHandler InvalidRequestHandler
}

func NewAlarmNotifyHandler(agent Agent, invalidRequestHandler InvalidRequestHandler) *AlarmNotifyHandler {
	if agent == nil {
		panic("agent == nil")
	}
	if invalidRequestHandler == nil {
		invalidRequestHandler = InvalidRequestHandlerFunc(defaultInvalidRequestHandlerFunc)
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

	var alarmData pay2.AlarmNotifyPostData
	if err := xml.Unmarshal(rawXMLMsg, &alarmData); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	wantAppId := agent.GetAppId()
	if len(alarmData.AppId) != len(wantAppId) {
		err = fmt.Errorf("AppId mismatch, have: %q, want: %q", alarmData.AppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(alarmData.AppId), []byte(wantAppId)) != 1 {
		err = fmt.Errorf("AppId mismatch, have: %q, want: %q", alarmData.AppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if err := alarmData.CheckSignature(agent.GetAppKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	agent.ServeAlarmNotification(w, r, &alarmData, rawXMLMsg)
}
