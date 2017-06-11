package pay

import (
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 提交刷卡支付.
func MicroPay(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/pay/micropay", req)
}
