package mall

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type AddShoppingListRequest struct {
	OpenId      string    `json:"user_open_id"`
	ProductList []Product `json:"sku_product_list"`
}

// 导入收藏
// 开发者可以在用户添加物品到购物车时，同步物品数据至好物圈的收藏列表。在首次接入时，开发者需导入用户购物车中的所有物品。导入数据有助于物品在搜索中获得更好的曝光，同时也保障用户侧获得一致的体验
func AddShoppingList(clt *core.Client, req *AddShoppingListRequest) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/mall/addshoppinglist?access_token="
	var result struct {
		core.Error
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
