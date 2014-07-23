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
	"net/url"
)

// 支付成功通知消息的 Handler
type NotifyHandler struct {
	paySignKey            string
	invalidRequestHandler InvalidRequestHandlerFunc
	notifyHandler         NotifyHandlerFunc
}

func NewNotifyHandler(
	paySignKey string,
	invalidRequestHandler InvalidRequestHandlerFunc,
	notifyHandler NotifyHandlerFunc,

) (handler *NotifyHandler) {

	if paySignKey == "" {
		panic(`paySignKey == ""`)
	}
	if invalidRequestHandler == nil {
		panic("invalidRequestHandler == nil")
	}
	if notifyHandler == nil {
		panic("notifyHandler == nil")
	}

	handler = &NotifyHandler{
		paySignKey:            paySignKey,
		invalidRequestHandler: invalidRequestHandler,
		notifyHandler:         notifyHandler,
	}

	return
}

// NotifyHandler 实现 http.Handler 接口
func (handler *NotifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		handler.invalidRequestHandler(w, r, err)
		return
	}

	if r.URL == nil {
		err := errors.New("r.URL == nil")
		handler.invalidRequestHandler(w, r, err)
		return
	}

	var postData pay.NotifyPostData
	if err := xml.NewDecoder(r.Body).Decode(&postData); err != nil {
		handler.invalidRequestHandler(w, r, err)
		return
	}

	if err := postData.Check(handler.paySignKey); err != nil {
		handler.invalidRequestHandler(w, r, err)
		return
	}

	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		handler.invalidRequestHandler(w, r, err)
		return
	}

	var urlData pay.NotifyURLData
	if err := urlData.CheckAndInit(values, handler.paySignKey); err != nil {
		handler.invalidRequestHandler(w, r, err)
		return
	}

	handler.notifyHandler(w, r, &postData, &urlData)
}
