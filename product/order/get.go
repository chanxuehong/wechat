package order

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

// Get 获取订单详情
func Get(clt *core.Client, orderId uint64) (order *model.Order, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/order/get?access_token="

	req := map[string]uint64{
		"order_id": orderId,
	}

	var result struct {
		core.Error
		Order *model.Order `json:"order"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	order = result.Order
	return
}
