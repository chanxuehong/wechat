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
type OrderNotifyHandler struct {
	paySignKey        string            // post 部分签名密钥
	getPartnerKey     pay.GetPartnerKey // url 部分签名密钥获取函数
	invalidHandler    InvalidRequestHandlerFunc
	notifyHandlerVer1 OrderNotifyHandlerFuncVer1
}

// NOTE: 所有参数必须有效
func NewOrderNotifyHandler(
	paySignKey string,
	getPartnerKey pay.GetPartnerKey,
	invalidHandler InvalidRequestHandlerFunc,
	notifyHandlerVer1 OrderNotifyHandlerFuncVer1,

) (handler *OrderNotifyHandler) {

	if paySignKey == "" {
		panic(`paySignKey == ""`)
	}
	if getPartnerKey == nil {
		panic("getPartnerKey == nil")
	}
	if invalidHandler == nil {
		panic("invalidHandler == nil")
	}
	if notifyHandlerVer1 == nil {
		panic("notifyHandlerVer1 == nil")
	}

	handler = &OrderNotifyHandler{
		paySignKey:        paySignKey,
		getPartnerKey:     getPartnerKey,
		invalidHandler:    invalidHandler,
		notifyHandlerVer1: notifyHandlerVer1,
	}

	return
}

// NotifyHandler 实现 http.Handler 接口
func (handler *OrderNotifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		handler.invalidHandler(w, r, err)
		return
	}

	if r.URL == nil {
		err := errors.New("r.URL == nil")
		handler.invalidHandler(w, r, err)
		return
	}

	var postData pay.OrderNotifyPostData
	if err := xml.NewDecoder(r.Body).Decode(&postData); err != nil {
		handler.invalidHandler(w, r, err)
		return
	}

	if err := postData.Check(handler.paySignKey); err != nil {
		handler.invalidHandler(w, r, err)
		return
	}

	values, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		handler.invalidHandler(w, r, err)
		return
	}

	// 确定接口版本号
	var serviceVersion string
	if serviceVersions := values["service_version"]; len(serviceVersions) > 0 && len(serviceVersions[0]) > 0 {
		serviceVersion = serviceVersions[0]
	} else {
		serviceVersion = "1.0"
	}

	switch serviceVersion {
	case "1.0":
		var urlData pay.OrderNotifyURLDataVer1
		if err := urlData.CheckAndInit(values, handler.getPartnerKey); err != nil {
			handler.invalidHandler(w, r, err)
			return
		}

		handler.notifyHandlerVer1(w, r, &postData, &urlData)

	default:
		err := fmt.Errorf("没有实现对接口版本号为 %s 的支持", serviceVersion)
		handler.invalidHandler(w, r, err)
		return
	}
}
