// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package request

import (
	"fmt"
	"strconv"
	"strings"
)

// 关注事件(普通关注)
type SubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event string `xml:"Event" json:"Event"` // 事件类型，subscribe(订阅)
}

func (req *Request) SubscribeEvent() (event *SubscribeEvent) {
	event = &SubscribeEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
	}
	return
}

// 取消关注
type UnsubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event string `xml:"Event" json:"Event"` // 事件类型，unsubscribe(取消订阅)
}

func (req *Request) UnsubscribeEvent() (event *UnsubscribeEvent) {
	event = &UnsubscribeEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
	}
	return
}

// 用户未关注时，扫描带参数二维码进行关注后的事件推送
type SubscribeByScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，subscribe
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，qrscene_为前缀，后面为二维码的参数值
	Ticket   string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

func (req *Request) SubscribeByScanEvent() (event *SubscribeByScanEvent) {
	event = &SubscribeByScanEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		EventKey:   req.EventKey,
		Ticket:     req.Ticket,
	}
	return
}

// 获取二维码参数
func (event *SubscribeByScanEvent) SceneId() (id uint32, err error) {
	const prefix = "qrscene_"

	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey 应该以 %s 为前缀, 但是现在是 %s", prefix, event.EventKey)
		return
	}

	id64, err := strconv.ParseUint(event.EventKey[len(prefix):], 10, 32)
	if err != nil {
		return
	}
	id = uint32(id64)
	return
}

// 用户已关注时，扫描带参数二维码的事件推送
type ScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event    string `xml:"Event"    json:"Event"`    // 事件类型，SCAN
	EventKey string `xml:"EventKey" json:"EventKey"` // 事件KEY值，是一个32位无符号整数，即创建二维码时的二维码scene_id
	Ticket   string `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket，可用来换取二维码图片
}

func (req *Request) ScanEvent() (event *ScanEvent) {
	event = &ScanEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		EventKey:   req.EventKey,
		Ticket:     req.Ticket,
	}
	return
}

// 获取二维码参数
func (event *ScanEvent) SceneId() (id uint32, err error) {
	id64, err := strconv.ParseUint(event.EventKey, 10, 32)
	if err != nil {
		return
	}
	id = uint32(id64)
	return
}

// 上报地理位置事件
type LocationEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event     string  `xml:"Event"     json:"Event"`     // 事件类型，LOCATION
	Latitude  float64 `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64 `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64 `xml:"Precision" json:"Precision"` // 地理位置精度
}

func (req *Request) LocationEvent() (event *LocationEvent) {
	event = &LocationEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		Latitude:   req.Latitude,
		Longitude:  req.Longitude,
		Precision:  req.Precision,
	}
	return
}

// 高级群发消息, 事件推送群发结果
type MassSendJobFinishEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event string `xml:"Event" json:"Event"` // 事件信息，此处为 MASSSENDJOBFINISH

	MsgId int64 `xml:"MsgId" json:"MsgId"` // 群发的消息ID, 64位整型

	// 群发的结构, 为 "send success" 或 "send fail" 或 "err(num)".
	// 但 send success 时, 也有可能因用户拒收公众号的消息, 系统错误等原因造成少量用户接收失败.
	// err(num) 是审核失败的具体原因, 可能的情况如下:
	// err(10001), //涉嫌广告
	// err(20001), //涉嫌政治
	// err(20004), //涉嫌社会
	// err(20002), //涉嫌色情
	// err(20006), //涉嫌违法犯罪
	// err(20008), //涉嫌欺诈
	// err(20013), //涉嫌版权
	// err(22000), //涉嫌互推(互相宣传)
	// err(21000), //涉嫌其他
	Status string `xml:"Status" json:"Status"`

	TotalCount int `xml:"TotalCount" json:"TotalCount"` // group_id 下粉丝数, 或者 openid_list 中的粉丝数

	// 过滤(过滤是指特定地区, 性别的过滤, 用户设置拒收的过滤; 用户接收已超4条的过滤）后,
	// 准备发送的粉丝数, 原则上, FilterCount = SentCount + ErrorCount
	FilterCount int `xml:"FilterCount" json:"FilterCount"`
	SentCount   int `xml:"SentCount"   json:"SentCount"`  // 发送成功的粉丝数
	ErrorCount  int `xml:"ErrorCount"  json:"ErrorCount"` // 发送失败的粉丝数
}

func (req *Request) MassSendJobFinishEvent() (event *MassSendJobFinishEvent) {
	event = &MassSendJobFinishEvent{
		CommonHead:  req.CommonHead,
		Event:       req.Event,
		MsgId:       req.MsgID, // NOTE
		Status:      req.Status,
		TotalCount:  req.TotalCount,
		FilterCount: req.FilterCount,
		SentCount:   req.SentCount,
		ErrorCount:  req.ErrorCount,
	}
	return
}

// 在模版消息发送任务完成后，微信服务器会将是否送达成功作为通知，发送到开发者中心中填写的服务器配置地址中。
type TemplateSendJobFinishEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event string `xml:"Event" json:"Event"` // 事件信息，此处为 TEMPLATESENDJOBFINISH

	MsgId int64 `xml:"MsgId" json:"MsgId"` // 模板消息ID

	// 送达成功时:                                     success
	// 送达由于用户拒收（用户设置拒绝接收公众号消息）而失败时:  failed:user block
	// 送达由于其他原因失败时:                            failed: system failed
	Status string `xml:"Status" json:"Status"`
}

func (req *Request) TemplateSendJobFinishEvent() (event *TemplateSendJobFinishEvent) {
	event = &TemplateSendJobFinishEvent{
		CommonHead: req.CommonHead,
		Event:      req.Event,
		MsgId:      req.MsgID, // NOTE
		Status:     req.Status,
	}
	return
}

// 微信小店, 订单付款通知
type MerchantOrderEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	Event       string `xml:"Event"       json:"Event"`       // 事件类型, merchant_order
	OrderId     string `xml:"OrderId"     json:"OrderId"`     // 订单 id
	OrderStatus int    `xml:"OrderStatus" json:"OrderStatus"` // 订单状态(2-待发货, 3-已发货, 5-已完成, 8-维权中)
	ProductId   string `xml:"ProductId"   json:"ProductId"`   // 商品 id
	SkuInfo     string `xml:"SkuInfo"     json:"SkuInfo"`     // sku 信息
}

func (req *Request) MerchantOrderEvent() (event *MerchantOrderEvent) {
	event = &MerchantOrderEvent{
		CommonHead:  req.CommonHead,
		Event:       req.Event,
		OrderId:     req.OrderId,
		OrderStatus: req.OrderStatus,
		ProductId:   req.ProductId,
		SkuInfo:     req.SkuInfo,
	}
	return
}
