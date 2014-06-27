// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package product

// ProductSKU.Id 的组合个数必须和 Product.AttrBase.SKUInfo 的长度一致
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
