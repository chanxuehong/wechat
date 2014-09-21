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
						"type":"view", 
						"name":"视频", 
						"url":"http://v.qq.com/"
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

	var subButtons = make([]Button, 3)
	subButtons[0].InitToViewButton("搜索", "http://www.soso.com/")
	subButtons[1].InitToViewButton("视频", "http://v.qq.com/")
	subButtons[2].InitToClickButton("赞一下我们", "V1001_GOOD")

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

	want = util.TrimSpace([]byte(`{
		"button":[
			{
				"name":"扫码", 
				"sub_button":[
					{
						"type":"scancode_waitmsg", 
						"name":"扫码带提示", 
						"key":"rselfmenu_0_0"
					}, 
					{
						"type":"scancode_push", 
						"name":"扫码推事件", 
						"key":"rselfmenu_0_1"
					}
				]
			}, 
			{
				"name":"发图", 
				"sub_button":[
					{
						"type":"pic_sysphoto", 
						"name":"系统拍照发图", 
						"key":"rselfmenu_1_0"
					}, 
					{
						"type":"pic_photo_or_album", 
						"name":"拍照或者相册发图", 
						"key":"rselfmenu_1_1"
					}, 
					{
						"type":"pic_weixin", 
						"name":"微信相册发图", 
						"key":"rselfmenu_1_2"
					}
				]
			}, 
			{
				"type":"location_select",
				"name":"发送位置", 
				"key":"rselfmenu_2_0"
			}
		]
	}`))

	var subButtons0 = make([]Button, 2)
	subButtons0[0].InitToScanCodeWaitMsgButton("扫码带提示", "rselfmenu_0_0")
	subButtons0[1].InitToScanCodePushButton("扫码推事件", "rselfmenu_0_1")

	var subButtons1 = make([]Button, 3)
	subButtons1[0].InitToPicSysPhotoButton("系统拍照发图", "rselfmenu_1_0")
	subButtons1[1].InitToPicPhotoOrAlbumButton("拍照或者相册发图", "rselfmenu_1_1")
	subButtons1[2].InitToPicWeixinButton("微信相册发图", "rselfmenu_1_2")

	var mn1 Menu
	mn1.Buttons = make([]Button, 3)
	mn1.Buttons[0].InitToSubMenuButton("扫码", subButtons0)
	mn1.Buttons[1].InitToSubMenuButton("发图", subButtons1)
	mn1.Buttons[2].InitToLocationSelectButton("发送位置", "rselfmenu_2_0")

	have, err = json.Marshal(mn1)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", mn1, err)
	} else if !bytes.Equal(have, want) {
		t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", mn1, have, want)
	}
}

func TestMenuJSONUnmarshal(t *testing.T) {
	src := []byte(`{
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
						"type":"view", 
						"name":"视频", 
						"url":"http://v.qq.com/"
					}, 
					{
						"type":"click", 
						"name":"赞一下我们", 
						"key":"V1001_GOOD"
					}
				]
			}
		]
	}`)

	var mn0 Menu
	if err := json.Unmarshal(src, &mn0); err != nil {
		t.Errorf("json.Unmarshal(%#q):\nError: %s\n", src, err)
	} else {
		var subButtons1 = make([]Button, 3)
		subButtons1[0].InitToViewButton("搜索", "http://www.soso.com/")
		subButtons1[1].InitToViewButton("视频", "http://v.qq.com/")
		subButtons1[2].InitToClickButton("赞一下我们", "V1001_GOOD")

		var mn0x Menu
		mn0x.Buttons = make([]Button, 2)
		mn0x.Buttons[0].InitToClickButton("今日歌曲", "V1001_TODAY_MUSIC")
		mn0x.Buttons[1].InitToSubMenuButton("菜单", subButtons1)

		if !menuEqual(mn0.Buttons, mn0x.Buttons) {
			t.Errorf("json.Unmarshal(%#q):\nhave %#q\nwant %#q\n", src, mn0, mn0x)
		}
	}

	src = []byte(`{
		"button":[
			{
				"name":"扫码", 
				"sub_button":[
					{
						"type":"scancode_waitmsg", 
						"name":"扫码带提示", 
						"key":"rselfmenu_0_0"
					}, 
					{
						"type":"scancode_push", 
						"name":"扫码推事件", 
						"key":"rselfmenu_0_1"
					}
				]
			}, 
			{
				"name":"发图", 
				"sub_button":[
					{
						"type":"pic_sysphoto", 
						"name":"系统拍照发图", 
						"key":"rselfmenu_1_0"
					}, 
					{
						"type":"pic_photo_or_album", 
						"name":"拍照或者相册发图", 
						"key":"rselfmenu_1_1"
					}, 
					{
						"type":"pic_weixin", 
						"name":"微信相册发图", 
						"key":"rselfmenu_1_2"
					}
				]
			}, 
			{
				"type":"location_select",
				"name":"发送位置", 
				"key":"rselfmenu_2_0"
			}
		]
	}`)

	var mn1 Menu
	if err := json.Unmarshal(src, &mn1); err != nil {
		t.Errorf("json.Unmarshal(%#q):\nError: %s\n", src, err)
	} else {
		var subButtons0 = make([]Button, 2)
		subButtons0[0].InitToScanCodeWaitMsgButton("扫码带提示", "rselfmenu_0_0")
		subButtons0[1].InitToScanCodePushButton("扫码推事件", "rselfmenu_0_1")

		var subButtons1 = make([]Button, 3)
		subButtons1[0].InitToPicSysPhotoButton("系统拍照发图", "rselfmenu_1_0")
		subButtons1[1].InitToPicPhotoOrAlbumButton("拍照或者相册发图", "rselfmenu_1_1")
		subButtons1[2].InitToPicWeixinButton("微信相册发图", "rselfmenu_1_2")

		var mn1x Menu
		mn1x.Buttons = make([]Button, 3)
		mn1x.Buttons[0].InitToSubMenuButton("扫码", subButtons0)
		mn1x.Buttons[1].InitToSubMenuButton("发图", subButtons1)
		mn1x.Buttons[2].InitToLocationSelectButton("发送位置", "rselfmenu_2_0")

		if !menuEqual(mn1.Buttons, mn1x.Buttons) {
			t.Errorf("json.Unmarshal(%#q):\nhave %#q\nwant %#q\n", src, mn1, mn1x)
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
