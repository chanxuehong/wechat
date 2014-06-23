package product

// 获取指定子分类的所有SKU 成功时返回结果的数据结构
type SKU struct {
	CategoryId int64  `json:"-"`
	Id         string `json:"id"`
	Name       string `json:"name"`
	Values     []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"value_list"`
}

type SKUInfo struct {
	// sku信息, 参照上述sku_table的定义;
	// 格式 : "id1:vid1;id2:vid2"
	// 规则 : id_info的组合个数必须与sku_table个数一致(若商品无sku信息, 即商品为统一规格，
	// 则此处赋值为空字符串即可)
	SKUId         string `json:"sku_id"`
	OriginalPrice int    `json:"ori_price"`    // sku原价(单位 : 分)
	Price         int    `json:"price"`        // sku微信价(单位 : 分, 微信价必须比原价小, 否则添加商品失败)
	IconURL       string `json:"icon_url"`     // sku iconurl(图片需调用图片上传接口获得图片URL)
	ProductCode   string `json:"product_code"` // 商家商品编码
	Quantity      int    `json:"quantity"`     // sku库存
}
