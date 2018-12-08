package pay

import (
	"github.com/chanxuehong/wechat/mch/core"
	"github.com/chanxuehong/wechat/util"
)

// Reverse 撤销订单.
//  NOTE: 请求需要双向证书.
func Reverse(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/secapi/pay/reverse", req)
}

type ReverseRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选参数，二选一
	TransactionId string `xml:"transaction_id"` // 微信的订单号，优先使用
	OutTradeNo    string `xml:"out_trade_no"`   // 商户系统内部订单号

	// 可选参数
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

type ReverseResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选返回
	Recall bool `xml:"recall"` // 是否需要继续调用撤销
}

// Reverse2 撤销订单.
//  NOTE: 请求需要双向证书.
func Reverse2(clt *core.Client, req *ReverseRequest) (resp *ReverseResponse, err error) {
	m1 := make(map[string]string, 8)
	if req.TransactionId != "" {
		m1["transaction_id"] = req.TransactionId
	}
	if req.OutTradeNo != "" {
		m1["out_trade_no"] = req.OutTradeNo
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = util.NonceStr()
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}

	m2, err := Reverse(clt, m1)
	if err != nil {
		return nil, err
	}

	resp = &ReverseResponse{}
	if recall := m2["recall"]; recall == "Y" || recall == "y" {
		resp.Recall = true
	}
	return resp, nil
}
