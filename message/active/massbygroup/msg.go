// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package massbygroup

type CommonHead struct {
	Filter struct {
		GroupId int64 `json:"group_id,string"`
	} `json:"filter"`
	MsgType string `json:"msgtype"`
}

// 文本消息
type Text struct {
	CommonHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

// 新建文本消息
func NewText(groupId int64, content string) *Text {
	var msg Text
	msg.Filter.GroupId = groupId
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

// 新建图片消息
//  mediaId 通过上传多媒体文件得到
func NewImage(groupId int64, mediaId string) *Image {
	var msg Image
	msg.Filter.GroupId = groupId
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
func NewVoice(groupId int64, mediaId string) *Voice {
	var msg Voice
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// 视频消息
type Video struct {
	CommonHead

	Video struct {
		MediaId string `json:"media_id"` // NOTE: mediaId 应该通过 Client.MediaCreateVideo 得到
	} `json:"mpvideo"`
}

// 新建视频消息
//  NOTE: mediaId 应该通过 Client.MediaCreateVideo 得到
func NewVideo(groupId int64, mediaId string) *Video {
	var msg Video
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId

	return &msg
}

// 图文消息
type News struct {
	CommonHead

	News struct {
		MediaId string `json:"media_id"` // NOTE: mediaId 应该通过 Client.MediaCreateNews 得到
	} `json:"mpnews"`
}

// 新建图文消息
//  NOTE: mediaId 应该通过 Client.MediaCreateNews 得到
func NewNews(groupId int64, mediaId string) *News {
	var msg News
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}
