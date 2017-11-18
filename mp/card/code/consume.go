// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package code

import (
	"gopkg.in/chanxuehong/wechat.v1/mp"
)

// 核销Code接口.
func Consume(clt *mp.Client, id *CardItemIdentifier) (cardId, openId string, err error) {
	var result struct {
		mp.Error
		Card struct {
			CardId string `json:"card_id"`
		} `json:"card"`
		OpenId string `json:"openid"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/consume?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	cardId = result.Card.CardId
	openId = result.OpenId
	return
}
