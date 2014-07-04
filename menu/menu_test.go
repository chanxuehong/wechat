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

	button0 := NewClickButton("今日歌曲", "V1001_TODAY_MUSIC")
	button1 := NewClickButton("歌手简介", "V1001_TODAY_SINGER")
	button2 := NewSubMenuButton("菜单", nil)
	button2.SubMenu = append(button2.SubMenu, NewViewButton("搜索", "http://www.soso.com/"))
	button2.SubMenu = append(button2.SubMenu, NewViewButton("视频", "http://v.qq.com/"))
	button2.SubMenu = append(button2.SubMenu, NewClickButton("赞一下我们", "V1001_GOOD"))

	var _menu struct {
		Menu `json:"button"`
	}
	_menu.Menu = make([]*Button, 3)
	_menu.Menu[0] = button0
	_menu.Menu[1] = button1
	_menu.Menu[2] = button2

	b, err := json.Marshal(_menu)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", _menu, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", _menu, b, want)
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

	var _menu struct {
		Menu `json:"button"`
	}
	if err := json.Unmarshal(src, &_menu); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", src, err)
	} else {
		button0 := NewClickButton("今日歌曲", "V1001_TODAY_MUSIC")
		button1 := NewClickButton("歌手简介", "V1001_TODAY_SINGER")
		button2 := NewSubMenuButton("菜单", nil)
		button2.SubMenu = append(button2.SubMenu, NewViewButton("搜索", "http://www.soso.com/"))
		button2.SubMenu = append(button2.SubMenu, NewViewButton("视频", "http://v.qq.com/"))
		button2.SubMenu = append(button2.SubMenu, NewClickButton("赞一下我们", "V1001_GOOD"))

		var _menu1 struct {
			Menu `json:"button"`
		}
		_menu1.Menu = make([]*Button, 3)
		_menu1.Menu[0] = button0
		_menu1.Menu[1] = button1
		_menu1.Menu[2] = button2

		if !menuEqual(_menu1.Menu, _menu.Menu) {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", src, _menu, _menu1)
		}
	}
}

func menuEqual(src, to Menu) bool {
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

		if !menuEqual(src[i].SubMenu, to[i].SubMenu) {
			return false
		}
	}

	return true
}
