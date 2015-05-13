// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     gaowenbin(gaowenbinmarr@gmail.com), chanxuehong(chanxuehong@gmail.com)

package card

import (
	"errors"

	"github.com/chanxuehong/wechat/mp"
)

type MovieTicketUpdateUserParameters struct {
	Code   string `json:"code"`              // 必须; 电影票的序列号。
	CardId string `json:"card_id,omitempty"` // 可选; 电影票card_id。自定义code 的电影票为必填，非自定义code 的电影票不必填。

	TicketClass   string `json:"ticket_class,omitempty"`   // 必须; 电影票的类别，如2D、3D
	ShowTime      int64  `json:"show_time,omitempty"`      // 必须; 电影放映时间对应的时间戳。
	Duration      int    `json:"duration,omitempty"`       // 必须; 放映时长，填写整数
	ScreeningRoom string `json:"screening_room,omitempty"` // 必须; 该场电影的影厅信息
	SeatNumber    string `json:"seat_number,omitempty"`    // 必须; 座位号
}

// 更新电影票.
//  领取电影票后通过调用“更新电影票”接口update 电影信息及用户选座信息
func (clt Client) MovieTicketUpdateUser(para *MovieTicketUpdateUserParameters) (err error) {
	if para == nil {
		return errors.New("nil MovieTicketUpdateUserParameters")
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/card/movieticket/updateuser?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
