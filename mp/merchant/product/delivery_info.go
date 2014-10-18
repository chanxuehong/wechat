// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package product

// 默认的邮费模板
type Express struct {
	Id    int64  `json:"id"`             // 快递id, 平邮: 10000027, 快递: 10000028, EMS: 10000029
	Name  string `json:"name,omitempty"` // 快递name
	Price int    `json:"price"`          // 运费(单位 : 分)
}

// 运费信息
//
//  "delivery_info": {
//      "delivery_type": 0,
//      "template_id": 0,
//      "express": [
//          {
//              "id": 10000027,
//              "price": 100
//          },
//          {
//              "id": 10000028,
//              "price": 100
//          },
//          {
//              "id": 10000029,
//              "price": 100
//          }
//      ]
//  }
type DeliveryInfo struct {
	// 运费类型(0-使用下面express字段的默认模板, 1-使用template_id代表的邮费模板, 详见邮费模板相关API)
	DeliveryType int       `json:"delivery_type"`
	TemplateId   int64     `json:"template_id,omitempty"` // 邮费模板ID
	Expresses    []Express `json:"express,omitempty"`
}

func (info *DeliveryInfo) SetWithExpresses(expresses []Express) {
	info.DeliveryType = 0
	info.Expresses = expresses
	info.TemplateId = 0
}
func (info *DeliveryInfo) SetWithTemplate(templateId int64) {
	info.DeliveryType = 1
	info.Expresses = nil
	info.TemplateId = templateId
}

func NewDeliveryInfoWithExpresses(expresses []Express) *DeliveryInfo {
	return &DeliveryInfo{
		DeliveryType: 0,
		Expresses:    expresses,
	}
}

func NewDeliveryInfoWithTemplate(templateId int64) *DeliveryInfo {
	return &DeliveryInfo{
		DeliveryType: 1,
		TemplateId:   templateId,
	}
}
