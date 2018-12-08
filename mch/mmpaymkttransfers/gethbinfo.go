package mmpaymkttransfers

import (
	"github.com/chanxuehong/wechat/mch/core"
)

// 红包查询接口.
//  NOTE: 请求需要双向证书
func GetRedPackInfo(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML(core.APIBaseURL()+"/mmpaymkttransfers/gethbinfo", req)
}
