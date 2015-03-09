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
	MsgTypeText  = "text"  // 文本消息
	MsgTypeImage = "image" // 图片消息
	MsgTypeVoice = "voice" // 语音消息
	MsgTypeVideo = "video" // 视频消息
	MsgTypeMusic = "music" // 音乐消息
	MsgTypeNews  = "news"  // 图文消息
)

type CommonMessageHeader struct {
	ToUser  string `json:"touser"` // 接收方 OpenID
	MsgType string `json:"msgtype"`
}

// 如果需要以某个客服帐号来发消息（在微信6.0.2及以上版本中显示自定义头像），
// 则需在JSON数据包的后半部分加入 customservice 参数
type CustomService struct {
	KfAccount string `json:"kfaccount"`
}

// 文本消息
type Text struct {
	CommonMessageHeader

	Text struct {
		Content string `json:"content"` // 支持换行符
	} `json:"text"`

	*CustomService `json:"customservice,omitempty"`
}

// 新建文本消息.
//  content 支持换行符;
//  如果不指定客服则 kfAccount 留空.
func NewText(toUser, content, kfAccount string) (text *Text) {
	text = &Text{
		CommonMessageHeader: CommonMessageHeader{
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
	CommonMessageHeader

	Image struct {
		MediaId string `json:"media_id"` // 通过上传多媒体文件得到的 MediaId
	} `json:"image"`

	*CustomService `json:"customservice,omitempty"`
}

// 新建图片消息.
//  mediaId 是通过上传多媒体文件得到;
//  如果不指定客服则 kfAccount 留空.
func NewImage(toUser, mediaId, kfAccount string) (image *Image) {
	image = &Image{
		CommonMessageHeader: CommonMessageHeader{
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
	CommonMessageHeader

	Voice struct {
		MediaId string `json:"media_id"` // 通过上传多媒体文件得到的 MediaId
	} `json:"voice"`

	*CustomService `json:"customservice,omitempty"`
}

// 新建语音消息.
//  mediaId 是通过上传多媒体文件得到;
//  如果不指定客服则 kfAccount 留空.
func NewVoice(toUser, mediaId, kfAccount string) (voice *Voice) {
	voice = &Voice{
		CommonMessageHeader: CommonMessageHeader{
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
	CommonMessageHeader

	Video struct {
		MediaId      string `json:"media_id"`              // 通过上传多媒体文件得到的 MediaId
		ThumbMediaId string `json:"thumb_media_id"`        // 缩略图的媒体id, 通过上传多媒体文件得到
		Title        string `json:"title,omitempty"`       // 视频消息的标题
		Description  string `json:"description,omitempty"` // 视频消息的描述
	} `json:"video"`

	*CustomService `json:"customservice,omitempty"`
}

// 新建视频消息.
//  mediaId, thumbMediaId 是通过上传多媒体文件得到;
//  title, description 可以为 "";
//  如果不指定客服则 kfAccount 留空.
func NewVideo(toUser, mediaId, thumbMediaId, title, description,
	kfAccount string) (video *Video) {

	video = &Video{
		CommonMessageHeader: CommonMessageHeader{
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
	CommonMessageHeader

	Music struct {
		Title        string `json:"title,omitempty"`          // 音乐标题
		Description  string `json:"description,omitempty"`    // 音乐描述
		MusicURL     string `json:"musicurl,omitempty"`       // 音乐链接
		HQMusicURL   string `json:"hqmusicurl,omitempty"`     // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `json:"thumb_media_id,omitempty"` // 缩略图的媒体id, 通过上传多媒体文件得到
	} `json:"music"`

	*CustomService `json:"customservice,omitempty"`
}

// 新建音乐消息.
//  thumbMediaId 通过上传多媒体文件得到;
//  title, description 可以为 "";
//  如果不指定客服则 kfAccount 留空.
func NewMusic(toUser, thumbMediaId, musicURL, HQMusicURL, title, description,
	kfAccount string) (music *Music) {

	music = &Music{
		CommonMessageHeader: CommonMessageHeader{
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
	PicURL      string `json:"picurl,omitempty"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
}

const (
	NewsArticleCountLimit = 10
)

// 图文消息
type News struct {
	CommonMessageHeader

	News struct {
		Articles []Article `json:"articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
	} `json:"news"`

	*CustomService `json:"customservice,omitempty"`
}

// 新建图文消息.
//  NOTE: articles 的长度不能超过 NewsArticleCountLimit;
//  如果不指定客服则 kfAccount 留空.
func NewNews(toUser string, articles []Article, kfAccount string) (news *News) {
	news = &News{
		CommonMessageHeader: CommonMessageHeader{
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

// 检查 News 是否有效，有效返回 nil，否则返回错误信息.
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
}
