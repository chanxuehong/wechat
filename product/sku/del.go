package sku

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

// Del 删除SKU
func Del(clt *core.Client, sku *model.Sku) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/sku/del?access_token="

	var result core.Error
	if err = clt.PostJSON(incompleteURL, sku, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
