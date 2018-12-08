package qrcode

import (
	"net/url"

	"github.com/chanxuehong/wechat/mp/core"
)

func QrcodePicURL(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(ticket)
}

type CreateParameters struct {
	CardId        string `json:"card_id"`                  // 必须; 卡券ID
	Code          string `json:"code,omitempty"`           // 可选; use_custom_code字段为true的卡券必须填写，非自定义code不必填写
	OpenId        string `json:"openid,omitempty"`         // 可选; 指定领取者的openid，只有该用户能领取。bind_openid字段为true的卡券必须填写，非指定openid不必填写。
	ExpireSeconds int    `json:"expire_seconds,omitempty"` // 可选; 指定二维码的有效时间，范围是60 ~ 1800秒。不填(值为0)默认为永久有效。
	IsUniqueCode  *bool  `json:"is_unique_code,omitempty"` // 可选; 指定下发二维码，生成的二维码随机分配一个code，领取后不可再次扫描。填写true或false。默认false。
	OuterId       *int64 `json:"outer_id,omitempty"`       // 可选; 领取场景值，用于领取渠道的数据统计，默认值为0，字段类型为整型，长度限制为60位数字。用户领取卡券后触发的事件推送中会带上此自定义场景值。
}

type QrcodeInfo struct {
	Ticket        string `json:"ticket"`
	URL           string `json:"url"`
	ExpireSeconds int    `json:"expire_seconds"` // 0 表示永久二维码
}

// 卡券投放, 创建二维码接口.
func Create(clt *core.Client, para *CreateParameters) (info *QrcodeInfo, err error) {
	request := struct {
		ActionName    string `json:"action_name"`
		ExpireSeconds int    `json:"expire_seconds,omitempty"`
		ActionInfo    struct {
			Card *CreateParameters `json:"card,omitempty"`
		} `json:"action_info"`
	}{
		ActionName:    "QR_CARD",
		ExpireSeconds: para.ExpireSeconds,
	}
	request.ActionInfo.Card = para

	var result struct {
		core.Error
		QrcodeInfo
	}

	incompleteURL := "https://api.weixin.qq.com/card/qrcode/create?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.QrcodeInfo
	return
}
