package express

import (
	"github.com/bububa/wechat/mp/core"
)

// Delivery 快递公司
type Delivery struct {
	// DeliveryID 快递公司ID
	DeliveryID string `json:"delivery_id,omitempty"`
	// DeliveryName 快递公司名称
	DeliveryName string `json:"delivery_name,omitempty"`
	// CanUseCash 是否支持散单, 1表示支持
	CanUseCash int `json:"can_use_cash,omitempty"`
	// CanGetQuota 是否支持查询面单余额, 1表示支持
	CanGetQuota int `json:"can_get_quota,omitempty"`
	// CashBizID 散单对应的bizid，当can_use_cash=1时有效
	CashBizID string `json:"cash_biz_id,omitempty"`
	// ServiceType 支持的服务类型
	ServiceType []ServiceType `json:"service_type,omitempty"`
}

// DeliveryGetAll 获取支持的快递公司列表
func DeliveryGetAll(clt *core.Client) (list []Delivery, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/express/business/delivery/getall?access_token="
	var result struct {
		core.Error
		// Data 快递公司信息列表
		Data []Delivery `json:"data,omitempty"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.Data
	return
}
