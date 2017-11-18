package promotion

import (
	"gopkg.in/chanxuehong/wechat.v2/mch/core"
)

// 企业付款.
//  NOTE: 请求需要双向证书
func Transfers(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/promotion/transfers", req)
}
