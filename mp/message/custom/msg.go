<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package custom

import (
	"errors"
	"fmt"
)

const (
	MsgTypeText   = "text"   // 文本消息
	MsgTypeImage  = "image"  // 图片消息
	MsgTypeVoice  = "voice"  // 语音消息
	MsgTypeVideo  = "video"  // 视频消息
	MsgTypeMusic  = "music"  // 音乐消息
	MsgTypeNews   = "news"   // 图文消息
	MsgTypeWxCard = "wxcard" // 卡卷消息
)

type MessageHeader struct {
	ToUser  string `json:"touser"` // 接收方 OpenID
	MsgType string `json:"msgtype"`
}

// 如果需要以某个客服帐号来发消息(在微信6.0.2及以上版本中显示自定义头像),
// 则需在JSON数据包的后半部分加入 customservice 参数
=======
package custom

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	MsgTypeText   core.MsgType = "text"   // 文本消息
	MsgTypeImage  core.MsgType = "image"  // 图片消息
	MsgTypeVoice  core.MsgType = "voice"  // 语音消息
	MsgTypeVideo  core.MsgType = "video"  // 视频消息
	MsgTypeMusic  core.MsgType = "music"  // 音乐消息
	MsgTypeNews   core.MsgType = "news"   // 图文消息
	MsgTypeMPNews core.MsgType = "mpnews" // 图文消息, 发送已经创建好的图文
	MsgTypeWxCard core.MsgType = "wxcard" // 卡卷消息
)

type MsgHeader struct {
	ToUser  string       `json:"touser"` // 接收方 OpenID
	MsgType core.MsgType `json:"msgtype"`
}

>>>>>>> github/v2
type CustomService struct {
	KfAccount string `json:"kf_account"`
}

// 文本消息
type Text struct {
<<<<<<< HEAD
	MessageHeader

	Text struct {
		Content string `json:"content"` // 支持换行符
	} `json:"text"`

	*CustomService `json:"customservice,omitempty"`
=======
	MsgHeader
	Text struct {
		Content string `json:"content"` // 支持换行符
	} `json:"text"`
	CustomService *CustomService `json:"customservice,omitempty"`
>>>>>>> github/v2
}

// 新建文本消息.
//  如果不指定客服则 kfAccount 留空.
func NewText(toUser, content, kfAccount string) (text *Text) {
	text = &Text{
<<<<<<< HEAD
		MessageHeader: MessageHeader{
=======
		MsgHeader: MsgHeader{
>>>>>>> github/v2
			ToUser:  toUser,
			MsgType: MsgTypeText,
		},
	}
	text.Text.Content = content

	if kfAccount != "" {
		text.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}

// 图片消息
type Image struct {
<<<<<<< HEAD
	MessageHeader

	Image struct {
		MediaId string `json:"media_id"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `json:"image"`

	*CustomService `json:"customservice,omitempty"`
=======
	MsgHeader
	Image struct {
		MediaId string `json:"media_id"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `json:"image"`
	CustomService *CustomService `json:"customservice,omitempty"`
>>>>>>> github/v2
}

// 新建图片消息.
//  如果不指定客服则 kfAccount 留空.
func NewImage(toUser, mediaId, kfAccount string) (image *Image) {
	image = &Image{
<<<<<<< HEAD
		MessageHeader: MessageHeader{
=======
		MsgHeader: MsgHeader{
>>>>>>> github/v2
			ToUser:  toUser,
			MsgType: MsgTypeImage,
		},
	}
	image.Image.MediaId = mediaId

	if kfAccount != "" {
		image.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}

// 语音消息
type Voice struct {
<<<<<<< HEAD
	MessageHeader

	Voice struct {
		MediaId string `json:"media_id"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `json:"voice"`

	*CustomService `json:"customservice,omitempty"`
=======
	MsgHeader
	Voice struct {
		MediaId string `json:"media_id"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `json:"voice"`
	CustomService *CustomService `json:"customservice,omitempty"`
>>>>>>> github/v2
}

// 新建语音消息.
//  如果不指定客服则 kfAccount 留空.
func NewVoice(toUser, mediaId, kfAccount string) (voice *Voice) {
	voice = &Voice{
<<<<<<< HEAD
		MessageHeader: MessageHeader{
=======
		MsgHeader: MsgHeader{
>>>>>>> github/v2
			ToUser:  toUser,
			MsgType: MsgTypeVoice,
		},
	}
	voice.Voice.MediaId = mediaId

	if kfAccount != "" {
		voice.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}

// 视频消息
type Video struct {
<<<<<<< HEAD
	MessageHeader

=======
	MsgHeader
>>>>>>> github/v2
	Video struct {
		MediaId      string `json:"media_id"`              // 通过素材管理接口上传多媒体文件得到 MediaId
		ThumbMediaId string `json:"thumb_media_id"`        // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
		Title        string `json:"title,omitempty"`       // 视频消息的标题, 可以为 ""
		Description  string `json:"description,omitempty"` // 视频消息的描述, 可以为 ""
	} `json:"video"`
<<<<<<< HEAD

	*CustomService `json:"customservice,omitempty"`
=======
	CustomService *CustomService `json:"customservice,omitempty"`
>>>>>>> github/v2
}

// 新建视频消息.
//  如果不指定客服则 kfAccount 留空.
func NewVideo(toUser, mediaId, thumbMediaId, title, description, kfAccount string) (video *Video) {
	video = &Video{
<<<<<<< HEAD
		MessageHeader: MessageHeader{
=======
		MsgHeader: MsgHeader{
>>>>>>> github/v2
			ToUser:  toUser,
			MsgType: MsgTypeVideo,
		},
	}
	video.Video.MediaId = mediaId
	video.Video.ThumbMediaId = thumbMediaId
	video.Video.Title = title
	video.Video.Description = description

	if kfAccount != "" {
		video.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}

// 音乐消息
type Music struct {
<<<<<<< HEAD
	MessageHeader

=======
	MsgHeader
>>>>>>> github/v2
	Music struct {
		Title        string `json:"title,omitempty"`       // 音乐标题, 可以为 ""
		Description  string `json:"description,omitempty"` // 音乐描述, 可以为 ""
		MusicURL     string `json:"musicurl"`              // 音乐链接
		HQMusicURL   string `json:"hqmusicurl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `json:"thumb_media_id"`        // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
	} `json:"music"`
<<<<<<< HEAD

	*CustomService `json:"customservice,omitempty"`
=======
	CustomService *CustomService `json:"customservice,omitempty"`
>>>>>>> github/v2
}

// 新建音乐消息.
//  如果不指定客服则 kfAccount 留空.
<<<<<<< HEAD
func NewMusic(toUser, thumbMediaId, musicURL, HQMusicURL, title, description,
	kfAccount string) (music *Music) {

	music = &Music{
		MessageHeader: MessageHeader{
=======
func NewMusic(toUser, thumbMediaId, musicURL, HQMusicURL, title, description, kfAccount string) (music *Music) {
	music = &Music{
		MsgHeader: MsgHeader{
>>>>>>> github/v2
			ToUser:  toUser,
			MsgType: MsgTypeMusic,
		},
	}
	music.Music.ThumbMediaId = thumbMediaId
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = HQMusicURL
	music.Music.Title = title
	music.Music.Description = description

	if kfAccount != "" {
		music.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}

// 图文消息里的 Article
type Article struct {
	Title       string `json:"title,omitempty"`       // 图文消息标题
	Description string `json:"description,omitempty"` // 图文消息描述
	URL         string `json:"url,omitempty"`         // 点击图文消息跳转链接
	PicURL      string `json:"picurl,omitempty"`      // 图文消息的图片链接, 支持JPG, PNG格式, 较好的效果为大图640*320, 小图80*80
}

<<<<<<< HEAD
const (
	NewsArticleCountLimit = 10
)

// 图文消息
type News struct {
	MessageHeader

	News struct {
		Articles []Article `json:"articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
	} `json:"news"`

	*CustomService `json:"customservice,omitempty"`
}

// 检查 News 是否有效, 有效返回 nil, 否则返回错误信息.
func (this *News) CheckValid() (err error) {
	n := len(this.News.Articles)
	if n <= 0 {
		err = errors.New("没有有效的图文消息")
		return
	}
	if n > NewsArticleCountLimit {
		err = fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, n)
		return
	}
	return
=======
// 图文消息
type News struct {
	MsgHeader
	News struct {
		Articles []Article `json:"articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过8, 则将会无响应
	} `json:"news"`
	CustomService *CustomService `json:"customservice,omitempty"`
>>>>>>> github/v2
}

// 新建图文消息.
//  如果不指定客服则 kfAccount 留空.
func NewNews(toUser string, articles []Article, kfAccount string) (news *News) {
	news = &News{
<<<<<<< HEAD
		MessageHeader: MessageHeader{
=======
		MsgHeader: MsgHeader{
>>>>>>> github/v2
			ToUser:  toUser,
			MsgType: MsgTypeNews,
		},
	}
	news.News.Articles = articles

	if kfAccount != "" {
		news.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}

<<<<<<< HEAD
// 卡券消息, 特别注意客服消息接口投放卡券仅支持非自定义Code码的卡券
type WxCard struct {
	MessageHeader

=======
type MPNews struct {
	MsgHeader
	MPNews struct {
		MediaId string `json:"media_id"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `json:"mpnews"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

// 新建图文消息.
//  如果不指定客服则 kfAccount 留空.
func NewMPNews(toUser, mediaId, kfAccount string) (mpnews *MPNews) {
	mpnews = &MPNews{
		MsgHeader: MsgHeader{
			ToUser:  toUser,
			MsgType: MsgTypeMPNews,
		},
	}
	mpnews.MPNews.MediaId = mediaId

	if kfAccount != "" {
		mpnews.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}

// 卡券消息, 特别注意客服消息接口投放卡券仅支持非自定义Code码的卡券
type WxCard struct {
	MsgHeader
>>>>>>> github/v2
	WxCard struct {
		CardId  string `json:"card_id"`
		CardExt string `json:"card_ext,omitempty"`
	} `json:"wxcard"`
<<<<<<< HEAD

	*CustomService `json:"customservice,omitempty"`
=======
	CustomService *CustomService `json:"customservice,omitempty"`
>>>>>>> github/v2
}

// 新建卡券消息.
//  如果不指定客服则 kfAccount 留空.
func NewWxCard(toUser, cardId, cardExt, kfAccount string) (card *WxCard) {
	card = &WxCard{
<<<<<<< HEAD
		MessageHeader: MessageHeader{
=======
		MsgHeader: MsgHeader{
>>>>>>> github/v2
			ToUser:  toUser,
			MsgType: MsgTypeWxCard,
		},
	}
	card.WxCard.CardId = cardId
	card.WxCard.CardExt = cardExt

	if kfAccount != "" {
		card.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}
