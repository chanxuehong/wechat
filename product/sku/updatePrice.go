package sku

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

// UpdatePrice 更新SKU价格
func UpdatePrice(clt *core.Client, req *model.Sku) (sku *model.Sku, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/sku/update_price?access_token="

	var result struct {
		core.Error
		Data *model.Sku `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	sku = result.Data
	return
}
