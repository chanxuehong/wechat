package menu

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 创建个性化菜单.
func AddConditionalMenu(clt *core.Client, menu *Menu) (menuId int64, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/addconditional?access_token="

	var result struct {
		core.Error
		MenuId int64 `json:"menuId"`
	}
	if err = clt.PostJSON(incompleteURL, menu, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	menuId = result.MenuId
	return
}

// 删除个性化菜单.
func DeleteConditionalMenu(clt *core.Client, menuId int64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/delconditional?access_token="

	var request = struct {
		MenuId int64 `json:"menuid"`
	}{
		MenuId: menuId,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 测试个性化菜单匹配结果.
//  userId 可以是粉丝的 OpenID, 也可以是粉丝的微信号
func TryMatch(clt *core.Client, userId string) (menu *Menu, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/trymatch?access_token="

	var request = struct {
		UserId string `json:"user_id"`
	}{
		UserId: userId,
	}
	var result struct {
		core.Error
		Menu `json:"menu"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = &result.Menu
	return
}
