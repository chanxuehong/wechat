// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"strconv"

	"github.com/chanxuehong/wechat/mp/message/passive/response"
	"github.com/chanxuehong/wechat/util"
)

// 把 text 回复消息 msg 写入 writer w
func WriteAESText(msg *response.Text, para *AESRequestParameters, appId, token string) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	if para == nil {
		return errors.New("para == nil")
	}
	if para.HTTPResponseWriter == nil || para.HTTPRequest == nil {
		return errors.New("para is invalid")
	}

	return writeAESResponse(msg, para, appId, token)
}

// 把 image 回复消息 msg 写入 writer w
func WriteAESImage(msg *response.Image, para *AESRequestParameters, appId, token string) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	if para == nil {
		return errors.New("para == nil")
	}
	if para.HTTPResponseWriter == nil || para.HTTPRequest == nil {
		return errors.New("para is invalid")
	}

	return writeAESResponse(msg, para, appId, token)
}

// 把 voice 回复消息 msg 写入 writer w
func WriteAESVoice(msg *response.Voice, para *AESRequestParameters, appId, token string) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	if para == nil {
		return errors.New("para == nil")
	}
	if para.HTTPResponseWriter == nil || para.HTTPRequest == nil {
		return errors.New("para is invalid")
	}

	return writeAESResponse(msg, para, appId, token)
}

// 把 video 回复消息 msg 写入 writer w
func WriteAESVideo(msg *response.Video, para *AESRequestParameters, appId, token string) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	if para == nil {
		return errors.New("para == nil")
	}
	if para.HTTPResponseWriter == nil || para.HTTPRequest == nil {
		return errors.New("para is invalid")
	}

	return writeAESResponse(msg, para, appId, token)
}

// 把 music 回复消息 msg 写入 writer w
func WriteAESMusic(msg *response.Music, para *AESRequestParameters, appId, token string) error {
	if msg == nil {
		return errors.New("msg == nil")
	}
	if para == nil {
		return errors.New("para == nil")
	}
	if para.HTTPResponseWriter == nil || para.HTTPRequest == nil {
		return errors.New("para is invalid")
	}

	return writeAESResponse(msg, para, appId, token)
}

// 把 news 回复消息 msg 写入 writer w
func WriteAESNews(msg *response.News, para *AESRequestParameters, appId, token string) (err error) {
	if msg == nil {
		return errors.New("msg == nil")
	}
	if para == nil {
		return errors.New("para == nil")
	}
	if para.HTTPResponseWriter == nil || para.HTTPRequest == nil {
		return errors.New("para is invalid")
	}

	return writeAESResponse(msg, para, appId, token)
}

// 把 TransferToCustomerService 回复消息 msg 写入 writer w
func WriteAESTransferToCustomerService(msg *response.TransferToCustomerService,
	para *AESRequestParameters, appId, token string) error {

	if msg == nil {
		return errors.New("msg == nil")
	}
	if para == nil {
		return errors.New("para == nil")
	}
	if para.HTTPResponseWriter == nil || para.HTTPRequest == nil {
		return errors.New("para is invalid")
	}

	return writeAESResponse(msg, para, appId, token)
}

// 把 TransferToSpecialCustomerService 回复消息 msg 写入 writer w
func WriteAESTransferToSpecialCustomerService(msg *response.TransferToSpecialCustomerService,
	para *AESRequestParameters, appId, token string) error {

	if msg == nil {
		return errors.New("msg == nil")
	}
	if para == nil {
		return errors.New("para == nil")
	}
	if para.HTTPResponseWriter == nil || para.HTTPRequest == nil {
		return errors.New("para is invalid")
	}

	return writeAESResponse(msg, para, appId, token)
}

func writeAESResponse(msg interface{}, para *AESRequestParameters, appId, token string) (err error) {

	rawXMLMsg, err := xml.Marshal(msg)
	if err != nil {
		return
	}

	EncryptedMsg := util.AESEncryptMsg(para.Random, rawXMLMsg, appId, para.AESKey)
	base64EncryptedMsg := base64.StdEncoding.EncodeToString(EncryptedMsg)

	var responseHttpBody response.ResponseHttpBody
	responseHttpBody.EncryptedMsg = base64EncryptedMsg
	responseHttpBody.TimeStamp = para.Timestamp
	responseHttpBody.Nonce = para.Nonce

	timestampStr := strconv.FormatInt(para.Timestamp, 10)
	responseHttpBody.MsgSignature = util.MsgSign(token, timestampStr, para.Nonce, base64EncryptedMsg)

	return xml.NewEncoder(para.HTTPResponseWriter).Encode(&responseHttpBody)
}
