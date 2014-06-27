// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package response

import (
	"bytes"
	"encoding/xml"
	"github.com/chanxuehong/util"
	"testing"
)

func TestMarshalAndNewFunc(t *testing.T) {
	var expectBytes []byte

	// 测试文本消息===============================================================

	expectBytes = []byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>text</MsgType>
		<Content>你好</Content>
	</xml>`)

	text := NewText("toUser", "fromUser", "你好")
	text.CreateTime = 12345678

	b, err := xml.Marshal(text)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", text, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", text, b, want)
		}
	}

	// 测试图片消息===============================================================

	expectBytes = []byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>image</MsgType>
		<Image>
			<MediaId>media_id</MediaId>
		</Image>
	</xml>`)

	image := NewImage("toUser", "fromUser", "media_id")
	image.CreateTime = 12345678

	b, err = xml.Marshal(image)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", image, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", image, b, want)
		}
	}

	// 测试语音消息===============================================================

	expectBytes = []byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>voice</MsgType>
		<Voice>
			<MediaId>media_id</MediaId>
		</Voice>
	</xml>`)

	voice := NewVoice("toUser", "fromUser", "media_id")
	voice.CreateTime = 12345678

	b, err = xml.Marshal(voice)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", voice, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", voice, b, want)
		}
	}

	// 测试视频消息===============================================================

	expectBytes = []byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>video</MsgType>
		<Video>
			<MediaId>media_id</MediaId>
			<Title>title</Title>
			<Description>description</Description>
		</Video> 
	</xml>`)

	video := NewVideo("toUser", "fromUser", "media_id", "title", "description")
	video.CreateTime = 12345678

	b, err = xml.Marshal(video)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", video, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", video, b, want)
		}
	}

	// 测试音乐消息===============================================================

	expectBytes = []byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>music</MsgType>
		<Music>
			<Title>TITLE</Title>
			<Description>DESCRIPTION</Description>
			<MusicUrl>MUSIC_Url</MusicUrl>
			<HQMusicUrl>HQ_MUSIC_Url</HQMusicUrl>
			<ThumbMediaId>media_id</ThumbMediaId>
		</Music>
	</xml>`)

	music := NewMusic("toUser", "fromUser", "media_id", "MUSIC_Url", "HQ_MUSIC_Url", "TITLE", "DESCRIPTION")
	music.CreateTime = 12345678

	b, err = xml.Marshal(music)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", music, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", music, b, want)
		}
	}

	// 测试图文消息===============================================================

	// 没有文章
	expectBytes = []byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>news</MsgType>
		<ArticleCount>0</ArticleCount>
		<Articles>
		</Articles>
	</xml>`)

	news := NewNews("toUser", "fromUser")
	news.CreateTime = 12345678

	b, err = xml.Marshal(news)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", news, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", news, b, want)
		}
	}

	// 增加一篇没有文章
	expectBytes = []byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>news</MsgType>
		<ArticleCount>1</ArticleCount>
		<Articles>
			<item>
				<Title>title1</Title> 
				<Description>description1</Description>
				<PicUrl>picurl</PicUrl>
				<Url>url</Url>
			</item>
		</Articles>
	</xml>`)

	news.AppendArticle(&NewsArticle{
		Title:       "title1",
		Description: "description1",
		PicURL:      "picurl",
		URL:         "url",
	})

	b, err = xml.Marshal(news)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", news, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", news, b, want)
		}
	}

	// 再增加一篇没有文章
	expectBytes = []byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>news</MsgType>
		<ArticleCount>2</ArticleCount>
		<Articles>
			<item>
				<Title>title1</Title> 
				<Description>description1</Description>
				<PicUrl>picurl</PicUrl>
				<Url>url</Url>
			</item>
			<item>
				<Title>title</Title>
				<Description>description</Description>
				<PicUrl>picurl</PicUrl>
				<Url>url</Url>
			</item>
		</Articles>
	</xml>`)

	news.AppendArticle(&NewsArticle{
		Title:       "title",
		Description: "description",
		PicURL:      "picurl",
		URL:         "url",
	})

	b, err = xml.Marshal(news)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", news, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", news, b, want)
		}
	}

	// 测试将消息转发到多客服=======================================================
	expectBytes = []byte(`<xml>
		<ToUserName>touser</ToUserName>
		<FromUserName>fromuser</FromUserName>
		<CreateTime>1399197672</CreateTime>
		<MsgType>transfer_customer_service</MsgType>
	</xml>`)

	transToCS := NewTransferCustomerService("touser", "fromuser")
	transToCS.CreateTime = 1399197672

	b, err = xml.Marshal(transToCS)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", transToCS, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", transToCS, b, want)
		}
	}
}
