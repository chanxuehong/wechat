package code

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 设置卡券失效接口.
func Unavailable(clt *core.Client, id *CardItemIdentifier) (err error) {
	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/card/code/unavailable?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
