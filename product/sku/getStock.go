package sku

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

// GetStock 获取SKU库存
func GetStock(clt *core.Client, req *model.Sku) (stockNum int, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/sku/stock/get?access_token="

	var result struct {
		core.Error
		Data struct {
			StockNum int `json:"stock_num"`
		} `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	stockNum = result.Data.StockNum
	return
}
