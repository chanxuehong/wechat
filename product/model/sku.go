package model

type Sku struct {
	Id           uint64      `json:"sku_id,omitempty"`         // 小商店内部SKU ID
	OutProductId string      `json:"out_product_id,omitempty"` // 商家自定义商品ID
	OutId        string      `json:"out_sku_id,omitempty"`     // sku_id
	ThumbImg     string      `json:"thumb_img,omitempty"`      // sku小图
	SalePrice    uint        `json:"sale_price,omitempty"`     // 售卖价格,以分为单位
	MarketPrice  uint        `json:"market_price,omitempty"`   // 市场价格,以分为单位
	StockNum     int         `json:"stock_num,omitempty"`      // 库存
	Barcode      string      `json:"barcode,omitempty"`        // 条形码
	SkuCode      string      `json:"sku_code,omitempty"`       // 商品编码
	Attrs        []Attribute `json:"sku_attrs,omitempty"`      // 属性自定义用
	UpdateType   int         `json:"type,omitempty"`           // 1:全量更新 2:增量更新
	CreateTime   string      `json:"create_time,omitempty"`
	UpdateTime   string      `json:"update_time,omitempty"`
}
