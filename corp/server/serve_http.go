// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"crypto/subtle"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chanxuehong/wechat/corp/message/passive/request"
	"github.com/chanxuehong/wechat/util"
)

// ServeHTTP 处理 http 消息请求
//  NOTE: 确保所有参数合法, r.Body 能正确读取数据
func ServeHTTP(w http.ResponseWriter, r *http.Request,
	urlValues url.Values, agent Agent, invalidRequestHandler InvalidRequestHandler) {

	switch r.Method {
	case "POST": // 消息处理
		msgSignature1, timestampStr, nonce, err := parsePostURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if len(msgSignature1) != 40 {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		timestamp, err := strconv.ParseInt(timestampStr, 10, 64)
		if err != nil {
			err = fmt.Errorf("can not parse timestamp(==%q) to int64, error: %s", timestampStr, err.Error())
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		var requestHttpBody request.RequestHttpBody
		if err := xml.NewDecoder(r.Body).Decode(&requestHttpBody); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		haveAgentCorpId := requestHttpBody.CorpId
		wantAgentCorpId := agent.GetCorpId()
		if len(haveAgentCorpId) != len(wantAgentCorpId) {
			err = fmt.Errorf("the message RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveAgentCorpId, wantAgentCorpId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		if subtle.ConstantTimeCompare([]byte(haveAgentCorpId), []byte(wantAgentCorpId)) != 1 {
			err = fmt.Errorf("the message RequestHttpBody's ToUserName mismatch, have: %s, want: %s", haveAgentCorpId, wantAgentCorpId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		haveAgentAgentId := requestHttpBody.AgentId
		wantAgentAgentId := agent.GetAgentId()
		if haveAgentAgentId != wantAgentAgentId && haveAgentAgentId != 0 {
			err = fmt.Errorf("the message RequestHttpBody's AgentId mismatch, have: %d, want: %d", haveAgentAgentId, wantAgentAgentId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		// 此时
		// 要么 haveAgentAgentId == agent.GetAgentId(),
		// 要么 haveAgentAgentId == 0

		msgSignature2 := util.MsgSign(agent.GetToken(), timestampStr, nonce, requestHttpBody.EncryptedMsg)
		if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(requestHttpBody.EncryptedMsg)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		random, rawXMLMsg, err := util.AESDecryptMsg(EncryptedMsgBytes, wantAgentCorpId, agent.GetAESKey())
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		var msgReq request.Request
		if err := xml.Unmarshal(rawXMLMsg, &msgReq); err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if haveAgentCorpId != msgReq.ToUserName {
			err = fmt.Errorf("the RequestHttpBody's ToUserName(==%s) mismatch the Request's ToUserName(==%s)", haveAgentCorpId, msgReq.ToUserName)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}
		if haveAgentAgentId != msgReq.AgentId {
			err = fmt.Errorf("the RequestHttpBody's AgentId(==%d) mismatch the Request's AgengId(==%d)", haveAgentAgentId, msgReq.AgentId)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if haveAgentAgentId != wantAgentAgentId { // 订阅/取消订阅 整个企业号, haveAgentAgentId == 0
			if msgReq.MsgType == request.MSG_TYPE_EVENT &&
				(msgReq.Event == request.EVENT_TYPE_SUBSCRIBE || msgReq.Event == request.EVENT_TYPE_UNSUBSCRIBE) {
				// do nothing
			} else {
				err = fmt.Errorf("the message Request's AgentId mismatch, have: %d, want: %d", haveAgentAgentId, wantAgentAgentId)
				invalidRequestHandler.ServeInvalidRequest(w, r, err)
				return
			}
		}

		msgDispatch(w, r, &msgReq, rawXMLMsg, timestamp, nonce, random, agent)

	case "GET": // 首次验证
		msgSignature1, timestamp, nonce, encryptedMsg, err := parseGetURLQuery(urlValues)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		if len(msgSignature1) != 40 {
			err = fmt.Errorf("the length of msg_signature mismatch, have: %d, want: 40", len(msgSignature1))
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		msgSignature2 := util.MsgSign(agent.GetToken(), timestamp, nonce, encryptedMsg)
		if subtle.ConstantTimeCompare([]byte(msgSignature1), []byte(msgSignature2)) != 1 {
			err = fmt.Errorf("check signature failed, input: %s, local: %s", msgSignature1, msgSignature2)
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		EncryptedMsgBytes, err := base64.StdEncoding.DecodeString(encryptedMsg)
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		_, echostr, err := util.AESDecryptMsg(EncryptedMsgBytes, agent.GetCorpId(), agent.GetAESKey())
		if err != nil {
			invalidRequestHandler.ServeInvalidRequest(w, r, err)
			return
		}

		w.Write(echostr)
	}
}
