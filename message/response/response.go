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
	MsgType      string `xml:"MsgType"      json:"MsgType"`      // text, image, voice, video, music, news, transfer_customer_service
}

// 文本消息
//
//  <xml>
//      <ToUserName>ovx6euNq-hN2do74jeVSqZB82DiE</ToUserName>
//      <FromUserName>gh_xxxxxxxxxxxx</FromUserName>
//      <CreateTime>1406609798</CreateTime>
//      <MsgType>text</MsgType>
//      <Content>文本回复测试</Content>
//  </xml>
type Text struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Content string `xml:"Content" json:"Content"` // 回复的消息内容(换行：在content中能够换行, 微信客户端就支持换行显示)
}

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
//
//  <xml>
//      <ToUserName>os-IKuHd9pJ6xsn4mS7GyL4HxqI4</ToUserName>
//      <FromUserName>gh_xxxxxxxxxxxx</FromUserName>
//      <CreateTime>1406609903</CreateTime>
//      <MsgType>image</MsgType>
//      <Image>
//          <MediaId>C-bBnTx9XFlVPTCMYWZ6_PeRBCWVfghkSJj2DXTG4faqgAyfjxqdHrtO0Jtpa7K-</MediaId>
//      </Image>
//  </xml>
type Image struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Image struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // 通过上传多媒体文件, 得到的id
	} `xml:"Image" json:"Image"`
}

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
//
//  <xml>
//      <ToUserName>os-IKuHd9pJ6xsn4mS7GyL4HxqI4</ToUserName>
//      <FromUserName>gh_xxxxxxxxxxxx</FromUserName>
//      <CreateTime>1406610000</CreateTime>
//      <MsgType>voice</MsgType>
//      <Voice>
//          <MediaId>GxIcE7umAGoJU29636XgsilpZmNYsiXngcA_RjIV3JJNkFw9fo2muf-94QsC37MT</MediaId>
//      </Voice>
//  </xml>
type Voice struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Voice struct {
		MediaId string `xml:"MediaId" json:"MediaId"` // 通过上传多媒体文件, 得到的id
	} `xml:"Voice" json:"Voice"`
}

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
//
//  <xml>
//      <ToUserName>os-IKuHd9pJ6xsn4mS7GyL4HxqI4</ToUserName>
//      <FromUserName>gh_xxxxxxxxxxxx</FromUserName>
//      <CreateTime>1406610204</CreateTime>
//      <MsgType>video</MsgType>
//      <Video>
//          <Title>标题</Title>
//          <Description>描述</Description>
//          <MediaId>kZ9bccrQaFVq1aa3TbLNdXnocPz-LfrfrI8Vrs-pKts8QOmmF66tsoihEW3qhpeP</MediaId>
//      </Video>
//  </xml>
type Video struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Video struct {
		Title       string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 视频消息的标题
		Description string `xml:"Description,omitempty" json:"Description,omitempty"` // 视频消息的描述
		MediaId     string `xml:"MediaId"               json:"MediaId"`               // 通过上传多媒体文件, 得到的id
	} `xml:"Video" json:"Video"`
}

// title, description 可以为 ""
func NewVideo(to, from, mediaId, title, description string) (video *Video) {
	video = &Video{
		CommonHead: CommonHead{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   time.Now().Unix(),
			MsgType:      MSG_TYPE_VIDEO,
		},
	}
	video.Video.Title = title
	video.Video.Description = description
	video.Video.MediaId = mediaId

	return
}

// 音乐消息
//
//  <xml>
//      <ToUserName>os-IKuHd9pJ6xsn4mS7GyL4HxqI4</ToUserName>
//      <FromUserName>gh_xxxxxxxxxxxx</FromUserName>
//      <CreateTime>1406610407</CreateTime>
//      <MsgType>music</MsgType>
//      <Music>
//          <Title>标题</Title>
//          <Description>描述</Description>
//          <MusicUrl>http://music.baidu.com/song/2191061</MusicUrl>
//          <HQMusicUrl>http://music.baidu.com/song/2191061</HQMusicUrl>
//          <ThumbMediaId>4lasRoqC1ydjrq7VhU74mra7KVwacWDVdF6PlS3caQkYdYhrj3rkt7P59GOoSKzX</ThumbMediaId>
//      </Music>
//  </xml>
type Music struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Music struct {
		Title        string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 音乐标题
		Description  string `xml:"Description,omitempty" json:"Description,omitempty"` // 音乐描述
		MusicURL     string `xml:"MusicUrl"              json:"MusicUrl"`              // 音乐链接
		HQMusicURL   string `xml:"HQMusicUrl"            json:"HQMusicUrl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `xml:"ThumbMediaId"          json:"ThumbMediaId"`          // 缩略图的媒体id, 通过上传多媒体文件, 得到的id
	} `xml:"Music" json:"Music"`
}

// title, description 可以为 ""
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
	music.Music.ThumbMediaId = thumbMediaId
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = HQMusicURL

	return
}

// 图文消息里的 Article
type NewsArticle struct {
	Title       string `xml:"Title,omitempty"       json:"Title,omitempty"`       // 图文消息标题
	Description string `xml:"Description,omitempty" json:"Description,omitempty"` // 图文消息描述
	PicURL      string `xml:"PicUrl,omitempty"      json:"PicUrl,omitempty"`      // 图片链接, 支持JPG, PNG格式, 较好的效果为大图360*200, 小图200*200
	URL         string `xml:"Url,omitempty"         json:"Url,omitempty"`         // 点击图文消息跳转链接
}

// 图文消息.
//  NOTE: Articles 赋值的同时也要更改 ArticleCount 字段, 建议用 NewNews() 和 News.AppendArticle()
//
//  <xml>
//      <ToUserName>os-IKuHd9pJ6xsn4mS7GyL4HxqI4</ToUserName>
//      <FromUserName>gh_xxxxxxxxxxxx</FromUserName>
//      <CreateTime>1406611521</CreateTime>
//      <MsgType>news</MsgType>
//      <ArticleCount>2</ArticleCount>
//      <Articles>
//          <item>
//              <Title>标题1</Title>
//              <Description>描述1</Description>
//              <PicUrl>http://news.baidu.com/resource/img/logo_news_137_46.png</PicUrl>
//              <Url>http://news.baidu.com/</Url>
//          </item>
//          <item>
//              <Title>标题2</Title>
//              <Description>描述2</Description>
//              <PicUrl>http://mat1.gtimg.com/news/news2013/LOGO.jpg</PicUrl>
//              <Url>http://news.qq.com/</Url>
//          </item>
//      </Articles>
//  </xml>
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

// n.ArticleCount = len(n.Articles)
func (n *News) UpdateArticleCount() {
	n.ArticleCount = len(n.Articles)
}

// 检查 News 是否有效，有效返回 nil，否则返回错误信息
func (n *News) CheckValid() (err error) {
	articleNum := len(n.Articles)

	if articleNum != n.ArticleCount {
		err = fmt.Errorf("图文消息的 ArticleCount == %d, 实际文章个数为 %d", n.ArticleCount, articleNum)
		return
	}
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
