// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

package massbygroup

type CommonHead struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	MsgType string `json:"msgtype"`
}

// text ========================================================================

type Text struct {
	CommonHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewText(groupId, content string) *Text {
	var msg Text
	msg.Filter.GroupId = groupId
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

func NewImage(groupId, mediaId string) *Image {
	var msg Image
	msg.Filter.GroupId = groupId
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

func NewVoice(groupId, mediaId string) *Voice {
	var msg Voice
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// video =======================================================================

type Video struct {
	CommonHead

	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

func NewVideo(groupId, mediaId string) *Video {
	var msg Video
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId

	return &msg
}

// news ========================================================================

type News struct {
	CommonHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewNews(groupId, mediaId string) *News {
	var msg News
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}
