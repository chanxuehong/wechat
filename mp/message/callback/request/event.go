package request

import (
	"fmt"
	"strings"

	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

const (
	// 普通事件类型
	EventTypeSubscribe   core.EventType = "subscribe"   // 关注事件, 包括点击关注和扫描二维码(公众号二维码和公众号带参数二维码)关注
	EventTypeUnsubscribe core.EventType = "unsubscribe" // 取消关注事件
	EventTypeScan        core.EventType = "SCAN"        // 已经关注的用户扫描带参数二维码事件
	EventTypeLocation    core.EventType = "LOCATION"    // 上报地理位置事件
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
