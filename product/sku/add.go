package sku

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

// Add 添加商品
func Add(clt *core.Client, req *model.Sku) (sku *model.Sku, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/sku/add?access_token="

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
