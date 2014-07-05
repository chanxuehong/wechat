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

// text ========================================================================

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

// image =======================================================================

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

// voice =======================================================================

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

// video =======================================================================

type Video struct {
	CommonHead

	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"video"`
}

// title, description 可以为空
func NewVideo(touser []string, mediaId, title, description string) *Video {
	var msg Video
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}

// news ========================================================================

type News struct {
	CommonHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewNews(touser []string, mediaId string) *News {
	var msg News
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}
