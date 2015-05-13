// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

//import (
//	"errors"
//
//	"github.com/chanxuehong/wechat/mp"
//)

//type LuckyMoneyUpdateUserBalanceParameters struct {
//	Code   string `json:"code"`              // 必须; 红包的序列号
//	CardId string `json:"card_id,omitempty"` // 可选; 自定义code 的卡券必填。非自定义code可不填。
//
//	Balance *int `json:"balance"` // 必须; 红包余额
//}

//// 更新红包金额.
////  支持领取红包后通过调用“更新红包”接口update 红包余额。
//func (clt Client) LuckyMoneyUpdateUserBalance(para *LuckyMoneyUpdateUserBalanceParameters) (err error) {
//	if para == nil {
//		return errors.New("nil LuckyMoneyUpdateUserBalanceParameters")
//	}
//
//	var result mp.Error
//
//	incompleteURL := "https://api.weixin.qq.com/card/luckymoney/updateuserbalance?access_token="
//	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
//		return
//	}
//
//	if result.ErrCode != mp.ErrCodeOK {
//		err = &result
//		return
//	}
//	return
//}
