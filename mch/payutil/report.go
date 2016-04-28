package payutil

import (
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 测速上报.
func Report(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.core.weixin.qq.com/payitil/report", req)
}
