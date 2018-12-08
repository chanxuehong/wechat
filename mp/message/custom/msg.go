package custom

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	MsgTypeText       core.MsgType = "text"   // 文本消息
	MsgTypeImage      core.MsgType = "image"  // 图片消息
	MsgTypeVoice      core.MsgType = "voice"  // 语音消息
	MsgTypeVideo      core.MsgType = "video"  // 视频消息
	MsgTypeMusic      core.MsgType = "music"  // 音乐消息
	MsgTypeNews       core.MsgType = "news"   // 图文消息
	MsgTypeMPNews     core.MsgType = "mpnews" // 图文消息, 发送已经创建好的图文
	MsgTypeWxCard     core.MsgType = "wxcard" // 卡卷消息
	MsgTypeWxMiniLink core.MsgType = "link"   // 小程序客服消息:图文链接
)

type MsgHeader struct {
	ToUser  string       `json:"touser"` // 接收方 OpenID
	MsgType core.MsgType `json:"msgtype"`
}

type CustomService struct {
	KfAccount string `json:"kf_account"`
}

// 文本消息
type Text struct {
	MsgHeader
	Text struct {
		Content string `json:"content"` // 支持换行符
	} `json:"text"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

// 新建文本消息.
//  如果不指定客服则 kfAccount 留空.
func NewText(toUser, content, kfAccount string) (text *Text) {
	text = &Text{
		MsgHeader: MsgHeader{
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
	MsgHeader
	Image struct {
		MediaId string `json:"media_id"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `json:"image"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

// 新建图片消息.
//  如果不指定客服则 kfAccount 留空.
func NewImage(toUser, mediaId, kfAccount string) (image *Image) {
	image = &Image{
		MsgHeader: MsgHeader{
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
	MsgHeader
	Voice struct {
		MediaId string `json:"media_id"` // 通过素材管理接口上传多媒体文件得到 MediaId
	} `json:"voice"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

// 新建语音消息.
//  如果不指定客服则 kfAccount 留空.
func NewVoice(toUser, mediaId, kfAccount string) (voice *Voice) {
	voice = &Voice{
		MsgHeader: MsgHeader{
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
	MsgHeader
	Video struct {
		MediaId      string `json:"media_id"`              // 通过素材管理接口上传多媒体文件得到 MediaId
		ThumbMediaId string `json:"thumb_media_id"`        // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
		Title        string `json:"title,omitempty"`       // 视频消息的标题, 可以为 ""
		Description  string `json:"description,omitempty"` // 视频消息的描述, 可以为 ""
	} `json:"video"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

// 新建视频消息.
//  如果不指定客服则 kfAccount 留空.
func NewVideo(toUser, mediaId, thumbMediaId, title, description, kfAccount string) (video *Video) {
	video = &Video{
		MsgHeader: MsgHeader{
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
	MsgHeader
	Music struct {
		Title        string `json:"title,omitempty"`       // 音乐标题, 可以为 ""
		Description  string `json:"description,omitempty"` // 音乐描述, 可以为 ""
		MusicURL     string `json:"musicurl"`              // 音乐链接
		HQMusicURL   string `json:"hqmusicurl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `json:"thumb_media_id"`        // 通过素材管理接口上传多媒体文件得到 ThumbMediaId
	} `json:"music"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

// 新建音乐消息.
//  如果不指定客服则 kfAccount 留空.
func NewMusic(toUser, thumbMediaId, musicURL, HQMusicURL, title, description, kfAccount string) (music *Music) {
	music = &Music{
		MsgHeader: MsgHeader{
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

// 图文消息
type News struct {
	MsgHeader
	News struct {
		Articles []Article `json:"articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过8, 则将会无响应
	} `json:"news"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

// 新建图文消息.
//  如果不指定客服则 kfAccount 留空.
func NewNews(toUser string, articles []Article, kfAccount string) (news *News) {
	news = &News{
		MsgHeader: MsgHeader{
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
	WxCard struct {
		CardId  string `json:"card_id"`
		CardExt string `json:"card_ext,omitempty"`
	} `json:"wxcard"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

// 新建卡券消息.
//  如果不指定客服则 kfAccount 留空.
func NewWxCard(toUser, cardId, cardExt, kfAccount string) (card *WxCard) {
	card = &WxCard{
		MsgHeader: MsgHeader{
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

type WxMiniLink struct {
	MsgHeader
	Link struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		URL         string `json:"url"`
		ThumbURL    string `json:"thumb_url"`
	} `json:"link"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

func NewMiniLink(toUser, title, desc, url, thumbUrl, kfAccount string) (link *WxMiniLink) {
	link = &WxMiniLink{
		MsgHeader: MsgHeader{
			ToUser:  toUser,
			MsgType: MsgTypeWxMiniLink,
		},
	}

	link.Link.Title = title
	link.Link.Description = desc
	link.Link.URL = url
	link.Link.ThumbURL = thumbUrl

	if kfAccount != "" {
		link.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}

type WxMiniPage struct {
	MsgHeader
	MiniProgramPage struct {
		Title        string `json:"title"`
		PagePath     string `json:"pagepath"`
		ThumbMediaId string `json:"thumb_media_id"`
	} `json:"miniprogrampage"`
	CustomService *CustomService `json:"customservice,omitempty"`
}

func NewMiniPage(toUser, title, pagePath, thumbMediaId, kfAccount string) (page *WxMiniPage) {
	page = &WxMiniPage{
		MsgHeader: MsgHeader{
			ToUser:  toUser,
			MsgType: MsgTypeWxMiniLink,
		},
	}

	page.MiniProgramPage.Title = title
	page.MiniProgramPage.PagePath = pagePath
	page.MiniProgramPage.ThumbMediaId = thumbMediaId
	if kfAccount != "" {
		page.CustomService = &CustomService{
			KfAccount: kfAccount,
		}
	}
	return
}
