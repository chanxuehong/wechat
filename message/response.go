package message

import (
	"encoding/xml"
	"time"
)

type commonResponseHead struct {
	ToUserName   string `xml:"ToUserName"   json:"touser"`  // 接收方帐号(收到的OpenID)
	FromUserName string `xml:"FromUserName" json:"-"`       // 开发者微信号
	CreateTime   int64  `xml:"CreateTime"   json:"-"`       // 消息创建时间(整型), unixtime
	MsgType      string `xml:"MsgType"      json:"msgtype"` // text, image, voice, video, music, news
}

// text ========================================================================

type textResponseBody struct {
	Content string `xml:"Content" json:"content"` // 回复的消息内容(换行：在content中能够换行, 微信客户端就支持换行显示)
}
type TextResponse struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	commonResponseHead
	textResponseBody `json:"text"`
}

func NewTextResponse(to, from, content string) *TextResponse {
	return &TextResponse{
		commonResponseHead: commonResponseHead{
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
type ImageResponse struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	commonResponseHead
	imageResponseBody `json:"image"`
}

func NewImageResponse(to, from, mediaId string) *ImageResponse {
	return &ImageResponse{
		commonResponseHead: commonResponseHead{
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
type VoiceResponse struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	commonResponseHead
	voiceResponseBody `json:"voice"`
}

func NewVoiceResponse(to, from, mediaId string) *VoiceResponse {
	return &VoiceResponse{
		commonResponseHead: commonResponseHead{
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
type VideoResponse struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	commonResponseHead
	videoResponseBody `json:"video"`
}

// title, description 可以为 ""
func NewVideoResponse(to, from, mediaId, title, description string) *VideoResponse {
	return &VideoResponse{
		commonResponseHead: commonResponseHead{
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
type MusicResponse struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	commonResponseHead
	musicResponseBody `json:"music"`
}

// title, description 可以为 ""
func NewMusicResponse(to, from, thumbMediaId, musicUrl, HQMusicUrl, title, description string) *MusicResponse {
	return &MusicResponse{
		commonResponseHead: commonResponseHead{
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
type NewsResponseArticle struct {
	Title       string `xml:"Title,omitempty"       json:"title,omitempty"`       // 图文消息标题
	Description string `xml:"Description,omitempty" json:"description,omitempty"` // 图文消息描述
	PicUrl      string `xml:"PicUrl,omitempty"      json:"picurl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	Url         string `xml:"Url,omitempty"         json:"url,omitempty"`         // 点击图文消息跳转链接
}

type newsResponseBody struct {
	ArticleCount int                    `xml:"ArticleCount"  json:"-"`        // 图文消息个数, 限制为10条以内
	Articles     []*NewsResponseArticle `xml:"Articles>item" json:"articles"` // 多条图文消息信息, 默认第一个item为大图,注意, 如果图文数超过10, 则将会无响应
}

// 图文消息
//  NOTE: 尽量用 NewNewsResponse 创建对象
type NewsResponse struct {
	XMLName xml.Name `xml:"xml" json:"-"`
	commonResponseHead
	newsResponseBody `json:"news"`
}

// NOTE: 如果图文消息数量大于微信的限制, 则把多余的截除.
func NewNewsResponse(to, from string, articles []*NewsResponseArticle) *NewsResponse {
	if len(articles) > NewsResponseArticleCountLimit {
		articles = articles[:NewsResponseArticleCountLimit]
	} else if articles == nil {
		articles = make([]*NewsResponseArticle, 0, NewsResponseArticleCountLimit)
	}

	return &NewsResponse{
		commonResponseHead: commonResponseHead{
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

// 如果总的按钮数超过限制, 则截除多余的.
func (msg *NewsResponse) AppendArticle(article ...*NewsResponseArticle) {
	if len(article) <= 0 {
		return
	}

	switch n := NewsResponseArticleCountLimit - len(msg.Articles); {
	case n > 0:
		if len(article) > n {
			article = article[:n]
		}
		msg.Articles = append(msg.Articles, article...)
	case n == 0:
	default: // n < 0
		msg.Articles = msg.Articles[:NewsResponseArticleCountLimit]
	}
}
