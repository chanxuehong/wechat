// 被动回复的基本消息的数据结构定义, 更多的消息定义在对应的业务模块内.
package response

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

const (
	MsgTypeText                    core.MsgType = "text"                      // 文本消息
	MsgTypeImage                   core.MsgType = "image"                     // 图片消息
	MsgTypeVoice                   core.MsgType = "voice"                     // 语音消息
	MsgTypeVideo                   core.MsgType = "video"                     // 视频消息
	MsgTypeMusic                   core.MsgType = "music"                     // 音乐消息
	MsgTypeNews                    core.MsgType = "news"                      // 图文消息
	MsgTypeTransferCustomerService core.MsgType = "transfer_customer_service" // 将消息转发到多客服
)

// 文本消息
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	Content string `xml:"Content" json:"Content"` // 回复的消息内容(换行: 在content中能够换行, 微信客户端支持换行显示)
}

func NewText(to, from string, timestamp int64, content string) (text *Text) {
	return &Text{
		MsgHeader: core.MsgHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeText,
		},
		Content: content,
	}
}

// 图片消息
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	Image struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Image" json:"Image"`
}

func NewImage(to, from string, timestamp int64, mediaId string) (image *Image) {
	image = &Image{
		MsgHeader: core.MsgHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeImage,
		},
	}
	image.Image.MediaId = mediaId
	return
}

// 语音消息
type Voice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	Voice struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `xml:"Voice" json:"Voice"`
}

func NewVoice(to, from string, timestamp int64, mediaId string) (voice *Voice) {
	voice = &Voice{
		MsgHeader: core.MsgHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeVoice,
		},
	}
	voice.Voice.MediaId = mediaId
	return
}

// 视频消息
type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	Video struct {
		MediaId     string `xml:"MediaId"               json:"MediaId"`               // 通过素材管理接口上传多媒体文件得到 MediaId
		Title       string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 视频消息的标题, 可以为空
		Description string `xml:"Description,omitempty" json:"Description,omitempty"` // 视频消息的描述, 可以为空
	} `xml:"Video" json:"Video"`
}

func NewVideo(to, from string, timestamp int64, mediaId, title, description string) (video *Video) {
	video = &Video{
		MsgHeader: core.MsgHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeVideo,
		},
	}
	video.Video.MediaId = mediaId
	video.Video.Title = title
	video.Video.Description = description
	return
}

// 音乐消息
type Music struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	Music struct {
		Title        string `xml:"Title,omitempty"        json:"Title,omitempty"`       // 音乐标题
		Description  string `xml:"Description,omitempty"  json:"Description,omitempty"` // 音乐描述
		MusicURL     string `xml:"MusicUrl"               json:"MusicUrl"`              // 音乐链接
		HQMusicURL   string `xml:"HQMusicUrl"             json:"HQMusicUrl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `xml:"ThumbMediaId"           json:"ThumbMediaId"`          // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
	} `xml:"Music" json:"Music"`
}

func NewMusic(to, from string, timestamp int64, thumbMediaId, musicURL, HQMusicURL, title, description string) (music *Music) {
	music = &Music{
		MsgHeader: core.MsgHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeMusic,
		},
	}
	music.Music.Title = title
	music.Music.Description = description
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = HQMusicURL
	music.Music.ThumbMediaId = thumbMediaId
	return
}

// 图文消息里的 Article
type Article struct {
	Title       string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 图文消息标题
	Description string `xml:"Description,omitempty" json:"Description,omitempty"` // 图文消息描述
	PicURL      string `xml:"PicUrl,omitempty"      json:"PicUrl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         string `xml:"Url,omitempty"         json:"Url,omitempty"`         // 点击图文消息跳转链接
}

// 图文消息
type News struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	ArticleCount int       `xml:"ArticleCount"            json:"ArticleCount"`       // 图文消息个数, 限制为10条以内
	Articles     []Article `xml:"Articles>item,omitempty" json:"Articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

func NewNews(to, from string, timestamp int64, articles []Article) (news *News) {
	news = &News{
		MsgHeader: core.MsgHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeNews,
		},
	}
	news.ArticleCount = len(articles)
	news.Articles = articles
	return
}

// 将消息转发到多客服, 参见多客服模块
type TransferToCustomerService struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	TransInfo *TransInfo `xml:"TransInfo,omitempty" json:"TransInfo,omitempty"`
}

type TransInfo struct {
	KfAccount string `xml:"KfAccount" json:"KfAccount"`
}

// 如果不指定客服则 kfAccount 留空.
func NewTransferToCustomerService(to, from string, timestamp int64, kfAccount string) (msg *TransferToCustomerService) {
	msg = &TransferToCustomerService{
		MsgHeader: core.MsgHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeTransferCustomerService,
		},
	}
	if kfAccount != "" {
		msg.TransInfo = &TransInfo{
			KfAccount: kfAccount,
		}
	}
	return
}
