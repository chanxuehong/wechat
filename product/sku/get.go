package sku

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

// Get 获取SKU
func Get(clt *core.Client, req *model.Sku) (sku *model.Sku, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/sku/get?access_token="

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
