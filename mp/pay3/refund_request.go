// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

// 退款申请 请求参数
type RefundRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId         string `xml:"appid"                    json:"appid"`                    // 必须, 微信分配的公众账号ID
	MerchantId    string `xml:"mch_id"                   json:"mch_id"`                   // 必须, 微信支付分配的商户号
	DeviceInfo    string `xml:"device_info,omitempty"    json:"device_info,omitempty"`    // 可选, 微信支付分配的终端设备号
	NonceStr      string `xml:"nonce_str"                json:"nonce_str"`                // 必须, 随机字符串，不长于32 位
	Signature     string `xml:"sign"                     json:"sign"`                     // 必须, 签名
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` // 可选, 微信的订单号，优先使用
	OutTradeNo    string `xml:"out_trade_no,omitempty"   json:"out_trade_no,omitempty"`   // 必须, 商户系统内部的订单号,transaction_id、out_trade_no 二选一，如果同时存在优先级: transaction_id > out_trade_no
	OutRefundNo   string `xml:"out_refund_no"            json:"out_refund_no"`            // 必须, 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔
	TotalFee      *int   `xml:"total_fee"                json:"total_fee,omitempty"`      // 必须, 订单总金额，单位为分
	RefundFee     *int   `xml:"refund_fee"               json:"refund_fee,omitempty"`     // 必须, 退款总金额，单位为分,可以做部分退款
	OpUserId      string `xml:"op_user_id"               json:"op_user_id"`               // 必须, 操作员帐号, 默认为商户号
}

// getter
func (this *RefundRequest) GetTotalFee() (n int, ok bool) {
	if this.TotalFee != nil {
		ok = true
		n = *this.TotalFee
		return
	}
	return
}
func (this *RefundRequest) GetRefundFee() (n int, ok bool) {
	if this.RefundFee != nil {
		ok = true
		n = *this.RefundFee
		return
	}
	return
}

// setter
func (this *RefundRequest) SetTotalFee(n int) {
	this.TotalFee = &n
}
func (this *RefundRequest) SetRefundFee(n int) {
	this.RefundFee = &n
}

// 设置签名字段.
//  appKey: 商户支付密钥Key
//
//  NOTE: 要求在 req *RefundRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req *RefundRequest) SetSignature(appKey string) (err error) {
	Hash := md5.New()
	Signature := make([]byte, md5.Size*2)

	// 字典序
	// appid
	// device_info
	// mch_id
	// nonce_str
	// op_user_id
	// out_refund_no
	// out_trade_no
	// refund_fee
	// total_fee
	// transaction_id
	if len(req.AppId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(req.AppId))
		Hash.Write([]byte{'&'})
	}
	if len(req.DeviceInfo) > 0 {
		Hash.Write([]byte("device_info="))
		Hash.Write([]byte(req.DeviceInfo))
		Hash.Write([]byte{'&'})
	}
	if len(req.MerchantId) > 0 {
		Hash.Write([]byte("mch_id="))
		Hash.Write([]byte(req.MerchantId))
		Hash.Write([]byte{'&'})
	}
	if len(req.NonceStr) > 0 {
		Hash.Write([]byte("nonce_str="))
		Hash.Write([]byte(req.NonceStr))
		Hash.Write([]byte{'&'})
	}
	if len(req.OpUserId) > 0 {
		Hash.Write([]byte("op_user_id="))
		Hash.Write([]byte(req.OpUserId))
		Hash.Write([]byte{'&'})
	}
	if len(req.OutRefundNo) > 0 {
		Hash.Write([]byte("out_refund_no="))
		Hash.Write([]byte(req.OutRefundNo))
		Hash.Write([]byte{'&'})
	}
	if len(req.OutTradeNo) > 0 {
		Hash.Write([]byte("out_trade_no="))
		Hash.Write([]byte(req.OutTradeNo))
		Hash.Write([]byte{'&'})
	}
	if req.RefundFee != nil {
		Hash.Write([]byte("refund_fee="))
		Hash.Write([]byte(strconv.FormatInt(int64(*req.RefundFee), 10)))
		Hash.Write([]byte{'&'})
	}
	if req.TotalFee != nil {
		Hash.Write([]byte("total_fee="))
		Hash.Write([]byte(strconv.FormatInt(int64(*req.TotalFee), 10)))
		Hash.Write([]byte{'&'})
	}
	if len(req.TransactionId) > 0 {
		Hash.Write([]byte("transaction_id="))
		Hash.Write([]byte(req.TransactionId))
		Hash.Write([]byte{'&'})
	}

	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(Signature, Hash.Sum(nil))
	Signature = bytes.ToUpper(Signature)

	req.Signature = string(Signature)
	return
}
