package coupon

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

type GetListRequest struct {
	StartCreateTime string             `json:"start_create_time,omitempty"`
	EndCreateTime   string             `json:"end_create_time,omitempty"`
	Status          model.CouponStatus `json:"status"`
	Page            int                `json:"page"`
	PageSize        int                `json:"page_size"`
}

// GetList 获取优惠券列表
func GetList(clt *core.Client, req *GetListRequest) (coupons []model.Coupon, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/coupon/get_list?access_token="

	var result struct {
		core.Error
		List  []model.Coupon `json:"coupons"`
		Total int            `json:"total_num"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	coupons = result.List
	return
}
