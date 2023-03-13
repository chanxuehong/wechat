package delivery

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

// Send 订单发货
func Send(clt *core.Client, orderId uint64, deliveries []model.DeliveryProduct) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/delivery/send?access_token="

	req := struct {
		OrderId    uint64                  `json:"order_id"`
		Deliveries []model.DeliveryProduct `json:"delivery_list"`
	}{
		OrderId:    orderId,
		Deliveries: deliveries,
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
