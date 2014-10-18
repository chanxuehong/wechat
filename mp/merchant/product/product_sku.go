// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package product

// 商品的 SKU 信息
//  NOTE: ProductSKU.Id 的组合个数必须和 Product.AttrBase.SKUInfo 的长度一致
//
//  "sku_list": [
//      {
//          "sku_id": "1075741873:1079742386",
//          "price": 30,
//          "icon_url": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl28bJj62XgfHPibY3ORKicN1oJ4CcoIr4BMbfA8LqyyjzOZzqrOGz3f5KWq1QGP3fo6TOTSYD3TBQjuw/0",
//          "product_code": "testing",
//          "ori_price": 9000000,
//          "quantity": 800
//      },
//      {
//          "sku_id": "1075741873:1079742363",
//          "price": 30,
//          "icon_url": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl28bJj62XgfHPibY3ORKicN1oJ4CcoIr4BMbfA8LqyyjzOZzqrOGz3f5KWq1QGP3fo6TOTSYD3TBQjuw/0",
//          "product_code": "testingtesting",
//          "ori_price": 9000000,
//          "quantity": 800
//      }
//  ],
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
