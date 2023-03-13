// 被动回复的基本消息的数据结构定义, 更多的消息定义在对应的业务模块内.
package response

import (
	"github.com/bububa/wechat/work/message"
)

const (
	MsgTypeText  message.MsgType = "text"  // 文本消息
	MsgTypeImage message.MsgType = "image" // 图片消息
	MsgTypeVoice message.MsgType = "voice" // 语音消息
	MsgTypeVideo message.MsgType = "video" // 视频消息
	MsgTypeNews  message.MsgType = "news"  // 图文消息
)

// 文本消息
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	Content message.CDATA `xml:"Content" json:"Content"` // 回复的消息内容(换行: 在content中能够换行, 微信客户端支持换行显示)
}

func NewText(to, from string, timestamp int64, content string) (text *Text) {
	return &Text{
		MsgHeader: message.MsgHeader{
			ToUserName:   message.CDATA(to),
			FromUserName: message.CDATA(from),
			CreateTime:   timestamp,
			MsgType:      MsgTypeText,
		},
		Content: message.CDATA(content),
	}
}

// 图片消息
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	Image struct {
		MediaId message.CDATA `xml:"MediaId" json:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Image" json:"Image"`
}

func NewImage(to, from string, timestamp int64, mediaId string) (image *Image) {
	image = &Image{
		MsgHeader: message.MsgHeader{
			ToUserName:   message.CDATA(to),
			FromUserName: message.CDATA(from),
			CreateTime:   timestamp,
			MsgType:      MsgTypeImage,
		},
	}
	image.Image.MediaId = message.CDATA(mediaId)
	return
}

// 语音消息
type Voice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	Voice struct {
		MediaId message.CDATA `xml:"MediaId" json:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Voice" json:"Voice"`
}

func NewVoice(to, from string, timestamp int64, mediaId string) (voice *Voice) {
	voice = &Voice{
		MsgHeader: message.MsgHeader{
			ToUserName:   message.CDATA(to),
			FromUserName: message.CDATA(from),
			CreateTime:   timestamp,
			MsgType:      MsgTypeVoice,
		},
	}
	voice.Voice.MediaId = message.CDATA(mediaId)
	return
}

// 视频消息
type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	Video struct {
		MediaId     message.CDATA `xml:"MediaId"               json:"MediaId"`               // 通过素材管理接口上传多媒体文件得到 MediaId
		Title       message.CDATA `xml:"Title,omitempty"       json:"Title,omitempty"`       // 视频消息的标题, 可以为空
		Description message.CDATA `xml:"Description,omitempty" json:"Description,omitempty"` // 视频消息的描述, 可以为空
	} `xml:"Video" json:"Video"`
}

func NewVideo(to, from string, timestamp int64, mediaId, title, description string) (video *Video) {
	video = &Video{
		MsgHeader: message.MsgHeader{
			ToUserName:   message.CDATA(to),
			FromUserName: message.CDATA(from),
			CreateTime:   timestamp,
			MsgType:      MsgTypeVideo,
		},
	}
	video.Video.MediaId = message.CDATA(mediaId)
	video.Video.Title = message.CDATA(title)
	video.Video.Description = message.CDATA(description)
	return
}

// 图文消息里的 Article
type Article struct {
	Title       message.CDATA `xml:"Title,omitempty"       json:"Title,omitempty"`       // 图文消息标题
	Description message.CDATA `xml:"Description,omitempty" json:"Description,omitempty"` // 图文消息描述
	PicURL      message.CDATA `xml:"PicUrl,omitempty"      json:"PicUrl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         message.CDATA `xml:"Url,omitempty"         json:"Url,omitempty"`         // 点击图文消息跳转链接
}

// 图文消息
type News struct {
	XMLName struct{} `xml:"xml" json:"-"`
	message.MsgHeader
	ArticleCount int       `xml:"ArticleCount"            json:"ArticleCount"`       // 图文消息个数, 限制为10条以内
	Articles     []Article `xml:"Articles>item,omitempty" json:"Articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

func NewNews(to, from string, timestamp int64, articles []Article) (news *News) {
	news = &News{
		MsgHeader: message.MsgHeader{
			ToUserName:   message.CDATA(to),
			FromUserName: message.CDATA(from),
			CreateTime:   timestamp,
			MsgType:      MsgTypeNews,
		},
	}
	news.ArticleCount = len(articles)
	news.Articles = articles
	return
}
