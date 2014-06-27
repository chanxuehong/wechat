package product

// 关于 Product.AttrBase.SKUInfo 和 Product.ProductSKUs 的关系说明:
// AttrBase.SKUInfo 是指定 Product 的 sku 是由哪些属性组合确定的, 这些属性就是
// SKUInfo([]SKU) 的每个 SKU 的 Id(属性id);
// ProductSKUs 就是具体这些属性的排列组合, SKUInfo 的每个属性取一个属性值做组合, 如:
// "id1:vid11;id2:vid21;id3:vid32"
// "id1:vid12;id2:vid22;id3:vid31"
// 这里的 id1,id2,id3 就是 SKUInfo 里的 SKU 的 Id;
// vid11,vid12,vid21,vid22,vid31,vid32 就是 SKUInfo 里的 SKU 的 vid;
//
// 规则是 ProductSKUs 的 skuid 的组合个数必须是和 SKUInfo 的 SKU 的个数一致!!!

type Product struct {
	Id     string `json:"product_id,omitempty"` // 商品id
	Status int    `json:"status,omitempty"`     // 商品状态

	AttrBase     AttrBase     `json:"product_base"`       // 基本属性
	AttrExt      *AttrExt     `json:"attrext,omitempty"`  // 商品其他属性
	ProductSKUs  []ProductSKU `json:"sku_list,omitempty"` // sku信息列表(可为多个)，每个sku信息串即为一个确定的商品，比如白色的37码的鞋子
	DeliveryInfo DeliveryInfo `json:"delivery_info"`      // 运费信息
}

// eg:
//	_product.SetDeliveryInfoWithExpresses([]Express{
//		{
//			Id:    10000027,
//			Price: 100,
//		},
//		{
//			Id:    10000028,
//			Price: 100,
//		},
//		{
//			Id:    10000029,
//			Price: 100,
//		},
//	})
func (p *Product) SetDeliveryInfoWithExpresses(expresses []Express) {
	p.DeliveryInfo.DeliveryType = 0
	p.DeliveryInfo.Expresses = expresses
	p.DeliveryInfo.TemplateId = 0
}
func (p *Product) SetDeliveryInfoWithTemplate(templateId int64) {
	p.DeliveryInfo.DeliveryType = 1
	p.DeliveryInfo.Expresses = nil
	p.DeliveryInfo.TemplateId = templateId
}
