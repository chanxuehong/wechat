package menu

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 创建自定义菜单.
func Create(clt *core.Client, menu *Menu) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/create?access_token="

	var result core.Error
	if err = clt.PostJSON(incompleteURL, menu, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 查询自定义菜单.
func Get(clt *core.Client) (menu *Menu, conditionalMenus []Menu, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/get?access_token="

	var result struct {
		core.Error
		Menu             Menu   `json:"menu"`
		ConditionalMenus []Menu `json:"conditionalmenu"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	menu = &result.Menu
	conditionalMenus = result.ConditionalMenus
	return
}

// 删除自定义菜单.
func Delete(clt *core.Client) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/menu/delete?access_token="

	var result core.Error
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
