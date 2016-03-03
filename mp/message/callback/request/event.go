package request

import (
	"fmt"
	"strings"

	"github.com/chanxuehong/wechat/mp/core"
)

const (
	// 微信服务器推送过来的事件类型
	EventTypeSubscribe   core.EventType = "subscribe"   // 订阅, 包括点击订阅和扫描二维码(公众号二维码和公众号带参数二维码)订阅
	EventTypeUnsubscribe core.EventType = "unsubscribe" // 取消订阅
	EventTypeScan        core.EventType = "SCAN"        // 已经订阅的用户扫描带参数二维码事件
	EventTypeLocation    core.EventType = "LOCATION"    // 上报地理位置事件
)

// 关注事件.
// 普通关注, 扫描公众号二维码(不是带参数二维码)关注
type SubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	Event    core.EventType `xml:"Event"              json:"Event"`              // subscribe(订阅)
	EventKey string         `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 为空值
}

func GetSubscribeEvent(msg *core.MixedMsg) *SubscribeEvent {
	return &SubscribeEvent{
		MsgHeader: msg.MsgHeader,
		Event:     msg.Event,
	}
}

// 取消关注
type UnsubscribeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	Event    core.EventType `xml:"Event"              json:"Event"`              // unsubscribe(取消订阅)
	EventKey string         `xml:"EventKey,omitempty" json:"EventKey,omitempty"` // 事件KEY值, 为空值
}

func GetUnsubscribeEvent(msg *core.MixedMsg) *UnsubscribeEvent {
	return &UnsubscribeEvent{
		MsgHeader: msg.MsgHeader,
		Event:     msg.Event,
	}
}

// 用户未关注时, 扫描带参数二维码进行关注后的事件
type SubscribeByScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	Event    core.EventType `xml:"Event"    json:"Event"`    // subscribe
	EventKey string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, qrscene_为前缀, 后面为二维码的参数值(scene_id, scene_str)
	Ticket   string         `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket, 可用来换取二维码图片
}

// 获取二维码参数
func (event *SubscribeByScanEvent) Scene() (scene string, err error) {
	const prefix = "qrscene_"
	if !strings.HasPrefix(event.EventKey, prefix) {
		err = fmt.Errorf("EventKey 应该以 %q 为前缀: %q", prefix, event.EventKey)
		return
	}
	scene = event.EventKey[len(prefix):]
	return
}

func GetSubscribeByScanEvent(msg *core.MixedMsg) *SubscribeByScanEvent {
	return &SubscribeByScanEvent{
		MsgHeader: msg.MsgHeader,
		Event:     msg.Event,
		EventKey:  msg.EventKey,
		Ticket:    msg.Ticket,
	}
}

// 用户已关注时, 扫描带参数二维码的事件
type ScanEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	Event    core.EventType `xml:"Event"    json:"Event"`    // SCAN
	EventKey string         `xml:"EventKey" json:"EventKey"` // 事件KEY值, 二维码的参数值(scene_id, scene_str)
	Ticket   string         `xml:"Ticket"   json:"Ticket"`   // 二维码的ticket, 可用来换取二维码图片
}

func GetScanEvent(msg *core.MixedMsg) *ScanEvent {
	return &ScanEvent{
		MsgHeader: msg.MsgHeader,
		Event:     msg.Event,
		EventKey:  msg.EventKey,
		Ticket:    msg.Ticket,
	}
}

// 上报地理位置事件
type LocationEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	Event     core.EventType `xml:"Event"     json:"Event"`     // LOCATION
	Latitude  float64        `xml:"Latitude"  json:"Latitude"`  // 地理位置纬度
	Longitude float64        `xml:"Longitude" json:"Longitude"` // 地理位置经度
	Precision float64        `xml:"Precision" json:"Precision"` // 地理位置精度(实际上应该是整数, 但是微信推送过来是浮点数形式)
}

func GetLocationEvent(msg *core.MixedMsg) *LocationEvent {
	return &LocationEvent{
		MsgHeader: msg.MsgHeader,
		Event:     msg.Event,
		Latitude:  msg.Latitude,
		Longitude: msg.Longitude,
		Precision: msg.Precision,
	}
}
