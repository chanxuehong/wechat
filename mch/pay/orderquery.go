package pay

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"

	"github.com/chanxuehong/rand"
	"github.com/chanxuehong/util"
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 查询订单.
func OrderQuery(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/orderquery", req)
}

type OrderQueryRequest struct {
	TransactionId string // 微信的订单号，优先使用
	OutTradeNo    string // 商户系统内部的订单号，当没提供transaction_id时需要传这个。
	NonceStr      string // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
}

type OrderQueryResponse struct {
	AppId string // 微信开放平台审核通过的应用APPID
	MchId string // 微信支付分配的商户号

	OpenId         string // 用户在商户appid下的唯一标识
	TradeType      string // 调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
	TradeState     string // 交易状态
	BankType       string // 银行类型，采用字符串类型的银行标识
	TotalFee       int64  // 订单总金额，单位为分
	CashFee        int64  // 现金支付金额订单现金支付金额，详见支付金额
	TransactionId  string // 微信支付订单号
	OutTradeNo     string // 商户系统的订单号，与请求一致。
	TimeEnd        string // 订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则
	TradeStateDesc string // 对当前查询订单状态的描述和下一步操作的指引

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	DeviceInfo         string        // 微信支付分配的终端设备号
	IsSubscribe        *bool         // 用户是否关注公众账号
	SettlementTotalFee *int64        // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	FeeType            string        // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFeeType        string        // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CouponFee          *int64        // “代金券”金额<=订单金额，订单金额-“代金券”金额=现金支付金额，详见支付金额
	CouponCount        *int          // 代金券使用数量
	Coupons            []OrderCoupon // 代金券列表
	Attach             string        // 附加数据，原样返回
}

type OrderCoupon struct {
	CouponBatchId string // 代金券批次ID
	CouponType    string // 代金券类型, CASH--充值代金券, NO_CASH---非充值代金券
	CouponId      string // 代金券ID
	CouponFee     int64  // 单个代金券支付金额
}

func OrderQuery2(clt *core.Client, req *OrderQueryRequest) (resp *OrderQueryResponse, err error) {
	m1 := make(map[string]string, 8)
	m1["appid"] = clt.AppId()
	m1["mch_id"] = clt.MchId()
	if req.TransactionId != "" {
		m1["transaction_id"] = req.TransactionId
	}
	if req.OutTradeNo != "" {
		m1["out_trade_no"] = req.OutTradeNo
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = string(rand.NewHex())
	}
	m1["sign"] = core.Sign(m1, clt.ApiKey(), md5.New)

	m2, err := OrderQuery(clt, m1)
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

	resp = &OrderQueryResponse{
		AppId: m2["appid"],
		MchId: m2["mch_id"],

		OpenId:         m2["openid"],
		TradeType:      m2["trade_type"],
		TradeState:     m2["trade_state"],
		BankType:       m2["bank_type"],
		TransactionId:  m2["transaction_id"],
		OutTradeNo:     m2["out_trade_no"],
		TimeEnd:        m2["time_end"],
		TradeStateDesc: m2["trade_state_desc"],
	}

	var (
		n   int64
		id  int
		str string
	)
	if n, err = strconv.ParseInt(m2["total_fee"], 10, 64); err != nil {
		err = fmt.Errorf("parse total_fee:%q to int64 failed: %s", m2["total_fee"], err.Error())
		return
	} else {
		resp.TotalFee = n
	}
	if n, err = strconv.ParseInt(m2["cash_fee"], 10, 64); err != nil {
		err = fmt.Errorf("parse cash_fee:%q to int64 failed: %s", m2["cash_fee"], err.Error())
		return
	} else {
		resp.CashFee = n
	}

	resp.DeviceInfo = m2["device_info"]
	if str = m2["is_subscribe"]; str != "" {
		if str == "Y" || str == "y" {
			resp.IsSubscribe = util.Bool(true)
		} else {
			resp.IsSubscribe = util.Bool(false)
		}
	}
	if str = m2["settlement_total_fee"]; str != "" {
		if n, err = strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse settlement_total_fee:%q to int64 failed: %s", str, err.Error())
			return
		} else {
			resp.SettlementTotalFee = util.Int64(n)
		}
	}
	resp.FeeType = m2["fee_type"]
	resp.CashFeeType = m2["cash_fee_type"]
	if str = m2["coupon_fee"]; str != "" {
		if n, err = strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse coupon_fee:%q to int64 failed: %s", str, err.Error())
			return
		} else {
			resp.CouponFee = util.Int64(n)
		}
	}
	if str = m2["coupon_count"]; str != "" {
		if n, err = strconv.ParseInt(str, 10, 32); err != nil {
			err = fmt.Errorf("parse coupon_count:%q to int failed: %s", str, err.Error())
			return
		} else {
			resp.CouponCount = util.Int(int(n))
		}
	}
	resp.Attach = m2["attach"]

	if resp.CouponCount == nil {
		return // 没有代金券, 直接返回
	}
	CouponCount := *resp.CouponCount
	if CouponCount < 0 {
		err = fmt.Errorf("invalid coupon_count: %s", m2["coupon_count"])
		return
	}
	Coupons := make([]OrderCoupon, CouponCount)

	for k, v := range m2 {
		switch {
		case strings.HasPrefix(k, "coupon_batch_id_"):
			id, err = parseCouponId(k, k[len("coupon_batch_id_"):], CouponCount)
			if err != nil {
				return
			}
			Coupons[id].CouponBatchId = v
		case strings.HasPrefix(k, "coupon_type_"):
			id, err = parseCouponId(k, k[len("coupon_type_"):], CouponCount)
			if err != nil {
				return
			}
			Coupons[id].CouponType = v
		case strings.HasPrefix(k, "coupon_id_"):
			id, err = parseCouponId(k, k[len("coupon_id_"):], CouponCount)
			if err != nil {
				return
			}
			Coupons[id].CouponId = v
		case strings.HasPrefix(k, "coupon_fee_"):
			id, err = parseCouponId(k, k[len("coupon_fee_"):], CouponCount)
			if err != nil {
				return
			}
			n, err = strconv.ParseInt(v, 10, 64)
			if err != nil {
				return
			}
			Coupons[id].CouponFee = n
		}
	}
	return
}

// parseCouponId 解析 xxx_{id} 的 id 到一个整数
//  key:         xxx_{id}
//  id:          {id}
//  couponCount: id 应该在 [0, couponCount) 范围内
func parseCouponId(key, id string, couponCount int) (int, error) {
	if len(id) == 0 {
		return 0, fmt.Errorf("不正确的参数名 %s", key)
	}
	if id[0] == '$' {
		id = id[1:]
	}
	if len(id) == 0 {
		return 0, fmt.Errorf("不正确的参数名 %s", key)
	}
	n, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("不正确的参数名 %s", key)
	}
	if n >= int64(couponCount) {
		return 0, fmt.Errorf("参数 %s 的 id 超过了 coupon_count:%d 的范围", key, couponCount)
	}
	return int(n), nil
}
