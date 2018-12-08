package shakearound

import (
	"unsafe"

	"github.com/chanxuehong/wechat/mp/core"
)

const (
	// 推送到公众号URL上的事件类型
	EventTypeUserShake core.EventType = "ShakearoundUserShake" // 摇一摇事件通知
)

type UserShakeEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType core.EventType `xml:"Event" json:"Event"` // 事件类型，ShakearoundUserShake

	ChosenBeacon  *ChosenBeacon  `xml:"ChosenBeacon,omitempty" json:"ChosenBeacon,omitempty"`
	AroundBeacons []AroundBeacon `xml:"AroundBeacons>AroundBeacon,omitempty" json:"AroundBeacons,omitempty"`
}

// 和 github.com/chanxuehong/wechat/mp/core.MixedMsg.ChosenBeacon 一样, 同步修改
type ChosenBeacon struct {
	UUID     string  `xml:"Uuid"     json:"Uuid"`
	Major    int     `xml:"Major"    json:"Major"`
	Minor    int     `xml:"Minor"    json:"Minor"`
	Distance float64 `xml:"Distance" json:"Distance"`
}

// 和 github.com/chanxuehong/wechat/mp/core.MixedMsg.AroundBeacon 一样, 同步修改
type AroundBeacon struct {
	UUID     string  `xml:"Uuid"     json:"Uuid"`
	Major    int     `xml:"Major"    json:"Major"`
	Minor    int     `xml:"Minor"    json:"Minor"`
	Distance float64 `xml:"Distance" json:"Distance"`
}

func GetUserShakeEvent(msg *core.MixedMsg) *UserShakeEvent {
	return &UserShakeEvent{
		MsgHeader:     msg.MsgHeader,
		EventType:     msg.EventType,
		ChosenBeacon:  (*ChosenBeacon)(msg.ChosenBeacon),
		AroundBeacons: *(*[]AroundBeacon)(unsafe.Pointer(&msg.AroundBeacons)),
	}
}
