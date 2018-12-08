package pay

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/mch/core"
	wechatutil "github.com/chanxuehong/wechat/util"
)

// OrderQuery 查询订单.
func OrderQuery(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/orderquery", req)
}

type OrderQueryRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 下面这些参数至少提供一个
	TransactionId string `xml:"transaction_id"` // 微信的订单号，优先使用
	OutTradeNo    string `xml:"out_trade_no"`   // 商户系统内部的订单号，当没提供transaction_id时需要传这个。

	// 可选参数
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

type OrderQueryResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选返回
	TradeState     string    `xml:"trade_state"`      // 交易状态
	TradeStateDesc string    `xml:"trade_state_desc"` // 对当前查询订单状态的描述和下一步操作的指引
	OpenId         string    `xml:"openid"`           // 用户在商户appid下的唯一标识
	TransactionId  string    `xml:"transaction_id"`   // 微信支付订单号
	OutTradeNo     string    `xml:"out_trade_no"`     // 商户系统的订单号，与请求一致。
	TradeType      string    `xml:"trade_type"`       // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	BankType       string    `xml:"bank_type"`        // 银行类型，采用字符串类型的银行标识
	TotalFee       int64     `xml:"total_fee"`        // 订单总金额，单位为分
	CashFee        int64     `xml:"cash_fee"`         // 现金支付金额订单现金支付金额，详见支付金额
	TimeEnd        time.Time `xml:"time_end"`         // 订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	DeviceInfo         string `xml:"device_info"`          // 微信支付分配的终端设备号
	IsSubscribe        *bool  `xml:"is_subscribe"`         // 用户是否关注公众账号
	SubOpenId          string `xml:"sub_openid"`           // 用户在子商户appid下的唯一标识
	SubIsSubscribe     *bool  `xml:"sub_is_subscribe"`     // 用户是否关注子公众账号
	SettlementTotalFee *int64 `xml:"settlement_total_fee"` // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	FeeType            string `xml:"fee_type"`             // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFeeType        string `xml:"cash_fee_type"`        // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	Detail             string `xml:"detail"`               // 商品详情
	Attach             string `xml:"attach"`               // 附加数据，原样返回
}

// OrderQuery2 查询订单.
//  NOTE: 该函数不支持 代金券 功能, 如果有 代金券 功能请使用 OrderQuery 函数.
func OrderQuery2(clt *core.Client, req *OrderQueryRequest) (resp *OrderQueryResponse, err error) {
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
		m1["nonce_str"] = wechatutil.NonceStr()
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}

	m2, err := OrderQuery(clt, m1)
	if err != nil {
		return nil, err
	}

	// 判断 trade_state
	tradeState := m2["trade_state"]
	if tradeState != "SUCCESS" {
		resp = &OrderQueryResponse{
			TradeState:     tradeState,
			TradeStateDesc: m2["trade_state_desc"],
			OutTradeNo:     m2["out_trade_no"],
			Attach:         m2["attach"],
		}
		return resp, nil
	}

	resp = &OrderQueryResponse{
		TradeState:     tradeState,
		TradeStateDesc: m2["trade_state_desc"],
		OpenId:         m2["openid"],
		TransactionId:  m2["transaction_id"],
		OutTradeNo:     m2["out_trade_no"],
		TradeType:      m2["trade_type"],
		BankType:       m2["bank_type"],
		DeviceInfo:     m2["device_info"],
		SubOpenId:      m2["sub_openid"],
		FeeType:        m2["fee_type"],
		CashFeeType:    m2["cash_fee_type"],
		Detail:         m2["detail"],
		Attach:         m2["attach"],
	}

	// 校验返回参数
	if req.TransactionId != "" && resp.TransactionId != "" && req.TransactionId != resp.TransactionId {
		err = fmt.Errorf("transaction_id mismatch, have: %s, want: %s", resp.TransactionId, req.TransactionId)
		return nil, err
	}
	if req.OutTradeNo != "" && resp.OutTradeNo != "" && req.OutTradeNo != resp.OutTradeNo {
		err = fmt.Errorf("out_trade_no mismatch, have: %s, want: %s", resp.OutTradeNo, req.OutTradeNo)
		return nil, err
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

	if str := m2["is_subscribe"]; str != "" {
		if str == "Y" || str == "y" {
			resp.IsSubscribe = wechatutil.Bool(true)
		} else {
			resp.IsSubscribe = wechatutil.Bool(false)
		}
	}
	if str := m2["sub_is_subscribe"]; str != "" {
		if str == "Y" || str == "y" {
			resp.SubIsSubscribe = wechatutil.Bool(true)
		} else {
			resp.SubIsSubscribe = wechatutil.Bool(false)
		}
	}
	if str := m2["settlement_total_fee"]; str != "" {
		if n, err := strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse settlement_total_fee:%q to int64 failed: %s", str, err.Error())
			return nil, err
		} else {
			resp.SettlementTotalFee = wechatutil.Int64(n)
		}
	}
	return resp, nil
}
