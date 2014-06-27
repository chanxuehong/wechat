package product

// 商品的其他属性
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
