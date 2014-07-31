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

func TestMarshalAndNewFunc(t *testing.T) {
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

	text := NewText(2, "CONTENT")

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
		"filter":{
			"group_id":"2"
		},
		"msgtype":"image",
		"image":{
			"media_id":"123dsdajkasd231jhksad"
		}
	}`)

	image := NewImage(2, "123dsdajkasd231jhksad")

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
		"filter":{
			"group_id":"2"
		},
		"msgtype":"voice",
		"voice":{
			"media_id":"123dsdajkasd231jhksad"
		}
	}`)

	voice := NewVoice(2, "123dsdajkasd231jhksad")

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
		"filter":{
			"group_id":"2"
		},
		"msgtype":"mpvideo",
		"mpvideo":{
			"media_id":"IhdaAQXuvJtGzwwc0abfXnzeezfO0NgPK6AQYShD8RQYMTtfzbLdBIQkQziv2XJc"
		}
	}`)

	video := NewVideo(2, "IhdaAQXuvJtGzwwc0abfXnzeezfO0NgPK6AQYShD8RQYMTtfzbLdBIQkQziv2XJc")

	b, err = json.Marshal(video)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", video, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", video, b, want)
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

	news := NewNews(2, "123dsdajkasd231jhksad")

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
