// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package custom

import (
	"errors"
	"fmt"
)

type CommonHead struct {
	ToUser  string `json:"touser"` // 接收方 OpenID
	MsgType string `json:"msgtype"`
}

// 文本消息
type Text struct {
	CommonHead

	Text struct {
		Content string `json:"content"` // 支持换行符
	} `json:"text"`
}

// 新建文本消息
//  content 支持换行符
func NewText(toUser, content string) (text *Text) {
	text = &Text{
		CommonHead: CommonHead{
			ToUser:  toUser,
			MsgType: MSG_TYPE_TEXT,
		},
	}
	text.Text.Content = content
	return
}

// 图片消息
type Image struct {
	CommonHead

	Image struct {
		MediaId string `json:"media_id"` // 通过上传多媒体文件得到的 MediaId
	} `json:"image"`
}

// 新建图片消息
//  mediaId 是通过上传多媒体文件得到
func NewImage(toUser, mediaId string) (image *Image) {
	image = &Image{
		CommonHead: CommonHead{
			ToUser:  toUser,
			MsgType: MSG_TYPE_IMAGE,
		},
	}
	image.Image.MediaId = mediaId
	return
}

// 语音消息
type Voice struct {
	CommonHead

	Voice struct {
		MediaId string `json:"media_id"` // 通过上传多媒体文件得到的 MediaId
	} `json:"voice"`
}

// 新建语音消息
//  mediaId 是通过上传多媒体文件得到
func NewVoice(toUser, mediaId string) (voice *Voice) {
	voice = &Voice{
		CommonHead: CommonHead{
			ToUser:  toUser,
			MsgType: MSG_TYPE_VOICE,
		},
	}
	voice.Voice.MediaId = mediaId
	return
}

// 视频消息
type Video struct {
	CommonHead

	Video struct {
		MediaId      string `json:"media_id"`              // 通过上传多媒体文件得到的 MediaId
		ThumbMediaId string `json:"thumb_media_id"`        // 缩略图的媒体id, 通过上传多媒体文件得到
		Title        string `json:"title,omitempty"`       // 视频消息的标题
		Description  string `json:"description,omitempty"` // 视频消息的描述
	} `json:"video"`
}

// 新建视频消息
//  mediaId, thumbMediaId 是通过上传多媒体文件得到
//  title, description 可以为 ""
func NewVideo(toUser, mediaId, thumbMediaId, title, description string) (video *Video) {
	video = &Video{
		CommonHead: CommonHead{
			ToUser:  toUser,
			MsgType: MSG_TYPE_VIDEO,
		},
	}
	video.Video.MediaId = mediaId
	video.Video.ThumbMediaId = thumbMediaId
	video.Video.Title = title
	video.Video.Description = description
	return
}

// 音乐消息
type Music struct {
	CommonHead

	Music struct {
		Title        string `json:"title,omitempty"`       // 音乐标题
		Description  string `json:"description,omitempty"` // 音乐描述
		MusicURL     string `json:"musicurl"`              // 音乐链接
		HQMusicURL   string `json:"hqmusicurl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `json:"thumb_media_id"`        // 缩略图的媒体id, 通过上传多媒体文件得到
	} `json:"music"`
}

// 新建音乐消息
//  thumbMediaId 通过上传多媒体文件得到
//  title, description 可以为 ""
func NewMusic(toUser, thumbMediaId, musicURL, HQMusicURL, title, description string) (music *Music) {
	music = &Music{
		CommonHead: CommonHead{
			ToUser:  toUser,
			MsgType: MSG_TYPE_MUSIC,
		},
	}
	music.Music.ThumbMediaId = thumbMediaId
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = HQMusicURL
	music.Music.Title = title
	music.Music.Description = description
	return
}

// 图文消息里的 Article
type NewsArticle struct {
	Title       string `json:"title,omitempty"`       // 图文消息标题
	Description string `json:"description,omitempty"` // 图文消息描述
	URL         string `json:"url,omitempty"`         // 点击图文消息跳转链接
	PicURL      string `json:"picurl,omitempty"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
}

func (this *NewsArticle) Init(title, description, url, picURL string) {
	this.Title = title
	this.Description = description
	this.URL = url
	this.PicURL = picURL
}

// 图文消息
type News struct {
	CommonHead

	News struct {
		Articles []NewsArticle `json:"articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
	} `json:"news"`
}

// 新建图文消息
//  NOTE: articles 的长度不能超过 NewsArticleCountLimit
func NewNews(toUser string, articles []NewsArticle) (news *News) {
	news = &News{
		CommonHead: CommonHead{
			ToUser:  toUser,
			MsgType: MSG_TYPE_NEWS,
		},
	}
	news.News.Articles = articles
	return
}

// 检查 News 是否有效，有效返回 nil，否则返回错误信息
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
