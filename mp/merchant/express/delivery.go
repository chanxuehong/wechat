// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package express

// 邮费模板
//
//  "delivery_template": {
//      "Name": "testexpress",
//      "Assumer": 0,
//      "Valuation": 0,
//      "TopFee": [
//          {
//              "Type": 10000027,
//              "Normal": {
//                  "StartStandards": 1,
//                  "StartFees": 2,
//                  "AddStandards": 3,
//                  "AddFees": 1
//              },
//              "Custom": [
//                  {
//                      "StartStandards": 1,
//                      "StartFees": 100,
//                      "AddStandards": 1,
//                      "AddFees": 3,
//                      "DestCountry": "中国",
//                      "DestProvince": "广东省",
//                      "DestCity": "广州市"
//                  }
//              ]
//          },
//          {
//              "Type": 10000028,
//              "Normal": {
//                  "StartStandards": 1,
//                  "StartFees": 3,
//                  "AddStandards": 3,
//                  "AddFees": 2
//              },
//              "Custom": [
//                  {
//                      "StartStandards": 1,
//                      "StartFees": 10,
//                      "AddStandards": 1,
//                      "AddFees": 30,
//                      "DestCountry": "中国",
//                      "DestProvince": "广东省",
//                      "DestCity": "广州市"
//                  }
//              ]
//          },
//          {
//              "Type": 10000029,
//              "Normal": {
//                  "StartStandards": 1,
//                  "StartFees": 4,
//                  "AddStandards": 3,
//                  "AddFees": 3
//              },
//              "Custom": [
//                  {
//                      "StartStandards": 1,
//                      "StartFees": 8,
//                      "AddStandards": 2,
//                      "AddFees": 11,
//                      "DestCountry": "中国",
//                      "DestProvince": "广东省",
//                      "DestCity": "广州市"
//                  }
//              ]
//          }
//      ]
//  }
type DeliveryTemplate struct {
	Id        int64    `json:"Id,omitempty"`     // 邮费模板id
	Name      string   `json:"Name"`             // 邮费模板名称
	Assumer   int      `json:"Assumer"`          // 支付方式(0-买家承担运费, 1-卖家承担运费)
	Valuation int      `json:"Valuation"`        // 计费单位(0-按件计费, 1-按重量计费, 2-按体积计费，目前只支持按件计费，默认为0)
	TopFees   []TopFee `json:"TopFee,omitempty"` // 具体运费计算
}

// 具体运费计算
type TopFee struct {
	ExpressId int64          `json:"Type"`             // 快递类型ID(参见增加商品/快递列表, 在 ../product 里有定义)
	Normal    TopFeeNormal   `json:"Normal"`           // 默认邮费计算方法
	Customs   []TopFeeCustom `json:"Custom,omitempty"` // 指定地区邮费计算方法
}

// 默认邮费计算方法
type TopFeeNormal struct {
	StartStandards int `json:"StartStandards"` // 起始计费数量(比如计费单位是按件, 填2代表起始计费为2件)
	StartFees      int `json:"StartFees"`      // 起始计费金额(单位: 分）
	AddStandards   int `json:"AddStandards"`   // 递增计费数量
	AddFees        int `json:"AddFees"`        // 递增计费金额(单位 : 分)
}

// 指定地区邮费计算方法
type TopFeeCustom struct {
	TopFeeNormal
	DestCountry  string `json:"DestCountry"`  // 指定国家(详见《地区列表》说明)
	DestProvince string `json:"DestProvince"` // 指定省份(详见《地区列表》说明)
	DestCity     string `json:"DestCity"`     // 指定城市(详见《地区列表》说明)
}
