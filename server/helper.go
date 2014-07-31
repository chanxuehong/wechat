// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/chanxuehong/wechat/message/response"
	"io"
	"time"
)

// 向 w io.Writer 写入文本回复消息
func WriteText(w io.Writer, toUser, fromUser, content string) error {
	if w == nil {
		return errors.New("w == nil")
	}

	text := response.Text{
		CommonHead: response.CommonHead{
			ToUserName:   toUser,
			FromUserName: fromUser,
			CreateTime:   time.Now().Unix(),
			MsgType:      response.MSG_TYPE_TEXT,
		},
	}
	text.Content = content

	return writeResponse(w, &text)
}

// 向 w io.Writer 写入图片回复消息
func WriteImage(w io.Writer, toUser, fromUser, mediaId string) error {
	if w == nil {
		return errors.New("w == nil")
	}

	image := response.Image{
		CommonHead: response.CommonHead{
			ToUserName:   toUser,
			FromUserName: fromUser,
			CreateTime:   time.Now().Unix(),
			MsgType:      response.MSG_TYPE_IMAGE,
		},
	}
	image.Image.MediaId = mediaId

	return writeResponse(w, &image)
}

// 向 w io.Writer 写入语音回复消息
func WriteVoice(w io.Writer, toUser, fromUser, mediaId string) error {
	if w == nil {
		return errors.New("w == nil")
	}

	voice := response.Voice{
		CommonHead: response.CommonHead{
			ToUserName:   toUser,
			FromUserName: fromUser,
			CreateTime:   time.Now().Unix(),
			MsgType:      response.MSG_TYPE_VOICE,
		},
	}
	voice.Voice.MediaId = mediaId

	return writeResponse(w, &voice)
}

// 向 w io.Writer 写入视频回复消息
func WriteVideo(w io.Writer, toUser, fromUser, mediaId, title, description string) error {
	if w == nil {
		return errors.New("w == nil")
	}

	video := response.Video{
		CommonHead: response.CommonHead{
			ToUserName:   toUser,
			FromUserName: fromUser,
			CreateTime:   time.Now().Unix(),
			MsgType:      response.MSG_TYPE_VIDEO,
		},
	}
	video.Video.Title = title
	video.Video.Description = description
	video.Video.MediaId = mediaId

	return writeResponse(w, &video)
}

// 向 w io.Writer 写入音乐回复消息
func WriteMusic(w io.Writer, toUser, fromUser, thumbMediaId, musicURL,
	HQMusicURL, title, description string) error {

	if w == nil {
		return errors.New("w == nil")
	}

	music := response.Music{
		CommonHead: response.CommonHead{
			ToUserName:   toUser,
			FromUserName: fromUser,
			CreateTime:   time.Now().Unix(),
			MsgType:      response.MSG_TYPE_MUSIC,
		},
	}
	music.Music.Title = title
	music.Music.Description = description
	music.Music.ThumbMediaId = thumbMediaId
	music.Music.MusicURL = musicURL
	music.Music.HQMusicURL = HQMusicURL

	return writeResponse(w, &music)
}

// 向 w io.Writer 写入图文回复消息
func WriteNews(w io.Writer, toUser, fromUser string, articles []response.NewsArticle) error {
	if w == nil {
		return errors.New("w == nil")
	}
	if len(articles) > response.NewsArticleCountLimit {
		return fmt.Errorf("图文消息的文章个数不能超过 %d, 现在为 %d", response.NewsArticleCountLimit, len(articles))
	}

	news := response.News{
		CommonHead: response.CommonHead{
			ToUserName:   toUser,
			FromUserName: fromUser,
			CreateTime:   time.Now().Unix(),
			MsgType:      response.MSG_TYPE_NEWS,
		},
	}
	news.Articles = articles
	news.ArticleCount = len(articles)

	return writeResponse(w, &news)
}

// 向 w io.Writer 写入转到多客服回复消息
func WriteTransferCustomerService(w io.Writer, toUser, fromUser string) error {
	if w == nil {
		return errors.New("w == nil")
	}

	tcs := response.TransferCustomerService{
		CommonHead: response.CommonHead{
			ToUserName:   toUser,
			FromUserName: fromUser,
			CreateTime:   time.Now().Unix(),
			MsgType:      response.MSG_TYPE_TRANSFER_CUSTOMER_SERVICE,
		},
	}

	return writeResponse(w, &tcs)
}

func writeResponse(w io.Writer, msg interface{}) error {
	return xml.NewEncoder(w).Encode(msg) // 只要 w 能正常的写, 不会返回错误
}
