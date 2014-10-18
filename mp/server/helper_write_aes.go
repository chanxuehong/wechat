// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/wechat/mp/message/passive/response"
	"io"
	"strconv"
)

// 把 text 回复消息 msg 写入 writer w
func WriteAESText(w io.Writer, msg *response.Text, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, AppId, Token string) error {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeAESResponse(w, msg, timestamp, nonce, AESKey, random, AppId, Token)
}

// 把 image 回复消息 msg 写入 writer w
func WriteAESImage(w io.Writer, msg *response.Image, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, AppId, Token string) error {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeAESResponse(w, msg, timestamp, nonce, AESKey, random, AppId, Token)
}

// 把 voice 回复消息 msg 写入 writer w
func WriteAESVoice(w io.Writer, msg *response.Voice, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, AppId, Token string) error {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeAESResponse(w, msg, timestamp, nonce, AESKey, random, AppId, Token)
}

// 把 video 回复消息 msg 写入 writer w
func WriteAESVideo(w io.Writer, msg *response.Video, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, AppId, Token string) error {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeAESResponse(w, msg, timestamp, nonce, AESKey, random, AppId, Token)
}

// 把 music 回复消息 msg 写入 writer w
func WriteAESMusic(w io.Writer, msg *response.Music, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, AppId, Token string) error {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeAESResponse(w, msg, timestamp, nonce, AESKey, random, AppId, Token)
}

// 把 news 回复消息 msg 写入 writer w
func WriteAESNews(w io.Writer, msg *response.News, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, AppId, Token string) (err error) {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return writeAESResponse(w, msg, timestamp, nonce, AESKey, random, AppId, Token)
}

// 把 TransferToCustomerService 回复消息 msg 写入 writer w
func WriteAESTransferToCustomerService(w io.Writer, msg *response.TransferToCustomerService,
	timestamp int64, nonce string, AESKey [32]byte, random []byte, AppId, Token string) error {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeAESResponse(w, msg, timestamp, nonce, AESKey, random, AppId, Token)
}

// 把 TransferToSpecialCustomerService 回复消息 msg 写入 writer w
func WriteAESTransferToSpecialCustomerService(w io.Writer, msg *response.TransferToSpecialCustomerService,
	timestamp int64, nonce string, AESKey [32]byte, random []byte, AppId, Token string) error {

	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return writeAESResponse(w, msg, timestamp, nonce, AESKey, random, AppId, Token)
}

func writeAESResponse(w io.Writer, msg interface{}, timestamp int64, nonce string,
	AESKey [32]byte, random []byte, AppId, Token string) (err error) {

	rawXMLMsg, err := xml.Marshal(msg)
	if err != nil {
		return
	}

	EncryptedMsg := aesEncryptMsg(random, rawXMLMsg, AppId, AESKey)
	base64EncryptedMsg := base64.StdEncoding.EncodeToString(EncryptedMsg)

	var responseHttpBody response.ResponseHttpBody
	responseHttpBody.EncryptedMsg = base64EncryptedMsg
	responseHttpBody.TimeStamp = timestamp
	responseHttpBody.Nonce = nonce

	timestampStr := strconv.FormatInt(timestamp, 10)
	responseHttpBody.MsgSignature = msgSignature(Token, timestampStr, nonce, base64EncryptedMsg)

	return xml.NewEncoder(w).Encode(&responseHttpBody)
}
