package order

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

type GetListRequest struct {
	StartCreateTime string            `json:"start_create_time,omitempty"`
	EndCreateTime   string            `json:"end_create_time,omitempty"`
	StartUpdateTime string            `json:"start_update_time,omitempty"`
	EndUpdateTime   string            `json:"end_update_time,omitempty"`
	Status          model.OrderStatus `json:"status"`
	Page            int               `json:"page"`
	PageSize        int               `json:"page_size"`
}

// GetList 获取订单列表
func GetList(clt *core.Client, req *GetListRequest) (total int, orders []model.Order, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/order/get_list?access_token="

	var result struct {
		core.Error
		List  []model.Order `json:"orders"`
		Total int           `json:"total_num"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	total = result.Total
	orders = result.List
	return
}
