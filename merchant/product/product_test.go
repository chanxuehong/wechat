// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package product

import (
	"bytes"
	"encoding/json"
	"github.com/chanxuehong/util"
	"testing"
)

func TestProductMarshal(t *testing.T) {
	expectBytes := []byte(`{
		"product_base":{
			"name":"testaddproduct",
			"category_id":[
				"537074298"
			],
			"main_img":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ulEKogfsiaua49pvLfUS8Ym0GSYjViaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			"img":[
				"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ulEKogfsiaua49pvLfUS8Ym0GSYjViaLic0FD3vN0V8PILcibEGb2fPfEOmw/0"
			],
			"detail":[
				{
					"text":"test first"
				},
				{
					"img":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ul1UcLcwxrFdwTKYhH9Q5YZoCfX4Ncx655ZK6ibnlibCCErbKQtReySaVA/0"
				},
				{
					"text":"test again"
				}
			],
			"property":[
				{
					"id":"1075741879",
					"vid":"1079749967"
				},
				{
					"id":"1075754127",
					"vid":"1079795198"
				},
				{
					"id":"1075777334",
					"vid":"1079837440"
				}
			],
			"sku_info":[
				{
					"id":"1075741873",
					"vid":[
						"1079742386",
						"1079742363"
					]
				}
			],
			"buy_limit":10
		},
		"sku_list":[
			{
				"sku_id":"1075741873:1079742386",
				"price":30,
				"icon_url":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl28bJj62XgfHPibY3ORKicN1oJ4CcoIr4BMbfA8LqyyjzOZzqrOGz3f5KWq1QGP3fo6TOTSYD3TBQjuw/0",
				"product_code":"testing",
				"ori_price":9000000,
				"quantity":800
			},
			{
				"sku_id":"1075741873:1079742363",
				"price":30,
				"icon_url":"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl28bJj62XgfHPibY3ORKicN1oJ4CcoIr4BMbfA8LqyyjzOZzqrOGz3f5KWq1QGP3fo6TOTSYD3TBQjuw/0",
				"product_code":"testingtesting",
				"ori_price":9000000,
				"quantity":800
			}
		],
		"attrext":{
			"location":{
				"country":"中国",
				"province":"广东省",
				"city":"广州市",
				"address":"T.I.T创意园"
			},
			"isPostFree":0,
			"isHasReceipt":1,
			"isUnderGuaranty":0,
			"isSupportReplace":0
		},
		"delivery_info":{
			"delivery_type":0,
			"express":[
				{
					"id":10000027,
					"price":100
				},
				{
					"id":10000028,
					"price":100
				},
				{
					"id":10000029,
					"price":100
				}
			]
		}
	}`)

	_product := Product{
		AttrBase: AttrBase{
			Name:        "testaddproduct",
			CategoryIds: []string{"537074298"},
			MainImage:   "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ulEKogfsiaua49pvLfUS8Ym0GSYjViaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			Images:      []string{"http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ulEKogfsiaua49pvLfUS8Ym0GSYjViaLic0FD3vN0V8PILcibEGb2fPfEOmw/0"},
			Details:     []Detail{{Text: "test first"}, {Image: "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ul1UcLcwxrFdwTKYhH9Q5YZoCfX4Ncx655ZK6ibnlibCCErbKQtReySaVA/0"}, {Text: "test again"}},
			Properties:  []Property{{"1075741879", "1079749967"}, {"1075754127", "1079795198"}, {"1075777334", "1079837440"}},
			SKUInfo:     []SKU{{Id: "1075741873", ValueIds: []string{"1079742386", "1079742363"}}},
			BuyLimit:    10,
		},
	}
	_product.AttrExt = &AttrExt{
		Location: Location{
			Country:  "中国",
			Province: "广东省",
			City:     "广州市",
			Address:  "T.I.T创意园",
		},
		IsPostFree:       0,
		IsHasReceipt:     1,
		IsUnderGuaranty:  0,
		IsSupportReplace: 0,
	}
	_product.ProductSKUs = []ProductSKU{
		{
			Id:            "1075741873:1079742386",
			Price:         30,
			IconURL:       "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl28bJj62XgfHPibY3ORKicN1oJ4CcoIr4BMbfA8LqyyjzOZzqrOGz3f5KWq1QGP3fo6TOTSYD3TBQjuw/0",
			ProductCode:   "testing",
			OriginalPrice: 9000000,
			Quantity:      800,
		},
		{
			Id:            "1075741873:1079742363",
			Price:         30,
			IconURL:       "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl28bJj62XgfHPibY3ORKicN1oJ4CcoIr4BMbfA8LqyyjzOZzqrOGz3f5KWq1QGP3fo6TOTSYD3TBQjuw/0",
			ProductCode:   "testingtesting",
			OriginalPrice: 9000000,
			Quantity:      800,
		},
	}
	_product.DeliveryInfo = NewDeliveryInfoWithExpresses([]Express{
		{
			Id:    10000027,
			Price: 100,
		},
		{
			Id:    10000028,
			Price: 100,
		},
		{
			Id:    10000029,
			Price: 100,
		},
	})

	b, err := json.Marshal(_product)
	if err != nil {
		t.Errorf("json.Marshal(%#q):\nError: %s\n", _product, err)
	} else {
		want := util.TrimSpace(expectBytes)
		if !bytes.Equal(b, want) {
			t.Errorf("json.Marshal(%#q):\nhave %#s\nwant %#s\n", _product, b, want)
		}
	}
}
