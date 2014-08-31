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
	n := len(head.ToUser)
	if n <= 0 {
		err = errors.New("用户列表是空的")
		return
	}
	if n > ToUserCountLimit {
		err = fmt.Errorf("用户列表的长度不能超过 %d, 现在为 %d", ToUserCountLimit, n)
		return
	}
	return
}

// 文本消息
type Text struct {
	CommonHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// 新建文本消息
func NewText(toUser []string, content string) *Text {
	var msg Text
	msg.ToUser = toUser
	msg.MsgType = MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// 图片消息
type Image struct {
	CommonHead

	Image struct {
		MediaId string `json:"media_id"` // mediaId 通过上传多媒体文件得到
	} `json:"image"`
}

// 图片消息
//  mediaId 通过上传多媒体文件得到
func NewImage(toUser []string, mediaId string) *Image {
	var msg Image
	msg.ToUser = toUser
	msg.MsgType = MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// 语音消息
type Voice struct {
	CommonHead

	Voice struct {
		MediaId string `json:"media_id"` // mediaId 通过上传多媒体文件得到
	} `json:"voice"`
}

// 新建语音消息
//  mediaId 通过上传多媒体文件得到
func NewVoice(toUser []string, mediaId string) *Voice {
	var msg Voice
	msg.ToUser = toUser
	msg.MsgType = MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// 视频消息
type Video struct {
	CommonHead

	Video struct {
		MediaId     string `json:"media_id"`              // NOTE: MediaId 应该通过 Client.MediaCreateVideo 得到
		Title       string `json:"title,omitempty"`       // 是否为多余？？？
		Description string `json:"description,omitempty"` // 是否为多余？？？
	} `json:"video"`
}

// 新建视频消息
//  NOTE:
//  MediaId 应该通过 Client.MediaCreateVideo 得到
//  title, description 可以为空, 是否为多余？？？
func NewVideo(toUser []string, mediaId, title, description string) *Video {
	var msg Video
	msg.ToUser = toUser
	msg.MsgType = MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}

// 图文消息
type News struct {
	CommonHead

	News struct {
		MediaId string `json:"media_id"` // NOTE: MediaId 应该通过 Client.MediaCreateNews 得到
	} `json:"mpnews"`
}

// 新建图文消息
//  NOTE: MediaId 应该通过 Client.MediaCreateNews 得到
func NewNews(toUser []string, mediaId string) *News {
	var msg News
	msg.ToUser = toUser
	msg.MsgType = MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}
