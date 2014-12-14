// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

const (
	MsgTypeSuiteTicket = "suite_ticket"
	MsgTypeChangeAuth  = "change_auth"
	MsgTypeCancelAuth  = "cancel_auth"
)

const ResponseSuccess = "success" // 应用提供商在收到推送消息后需要返回字符串success

type CommonHead struct {
	SuiteId   string `xml:"SuiteId" json:"SuiteId"`
	InfoType  string `xml:"InfoType" json:"InfoType"`
	TimeStamp int64  `xml:"TimeStamp" json:"TimeStamp"`
}

type mixedRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	SuiteTicket string `xml:"SuiteTicket" json:"SuiteTicket"`
	AuthCorpId  string `xml:"AuthCorpId" json:"AuthCorpId"`
}

func (this *mixedRequest) GetSuiteTicket() *SuiteTicket {
	return &SuiteTicket{
		CommonHead:  this.CommonHead,
		SuiteTicket: this.SuiteTicket,
	}
}
func (this *mixedRequest) GetChangeAuth() *ChangeAuth {
	return &ChangeAuth{
		CommonHead: this.CommonHead,
		AuthCorpId: this.AuthCorpId,
	}
}
func (this *mixedRequest) GetCancelAuth() *CancelAuth {
	return &CancelAuth{
		CommonHead: this.CommonHead,
		AuthCorpId: this.AuthCorpId,
	}
}

type SuiteTicket struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	SuiteTicket string `xml:"SuiteTicket" json:"SuiteTicket"`
}

type ChangeAuth struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	AuthCorpId string `xml:"AuthCorpId" json:"AuthCorpId"`
}

type CancelAuth struct {
	XMLName struct{} `xml:"xml" json:"-"`
	CommonHead

	AuthCorpId string `xml:"AuthCorpId" json:"AuthCorpId"`
}
