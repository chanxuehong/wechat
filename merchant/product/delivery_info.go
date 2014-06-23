package product

// 默认的邮费模板
type Express struct {
	Id    int64  `json:"id"`             // 快递id, 平邮: 10000027, 快递: 10000028, EMS: 10000029
	Name  string `json:"name,omitempty"` // 快递name
	Price int    `json:"price"`          // 运费(单位 : 分)
}

// 运费信息
type DeliveryInfo struct {
	// 运费类型(0-使用下面express字段的默认模板, 1-使用template_id代表的邮费模板, 详见邮费模板相关API)
	DeliveryType int       `json:"delivery_type"`
	Expresses    []Express `json:"express,omitempty"`
	TemplateId   int64     `json:"template_id,omitempty"` // 邮费模板ID
}

func NewDeliveryInfoFromExpresses(expresses []Express) *DeliveryInfo {
	return &DeliveryInfo{
		DeliveryType: 0,
		Expresses:    expresses,
	}
}
func NewDeliveryInfoFromTemplate(templateId int64) *DeliveryInfo {
	return &DeliveryInfo{
		DeliveryType: 1,
		TemplateId:   templateId,
	}
}
