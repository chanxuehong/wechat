package spu

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

// GetList 获取商品列表
// status: 商品状态
// page: 第几页（最小填1）
// pageSize: 每页数量(不超过10,000)
func GetList(clt *core.Client, status int, page int, pageSize int) (products []model.Spu, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/spu/get_list?access_token="

	req := struct {
		Status   int `json:"status"`
		Page     int `json:"page"`
		PageSize int `json:"page_size"`
	}{
		Status:   status,
		Page:     page,
		PageSize: pageSize,
	}

	var result struct {
		core.Error
		List []model.Spu `json:"spus"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	products = result.List
	return
}
