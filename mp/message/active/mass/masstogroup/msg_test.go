// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package masstogroup

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestMarshalAndNewFunc(t *testing.T) {

	// 测试文本消息===============================================================

	want := util.TrimSpace([]byte(`{
	    "filter":{
	        "group_id":"2"
	    }, 
	    "msgtype":"text", 
	    "text":{
	        "content":"CONTENT"
	    }
	}`))

	text := NewText(2, "CONTENT")

	have, err := json.Marshal(text)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", text, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", text, have, want)
	}

	// 测试图片消息===============================================================

	want = util.TrimSpace([]byte(`{
	    "filter":{
	        "group_id":"2"
	    }, 
	    "msgtype":"image", 
	    "image":{
	        "media_id":"123dsdajkasd231jhksad"
	    }
	}`))

	image := NewImage(2, "123dsdajkasd231jhksad")

	have, err = json.Marshal(image)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", image, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", image, have, want)
	}

	// 测试语音消息===============================================================

	want = util.TrimSpace([]byte(`{
	    "filter":{
	        "group_id":"2"
	    }, 
	    "msgtype":"voice", 
	    "voice":{
	        "media_id":"123dsdajkasd231jhksad"
	    }
	}`))

	voice := NewVoice(2, "123dsdajkasd231jhksad")

	have, err = json.Marshal(voice)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", voice, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", voice, have, want)
	}

	// 测试视频消息===============================================================

	want = util.TrimSpace([]byte(`{
	    "filter":{
	        "group_id":"2"
	    }, 
	    "msgtype":"mpvideo", 
	    "mpvideo":{
	        "media_id":"IhdaAQXuvJtGzwwc0abfXnzeezfO0NgPK6AQYShD8RQYMTtfzbLdBIQkQziv2XJc"
	    }
	}`))

	video := NewVideo(2, "IhdaAQXuvJtGzwwc0abfXnzeezfO0NgPK6AQYShD8RQYMTtfzbLdBIQkQziv2XJc")

	have, err = json.Marshal(video)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", video, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", video, have, want)
	}

	// 测试图文消息===============================================================

	want = util.TrimSpace([]byte(`{
	    "filter":{
	        "group_id":"2"
	    }, 
	    "msgtype":"mpnews", 
	    "mpnews":{
	        "media_id":"123dsdajkasd231jhksad"
	    }
	}`))

	news := NewNews(2, "123dsdajkasd231jhksad")

	have, err = json.Marshal(news)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", news, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", news, have, want)
	}
}
