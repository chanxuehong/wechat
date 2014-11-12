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
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/chanxuehong/wechat/mp/pay/feedback"
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
		invalidRequestHandler = DefaultInvalidRequestHandler
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

	ServeFeedbackHTTP(w, r, nil, agent, invalidRequestHandler)
}

// ServeFeedbackHTTP 处理 http 消息请求
//  NOTE: 确保所有参数合法, r.Body 能正确读取数据
func ServeFeedbackHTTP(w http.ResponseWriter, r *http.Request,
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

	var mixedReq feedback.MixedRequest
	if err := xml.Unmarshal(postRawXMLMsg, &mixedReq); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	haveAppId := mixedReq.AppId
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

	if err := mixedReq.CheckSignature(agent.GetAppKey()); err != nil {
		invalidRequestHandler.ServeInvalidRequest(w, r, err)
		return
	}

	switch mixedReq.MsgType {
	case feedback.MSG_TYPE_COMPLAIN:
		agent.ServeFeedbackComplaint(w, r, mixedReq.GetComplaint(), postRawXMLMsg)

	case feedback.MSG_TYPE_CONFIRM:
		agent.ServeFeedbackConfirmation(w, r, mixedReq.GetConfirmation(), postRawXMLMsg)

	case feedback.MSG_TYPE_REJECT:
		agent.ServeFeedbackRejection(w, r, mixedReq.GetRejection(), postRawXMLMsg)

	default:
		agent.ServeUnknownMsg(w, r, postRawXMLMsg)
	}
}
