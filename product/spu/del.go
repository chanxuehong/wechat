package spu

import (
	"github.com/bububa/wechat/product/core"
)

// Del 删除商品
func Del(clt *core.Client, spuId uint64, outId string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/spu/del?access_token="

	req := struct {
		Id    uint64 `json:"product_id,omitempty"`
		OutId string `json:"out_product_id,omitempty"`
	}{
		Id:    spuId,
		OutId: outId,
	}

	var result core.Error
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
