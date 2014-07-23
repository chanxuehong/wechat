// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/pay"
	"net/http"
	"net/url"
)

// 支付成功通知消息的 Handler
type NotifyHandler struct {
	paySignKey            string         // post 部分签名密钥
	getSignKey            pay.GetSignKey // url 部分签名密钥获取函数
	invalidRequestHandler InvalidRequestHandlerFunc
	notifyHandlerVer1     NotifyHandlerFuncVer1
}

func NewNotifyHandler(
	paySignKey string,
	getSignKey pay.GetSignKey,
	invalidRequestHandler InvalidRequestHandlerFunc,
	notifyHandlerVer1 NotifyHandlerFuncVer1,

) (handler *NotifyHandler) {

	if paySignKey == "" {
		panic(`paySignKey == ""`)
	}
	if getSignKey == nil {
		panic("getSignKey == nil")
	}
	if invalidRequestHandler == nil {
		panic("invalidRequestHandler == nil")
	}
	if notifyHandlerVer1 == nil {
		panic("notifyHandlerVer1 == nil")
	}

	handler = &NotifyHandler{
		getSignKey:            getSignKey,
		invalidRequestHandler: invalidRequestHandler,
		notifyHandlerVer1:     notifyHandlerVer1,
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

	var serviceVersion string
	if serviceVersions := values["service_version"]; len(serviceVersions) > 0 && len(serviceVersions[0]) > 0 {
		serviceVersion = serviceVersions[0]
	} else {
		serviceVersion = "1.0"
	}

	switch serviceVersion {
	case "1.0":
		var urlData pay.NotifyURLDataVer1
		if err := urlData.CheckAndInit(values, handler.getSignKey); err != nil {
			handler.invalidRequestHandler(w, r, err)
			return
		}

		handler.notifyHandlerVer1(w, r, &postData, &urlData)

	default:
		err := fmt.Errorf("没有实现对接口版本号为 %s 的支持", serviceVersion)
		handler.invalidRequestHandler(w, r, err)
		return
	}
}
