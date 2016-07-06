package pay

import (
	"crypto/md5"
	"strconv"

	"github.com/chanxuehong/rand"
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 统一下单.
func UnifiedOrder(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/unifiedorder", req)
}

type UnifiedOrderRequest struct {
	DeviceInfo     string // 终端设备号(门店号或收银设备ID)，注意：PC网页或公众号内支付请传"WEB"
	NonceStr       string // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	Body           string // 商品或支付单简要描述
	Detail         string // 商品名称明细列表
	Attach         string // 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据
	OutTradeNo     string // 商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
	FeeType        string // 符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	TotalFee       int64  // 订单总金额，单位为分，详见支付金额
	SpbillCreateIP string // APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP。
	TimeStart      string // 订单生成时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TimeExpire     string // 订单失效时间，格式为yyyyMMddHHmmss，如2009年12月27日9点10分10秒表示为20091227091010。其他详见时间规则
	GoodsTag       string // 商品标记，代金券或立减优惠功能的参数，说明详见代金券或立减优惠
	NotifyURL      string // 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数。
	TradeType      string // 取值如下：JSAPI，NATIVE，APP，详细说明见参数规定
	ProductId      string // trade_type=NATIVE，此参数必传。此id为二维码中包含的商品ID，商户自行定义。
	LimitPay       string // no_credit--指定不能使用信用卡支付
	OpenId         string // rade_type=JSAPI，此参数必传，用户在商户appid下的唯一标识。
}

type UnifiedOrderResponse struct {
	AppId string // 微信开放平台审核通过的应用APPID
	MchId string // 微信支付分配的商户号

	TradeType string // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，详细说明见参数规定
	PrepayId  string // 微信生成的预支付回话标识，用于后续接口调用中使用，该值有效期为2小时

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	DeviceInfo string // 调用接口提交的终端设备号。
	CodeURL    string // trade_type 为 NATIVE 时有返回，可将该参数值生成二维码展示出来进行扫码支付
	MWebURL    string // trade_type 为 MWEB 时有返回
}

func UnifiedOrder2(clt *core.Client, req *UnifiedOrderRequest) (resp *UnifiedOrderResponse, err error) {
	m1 := make(map[string]string, 20)
	m1["appid"] = clt.AppId()
	m1["mch_id"] = clt.MchId()
	if req.DeviceInfo != "" {
		m1["device_info"] = req.DeviceInfo
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = string(rand.NewHex())
	}
	m1["body"] = req.Body
	if req.Detail != "" {
		m1["detail"] = req.Detail
	}
	if req.Attach != "" {
		m1["attach"] = req.Attach
	}
	m1["out_trade_no"] = req.OutTradeNo
	if req.FeeType != "" {
		m1["fee_type"] = req.FeeType
	}
	m1["total_fee"] = strconv.FormatInt(req.TotalFee, 10)
	m1["spbill_create_ip"] = req.SpbillCreateIP
	if req.TimeStart != "" {
		m1["time_start"] = req.TimeStart
	}
	if req.TimeExpire != "" {
		m1["time_expire"] = req.TimeExpire
	}
	if req.GoodsTag != "" {
		m1["goods_tag"] = req.GoodsTag
	}
	m1["notify_url"] = req.NotifyURL
	m1["trade_type"] = req.TradeType
	if req.ProductId != "" {
		m1["product_id"] = req.ProductId
	}
	if req.LimitPay != "" {
		m1["limit_pay"] = req.LimitPay
	}
	if req.OpenId != "" {
		m1["openid"] = req.OpenId
	}
	m1["sign"] = core.Sign(m1, clt.ApiKey(), md5.New)

	m2, err := UnifiedOrder(clt, m1)
	if err != nil {
		return
	}

	// 判断业务状态
	resultCode, ok := m2["result_code"]
	if !ok {
		err = core.ErrNotFoundResultCode
		return
	}
	if resultCode != core.ResultCodeSuccess {
		err = &core.BizError{
			ResultCode:  resultCode,
			ErrCode:     m2["err_code"],
			ErrCodeDesc: m2["err_code_des"],
		}
		return
	}

	resp = &UnifiedOrderResponse{
		AppId: m2["appid"],
		MchId: m2["mch_id"],

		TradeType:  m2["trade_type"],
		PrepayId:   m2["prepay_id"],
		DeviceInfo: m2["device_info"],
		CodeURL:    m2["code_url"],
		MWebURL:    m2["mweb_url"],
	}
	return
}
