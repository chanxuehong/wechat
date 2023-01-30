package union

// Sku 商品SKU
type Sku struct {
	// SkuID 商品SKU ID
	SkuID string `json:"skuId,omitempty"`
	// ProductSkuInfo 商品SKU 小图
	ProductSkuInfo *ProductSkuInfo `json:"productSkuInfo,omitempty"`
}

// ProductSkuInfo 商品SKU 小图
type ProductSkuInfo struct {
	// ThumbImg 商品SKU 小图
	ThumbImg string `json:"thumbImg,omitempty"`
	// SalePrice 商品SKU 销售价格，单位分
	SalePrice int64 `json:"salePrice,omitempty"`
	// MarketPrice 商品SKU 市场价格，单位分
	MarketPrice int64 `json:"marketPrice,omitempty"`
	// StockInfo 商品SKU 库存
	StockInfo *StockInfo `json:"stockInfo,omitempty"`
}

// StockInfo 商品SKU
type StockInfo struct {
	// StockNum 库存
	StockNum int64 `json:"stockNum,omitempty"`
}
