// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"github.com/chanxuehong/wechat/mp"
)

// 卡券核销, 消耗code
//  消耗code 接口是核销卡券的唯一接口，仅支持核销有效期内的卡券，否则会返回错误码invalid time。
//  自定义code（use_custom_code 为true）的优惠券，在code 被核销时，必须调用此接口。
//  用于将用户客户端的code 状态变更。自定义code 的卡券调用接口时， post 数据中需包含card_id，
//  非自定义code 不需上报。
//
//  code:   要消耗序列号
//  cardId: 卡券ID。创建卡券时use_custom_code 填写true时必填。非自定义code 不必填写。
func (clt Client) CardCodeConsume(code, cardId string) (_cardId, openId string, err error) {
	var request = struct {
		Code   string `json:"code"`
		CardId string `json:"card_id,omitempty"`
	}{
		Code:   code,
		CardId: cardId,
	}

	var result struct {
		mp.Error
		Card struct {
			CardId string `json:"card_id"`
		} `json:"card"`
		OpenId string `json:"openid"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/consume?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	_cardId = result.Card.CardId
	openId = result.OpenId
	return
}

// code 解码接口
//  code 解码接口支持两种场景：
//  1.商家获取choos_card_info 后，将card_id 和encrypt_code 字段通过解码接口，获取真实code。
//  2.卡券内跳转外链的签名中会对code 进行加密处理，通过调用解码接口获取真实code。
func (clt Client) CardCodeDecrypt(encryptCode string) (code string, err error) {
	var request = struct {
		EncryptCode string `json:"encrypt_code"`
	}{
		EncryptCode: encryptCode,
	}

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

// 某一张特定的卡卷
type CardCode struct {
	Code      string `json:"code"`       // code
	CardId    string `json:"card_id"`    // 卡券ID
	BeginTime int64  `json:"begin_time"` // 起始使用时间
	EndTime   int64  `json:"end_time"`   // 结束时间
}

// 查询code
//  code:   要查询的序列号
//  cardId: 要消耗序列号所述的card_id， 生成券时use_custom_code 填写true 时必填。非自定义code 不必填写。
func (clt Client) CardCodeGet(code, cardId string) (card *CardCode, openId string, err error) {
	var request = struct {
		Code   string `json:"code"`
		CardId string `json:"card_id,omitempty"`
	}{
		Code:   code,
		CardId: cardId,
	}

	var result struct {
		mp.Error
		Card   CardCode `json:"card"`
		OpenId string   `json:"openid"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/code/get?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	result.Card.Code = code //
	card = &result.Card
	openId = result.OpenId
	return
}

// 更改code.
//  为确保转赠后的安全性，微信允许自定义code的商户对已下发的code进行更改。
//  注：为避免用户疑惑，建议仅在发生转赠行为后（发生转赠后，微信会通过事件推送的方
//  式告知商户被转赠的卡券code）对用户的code进行更改。
func (clt Client) CardCodeUpdate(code, cardId, newCode string) (err error) {
	var request = struct {
		Code    string `json:"code"`
		CardId  string `json:"card_id,omitempty"`
		NewCode string `json:"new_code,omitempty"`
	}{
		Code:    code,
		CardId:  cardId,
		NewCode: newCode,
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

// 设置卡券失效接口.
//  为满足改票、退款等异常情况，可调用卡券失效接口将用户的卡券设置为失效状态。
//  注：设置卡券失效的操作不可逆，即无法将设置为失效的卡券调回有效状态，商家须慎重调用该接口。
func (clt Client) CardCodeUnavailable(code, cardId string) (err error) {
	var request = struct {
		Code   string `json:"code"`
		CardId string `json:"card_id,omitempty"`
	}{
		Code:   code,
		CardId: cardId,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/code/unavailable?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
