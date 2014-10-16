// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/wechat/message/passive/response"
	"io"
	"strconv"
)

// 把 text 回复消息 msg 写入 writer w
func (this *DefaultAgent) WriteAESText(w io.Writer, msg *response.Text, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return this.writeAESResponse(w, msg, timestamp, nonce, AESKey, random)
}

// 把 image 回复消息 msg 写入 writer w
func (this *DefaultAgent) WriteAESImage(w io.Writer, msg *response.Image, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return this.writeAESResponse(w, msg, timestamp, nonce, AESKey, random)
}

// 把 voice 回复消息 msg 写入 writer w
func (this *DefaultAgent) WriteAESVoice(w io.Writer, msg *response.Voice, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return this.writeAESResponse(w, msg, timestamp, nonce, AESKey, random)
}

// 把 video 回复消息 msg 写入 writer w
func (this *DefaultAgent) WriteAESVideo(w io.Writer, msg *response.Video, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return this.writeAESResponse(w, msg, timestamp, nonce, AESKey, random)
}

// 把 music 回复消息 msg 写入 writer w
func (this *DefaultAgent) WriteAESMusic(w io.Writer, msg *response.Music, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return this.writeAESResponse(w, msg, timestamp, nonce, AESKey, random)
}

// 把 news 回复消息 msg 写入 writer w
func (this *DefaultAgent) WriteAESNews(w io.Writer, msg *response.News, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) (err error) {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	if err = msg.CheckValid(); err != nil {
		return
	}
	return this.writeAESResponse(w, msg, timestamp, nonce, AESKey, random)
}

// 把 TransferToCustomerService 回复消息 msg 写入 writer w
func (this *DefaultAgent) WriteAESTransferToCustomerService(w io.Writer, msg *response.TransferToCustomerService, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return this.writeAESResponse(w, msg, timestamp, nonce, AESKey, random)
}

// 把 TransferToSpecialCustomerService 回复消息 msg 写入 writer w
func (this *DefaultAgent) WriteAESTransferToSpecialCustomerService(w io.Writer, msg *response.TransferToSpecialCustomerService, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return this.writeAESResponse(w, msg, timestamp, nonce, AESKey, random)
}

func (this *DefaultAgent) writeAESResponse(w io.Writer, msg interface{}, timestamp int64, nonce string, AESKey [32]byte, random [16]byte) (err error) {
	rawXMLMsg, err := xml.Marshal(msg)
	if err != nil {
		return
	}

	EncryptMsg := encryptMsg(random, rawXMLMsg, this.Id, AESKey)
	base64EncryptMsg := base64.StdEncoding.EncodeToString(EncryptMsg)

	var responseHttpBody response.ResponseHttpBody
	responseHttpBody.EncryptMsg = base64EncryptMsg
	responseHttpBody.TimeStamp = timestamp
	responseHttpBody.Nonce = nonce

	timestampStr := strconv.FormatInt(timestamp, 10)
	responseHttpBody.Signature = msgSignature(this.Token, timestampStr, nonce, base64EncryptMsg)

	return xml.NewEncoder(w).Encode(&responseHttpBody)
}
