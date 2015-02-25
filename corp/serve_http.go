// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package corp

import (
	"bytes"
	"crypto/subtle"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chanxuehong/wechat/util"
)

var zeroAESKey [32]byte

// 微信服务器请求 http body
type RequestHttpBody struct {
	XMLName      struct{} `xml:"xml" json:"-"`
	CorpId       string   `xml:"ToUserName"`
	AgentId      int64    `xml:"AgentID"`
	EncryptedMsg string   `xml:"Encrypt"`
}

// ServeHTTP 处理 http 消息请求
//  NOTE: 调用者保证所有参数有效
func ServeHTTP(w http.ResponseWriter, r *http.Request, urlValues url.Values,
	agentServer AgentServer, invalidRequestHandler InvalidRequestHandler) {

	switch r.Method {
	case "POST": // 消息处理
		msgSignature1, timestampStr, nonce, err := parsePostURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 首先判断签名长度是否合法
		if len(msgSignature1) != 40 {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			err = errors.New("can not parse timestamp to int64: " + timestampStr)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 解析 RequestHttpBody
		var requestHttpBody RequestHttpBody
		if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		haveCorpId := requestHttpBody.CorpId
		wantCorpId := agentServer.CorpId()
		if len(haveCorpId) != len(wantCorpId) {
			err = fmt.Errorf("the RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveCorpId, wantCorpId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		if subtle.ConstantTimeCompare([]byte(haveCorpId), []byte(wantCorpId)) != 1 {
			err = fmt.Errorf("the RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveCorpId, wantCorpId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		haveAgentId := requestHttpBody.AgentId
		wantAgentId := agentServer.AgentId()
		if haveAgentId != wantAgentId && haveAgentId != 0 {
			err = fmt.Errorf("the RequestHttpBody's AgentId mismatch, have: %d, want: %d", haveAgentId, wantAgentId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 此时
		// 要么 haveAgentId == wantAgentId,
		// 要么 haveAgentId == 0

		agentToken := agentServer.Token()

		// 验证签名
		msgSignature2 := util.MsgSign(agentToken, timestampStr, nonce, requestHttpBody.EncryptedMsg)
		if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 解密
		EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptedMsg)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		AESKey := agentServer.CurrentAESKey()
		Random, RawMsgXML, err := util.AESDecryptMsg(EncryptedMsgBytes, wantCorpId, AESKey)
		if err != nil {
			// 尝试用上一次的 AESKey 来解密
			LastAESKey := agentServer.LastAESKey()
			if bytes.Equal(AESKey[:], LastAESKey[:]) || bytes.Equal(zeroAESKey[:], LastAESKey[:]) {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}

			AESKey = LastAESKey // NOTE
			Random, RawMsgXML, err = util.AESDecryptMsg(EncryptedMsgBytes, wantCorpId, AESKey)
			if err != nil {
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
		}

		// 解密成功, 解析 MixedMessage
		var MixedMsg MixedMessage
		if err = xml.Unmarshal(RawMsgXML, &MixedMsg); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 安全考虑再次验证
		if haveCorpId != MixedMsg.ToUserName {
			err = fmt.Errorf("the RequestHttpBody's ToUserName(==%s) mismatch the MixedMessage's ToUserName(==%s)", haveCorpId, MixedMsg.ToUserName)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		if haveAgentId != MixedMsg.AgentId {
			err = fmt.Errorf("the RequestHttpBody's AgentId(==%d) mismatch the MixedMessage's AgengId(==%d)", haveAgentId, MixedMsg.AgentId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 如果是订阅/取消订阅 整个企业号, haveAgentId == 0
		if haveAgentId != wantAgentId {
			if MixedMsg.MsgType == "event" &&
				(MixedMsg.Event == "subscribe" || MixedMsg.Event == "unsubscribe") {
				// do nothing
			} else {
				err = fmt.Errorf("the RequestHttpBody's AgentId mismatch, have: %d, want: %d", haveAgentId, wantAgentId)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
		}

		// 成功, 交给 MessageHandler
		r := &Request{
			HttpRequest: r,

			MsgSignature: msgSignature1,
			TimeStamp:    timestamp,
			Nonce:        nonce,
			RawMsgXML:    RawMsgXML,
			MixedMsg:     &MixedMsg,

			AESKey: AESKey,
			Random: Random,

			CorpId:     wantCorpId,
			AgentId:    wantAgentId,
			AgentToken: agentToken,
		}
		agentServer.MessageHandler().ServeMessage(w, r)

	case "GET": // 首次验证
		msgSignature1, timestamp, nonce, encryptedMsg, err := parseGetURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 验证签名
		if len(msgSignature1) != 40 {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msgSignature2 := util.MsgSign(agentServer.Token(), timestamp, nonce, encryptedMsg)
		if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 解密
		EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(encryptedMsg)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		CorpId := agentServer.CorpId()
		AESKey := agentServer.CurrentAESKey()
		_, echostr, err := util.AESDecryptMsg(EncryptedMsgBytes, CorpId, AESKey)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		w.Write(echostr)
	}
}
