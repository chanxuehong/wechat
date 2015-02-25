// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package preview

const (
	MsgTypeText  = "text"
	MsgTypeImage = "image"
	MsgTypeVoice = "voice"
	MsgTypeVideo = "mpvideo"
	MsgTypeNews  = "mpnews"
)

type CommonMessageHeader struct {
	ToUser  string `json:"touser"`
	MsgType string `json:"msgtype"`
}

type Text struct {
	CommonMessageHeader
	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewText(touser, content string) *Text {
	var msg Text
	msg.MsgType = MsgTypeText
	msg.ToUser = touser
	msg.Text.Content = content
	return &msg
}

type Image struct {
	CommonMessageHeader
	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImage(touser, mediaId string) *Image {
	var msg Image
	msg.MsgType = MsgTypeImage
	msg.ToUser = touser
	msg.Image.MediaId = mediaId
	return &msg
}

type Voice struct {
	CommonMessageHeader
	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewVoice(touser, mediaId string) *Voice {
	var msg Voice
	msg.MsgType = MsgTypeVoice
	msg.ToUser = touser
	msg.Voice.MediaId = mediaId
	return &msg
}

type Video struct {
	CommonMessageHeader
	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

// 新建视频消息
//  NOTE: mediaId 应该通过 media.Client.MediaCreateVideo 得到
func NewVideo(touser, mediaId string) *Video {
	var msg Video
	msg.MsgType = MsgTypeVideo
	msg.ToUser = touser
	msg.Video.MediaId = mediaId
	return &msg
}

// 图文消息
type News struct {
	CommonMessageHeader
	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

// 新建图文消息
//  NOTE: mediaId 应该通过 media.Client.MediaCreateNews 得到
func NewNews(touser, mediaId string) *News {
	var msg News
	msg.MsgType = MsgTypeNews
	msg.ToUser = touser
	msg.News.MediaId = mediaId
	return &msg
}
