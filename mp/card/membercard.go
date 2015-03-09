// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"errors"

	"github.com/chanxuehong/wechat/mp"
)

type MemberCardActivateParameters struct {
	InitBonus   int `json:"init_bonus,omitempty"`   // 可选; 初始积分，不填为0
	InitBalance int `json:"init_balance,omitempty"` // 可选; 初始余额，不填为0

	BonusURL   string `json:"bonus_url,omitempty"`   // 可选; 积分查询，仅用于init_bonus 无法同步的情况填写，调转外链查询积分
	BalanceURL string `json:"balance_url,omitempty"` // 可选; 余额查询，仅用于init_balance 无法同步的情况填写，调转外链查询积分

	MembershipNumber string `json:"membership_number"` // 必填，会员卡编号，作为序列号显示在用户的卡包里。
	Code             string `json:"code"`              // 必填，创建会员卡时获取的code。
	CardId           string `json:"card_id,omitempty"` // 可选; 卡券ID。自定义code 的会员卡必填card_id，非自定义code 的会员卡不必填。
}

// 激活/绑定会员卡
func (clt *Client) MemberCardActivate(para *MemberCardActivateParameters) (err error) {
	if para == nil {
		return errors.New("nil MemberCardActivateParameters")
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/membercard/activate?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

type MemberCardUpdateUserParameters struct {
	Code   string `json:"code"`              // 要消耗的序列号。
	CardId string `json:"card_id,omitempty"` // 要消耗序列号所述的card_id。自定义code 的会员卡必填

	AddBonus      int    `json:"add_bonus,omitempty"`      // 需要变更的积分，扣除积分用“-“表示。
	RecordBonus   string `json:"record_bonus,omitempty"`   // 商家自定义积分消耗记录，不超过14 个汉字
	AddBalance    int    `json:"add_balance,omitempty"`    // 需要变更的余额，扣除金额用“-”表示。单位为分
	RecordBalance string `json:"record_balance,omitempty"` // 商家自定义金额消耗记录，不超过14 个汉字
}

type MemberCardUpdateUserResult struct {
	ResultBonus   int    `json:"result_bonus"`   // 当前用户积分总额
	ResultBalance int    `json:"result_balance"` // 当前用户预存总金额
	OpenId        string `json:"openid"`         // 用户openid
}

// 会员卡交易.
//  会员卡交易后每次积分及余额变更需通过接口通知微信，便于后续消息通知及其他扩展功能。
func (clt *Client) MemberCardUpdateUser(para *MemberCardUpdateUserParameters) (rst *MemberCardUpdateUserResult, err error) {
	if para == nil {
		err = errors.New("nil MemberCardUpdateUserParameters")
		return
	}

	var result struct {
		mp.Error
		MemberCardUpdateUserResult
	}

	incompleteURL := "https://api.weixin.qq.com/card/membercard/updateuser?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	rst = &result.MemberCardUpdateUserResult
	return
}
