package model

type Spu struct {
	Id         uint64   `json:"product_id,omitempty"`     // 小商店内部商品ID
	OutId      string   `json:"out_product_id,omitempty"` // 商家自定义商品ID，与product_id二选一
	Title      string   `json:"title,omitempty"`          // 标题
	SubTitle   string   `json:"sub_title,omitempty"`      // 副标题
	HeadImages []string `json:"head_img,omitempty"`       // 主图,多张,列表
	DescInfo   *struct {
		Imgs []string `json:"imgs,omitempty"`
	} `json:"desc_info,omitempty"` // 商品详情，图文(目前只支持图片)
	BrandId     uint64     `json:"brand_id,omitempty"` // 商家需要申请品牌
	Cats        []Category `json:"cats,omitempty"`     // 商家需要先申请可使用类目
	Attrs       []Attribute
	Model       string           `json:"model,omitempty"`        // 商品型号
	ExpressInfo *FreightTemplate `json:"express_info,omitempty"` // 运费模板ID（先通过获取运费模板接口拿到）
	Skus        []Sku            `json:"skus,omitempty"`         // 该 skus 列表非必填，可另行通过 BatchAddSKU 添加 SKU
	CreateTime  string           `json:"create_time,omitmepty"`  // 创建时间
	UpdateTime  string           `json:"update_time,omitempty"`  // 更新时间
}
