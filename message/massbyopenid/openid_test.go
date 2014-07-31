// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package massbyopenid

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
		"touser":[
			"oR5Gjjl_eiZoUpGozMo7dbBJ362A",
			"oR5Gjjo5rXlMUocSEXKT7Q5RQ63Q"
		],
		"msgtype":"text", 
		"text":{
			"content":"hello from boxer."
		}
	}`)

	text := NewText([]string{"oR5Gjjl_eiZoUpGozMo7dbBJ362A", "oR5Gjjo5rXlMUocSEXKT7Q5RQ63Q"},
		"hello from boxer.")

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
		"touser":[
			"OPENID1",
			"OPENID2"
		],
		"msgtype":"image",
		"image":{
			"media_id":"BTgN0opcW3Y5zV_ZebbsD3NFKRWf6cb7OPswPi9Q83fOJHK2P67dzxn11Cp7THat"
		}
	}`)

	image := NewImage([]string{"OPENID1", "OPENID2"},
		"BTgN0opcW3Y5zV_ZebbsD3NFKRWf6cb7OPswPi9Q83fOJHK2P67dzxn11Cp7THat")

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
		"touser":[
			"OPENID1",
			"OPENID2"
		],
		"msgtype":"voice",
		"voice":{
			"media_id":"mLxl6paC7z2Tl-NJT64yzJve8T9c8u9K2x-Ai6Ujd4lIH9IBuF6-2r66mamn_gIT"
		}
	}`)

	voice := NewVoice([]string{"OPENID1", "OPENID2"},
		"mLxl6paC7z2Tl-NJT64yzJve8T9c8u9K2x-Ai6Ujd4lIH9IBuF6-2r66mamn_gIT")

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
		"touser":[
			"OPENID1",
			"OPENID2"
		],
		"msgtype":"video",
		"video":{
			"media_id":"123dsdajkasd231jhksad",
			"title":"TITLE",
			"description":"DESCRIPTION"
		}
	}`)

	video := NewVideo(
		[]string{"OPENID1", "OPENID2"},
		"123dsdajkasd231jhksad",
		"TITLE",
		"DESCRIPTION",
	)

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
		"touser":[
			"OPENID1",
			"OPENID2"
		],
		"msgtype":"mpnews",
		"mpnews":{
			"media_id":"123dsdajkasd231jhksad"
		}
	}`)

	news := NewNews(
		[]string{"OPENID1", "OPENID2"},
		"123dsdajkasd231jhksad",
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
