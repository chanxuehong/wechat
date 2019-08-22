package mall

type Order struct {
	Id            string   `json:"order_id"`                  // 订单id，需要保证唯一性
	CreateTime    int64    `json:"create_time,omitempty"`     // 订单创建时间，unix时间戳
	PayFinishTime int64    `json:"pay_finish_time,omitempty"` // 支付完成时间，unix时间戳
	Desc          string   `json:"desc,omitempty"`            // 订单备注
	Fee           int      `json:"fee,omitempty"`             // 订单金额，单位：分
	TransId       string   `json:"trans_id,omitempty"`        // 微信支付订单id，对于使用微信支付的订单，该字段必填
	Status        uint     `json:"status"`                    // 订单状态，3：支付完成 4：已发货 5：已退款 100: 已完成
	Ext           OrderExt `json:"ext_info"`                  // 订单扩展信息
}

type OrderExt struct {
	Product       *OrderProductList `json:"product_info,omitempty"`      // 物品相关信息
	Express       *Express          `json:"express_info,omitempty"`      // 快递信息
	Promotion     *Promotion        `json:"promotion_info,omitempty"`    // 订单优惠信息
	Brand         *Brand            `json:"brand_info,omitempty"`        // 商家信息
	Invoice       *Invoice          `json:"invoce_info,omitempty"`       // 发票信息，对于开发票订单，该字段必填
	PaymentMethod uint              `json:"payment_method,omitempty"`    // 订单支付方式，0：未知方式 1：微信支付 2：其他支付方式
	OpenId        string            `json:"user_open_id"`                // 用户的openid
	Page          *DetailPage       `json:"order_detail_page,omitempty"` // 订单详情页（小程序页面）
	TotalFee      int               `json:"total_fee,omitempty"`         // 订单物品合计金额（优惠前金额，不填写的话，平台默认用物品的total fee累加）
}

type OrderProductList struct {
	ItemList []OrderProduct `json:"item_list"` // 包含订单中所有物品的信息
}

type OrderProduct struct {
	ItemCode             string      `json:"item_code"`                        // 物品ID（SPU ID），要求appid下全局唯一
	SkuId                string      `json:"sku_id"`                           // sku_id
	Amount               uint        `json:"amount"`                           // 物品数量
	TotalFee             int         `json:"total_fee"`                        // 物品总价，单位：分
	ThumbUrl             string      `json:"thumb_url"`                        // 物品图片，图片宽度必须大于750px，宽高比建议4:3 - 1:1之间
	Title                string      `json:"title"`                            // 物品名称
	Desc                 string      `json:"desc,omitempty"`                   // 物品详细描述
	UnitPrice            int         `json:"unit_price"`                       // 物品单价（实际售价），单位：分
	OriginalPrice        int         `json:"original_price"`                   // 物品原价，单位：分
	StockAttrInfo        []StockAttr `json:"stock_attr_info,omitempty"`        // 物品属性列表
	CategoryList         []string    `json:"category_list"`                    // 物品类目列表
	Page                 DetailPage  `json:"item_detail_page"`                 // 物品详情页（小程序页面）
	CanBeSearch          bool        `json:"can_be_search"`                    // 物品能否被搜索（默认true可以被搜索）
	BarCodeInfo          *BarCode    `json:"bar_code_info,omitempty"`          // 物品的条形码信息
	PlatformCategoryList []Category  `json:"platform_category_list,omitempty"` // 物品平台类目列表，填写的每个类目必须在好物圈物品类目表列出，多级类目只填最后一级（如完整类目为"运动户外-运动服饰-运动裤"，只需要填"运动裤"的类目ID与类目名）
}

type StockAttr struct {
	Name struct {
		Name string `json:"name"` // 属性名称
	} `json:"attr_name"` // 属性名称
	Value struct {
		Name string `json:"name"` // 属性值
	} `json:"attr_value"` // 属性内容
}

type DetailPage struct {
	Path   string `json:"path,omitempty"`        // 小程序物品详情页跳转链接
	H5Path string `json:"src_h5_path,omitempty"` // h5物品详情页跳转链接
	KfType uint   `json:"kf_type,omitempty"`     // 在线客服类型 1 没有在线客服; 2 微信客服消息; 3 小程序自有客服; 4 公众号h5自有客服
}

type Express struct {
	Name         string           `json:"name,omitempty"`                     // 收件人姓名
	Phone        string           `json:"phone,omitempty"`                    // 收件人联系电话
	Address      string           `json:"address,omitempty"`                  // 收件人地址
	Price        int              `json:"price"`                              // 运费，单位：分
	NationalCode string           `json:"national_code,omitempty"`            // 行政区划代码
	Country      string           `json:"country,omitempty"`                  // 国家
	Province     string           `json:"province,omitempty"`                 // 省份
	City         string           `json:"city,omitempty"`                     // 城市
	District     string           `json:"district,omitempty"`                 // 区
	PackageList  []ExpressPackage `json:"express_package_info_list,omitepty"` // 包裹中的物品信息
}

type ExpressPackage struct {
	CompanyId   int           `json:"express_company_id"`      // 快递公司编号，参见快递公司信息
	CompanyName string        `json:"express_company_name"`    // 快递公司名
	Code        string        `json:"express_code"`            // 快递单号
	ShipTime    int64         `json:"ship_time"`               // 发货时间，unix时间戳
	Page        DetailPage    `json:"express_page"`            // 快递详情页（小程序页面）
	GoodsList   []ExpressGood `json:"express_goods_info_list"` // 包裹物品信息
}

type ExpressGood struct {
	ItemCode string `json:"item_code"` // 物品id
	SkuId    string `json:"sku_id"`    // SkuID
}

type Promotion struct {
	Discount int `json:"discount"` // 优惠金额
}

type Invoice struct {
	Type           uint       `json:"type"`                          // 抬头类型，0：单位，1：个人
	Title          string     `json:"title"`                         // 发票抬头
	TaxNumber      string     `json:"tax_number,omitempty"`          // 发票税号
	CompanyAddress string     `json:"company_address,omitempty"`     // 单位地址
	Telephone      string     `json:"telephone,omitempty"`           // 手机号码
	BankName       string     `json:"bank_name,omitempty"`           // 银行名称
	BankAccount    string     `json:"bank_account,omitempty"`        // 银行账号
	Page           DetailPage `json:"invoice_detail_page,omitempty"` // 发票详情页（小程序页面）
}
