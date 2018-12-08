package pay

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/mch/core"
	"github.com/chanxuehong/wechat/util"
)

// UnifiedOrder 统一下单.
func UnifiedOrder(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/unifiedorder", req)
}

type UnifiedOrderRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选参数
	Body           string `xml:"body"`             // 商品或支付单简要描述
	OutTradeNo     string `xml:"out_trade_no"`     // 商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
	TotalFee       int64  `xml:"total_fee"`        // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string `xml:"spbill_create_ip"` // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	NotifyURL      string `xml:"notify_url"`       // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。
	TradeType      string `xml:"trade_type"`       // 取值如下：JSAPI，NATIVE，APP，详细说明见参数规定

	// 可选参数
	DeviceInfo string    `xml:"device_info"` // 终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
	NonceStr   string    `xml:"nonce_str"`   // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType   string    `xml:"sign_type"`   // 签名类型，默认为MD5，支持HMAC-SHA256和MD5。
	Detail     string    `xml:"detail"`      // 商品名称明细列表
	Attach     string    `xml:"attach"`      // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	FeeType    string    `xml:"fee_type"`    // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TimeStart  time.Time `xml:"time_start"`  // 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeExpire time.Time `xml:"time_expire"` // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则
	GoodsTag   string    `xml:"goods_tag"`   // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	ProductId  string    `xml:"product_id"`  // trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义。
	LimitPay   string    `xml:"limit_pay"`   // no_credit--指定不能使用信用卡支付
	OpenId     string    `xml:"openid"`      // rade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。
	SubOpenId  string    `xml:"sub_openid"`  // trade_type=JSAPI，此参数必传，用户在子商户appid下的唯一标识。openid和sub_openid可以选传其中之一，如果选择传sub_openid,则必须传sub_appid。
	SceneInfo  string    `xml:"scene_info"`  // 该字段用于上报支付的场景信息,针对H5支付有以下三种场景,请根据对应场景上报,H5支付不建议在APP端使用，针对场景1，2请接入APP支付，不然可能会出现兼容性问题
}

type UnifiedOrderResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选返回
	PrepayId  string `xml:"prepay_id"`  // 微信生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时
	TradeType string `xml:"trade_type"` // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，详细说明见参数规定

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	DeviceInfo string `xml:"device_info"` // 调用接口提交的终端设备号。
	CodeURL    string `xml:"code_url"`    // trade_type 为 NATIVE 时有返回，可将该参数值生成二维码展示出来进行扫码支付
	MWebURL    string `xml:"mweb_url"`    // trade_type 为 MWEB 时有返回
}

// UnifiedOrder2 统一下单.
func UnifiedOrder2(clt *core.Client, req *UnifiedOrderRequest) (resp *UnifiedOrderResponse, err error) {
	m1 := make(map[string]string, 24)
	m1["body"] = req.Body
	m1["out_trade_no"] = req.OutTradeNo
	m1["total_fee"] = strconv.FormatInt(req.TotalFee, 10)
	m1["spbill_create_ip"] = req.SpbillCreateIP
	m1["notify_url"] = req.NotifyURL
	m1["trade_type"] = req.TradeType
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
	if !req.TimeStart.IsZero() {
		m1["time_start"] = core.FormatTime(req.TimeStart)
	}
	if !req.TimeExpire.IsZero() {
		m1["time_expire"] = core.FormatTime(req.TimeExpire)
	}
	if req.GoodsTag != "" {
		m1["goods_tag"] = req.GoodsTag
	}
	if req.ProductId != "" {
		m1["product_id"] = req.ProductId
	}
	if req.LimitPay != "" {
		m1["limit_pay"] = req.LimitPay
	}
	if req.OpenId != "" {
		m1["openid"] = req.OpenId
	}
	if req.SubOpenId != "" {
		m1["sub_openid"] = req.SubOpenId
	}
	if req.SceneInfo != "" {
		m1["scene_info"] = req.SceneInfo
	}

	m2, err := UnifiedOrder(clt, m1)
	if err != nil {
		return nil, err
	}

	// 校验 trade_type
	respTradeType := m2["trade_type"]
	if respTradeType != req.TradeType {
		err = fmt.Errorf("trade_type mismatch, have: %s, want: %s", respTradeType, req.TradeType)
		return nil, err
	}

	resp = &UnifiedOrderResponse{
		PrepayId:   m2["prepay_id"],
		TradeType:  respTradeType,
		DeviceInfo: m2["device_info"],
		CodeURL:    m2["code_url"],
		MWebURL:    m2["mweb_url"],
	}
	return resp, nil
}
