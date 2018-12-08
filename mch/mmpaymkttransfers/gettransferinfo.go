package mmpaymkttransfers

import (
	"github.com/chanxuehong/wechat/mch/core"
)

// 查询企业付款.
//  NOTE: 请求需要双向证书
func GetTransferInfo(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/gettransferinfo", req)
}
