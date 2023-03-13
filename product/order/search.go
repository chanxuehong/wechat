package order

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

type SearchRequest struct {
	StartPayTime          string            `json:"start_pay_time,omitempty"` // 订单支付时间的搜索开始时间
	EndPayTime            string            `json:"end_pay_time,omitempty"`   // 订单支付时间的搜索结束时间
	Title                 string            `json:"title,omitempty"`          // 商品标题关键词
	SkuCode               string            `json:"sku_code,omitempty"`       // 商品编码
	UserName              string            `json:"user_name,omitempty"`      // 收件人
	Mobile                string            `json:"tel_number,omitempty"`     // 收件人电话
	OnAftersaleOrderExist *int              `json:"on_aftersale_order_exist"` // 不填该参数:全部订单 0:没有正在售后的订单, 1:正在售后单数量大于等于1的订单
	Status                model.OrderStatus `json:"status,omitempty"`         // 订单状态
	Page                  int               `json:"page"`
	PageSize              int               `json:"page_size"`
}

// Search 搜索订单
func Search(clt *core.Client, req *SearchRequest) (total int, orders []model.Order, err error) {
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
