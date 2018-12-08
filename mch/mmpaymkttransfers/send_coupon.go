package mmpaymkttransfers

import (
	"github.com/chanxuehong/wechat/mch/core"
)

// 发放代金券.
//  请求需要双向证书
func SendCoupon(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/send_coupon", req)
}
