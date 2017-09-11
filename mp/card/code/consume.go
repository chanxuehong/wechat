package code

import (
	"github.com/mingjunyang/wechat.v2/mp/core"
)

// 核销Code接口.
func Consume(clt *core.Client, id *CardItemIdentifier) (cardId, openId string, err error) {
	var result struct {
		core.Error
		Card struct {
			CardId string `json:"card_id"`
		} `json:"card"`
		OpenId string `json:"openid"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/consume?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	cardId = result.Card.CardId
	openId = result.OpenId
	return
}
