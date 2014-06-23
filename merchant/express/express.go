package express

// 邮费模板
type DeliveryTemplate struct {
	Id        int64     `json:"Id,omitempty"` // 邮费模板id
	Name      string    `json:"Name"`         // 邮费模板名称
	Assumer   int       `json:"Assumer"`      // 支付方式(0-买家承担运费, 1-卖家承担运费)
	Valuation int       `json:"Valuation"`    // 计费单位(0-按件计费, 1-按重量计费, 2-按体积计费，目前只支持按件计费，默认为0)
	TopFees   []*TopFee `json:"TopFee"`       // 具体运费计算
}

// 具体运费计算
type TopFee struct {
	ExpressId int             `json:"Type"`             // 快递类型ID(参见增加商品/快递列表)
	Normal    TopFeeNormal    `json:"Normal"`           // 默认邮费计算方法
	Customs   []*TopFeeCustom `json:"Custom,omitempty"` // 指定地区邮费计算方法
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

type DeliveryInfo struct {
	DeliveryType int       `json:"delivery_type"`         // 运费类型(0-使用下面express字段的默认模板, 1-使用template_id代表的邮费模板, 详见邮费模板相关API)
	TemplateId   int64     `json:"template_id,omitempty"` // 邮费模板ID
	Expresses    []Express `json:"express,omitempty"`
}

func NewDeliveryInfoFromTemplate(templateId int64) *DeliveryInfo {
	return &DeliveryInfo{
		DeliveryType: 1,
		TemplateId:   templateId,
	}
}
func NewDeliveryInfoFromExpresses(expresses []Express) *DeliveryInfo {
	return &DeliveryInfo{
		DeliveryType: 0,
		Expresses:    expresses,
	}
}

type Express struct {
	Id    int    `json:"id"`             // 快递id
	Name  string `json:"name,omitempty"` // 快递name
	Price int    `json:"price"`          // 运费(单位 : 分)
}
