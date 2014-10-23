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

// 支付完成后，微信会把相关支付和用户信息发送到该 notify URL，商户需要接收处理信息。
// 对后台通知交互时，如果微信收到商户的应答不是成功或超时，微信认为通知失败，微
// 信会通过一定的策略（如30 分钟共8 次）定期重新发起通知，尽可能提高通知的成功率，
// 但微信不保证通知最终能成功。
// 由于存在重新发送后台通知的情况，因此同样的通知可能会多次发送给商户系统。商户
// 系统必须能够正确处理重复的通知。
//
// 这是通知的 post 参数的数据结构
type OrderNotifyPostData struct {
	XMLName struct{} `xml:"xml" json:"-"`

	RetCode string `xml:"return_code"           json:"return_code"`          // 必须, SUCCESS/FAIL; 此字段是通信标识，非交易标识，交易是否成功需要查看result_code 来判断
	RetMsg  string `xml:"return_msg,omitempty"  json:"return_msg,omitempty"` // 可选, 返回信息，如非空，为错误原因: 签名失败, 参数格式校验错误

	// 以下字段在 RetCode 为 SUCCESS 的时候有返回
	AppId       string `xml:"appid"                  json:"appid"`                  // 必须, 微信分配的公众账号ID
	MerchantId  string `xml:"mch_id"                 json:"mch_id"`                 // 必须, 微信支付分配的商户号
	DeviceInfo  string `xml:"device_info,omitempty"  json:"device_info,omitempty"`  // 可选, 微信支付分配的终端设备号
	NonceStr    string `xml:"nonce_str"              json:"nonce_str"`              // 必须, 随机字符串，不长于32 位
	Signature   string `xml:"sign"                   json:"sign"`                   // 必须, 签名
	ResultCode  string `xml:"result_code"            json:"result_code"`            // 必须, SUCCESS/FAIL
	ErrCode     string `xml:"err_code,omitempty"     json:"err_code,omitempty"`     // 可选, 错误代码
	ErrCodeDesc string `xml:"err_code_des,omitempty" json:"err_code_des,omitempty"` // 可选, 错误代码详细描述

	// 以下字段在 RetCode 和 ResultCode 都为 SUCCESS 的时候有返回
	OpenId        string `xml:"openid"             json:"openid"`               // 必须, 用户在商户appid 下的唯一标识
	IsSubscribe   string `xml:"is_subscribe"       json:"is_subscribe"`         // 必须, 用户是否关注公众账号，Y-关注，N-未关注，仅在公众账号类型支付有效
	TradeType     string `xml:"trade_type"         json:"trade_type"`           // 必须, JSAPI、NATIVE、MICROPAY、APP
	BankType      string `xml:"bank_type"          json:"bank_type"`            // 必须, 银行类型，采用字符串类型的银行标识
	TotalFee      *int   `xml:"total_fee"          json:"total_fee,omitempty"`  // 必须, 订单总金额，单位为分
	CouponFee     *int   `xml:"coupon_fee"         json:"coupon_fee,omitempty"` // 可选, 现金券支付金额<=订单总金额，订单总金额-现金券金额为现金支付金额
	FeeType       string `xml:"fee_type,omitempty" json:"fee_type,omitempty"`   // 可选, 货币类型，符合ISO 4217 标准的三位字母代码，默认人民币：CNY
	TransactionId string `xml:"transaction_id"     json:"transaction_id"`       // 必须, 微信支付订单号
	OutTradeNo    string `xml:"out_trade_no"       json:"out_trade_no"`         // 必须, 商户系统的订单号，与请求一致。
	Attach        string `xml:"attach,omitempty"   json:"attach,omitempty"`     // 可选, 商家数据包，原样返回
	TimeEnd       string `xml:"time_end"           json:"time_end"`             // 必须, 支付完成时间， 格式为yyyyMMddhhmmss。时区为GMT+8 beijing。该时间取自微信支付服务器
}

func (this *OrderNotifyPostData) GetTimeEnd() (time.Time, error) {
	return util.ParseTime(this.TimeEnd)
}

// 检查 data *OrderNotifyPostData 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 商户支付密钥Key
func (data *OrderNotifyPostData) CheckSignature(appKey string) (err error) {
	if data.RetCode != RET_CODE_SUCCESS {
		return
	}

	if len(data.Signature) != md5.Size*2 {
		err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
			data.Signature, len(data.Signature), md5.Size*2)
		return
	}

	Hash := md5.New()
	Signature := make([]byte, md5.Size*2)

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
	// sign
	// time_end
	// total_fee
	// trade_type
	// transaction_id
	if len(data.AppId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(data.AppId))
		Hash.Write([]byte{'&'})
	}
	if len(data.Attach) > 0 {
		Hash.Write([]byte("attach="))
		Hash.Write([]byte(data.Attach))
		Hash.Write([]byte{'&'})
	}
	if len(data.BankType) > 0 {
		Hash.Write([]byte("bank_type="))
		Hash.Write([]byte(data.BankType))
		Hash.Write([]byte{'&'})
	}
	if data.CouponFee != nil {
		Hash.Write([]byte("coupon_fee="))
		Hash.Write([]byte(strconv.FormatInt(int64(*data.CouponFee), 10)))
		Hash.Write([]byte{'&'})
	}
	if len(data.DeviceInfo) > 0 {
		Hash.Write([]byte("device_info="))
		Hash.Write([]byte(data.DeviceInfo))
		Hash.Write([]byte{'&'})
	}
	if len(data.ErrCode) > 0 {
		Hash.Write([]byte("err_code="))
		Hash.Write([]byte(data.ErrCode))
		Hash.Write([]byte{'&'})
	}
	if len(data.ErrCodeDesc) > 0 {
		Hash.Write([]byte("err_code_des="))
		Hash.Write([]byte(data.ErrCodeDesc))
		Hash.Write([]byte{'&'})
	}
	if len(data.FeeType) > 0 {
		Hash.Write([]byte("fee_type="))
		Hash.Write([]byte(data.FeeType))
		Hash.Write([]byte{'&'})
	}
	if len(data.IsSubscribe) > 0 {
		Hash.Write([]byte("is_subscribe="))
		Hash.Write([]byte(data.IsSubscribe))
		Hash.Write([]byte{'&'})
	}
	if len(data.MerchantId) > 0 {
		Hash.Write([]byte("mch_id="))
		Hash.Write([]byte(data.MerchantId))
		Hash.Write([]byte{'&'})
	}
	if len(data.NonceStr) > 0 {
		Hash.Write([]byte("nonce_str="))
		Hash.Write([]byte(data.NonceStr))
		Hash.Write([]byte{'&'})
	}
	if len(data.OpenId) > 0 {
		Hash.Write([]byte("openid="))
		Hash.Write([]byte(data.OpenId))
		Hash.Write([]byte{'&'})
	}
	if len(data.OutTradeNo) > 0 {
		Hash.Write([]byte("out_trade_no="))
		Hash.Write([]byte(data.OutTradeNo))
		Hash.Write([]byte{'&'})
	}
	if len(data.ResultCode) > 0 {
		Hash.Write([]byte("result_code="))
		Hash.Write([]byte(data.ResultCode))
		Hash.Write([]byte{'&'})
	}
	if len(data.RetCode) > 0 {
		Hash.Write([]byte("return_code="))
		Hash.Write([]byte(data.RetCode))
		Hash.Write([]byte{'&'})
	}
	if len(data.RetMsg) > 0 {
		Hash.Write([]byte("return_msg="))
		Hash.Write([]byte(data.RetMsg))
		Hash.Write([]byte{'&'})
	}
	if len(data.Signature) > 0 {
		Hash.Write([]byte("sign="))
		Hash.Write([]byte(data.Signature))
		Hash.Write([]byte{'&'})
	}
	if len(data.TimeEnd) > 0 {
		Hash.Write([]byte("time_end="))
		Hash.Write([]byte(data.TimeEnd))
		Hash.Write([]byte{'&'})
	}
	if data.TotalFee != nil {
		Hash.Write([]byte("total_fee="))
		Hash.Write([]byte(strconv.FormatInt(int64(*data.TotalFee), 10)))
		Hash.Write([]byte{'&'})
	}
	if len(data.TradeType) > 0 {
		Hash.Write([]byte("trade_type="))
		Hash.Write([]byte(data.TradeType))
		Hash.Write([]byte{'&'})
	}
	if len(data.TransactionId) > 0 {
		Hash.Write([]byte("transaction_id="))
		Hash.Write([]byte(data.TransactionId))
		Hash.Write([]byte{'&'})
	}

	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(Signature, Hash.Sum(nil))
	Signature = bytes.ToUpper(Signature)

	if subtle.ConstantTimeCompare(Signature, []byte(data.Signature)) != 1 {
		err = fmt.Errorf("不正确的签名, \r\nhave: %q, \r\nwant: %q", Signature, data.Signature)
		return
	}
	return
}
