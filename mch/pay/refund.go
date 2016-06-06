package pay

import (
	"crypto/md5"
	"fmt"
	"strconv"

	"github.com/chanxuehong/rand"
	"github.com/chanxuehong/util"
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 申请退款.
//  NOTE: 请求需要双向证书.
func Refund(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/secapi/pay/refund", req)
}

type RefundRequest struct {
	DeviceInfo    string // 终端设备号
	NonceStr      string // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	TransactionId string // 微信生成的订单号，在支付通知中有返回
	OutTradeNo    string // 商户侧传给微信的订单号
	OutRefundNo   string // 商户系统内部的退款单号，商户系统内部唯一，同一退款单号多次请求只退一笔
	TotalFee      int64  // 订单总金额，单位为分，只能为整数，详见支付金额
	RefundFee     int64  // 退款总金额，订单总金额，单位为分，只能为整数，详见支付金额
	RefundFeeType string // 货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	OperUserId    string // 操作员帐号, 默认为商户号
}

type RefundResponse struct {
	AppId string // 微信分配的公众账号ID
	MchId string // 微信支付分配的商户号

	TransactionId string // 微信订单号
	OutTradeNo    string // 商户系统内部的订单号
	OutRefundNo   string // 商户退款单号
	RefundId      string // 微信退款单号
	RefundFee     int64  // 退款总金额,单位为分,可以做部分退款
	TotalFee      int64  // 订单总金额，单位为分，只能为整数，详见支付金额
	CashFee       int64  // 现金支付金额，单位为分，只能为整数，详见支付金额

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	DeviceInfo          string // 微信支付分配的终端设备号，与下单一致
	RefundChannel       string // 退款渠道
	FeeType             string // 订单金额货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	SettlementRefundFee *int64 // 退款金额=申请退款金额-非充值代金券退款金额，退款金额<=申请退款金额
	SettlementTotalFee  *int64 // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	CashRefundFee       *int64 // 现金退款金额，单位为分，只能为整数，详见支付金额

	// TODO: 增加代金券相关的数据结构
}

func Refund2(clt *core.Client, req *RefundRequest) (resp *RefundResponse, err error) {
	m1 := make(map[string]string, 16)
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
	if req.TransactionId != "" {
		m1["transaction_id"] = req.TransactionId
	}
	if req.OutTradeNo != "" {
		m1["out_trade_no"] = req.OutTradeNo
	}
	m1["out_refund_no"] = req.OutRefundNo
	m1["total_fee"] = strconv.FormatInt(req.TotalFee, 10)
	m1["refund_fee"] = strconv.FormatInt(req.RefundFee, 10)
	if req.RefundFeeType != "" {
		m1["refund_fee_type"] = req.RefundFeeType
	}
	if req.OperUserId != "" {
		m1["op_user_id"] = req.OperUserId
	} else {
		m1["op_user_id"] = clt.MchId()
	}
	m1["sign"] = core.Sign(m1, clt.ApiKey(), md5.New)

	m2, err := Refund(clt, m1)
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

	resp = &RefundResponse{
		AppId: m2["appid"],
		MchId: m2["mch_id"],

		TransactionId: m2["transaction_id"],
		OutTradeNo:    m2["out_trade_no"],
		OutRefundNo:   m2["out_refund_no"],
		RefundId:      m2["refund_id"],

		DeviceInfo:    m2["device_info"],
		RefundChannel: m2["refund_channel"],
		FeeType:       m2["fee_type"],
	}

	var (
		n   int64
		str string
	)
	if n, err = strconv.ParseInt(m2["refund_fee"], 10, 64); err != nil {
		err = fmt.Errorf("parse refund_fee:%q to int64 failed: %s", m2["refund_fee"], err.Error())
		return
	} else {
		resp.RefundFee = n
	}
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
	if str = m2["settlement_refund_fee"]; str != "" {
		if n, err = strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse settlement_refund_fee:%q to int64 failed: %s", str, err.Error())
			return
		} else {
			resp.SettlementRefundFee = util.Int64(n)
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
	if str = m2["cash_refund_fee"]; str != "" {
		if n, err = strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse cash_refund_fee:%q to int64 failed: %s", str, err.Error())
			return
		} else {
			resp.CashRefundFee = util.Int64(n)
		}
	}
	return
}
