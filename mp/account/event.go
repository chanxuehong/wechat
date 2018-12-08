package account

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	EventTypeQualificationVerifySuccess core.EventType = "qualification_verify_success" // 资质认证成功（此时立即获得接口权限）
	EventTypeQualificationVerifyFail    core.EventType = "qualification_verify_fail"    // 资质认证失败
	EventTypeNamingVerifySuccess        core.EventType = "naming_verify_success"        // 名称认证成功（即命名成功）
	EventTypeNamingVerifyFail           core.EventType = "naming_verify_fail"           // 名称认证失败（这时虽然客户端不打勾，但仍有接口权限）
	EventTypeAnnualRenew                core.EventType = "annual_renew"                 // 年审通知
	EventTypeVerifyExpired              core.EventType = "verify_expired"               // 认证过期失效通知
)

// 资质认证成功（此时立即获得接口权限）事件
type QualificationVerifySuccessEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType   core.EventType `xml:"Event"       json:"Event"`
	ExpiredTime int64          `xml:"ExpiredTime" json:"ExpiredTime"` // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
}

func GetQualificationVerifySuccessEvent(msg *core.MixedMsg) *QualificationVerifySuccessEvent {
	return &QualificationVerifySuccessEvent{
		MsgHeader:   msg.MsgHeader,
		EventType:   msg.EventType,
		ExpiredTime: msg.ExpiredTime,
	}
}

// 资质认证失败事件
type QualificationVerifyFailEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType  core.EventType `xml:"Event"      json:"Event"`
	FailTime   int64          `xml:"FailTime"   json:"FailTime"`   // 失败发生时间 (整形)，时间戳
	FailReason string         `xml:"FailReason" json:"FailReason"` // 认证失败的原因
}

func GetQualificationVerifyFailEvent(msg *core.MixedMsg) *QualificationVerifyFailEvent {
	return &QualificationVerifyFailEvent{
		MsgHeader:  msg.MsgHeader,
		EventType:  msg.EventType,
		FailTime:   msg.FailTime,
		FailReason: msg.FailReason,
	}
}

// 名称认证成功（即命名成功）事件
type NamingVerifySuccessEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType   core.EventType `xml:"Event"       json:"Event"`
	ExpiredTime int64          `xml:"ExpiredTime" json:"ExpiredTime"` // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期
}

func GetNamingVerifySuccessEvent(msg *core.MixedMsg) *NamingVerifySuccessEvent {
	return &NamingVerifySuccessEvent{
		MsgHeader:   msg.MsgHeader,
		EventType:   msg.EventType,
		ExpiredTime: msg.ExpiredTime,
	}
}

// 名称认证失败（这时虽然客户端不打勾，但仍有接口权限）事件
type NamingVerifyFailEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType  core.EventType `xml:"Event"      json:"Event"`
	FailTime   int64          `xml:"FailTime"   json:"FailTime"`   // 失败发生时间 (整形)，时间戳
	FailReason string         `xml:"FailReason" json:"FailReason"` // 认证失败的原因
}

func GetNamingVerifyFailEvent(msg *core.MixedMsg) *NamingVerifyFailEvent {
	return &NamingVerifyFailEvent{
		MsgHeader:  msg.MsgHeader,
		EventType:  msg.EventType,
		FailTime:   msg.FailTime,
		FailReason: msg.FailReason,
	}
}

// 年审通知事件
type AnnualRenewEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType   core.EventType `xml:"Event"       json:"Event"`
	ExpiredTime int64          `xml:"ExpiredTime" json:"ExpiredTime"` // 有效期 (整形)，指的是时间戳，将于该时间戳认证过期，需尽快年审
}

func GetAnnualRenewEvent(msg *core.MixedMsg) *AnnualRenewEvent {
	return &AnnualRenewEvent{
		MsgHeader:   msg.MsgHeader,
		EventType:   msg.EventType,
		ExpiredTime: msg.ExpiredTime,
	}
}

// 认证过期失效通知事件
type VerifyExpiredEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	EventType   core.EventType `xml:"Event"       json:"Event"`
	ExpiredTime int64          `xml:"ExpiredTime" json:"ExpiredTime"` // 有效期 (整形)，指的是时间戳，表示已于该时间戳认证过期，需要重新发起微信认证
}

func GetVerifyExpiredEvent(msg *core.MixedMsg) *VerifyExpiredEvent {
	return &VerifyExpiredEvent{
		MsgHeader:   msg.MsgHeader,
		EventType:   msg.EventType,
		ExpiredTime: msg.ExpiredTime,
	}
}
