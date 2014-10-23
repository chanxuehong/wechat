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
)

// 统一支付接口 返回参数
type UnifiedOrderResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	RetCode string `xml:"return_code"          json:"return_code"`          // 必须, SUCCESS/FAIL; 此字段是通信标识，非交易标识，交易是否成功需要查看result_code 来判断
	RetMsg  string `xml:"return_msg,omitempty" json:"return_msg,omitempty"` // 可选, 返回信息，如非空，为错误原因: 签名失败, 参数格式校验错误

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
	TradeType string `xml:"trade_type"         json:"trade_type"`         // 必须, JSAPI、NATIVE、APP
	PrepayId  string `xml:"prepay_id"          json:"prepay_id"`          // 必须, 微信生成的预支付ID，用于后续接口调用中使用
	CodeURL   string `xml:"code_url,omitempty" json:"code_url,omitempty"` // 可选, trade_type 为NATIVE 是有返回，此参数可直接生成二维码展示出来进行扫码支付
}

// 检查 resp *UnifiedOrderResponse 的签名是否正确, 正确时返回 nil, 否则返回错误信息.
//  appKey: 商户支付密钥Key
func (resp *UnifiedOrderResponse) CheckSignature(appKey string) (err error) {
	if resp.RetCode != RET_CODE_SUCCESS {
		return
	}

	if len(resp.Signature) != md5.Size*2 {
		err = fmt.Errorf(`不正确的签名: %q, 长度不对, have: %d, want: %d`,
			resp.Signature, len(resp.Signature), md5.Size*2)
		return
	}

	Hash := md5.New()
	Signature := make([]byte, md5.Size*2)

	// 字典序
	// appid
	// code_url // resp.RetCode == RET_CODE_SUCCESS && resp.ResultCode == RESULT_CODE_SUCCESS
	// device_info
	// err_code
	// err_code_des
	// mch_id
	// nonce_str
	// prepay_id // resp.RetCode == RET_CODE_SUCCESS && resp.ResultCode == RESULT_CODE_SUCCESS
	// result_code
	// return_code
	// return_msg
	// trade_type // resp.RetCode == RET_CODE_SUCCESS && resp.ResultCode == RESULT_CODE_SUCCESS
	if len(resp.AppId) > 0 {
		Hash.Write([]byte("appid="))
		Hash.Write([]byte(resp.AppId))
		Hash.Write([]byte{'&'})
	}
	if len(resp.CodeURL) > 0 && resp.ResultCode == RESULT_CODE_SUCCESS {
		Hash.Write([]byte("code_url="))
		Hash.Write([]byte(resp.CodeURL))
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
	if len(resp.PrepayId) > 0 && resp.ResultCode == RESULT_CODE_SUCCESS {
		Hash.Write([]byte("prepay_id="))
		Hash.Write([]byte(resp.PrepayId))
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
	if len(resp.TradeType) > 0 && resp.ResultCode == RESULT_CODE_SUCCESS {
		Hash.Write([]byte("trade_type="))
		Hash.Write([]byte(resp.TradeType))
		Hash.Write([]byte{'&'})
	}

	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(Signature, Hash.Sum(nil))
	Signature = bytes.ToUpper(Signature)

	if subtle.ConstantTimeCompare(Signature, []byte(resp.Signature)) != 1 {
		err = fmt.Errorf("不正确的签名, \r\nhave: %q, \r\nwant: %q", Signature, resp.Signature)
		return
	}
	return
}
