package menu

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 创建自定义菜单.
func CreateMenu(clt *core.Client, menu Menu) (err error) {
	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="
	if err = clt.PostJSON(incompleteURL, &menu, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 删除自定义菜单
func DeleteMenu(clt *core.Client) (err error) {
	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 获取自定义菜单
func GetMenu(clt *core.Client) (menu Menu, err error) {
	var result struct {
		core.Error
		Menu Menu `json:"menu"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/menu/get?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = result.Menu
	return
}

// 获取自定义菜单配置接口
func GetMenuInfo(clt *core.Client) (info MenuInfo, isMenuOpen bool, err error) {
	var result struct {
		core.Error
		IsMenuOpen int      `json:"is_menu_open"`
		MenuInfo   MenuInfo `json:"selfmenu_info"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/get_current_selfmenu_info?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = result.MenuInfo
	if result.IsMenuOpen == 1 {
		isMenuOpen = true
	}
	return
}
