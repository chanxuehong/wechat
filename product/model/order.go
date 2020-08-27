package model

// 订单状态
type OrderStatus = int

const (
	PENDING_PAY_ORDER      OrderStatus = 10  // 待付款
	PENDING_DELIVERY_ORDER OrderStatus = 20  // 待发货
	PENDING_RECEIVE_ORDER  OrderStatus = 30  // 待收货
	COMPLETE_ORDER         OrderStatus = 100 // 完成
	AFTERSALE_CANCEL_ORDER OrderStatus = 200 // 全部商品售后之后，订单取消
	CANCEL_ORDER           OrderStatus = 250 // 用户主动取消或待付款超时取消
)

type Order struct {
	Id         uint64          `json:"order_id,omitempty"`    // 订单ID
	Status     OrderStatus     `json:"status,omitempty"`      // 订单状态
	CreateTime string          `json:"create_time,omitempty"` // 创建时间
	UpdateTime string          `json:"update_time,omitempty"` // 更新时间
	Detail     *OrderDetail    `json:"order_detail,omitempty"`
	Pay        *OrderPay       `json:"pay_info,omitempty"`
	Price      *OrderPrice     `json:"price_info,omitempty"`
	Delivery   *OrderDelivery  `json:"delivery_info,omitempty"`
	Aftersale  *OrderAftersale `json:"aftersale_detail,omitempty"`
	OpenId     string          `json:"openid,omitempty"` // 用户的openid，用于物流助手接口
}

type OrderDetail struct {
	Products []OrderProduct `json:"product_infos,omitempty"`
}

type OrderProduct struct {
	SpuId                 uint64      `json:"product_id,omitempty"`               // 小商店内部商品ID
	SkuId                 uint64      `json:"sku_id,omitempty"`                   // 小商店内部skuID
	ThumbImg              string      `json:"thumb_img,omitempty"`                // sku小图
	SkuCnt                int         `json:"sku_cnt,omitempty"`                  // sku数量
	OnAftersaleSkuCnt     int         `json:"on_aftersale_sku_cnt,omitempty"`     // 正在售后/退款流程中的sku数量
	FinishAftersaleSkuCnt int         `json:"finish_aftersale_sku_cnt,omitempty"` // 完成售后/退款的sku数量
	SalePrice             uint        `json:"sale_price,omitempty"`               // 售卖价格,以分为单位
	Attrs                 []Attribute `json:"sku_attrs,omtitempty"`               // 属性自定义用
}

type OrderPay struct {
	Method        string `json:"pay_method,omitempty"`     // 支付方式（目前只有"微信支付"）
	PrepayId      uint64 `json:"prepay_id,omitempty"`      // 预支付ID
	TransactionId uint64 `json:"transaction_id,omitempty"` // 支付订单号
	PrepayTime    string `json:"prepay_time,omitempty"`    // 预付款时间
	PayTime       string `json:"pay_time,omitempty"`       // 付款时间
}

type OrderPrice struct {
	ProductPrice  uint `json:"product_price,omitempty"`  // 商品金额（单位：分）
	OrderPrice    uint `json:"order_price,omitempty"`    // 订单金额（单位：分）
	Freight       uint `json:"freight,omitempty"`        // 运费（单位：分）
	DiscountPrice uint `json:"discount_price,omitempty"` // 优惠金额（单位：分）
	IsDiscounted  bool `json:"is_discounted,omitempty"`  // 是否有优惠（false：无优惠/true：有优惠）
}

type OrderDelivery struct {
	Method       string           `json:"delivery_method,omitempty"` // 快递方式（目前只有"快递"）
	DeliveryTime string           `json:"delivery_time,omitempty"`   // 发货时间
	Product      *DeliveryProduct `json:"delivery_product_info,omitempty"`
	Address      *DeliveryAddress `json:"address_info,omitempty"`
}

type DeliveryProduct struct {
	WaybillId  string `json:"waybill_id,omitempty"`  // 快递单号
	DeliveryId string `json:"delivery_id,omitempty"` // 快递公司编号
}

type DeliveryAddress struct {
	Username     string `json:"username,omitempty"`      // 收货人姓名
	PostalCode   string `json:"postal_code,omitempty"`   // 邮编
	ProvinceName string `json:"province_name,omitempty"` // 国标收货地址第一级地址
	CityName     string `json:"city_name,omitempty"`     // 国标收货地址第二级地址
	CountyName   string `json:"county_name,omitempty"`   // 国标收货地址第三级地址
	Detail       string `json:"detail_info,omitempty"`   // 详细收货地址信息
	NationalCode string `json:"national_code,omitempty"` // 收货地址国家码
	Mobile       string `json:"tel_number,omitempty"`    // 收货人手机号码
}

type OrderAftersale struct {
	OrderList []struct {
		OrderId uint64 `json:"aftersale_order_id,omitempty"`
	} `json:"aftersale_order_list,omitempty"` // 售后单ID（售后接口开放后可用于查询）
	OrderCnt int `json:"on_aftersale_order_cnt,omitempty"` // 正在售后流程的售后单数
}
