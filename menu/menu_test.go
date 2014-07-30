// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package menu

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestMenuMarshal(t *testing.T) {
	expectBytes := []byte(`{
		"button":[
			{
				"name":"今日歌曲",
				"type":"click",
				"key":"V1001_TODAY_MUSIC"
			},
			{
				"name":"歌手简介",
				"type":"click",
				"key":"V1001_TODAY_SINGER"
			},
			{
				"name":"菜单",
				"sub_button":[
					{
						"name":"搜索",
						"type":"view",
						"url":"http://www.soso.com/"
					},
					{
						"name":"视频",
						"type":"view",
						"url":"http://v.qq.com/"
					},
					{
						"name":"赞一下我们",
						"type":"click",
						"key":"V1001_GOOD"
					}
				]
			}
		]
	}`)

	var buttons = make([]Button, 3)
	buttons[0].InitToViewButton("搜索", "http://www.soso.com/")
	buttons[1].InitToViewButton("视频", "http://v.qq.com/")
	buttons[2].InitToClickButton("赞一下我们", "V1001_GOOD")

	var mn Menu
	mn.Buttons = make([]Button, 3)
	mn.Buttons[0].InitToClickButton("今日歌曲", "V1001_TODAY_MUSIC")
	mn.Buttons[1].InitToClickButton("歌手简介", "V1001_TODAY_SINGER")
	mn.Buttons[2].InitToSubMenuButton("菜单", buttons)

	b, err := json.Marshal(mn)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", mn, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", mn, b, want)
		}
	}
}

func TestMenuUnmarshal(t *testing.T) {
	src := []byte(`{
		"button": [
			{
				"type": "click",
				"name": "今日歌曲",
				"key": "V1001_TODAY_MUSIC",
				"sub_button": []
			},
			{
				"type": "click",
				"name": "歌手简介",
				"key": "V1001_TODAY_SINGER",
				"sub_button": []
			},
			{
				"name": "菜单",
				"sub_button": [
					{
						"type": "view",
						"name": "搜索",
						"url": "http://www.soso.com/",
						"sub_button": []
					},
					{
						"type": "view",
						"name": "视频",
						"url": "http://v.qq.com/",
						"sub_button": []
					},
					{
						"type": "click",
						"name": "赞一下我们",
						"key": "V1001_GOOD",
						"sub_button": []
					}
				]
			}
		]
	}`)

	var mn Menu
	if err := json.Unmarshal(src, &mn); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", src, err)
	} else {
		var buttons = make([]Button, 3)
		buttons[0].InitToViewButton("搜索", "http://www.soso.com/")
		buttons[1].InitToViewButton("视频", "http://v.qq.com/")
		buttons[2].InitToClickButton("赞一下我们", "V1001_GOOD")

		var mn1 Menu
		mn1.Buttons = make([]Button, 3)
		mn1.Buttons[0].InitToClickButton("今日歌曲", "V1001_TODAY_MUSIC")
		mn1.Buttons[1].InitToClickButton("歌手简介", "V1001_TODAY_SINGER")
		mn1.Buttons[2].InitToSubMenuButton("菜单", buttons)

		if !menuEqual(mn1.Buttons, mn.Buttons) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", src, mn, mn1)
		}
	}
}

func menuEqual(src, to []Button) bool {
	if len(src) != len(to) {
		return false
	}
	for i := 0; i < len(src); i++ {
		if src[i].Name != to[i].Name {
			return false
		}
		if src[i].Type != to[i].Type {
			return false
		}
		if src[i].Key != to[i].Key {
			return false
		}
		if src[i].URL != to[i].URL {
			return false
		}

		if !menuEqual(src[i].SubButtons, to[i].SubButtons) {
			return false
		}
	}

	return true
}
