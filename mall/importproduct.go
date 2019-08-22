package mall

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type ImportProductRequest struct {
	ProductList []Product `json:"product_list"`
}

// 更新或导入物品信息
// 开发者可以对好物圈收藏/搜索场景下物品信息进行导入或更新，如上架状态改变、物品售罄、价格更新等。如果物品仅支持到店提货或到家送货，poi_list必填；如果物品同时支持线上物流配送，该字段应为空。
func ImportProduct(clt *core.Client, req *ImportProductRequest) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/mall/importproduct?access_token="
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
