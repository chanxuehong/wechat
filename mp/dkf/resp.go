package dkf

import (
	"github.com/chanxuehong/wechat/mp/core"
)

const (
	MsgTypeTransferCustomerService core.MsgType = "transfer_customer_service" // 将消息转发到多客服
)

// 将消息转发到多客服消息
type TransferToCustomerService struct {
	XMLName struct{} `xml:"xml" json:"-"`
	core.MsgHeader
	TransInfo *TransInfo `xml:"TransInfo,omitempty" json:"TransInfo,omitempty"`
}

type TransInfo struct {
	KfAccount string `xml:"KfAccount" json:"KfAccount"`
}

// 如果不指定客服则 kfAccount 留空.
func NewTransferToCustomerService(to, from string, timestamp int64, kfAccount string) (msg *TransferToCustomerService) {
	msg = &TransferToCustomerService{
		MsgHeader: core.MsgHeader{
			ToUserName:   to,
			FromUserName: from,
			CreateTime:   timestamp,
			MsgType:      MsgTypeTransferCustomerService,
		},
	}
	if kfAccount != "" {
		msg.TransInfo = &TransInfo{
			KfAccount: kfAccount,
		}
	}
	return
}
