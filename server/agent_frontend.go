// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"crypto/subtle"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/message/passive/request"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Agent 的前端, 负责处理 http 请求, net/http.Handler 的实现
//  NOTE: 只能处理一个公众号的消息
type AgentFrontend struct {
	agent                 Agent
	invalidRequestHandler InvalidRequestHandler
}

func NewAgentFrontend(agent Agent, invalidRequestHandler InvalidRequestHandler) *AgentFrontend {
	if agent == nil {
		panic("agent == nil")
	}
	if invalidRequestHandler == nil {
		invalidRequestHandler = InvalidRequestHandlerFunc(defaultInvalidRequestHandlerFunc)
	}

	return &AgentFrontend{
		agent: agent,
		invalidRequestHandler: invalidRequestHandler,
	}
}

func (this *AgentFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	agent := this.agent
	invalidRequestHandler := this.invalidRequestHandler

	switch r.Method {
	case "POST": // 处理从微信服务器推送过来的消息(事件) ==============================
		signature1, timestampStr, nonce, err := parsePostURLQuery(r.URL)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		signature2 := signature(agent.GetToken(), timestampStr, nonce)
		if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("check signature failed"))
			return
		}

		rawXMLMsg, err := ioutil.ReadAll(r.Body)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		var msgReq request.Request
		if err = xml.Unmarshal(rawXMLMsg, &msgReq); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		wantToUserName := agent.GetId()
		if len(msgReq.ToUserName) != len(wantToUserName) {
			err = fmt.Errorf("the message Request's ToUserName mismatch, have: %s, want: %s", msgReq.ToUserName, wantToUserName)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		if subtle.ConstantTimeCompare([]byte(msgReq.ToUserName), []byte(wantToUserName)) != 1 {
			err = fmt.Errorf("the message Request's ToUserName mismatch, have: %s, want: %s", msgReq.ToUserName, wantToUserName)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msgDispatch(w, r, &msgReq, rawXMLMsg, timestamp, agent)

	case "GET": // 首次验证 ======================================================
		signature1, timestamp, nonce, echostr, err := parseGetURLQuery(r.URL)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		signature2 := signature(agent.GetToken(), timestamp, nonce)
		if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("check signature failed"))
			return
		}

		io.WriteString(w, echostr)
	}
}
