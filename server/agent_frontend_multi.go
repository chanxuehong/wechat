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
	"sync"
)

// 定义回调 URL 上指定 Agent 的查询参数名
const URLQueryAgentKeyName = "agentkey"

// 多个 Agent 的前端, 负责处理 http 请求, net/http.Handler 的实现
//
//  NOTE:
//  MultiAgentFrontend 可以处理多个公众号的消息（事件），但是要求在回调 URL 上加上一个查询
//  参数，一般为 agentkey（参考常量 URLQueryAgentKeyName），这个参数的值就是 MultiAgentFrontend
//  索引 Agent 的 key。
//  例如回调 URL 为 http://www.xxx.com/weixin?agentkey=1234567890，那么就可以在后端调用
//
//    MultiAgentFrontend.SetAgent("1234567890", agent)
//
//  来增加一个 Agent 来处理 agentkey=1234567890 的消息（事件）。
//
//  MultiAgentFrontend 并发安全，可以在运行中动态增加和删除 Agent。
type MultiAgentFrontend struct {
	rwmutex               sync.RWMutex
	agentMap              map[string]Agent
	invalidRequestHandler InvalidRequestHandler
}

// 设置 InvalidRequestHandler, 如果 handler == nil 则使用默认的 DefaultInvalidRequestHandlerFunc
func (this *MultiAgentFrontend) SetInvalidRequestHandler(handler InvalidRequestHandler) {
	this.rwmutex.Lock()
	if handler == nil {
		this.invalidRequestHandler = InvalidRequestHandlerFunc(defaultInvalidRequestHandlerFunc)
	} else {
		this.invalidRequestHandler = handler
	}
	this.rwmutex.Unlock()
}

// 添加（设置） agentkey-agent pair, 如果 agent == nil 则不做任何操作
func (this *MultiAgentFrontend) SetAgent(agentkey string, agent Agent) {
	if agent == nil {
		return
	}

	this.rwmutex.Lock()
	if this.agentMap == nil {
		this.agentMap = make(map[string]Agent)
	}
	this.agentMap[agentkey] = agent
	this.rwmutex.Unlock()
}

// 删除 agentkey 对应的 Agent
func (this *MultiAgentFrontend) DeleteAgent(agentkey string) {
	this.rwmutex.Lock()
	delete(this.agentMap, agentkey)
	this.rwmutex.Unlock()
}

// 删除所有的 Agent
func (this *MultiAgentFrontend) DeleteAllAgent() {
	this.rwmutex.Lock()
	this.agentMap = nil
	this.rwmutex.Unlock()
}

func (this *MultiAgentFrontend) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	this.rwmutex.RLock()
	defer this.rwmutex.RUnlock()

	invalidRequestHandler := this.invalidRequestHandler
	if invalidRequestHandler == nil {
		invalidRequestHandler = InvalidRequestHandlerFunc(defaultInvalidRequestHandlerFunc)
	}
	if len(this.agentMap) == 0 {
		invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("no Agent"))
		return
	}

	switch r.Method {
	case "POST": // 处理从微信服务器推送过来的消息(事件) ==============================
		agentkey, signature1, timestampStr, nonce, encryptType, msgSignature1, err := parsePostURLQueryEx(r.URL)
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

			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			agent := this.agentMap[agentkey]
			if agent == nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("Not found Agent for %s == %s", URLQueryAgentKeyName, agentkey))
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

			timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			agent := this.agentMap[agentkey]
			if agent == nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("Not found Agent for %s == %s", URLQueryAgentKeyName, agentkey))
				return
			}

			var requestHttpBody request.RequestHttpBody
			if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			msgSignature2 := msgSignature(agent.GetToken(), timestampStr, nonce, requestHttpBody.EncryptMsg)
			if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
				invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("check signature failed"))
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

			EncryptMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptMsg)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			AESKey := agent.GetCurrentAESKey()

			random, rawXMLMsg, err := decryptMsg(EncryptMsgBytes, agent.GetId(), AESKey[:])
			if err != nil {
				LastAESKey := agent.GetLastAESKey()
				if bytes.Equal(zeroAESKey[:], LastAESKey[:]) || bytes.Equal(AESKey[:], LastAESKey[:]) {
					invalidRequestHandler.ServeInvalidRequest(w, r, err)
					return
				}

				AESKey = LastAESKey // !!!

				random, rawXMLMsg, err = decryptMsg(EncryptMsgBytes, agent.GetId(), AESKey[:])
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

			aesMsgDispatch(w, r, &msgReq, rawXMLMsg, timestamp, nonce, AESKey[:], random, agent)

		default: // 未知的加密类型
			invalidRequestHandler.ServeInvalidRequest(w, r, errors.New("unknown encrypt_type"))
			return
		}

	case "GET": // 首次验证 ======================================================
		agentkey, signature1, timestamp, nonce, echostr, err := parseGetURLQueryEx(r.URL)
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

		agent := this.agentMap[agentkey]
		if agent == nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, fmt.Errorf("Not found Agent for %s == %s", URLQueryAgentKeyName, agentkey))
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
