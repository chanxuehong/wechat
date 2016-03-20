<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package code

import (
	"github.com/chanxuehong/wechat/mp"
)

// 更改Code接口.
func Update(clt *mp.Client, id *CardItemIdentifier, newCode string) (err error) {
=======
package code

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 更改Code接口.
func Update(clt *core.Client, id *CardItemIdentifier, newCode string) (err error) {
>>>>>>> github/v2
	request := struct {
		*CardItemIdentifier
		NewCode string `json:"new_code,omitempty"`
	}{
		CardItemIdentifier: id,
		NewCode:            newCode,
	}

<<<<<<< HEAD
	var result mp.Error
=======
	var result core.Error
>>>>>>> github/v2

	incompleteURL := "https://api.weixin.qq.com/card/code/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result
		return
	}
	return
}
