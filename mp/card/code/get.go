package code

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 查询code.
func Get(clt *core.Client, id *CardItemIdentifier) (info *CardItem, err error) {
	var result struct {
		core.Error
		CardItem
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/get?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	result.CardItem.Code = id.Code
	info = &result.CardItem
	return
}
