// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package response

import (
	"errors"
	"fmt"
	"time"
)

type CommonHead struct {
	ToUserName   string `xml:"ToUserName"   json:"ToUserName"`   // 接收方帐号(OpenID)
	FromUserName string `xml:"FromUserName" json:"FromUserName"` // 开发者微信号
	CreateTime   int64  `xml:"CreateTime"   json:"CreateTime"`   // 消息创建时间(整型), unixtime
	MsgType      string `xml:"MsgType"      json:"MsgType"`      // 消息类型
}

// 文本消息
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Content string `xml:"Content" json:"Content"` // 回复的消息内容, 支持换行符
}

// 新建文本消息
//  NOTE: content 支持换行符
func NewText(to, from, content string) (text *Text) {
	text = &Text{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_TEXT,
		},
	}
	text.Content = content
	return
}

// 图片消息
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Image struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // MediaId 通过上传多媒体文件得到
	} `xml:"Image" json:"Image"`
}

// 新建图片消息
//  MediaId 通过上传多媒体文件得到
func NewImage(to, from, mediaId string) (image *Image) {
	image = &Image{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_IMAGE,
		},
	}
	image.Image.MediaId = mediaId
	return
}

// 语音消息
type Voice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Voice struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // MediaId 通过上传多媒体文件得到
	} `xml:"Voice" json:"Voice"`
}

// 新建语音消息
//  MediaId 通过上传多媒体文件得到
func NewVoice(to, from, mediaId string) (voice *Voice) {
	voice = &Voice{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_VOICE,
		},
	}
	voice.Voice.MediaId = mediaId
	return
}

// 视频消息
type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Video struct {
		MediaId     string `xml:"MediaId"               json:"MediaId"`               // MediaId 通过上传多媒体文件得到
		Title       string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 视频消息的标题
		Description string `xml:"Description,omitempty" json:"Description,omitempty"` // 视频消息的描述
	} `xml:"Video" json:"Video"`
}

// 新建视频消息
//  MediaId 通过上传多媒体文件得到
//  title, description 可以为 ""
func NewVideo(to, from, mediaId, title, description string) (video *Video) {
	video = &Video{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_VIDEO,
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
	CommonHead

	Music struct {
		Title        string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 音乐标题
		Description  string `xml:"Description,omitempty" json:"Description,omitempty"` // 音乐描述
		MusicURL     string `xml:"MusicUrl"              json:"MusicUrl"`              // 音乐链接
		HQMusicURL   string `xml:"HQMusicUrl"            json:"HQMusicUrl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `xml:"ThumbMediaId"          json:"ThumbMediaId"`          // 缩略图的媒体id, 通过上传多媒体文件得到
	} `xml:"Music" json:"Music"`
}

// 新建音乐消息
//  thumbMediaId 通过上传多媒体文件得到
//  title, description 可以为 ""
func NewMusic(to, from, thumbMediaId, musicURL,
	HQMusicURL, title, description string) (music *Music) {

	music = &Music{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_MUSIC,
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
type NewsArticle struct {
	Title       string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 图文消息标题
	Description string `xml:"Description,omitempty" json:"Description,omitempty"` // 图文消息描述
	PicURL      string `xml:"PicUrl,omitempty"      json:"PicUrl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         string `xml:"Url,omitempty"         json:"Url,omitempty"`         // 点击图文消息跳转链接
}

func (this *NewsArticle) Init(title, description, url, picURL string) {
	this.Title = title
	this.Description = description
	this.URL = url
	this.PicURL = picURL
}

// 图文消息.
//  NOTE: Articles 赋值的同时也要更改 ArticleCount 字段, 建议用 NewNews() 和 News.AppendArticle()
type News struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	ArticleCount int           `xml:"ArticleCount"            json:"ArticleCount"`       // 图文消息个数, 限制为10条以内
	Articles     []NewsArticle `xml:"Articles>item,omitempty" json:"Articles,omitempty"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
}

// NOTE: articles 的长度不能超过 NewsArticleCountLimit
func NewNews(to, from string, articles []NewsArticle) (news *News) {
	news = &News{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_NEWS,
		},
	}
	news.Articles = articles
	news.ArticleCount = len(articles)

	return
}

// 更新 this.ArticleCount 字段, 使其等于 len(this.Articles)
func (this *News) UpdateArticleCount() {
	this.ArticleCount = len(this.Articles)
}

// 增加文章到图文消息中, 该方法会自动更新 News.ArticleCount 字段
func (this *News) AppendArticle(article ...NewsArticle) {
	this.Articles = append(this.Articles, article...)
	this.ArticleCount = len(this.Articles)
}

// 检查 News 是否有效，有效返回 nil，否则返回错误信息
func (this *News) CheckValid() (err error) {
	n := len(this.Articles)

	if n != this.ArticleCount {
		err = fmt.Errorf("图文消息的 ArticleCount == %d, 实际文章个数为 %d", this.ArticleCount, n)
		return
	}
	if n <= 0 {
		err = errors.New("图文消息里没有文章")
		return
	}
	if n > NewsArticleCountLimit {
		err = fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", NewsArticleCountLimit, n)
		return
	}
	return
}

// 将消息转发到多客服
type TransferToCustomerService struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead
}

func NewTransferToCustomerService(to, from string) *TransferToCustomerService {
	return &TransferToCustomerService{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_TRANSFER_CUSTOMER_SERVICE,
		},
	}
}

// 将消息转发到指定客服
type TransferToSpecialCustomerService struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	TransInfo struct {
		KfAccount string `xml:"KfAccount"         json:"KfAccount"`
	} `xml:"TransInfo"         json:"TransInfo"`
}

func NewTransferToSpecialCustomerService(to, from, KfAccount string) (msg *TransferToSpecialCustomerService) {
	msg = &TransferToSpecialCustomerService{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_TRANSFER_CUSTOMER_SERVICE,
		},
	}
	msg.TransInfo.KfAccount = KfAccount
	return
}
