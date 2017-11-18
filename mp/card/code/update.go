// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package code

import (
	"gopkg.in/chanxuehong/wechat.v1/mp"
)

// 更改Code接口.
func Update(clt *mp.Client, id *CardItemIdentifier, newCode string) (err error) {
	request := struct {
		*CardItemIdentifier
		NewCode string `json:"new_code,omitempty"`
	}{
		CardItemIdentifier: id,
		NewCode:            newCode,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/code/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
