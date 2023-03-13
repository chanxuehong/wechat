package express

import (
	"github.com/chanxuehong/wechat/component/core"
)

type OrderCancelRequest struct {
	// OrderID 订单ID，须保证全局唯一，不超过512字节
	OrderID string `json:"order_id,omitempty"`
	// OpenID 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	OpenID string `json:"openid,omitempty"`
	// DeliveryID 快递公司ID，参见getAllDelivery
	DeliveryID string `json:"delivery_id,omitempty"`
	// WaybillID 运单ID
	WaybillID string `json:"waybill_id,omitempty"`
}

// OrderCancelResult 订单取消结果
type OrderCancelResult struct {
	// DeliveryResultCode 快递侧错误码，取消失败时返回
	DeliveryResultCode int `json:"delivery_resultcode,omitempty"`
	// DeliveryResultMsg 快递侧错误信息，取消失败时返回
	DeliveryResultMsg string `json:"delivery_resultmsg,omitempty"`
}

// OrderCancel 取消运单
func OrderCancel(clt *core.Client, req *OrderCancelRequest) (res *OrderCancelResult, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/express/business/order/cancel?access_token="
	var result struct {
		core.Error
		OrderCancelResult
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	res = &result.OrderCancelResult
	return
}
