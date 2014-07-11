// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"github.com/chanxuehong/wechat/menu"
	"testing"
)

func TestMenu(t *testing.T) {
	// 获取老的菜单
	oldMenu, _ := _test_client.MenuGet()

	// 首先删除原来的菜单
	err := _test_client.MenuDelete()
	if err != nil {
		t.Error(err)
		return
	}

	// 应该出错
	_, err = _test_client.MenuGet()
	if err == nil {
		t.Error("菜单已经删除, 这里应该出错")
		return
	}

	// 创建菜单
	mn0 := menu.Menu{
		Buttons: make([]*menu.Button, 3),
	}

	mn0.Buttons[0] = menu.NewClickButton("click", "key")
	mn0.Buttons[1] = menu.NewViewButton("view", "http://www.qq.com")

	subButtons := make([]*menu.Button, 2)
	subButtons[0] = menu.NewClickButton("sub_click", "key")
	subButtons[1] = menu.NewViewButton("sub_view", "http://www.qq.com")

	mn0.Buttons[2] = menu.NewSubMenuButton("sub_menu", subButtons)

	err = _test_client.MenuCreate(mn0)
	if err != nil {
		t.Error(err)
		return
	}

	// 再次获取菜单
	mn1, err := _test_client.MenuGet()
	if err != nil {
		t.Error(err)
		return
	}

	if !menuEqual(mn0.Buttons, mn1.Buttons) {
		t.Error("创建的菜单和获取的菜单不一致")
		return
	}

	// 删除现在的菜单
	err = _test_client.MenuDelete()
	if err != nil {
		t.Error(err)
		return
	}

	// 还原原来的菜单
	err = _test_client.MenuCreate(oldMenu)
	if err != nil {
		t.Error(err)
		return
	}
}

func menuEqual(src, to []*menu.Button) bool {
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
