// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"errors"

	"github.com/chanxuehong/wechat/mp"
)

type BoardingPassCheckinParameters struct {
	Code   string `json:"code"`              // 飞机票的序列号
	CardId string `json:"card_id,omitempty"` // 需办理值机的机票card_id。自定义code 的飞机票为必填

	PassengerName string `json:"passenger_name,omitempty"` // 乘客姓名，上限为15 个汉字。
	Class         string `json:"class,omitempty"`          // 舱等，如头等舱等，上限为5 个汉字。
	Seat          string `json:"seat,omitempty"`           // 乘客座位号。
	ETKT_NBR      string `json:"etkt_bnr,omitempty"`       // 电子客票号，上限为14 个数字
	QRCodeData    string `json:"qrcode_data,omitempty"`    // 二维码数据。乘客用于值机的二维码字符串，微信会通过此数据为用户生成值机用的二维码。
	IsCancel      *bool  `json:"is_cancel,omitempty"`      // 是否取消值机。填写true 或false。true 代表取消，如填写true 上述字段（如calss 等）均不做判断，机票返回未值机状态，乘客可重新值机。默认填写false
}

// 在线值机接口.
//  领取电影票后通过调用“更新电影票”接口update 电影信息及用户选座信息
func (clt *Client) BoardingPassCheckin(para *BoardingPassCheckinParameters) (err error) {
	if para == nil {
		return errors.New("nil BoardingPassCheckinParameters")
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/boardingpass/checkin?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
