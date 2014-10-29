// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package pay3

import (
	"bytes"
	"crypto/md5"
	"crypto/subtle"
	"encoding/hex"
	"fmt"
	"strconv"
)

// 申请退款 返回参数
type RefundResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	RetCode string `xml:"return_code"           json:"return_code"`          // 必须, SUCCESS/FAIL; 此字段是通信标识，非交易标识，交易是否成功需要查看trade_state 来判断
	RetMsg  string `xml:"return_msg,omitempty"  json:"return_msg,omitempty"` // 可选, 返回信息，如非空，为错误原因: 签名失败, 参数格式校验错误

	// 以下字段在 RetCode 为 SUCCESS 的时候有返回
	AppId       string `xml:"appid"                  json:"appid"`                   // 必须, 微信分配的公众账号ID
	MerchantId  string `xml:"mch_id"                   json:"mch_id"`                // 必须, 微信支付分配的商户号
	DeviceInfo  string `xml:"device_info,omitempty"    json:"device_info,omitempty"` // 可选, 微信支付分配的终端设备号
	NonceStr    string `xml:"nonce_str"              json:"nonce_str"`               // 必须, 随机字符串，不长于32 位
	Signature   string `xml:"sign"                   json:"sign"`                    // 必须, 签名
	ResultCode  string `xml:"result_code"            json:"result_code"`             // 必须, SUCCESS/FAIL; SUCCESS 退款申请接收成功，结果通过退款查询接口查询FAIL
	ErrCode     string `xml:"err_code,omitempty"     json:"err_code,omitempty"`      // 可选, 错误代码
	ErrCodeDesc string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"`  // 可选, 错误代码详细描述

	TransactionId   string `xml:"transaction_id"              json:"transaction_id"`              // 必须, 微信订单号
	OutTradeNo      string `xml:"out_trade_no"                json:"out_trade_no"`                // 必须, 商户系统内部的订单号
	OutRefundNo     string `xml:"out_refund_no"               json:"out_refund_no"`               // 必须, 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔
	RefundId        string `xml:"refund_id"                   json:"refund_id"`                   // 必须, 微信退款单号
	RefundChannel   string `xml:"refund_channel,omitempty"    json:"refund_channel,omitempty"`    // 可选, 退款渠道, ORIGINAL—原路退款，默认BALANCE—退回到余额
	RefundFee       string `xml:"refund_fee"                  json:"refund_fee"`                  // 必须, 退款总金额,单位为分,可以做部分退款
	CouponRefundFee string `xml:"coupon_refund_fee,omitempty" json:"coupon_refund_fee,omitempty"` // 可选, 现金券退款金额<= 退款金额，退款金额-现金券退款金额为现金
}

func (this *RefundResponse) GetCouponRefundFee() (n int64, err error) {
	return strconv.ParseInt(this.CouponRefundFee, 10, 64)
}
func (this *RefundResponse) GetRefundFee() (n int64, err error) {
	return strconv.ParseInt(this.RefundFee, 10, 64)
}

// 检查 resp *RefundResponse 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 商户支付密钥Key
func (resp *RefundResponse) CheckSignature(appKey string) (err error) {
	if resp.RetCode != RET_CODE_SUCCESS {
		return
	}

	if len(resp.Signature) != md5.Size*2 {
		err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
			resp.Signature, len(resp.Signature), md5.Size*2)
		return
	}

	Hash := md5.New()
	hashsum := make([]byte, md5.Size*2)

	// 字典序
	// appid
	// coupon_refund_fee
	// device_info
	// err_code
	// err_code_des
	// mch_id
	// nonce_str
	// out_refund_no
	// out_trade_no
	// refund_channel
	// refund_fee
	// refund_id
	// result_code
	// return_code
	// return_msg
	// transaction_id
	if len(resp.AppId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(resp.AppId))
		Hash.Write([]byte{'&'})
	}
	if len(resp.CouponRefundFee) > 0 {
		Hash.Write([]byte("coupon_refund_fee="))
		Hash.Write([]byte(resp.CouponRefundFee))
		Hash.Write([]byte{'&'})
	}
	if len(resp.DeviceInfo) > 0 {
		Hash.Write([]byte("device_info="))
		Hash.Write([]byte(resp.DeviceInfo))
		Hash.Write([]byte{'&'})
	}
	if len(resp.ErrCode) > 0 {
		Hash.Write([]byte("err_code="))
		Hash.Write([]byte(resp.ErrCode))
		Hash.Write([]byte{'&'})
	}
	if len(resp.ErrCodeDesc) > 0 {
		Hash.Write([]byte("err_code_des="))
		Hash.Write([]byte(resp.ErrCodeDesc))
		Hash.Write([]byte{'&'})
	}
	if len(resp.MerchantId) > 0 {
		Hash.Write([]byte("mch_id="))
		Hash.Write([]byte(resp.MerchantId))
		Hash.Write([]byte{'&'})
	}
	if len(resp.NonceStr) > 0 {
		Hash.Write([]byte("nonce_str="))
		Hash.Write([]byte(resp.NonceStr))
		Hash.Write([]byte{'&'})
	}
	if len(resp.OutRefundNo) > 0 {
		Hash.Write([]byte("out_refund_no="))
		Hash.Write([]byte(resp.OutRefundNo))
		Hash.Write([]byte{'&'})
	}
	if len(resp.OutTradeNo) > 0 {
		Hash.Write([]byte("out_trade_no="))
		Hash.Write([]byte(resp.OutTradeNo))
		Hash.Write([]byte{'&'})
	}
	if len(resp.RefundChannel) > 0 {
		Hash.Write([]byte("refund_channel="))
		Hash.Write([]byte(resp.RefundChannel))
		Hash.Write([]byte{'&'})
	}
	if len(resp.RefundFee) > 0 {
		Hash.Write([]byte("refund_fee="))
		Hash.Write([]byte(resp.RefundFee))
		Hash.Write([]byte{'&'})
	}
	if len(resp.RefundId) > 0 {
		Hash.Write([]byte("refund_id="))
		Hash.Write([]byte(resp.RefundId))
		Hash.Write([]byte{'&'})
	}
	if len(resp.ResultCode) > 0 {
		Hash.Write([]byte("result_code="))
		Hash.Write([]byte(resp.ResultCode))
		Hash.Write([]byte{'&'})
	}
	if len(resp.RetCode) > 0 {
		Hash.Write([]byte("return_code="))
		Hash.Write([]byte(resp.RetCode))
		Hash.Write([]byte{'&'})
	}
	if len(resp.RetMsg) > 0 {
		Hash.Write([]byte("return_msg="))
		Hash.Write([]byte(resp.RetMsg))
		Hash.Write([]byte{'&'})
	}
	if len(resp.TransactionId) > 0 {
		Hash.Write([]byte("transaction_id="))
		Hash.Write([]byte(resp.TransactionId))
		Hash.Write([]byte{'&'})
	}

	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(hashsum, Hash.Sum(nil))
	hashsum = bytes.ToUpper(hashsum)

	if subtle.ConstantTimeCompare(hashsum, []byte(resp.Signature)) != 1 {
		err = fmt.Errorf("签名不匹配, \r\nlocal: %q, \r\ninput: %q", hashsum, resp.Signature)
		return
	}
	return
}
