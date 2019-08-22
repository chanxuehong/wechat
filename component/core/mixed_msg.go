package core

// 微信服务器推送过来的消息(事件)的通用消息头.
type MsgHeader struct {
	AppId      string `xml:"AppId" json:"AppId"`
	InfoType   string `xml:"InfoType" json:"InfoType"`
	CreateTime int64  `xml:"CreateTime" json:"CreateTime"`
}

// 微信服务器推送过来的消息(事件)的合集.
type MixedMsg struct {
	XMLName struct{} `xml:"xml" json:"-"`
	MsgHeader
	ComponentVerifyTicket        string `xml:"ComponentVerifyTicket,omitempty"      json:"ComponentVerifyTicket,omitempty"`
	AuthorizerAppid              string `xml:"AuthorizerAppid,omitempty" json:"AuthorizerAppid,omitempty"`                           // 公众号或小程序
	AuthorizationCode            string `xml:"AuthorizationCode,omitempty" json:"AuthorizationCode,omitempty"`                       // 授权码，可用于换取公众号的接口调用凭据
	AuthorizationCodeExpiredTime string `xml:"AuthorizationCodeExpiredTime,omitempty" json:"AuthorizationCodeExpiredTime,omitempty"` // 授权码过期时间
	PreAuthCode                  string `xml:"PreAuthCode,omitempty" json:"PreAuthCode,omitempty"`                                   // 预授权码
}
