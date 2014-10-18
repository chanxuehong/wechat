// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay

import (
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/wechat/mp/pay/feedback"
	"net/http"
)

// 维权接口的 Handler
type FeedbackHandler struct {
	paySignKey     string
	invalidHandler InvalidRequestHandlerFunc
	requestHandler FeedbackRequestHandlerFunc
	confirmHandler FeedbackConfirmHandlerFunc
	rejectHandler  FeedbackRejectHandlerFunc
}

// NOTE: 所有参数必须有效
func NewFeedbackHandler(
	paySignKey string,
	invalidHandler InvalidRequestHandlerFunc,
	requestHandler FeedbackRequestHandlerFunc,
	confirmHandler FeedbackConfirmHandlerFunc,
	rejectHandler FeedbackRejectHandlerFunc,

) (handler *FeedbackHandler) {

	if paySignKey == "" {
		panic(`paySignKey == ""`)
	}
	if invalidHandler == nil {
		panic("invalidHandler == nil")
	}
	if requestHandler == nil {
		panic("requestHandler == nil")
	}
	if confirmHandler == nil {
		panic("confirmHandler == nil")
	}
	if rejectHandler == nil {
		panic("rejectHandler == nil")
	}

	handler = &FeedbackHandler{
		paySignKey:     paySignKey,
		invalidHandler: invalidHandler,
		requestHandler: requestHandler,
		confirmHandler: confirmHandler,
		rejectHandler:  rejectHandler,
	}

	return
}

// FeedbackHandler 实现 http.Handler 接口
func (handler *FeedbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		err := errors.New("request method is not POST")
		handler.invalidHandler(w, r, err)
		return
	}

	var msgReq feedback.MsgRequest
	if err := xml.NewDecoder(r.Body).Decode(&msgReq); err != nil {
		handler.invalidHandler(w, r, err)
		return
	}
	if err := msgReq.Check(handler.paySignKey); err != nil {
		handler.invalidHandler(w, r, err)
		return
	}

	switch msgReq.MsgType {
	case feedback.MSG_TYPE_REQUEST:
		req := msgReq.GetRequest()
		handler.requestHandler(w, r, req)
		return

	case feedback.MSG_TYPE_CONFIRM:
		cfm := msgReq.GetConfirm()
		handler.confirmHandler(w, r, cfm)
		return

	case feedback.MSG_TYPE_REJECT:
		rjt := msgReq.GetReject()
		handler.rejectHandler(w, r, rjt)
		return

	default:
		err := errors.New("未知的消息类型" + msgReq.MsgType)
		handler.invalidHandler(w, r, err)
		return
	}
}
