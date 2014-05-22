// 回复消息, 包括被动回复和主动回复

package message

import (
	"encoding/xml"
	"fmt"
	"time"
)

type responseCommonHead struct {
	ToUserName   string `xml:"ToUserName"   json:"touser"`  // 接收方帐号(收到的OpenID)
	FromUserName string `xml:"FromUserName" json:"-"`       // 开发者微信号
	CreateTime   int64  `xml:"CreateTime"   json:"-"`       // 消息创建时间(整型), unixtime
	MsgType      string `xml:"MsgType"      json:"msgtype"` // text, image, voice, video, music, news
}

// text ========================================================================

type textResponseBody struct {
	Content string `xml:"Content" json:"content"` // 回复的消息内容(换行：在content中能够换行, 微信客户端就支持换行显示)
}
type TextResponseMsg struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	responseCommonHead
	textResponseBody `json:"text"`
}

func NewTextResponseMsg(to, from, content string) *TextResponseMsg {
	return &TextResponseMsg{
		responseCommonHead: responseCommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      RESP_MSG_TYPE_TEXT,
		},
		textResponseBody: textResponseBody{
			Content: content,
		},
	}
}

// image =======================================================================

type imageResponseBody struct {
	MediaId string `xml:"Image>MediaId" json:"media_id"` // 通过上传多媒体文件, 得到的id
}
type ImageResponseMsg struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	responseCommonHead
	imageResponseBody `json:"image"`
}

func NewImageResponseMsg(to, from, mediaId string) *ImageResponseMsg {
	return &ImageResponseMsg{
		responseCommonHead: responseCommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      RESP_MSG_TYPE_IMAGE,
		},
		imageResponseBody: imageResponseBody{
			MediaId: mediaId,
		},
	}
}

// voice =======================================================================

type voiceResponseBody struct {
	MediaId string `xml:"Voice>MediaId" json:"media_id"` // 通过上传多媒体文件, 得到的id
}
type VoiceResponseMsg struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	responseCommonHead
	voiceResponseBody `json:"voice"`
}

func NewVoiceResponseMsg(to, from, mediaId string) *VoiceResponseMsg {
	return &VoiceResponseMsg{
		responseCommonHead: responseCommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      RESP_MSG_TYPE_VOICE,
		},
		voiceResponseBody: voiceResponseBody{
			MediaId: mediaId,
		},
	}
}

// video =======================================================================

type videoResponseBody struct {
	MediaId     string `xml:"Video>MediaId"               json:"media_id"`              // 通过上传多媒体文件, 得到的id
	Title       string `xml:"Video>Title,omitempty"       json:"title,omitempty"`       // 视频消息的标题
	Description string `xml:"Video>Description,omitempty" json:"description,omitempty"` // 视频消息的描述
}
type VideoResponseMsg struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	responseCommonHead
	videoResponseBody `json:"video"`
}

// title, description 可以为 ""
func NewVideoResponseMsg(to, from, mediaId, title, description string) *VideoResponseMsg {
	return &VideoResponseMsg{
		responseCommonHead: responseCommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      RESP_MSG_TYPE_VIDEO,
		},
		videoResponseBody: videoResponseBody{
			MediaId:     mediaId,
			Title:       title,
			Description: description,
		},
	}
}

// music =======================================================================

type musicResponseBody struct {
	Title        string `xml:"Music>Title,omitempty"       json:"title,omitempty"`       // 音乐标题
	Description  string `xml:"Music>Description,omitempty" json:"description,omitempty"` // 音乐描述
	MusicUrl     string `xml:"Music>MusicUrl"              json:"musicurl"`              // 音乐链接
	HQMusicUrl   string `xml:"Music>HQMusicUrl"            json:"hqmusicurl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
	ThumbMediaId string `xml:"Music>ThumbMediaId"          json:"thumb_media_id"`        // 缩略图的媒体id, 通过上传多媒体文件, 得到的id
}
type MusicResponseMsg struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	responseCommonHead
	musicResponseBody `json:"music"`
}

// title, description 可以为 ""
func NewMusicResponseMsg(to, from, thumbMediaId, musicUrl, HQMusicUrl, title, description string) *MusicResponseMsg {
	return &MusicResponseMsg{
		responseCommonHead: responseCommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      RESP_MSG_TYPE_MUSIC,
		},
		musicResponseBody: musicResponseBody{
			Title:        title,
			Description:  description,
			MusicUrl:     musicUrl,
			HQMusicUrl:   HQMusicUrl,
			ThumbMediaId: thumbMediaId,
		},
	}
}

// news ========================================================================

// 图文消息里的 item
type Article struct {
	Title       string `xml:"Title,omitempty"       json:"title,omitempty"`       // 图文消息标题
	Description string `xml:"Description,omitempty" json:"description,omitempty"` // 图文消息描述
	PicUrl      string `xml:"PicUrl,omitempty"      json:"picurl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	Url         string `xml:"Url,omitempty"         json:"url,omitempty"`         // 点击图文消息跳转链接
}
type newsResponseBody struct {
	ArticleCount int        `xml:"ArticleCount"  json:"-"`        // 图文消息个数, 限制为10条以内
	Articles     []*Article `xml:"Articles>item" json:"articles"` // 多条图文消息信息, 默认第一个item为大图,注意, 如果图文数超过10, 则将会无响应
}

// 图文消息
//  NOTE: 尽量用 NewNewsResponseMsg 创建对象
type NewsResponseMsg struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	responseCommonHead
	newsResponseBody `json:"news"`
}

// NOTE: 如果图文消息数量大于微信的限制, 则把多余的截除.
func NewNewsResponseMsg(to, from string, articles []*Article) *NewsResponseMsg {
	if len(articles) > newsResponseMsgArticleCountLimit {
		articles = articles[:newsResponseMsgArticleCountLimit]
	} else if articles == nil {
		articles = make([]*Article, 0, newsResponseMsgArticleCountLimit)
	}

	return &NewsResponseMsg{
		responseCommonHead: responseCommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      RESP_MSG_TYPE_NEWS,
		},
		newsResponseBody: newsResponseBody{
			ArticleCount: len(articles),
			Articles:     articles,
		},
	}
}

// 如果当前的图文数量已经达到了上限, 则返回错误, msg *NewsResponseMsg 不做修改;
// 否则返回 nil.
func (msg *NewsResponseMsg) AppendArticle(article *Article) error {
	if article == nil {
		return nil
	}
	if msg.ArticleCount >= newsResponseMsgArticleCountLimit {
		return fmt.Errorf("当前的图文消息已经达到数量上限: %d\n", newsResponseMsgArticleCountLimit)
	}

	msg.Articles = append(msg.Articles, article)
	msg.ArticleCount = len(msg.Articles)
	return nil
}
