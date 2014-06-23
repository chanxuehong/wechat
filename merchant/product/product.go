package product

import (
	"github.com/chanxuehong/wechat/merchant/express"
)

/*
{
    "product_base": {
        "category_id": [
            "537074298"
        ],
        "property": [
            {
                "id": "1075741879",
                "vid": "1079749967"
            },
            {
                "id": "1075754127",
                "vid": "1079795198"
            },
            {
                "id": "1075777334",
                "vid": "1079837440"
            }
        ],
        "name": "testaddproduct",
        "sku_info": [
            {
                "id": "1075741873",
                "vid": [
                    "1079742386",
                    "1079742363"
                ]
            }
        ],
        "main_img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ulEKogfsiaua49pvLfUS8Ym0GSYjViaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
        "img": [
            "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ulEKogfsiaua49pvLfUS8Ym0GSYjViaLic0FD3vN0V8PILcibEGb2fPfEOmw/0"
        ],
        "detail": [
            {
                "text": "test first"
            },
            {
                "img": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjD3ul1UcLcwxrFdwTKYhH9Q5YZoCfX4Ncx655ZK6ibnlibCCErbKQtReySaVA/0"
            },
            {
                "text": "test again"
            }
        ],
        "buy_limit": 10
    },
    "sku_list": [
        {
            "sku_id": "1075741873:1079742386",
            "price": 30,
            "icon_url": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl28bJj62XgfHPibY3ORKicN1oJ4CcoIr4BMbfA8LqyyjzOZzqrOGz3f5KWq1QGP3fo6TOTSYD3TBQjuw/0",
            "product_code": "testing",
            "ori_price": 9000000,
            "quantity": 800
        },
        {
            "sku_id": "1075741873:1079742363",
            "price": 30,
            "icon_url": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl28bJj62XgfHPibY3ORKicN1oJ4CcoIr4BMbfA8LqyyjzOZzqrOGz3f5KWq1QGP3fo6TOTSYD3TBQjuw/0",
            "product_code": "testingtesting",
            "ori_price": 9000000,
            "quantity": 800
        }
    ],
    "attrext": {
        "location": {
            "country": "中国",
            "province": "广东省",
            "city": "广州市",
            "address": "T.I.T创意园"
        },
        "isPostFree": 0,
        "isHasReceipt": 1,
        "isUnderGuaranty": 0,
        "isSupportReplace": 0
    },
    "delivery_info": {
        "delivery_type": 0,
        "template_id": 0,
        "express": [
            {
                "id": 10000027,
                "price": 100
            },
            {
                "id": 10000028,
                "price": 100
            },
            {
                "id": 10000029,
                "price": 100
            }
        ]
    }
}
*/
type Product struct {
	Id   string `json:"product_id,omitempty"` // 商品id
	Attr struct {
		Name        string         `json:"name"`                // 商品名称
		CategoryIds []string       `json:"category_id"`         // 商品分类id，商品分类列表请通过《获取指定分类的所有子分类》获取
		MainImage   string         `json:"main_img"`            // 商品主图(图片需调用图片上传接口获得图片Url填写至此，否则无法添加商品。图片分辨率推荐尺寸为640×600)
		Images      []string       `json:"img"`                 // 商品图片列表(图片需调用图片上传接口获得图片Url填写至此，否则无法添加商品。图片分辨率推荐尺寸为640×600)
		Details     []Detail       `json:"detail"`              // 商品详情列表，显示在客户端的商品详情页内
		Properties  []AttrProperty `json:"property,omitempty"`  // 商品属性列表，属性列表请通过《获取指定分类的所有属性》获取
		SKUs        []AttrSKU      `json:"sku_info,omitempty"`  // 商品sku定义，SKU列表请通过《获取指定子分类的所有SKU》获取
		BuyLimit    int            `json:"buy_limit,omitempty"` // 用户商品限购数量
	} `json:"product_base"` // 基本属性

	AttrExt      *AttrExt              `json:"attrext,omitempty"`       // 商品其他属性
	SKUInfos     []SKUInfo             `json:"sku_list,omitempty"`      // sku信息列表(可为多个)，每个sku信息串即为一个确定的商品，比如白色的37码的鞋子
	DeliveryInfo *express.DeliveryInfo `json:"delivery_info,omitempty"` // 运费信息
}

type AttrProperty struct {
	Id      string `json:"id"`  // 属性id
	ValueId string `json:"vid"` // 属性值id
}
type AttrSKU struct {
	Id       string   `json:"id"`  // sku属性(SKU列表中id, 支持自定义SKU，格式为"$xxx"，xxx即为显示在客户端中的字符串)
	ValueIds []string `json:"vid"` // sku值(SKU列表中vid, 如需自定义SKU，格式为"$xxx"，xxx即为显示在客户端中的字符串)
}

// 同一时刻只能设置一个值, 如果两个都设置则 json.Marshal 的时候只有 Text 有效
type Detail struct {
	Text  string `json:"text,omitempty"` // 文字描述
	Image string `json:"img,omitempty"`  // 图片(图片需调用图片上传接口获得图片Url填写至此，否则无法添加商品)
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
	Location struct {
		Country  string `json:"country"`  // 国家(详见《地区列表》说明)
		Province string `json:"province"` // 省份(详见《地区列表》说明)
		City     string `json:"city"`     // 城市(详见《地区列表》说明)
		Address  string `json:"address"`  // 地址
	} `json:"location"` // 商品所在地地址

	IsPostFree       int `json:"isPostFree"`       // 是否包邮(0-否, 1-是), 如果包邮delivery_info字段可省略
	IsHasReceipt     int `json:"isHasReceipt"`     // 是否提供发票(0-否, 1-是)
	IsUnderGuaranty  int `json:"isUnderGuaranty"`  // 是否保修(0-否, 1-是)
	IsSupportReplace int `json:"isSupportReplace"` // 是否支持退换货(0-否, 1-是)
}
