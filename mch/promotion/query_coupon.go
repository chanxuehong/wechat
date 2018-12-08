package promotion

import (
	"github.com/chanxuehong/wechat/mch/core"
)

// 查询代金券信息.
func QueryCoupon(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/promotion/query_coupon", req)
}
