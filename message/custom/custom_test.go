// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package custom

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestMarshalAndNewFunc(t *testing.T) {
	var expectBytes []byte

	// 测试文本消息===============================================================

	expectBytes = []byte(`{
		"touser":"toUser",
		"msgtype":"text",
		"text":
		{
			"content":"你好"
		}
	}`)

	text := NewText("toUser", "你好")

	b, err := json.Marshal(text)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", text, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", text, b, want)
		}
	}

	// 测试图片消息===============================================================

	expectBytes = []byte(`{
		"touser":"toUser",
		"msgtype":"image",
		"image":
		{
			"media_id":"media_id"
		}
	}`)

	image := NewImage("toUser", "media_id")

	b, err = json.Marshal(image)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", image, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", image, b, want)
		}
	}

	// 测试语音消息===============================================================

	expectBytes = []byte(`{
		"touser":"toUser",
		"msgtype":"voice",
		"voice":
		{
			"media_id":"media_id"
		}
	}`)

	voice := NewVoice("toUser", "media_id")

	b, err = json.Marshal(voice)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", voice, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", voice, b, want)
		}
	}

	// 测试视频消息===============================================================

	expectBytes = []byte(`{
		"touser":"toUser",
		"msgtype":"video",
		"video":
		{
			"media_id":"media_id",
			"title":"title",
			"description":"description"
		}
	}`)

	video := NewVideo("toUser", "media_id", "title", "description")

	b, err = json.Marshal(video)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", video, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", video, b, want)
		}
	}

	// 测试音乐消息===============================================================

	expectBytes = []byte(`{
		"touser":"toUser",
		"msgtype":"music",
		"music":
		{
			"title":"TITLE",
			"description":"DESCRIPTION",
			"musicurl":"MUSIC_Url",
			"hqmusicurl":"HQ_MUSIC_Url",
			"thumb_media_id":"media_id" 
		}
	}`)

	music := NewMusic("toUser", "media_id", "MUSIC_Url", "HQ_MUSIC_Url", "TITLE", "DESCRIPTION")

	b, err = json.Marshal(music)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", music, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", music, b, want)
		}
	}

	// 测试图文消息===============================================================

	// 没有文章
	expectBytes = []byte(`{
		"touser":"toUser",
		"msgtype":"news",
		"news":{}
	}`)

	news := NewNews("toUser", nil)

	b, err = json.Marshal(news)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", news, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", news, b, want)
		}
	}

	// 增加一篇没有文章
	expectBytes = []byte(`{
		"touser":"toUser",
		"msgtype":"news",
		"news":{
			"articles":[
				{
					"title":"title1",
					"description":"description1",
					"picurl":"picurl",
					"url":"url"
				}
			]
		}
	}`)

	news.News.Articles = append(
		news.News.Articles,
		NewsArticle{
			Title:       "title1",
			Description: "description1",
			PicURL:      "picurl",
			URL:         "url",
		},
	)

	b, err = json.Marshal(news)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", news, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", news, b, want)
		}
	}

	// 再增加一篇没有文章
	expectBytes = []byte(`{
		"touser":"toUser",
		"msgtype":"news",
		"news":{
			"articles":[
				{
					"title":"title1",
					"description":"description1",
					"picurl":"picurl",
					"url":"url"
				},
				{
					"title":"title",
					"description":"description",
					"picurl":"picurl",
					"url":"url"
				}
			]
		}
	}`)

	news.News.Articles = append(
		news.News.Articles,
		NewsArticle{
			Title:       "title",
			Description: "description",
			PicURL:      "picurl",
			URL:         "url",
		},
	)

	b, err = json.Marshal(news)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", news, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", news, b, want)
		}
	}
}
