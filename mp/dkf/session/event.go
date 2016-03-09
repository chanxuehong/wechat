package session

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	EventTypeKfCreateSession core.EventType = "kf_create_session" // 接入会话
	EventTypeKfCloseSession  core.EventType = "kf_close_session"  // 关闭会话
	EventTypeKfSwitchSession core.EventType = "kf_switch_session" // 转接会话
)

type KfCreateSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"     json:"Event"`
	KfAccount string         `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCreateSessionEvent(msg *core.MixedMsg) *KfCreateSessionEvent {
	return &KfCreateSessionEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		KfAccount: msg.KfAccount,
	}
}

type KfCloseSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType core.EventType `xml:"Event"     json:"Event"`
	KfAccount string         `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCloseSessionEvent(msg *core.MixedMsg) *KfCloseSessionEvent {
	return &KfCloseSessionEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		KfAccount: msg.KfAccount,
	}
}

type KfSwitchSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType     core.EventType `xml:"Event"         json:"Event"`
	FromKfAccount string         `xml:"FromKfAccount" json:"FromKfAccount"`
	ToKfAccount   string         `xml:"ToKfAccount"   json:"ToKfAccount"`
}

func GetKfSwitchSessionEvent(msg *core.MixedMsg) *KfSwitchSessionEvent {
	return &KfSwitchSessionEvent{
		MsgHeader:     msg.MsgHeader,
		EventType:     msg.EventType,
		FromKfAccount: msg.FromKfAccount,
		ToKfAccount:   msg.ToKfAccount,
	}
}
