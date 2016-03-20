<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package bizwifi

import (
	"github.com/chanxuehong/wechat/mp"
=======
package bizwifi

import (
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

const (
	// 推送到公众号URL上的事件类型
<<<<<<< HEAD
	EventTypeWifiConnected = "WifiConnected" // Wi-Fi连网成功事件
=======
	EventTypeWifiConnected core.EventType = "WifiConnected" // Wi-Fi连网成功事件
>>>>>>> github/v2
)

type WifiConnectedEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型，WifiConnected (Wi-Fi连网成功)
=======
	core.MsgHeader

	EventType core.EventType `xml:"Event" json:"Event"` // 事件类型，WifiConnected (Wi-Fi连网成功)
>>>>>>> github/v2

	ConnectTime int64  `xml:"ConnectTime" json:"ConnectTime"` // 连网时间（整型）
	ExpireTime  int64  `xml:"ExpireTime"  json:"ExpireTime"`  // 系统保留字段，固定值
	VendorId    string `xml:"VendorId"    json:"VendorId"`    // 系统保留字段，固定值
	PlaceId     int64  `xml:"PlaceId"     json:"PlaceId"`     // 连网的门店id
	DeviceNo    string `xml:"DeviceNo"    json:"DeviceNo"`    // 连网的设备无线mac地址，对应bssid
}

<<<<<<< HEAD
func GetWifiConnectedEvent(msg *mp.MixedMessage) *WifiConnectedEvent {
	return &WifiConnectedEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		ConnectTime:   msg.ConnectTime,
		ExpireTime:    msg.ExpireTime,
		VendorId:      msg.VendorId,
		PlaceId:       msg.PlaceId,
		DeviceNo:      msg.DeviceNo,
=======
func GetWifiConnectedEvent(msg *core.MixedMsg) *WifiConnectedEvent {
	return &WifiConnectedEvent{
		MsgHeader:   msg.MsgHeader,
		EventType:   msg.EventType,
		ConnectTime: msg.ConnectTime,
		ExpireTime:  msg.ExpireTime,
		VendorId:    msg.VendorId,
		PlaceId:     msg.PlaceId,
		DeviceNo:    msg.DeviceNo,
>>>>>>> github/v2
	}
}
