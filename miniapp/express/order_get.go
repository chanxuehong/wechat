package express

import (
	"github.com/bububa/wechat/component/core"
)

// OrderGetRequest 获取运单数据 API Request
type OrderGetRequest struct {
	// OrderID 订单 ID，需保证全局唯一
	OrderID string `json:"order_id,omitempty"`
	// OpenID 该参数仅在 getOrder 接口生效，batchGetOrder接口不生效。用户openid，当add_source=2时无需填写（不发送物流服务通知）
	OpenID string `json:"open_id,omitempty"`
	// DeliveryID 快递公司ID，参见 getAllDelivery , 必须和waybill_id对应
	DeliveryID string `json:"delivery_id,omitempty"`
	// WaybillID 运单ID
	WaybillID string `json:"waybill_id,omitempty"`
	// PrintType 该参数仅在 getOrder 接口生效，batchGetOrder接口不生效。获取打印面单类型【1：一联单，0：二联单】，默认获取二联单
	PrintType int `json:"print_type,omitempty"`
	// CustomRemark
	CustomRemark string `json:"custom_remark,omitempty"`
}

// OrderGet 获取运单数据
func OrderGet(clt *core.Client, req *OrderGetRequest) (order *Order, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/express/business/order/get?access_token="
	var result struct {
		core.Error
		Order
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	order = &result.Order
	return
}
