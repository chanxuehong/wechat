package service

// 服务类型
type ServiceType = int

const (
	PRODUCT_SERVICE  ServiceType = 2 // 商品
	ORDER_SERVICE    ServiceType = 3 // 订单
	DELIVERY_SERVICE ServiceType = 4 // 物流
)

type Service struct {
	Id         uint64      `json:"service_id,omitempty"`       // 服务ID
	SpecificId string      `json:"specification_id,omitempty"` // 规格英文名称
	OrderId    uint64      `json:"service_order_id,omitempty"` // 服务订单ID
	Name       string      `json:"service_name,omitempty"`     // 服务名称
	CreateTime string      `json:"create_time,omitempty"`      // 创建时间
	ExpireTime string      `json:"expire_time,omitempty"`      // 过期时间
	TotalPrice uint        `json:"total_price,omitempty"`      // 订单总价格
	Type       ServiceType `json:"service_type,omitempty"`     // 服务类型
}
