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

// 统一支付接口 请求参数
type UnifiedOrderRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	AppId      string `xml:"appid"                 json:"appid"`                 // 必须, 微信分配的公众账号ID
	MerchantId string `xml:"mch_id"                json:"mch_id"`                // 必须, 微信支付分配的商户号
	DeviceInfo string `xml:"device_info,omitempty" json:"device_info,omitempty"` // 可选, 微信支付分配的终端设备号

	TradeType string `xml:"trade_type" json:"trade_type"` // 必须, JSAPI、NATIVE、APP
	NotifyURL string `xml:"notify_url" json:"notify_url"` // 必须, 接收微信支付成功通知

	OpenId       string `xml:"openid,omitempty"      json:"openid,omitempty"`      // 可选, 用户在商户appid 下的唯一标识，trade_type 为JSAPI时，此参数必传
	BillCreateIP string `xml:"spbill_create_ip"      json:"spbill_create_ip"`      // 必须, 订单生成的机器IP
	OutTradeNo   string `xml:"out_trade_no"          json:"out_trade_no"`          // 必须, 商户系统内部的订单号,32个字符内、可包含字母,确保在商户系统唯一
	ProductId    string `xml:"product_id,omitempty"  json:"product_id,omitempty"`  // 可选, 只在trade_type 为NATIVE时需要填写。此id 为二维码中包含的商品ID，商户自行维护。
	ProductTag   string `xml:"goods_tag,omitempty"   json:"goods_tag,omitempty"`   // 可选, 商品标记，该字段不能随便填，不使用请填空
	Body         string `xml:"body"                  json:"body"`                  // 必须, 商品描述
	Attach       string `xml:"attach,omitempty"      json:"attach,omitempty"`      // 可选, 附加数据，原样返回
	TotalFee     int    `xml:"total_fee"             json:"total_fee"`             // 必须, 订单总金额，单位为分，不能带小数点
	TimeStart    string `xml:"time_start,omitempty"  json:"time_start,omitempty"`  // 可选, 订单生成时间, see github.com/chanxuehong/wechat/util/FormatTime
	TimeExpire   string `xml:"time_expire,omitempty" json:"time_expire,omitempty"` // 可选, 订单失效时间, see github.com/chanxuehong/wechat/util/FormatTime

	NonceStr  string `xml:"nonce_str" json:"nonce_str"` // 必须, 随机字符串，不长于32 位
	Signature string `xml:"sign"      json:"sign"`      // 必须, 签名
}

// 设置签名字段.
//  appKey: 商户支付密钥Key
//
//  NOTE: 要求在 req *UnifiedOrderRequest 其他字段设置完毕后才能调用这个函数, 否则签名就不正确.
func (req *UnifiedOrderRequest) SetSignature(appKey string) (err error) {
	Hash := md5.New()
	Signature := make([]byte, md5.Size*2)

	// 字典序
	// appid
	// attach
	// body
	// device_info
	// goods_tag
	// mch_id
	// nonce_str
	// notify_url
	// openid
	// out_trade_no
	// product_id
	// spbill_create_ip
	// time_expire
	// time_start
	// total_fee
	// trade_type
	hasWrite := false
	if len(req.AppId) > 0 {
		hasWrite = true // 第一个不用判断

		Hash.Write([]byte("appid="))
		Hash.Write([]byte(req.AppId))
	}
	if len(req.Attach) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("attach="))
		Hash.Write([]byte(req.Attach))
	}
	if len(req.Body) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("body="))
		Hash.Write([]byte(req.Body))
	}
	if len(req.DeviceInfo) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("device_info="))
		Hash.Write([]byte(req.DeviceInfo))
	}
	if len(req.ProductTag) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("goods_tag="))
		Hash.Write([]byte(req.ProductTag))
	}
	if len(req.MerchantId) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("mch_id="))
		Hash.Write([]byte(req.MerchantId))
	}
	if len(req.NonceStr) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("nonce_str="))
		Hash.Write([]byte(req.NonceStr))
	}
	if len(req.NotifyURL) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("notify_url="))
		Hash.Write([]byte(req.NotifyURL))
	}
	if len(req.OpenId) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("openid="))
		Hash.Write([]byte(req.OpenId))
	}
	if len(req.OutTradeNo) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("out_trade_no="))
		Hash.Write([]byte(req.OutTradeNo))
	}
	if len(req.ProductId) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("product_id="))
		Hash.Write([]byte(req.ProductId))
	}
	if len(req.BillCreateIP) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("spbill_create_ip="))
		Hash.Write([]byte(req.BillCreateIP))
	}
	if len(req.TimeExpire) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("time_expire="))
		Hash.Write([]byte(req.TimeExpire))
	}
	if len(req.TimeStart) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("time_start="))
		Hash.Write([]byte(req.TimeStart))
	}
	Hash.Write([]byte("total_fee="))
	Hash.Write([]byte(strconv.FormatInt(int64(req.TotalFee), 10)))
	if len(req.TradeType) > 0 {
		if !hasWrite {
			hasWrite = true
		} else {
			Hash.Write([]byte{'&'})
		}
		Hash.Write([]byte("trade_type="))
		Hash.Write([]byte(req.TradeType))
	}

	if hasWrite {
		Hash.Write([]byte{'&'})
	}
	Hash.Write([]byte("key="))
	Hash.Write([]byte(appKey))

	hex.Encode(Signature, Hash.Sum(nil))
	Signature = bytes.ToUpper(Signature)

	req.Signature = string(Signature)
	return
}
