package spu

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

// Add 添加商品
func Add(clt *core.Client, spu *model.Spu) (product *model.Spu, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/spu/add?access_token="

	var result struct {
		core.Error
		Data *model.Spu `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, spu, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	product = result.Data
	return
}
