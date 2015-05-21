// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package suite

const (
	// 微信服务器推送过来的消息类型
	SuiteMsgTypeSuiteTicket = "suite_ticket" // 推送suite_ticket协议
	SuiteMsgTypeChangeAuth  = "change_auth"  // 变更授权的通知
	SuiteMsgTypeCancelAuth  = "cancel_auth"  // 取消授权的通知
)

type SuiteTicketMessage struct {
	XMLName struct{} `xml:"xml" json:"-"`

	SuiteId   string `xml:"SuiteId"   json:"SuiteId"`
	InfoType  string `xml:"InfoType"  json:"InfoType"`
	Timestamp int64  `xml:"TimeStamp" json:"TimeStamp"`

	SuiteTicket string `xml:"SuiteTicket" json:"SuiteTicket"`
}

func GetSuiteTicketMessage(msg *MixedSuiteMessage) *SuiteTicketMessage {
	return &SuiteTicketMessage{
		SuiteId:     msg.SuiteId,
		InfoType:    msg.InfoType,
		Timestamp:   msg.Timestamp,
		SuiteTicket: msg.SuiteTicket,
	}
}

type ChangeAuthMessage struct {
	XMLName struct{} `xml:"xml" json:"-"`

	SuiteId   string `xml:"SuiteId"   json:"SuiteId"`
	InfoType  string `xml:"InfoType"  json:"InfoType"`
	Timestamp int64  `xml:"TimeStamp" json:"TimeStamp"`

	AuthCorpId string `xml:"AuthCorpId"  json:"AuthCorpId"`
}

func GetChangeAuthMessage(msg *MixedSuiteMessage) *ChangeAuthMessage {
	return &ChangeAuthMessage{
		SuiteId:    msg.SuiteId,
		InfoType:   msg.InfoType,
		Timestamp:  msg.Timestamp,
		AuthCorpId: msg.AuthCorpId,
	}
}

type CancelAuthMessage struct {
	XMLName struct{} `xml:"xml" json:"-"`

	SuiteId   string `xml:"SuiteId"   json:"SuiteId"`
	InfoType  string `xml:"InfoType"  json:"InfoType"`
	Timestamp int64  `xml:"TimeStamp" json:"TimeStamp"`

	AuthCorpId string `xml:"AuthCorpId"  json:"AuthCorpId"`
}

func GetCancelAuthMessage(msg *MixedSuiteMessage) *CancelAuthMessage {
	return &CancelAuthMessage{
		SuiteId:    msg.SuiteId,
		InfoType:   msg.InfoType,
		Timestamp:  msg.Timestamp,
		AuthCorpId: msg.AuthCorpId,
	}
}
