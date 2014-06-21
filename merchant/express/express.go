package express

type DeliveryTemplate struct {
	Id        int       `json:"Id,omitempty"` // 邮费模板id
	Name      string    `json:"Name"`         // 邮费模板名称
	Assumer   int       `json:"Assumer"`      // 支付方式(0-买家承担运费, 1-卖家承担运费)
	Valuation int       `json:"Valuation"`    // 计费单位(0-按件计费, 1-按重量计费, 2-按体积计费，目前只支持按件计费，默认为0)
	TopFee    []*TopFee `json:"TopFee"`       // 具体运费计算
}

// 具体运费计算
type TopFee struct {
	Type   int             `json:"Type"`   // 快递类型ID(参见增加商品/快递列表)
	Normal TopFeeNormal    `json:"Normal"` // 默认邮费计算方法
	Custom []*TopFeeCustom `json:"Custom"` // 指定地区邮费计算方法
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
