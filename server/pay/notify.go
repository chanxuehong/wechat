// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"net/http"
)

// 根据密钥编号获取密钥
type GetSignKeyFunc func(keyIndex int) string

// 支付成功通知消息的 Handler
type NotifyHandler struct {
	paySignKey            string
	getSignKey            GetSignKeyFunc
	invalidRequestHandler InvalidRequestHandlerFunc
	notifyHandler         NotifyHandlerFunc
}

func NewNotifyHandler(
	paySignKey string,
	getSignKey GetSignKeyFunc,
	invalidRequestHandler InvalidRequestHandlerFunc,
	notifyHandler NotifyHandlerFunc,

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
	if notifyHandler == nil {
		panic("notifyHandler == nil")
	}

	handler = &NotifyHandler{
		paySignKey:            paySignKey,
		getSignKey:            getSignKey,
		invalidRequestHandler: invalidRequestHandler,
		notifyHandler:         notifyHandler,
	}

	return
}

// NotifyHandler 实现 http.Handler 接口
func (handler *NotifyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
