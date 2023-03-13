package mall

import (
	"github.com/bububa/wechat/mp/core"
)

type DeleteShoppingListRequest struct {
	OpenId         string       `json:"user_open_id"`
	SkuProductList []SkuProduct `json:"sku_product_list"`
}

type SkuProduct struct {
	ItemCode string `json:"item_code"` // 物品ID（SPU ID），要求appid下全局唯一
	SkuId    string `json:"sku_id"`    // 物品sku_id，特殊情况下可以填入与item_code一致
}

// 删除收藏
// 开发者可以在用户从购物车删除物品时，同步将物品数据从好物圈的收藏中删除
func DeleteShoppingList(clt *core.Client, req *DeleteShoppingListRequest) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/mall/deleteshoppinglist?access_token="
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
