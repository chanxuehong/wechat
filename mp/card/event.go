package card

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	// 推送到公众号URL上的事件类型
	EventTypeCardPassCheck            core.EventType = "card_pass_check"              // 卡券通过审核
	EventTypeCardNotPassCheck         core.EventType = "card_not_pass_check"          // 卡券未通过审核
	EventTypeUserGiftingCard          core.EventType = "user_gifting_card"            // 转赠事件推送
	EventTypeUserGetCard              core.EventType = "user_get_card"                // 领取事件推送
	EventTypeUserDelCard              core.EventType = "user_del_card"                // 删除事件推送
	EventTypeUserConsumeCard          core.EventType = "user_consume_card"            // 核销事件推送
	EventTypeUserViewCard             core.EventType = "user_view_card"               // 进入会员卡事件推送
	EventTypeUserEnterSessionFromCard core.EventType = "user_enter_session_from_card" // 从卡券进入公众号会话事件推送
	EventTypeCardSkuRemind            core.EventType = "card_sku_remind"              // 库存报警事件

	EventTypeGiftCardPayDone    core.EventType = "giftcard_pay_done"    // 用户购买礼品卡付款成功
	EventTypeGiftCardUserAccept core.EventType = "giftcard_user_accept" // 用户领取礼品卡成功
)

// 卡券通过审核, 微信会把这个事件推送到开发者填写的URL
type CardPassCheckEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType core.EventType `xml:"Event"  json:"Event"`  // 事件类型, card_pass_check
	CardId    string         `xml:"CardId" json:"CardId"` // 卡券ID
}

func GetCardPassCheckEvent(msg *core.MixedMsg) *CardPassCheckEvent {
	return &CardPassCheckEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		CardId:    msg.CardId,
	}
}

// 卡券未通过审核, 微信会把这个事件推送到开发者填写的URL
type CardNotPassCheckEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType    core.EventType `xml:"Event"        json:"Event"`  // 事件类型, card_not_pass_check
	CardId       string         `xml:"CardId"       json:"CardId"` // 卡券ID
	RefuseReason string         `xml:"RefuseReason" json:"RefuseReason"`
}

func GetCardNotPassCheckEvent(msg *core.MixedMsg) *CardNotPassCheckEvent {
	return &CardNotPassCheckEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		CardId:       msg.CardId,
		RefuseReason: msg.RefuseReason,
	}
}

// 用户在领取卡券时, 微信会把这个事件推送到开发者填写的URL.
type UserGetCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType           core.EventType `xml:"Event"               json:"Event"`           // 事件类型, user_get_card
	CardId              string         `xml:"CardId"              json:"CardId"`          // 卡券ID
	IsGiveByFriend      int            `xml:"IsGiveByFriend"      json:"IsGiveByFriend"`  // 是否为转赠, 1 代表是, 0 代表否.
	FriendUserName      string         `xml:"FriendUserName"      json:"FriendUserName"`  // 赠送方账号(一个OpenID), "IsGiveByFriend"为1 时填写该参数.
	UserCardCode        string         `xml:"UserCardCode"        json:"UserCardCode"`    // code 序列号. 自定义code 及非自定义code的卡券被领取后都支持事件推送.
	OldUserCardCode     string         `xml:"OldUserCardCode"     json:"OldUserCardCode"` // 转赠前的code序列号。
	OuterId             int64          `xml:"OuterId"             json:"OuterId"`         // 领取场景值, 用于领取渠道数据统计. 可在生成二维码接口及添加JS API 接口中自定义该字段的整型值.
	OuterStr            string         `xml:"OuterStr"            json:"OuterStr"`
	IsRestoreMemberCard int            `xml:"IsRestoreMemberCard" json:"IsRestoreMemberCard"`
	IsRecommendByFriend int            `xml:"IsRecommendByFriend" json:"IsRecommendByFriend"`
}

func GetUserGetCardEvent(msg *core.MixedMsg) *UserGetCardEvent {
	return &UserGetCardEvent{
		MsgHeader:           msg.MsgHeader,
		EventType:           msg.EventType,
		CardId:              msg.CardId,
		IsGiveByFriend:      msg.IsGiveByFriend,
		FriendUserName:      msg.FriendUserName,
		UserCardCode:        msg.UserCardCode,
		OldUserCardCode:     msg.OldUserCardCode,
		OuterId:             msg.OuterId,
		OuterStr:            msg.OuterStr,
		IsRestoreMemberCard: msg.IsRestoreMemberCard,
		IsRecommendByFriend: msg.IsRecommendByFriend,
	}
}

// 转赠事件推送.
type UserGiftingCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType      core.EventType `xml:"Event"          json:"Event"`
	CardId         string         `xml:"CardId"         json:"CardId"`
	UserCardCode   string         `xml:"UserCardCode"   json:"UserCardCode"`
	IsReturnBack   int            `xml:"IsReturnBack"   json:"IsReturnBack"`
	FriendUserName string         `xml:"FriendUserName" json:"FriendUserName"`
	IsChatRoom     int            `xml:"IsChatRoom"     json:"IsChatRoom"`
}

func GetUserGiftingCardEvent(msg *core.MixedMsg) *UserGiftingCardEvent {
	return &UserGiftingCardEvent{
		MsgHeader:      msg.MsgHeader,
		EventType:      msg.EventType,
		CardId:         msg.CardId,
		UserCardCode:   msg.UserCardCode,
		IsReturnBack:   msg.IsReturnBack,
		FriendUserName: msg.FriendUserName,
		IsChatRoom:     msg.IsChatRoom,
	}
}

// 用户在删除卡券时, 微信会把这个事件推送到开发者填写的URL.
type UserDelCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType    core.EventType `xml:"Event"        json:"Event"`        // 事件类型, user_del_card
	CardId       string         `xml:"CardId"       json:"CardId"`       // 卡券ID
	UserCardCode string         `xml:"UserCardCode" json:"UserCardCode"` // 商户自定义code 值. 非自定code 推送为空串
}

func GetUserDelCardEvent(msg *core.MixedMsg) *UserDelCardEvent {
	return &UserDelCardEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		CardId:       msg.CardId,
		UserCardCode: msg.UserCardCode,
	}
}

// 核销事件推送
type UserConsumeCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType     core.EventType `xml:"Event"         json:"Event"`         // 事件类型, user_consume_card
	CardId        string         `xml:"CardId"        json:"CardId"`        // 卡券ID
	UserCardCode  string         `xml:"UserCardCode"  json:"UserCardCode"`  // 商户自定义code 值. 非自定code 推送为空串
	ConsumeSource string         `xml:"ConsumeSource" json:"ConsumeSource"` // 核销来源。支持开发者统计API核销（FROM_API）、公众平台核销（FROM_MP）、卡券商户助手核销（FROM_MOBILE_HELPER）（核销员微信号）
	LocationName  string         `xml:"LocationName"  json:"LocationName"`  // 门店名称，当前卡券核销的门店名称（只有通过自助核销和买单核销时才会出现该字段）
	StaffOpenId   string         `xml:"StaffOpenId"   json:"StaffOpenId"`   // 核销该卡券核销员的openid（只有通过卡券商户助手核销时才会出现）
	VerifyCode    string         `xml:"VerifyCode"    json:"VerifyCode"`    // 自助核销时，用户输入的验证码
	RemarkAmount  string         `xml:"RemarkAmount"  json:"RemarkAmount"`  // 自助核销时，用户输入的备注金额
	OuterStr      string         `xml:"OuterStr"      json:"OuterStr"`      // 开发者发起核销时传入的自定义参数，用于进行核销渠道统计
}

func GetUserConsumeCardEvent(msg *core.MixedMsg) *UserConsumeCardEvent {
	return &UserConsumeCardEvent{
		MsgHeader:     msg.MsgHeader,
		EventType:     msg.EventType,
		CardId:        msg.CardId,
		UserCardCode:  msg.UserCardCode,
		ConsumeSource: msg.ConsumeSource,
		LocationName:  msg.LocationName,
		StaffOpenId:   msg.StaffOpenId,
		VerifyCode:    msg.VerifyCode,
		RemarkAmount:  msg.RemarkAmount,
		OuterStr:      msg.OuterStr,
	}
}

// 用户在进入会员卡时, 微信会把这个事件推送到开发者填写的URL
type UserViewCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType    core.EventType `xml:"Event"        json:"Event"`        // 事件类型, user_view_card
	CardId       string         `xml:"CardId"       json:"CardId"`       // 卡券ID
	UserCardCode string         `xml:"UserCardCode" json:"UserCardCode"` // 商户自定义code 值. 非自定code 推送为空串
}

func GetUserViewCardEvent(msg *core.MixedMsg) *UserViewCardEvent {
	return &UserViewCardEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		CardId:       msg.CardId,
		UserCardCode: msg.UserCardCode,
	}
}

// 从卡券进入公众号会话事件推送
type UserEnterSessionFromCardEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType    core.EventType `xml:"Event"        json:"Event"`        // 事件类型, user_view_card
	CardId       string         `xml:"CardId"       json:"CardId"`       // 卡券ID
	UserCardCode string         `xml:"UserCardCode" json:"UserCardCode"` // 商户自定义code 值. 非自定code 推送为空串
}

func GetUserEnterSessionFromCardEvent(msg *core.MixedMsg) *UserEnterSessionFromCardEvent {
	return &UserEnterSessionFromCardEvent{
		MsgHeader:    msg.MsgHeader,
		EventType:    msg.EventType,
		CardId:       msg.CardId,
		UserCardCode: msg.UserCardCode,
	}
}

// 库存报警事件
type CardSkuRemindEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType core.EventType `xml:"Event"  json:"Event"`  // 事件类型, card_sku_remind
	CardId    string         `xml:"CardId" json:"CardId"` // 卡券ID
	Detail    string         `xml:"Detail" json:"Detail"` // 报警详细信息
}

func GetCardSkuRemindEvent(msg *core.MixedMsg) *CardSkuRemindEvent {
	return &CardSkuRemindEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		CardId:    msg.CardId,
		Detail:    msg.Detail,
	}
}

// 用户购买礼品卡付款成功
type GiftCardPayDoneEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType core.EventType `xml:"Event"   json:"Event"`   // 事件类型，此处为giftcard_pay_done标识订单完成事件
	PageId    string         `xml:"PageId"  json:"PageId"`  // 货架的id
	OrderId   string         `xml:"OrderId" json:"OrderId"` // 订单号
}

func GetGiftCardPayDoneEvent(msg *core.MixedMsg) *GiftCardPayDoneEvent {
	return &GiftCardPayDoneEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		PageId:    msg.PageId,
		OrderId:   msg.OrderId,
	}
}

// 用户领取礼品卡成功
type GiftCardUserAcceptEvent struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader

	EventType core.EventType `xml:"Event"   json:"Event"`   // 事件类型，此处为giftcard_user_accept标识订单完成事件
	PageId    string         `xml:"PageId"  json:"PageId"`  // 货架的id
	OrderId   string         `xml:"OrderId" json:"OrderId"` // 订单号
}

func GetGiftCardUserAcceptEvent(msg *core.MixedMsg) *GiftCardUserAcceptEvent {
	return &GiftCardUserAcceptEvent{
		MsgHeader: msg.MsgHeader,
		EventType: msg.EventType,
		PageId:    msg.PageId,
		OrderId:   msg.OrderId,
	}
}
