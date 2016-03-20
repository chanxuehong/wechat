<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package session

import (
	"github.com/chanxuehong/wechat/mp"
)

const (
	EventTypeKfCreateSession = "kf_create_session" // 接入会话
	EventTypeKfCloseSession  = "kf_close_session"  // 关闭会话
	EventTypeKfSwitchSession = "kf_switch_session" // 转接会话
=======
package session

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	EventTypeKfCreateSession core.EventType = "kf_create_session" // 接入会话
	EventTypeKfCloseSession  core.EventType = "kf_close_session"  // 关闭会话
	EventTypeKfSwitchSession core.EventType = "kf_switch_session" // 转接会话
>>>>>>> github/v2
)

type KfCreateSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event     string `xml:"Event"     json:"Event"`
	KfAccount string `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCreateSessionEvent(msg *mp.MixedMessage) *KfCreateSessionEvent {
	return &KfCreateSessionEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		KfAccount:     msg.KfAccount,
=======
	core.MsgHeader
	EventType core.EventType `xml:"Event"     json:"Event"`
	KfAccount string         `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCreateSessionEvent(msg *core.MixedMsg) *KfCreateSessionEvent {
	return &KfCreateSessionEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		KfAccount: msg.KfAccount,
>>>>>>> github/v2
	}
}

type KfCloseSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event     string `xml:"Event"     json:"Event"`
	KfAccount string `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCloseSessionEvent(msg *mp.MixedMessage) *KfCloseSessionEvent {
	return &KfCloseSessionEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
		KfAccount:     msg.KfAccount,
=======
	core.MsgHeader
	EventType core.EventType `xml:"Event"     json:"Event"`
	KfAccount string         `xml:"KfAccount" json:"KfAccount"`
}

func GetKfCloseSessionEvent(msg *core.MixedMsg) *KfCloseSessionEvent {
	return &KfCloseSessionEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		KfAccount: msg.KfAccount,
>>>>>>> github/v2
	}
}

type KfSwitchSessionEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
<<<<<<< HEAD
	mp.MessageHeader

	Event         string `xml:"Event"         json:"Event"`
	FromKfAccount string `xml:"FromKfAccount" json:"FromKfAccount"`
	ToKfAccount   string `xml:"ToKfAccount"   json:"ToKfAccount"`
}

func GetKfSwitchSessionEvent(msg *mp.MixedMessage) *KfSwitchSessionEvent {
	return &KfSwitchSessionEvent{
		MessageHeader: msg.MessageHeader,
		Event:         msg.Event,
=======
	core.MsgHeader
	EventType     core.EventType `xml:"Event"         json:"Event"`
	FromKfAccount string         `xml:"FromKfAccount" json:"FromKfAccount"`
	ToKfAccount   string         `xml:"ToKfAccount"   json:"ToKfAccount"`
}

func GetKfSwitchSessionEvent(msg *core.MixedMsg) *KfSwitchSessionEvent {
	return &KfSwitchSessionEvent{
		MsgHeader:     msg.MsgHeader,
		EventType:     msg.EventType,
>>>>>>> github/v2
		FromKfAccount: msg.FromKfAccount,
		ToKfAccount:   msg.ToKfAccount,
	}
}
