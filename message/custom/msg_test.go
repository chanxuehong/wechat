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

	// 测试文本消息===============================================================

	want := util.TrimSpace([]byte(`{
	    "touser":"OPENID", 
	    "msgtype":"text", 
	    "text":{
	        "content":"Hello World"
	    }
	}`))

	text := NewText("OPENID", "Hello World")

	have, err := json.Marshal(text)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", text, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", text, have, want)
	}

	// 测试图片消息===============================================================

	want = util.TrimSpace([]byte(`{
	    "touser":"OPENID", 
	    "msgtype":"image", 
	    "image":{
	        "media_id":"MEDIA_ID"
	    }
	}`))

	image := NewImage("OPENID", "MEDIA_ID")

	have, err = json.Marshal(image)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", image, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", image, have, want)
	}

	// 测试语音消息===============================================================

	want = util.TrimSpace([]byte(`{
	    "touser":"OPENID", 
	    "msgtype":"voice", 
	    "voice":{
	        "media_id":"MEDIA_ID"
	    }
	}`))

	voice := NewVoice("OPENID", "MEDIA_ID")

	have, err = json.Marshal(voice)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", voice, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", voice, have, want)
	}

	// 测试视频消息===============================================================

	want = util.TrimSpace([]byte(`{
	    "touser":"OPENID", 
	    "msgtype":"video", 
	    "video":{
	        "media_id":"MEDIA_ID", 
	        "thumb_media_id":"THUMB_MEDIA_ID", 
	        "title":"TITLE", 
	        "description":"DESCRIPTION"
	    }
	}`))

	video := NewVideo("OPENID", "MEDIA_ID", "THUMB_MEDIA_ID", "TITLE", "DESCRIPTION")

	have, err = json.Marshal(video)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", video, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", video, have, want)
	}

	// 测试音乐消息===============================================================

	want = util.TrimSpace([]byte(`{
	    "touser":"OPENID", 
	    "msgtype":"music", 
	    "music":{
	        "title":"MUSIC_TITLE", 
	        "description":"MUSIC_DESCRIPTION", 
	        "musicurl":"MUSIC_URL", 
	        "hqmusicurl":"HQ_MUSIC_URL", 
	        "thumb_media_id":"THUMB_MEDIA_ID"
	    }
	}`))

	music := NewMusic("OPENID", "THUMB_MEDIA_ID", "MUSIC_URL", "HQ_MUSIC_URL", "MUSIC_TITLE", "MUSIC_DESCRIPTION")

	have, err = json.Marshal(music)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", music, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", music, have, want)
	}

	articles := make([]NewsArticle, 2)
	articles[0].Init("Happy Day1", "Is Really A Happy Day1", "URL1", "PIC_URL1")
	articles[1].Init("Happy Day2", "Is Really A Happy Day2", "URL2", "PIC_URL2")

	// 测试图文消息, 没有文章=======================================================

	want = util.TrimSpace([]byte(`{
	    "touser":"OPENID", 
	    "msgtype":"news", 
	    "news":{}
	}`))

	news := NewNews("OPENID", articles[:0])

	have, err = json.Marshal(news)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", news, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", news, have, want)
	}

	// 测试图文消息, 1篇没有文章====================================================

	want = util.TrimSpace([]byte(`{
	    "touser":"OPENID", 
	    "msgtype":"news", 
	    "news":{
	        "articles":[
	            {
	                "title":"Happy Day1", 
	                "description":"Is Really A Happy Day1", 
	                "url":"URL1", 
	                "picurl":"PIC_URL1"
	            }
	        ]
	    }
	}`))

	news = NewNews("OPENID", articles[:1])

	have, err = json.Marshal(news)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", news, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", news, have, want)
	}

	// 测试图文消息, 2篇没有文章====================================================

	want = util.TrimSpace([]byte(`{
	    "touser":"OPENID", 
	    "msgtype":"news", 
	    "news":{
	        "articles":[
	            {
	                "title":"Happy Day1", 
	                "description":"Is Really A Happy Day1", 
	                "url":"URL1", 
	                "picurl":"PIC_URL1"
	            }, 
	            {
	                "title":"Happy Day2", 
	                "description":"Is Really A Happy Day2", 
	                "url":"URL2", 
	                "picurl":"PIC_URL2"
	            }
	        ]
	    }
	}`))

	news = NewNews("OPENID", articles[:2])

	have, err = json.Marshal(news)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", news, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", news, have, want)
	}
}
