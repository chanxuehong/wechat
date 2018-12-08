package bizwifi

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	// 推送到公众号URL上的事件类型
	EventTypeWifiConnected core.EventType = "WifiConnected" // Wi-Fi连网成功事件
)

type WifiConnectedEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType core.EventType `xml:"Event" json:"Event"` // 事件类型，WifiConnected (Wi-Fi连网成功)

	ConnectTime int64  `xml:"ConnectTime" json:"ConnectTime"` // 连网时间（整型）
	ExpireTime  int64  `xml:"ExpireTime"  json:"ExpireTime"`  // 系统保留字段，固定值
	VendorId    string `xml:"VendorId"    json:"VendorId"`    // 系统保留字段，固定值
	PlaceId     int64  `xml:"PlaceId"     json:"PlaceId"`     // 连网的门店id
	DeviceNo    string `xml:"DeviceNo"    json:"DeviceNo"`    // 连网的设备无线mac地址，对应bssid
}

func GetWifiConnectedEvent(msg *core.MixedMsg) *WifiConnectedEvent {
	return &WifiConnectedEvent{
		MsgHeader:   msg.MsgHeader,
		EventType:   msg.EventType,
		ConnectTime: msg.ConnectTime,
		ExpireTime:  msg.ExpireTime,
		VendorId:    msg.VendorId,
		PlaceId:     msg.PlaceId,
		DeviceNo:    msg.DeviceNo,
	}
}
