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
	"strconv"
)

// 根据密钥 index 获取密钥, 找不到合法的密钥返回空值 ""
type GetSignKey func(keyIndex int) string

// 支付成功通知消息的 Handler
type NotifyHandler struct {
	paySignKey            string
	getSignKey            GetSignKey
	invalidRequestHandler InvalidRequestHandlerFunc
	notifyHandlerVer1     NotifyHandlerFuncVer1
}

func NewNotifyHandler(
	paySignKey string,
	getSignKey GetSignKey,
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

	var signKeyIndex int
	if signKeyIndexes := values["sign_key_index"]; len(signKeyIndexes) > 0 && len(signKeyIndexes[0]) > 0 {
		index, err := strconv.ParseInt(signKeyIndexes[0], 10, 64)
		if err != nil {
			err = fmt.Errorf("获取密钥 index 出错: %s", err.Error())
			handler.invalidRequestHandler(w, r, err)
			return
		}
		signKeyIndex = int(index)
	} else {
		signKeyIndex = 1
	}

	urlSignKey := handler.getSignKey(signKeyIndex)
	if urlSignKey == "" {
		err = fmt.Errorf("获取index 为 %d 的密钥失败", signKeyIndex)
		handler.invalidRequestHandler(w, r, err)
		return
	}

	switch serviceVersion {
	case "1.0":
		var urlData pay.NotifyURLDataVer1
		if err := urlData.CheckAndInit(values, urlSignKey); err != nil {
			handler.invalidRequestHandler(w, r, err)
			return
		}
		handler.notifyHandlerVer1(w, r, &postData, &urlData)

	default:
		err := fmt.Errorf("没有实现接口版本号为 %s 的支持", serviceVersion)
		handler.invalidRequestHandler(w, r, err)
		return
	}
}
