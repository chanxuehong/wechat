// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package component

const (
	// 微信服务器推送过来的消息类型
	ComponentMsgTypeComponentVerifyTicket = "component_verify_ticket" // 推送 component_verify_ticket 协议
	ComponentMsgTypeUnauthorized          = "unauthorized"            // 取消授权的通知
)

type ComponentVerifyTicket struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ComponentAppId string `xml:"AppId"      json:"AppId"`
	CreateTime     int64  `xml:"CreateTime" json:"CreateTime"`
	InfoType       string `xml:"InfoType"   json:"InfoType"`

	ComponentVerifyTicket string `xml:"ComponentVerifyTicket" json:"ComponentVerifyTicket"`
}

func GetComponentVerifyTicket(msg *MixedComponentMessage) *ComponentVerifyTicket {
	return &ComponentVerifyTicket{
		ComponentAppId:        msg.ComponentAppId,
		CreateTime:            msg.CreateTime,
		InfoType:              msg.InfoType,
		ComponentVerifyTicket: msg.ComponentVerifyTicket,
	}
}

type Unauthorized struct {
	XMLName struct{} `xml:"xml" json:"-"`

	ComponentAppId string `xml:"AppId"      json:"AppId"`
	CreateTime     int64  `xml:"CreateTime" json:"CreateTime"`
	InfoType       string `xml:"InfoType"   json:"InfoType"`

	AuthorizerAppid string `xml:"AuthorizerAppid" json:"AuthorizerAppid"`
}

func GetUnauthorized(msg *MixedComponentMessage) *Unauthorized {
	return &Unauthorized{
		ComponentAppId:  msg.ComponentAppId,
		CreateTime:      msg.CreateTime,
		InfoType:        msg.InfoType,
		AuthorizerAppid: msg.AuthorizerAppid,
	}
}
