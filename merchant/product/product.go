// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong@gmail.com

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

	AttrBase     AttrBase      `json:"product_base"`            // 基本属性
	ProductSKUs  []ProductSKU  `json:"sku_list,omitempty"`      // sku信息列表(可为多个)，每个sku信息串即为一个确定的商品，比如白色的37码的鞋子
	AttrExt      *AttrExt      `json:"attrext,omitempty"`       // 商品其他属性
	DeliveryInfo *DeliveryInfo `json:"delivery_info,omitempty"` // 运费信息, 如果在 AttrExt 设置包邮, 则这个字段可以省略
}

//==============================================================================

// 是否包邮, true--包邮, false--不包邮
func (p *Product) IsPostFree() bool {
	if p.AttrExt == nil || p.AttrExt.IsPostFree == 0 {
		return false
	}
	return true
}

// 是否提供发票, true--提供发票, false--不提供发票
func (p *Product) IsHasReceipt() bool {
	if p.AttrExt == nil || p.AttrExt.IsHasReceipt == 0 {
		return false
	}
	return true
}

// 是否保修, true--保修, false--不保修
func (p *Product) IsUnderGuaranty() bool {
	if p.AttrExt == nil || p.AttrExt.IsUnderGuaranty == 0 {
		return false
	}
	return true
}

// 是否支持退货, true--支持退货, false--不支持退货
func (p *Product) IsSupportReplace() bool {
	if p.AttrExt == nil || p.AttrExt.IsSupportReplace == 0 {
		return false
	}
	return true
}

//==============================================================================

// 设置是否 包邮, true--包邮, false--不包邮
func (p *Product) SetPostFree(b bool) {
	if p.AttrExt == nil {
		p.AttrExt = new(AttrExt)
	}
	p.AttrExt.SetPostFree(b)
}

// 设置是否 提供发票, true--提供发票, false--不提供发票
func (p *Product) SetHasReceipt(b bool) {
	if p.AttrExt == nil {
		p.AttrExt = new(AttrExt)
	}
	p.AttrExt.SetHasReceipt(b)
}

// 设置是否保修, true--保修, false--不保修
func (p *Product) SetUnderGuaranty(b bool) {
	if p.AttrExt == nil {
		p.AttrExt = new(AttrExt)
	}
	p.AttrExt.SetUnderGuaranty(b)
}

// 设置是否支持退货, true--支持退货, false--不支持退货
func (p *Product) SetSupportReplace(b bool) {
	if p.AttrExt == nil {
		p.AttrExt = new(AttrExt)
	}
	p.AttrExt.SetSupportReplace(b)
}

//==============================================================================

// 使用默认邮费模板
func (p *Product) SetDeliveryInfoWithExpresses(expresses []Express) {
	if p.DeliveryInfo == nil {
		p.DeliveryInfo = new(DeliveryInfo)
	}
	p.DeliveryInfo.SetWithExpresses(expresses)
}

// 使用自定义的邮费模板
func (p *Product) SetDeliveryInfoWithTemplate(templateId int64) {
	if p.DeliveryInfo == nil {
		p.DeliveryInfo = new(DeliveryInfo)
	}
	p.DeliveryInfo.SetWithTemplate(templateId)
}
