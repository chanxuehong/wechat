package service

import (
	"github.com/bububa/wechat/product/core"
)

type GetOrderListRequest struct {
	StartCreateTime string `json:"start_create_time"`
	EndCreateTime   string `json:"end_create_time"`
	Page            int    `json:"page"`
	PageSize        int    `json:"page_size"`
}

// GetOrderList 获取用户购买的服务列表
func GetOrderList(clt *core.Client, req *GetOrderListRequest) (services []Service, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/service/get_order_list?access_token="

	var result struct {
		core.Error
		List []Service `json:"service_order_list"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	services = result.List
	return
}
