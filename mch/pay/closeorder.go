package pay

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"

	"github.com/chanxuehong/wechat.v2/mch/core"
	"github.com/chanxuehong/wechat.v2/util"
)

// CloseOrder 关闭订单.
func CloseOrder(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/closeorder", req)
}

type CloseOrderRequest struct {
	OutTradeNo string `xml:"out_trade_no"` // 商户系统内部订单号
	NonceStr   string `xml:"nonce_str"`    // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType   string `xml:"sign_type"`    // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

// CloseOrder2 关闭订单.
func CloseOrder2(clt *core.Client, req *CloseOrderRequest) (err error) {
	m1 := make(map[string]string, 8)
	m1["appid"] = clt.AppId()
	m1["mch_id"] = clt.MchId()
	if req.OutTradeNo != "" {
		m1["out_trade_no"] = req.OutTradeNo
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = util.NonceStr()
	}

	// 签名
	switch req.SignType {
	case "":
		m1["sign"] = core.Sign2(m1, clt.ApiKey(), md5.New())
	case "MD5":
		m1["sign_type"] = "MD5"
		m1["sign"] = core.Sign2(m1, clt.ApiKey(), md5.New())
	case "HMAC-SHA256":
		m1["sign_type"] = "HMAC-SHA256"
		m1["sign"] = core.Sign2(m1, clt.ApiKey(), hmac.New(sha256.New, []byte(clt.ApiKey())))
	default:
		err = fmt.Errorf("unsupported sign_type: %s", req.SignType)
		return err
	}

	_, err = CloseOrder(clt, m1)
	return
}
