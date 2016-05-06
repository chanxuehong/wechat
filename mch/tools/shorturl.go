package tools

import (
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 转换短链接.
func ShortURL(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/tools/shorturl", req)
}
