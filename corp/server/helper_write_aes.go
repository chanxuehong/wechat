// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"io"
	"strconv"

	"github.com/chanxuehong/wechat/corp/message/passive/response"
	"github.com/chanxuehong/wechat/util"
)

// 把 text 回复消息 msg 写入 writer w
func WriteText(w io.Writer, msg *response.Text, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, CorpId, Token string) (err error) {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeResponse(w, msg, timestamp, nonce, AESKey, random, CorpId, Token)
}

// 把 image 回复消息 msg 写入 writer w
func WriteImage(w io.Writer, msg *response.Image, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, CorpId, Token string) (err error) {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeResponse(w, msg, timestamp, nonce, AESKey, random, CorpId, Token)
}

// 把 voice 回复消息 msg 写入 writer w
func WriteVoice(w io.Writer, msg *response.Voice, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, CorpId, Token string) (err error) {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeResponse(w, msg, timestamp, nonce, AESKey, random, CorpId, Token)
}

// 把 video 回复消息 msg 写入 writer w
func WriteVideo(w io.Writer, msg *response.Video, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, CorpId, Token string) (err error) {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeResponse(w, msg, timestamp, nonce, AESKey, random, CorpId, Token)
}

// 把 news 回复消息 msg 写入 writer w
func WriteNews(w io.Writer, msg *response.News, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, CorpId, Token string) (err error) {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return writeResponse(w, msg, timestamp, nonce, AESKey, random, CorpId, Token)
}

func writeResponse(w io.Writer, msg interface{}, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, CorpId, Token string) (err error) {

	rawXMLMsg, err := xml.Marshal(msg)
	if err != nil {
		return
	}

	EncryptedMsg := util.AESEncryptMsg(random, rawXMLMsg, CorpId, AESKey)
	base64EncryptedMsg := base64.StdEncoding.EncodeToString(EncryptedMsg)

	var responseHttpBody response.ResponseHttpBody
	responseHttpBody.EncryptedMsg = base64EncryptedMsg
	responseHttpBody.TimeStamp = timestamp
	responseHttpBody.Nonce = nonce

	timestampStr := strconv.FormatInt(timestamp, 10)
	responseHttpBody.MsgSignature = util.MsgSign(Token, timestampStr, nonce, base64EncryptedMsg)

	return xml.NewEncoder(w).Encode(&responseHttpBody)
}
