package mall

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 小程序的物品是否可被搜索
// 接入好物圈的小程序默认打开物品搜索，开发者可通过此接口进行调整。当“小程序打开物品搜索”且“物品的状态为可以被搜索”的情况下，小程序的物品可被搜索。
func SetCanBeSearch(clt *core.Client, canBeSearch bool) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/mall/brandmanage?action=set_biz_can_be_search&access_token="
	var result struct {
		core.Error
	}
	req := map[string]bool{
		"can_be_search": canBeSearch,
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
