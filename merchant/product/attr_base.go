// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package product

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

// 商品详情的一个单元, 多个这样的 Detail 组成商品的详情.
// 同一时刻只能设置一个值, 如果两个都设置则 json.Marshal 的时候只有 Text 有效
type Detail struct {
	Text  string `json:"text,omitempty"` // 文字描述
	Image string `json:"img,omitempty"`  // 图片(图片需调用图片上传接口获得图片URL填写至此，否则无法添加商品)
}

// 实现 json.Marshaler.
// text 和 image 同一时刻只 marshal 一个, 优先 marshal text.
func (detail Detail) MarshalJSON() (result []byte, err error) {
	if len(detail.Text) > 0 {
		result = make([]byte, 0, 11+len(detail.Text))
		result = append(result, `{"text":"`...)
		result = append(result, detail.Text...)
		result = append(result, `"}`...)

		return
	}

	if len(detail.Image) > 0 {
		result = make([]byte, 0, 10+len(detail.Image))
		result = append(result, `{"img":"`...)
		result = append(result, detail.Image...)
		result = append(result, `"}`...)

		return
	}

	result = []byte(`{"text":""}`)
	return
}
