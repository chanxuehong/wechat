package express

import (
	"github.com/bububa/wechat/mp/core"
)

type OrderAddRequest struct {
	// OrderID 订单ID，须保证全局唯一，不超过512字节
	OrderID string `json:"order_id,omitempty"`
	// OpenID 用户openid，当add_source=2时无需填写（不发送物流服务通知）
	OpenID string `json:"openid,omitempty"`
	// DeliveryID 快递公司ID，参见getAllDelivery
	DeliveryID string `json:"delivery_id,omitempty"`
	// BizID 快递客户编码或者现付编码
	BizID string `json:"biz_id,omitempty"`
	// CustomRemask 快递备注信息，比如"易碎物品"，不超过1024字节
	CustomRemark string `json:"custom_remark,omitempty"`
	// TagID 订单标签id，用于平台型小程序区分平台上的入驻方，tagid须与入驻方账号一一对应，非平台型小程序无需填写该字段
	TagID uint64 `json:"tag_id,omitempty"`
	// AddSource 订单来源，0为小程序订单，2为 App 或H5订单，填2则不发送物流服务通知
	AddSource int `json:"add_source"`
	// WxAppID App或H5的appid，add_source=2时必填，需和开通了物流助手的小程序绑定同一 open 帐号
	WxAppID string `json:"wx_appid,omitempty"`
	// Sender 发件人信息
	Sender *ContactInfo `json:"sender,omitempty"`
	// Receiver 收件人信息
	Receiver *ContactInfo `json:"receiver,omitempty"`
	// Cargo 包裹信息，将传递给快递公司
	Cargo *Cargo `json:"cargo,omitempty"`
	// Shop 商品信息，会展示到物流服务通知和电子面单中
	Shop *Shop `json:"shop,omitempty"`
	// Insured 保价信息
	Insured *InsureInfo `json:"insured,omitempty"`
	// Service 服务类型
	Service *ServiceType `json:"service,omitempty"`
	// ExpectTime Unix 时间戳, 单位秒，顺丰必须传。 预期的上门揽件时间，0表示已事先约定取件时间；否则请传预期揽件时间戳，需大于当前时间，收件员会在预期时间附近上门。例如expect_time为“1557989929”，表示希望收件员将在2019年05月16日14:58:49-15:58:49内上门取货。说明：若选择 了预期揽件时间，请不要自己打单，由上门揽件的时候打印。如果是下顺丰散单，则必传此字段，否则不会有收件员上门揽件。
	ExpectTime int64 `json:"expect_time,omitempty"`
	// TakeMode 分单策略，【0：线下网点签约，1：总部签约结算】，不传默认线下网点签约。目前支持圆通。
	TakdMode int `json:"take_mode,omitempty"`
}

// ContactInfo 发/收件人信息
type ContactInfo struct {
	// Name 发件人姓名，不超过64字节
	Name string `json:"name,omitempty"`
	// Tel 发件人座机号码，若不填写则必须填写 mobile，不超过32字节
	Tel string `json:"tel,omitempty"`
	// Mobile 发件人手机号码，若不填写则必须填写 tel，不超过32字节
	Mobile string `json:"mobile,omitempty"`
	// Company 发件人公司名称，不超过64字节
	Company string `json:"company,omitempty"`
	// PostCode 发件人邮编，不超过10字节
	PostCode string `json:"post_code,omitempty"`
	// Country 发件人国家，不超过64字节
	Country string `json:"country,omitempty"`
	// Province 发件人省份，比如："广东省"，不超过64字节
	Province string `json:"province,omitempty"`
	// City 发件人市/地区，比如："广州市"，不超过64字节
	City string `json:"city,omitempty"`
	// Area 发件人区/县，比如："海珠区"，不超过64字节
	Area string `json:"area,omitempty"`
	// Address 发件人详细地址，比如："XX路 XX 号XX大厦XX"，不超过512字节
	Address string `json:"address,omitempty"`
}

// Cargo 包裹信息，将传递给快递公司
type Cargo struct {
	// Count 包裹数量, 默认为1
	Count int `json:"count,omitempty"`
	// Weight 包裹总重量，单位是千克(kg)
	Weight float64 `json:"weight,omitempty"`
	// SpaceX
	SpaceX float64 `json:"space_x,omitempty"`
	// SpaceY
	SpaceY float64 `json:"space_y,omitempty"`
	// SpaceZ
	SpaceZ float64 `json:"space_z,omitempty"`
	// DetailList 包裹总重量，单位是千克(kg)
	DetailList []CargoDetail `json:"detail_list,omitempty"`
}

type CargoDetail struct {
	// Name 商品名，不超过128字节
	Name string `json:"name,omitempty"`
	// Count 商品数量
	Count int `json:"count,omitempty"`
}

// Shop 商品信息，会展示到物流服务通知和电子面单中
type Shop struct {
	// WxaPath 商家小程序的路径，建议为订单页面
	WxaPath string `json:"wxa_path,omitempty"`
	// ImgURL 商品缩略图 url；shop.detail_list为空则必传，shop.detail_list非空可不传。
	ImgURL string `json:"img_url,omitempty"`
	// GoodsName 商品名称, 不超过128字节；shop.detail_list为空则必传，shop.detail_list非空可不传。
	GoodsName string `json:"goods_name,omitempty"`
	// GoodsCount 商品数量；shop.detail_list为空则必传。shop.detail_list非空可不传，默认取shop.detail_list的size
	GoodsCount int `json:"goods_count,omitempty"`
	// DetailList 商品详情列表，适配多商品场景，用以消息落地页展示。（新规范，新接入商家建议用此字段）
	DetailList []ShopDetail `json:"detail_list,omitempty"`
}

type ShopDetail struct {
	// GoodsName 商品名称, 不超过128字节；shop.detail_list为空则必传，shop.detail_list非空可不传。
	GoodsName string `json:"goods_name,omitempty"`
	// GoodsImgURL
	GoodsImgURL string `json:"goods_img_url,omitempty"`
	// GoodsDesc
	GoodsDesc string `json:"goods_desc,omitempty"`
	// GoodsCount 商品数量；shop.detail_list为空则必传。shop.detail_list非空可不传，默认取shop.detail_list的size
	GoodsCount int `json:"goods_count,omitempty"`
}

// InsureInfo 保价信息
type InsureInfo struct {
	// UseInsured 是否保价，0 表示不保价，1 表示保价
	UseInsured int `json:"user_insured,omitempty"`
	// InsuredValue 保价金额，单位是分，比如: 10000 表示 100 元
	InsuredValue int64 `json:"insured_value,omitempty"`
}

// ServiceType 服务类型
type ServiceType struct {
	// ServiceType 服务类型ID
	ServiceType int `json:"service_type,omitempty"`
	// ServiceName 服务名称
	ServiceName string `json:"service_name,omitempty"`
}

// Order 订单信息
type Order struct {
	// OrderID 订单ID，下单成功时返回
	OrderID string `json:"order_id,omitempty"`
	// WaybillID 运单ID，下单成功时返回
	WaybillID string `json:"waybill_id,omitempty"`
	// DeliveryResultCode 快递侧错误码，下单失败时返回
	DeliveryResultCode int `json:"delivery_resultcode,omitempty"`
	// DeliveryResultMsg 快递侧错误信息，下单失败时返回
	DeliveryResultMsg string `json:"delivery_resultmsg,omitempty"`
	// WaybillData 运单信息，下单成功时返回
	WaybillData []WaybillKV `json:"waybill_data,omitempty"`
	// PrintHtml 运单 html 的 BASE64 结果
	PrintHtml string `json:"print_html,omitempty"`
	// OrderStatus 运单状态, 0正常，1取消
	OrderStatus int `json:"order_status,omitempty"`
}

// WaybillKV 运单信息
type WaybillKV struct {
	// Key 运单信息 key
	Key string `json:"key,omitempty"`
	// Value 运单信息 value
	Value string `json:"value,omitempty"`
}

// OrderAdd 生成运单
func OrderAdd(clt *core.Client, req *OrderAddRequest) (order *Order, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/express/business/order/add?access_token="
	var result struct {
		core.Error
		Order
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	order = &result.Order
	return
}
