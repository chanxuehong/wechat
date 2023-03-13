package mall

import (
	"github.com/bububa/wechat/mp/core"
)

// 删除用户的所有收藏
// 开发者可以删除用户在好物圈中指定商家的所有收藏物品
func DeleteAllShoppingList(clt *core.Client, openId string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/mall/deletebizallshoppinglist?access_token="
	var result struct {
		core.Error
	}
	req := map[string]string{
		"user_open_id": openId,
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
