// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package product

// 商品的其他属性
//
//  "attrext": {
//      "location": {
//          "country": "中国",
//          "province": "广东省",
//          "city": "广州市",
//          "address": "T.I.T创意园"
//      },
//      "isPostFree": 0,
//      "isHasReceipt": 1,
//      "isUnderGuaranty": 0,
//      "isSupportReplace": 0
//  },
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

// 设置是否 包邮, true--包邮, false--不包邮
func (attr *AttrExt) SetPostFree(b bool) {
	if b {
		attr.IsPostFree = 1
	} else {
		attr.IsPostFree = 0
	}
}

// 设置是否 提供发票, true--提供发票, false--不提供发票
func (attr *AttrExt) SetHasReceipt(b bool) {
	if b {
		attr.IsHasReceipt = 1
	} else {
		attr.IsHasReceipt = 0
	}
}

// 设置是否保修, true--保修, false--不保修
func (attr *AttrExt) SetUnderGuaranty(b bool) {
	if b {
		attr.IsUnderGuaranty = 1
	} else {
		attr.IsUnderGuaranty = 0
	}
}

// 设置是否支持退货, true--支持退货, false--不支持退货
func (attr *AttrExt) SetSupportReplace(b bool) {
	if b {
		attr.IsSupportReplace = 1
	} else {
		attr.IsSupportReplace = 0
	}
}
