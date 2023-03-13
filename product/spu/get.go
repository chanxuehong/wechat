package spu

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

// Get 获取商品
func Get(clt *core.Client, spuId uint64, outId string) (product *model.Spu, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/spu/get?access_token="

	req := struct {
		Id    uint64 `json:"product_id,omitempty"`
		OutId string `json:"out_product_id,omitempty"`
	}{
		Id:    spuId,
		OutId: outId,
	}

	var result struct {
		core.Error
		Data struct {
			Spu *model.Spu `json:"spu"`
		} `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	product = result.Data.Spu
	return
}
