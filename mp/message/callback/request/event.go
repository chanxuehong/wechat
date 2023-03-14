package request

import (
	"fmt"
	"strings"

	"github.com/bububa/wechat/mp/core"
)

const (
	// 普通事件类型
	EventTypeSubscribe              core.EventType = "subscribe"                 // 关注事件, 包括点击关注和扫描二维码(公众号二维码和公众号带参数二维码)关注
	EventTypeUnsubscribe            core.EventType = "unsubscribe"               // 取消关注事件
	EventTypeScan                   core.EventType = "SCAN"                      // 已经关注的用户扫描带参数二维码事件
	EventTypeLocation               core.EventType = "LOCATION"                  // 上报地理位置事件
	EventTypeMessageSendJobFinish   core.EventType = "MASSSENDJOBFINISH"         // 事件推送群发结果
	EventTypeAddExpressPath         core.EventType = "add_express_path"          // 运单轨迹更新事件
	EventTypeSubscribeMsgPopupEvent core.EventType = "subscribe_msg_popup_event" // 小程序订阅消息事件
)

// 关注事件
type SubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event" json:"Event"` // subscribe

	// 下面两个字段只有在扫描带参数二维码进行关注时才有值, 否则为空值!
	EventKey string `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 格式为: qrscene_二维码的参数值
	Ticket   string `xml:"Ticket,omitempty"   json:"Ticket,omitempty"`   // 二维码的ticket, 可用来换取二维码图片
}

func GetSubscribeEvent(msg *core.MixedMsg) *SubscribeEvent {
	return &SubscribeEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
		Ticket:    msg.Ticket,
	}
}

// 获取二维码参数
func (event *SubscribeEvent) Scene() (scene string, err error) {
	const prefix = "qrscene_"
	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey 应该以 %s 为前缀: %s", prefix, event.EventKey)
		return
	}
	scene = event.EventKey[len(prefix):]
	return
}

// 取消关注事件
type UnsubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"              json:"Event"`              // unsubscribe
	EventKey  string         `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 空值
}

func GetUnsubscribeEvent(msg *core.MixedMsg) *UnsubscribeEvent {
	return &UnsubscribeEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
	}
}

// 用户已关注时, 扫描带参数二维码的事件
type ScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"    json:"Event"`    // SCAN
	EventKey  string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 二维码的参数值(scene_id, scene_str)
	Ticket    string         `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket, 可用来换取二维码图片
}

func GetScanEvent(msg *core.MixedMsg) *ScanEvent {
	return &ScanEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		EventKey:  msg.EventKey,
		Ticket:    msg.Ticket,
	}
}

// 上报地理位置事件
type LocationEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"     json:"Event"`     // LOCATION
	Latitude  float64        `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64        `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64        `xml:"Precision" json:"Precision"` // 地理位置精度(整数? 但是微信推送过来是浮点数形式)
}

func GetLocationEvent(msg *core.MixedMsg) *LocationEvent {
	return &LocationEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		Latitude:  msg.Latitude,
		Longitude: msg.Longitude,
		Precision: msg.Precision,
	}
}

// 事件推送群发结果
type MessageSendJobFinishEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType        core.EventType         `xml:"Event" json:"Event"`       // MASSSENDJOBFINISH
	MsgId            int64                  `xml:"MsgId"       json:"MsgId"` // 消息id, 64位整型
	Status           string                 `xml:"Status" json:"Status"`     // 群发的结果，为“send success”或“send fail”或“err(num)”。但send success时，也有可能因用户拒收公众号的消息、系统错误等原因造成少量用户接收失败。err(num)是审核失败的具体原因，可能的情况如下：err(10001):涉嫌广告, err(20001):涉嫌政治, err(20004):涉嫌社会, err(20002):涉嫌色情, err(20006):涉嫌违法犯罪, err(20008):涉嫌欺诈, err(20013):涉嫌版权, err(22000):涉嫌互推(互相宣传), err(21000):涉嫌其他, err(30001):原创校验出现系统错误且用户选择了被判为转载就不群发, err(30002): 原创校验被判定为不能群发, err(30003): 原创校验被判定为转载文且用户选择了被判为转载就不群发, err(40001)：管理员拒绝, err(40002)：管理员30分钟内无响应，超时
	TotalCount       int                    `xml:"TotalCount"  json:"TotalCount"`
	FilterCount      int                    `xml:"FilterCount" json:"FilterCount"`
	SentCount        int                    `xml:"SentCount"   json:"SentCount"`
	ErrorCount       int                    `xml:"ErrorCount"  json:"ErrorCount"`
	ArticleUrlResult *core.ArticleUrlResult `xml:"ArticleUrlResult" json:"ArticleUrlResult"`
}

func GetMessageSendJobFinishEvent(msg *core.MixedMsg) *MessageSendJobFinishEvent {
	return &MessageSendJobFinishEvent{
		MsgHeader:        msg.MsgHeader,
		EventType:        msg.EventType,
		MsgId:            msg.MsgId,
		Status:           msg.Status,
		TotalCount:       msg.TotalCount,
		FilterCount:      msg.FilterCount,
		SentCount:        msg.SentCount,
		ErrorCount:       msg.ErrorCount,
		ArticleUrlResult: msg.ArticleUrlResult,
	}
}

// 运单轨迹更新事件
type AddExpressPathEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType  core.EventType              `xml:"Event" json:"Event"`           // add_express_path
	MsgId      int64                       `xml:"MsgId"       json:"MsgId"`     // 消息id, 64位整型
	DeliveryID string                      `xml:"DeliveryID" json:"DeliveryID"` // 快递公司ID
	WaybillID  string                      `xml:"WayBillId" json:"WayBillId"`   // 运单ID
	OrderID    string                      `xml:"OrderId" json:"OrderId"`       // 订单ID
	Version    int                         `xml:"Version" json:"Version"`       // 轨迹版本号（整型）
	Count      int                         `xml:"Count" json:"Count"`           // 轨迹节点数（整型）
	Actions    []core.AddExpressPathAction `xml:"Actions" json:"Actions"`       // 轨迹列表
}

func GetAddExpressPathEvent(msg *core.MixedMsg) *AddExpressPathEvent {
	return &AddExpressPathEvent{
		MsgHeader:  msg.MsgHeader,
		EventType:  msg.EventType,
		MsgId:      msg.MsgId,
		DeliveryID: msg.DeliveryID,
		WaybillID:  msg.WaybillID,
		OrderID:    msg.OrderId,
		Version:    msg.Version,
		Count:      msg.Count,
		Actions:    msg.Actions,
	}
}

// 小程序订阅消息事件
type SubscribeMsgPopupEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType              core.EventType `xml:"Event" json:"Event"` // subscribe_msg_popup_event
	SubscribeMsgPopupEvent *core.SubscribeMsgPopupEvent
}

func GetSubscribeMsgPopupEvent(msg *core.MixedMsg) *SubscribeMsgPopupEvent {
	return &SubscribeMsgPopupEvent{
		MsgHeader:              msg.MsgHeader,
		EventType:              msg.EventType,
		SubscribeMsgPopupEvent: msg.SubscribeMsgPopupEvent,
	}
}
