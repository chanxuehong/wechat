// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/wechat/pay"
	"net/http"
)

// 告警处理 Handler
type AlarmNotifyHandler struct {
	paySignKey        string
	invalidHandler    InvalidRequestHandlerFunc
	notifyHandlerFunc AlarmNotifyHandlerFunc
}

// NOTE: 所有参数必须有效
func NewAlarmNotifyHandler(
	paySignKey string,
	invalidHandler InvalidRequestHandlerFunc,
	notifyHandlerFunc AlarmNotifyHandlerFunc,

) (handler *AlarmNotifyHandler) {

	if paySignKey == "" {
		panic(`paySignKey == ""`)
	}
	if invalidHandler == nil {
		panic("invalidHandler == nil")
	}
	if notifyHandlerFunc == nil {
		panic("notifyHandlerFunc == nil")
	}

	handler = &AlarmNotifyHandler{
		paySignKey:        paySignKey,
		invalidHandler:    invalidHandler,
		notifyHandlerFunc: notifyHandlerFunc,
	}

	return
}

// AlarmNotifyHandler 实现 http.Handler 接口
func (handler *AlarmNotifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		handler.invalidHandler(w, r, err)
		return
	}

	var alarmData pay.AlarmNotifyData
	if err := xml.NewDecoder(r.Body).Decode(&alarmData); err != nil {
		handler.invalidHandler(w, r, err)
		return
	}

	if err := alarmData.Check(handler.paySignKey); err != nil {
		handler.invalidHandler(w, r, err)
		return
	}

	handler.notifyHandlerFunc(w, r, &alarmData)
}
