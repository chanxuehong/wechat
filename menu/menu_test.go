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

	menuButton0 := MenuButton{
		Name: "今日歌曲",
		Type: MENUBUTTON_TYPE_CLICK,
		Key:  "V1001_TODAY_MUSIC",
	}
	menuButton1 := MenuButton{
		Name: "歌手简介",
		Type: MENUBUTTON_TYPE_CLICK,
		Key:  "V1001_TODAY_SINGER",
	}
	menuButton2 := MenuButton{
		Name: "菜单",
		SubButton: []*MenuButton{
			&MenuButton{
				Name: "搜索",
				Type: MENUBUTTON_TYPE_VIEW,
				URL:  "http://www.soso.com/",
			},
			&MenuButton{
				Name: "视频",
				Type: MENUBUTTON_TYPE_VIEW,
				URL:  "http://v.qq.com/",
			},
			&MenuButton{
				Name: "赞一下我们",
				Type: MENUBUTTON_TYPE_CLICK,
				Key:  "V1001_GOOD",
			},
		},
	}

	var _menu Menu
	_menu.AppendButton(&menuButton0, &menuButton1, &menuButton2)

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

	var _menu Menu
	if err := json.Unmarshal(src, &_menu); err != nil {
		t.Errorf("unmarshal(%#q):\nError: %s\n", src, err)
	} else {
		expectMenu := Menu{
			Button: []*MenuButton{
				&MenuButton{
					Name: "今日歌曲",
					Type: MENUBUTTON_TYPE_CLICK,
					Key:  "V1001_TODAY_MUSIC",
				},
				&MenuButton{
					Name: "歌手简介",
					Type: MENUBUTTON_TYPE_CLICK,
					Key:  "V1001_TODAY_SINGER",
				},
				&MenuButton{
					Name: "菜单",
					SubButton: []*MenuButton{
						&MenuButton{
							Name: "搜索",
							Type: MENUBUTTON_TYPE_VIEW,
							URL:  "http://www.soso.com/",
						},
						&MenuButton{
							Name: "视频",
							Type: MENUBUTTON_TYPE_VIEW,
							URL:  "http://v.qq.com/",
						},
						&MenuButton{
							Name: "赞一下我们",
							Type: MENUBUTTON_TYPE_CLICK,
							Key:  "V1001_GOOD",
						},
					},
				},
			},
		}

		isEqual := func(src, to Menu) bool {
			if len(_menu.Button) != 3 {
				return false
			}
			if len(_menu.Button[2].SubButton) != 3 {
				return false
			}

			if _menu.Button[0].Name != expectMenu.Button[0].Name ||
				_menu.Button[0].Type != expectMenu.Button[0].Type ||
				_menu.Button[0].Key != expectMenu.Button[0].Key {
				return false
			}
			if _menu.Button[1].Name != expectMenu.Button[1].Name ||
				_menu.Button[1].Type != expectMenu.Button[1].Type ||
				_menu.Button[1].Key != expectMenu.Button[1].Key {
				return false
			}

			if _menu.Button[2].Name != expectMenu.Button[2].Name {
				return false
			}
			if _menu.Button[2].SubButton[0].Name != expectMenu.Button[2].SubButton[0].Name ||
				_menu.Button[2].SubButton[0].Type != expectMenu.Button[2].SubButton[0].Type ||
				_menu.Button[2].SubButton[0].URL != expectMenu.Button[2].SubButton[0].URL {
				return false
			}
			if _menu.Button[2].SubButton[1].Name != expectMenu.Button[2].SubButton[1].Name ||
				_menu.Button[2].SubButton[1].Type != expectMenu.Button[2].SubButton[1].Type ||
				_menu.Button[2].SubButton[1].URL != expectMenu.Button[2].SubButton[1].URL {
				return false
			}
			if _menu.Button[2].SubButton[2].Name != expectMenu.Button[2].SubButton[2].Name ||
				_menu.Button[2].SubButton[2].Type != expectMenu.Button[2].SubButton[2].Type ||
				_menu.Button[2].SubButton[2].Key != expectMenu.Button[2].SubButton[2].Key {
				return false
			}

			return true
		}(_menu, expectMenu)

		if !isEqual {
			t.Errorf("unmarshal(%#q):\nhave %#q\nwant %#q\n", src, _menu, expectMenu)
		}
	}
}
