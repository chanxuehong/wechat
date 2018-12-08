package pay

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/mch/core"
	"github.com/chanxuehong/wechat/util"
)

// MicroPay 提交刷卡支付.
func MicroPay(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/micropay", req)
}

type MicroPayRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选参数
	Body           string `xml:"body"`             // 商品或支付单简要描述
	OutTradeNo     string `xml:"out_trade_no"`     // 商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
	TotalFee       int64  `xml:"total_fee"`        // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string `xml:"spbill_create_ip"` // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	AuthCode       string `xml:"auth_code"`        // 扫码支付授权码，设备读取用户微信中的条码或者二维码信息

	// 可选参数
	DeviceInfo string `xml:"device_info"` // 终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
	NonceStr   string `xml:"nonce_str"`   // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType   string `xml:"sign_type"`   // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	Detail     string `xml:"detail"`      // 商品名称明细列表
	Attach     string `xml:"attach"`      // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType    string `xml:"fee_type"`    // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	GoodsTag   string `xml:"goods_tag"`   // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	LimitPay   string `xml:"limit_pay"`   // no_credit--指定不能使用信用卡支付
	SceneInfo  string `xml:"scene_info"`  // 场景信息
}

type MicroPayResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选返回
	OpenId        string    `xml:"openid"`         // 用户在商户appid下的唯一标识
	IsSubscribe   bool      `xml:"is_subscribe"`   // 用户是否关注公众账号
	TradeType     string    `xml:"trade_type"`     // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	BankType      string    `xml:"bank_type"`      // 银行类型，采用字符串类型的银行标识
	TotalFee      int64     `xml:"total_fee"`      // 订单总金额，单位为分
	CashFee       int64     `xml:"cash_fee"`       // 现金支付金额订单现金支付金额，详见支付金额
	TransactionId string    `xml:"transaction_id"` // 微信支付订单号
	OutTradeNo    string    `xml:"out_trade_no"`   // 商户系统的订单号，与请求一致。
	TimeEnd       time.Time `xml:"time_end"`       // 订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	DeviceInfo         string `xml:"device_info"`          // 微信支付分配的终端设备号
	SubOpenId          string `xml:"sub_openid"`           // 子商户appid下用户唯一标识，如需返回则请求时需要传sub_appid
	SubIsSubscribe     *bool  `xml:"sub_is_subscribe"`     // 用户是否关注子公众账号，仅在公众账号类型支付有效，取值范围：Y或N;Y-关注;N-未关注
	FeeType            string `xml:"fee_type"`             // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	SettlementTotalFee *int64 `xml:"settlement_total_fee"` // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	CouponFee          *int64 `xml:"coupon_fee"`           // 代金券金额
	CashFeeType        string `xml:"cash_fee_type"`        // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	Attach             string `xml:"attach"`               // 附加数据，原样返回
	PromotionDetail    string `xml:"promotion_detail"`     // 营销详情
}

// MicroPay2 提交刷卡支付.
func MicroPay2(clt *core.Client, req *MicroPayRequest) (resp *MicroPayResponse, err error) {
	m1 := make(map[string]string, 24)
	m1["body"] = req.Body
	m1["out_trade_no"] = req.OutTradeNo
	m1["total_fee"] = strconv.FormatInt(req.TotalFee, 10)
	m1["spbill_create_ip"] = req.SpbillCreateIP
	m1["auth_code"] = req.AuthCode
	if req.DeviceInfo != "" {
		m1["device_info"] = req.DeviceInfo
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = util.NonceStr()
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}
	if req.Detail != "" {
		m1["detail"] = req.Detail
	}
	if req.Attach != "" {
		m1["attach"] = req.Attach
	}
	if req.FeeType != "" {
		m1["fee_type"] = req.FeeType
	}
	if req.GoodsTag != "" {
		m1["goods_tag"] = req.GoodsTag
	}
	if req.LimitPay != "" {
		m1["limit_pay"] = req.LimitPay
	}
	if req.SceneInfo != "" {
		m1["scene_info"] = req.SceneInfo
	}

	m2, err := MicroPay(clt, m1)
	if err != nil {
		return nil, err
	}

	resp = &MicroPayResponse{
		OpenId:          m2["openid"],
		TradeType:       m2["trade_type"],
		BankType:        m2["bank_type"],
		TransactionId:   m2["transaction_id"],
		OutTradeNo:      m2["out_trade_no"],
		DeviceInfo:      m2["device_info"],
		SubOpenId:       m2["sub_openid"],
		FeeType:         m2["fee_type"],
		CashFeeType:     m2["cash_fee_type"],
		Attach:          m2["attach"],
		PromotionDetail: m2["promotion_detail"],
	}

	if str := m2["is_subscribe"]; str != "" {
		if str == "Y" || str == "y" {
			resp.IsSubscribe = true
		}
	}
	if str := m2["total_fee"]; str != "" {
		if n, err := strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse total_fee:%q to int64 failed: %s", str, err.Error())
			return nil, err
		} else {
			resp.TotalFee = n
		}
	}
	if str := m2["cash_fee"]; str != "" {
		if n, err := strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse cash_fee:%q to int64 failed: %s", str, err.Error())
			return nil, err
		} else {
			resp.CashFee = n
		}
	}
	if str := m2["time_end"]; str != "" {
		if t, err := core.ParseTime(str); err != nil {
			err = fmt.Errorf("parse time_end:%q to time.Time failed: %s", str, err.Error())
			return nil, err
		} else {
			resp.TimeEnd = t
		}
	}

	if str := m2["sub_is_subscribe"]; str != "" {
		if str == "Y" || str == "y" {
			resp.SubIsSubscribe = util.Bool(true)
		} else {
			resp.SubIsSubscribe = util.Bool(false)
		}
	}
	if str := m2["settlement_total_fee"]; str != "" {
		if n, err := strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse settlement_total_fee:%q to int64 failed: %s", str, err.Error())
			return nil, err
		} else {
			resp.SettlementTotalFee = util.Int64(n)
		}
	}
	if str := m2["coupon_fee"]; str != "" {
		if n, err := strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse coupon_fee:%q to int64 failed: %s", str, err.Error())
			return nil, err
		} else {
			resp.CouponFee = util.Int64(n)
		}
	}

	// 校验返回参数
	if req.OutTradeNo != resp.OutTradeNo {
		err = fmt.Errorf("out_trade_no mismatch, have: %s, want: %s", resp.OutTradeNo, req.OutTradeNo)
		return nil, err
	}
	if req.TotalFee != resp.TotalFee {
		err = fmt.Errorf("total_fee mismatch, have: %d, want: %d", resp.TotalFee, req.TotalFee)
		return nil, err
	}
	return resp, nil
}
