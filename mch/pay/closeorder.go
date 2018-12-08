package pay

import (
	"github.com/chanxuehong/wechat/mch/core"
	"github.com/chanxuehong/wechat/util"
)

// CloseOrder 关闭订单.
func CloseOrder(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/closeorder", req)
}

type CloseOrderRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选参数
	OutTradeNo string `xml:"out_trade_no"` // 商户系统内部订单号

	// 可选参数
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

// CloseOrder2 关闭订单.
func CloseOrder2(clt *core.Client, req *CloseOrderRequest) (err error) {
	m1 := make(map[string]string, 8)
	m1["out_trade_no"] = req.OutTradeNo
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = util.NonceStr()
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}

	_, err = CloseOrder(clt, m1)
	return
}
