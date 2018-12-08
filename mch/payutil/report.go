package payutil

import (
	"strconv"
	"time"

	"github.com/chanxuehong/wechat/mch/core"
	"github.com/chanxuehong/wechat/util"
)

// Report 交易保障.
func Report(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/payitil/report", req)
}

type ReportRequest struct {
	XMLName      struct{}  `xml:"xml" json:"-"`
	DeviceInfo   string    `xml:"device_info"`   // 微信支付分配的终端设备号，商户自定义
	NonceStr     string    `xml:"nonce_str"`     // 随机字符串，不长于32位。NOTE: 如果为空则系统会自动生成一个随机字符串。
	SignType     string    `xml:"sign_type"`     // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	InterfaceURL string    `xml:"interface_url"` // 刷卡支付终端上报统一填：https://api.mch.weixin.qq.com/pay/batchreport/micropay/total
	UserIP       string    `xml:"user_ip"`       // 发起接口调用时的机器IP
	Trades       string    `xml:"trades"`        // 上报数据包
	ExecuteTime  int       `xml:"execute_time"`  // 接口耗时情况，单位为毫秒
	ReturnCode   string    `xml:"return_code"`   // 返回状态码
	ReturnMsg    string    `xml:"return_msg"`    // 返回信息
	ResultCode   string    `xml:"result_code"`   // 业务结果
	ErrCode      string    `xml:"err_code"`      // 错误代码
	ErrCodeDesc  string    `xml:"err_code_des"`  // 错误代码描述
	OutTradeNo   string    `xml:"out_trade_no"`  // 商户订单号
	Time         time.Time `xml:"time"`          // 商户上报时间
}

// Report2 交易保障.
func Report2(clt *core.Client, req *ReportRequest) (err error) {
	m1 := make(map[string]string, 24)
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
	if req.InterfaceURL != "" {
		m1["interface_url"] = req.InterfaceURL
	}
	if req.UserIP != "" {
		m1["user_ip"] = req.UserIP
	}
	if req.Trades != "" {
		m1["trades"] = req.Trades
	}
	if req.ExecuteTime > 0 {
		m1["execute_time"] = strconv.Itoa(req.ExecuteTime)
	}
	if req.ReturnCode != "" {
		m1["return_code"] = req.ReturnCode
	}
	if req.ReturnMsg != "" {
		m1["return_msg"] = req.ReturnMsg
	}
	if req.ResultCode != "" {
		m1["result_code"] = req.ResultCode
	}
	if req.ErrCode != "" {
		m1["err_code"] = req.ErrCode
	}
	if req.ErrCodeDesc != "" {
		m1["err_code_des"] = req.ErrCodeDesc
	}
	if req.OutTradeNo != "" {
		m1["out_trade_no"] = req.OutTradeNo
	}
	if !req.Time.IsZero() {
		m1["time"] = core.FormatTime(req.Time)
	}

	_, err = Report(clt, m1)
	return
}
