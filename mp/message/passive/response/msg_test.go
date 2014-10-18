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

	// 测试文本消息===============================================================

	want := util.TrimSpace([]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>text</MsgType>
		<Content>你好</Content>
	</xml>`))

	text := NewText("toUser", "fromUser", "你好", 12345678)

	have, err := xml.Marshal(text)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", text, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", text, have, want)
	}

	// 测试图片消息===============================================================

	want = util.TrimSpace([]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>image</MsgType>
		<Image>
			<MediaId>media_id</MediaId>
		</Image>
	</xml>`))

	image := NewImage("toUser", "fromUser", "media_id", 12345678)

	have, err = xml.Marshal(image)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", image, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", image, have, want)
	}

	// 测试语音消息===============================================================

	want = util.TrimSpace([]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>voice</MsgType>
		<Voice>
			<MediaId>media_id</MediaId>
		</Voice>
	</xml>`))

	voice := NewVoice("toUser", "fromUser", "media_id", 12345678)

	have, err = xml.Marshal(voice)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", voice, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", voice, have, want)
	}

	// 测试视频消息===============================================================

	want = util.TrimSpace([]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>video</MsgType>
		<Video>
			<MediaId>media_id</MediaId>
			<Title>title</Title>
			<Description>description</Description>
		</Video> 
	</xml>`))

	video := NewVideo("toUser", "fromUser", "media_id", "title", "description", 12345678)

	have, err = xml.Marshal(video)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", video, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", video, have, want)
	}

	// 测试音乐消息===============================================================

	want = util.TrimSpace([]byte(`<xml>
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
	</xml>`))

	music := NewMusic("toUser", "fromUser", "media_id",
		"MUSIC_Url", "HQ_MUSIC_Url", "TITLE", "DESCRIPTION", 12345678)

	have, err = xml.Marshal(music)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", music, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", music, have, want)
	}

	// 测试图文消息, 没有文章=======================================================

	want = util.TrimSpace([]byte(`<xml>
		<ToUserName>toUser</ToUserName>
		<FromUserName>fromUser</FromUserName>
		<CreateTime>12345678</CreateTime>
		<MsgType>news</MsgType>
		<ArticleCount>0</ArticleCount>
		<Articles>
		</Articles>
	</xml>`))

	news := NewNews("toUser", "fromUser", make([]NewsArticle, 0, 2), 12345678)

	have, err = xml.Marshal(news)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", news, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", news, have, want)
	}

	// 测试图文消息, 1篇文章=======================================================

	want = util.TrimSpace([]byte(`<xml>
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
	</xml>`))

	news.AppendArticle(NewsArticle{
		Title:       "title1",
		Description: "description1",
		PicURL:      "picurl",
		URL:         "url",
	})

	have, err = xml.Marshal(news)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", news, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", news, have, want)
	}

	// 测试图文消息, 2篇文章=======================================================

	want = util.TrimSpace([]byte(`<xml>
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
	</xml>`))

	news.AppendArticle(NewsArticle{
		Title:       "title",
		Description: "description",
		PicURL:      "picurl",
		URL:         "url",
	})

	have, err = xml.Marshal(news)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", news, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", news, have, want)
	}

	// 测试将消息转发到多客服=======================================================

	want = util.TrimSpace([]byte(`<xml>
		<ToUserName>touser</ToUserName>
		<FromUserName>fromuser</FromUserName>
		<CreateTime>1399197672</CreateTime>
		<MsgType>transfer_customer_service</MsgType>
	</xml>`))

	transToCS := NewTransferToCustomerService("touser", "fromuser", 1399197672)

	have, err = xml.Marshal(transToCS)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", transToCS, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", transToCS, have, want)
	}

	// 测试将消息转发到指定客服=======================================================

	want = util.TrimSpace([]byte(`<xml> 
		<ToUserName>touser</ToUserName>  
		<FromUserName>fromuser</FromUserName>  
		<CreateTime>1399197672</CreateTime>  
		<MsgType>transfer_customer_service</MsgType>  
		<TransInfo> 
			<KfAccount>test1@test</KfAccount> 
		</TransInfo> 
	</xml>`))

	transToSCS := NewTransferToSpecialCustomerService("touser", "fromuser", "test1@test", 1399197672)

	have, err = xml.Marshal(transToSCS)
	if err != nil {
		t.Errorf("xml.Marshal(%#q):\nError: %s\n", transToSCS, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("xml.Marshal(%#q):\nhave %#s\nwant %#s\n", transToSCS, have, want)
	}
}
