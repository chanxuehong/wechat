// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package massbygroup

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestJSONMarshal(t *testing.T) {
	var expectBytes []byte

	// 测试文本消息===============================================================

	expectBytes = []byte(`{
		"filter":{
			"group_id":"2"
		},
		"msgtype":"text",
		"text":{
			"content":"CONTENT"
		}
	}`)

	var text Text
	text.Filter.GroupId = 2
	text.MsgType = MSG_TYPE_TEXT
	text.Text.Content = "CONTENT"

	b, err := json.Marshal(&text)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", &text, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", &text, b, want)
		}
	}

	// 测试图片消息===============================================================

	expectBytes = []byte(`{
		"filter":{
			"group_id":"2"
		},
		"msgtype":"image",
		"image":{
			"media_id":"123dsdajkasd231jhksad"
		}
	}`)

	var image Image
	image.Filter.GroupId = 2
	image.MsgType = MSG_TYPE_IMAGE
	image.Image.MediaId = "123dsdajkasd231jhksad"

	b, err = json.Marshal(&image)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", &image, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", &image, b, want)
		}
	}

	// 测试语音消息===============================================================

	expectBytes = []byte(`{
		"filter":{
			"group_id":"2"
		},
		"msgtype":"voice",
		"voice":{
			"media_id":"123dsdajkasd231jhksad"
		}
	}`)

	var voice Voice
	voice.Filter.GroupId = 2
	voice.MsgType = MSG_TYPE_VOICE
	voice.Voice.MediaId = "123dsdajkasd231jhksad"

	b, err = json.Marshal(&voice)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", &voice, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", &voice, b, want)
		}
	}

	// 测试视频消息===============================================================

	expectBytes = []byte(`{
		"filter":{
			"group_id":"2"
		},
		"msgtype":"mpvideo",
		"mpvideo":{
			"media_id":"IhdaAQXuvJtGzwwc0abfXnzeezfO0NgPK6AQYShD8RQYMTtfzbLdBIQkQziv2XJc"
		}
	}`)

	var video Video
	video.Filter.GroupId = 2
	video.MsgType = MSG_TYPE_VIDEO
	video.Video.MediaId = "IhdaAQXuvJtGzwwc0abfXnzeezfO0NgPK6AQYShD8RQYMTtfzbLdBIQkQziv2XJc"

	b, err = json.Marshal(&video)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", &video, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", &video, b, want)
		}
	}

	// 测试图文消息===============================================================

	expectBytes = []byte(`{
		"filter":{
			"group_id":"2"
		},
		"msgtype":"mpnews",
		"mpnews":{
			"media_id":"123dsdajkasd231jhksad"
		}
	}`)

	var news News
	news.Filter.GroupId = 2
	news.MsgType = MSG_TYPE_NEWS
	news.News.MediaId = "123dsdajkasd231jhksad"

	b, err = json.Marshal(&news)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", &news, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", &news, b, want)
		}
	}
}
