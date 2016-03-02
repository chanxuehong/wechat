package code

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 更改Code接口.
func Update(clt *core.Client, id *CardItemIdentifier, newCode string) (err error) {
	request := struct {
		*CardItemIdentifier
		NewCode string `json:"new_code,omitempty"`
	}{
		CardItemIdentifier: id,
		NewCode:            newCode,
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/card/code/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
