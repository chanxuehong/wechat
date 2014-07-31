// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package massbyopenid

import (
	"errors"
	"fmt"
)

type CommonHead struct {
	ToUser  []string `json:"touser,omitempty"` // 长度不能超过 ToUserCountLimit
	MsgType string   `json:"msgtype"`
}

// 检查 CommonHead 是否有效，有效返回 nil，否则返回错误信息
func (head *CommonHead) CheckValid() (err error) {
	touserNum := len(head.ToUser)
	if touserNum == 0 {
		err = errors.New("用户列表是空的")
		return
	}
	if touserNum > ToUserCountLimit {
		err = fmt.Errorf("用户列表的长度不能超过 %d, 现在为 %d", ToUserCountLimit, touserNum)
		return
	}
	return
}

// 文本消息
//
//  {
//      "touser": [
//          "oR5Gjjl_eiZoUpGozMo7dbBJ362A",
//          "oR5Gjjo5rXlMUocSEXKT7Q5RQ63Q"
//      ],
//      "msgtype": "text",
//      "text": {
//          "content": "hello from boxer."
//      }
//  }
type Text struct {
	CommonHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewText(touser []string, content string) *Text {
	var msg Text
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// 图片消息
//
//  {
//      "touser": [
//          "OPENID1",
//          "OPENID2"
//      ],
//      "msgtype": "image"
//      "image": {
//          "media_id": "BTgN0opcW3Y5zV_ZebbsD3NFKRWf6cb7OPswPi9Q83fOJHK2P67dzxn11Cp7THat"
//      },
//  }
type Image struct {
	CommonHead

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImage(touser []string, mediaId string) *Image {
	var msg Image
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// 语音消息
//
//  {
//      "touser": [
//          "OPENID1",
//          "OPENID2"
//      ],
//      "msgtype": "voice"
//      "voice": {
//          "media_id": "mLxl6paC7z2Tl-NJT64yzJve8T9c8u9K2x-Ai6Ujd4lIH9IBuF6-2r66mamn_gIT"
//      },
//  }
type Voice struct {
	CommonHead

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewVoice(touser []string, mediaId string) *Voice {
	var msg Voice
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// 视频消息
//  NOTE: MediaId 应该通过 Client.MediaCreateVideo 得到
//
//  {
//      "touser": [
//          "OPENID1",
//          "OPENID2"
//      ],
//      "msgtype": "video"
//      "video": {
//          "media_id": "123dsdajkasd231jhksad",
//          "title": "TITLE",
//          "description": "DESCRIPTION"
//      },
//  }
type Video struct {
	CommonHead

	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"video"`
}

// 创建视频消息
//  NOTE: MediaId 应该通过 Client.MediaCreateVideo 得到
//  title, description 可以为空
func NewVideo(touser []string, mediaId, title, description string) *Video {
	var msg Video
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}

// 图文消息
//  NOTE: MediaId 应该通过 Client.MediaCreateNews 得到
//
//  {
//      "touser": [
//          "OPENID1",
//          "OPENID2"
//      ],
//      "msgtype": "mpnews"
//      "mpnews": {
//          "media_id": "123dsdajkasd231jhksad"
//      },
//  }
type News struct {
	CommonHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

// 创建图文消息
//  NOTE: MediaId 应该通过 Client.MediaCreateNews 得到
func NewNews(touser []string, mediaId string) *News {
	var msg News
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}
