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
	ToUser  string `json:"touser"`  // 接收方帐号(OpenID)
	MsgType string `json:"msgtype"` // text, image, voice, video, music, news
}

// 文本消息
//
//  {
//      "touser": "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
//      "msgtype": "text",
//      "text": {
//          "content": "测试文本"
//      }
//  }
type Text struct {
	CommonHead

	Text struct {
		Content string `json:"content"` // 回复的消息内容(换行：在content中能够换行, 微信客户端就支持换行显示)
	} `json:"text"`
}

func NewText(to, content string) (text *Text) {
	text = &Text{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_TEXT,
		},
	}
	text.Text.Content = content

	return
}

// 图片消息
//
//  {
//      "touser": "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
//      "msgtype": "image",
//      "image": {
//          "media_id": "C-bBnTx9XFlVPTCMYWZ6_PeRBCWVfghkSJj2DXTG4faqgAyfjxqdHrtO0Jtpa7K-"
//      }
//  }
type Image struct {
	CommonHead

	Image struct {
		MediaId string `json:"media_id"` // 通过上传多媒体文件, 得到的id
	} `json:"image"`
}

func NewImage(to, mediaId string) (image *Image) {
	image = &Image{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_IMAGE,
		},
	}
	image.Image.MediaId = mediaId

	return
}

// 语音消息
//
//  {
//      "touser": "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
//      "msgtype": "voice",
//      "voice": {
//          "media_id": "GxIcE7umAGoJU29636XgsilpZmNYsiXngcA_RjIV3JJNkFw9fo2muf-94QsC37MT"
//      }
//  }
type Voice struct {
	CommonHead

	Voice struct {
		MediaId string `json:"media_id"` // 通过上传多媒体文件, 得到的id
	} `json:"voice"`
}

func NewVoice(to, mediaId string) (voice *Voice) {
	voice = &Voice{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_VOICE,
		},
	}
	voice.Voice.MediaId = mediaId

	return
}

// 视频消息
//
//  {
//      "touser": "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
//      "msgtype": "video",
//      "video": {
//          "media_id": "kZ9bccrQaFVq1aa3TbLNdXnocPz-LfrfrI8Vrs-pKts8QOmmF66tsoihEW3qhpeP",
//          "title": "标题",
//          "description": "描述"
//      }
//  }
type Video struct {
	CommonHead

	Video struct {
		MediaId     string `json:"media_id"`              // 通过上传多媒体文件, 得到的id
		Title       string `json:"title,omitempty"`       // 视频消息的标题
		Description string `json:"description,omitempty"` // 视频消息的描述
	} `json:"video"`
}

// 新建一个视频消息
//  title, description 可以为 ""
func NewVideo(to, mediaId, title, description string) (video *Video) {
	video = &Video{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_VIDEO,
		},
	}
	video.Video.Title = title
	video.Video.Description = description
	video.Video.MediaId = mediaId

	return
}

// 音乐消息
//
//  {
//      "touser": "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
//      "msgtype": "music",
//      "music": {
//          "title": "标题",
//          "description": "描述",
//          "musicurl": "http://music.baidu.com/song/2191061",
//          "hqmusicurl": "http://music.baidu.com/song/2191061",
//          "thumb_media_id": "4lasRoqC1ydjrq7VhU74mra7KVwacWDVdF6PlS3caQkYdYhrj3rkt7P59GOoSKzX"
//      }
//  }
type Music struct {
	CommonHead

	Music struct {
		Title        string `json:"title,omitempty"`       // 音乐标题
		Description  string `json:"description,omitempty"` // 音乐描述
		MusicURL     string `json:"musicurl"`              // 音乐链接
		HQMusicURL   string `json:"hqmusicurl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `json:"thumb_media_id"`        // 缩略图的媒体id, 通过上传多媒体文件, 得到的id
	} `json:"music"`
}

// 新建一个音乐消息
//  title, description 可以为 ""
func NewMusic(to, thumbMediaId, musicURL, HQMusicURL, title, description string) (music *Music) {
	music = &Music{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_MUSIC,
		},
	}
	music.Music.Title = title
	music.Music.Description = description
	music.Music.ThumbMediaId = thumbMediaId
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = HQMusicURL

	return
}

// 图文消息里的 Article
type NewsArticle struct {
	Title       string `json:"title,omitempty"`       // 图文消息标题
	Description string `json:"description,omitempty"` // 图文消息描述
	PicURL      string `json:"picurl,omitempty"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
	URL         string `json:"url,omitempty"`         // 点击图文消息跳转链接
}

// 图文消息
//
//  {
//      "touser": "os-IKuHd9pJ6xsn4mS7GyL4HxqI4",
//      "msgtype": "news",
//      "news": {
//          "articles": [
//              {
//                  "title": "标题1",
//                  "description": "描述1",
//                  "picurl": "http://news.baidu.com/resource/img/logo_news_137_46.png",
//                  "url": "http://news.baidu.com/"
//              },
//              {
//                  "title": "标题2",
//                  "description": "描述2",
//                  "picurl": "http://mat1.gtimg.com/news/news2013/LOGO.jpg",
//                  "url": "http://news.qq.com/"
//              }
//          ]
//      }
//  }
type News struct {
	CommonHead

	News struct {
		Articles []NewsArticle `json:"articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
	} `json:"news"`
}

// 新建图文消息
//  NOTE: articles 的长度不能超过 NewsArticleCountLimit
func NewNews(to string, articles []NewsArticle) (news *News) {
	news = &News{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_NEWS,
		},
	}
	news.News.Articles = articles

	return
}

// 检查 News 是否有效，有效返回 nil，否则返回错误信息
func (n *News) CheckValid() (err error) {
	articleNum := len(n.News.Articles)
	if articleNum == 0 {
		err = errors.New("图文消息是空的")
		return
	}
	if articleNum > NewsArticleCountLimit {
		err = fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, articleNum)
		return
	}
	return
}
