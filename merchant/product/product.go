package product

type Product struct {
	Id     string `json:"product_id,omitempty"` // 商品id
	Status int    `json:"status,omitempty"`     // 商品状态

	AttrBase     AttrBase      `json:"product_base"`            // 基本属性
	AttrExt      *AttrExt      `json:"attrext,omitempty"`       // 商品其他属性
	ProductSKUs  []ProductSKU  `json:"sku_list,omitempty"`      // sku信息列表(可为多个)，每个sku信息串即为一个确定的商品，比如白色的37码的鞋子
	DeliveryInfo *DeliveryInfo `json:"delivery_info,omitempty"` // 运费信息
}

type AttrBase struct {
	Name        string     `json:"name"`                  // 商品名称
	CategoryIds []string   `json:"category_id,omitempty"` // 商品分类id，商品分类列表请通过《获取指定分类的所有子分类》获取
	MainImage   string     `json:"main_img"`              // 商品主图(图片需调用图片上传接口获得图片URL填写至此，否则无法添加商品。图片分辨率推荐尺寸为640×600)
	Images      []string   `json:"img,omitempty"`         // 商品图片列表(图片需调用图片上传接口获得图片URL填写至此，否则无法添加商品。图片分辨率推荐尺寸为640×600)
	Details     []Detail   `json:"detail,omitempty"`      // 商品详情列表，显示在客户端的商品详情页内
	Properties  []Property `json:"property,omitempty"`    // 商品属性列表，属性列表请通过《获取指定分类的所有属性》获取
	SKUInfo     []SKU      `json:"sku_info,omitempty"`    // 商品sku定义，SKU列表请通过《获取指定子分类的所有SKU》获取
	BuyLimit    int        `json:"buy_limit,omitempty"`   // 用户商品限购数量
}

// 参考 category/Property
type Property struct {
	Id      string `json:"id"`  // 属性id
	ValueId string `json:"vid"` // 属性值id
}

// 参考 category/SKU
type SKU struct {
	Id       string   `json:"id"`            // sku属性(SKU列表中id, 支持自定义SKU，格式为"$xxx"，xxx即为显示在客户端中的字符串)
	ValueIds []string `json:"vid,omitempty"` // sku值(SKU列表中vid, 如需自定义SKU，格式为"$xxx"，xxx即为显示在客户端中的字符串)
}

// ProductSKU.Id 的组合个数和 Product.Attr.SKUs 的个数一致
type ProductSKU struct {
	// sku信息, 参照上述sku_table的定义;
	// 格式 : "id1:vid1;id2:vid2"
	// 规则 : id_info的组合个数必须与sku_table个数一致(若商品无sku信息, 即商品为统一规格，
	// 则此处赋值为空字符串即可)
	Id            string `json:"sku_id"`
	Price         int    `json:"price"`        // sku微信价(单位 : 分, 微信价必须比原价小, 否则添加商品失败)
	IconURL       string `json:"icon_url"`     // sku iconurl(图片需调用图片上传接口获得图片URL)
	ProductCode   string `json:"product_code"` // 商家商品编码
	OriginalPrice int    `json:"ori_price"`    // sku原价(单位 : 分)
	Quantity      int    `json:"quantity"`     // sku库存
}

// 商品详情的一个单元, 多个这样的 Detail 组成商品的详情.
// 同一时刻只能设置一个值, 如果两个都设置则 json.Marshal 的时候只有 Text 有效
type Detail struct {
	Text  string `json:"text,omitempty"` // 文字描述
	Image string `json:"img,omitempty"`  // 图片(图片需调用图片上传接口获得图片URL填写至此，否则无法添加商品)
}

// 实现 json.Marshaler.
// text 和 image 同一时刻只 marshal 一个, 优先 marshal text.
func (detail Detail) MarshalJSON() ([]byte, error) {
	if len(detail.Text) > 0 {
		ret := make([]byte, 0, 11+len(detail.Text))
		ret = append(ret, `{"text":"`...)
		ret = append(ret, detail.Text...)
		ret = append(ret, `"}`...)

		return ret, nil
	}

	if len(detail.Image) > 0 {
		ret := make([]byte, 0, 10+len(detail.Image))
		ret = append(ret, `{"img":"`...)
		ret = append(ret, detail.Image...)
		ret = append(ret, `"}`...)

		return ret, nil
	}

	return []byte(`{"text":""}`), nil
}

// 商品的其他属性
type AttrExt struct {
	Location Location `json:"location"` // 商品所在地地址

	IsPostFree       int `json:"isPostFree"`       // 是否包邮(0-否, 1-是), 如果包邮delivery_info字段可省略
	IsHasReceipt     int `json:"isHasReceipt"`     // 是否提供发票(0-否, 1-是)
	IsUnderGuaranty  int `json:"isUnderGuaranty"`  // 是否保修(0-否, 1-是)
	IsSupportReplace int `json:"isSupportReplace"` // 是否支持退换货(0-否, 1-是)
}

type Location struct {
	Country  string `json:"country"`  // 国家(详见《地区列表》说明)
	Province string `json:"province"` // 省份(详见《地区列表》说明)
	City     string `json:"city"`     // 城市(详见《地区列表》说明)
	Address  string `json:"address"`  // 地址
}
