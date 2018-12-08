package user

import (
	"github.com/chanxuehong/wechat/mp/card/code"
	"github.com/chanxuehong/wechat/mp/core"
)

// 获取用户已领取卡券接口
//  openid: 需要查询的用户openid
//  cardid: 卡券ID。不填写时默认查询当前appid下的卡券。
func GetCardList(clt *core.Client, openid, cardid string) (list []code.CardItemIdentifier, err error) {
	request := struct {
		OpenId string `json:"openid"`
		CardId string `json:"card_id,omitempty"`
	}{
		OpenId: openid,
		CardId: cardid,
	}

	var result struct {
		core.Error
		CardList []code.CardItemIdentifier `json:"card_list"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/user/getcardlist?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.CardList
	return
}
