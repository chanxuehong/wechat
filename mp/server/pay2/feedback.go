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
	"github.com/chanxuehong/wechat/mp/pay/feedback"
	"io/ioutil"
	"net/http"
)

// 用户维权的 Handler
type FeedbackHandler struct {
	agent                 Agent
	invalidRequestHandler InvalidRequestHandler
}

func NewFeedbackHandler(agent Agent, invalidRequestHandler InvalidRequestHandler) *FeedbackHandler {
	if agent == nil {
		panic("agent == nil")
	}
	if invalidRequestHandler == nil {
		invalidRequestHandler = InvalidRequestHandlerFunc(defaultInvalidRequestHandlerFunc)
	}

	return &FeedbackHandler{
		agent: agent,
		invalidRequestHandler: invalidRequestHandler,
	}
}

// FeedbackHandler 实现 http.Handler 接口
func (handler *FeedbackHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	var mixedReq feedback.MixedRequest
	if err := xml.Unmarshal(rawXMLMsg, &mixedReq); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	wantAppId := agent.GetAppId()
	if len(mixedReq.AppId) != len(wantAppId) {
		err = fmt.Errorf("AppId mismatch, have: %q, want: %q", mixedReq.AppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}
	if subtle.ConstantTimeCompare([]byte(mixedReq.AppId), []byte(wantAppId)) != 1 {
		err = fmt.Errorf("AppId mismatch, have: %q, want: %q", mixedReq.AppId, wantAppId)
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	if err := mixedReq.CheckSignature(agent.GetAppKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	switch mixedReq.MsgType {
	case feedback.MSG_TYPE_COMPLAIN:
		agent.ServeFeedbackComplaint(w, r, mixedReq.GetComplaint(), rawXMLMsg)

	case feedback.MSG_TYPE_CONFIRM:
		agent.ServeFeedbackConfirmation(w, r, mixedReq.GetConfirmation(), rawXMLMsg)

	case feedback.MSG_TYPE_REJECT:
		agent.ServeFeedbackRejection(w, r, mixedReq.GetRejection(), rawXMLMsg)

	default:
		agent.ServeUnknownMsg(w, r, rawXMLMsg)
	}
}
