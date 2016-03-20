<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com), magicshui(shuiyuzhe@gmail.com), Harry Rong(harrykobe@gmail.com)

=======
>>>>>>> github/v2
package shakearound

import (
	"unsafe"

<<<<<<< HEAD
	"github.com/chanxuehong/wechat/mp"
=======
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

const (
	// 推送到公众号URL上的事件类型
<<<<<<< HEAD
	EventTypeUserShake = "ShakearoundUserShake" // 摇一摇事件通知
=======
	EventTypeUserShake core.EventType = "ShakearoundUserShake" // 摇一摇事件通知
>>>>>>> github/v2
)

type UserShakeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event string `xml:"Event" json:"Event"` // 事件类型，ShakearoundUserShake

	ChosenBeacon  ChosenBeacon   `xml:"ChosenBeacon"                         json:"ChosenBeacon"`
	AroundBeacons []AroundBeacon `xml:"AroundBeacons>AroundBeacon,omitempty" json:"AroundBeacons,omitempty"`
}

// 和 github.com/chanxuehong/wechat/mp.ChosenBeacon 一样, 同步修改
=======
	core.MsgHeader

	EventType core.EventType `xml:"Event" json:"Event"` // 事件类型，ShakearoundUserShake

	ChosenBeacon  *ChosenBeacon  `xml:"ChosenBeacon,omitempty" json:"ChosenBeacon,omitempty"`
	AroundBeacons []AroundBeacon `xml:"AroundBeacons>AroundBeacon,omitempty" json:"AroundBeacons,omitempty"`
}

// 和 github.com/chanxuehong/wechat/mp/core.MixedMsg.ChosenBeacon 一样, 同步修改
>>>>>>> github/v2
type ChosenBeacon struct {
	UUID     string  `xml:"Uuid"     json:"Uuid"`
	Major    int     `xml:"Major"    json:"Major"`
	Minor    int     `xml:"Minor"    json:"Minor"`
	Distance float64 `xml:"Distance" json:"Distance"`
}

<<<<<<< HEAD
// 和 github.com/chanxuehong/wechat/mp.AroundBeacon 一样, 同步修改
=======
// 和 github.com/chanxuehong/wechat/mp/core.MixedMsg.AroundBeacon 一样, 同步修改
>>>>>>> github/v2
type AroundBeacon struct {
	UUID     string  `xml:"Uuid"     json:"Uuid"`
	Major    int     `xml:"Major"    json:"Major"`
	Minor    int     `xml:"Minor"    json:"Minor"`
	Distance float64 `xml:"Distance" json:"Distance"`
}

<<<<<<< HEAD
func GetUserShakeEvent(msg *mp.MixedMessage) *UserShakeEvent {
	return &UserShakeEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		ChosenBeacon:  ChosenBeacon(msg.ChosenBeacon),
=======
func GetUserShakeEvent(msg *core.MixedMsg) *UserShakeEvent {
	return &UserShakeEvent{
		MsgHeader:     msg.MsgHeader,
		EventType:     msg.EventType,
		ChosenBeacon:  (*ChosenBeacon)(msg.ChosenBeacon),
>>>>>>> github/v2
		AroundBeacons: *(*[]AroundBeacon)(unsafe.Pointer(&msg.AroundBeacons)),
	}
}
