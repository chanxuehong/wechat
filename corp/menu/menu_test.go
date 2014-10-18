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

func TestMenuJSONMarshal(t *testing.T) {
	want := util.TrimSpace([]byte(`{
		"button":[
			{
				"type":"click", 
				"name":"今日歌曲", 
				"key":"V1001_TODAY_MUSIC"
			}, 
			{
				"name":"菜单", 
				"sub_button":[
					{
						"type":"view", 
						"name":"搜索", 
						"url":"http://www.soso.com/"
					}, 
					{
						"type":"click", 
						"name":"赞一下我们", 
						"key":"V1001_GOOD"
					}
				]
			}
		]
	}`))

	var subButtons = make([]Button, 2)
	subButtons[0].InitToViewButton("搜索", "http://www.soso.com/")
	subButtons[1].InitToClickButton("赞一下我们", "V1001_GOOD")

	var mn Menu
	mn.Buttons = make([]Button, 2)
	mn.Buttons[0].InitToClickButton("今日歌曲", "V1001_TODAY_MUSIC")
	mn.Buttons[1].InitToSubMenuButton("菜单", subButtons)

	have, err := json.Marshal(mn)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", mn, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", mn, have, want)
	}
}

func TestMenuJSONUnmarshal(t *testing.T) {
	src := []byte(`{
		"button": [
			{
				"type": "click", 
				"name": "今日歌曲", 
				"key": "V1001_TODAY_MUSIC"
			}, 
			{
				"name": "菜单", 
				"sub_button": [
					{
						"type": "view", 
						"name": "搜索", 
						"url": "http://www.soso.com/"
					}, 
					{
						"type": "click", 
						"name": "赞一下我们", 
						"key": "V1001_GOOD"
					}
				]
			}
		]
	}`)

	var mn0 Menu
	if err := json.Unmarshal(src, &mn0); err != nil {
		t.Errorf("json.Unmarshal(%#q):\nError: %s\n", src, err)
	} else {
		var subButtons1 = make([]Button, 2)
		subButtons1[0].InitToViewButton("搜索", "http://www.soso.com/")
		subButtons1[1].InitToClickButton("赞一下我们", "V1001_GOOD")

		var mn0x Menu
		mn0x.Buttons = make([]Button, 2)
		mn0x.Buttons[0].InitToClickButton("今日歌曲", "V1001_TODAY_MUSIC")
		mn0x.Buttons[1].InitToSubMenuButton("菜单", subButtons1)

		if !menuEqual(mn0.Buttons, mn0x.Buttons) {
			t.Errorf("json.Unmarshal(%#q):\nhave %#q\nwant %#q\n", src, mn0, mn0x)
		}
	}
}

func menuEqual(btns1, btns2 []Button) bool {
	if len(btns1) != len(btns2) {
		return false
	}
	for i := 0; i < len(btns1); i++ {
		if btns1[i].Name != btns2[i].Name {
			return false
		}
		if btns1[i].Type != btns2[i].Type {
			return false
		}
		if btns1[i].Key != btns2[i].Key {
			return false
		}
		if btns1[i].URL != btns2[i].URL {
			return false
		}

		if !menuEqual(btns1[i].SubButtons, btns2[i].SubButtons) {
			return false
		}
	}

	return true
}
