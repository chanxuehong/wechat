// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package response

import (
	"time"
)

type CommonHead struct {
	ToUserName   string `xml:"ToUserName"`   // 接收方帐号(OpenID)
	FromUserName string `xml:"FromUserName"` // 开发者微信号
	CreateTime   int64  `xml:"CreateTime"`   // 消息创建时间(整型), unixtime
	MsgType      string `xml:"MsgType"`      // text, image, voice, video, music, news, transfer_customer_service
}

// text ========================================================================

type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Content string `xml:"Content"` // 回复的消息内容(换行：在content中能够换行, 微信客户端就支持换行显示)
}

func NewText(to, from, content string) *Text {
	msg := Text{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_TEXT,
		},
	}
	msg.Content = content

	return &msg
}

// image =======================================================================

type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Image struct {
		MediaId string `xml:"MediaId"` // 通过上传多媒体文件, 得到的id
	} `xml:"Image"`
}

func NewImage(to, from, mediaId string) *Image {
	msg := Image{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_IMAGE,
		},
	}
	msg.Image.MediaId = mediaId

	return &msg
}

// voice =======================================================================

type Voice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Voice struct {
		MediaId string `xml:"MediaId"` // 通过上传多媒体文件, 得到的id
	} `xml:"Voice"`
}

func NewVoice(to, from, mediaId string) *Voice {
	msg := Voice{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_VOICE,
		},
	}
	msg.Voice.MediaId = mediaId

	return &msg
}

// video =======================================================================

type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Video struct {
		MediaId     string `xml:"MediaId"`               // 通过上传多媒体文件, 得到的id
		Title       string `xml:"Title,omitempty"`       // 视频消息的标题
		Description string `xml:"Description,omitempty"` // 视频消息的描述
	} `xml:"Video"`
}

// title, description 可以为 ""
func NewVideo(to, from, mediaId, title, description string) *Video {
	msg := Video{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_VIDEO,
		},
	}
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}

// music =======================================================================

type Music struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Music struct {
		Title        string `xml:"Title,omitempty"`       // 音乐标题
		Description  string `xml:"Description,omitempty"` // 音乐描述
		MusicURL     string `xml:"MusicUrl"`              // 音乐链接
		HQMusicURL   string `xml:"HQMusicUrl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `xml:"ThumbMediaId"`          // 缩略图的媒体id, 通过上传多媒体文件, 得到的id
	} `xml:"Music"`
}

// title, description 可以为 ""
func NewMusic(to, from, thumbMediaId, musicURL, HQMusicURL, title, description string) *Music {
	msg := Music{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_MUSIC,
		},
	}
	msg.Music.ThumbMediaId = thumbMediaId
	msg.Music.MusicURL = musicURL
	msg.Music.HQMusicURL = HQMusicURL
	msg.Music.Title = title
	msg.Music.Description = description

	return &msg
}

// news ========================================================================

// 图文消息里的 item
type NewsArticle struct {
	Title       string `xml:"Title,omitempty"`       // 图文消息标题
	Description string `xml:"Description,omitempty"` // 图文消息描述
	PicURL      string `xml:"PicUrl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         string `xml:"Url,omitempty"`         // 点击图文消息跳转链接
}

// 图文消息
type News struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	ArticleCount int            `xml:"ArticleCount"`            // 图文消息个数, 限制为10条以内
	Articles     []*NewsArticle `xml:"Articles>item,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

// NOTE: 如果图文消息数量大于微信的限制, 则把多余的截除.
func NewNews(to, from string, article ...*NewsArticle) *News {
	if len(article) > NewsArticleCountLimit {
		article = article[:NewsArticleCountLimit]
	} else if article == nil {
		article = make([]*NewsArticle, 0, NewsArticleCountLimit)
	}

	msg := News{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_NEWS,
		},
	}
	msg.ArticleCount = len(article)
	msg.Articles = article

	return &msg
}

// NOTE: 如果总的按钮数超过限制, 则截除多余的.
func (msg *News) AppendArticle(article ...*NewsArticle) {
	if len(article) <= 0 {
		return
	}

	switch n := NewsArticleCountLimit - len(msg.Articles); {
	case n > 0:
		if len(article) > n {
			article = article[:n]
		}
		msg.Articles = append(msg.Articles, article...)
		msg.ArticleCount = len(msg.Articles)
	case n == 0:
	default: // n < 0
		msg.Articles = msg.Articles[:NewsArticleCountLimit]
		msg.ArticleCount = NewsArticleCountLimit
	}
}

// transfer_customer_service ===================================================

// 将消息转发到多客服
type TransferCustomerService struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead
}

func NewTransferCustomerService(to, from string) *TransferCustomerService {
	return &TransferCustomerService{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_TRANSFER_CUSTOMER_SERVICE,
		},
	}
}
