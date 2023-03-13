package sku

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

// batchAdd 批量添加商品
func batchAdd(clt *core.Client, reqSkus []model.Sku) (skus []model.Sku, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/sku/batch_add?access_token="

	req := map[string][]model.Sku{
		"skus": reqSkus,
	}
	var result struct {
		core.Error
		Data []model.Sku `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	skus = result.Data
	return
}
