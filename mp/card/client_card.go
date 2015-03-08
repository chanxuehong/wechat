// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com)
package card

import (
	"github.com/chanxuehong/wechat/mp"
)

//  获取颜色列表接口
func (clt *Client) CardGetColors() (colors []Color, err error) {
	var result struct {
		Colors []Color `json:"colors"`
		mp.Error
	}
	incompleteURL := "https://api.weixin.qq.com/card/getcolors?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	colors = result.Colors
	return

}

//  创建卡券
//  Card: 卡券信息
func (clt *Client) CardCreate(card Card) (cardId string, err error) {
	var request struct {
		Card Card `json:"card"`
	}
	request.Card = card
	var result struct {
		CardId string `json:"card_id"`
		mp.Error
	}
	incompleteURL := "https://api.weixin.qq.com/card/create?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	cardId = result.CardId
	return
}

// 删除卡券
// cardId: 卡券Id
func (clt *Client) CardDelete(cardId string) (err error) {
	var request struct {
		CardId string `json:"card_id"`
	}
	request.CardId = cardId
	var result mp.Error
	incompleteURL := "https://api.weixin.qq.com/card/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 查询code
func (clt *Client) CardGet(cardId string) (card *ResultCard, err error) {
	var request struct {
		CardId string `json:"card_id"`
	}
	request.CardId = cardId

	var result struct {
		mp.Error
		OpenId string      `json:"openid"`
		Card   *ResultCard `json:"card"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/get?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	card = result.Card
	return
}

// 解码接口
func (clt *Client) CardDecrypt(encryptCode string) (code string, err error) {
	var request struct {
		EncryptCode string `json:"encrypt_code"`
	}
	request.EncryptCode = encryptCode

	var result struct {
		mp.Error
		Code string `json:"code"`
	}
	incompleteURL := "https://api.weixin.qq.com/card/code/decrypt?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	code = result.Code
	return
}

// 批量查询卡列表
// offset:查询起始偏移量
// count :总数
func (clt *Client) CardBatchGet(offset, count int) (cardsIdList []string, err error) {
	var request struct {
		Offset int `json:"offset"`
		Count  int `json:"count"`
	}
	request.Offset = offset
	request.Count = count

	var result struct {
		mp.Error
		CardIdList []string `json:"card_id_list"`
		TotalSum   int      `json:"total_sum"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/batchget?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	cardsIdList = result.CardIdList
	return

}

//更改code，为确保转赠后的安全性，微信允许自定义code的商户对已下发的code进行更改。
//注：为避免用户疑惑，建议仅在发生转赠行为后（发生转赠后，微信会通过事件推送的方
//式告知商户被转赠的卡券code）对用户的code进行更改

func (clt *Client) CardCodeUpdate(cardId string) (err error) {
	var request struct {
		Code    string `json:"code"`
		CardId  string `json:"card_id"`
		NewCode string `json:"new_code"`
	}
	request.CardId = cardId

	var result struct {
		mp.Error
	}
	incompleteURL := "https://api.weixin.qq.com/card/code/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	return

}

// 设置卡券失效接口

func (clt *Client) CardCodeUnavailable(code, cardId string) (err error) {
	var request struct {
		Code   string `json:"code"`
		CardId string `json:"card_id"`
	}
	request.Code = code
	request.CardId = cardId

	var result struct {
		mp.Error
	}
	incompleteURL := "https://api.weixin.qq.com/card/code/unavailable?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	return

}

// 更改卡券信息接口
func (clt *Client) CardUpdate(cardId string, baseInfo BaseInfo) (err error) {
	var request struct {
		CardId     string `json:"card_id"`
		MemberCard struct {
			baseInfo BaseInfo `json:"base_info"`
		} `json:"member_card"`
	}
	request.CardId = cardId
	request.MemberCard.baseInfo = baseInfo

	var result struct {
		mp.Error
	}
	incompleteURL := "https://api.weixin.qq.com/card/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 库存修改接口
// cardId：卡券ID
// increaseStock：增加多少库存
// reduceStock:减少多少库存
func (clt *Client) CardModifyStock(cardId string, increaseStock, reduceStock int) (err error) {
	var request struct {
		CardId        string `json:"card_id"`
		IncreaseStock int    `json:"increase_stock_value"`
		ReduceStock   int    `json:"reduce_stock_value"`
	}
	request.CardId = cardId
	request.IncreaseStock = increaseStock
	request.ReduceStock = reduceStock

	var result struct {
		mp.Error
	}
	incompleteURL := "https://api.weixin.qq.com/card/modifystock?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 卡券投放，创建二维码
// CardQRCode:       卡券信息
// ExpireSeconds: 二维码的有效时间，以秒为单位。
func (clt *Client) CreateCardQRCode(cardQRCode CardQRCode, ExpireSeconds int) (ticket string, err error) {
	var request struct {
		ExpireSeconds int    `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Card CardQRCode `json:"card"`
		} `json:"action_info"`
	}
	request.ExpireSeconds = ExpireSeconds
	request.ActionName = "QR_CARD"
	request.ActionInfo.Card = cardQRCode

	var result struct {
		mp.Error
		TemporaryQRCode string `json:"ticket"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	ticket = result.TemporaryQRCode
	return
}

// 激活绑定会员卡
// initBonus:初始积分
// InitBalance:初始余额
// memberShipNumber:会员卡编号
// code:创建会员卡时候获取的code
// card_id:卡券id
func (clt *Client) MemberCardActivate(cardId, code, memberShipNumber string, initBonus, initBalance int) (err error) {
	var request struct {
		InitBonus        int    `json:"init_bonus"`
		InitBalance      int    `json:"init_balance"`
		MemberShipNumber string `json:"membership_number"`
		Code             string `json:"code"`
		CardId           string `json:"card_id"`
	}
	request.InitBonus = initBonus
	request.InitBalance = initBalance
	request.MemberShipNumber = memberShipNumber
	request.Code = code
	request.CardId = cardId

	var result struct {
		mp.Error
	}
	incompleteURL := "https://api.weixin.qq.com/card/membercard/activate?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 会员卡交易，会员卡交易后每次积分及余额变更需通过接口通知微信， 便于后续消息通知及其他扩展功能
// code :要消耗的序列号
// addBonus : 需要变更的积分，扣除积分用-
// recordBonus: 商家自定义积分消耗记录，不超过14个汉子
// addBalance:需要变更的金额，扣除金额用-表示
// record_balance:商家自定义金额消费记录
// cardId:要消耗序列号所述的card_id。 自定义code的会员卡必填
func (clt *Client) MemberCardUpdate(cardId, code, recordBonus, recordbalance, addBonus, addBalance string) (err error, resultBonus, resultBalance, openid string) {
	var request struct {
		Code          string `json:"code"`
		CardId        string `json:"card_id"`
		RecordBonus   string `json:"record_bonus"`
		AddBonus      string `json:"add_bonus"`
		AddBalance    string `json:"add_balance"`
		RecordBalance string `json:"record_balance"`
	}
	request.Code = code
	request.CardId = cardId
	request.RecordBonus = recordBonus
	request.AddBonus = addBonus
	request.AddBalance = addBalance
	request.RecordBalance = recordbalance

	var result struct {
		mp.Error
		ResultBonus   string `json:"result_bonus"`
		ResultBalance string `json:"result_balance"`
		Openid        string `json:"openid"`
	}
	incompleteURL := "https://api.weixin.qq.com/card/membercard/updateuser?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	resultBonus = result.ResultBonus
	resultBalance = result.ResultBalance
	openid = result.Openid
	return
}
