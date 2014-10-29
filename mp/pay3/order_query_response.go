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
	"github.com/chanxuehong/wechat/util"
	"strconv"
	"time"
)

// 订单查询接口 返回参数
type OrderQueryResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	RetCode string `xml:"return_code"           json:"return_code"`          // 必须, SUCCESS/FAIL; 此字段是通信标识，非交易标识，交易是否成功需要查看trade_state 来判断
	RetMsg  string `xml:"return_msg,omitempty"  json:"return_msg,omitempty"` // 可选, 返回信息，如非空，为错误原因: 签名失败, 参数格式校验错误

	// 以下字段在 RetCode 为 SUCCESS 的时候有返回
	AppId       string `xml:"appid"                  json:"appid"`                  // 必须, 微信分配的公众账号ID
	MerchantId  string `xml:"mch_id"                 json:"mch_id"`                 // 必须, 微信支付分配的商户号
	NonceStr    string `xml:"nonce_str"              json:"nonce_str"`              // 必须, 随机字符串，不长于32 位
	Signature   string `xml:"sign"                   json:"sign"`                   // 必须, 签名
	ResultCode  string `xml:"result_code"            json:"result_code"`            // 必须, SUCCESS/FAIL
	ErrCode     string `xml:"err_code,omitempty"     json:"err_code,omitempty"`     // 可选, 错误代码
	ErrCodeDesc string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"` // 可选, 错误代码详细描述

	// 以下字段在 RetCode 和 ResultCode 都为 SUCCESS 的时候有返回

	// SUCCESS—支付成功
	// REFUND—转入退款
	// NOTPAY—未支付
	// CLOSED—已关闭
	// REVOKED—已撤销
	// USERPAYING--用户支付中
	// NOPAY--未支付(输入密码或确认支付超时)
	// PAYERROR--支付失败(其他原因，如银行返回失败)
	TradeState    string `xml:"trade_state"              json:"trade_state"`              // 必须
	DeviceInfo    string `xml:"device_info"              json:"device_info"`              // 可选, 微信支付分配的终端设备号
	OpenId        string `xml:"openid"                   json:"openid"`                   // 必须, 用户在商户appid 下的唯一标识
	IsSubscribe   string `xml:"is_subscribe"             json:"is_subscribe"`             // 必须, 用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	TradeType     string `xml:"trade_type"               json:"trade_type"`               // 必须, JSAPI、NATIVE、MICROPAY、APP
	BankType      string `xml:"bank_type"                json:"bank_type"`                // 必须, 银行类型，采用字符串类型的银行标识
	TotalFee      string `xml:"total_fee"                json:"total_fee"`                // 必须, 订单总金额，单位为分
	CouponFee     string `xml:"coupon_fee,omitempty"     json:"coupon_fee,omitempty"`     // 可选, 现金券支付金额<=订单总金额，订单总金额-现金券金额为现金支付金额
	FeeType       string `xml:"fee_type,omitempty"       json:"fee_type,omitempty"`       // 可选, 货币类型，符合ISO 4217 标准的三位字母代码，默认人民币：CNY
	TransactionId string `xml:"transaction_id,omitempty" json:"transaction_id,omitempty"` // 可选, 微信支付订单号
	OutTradeNo    string `xml:"out_trade_no,omitempty"   json:"out_trade_no,omitempty"`   // 可选, 商户系统的订单号，与请求一致。
	Attach        string `xml:"attach,omitempty"         json:"attach,omitempty"`         // 可选, 商家数据包，原样返回
	TimeEnd       string `xml:"time_end"                 json:"time_end"`                 // 必须, 支付完成时间， 格式为yyyyMMddhhmmss。时区为GMT+8 beijing。该时间取自微信支付服务器
}

// getter
func (this *OrderQueryResponse) GetTotalFee() (n int64, err error) {
	return strconv.ParseInt(this.TotalFee, 10, 64)
}
func (this *OrderQueryResponse) GetCouponFee() (n int64, err error) {
	return strconv.ParseInt(this.CouponFee, 10, 64)
}
func (this *OrderQueryResponse) GetTimeEnd() (time.Time, error) {
	return util.ParseTime(this.TimeEnd)
}

// 检查 resp *OrderQueryResponse 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 商户支付密钥Key
func (resp *OrderQueryResponse) CheckSignature(appKey string) (err error) {
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
	// attach
	// bank_type
	// coupon_fee
	// device_info
	// err_code
	// err_code_des
	// fee_type
	// is_subscribe
	// mch_id
	// nonce_str
	// openid
	// out_trade_no
	// result_code
	// return_code
	// return_msg
	// time_end
	// total_fee
	// trade_state
	// trade_type
	// transaction_id
	if len(resp.AppId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(resp.AppId))
		Hash.Write([]byte{'&'})
	}
	if len(resp.Attach) > 0 {
		Hash.Write([]byte("attach="))
		Hash.Write([]byte(resp.Attach))
		Hash.Write([]byte{'&'})
	}
	if len(resp.BankType) > 0 {
		Hash.Write([]byte("bank_type="))
		Hash.Write([]byte(resp.BankType))
		Hash.Write([]byte{'&'})
	}
	if len(resp.CouponFee) > 0 {
		Hash.Write([]byte("coupon_fee="))
		Hash.Write([]byte(resp.CouponFee))
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
	if len(resp.FeeType) > 0 {
		Hash.Write([]byte("fee_type="))
		Hash.Write([]byte(resp.FeeType))
		Hash.Write([]byte{'&'})
	}
	if len(resp.IsSubscribe) > 0 {
		Hash.Write([]byte("is_subscribe="))
		Hash.Write([]byte(resp.IsSubscribe))
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
	if len(resp.OpenId) > 0 {
		Hash.Write([]byte("openid="))
		Hash.Write([]byte(resp.OpenId))
		Hash.Write([]byte{'&'})
	}
	if len(resp.OutTradeNo) > 0 {
		Hash.Write([]byte("out_trade_no="))
		Hash.Write([]byte(resp.OutTradeNo))
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
	if len(resp.TimeEnd) > 0 {
		Hash.Write([]byte("time_end="))
		Hash.Write([]byte(resp.TimeEnd))
		Hash.Write([]byte{'&'})
	}
	if len(resp.TotalFee) > 0 {
		Hash.Write([]byte("total_fee="))
		Hash.Write([]byte(resp.TotalFee))
		Hash.Write([]byte{'&'})
	}
	if len(resp.TradeState) > 0 {
		Hash.Write([]byte("trade_state="))
		Hash.Write([]byte(resp.TradeState))
		Hash.Write([]byte{'&'})
	}
	if len(resp.TradeType) > 0 {
		Hash.Write([]byte("trade_type="))
		Hash.Write([]byte(resp.TradeType))
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
