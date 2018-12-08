package poi

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	EventTypePoiCheckNotify core.EventType = "poi_check_notify" // 审核事件推送
)

// 创建门店审核事件推送
type PoiCheckNotifyEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"  json:"Event"`  // 事件类型，poi_check_notify
	UniqId    string         `xml:"UniqId" json:"UniqId"` // 商户自己内部ID，即字段中的sid
	PoiId     int64          `xml:"PoiId"  json:"PoiId"`  // 微信的门店ID，微信内门店唯一标示ID
	Result    string         `xml:"Result" json:"Result"` // 审核结果，成功succ 或失败fail
	Msg       string         `xml:"Msg"    json:"Msg"`    // 成功的通知信息，或审核失败的驳回理由
}

func GetPoiCheckNotifyEvent(msg *core.MixedMsg) *PoiCheckNotifyEvent {
	return &PoiCheckNotifyEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		UniqId:    msg.UniqId,
		PoiId:     msg.PoiId,
		Result:    msg.Result,
		Msg:       msg.Msg,
	}
}
