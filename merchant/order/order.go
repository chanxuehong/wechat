package order

type Order struct {
	Id           string `json:"order_id"`            // 订单ID
	Status       int    `json:"order_status"`        // 订单状态
	TotalPrice   int    `json:"order_total_price"`   // 订单总价格(单位 : 分)
	CreateTime   int64  `json:"order_create_time"`   // 订单创建时间
	ExpressPrice int    `json:"order_express_price"` // 订单运费价格(单位 : 分)

	BuyerOpenId      string `json:"buyer_openid"`      // 买家微信OPENID
	BuyerNick        string `json:"buyer_nick"`        // 买家微信昵称
	ReceiverName     string `json:"receiver_name"`     // 收货人姓名
	ReceiverProvince string `json:"receiver_province"` // 收货地址省份
	ReceiverCity     string `json:"receiver_city"`     // 收货地址城市
	ReceiverAddress  string `json:"receiver_address"`  // 收货详细地址
	ReceiverMobile   string `json:"receiver_mobile"`   // 收货人移动电话
	ReceiverPhone    string `json:"receiver_phone"`    // 收货人固定电话

	ProductId    string `json:"product_id"`    // 商品ID
	ProductName  string `json:"product_name"`  // 商品名称
	ProductPrice int    `json:"product_price"` // 商品价格(单位 : 分)
	ProductSku   string `json:"product_sku"`   // 商品SKU
	ProductCount int    `json:"product_count"` // 商品个数
	ProductImage string `json:"product_img"`   // 商品图片

	DeliveryId      string `json:"delivery_id"`      // 运单ID
	DeliveryCompany string `json:"delivery_company"` // 物流公司编码

	TransactionId string `json:"trans_id"` // 交易ID
}
