<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package code

import (
	"github.com/chanxuehong/wechat/mp"
)

// 核销Code接口.
func Consume(clt *mp.Client, id *CardItemIdentifier) (cardId, openId string, err error) {
	var result struct {
		mp.Error
=======
package code

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 核销Code接口.
func Consume(clt *core.Client, id *CardItemIdentifier) (cardId, openId string, err error) {
	var result struct {
		core.Error
>>>>>>> github/v2
		Card struct {
			CardId string `json:"card_id"`
		} `json:"card"`
		OpenId string `json:"openid"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/consume?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result.Error
		return
	}
	cardId = result.Card.CardId
	openId = result.OpenId
	return
}
