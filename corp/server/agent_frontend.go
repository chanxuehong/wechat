// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/corp/message/passive/request"
	"net/http"
	"strconv"
)

// Agent 的前端, 负责处理 http 请求, net/http.Handler 的实现
//  NOTE: 只能处理一个企业号应用的消息
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
		msgSignature1, timestampStr, nonce, err := parsePostURLQuery(r.URL)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		const signatureLen = sha1.Size * 2
		if len(msgSignature1) != signatureLen {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: %d", len(msgSignature1), signatureLen)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("can not parse timestamp(==%q) to int64, error: %s", timestampStr, err.Error()))
			return
		}

		var requestHttpBody request.RequestHttpBody
		if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		wantCorpId := agent.GetCorpId()
		if len(requestHttpBody.CorpId) != len(wantCorpId) {
			err = fmt.Errorf("the message RequestHttpBody's ToUserName mismatch, have: %s, want: %s", requestHttpBody.CorpId, wantCorpId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		if subtle.ConstantTimeCompare([]byte(requestHttpBody.CorpId), []byte(wantCorpId)) != 1 {
			err = fmt.Errorf("the message RequestHttpBody's ToUserName mismatch, have: %s, want: %s", requestHttpBody.CorpId, wantCorpId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		wantAgentId := agent.GetAgentId()
		if requestHttpBody.AgentId != wantAgentId && requestHttpBody.AgentId != 0 {
			err = fmt.Errorf("the message RequestHttpBody's AgentId mismatch, have: %d, want: %d", requestHttpBody.AgentId, wantAgentId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msgSignature2 := msgSignature(agent.GetToken(), timestampStr, nonce, requestHttpBody.EncryptedMsg)
		if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("check signature failed"))
			return
		}

		EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptedMsg)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		random, rawXMLMsg, err := aesDecryptMsg(EncryptedMsgBytes, agent.GetCorpId(), agent.GetAESKey())
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		var msgReq request.Request
		if err := xml.Unmarshal(rawXMLMsg, &msgReq); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if requestHttpBody.CorpId != msgReq.ToUserName {
			err = fmt.Errorf("the RequestHttpBody's ToUserName(==%d) mismatch the Request's ToUserName(==%d)", requestHttpBody.CorpId, msgReq.ToUserName)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if requestHttpBody.AgentId != msgReq.AgentId {
			err = fmt.Errorf("the RequestHttpBody's AgentId(==%d) mismatch the Request's AgengId(==%d)", requestHttpBody.AgentId, msgReq.AgentId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 此时要么 msgReq.AgentId == wantAgentId, 要么 msgReq.AgentId == 0

		if msgReq.AgentId == 0 {
			// 订阅/取消订阅 整个企业号
			if msgReq.MsgType == request.MSG_TYPE_EVENT &&
				(msgReq.Event == request.EVENT_TYPE_SUBSCRIBE || msgReq.Event == request.EVENT_TYPE_UNSUBSCRIBE) {
				// do nothing
			} else {
				err = fmt.Errorf("the message Request's AgentId mismatch, have: %d, want: %d", msgReq.AgentId, wantAgentId)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
		}

		msgDispatch(w, r, &msgReq, rawXMLMsg, timestamp, nonce, random, agent)

	case "GET": // 首次验证 ======================================================
		msgSignature1, timestamp, nonce, encryptedMsg, err := parseGetURLQuery(r.URL)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		const signatureLen = sha1.Size * 2
		if len(msgSignature1) != signatureLen {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: %d", len(msgSignature1), signatureLen)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msgSignature2 := msgSignature(agent.GetToken(), timestamp, nonce, encryptedMsg)
		if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("check signature failed"))
			return
		}

		EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(encryptedMsg)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		_, echostr, err := aesDecryptMsg(EncryptedMsgBytes, agent.GetCorpId(), agent.GetAESKey())
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		w.Write(echostr)
	}
}
