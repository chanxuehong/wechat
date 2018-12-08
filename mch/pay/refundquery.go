package pay

import (
	"fmt"
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/mch/core"
	wechatutil "github.com/chanxuehong/wechat/util"
)

// RefundQuery 查询退款.
func RefundQuery(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/refundquery", req)
}

type RefundQueryRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选参数, 四选一
	TransactionId string `xml:"transaction_id"` // 微信订单号
	OutTradeNo    string `xml:"out_trade_no"`   // 商户订单号
	OutRefundNo   string `xml:"out_refund_no"`  // 商户退款单号
	RefundId      string `xml:"refund_id"`      // 微信退款单号

	// 可选参数
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType string `xml:"sign_type"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
}

type RefundQueryResponse struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选返回
	TransactionId string       `xml:"transaction_id"` // 微信订单号
	OutTradeNo    string       `xml:"out_trade_no"`   // 商户系统内部的订单号
	TotalFee      int64        `xml:"total_fee"`      // 订单总金额，单位为分，只能为整数，详见支付金额
	CashFee       int64        `xml:"cash_fee"`       // 现金支付金额，单位为分，只能为整数，详见支付金额
	RefundCount   int          `xml:"refund_count"`   // 退款笔数
	RefundList    []RefundItem `xml:"refund_list"`    // 退款列表

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	SettlementTotalFee *int64 `xml:"settlement_total_fee"` // 应结订单金额=订单金额-非充值代金券金额，应结订单金额<=订单金额。
	FeeType            string `xml:"fee_type"`             // 订单金额货币类型，符合ISO 4217标准的三位字母代码，默认人民币：CNY，其他值列表详见货币类型
	CashFeeType        string `xml:"cash_fee_type"`        // 现金支付货币类型
}

type RefundItem struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选返回
	OutRefundNo      string `xml:"out_refund_no"`      // 商户退款单号
	RefundId         string `xml:"refund_id"`          // 微信退款单号
	RefundFee        int64  `xml:"refund_fee"`         // 申请退款金额
	RefundStatus     string `xml:"refund_status"`      // 退款状态
	RefundRecvAccout string `xml:"refund_recv_accout"` // 退款入账账户

	// 下面字段都是可选返回的(详细见微信支付文档), 为空值表示没有返回, 程序逻辑里需要判断
	RefundChannel       string    `xml:"refund_channel"`        // 退款渠道
	SettlementRefundFee *int64    `xml:"settlement_refund_fee"` // 退款金额
	RefundAccount       string    `xml:"refund_account"`        // 退款资金来源
	RefundSuccessTime   time.Time `xml:"refund_success_time"`   // 退款成功时间
}

// RefundQuery2 查询退款.
//  NOTE: 该函数不支持 代金券 功能, 如果有 代金券 功能请使用 RefundQuery 函数.
func RefundQuery2(clt *core.Client, req *RefundQueryRequest) (resp *RefundQueryResponse, err error) {
	m1 := make(map[string]string, 16)
	if req.TransactionId != "" {
		m1["transaction_id"] = req.TransactionId
	}
	if req.OutTradeNo != "" {
		m1["out_trade_no"] = req.OutTradeNo
	}
	if req.OutRefundNo != "" {
		m1["out_refund_no"] = req.OutRefundNo
	}
	if req.RefundId != "" {
		m1["refund_id"] = req.RefundId
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = wechatutil.NonceStr()
	}
	if req.SignType != "" {
		m1["sign_type"] = req.SignType
	}

	m2, err := RefundQuery(clt, m1)
	if err != nil {
		return nil, err
	}

	resp = &RefundQueryResponse{
		TransactionId: m2["transaction_id"],
		OutTradeNo:    m2["out_trade_no"],
		FeeType:       m2["fee_type"],
		CashFeeType:   m2["cash_fee_type"],
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
	if str := m2["refund_count"]; str != "" {
		if n, err := strconv.ParseInt(str, 10, 64); err != nil {
			err = fmt.Errorf("parse refund_count:%q to int64 failed: %s", str, err.Error())
			return nil, err
		} else {
			resp.RefundCount = int(n)
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

	resp.RefundList = make([]RefundItem, resp.RefundCount)
	for i := 0; i < resp.RefundCount; i++ {
		resp.RefundList[i].OutRefundNo = m2["out_refund_no_"+strconv.Itoa(i)]
		resp.RefundList[i].RefundId = m2["refund_id_"+strconv.Itoa(i)]
		resp.RefundList[i].RefundStatus = m2["refund_status_"+strconv.Itoa(i)]
		resp.RefundList[i].RefundRecvAccout = m2["refund_recv_accout_"+strconv.Itoa(i)]
		resp.RefundList[i].RefundChannel = m2["refund_channel_"+strconv.Itoa(i)]
		resp.RefundList[i].RefundAccount = m2["refund_account_"+strconv.Itoa(i)]

		if str := m2["refund_fee_"+strconv.Itoa(i)]; str != "" {
			if n, err := strconv.ParseInt(str, 10, 64); err != nil {
				err = fmt.Errorf("parse refund_fee_%d:%q to int64 failed: %s", i, str, err.Error())
				return nil, err
			} else {
				resp.RefundList[i].RefundFee = n
			}
		}
		if str := m2["settlement_refund_fee_"+strconv.Itoa(i)]; str != "" {
			if n, err := strconv.ParseInt(str, 10, 64); err != nil {
				err = fmt.Errorf("parse settlement_refund_fee_%d:%q to int64 failed: %s", i, str, err.Error())
				return nil, err
			} else {
				resp.RefundList[i].SettlementRefundFee = wechatutil.Int64(n)
			}
		}
		if str := m2["refund_success_time_"+strconv.Itoa(i)]; str != "" {
			// 2016-07-25 15:26:26
			if t, err := time.ParseInLocation("2006-01-02 15:04:05", str, wechatutil.BeijingLocation); err != nil {
				err = fmt.Errorf("parse refund_success_time_%d:%q to time.Time failed: %s", i, str, err.Error())
				return nil, err
			} else {
				resp.RefundList[i].RefundSuccessTime = t
			}
		}
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

	return resp, nil
}
