package pay

import (
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 提交刷卡支付.
func MicroPay(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/pay/micropay", req)
}

// 撤销订单.
//  NOTE: 请求需要双向证书.
func Reverse(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/secapi/pay/reverse", req)
}
