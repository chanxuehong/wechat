// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"bytes"
	"crypto/sha1"
	"crypto/subtle"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/message/passive/request"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

var zeroAESKey [32]byte

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
		signature1, timestampStr, nonce, encryptType, msgSignature1, err := parsePostURLQuery(r.URL)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		switch encryptType {
		case "", "raw": // 明文模式
			const signatureLen = sha1.Size * 2
			if len(signature1) != signatureLen {
				err = fmt.Errorf("the length of signature mismatch, have: %d, want: %d", len(signature1), signatureLen)
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

			rawMsgDispatch(w, r, &msgReq, rawXMLMsg, timestamp, agent)

		case "aes": // 兼容模式, 安全模式
			const signatureLen = sha1.Size * 2
			//if len(signature1) != signatureLen {
			//	err = fmt.Errorf("the length of signature mismatch, have: %d, want: %d", len(signature1), signatureLen)
			//	invalidRequestHandler.ServeInvalidRequest(w, r, err)
			//	return
			//}
			if len(msgSignature1) != signatureLen {
				err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: %d", len(msgSignature1), signatureLen)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			//signature2 := signature(agent.GetToken(), timestampStr, nonce)
			//if subtle.ConstantTimeCompare([]byte(signature1), []byte(signature2)) != 1 {
			//	invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("check signature failed"))
			//	return
			//}

			var requestHttpBody request.RequestHttpBody
			if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			wantToUserName := agent.GetId()
			if len(requestHttpBody.ToUserName) != len(wantToUserName) {
				err = fmt.Errorf("the message RequestHttpBody's ToUserName mismatch, have: %s, want: %s", requestHttpBody.ToUserName, wantToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
			if subtle.ConstantTimeCompare([]byte(requestHttpBody.ToUserName), []byte(wantToUserName)) != 1 {
				err = fmt.Errorf("the message RequestHttpBody's ToUserName mismatch, have: %s, want: %s", requestHttpBody.ToUserName, wantToUserName)
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

			AESKey := agent.GetCurrentAESKey()

			random, rawXMLMsg, err := aesDecryptMsg(EncryptedMsgBytes, agent.GetId(), AESKey)
			if err != nil {
				LastAESKey := agent.GetLastAESKey()
				if bytes.Equal(zeroAESKey[:], LastAESKey[:]) || bytes.Equal(AESKey[:], LastAESKey[:]) {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}

				AESKey = LastAESKey // !!!

				random, rawXMLMsg, err = aesDecryptMsg(EncryptedMsgBytes, agent.GetId(), AESKey)
				if err != nil {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}
			}

			var msgReq request.Request
			if err = xml.Unmarshal(rawXMLMsg, &msgReq); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			if requestHttpBody.ToUserName != msgReq.ToUserName {
				err = fmt.Errorf("the RequestHttpBody's ToUserName(==%d) mismatch the Request's ToUserName(==%d)", requestHttpBody.ToUserName, msgReq.ToUserName)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			aesMsgDispatch(w, r, &msgReq, rawXMLMsg, timestamp, nonce, AESKey, random, agent)

		default: // 未知的加密类型
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("unknown encrypt_type"))
			return
		}

	case "GET": // 首次验证 ======================================================
		signature1, timestamp, nonce, echostr, err := parseGetURLQuery(r.URL)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		const signatureLen = sha1.Size * 2
		if len(signature1) != signatureLen {
			err = fmt.Errorf("the length of signature mismatch, have: %d, want: %d", len(signature1), signatureLen)
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
