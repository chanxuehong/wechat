// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"errors"
	"net/url"

	"github.com/chanxuehong/wechat/mp"
)

// 根据二维码的ticket得到二维码图片的url
func QRCodePicURL(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(ticket)
}

// 创建二维码的参数
type CardQRCodeInfo struct {
	CardId        string `json:"card_id"`                  // 必须; 卡券ID
	Code          string `json:"code,omitempty"`           // 可选; 指定卡券code码, 只能被领一次. use_custom_code 字段为true 的卡券必须填写, 非自定义code 不必填写.
	OpenId        string `json:"openid,omitempty"`         // 可选; 指定领取者的openid, 只有该用户能领取. bind_openid字段为true 的卡券必须填写, 非自定义openid 不必填写.
	ExpireSeconds *int   `json:"expire_seconds,omitempty"` // 可选; 指定二维码的有效时间, 范围是60 ~ 1800 秒. 不填默认为永久有效.
	IsUniqueCode  *bool  `json:"is_unique_code,omitempty"` // 可选; 指定下发二维码, 生成的二维码随机分配一个code, 领取后不可再次扫描. 填写true 或false. 默认false.
	Balance       *int   `json:"balance,omitempty"`        // 可选; 红包余额, 以分为单位. 红包类型必填(LUCKY_MONEY), 其他卡券类型不填.
	OuterId       *int64 `json:"outer_id,omitempty"`       // 可选; 领取场景值, 用于领取渠道的数据统计, 默认值为0, 字段类型为整型. 用户领取卡券后触发的事件推送中会带上此自定义场景值.
}

// 卡券投放, 创建二维码.
//  创建卡券后, 商户可通过接口生成一张卡券二维码供用户扫码后添加卡券到卡包.
func (clt Client) CardQRCodeCreate(info *CardQRCodeInfo) (ticket string, err error) {
	if info == nil {
		err = errors.New("nil CardQRCodeInfo")
		return
	}

	var request struct {
		ActionName string `json:"action_name"`
		ActionInfo struct {
			Card *CardQRCodeInfo `json:"card,omitempty"`
		} `json:"action_info"`
	}
	request.ActionName = "QR_CARD"
	request.ActionInfo.Card = info

	var result struct {
		mp.Error
		Ticket string `json:"ticket"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/qrcode/create?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	ticket = result.Ticket
	return
}
