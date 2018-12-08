package template

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	EventTypeTemplateSendJobFinish core.EventType = "TEMPLATESENDJOBFINISH"
)

const (
	TemplateSendStatusSuccess            = "success"               // 送达成功时
	TemplateSendStatusFailedUserBlock    = "failed:user block"     // 送达由于用户拒收(用户设置拒绝接收公众号消息)而失败
	TemplateSendStatusFailedSystemFailed = "failed: system failed" // 送达由于其他原因失败
)

type TemplateSendJobFinishEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"  json:"Event"`  // 此处为 TEMPLATESENDJOBFINISH
	MsgId     int64          `xml:"MsgId"  json:"MsgId"`  // 模板消息ID
	Status    string         `xml:"Status" json:"Status"` // 发送状态
}

func GetTemplateSendJobFinishEvent(msg *core.MixedMsg) *TemplateSendJobFinishEvent {
	return &TemplateSendJobFinishEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		MsgId:     msg.MsgID, // NOTE
		Status:    msg.Status,
	}
}
