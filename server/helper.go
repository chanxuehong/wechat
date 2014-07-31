// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"encoding/xml"
	"errors"
	"github.com/chanxuehong/wechat/message/response"
	"io"
)

// 把 text 回复消息 msg 写入 writer w
func CopyText(w io.Writer, msg *response.Text) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return copyResponse(w, msg)
}

// 把 image 回复消息 msg 写入 writer w
func CopyImage(w io.Writer, msg *response.Image) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return copyResponse(w, msg)
}

// 把 voice 回复消息 msg 写入 writer w
func CopyVoice(w io.Writer, msg *response.Voice) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return copyResponse(w, msg)
}

// 把 video 回复消息 msg 写入 writer w
func CopyVideo(w io.Writer, msg *response.Video) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return copyResponse(w, msg)
}

// 把 music 回复消息 msg 写入 writer w
func CopyMusic(w io.Writer, msg *response.Music) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return copyResponse(w, msg)
}

// 把 news 回复消息 msg 写入 writer w
func CopyNews(w io.Writer, msg *response.News) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return copyResponse(w, msg)
}

// 把 TransferCustomerService 回复消息 msg 写入 writer w
func CopyTransferCustomerService(w io.Writer, msg *response.TransferCustomerService) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if msg == nil {
		return errors.New("msg == nil")
	}
	return copyResponse(w, msg)
}

func copyResponse(w io.Writer, msg interface{}) error {
	return xml.NewEncoder(w).Encode(msg) // 只要 w 能正常的写, 不会返回错误
}
