package boardingpass

import (
	"github.com/mingjunyang/wechat.v2/mp/core"
)

type CheckinParameters struct {
	Code   string `json:"code"`              // 必须; 卡券Code码。
	CardId string `json:"card_id,omitempty"` // 可选; 卡券ID，自定义Code码的卡券必填。

	PassengerName string `json:"passenger_name,omitempty"` // 必须; 乘客姓名, 上限为15 个汉字.
	Class         string `json:"class,omitempty"`          // 必须; 舱等，如头等舱等，上限为5个汉字。
	ETKT_NBR      string `json:"etkt_bnr,omitempty"`       // 必须; 电子客票号，上限为14个数字。
	Seat          string `json:"seat,omitempty"`           // 可选; 乘客座位号。
	QrcodeData    string `json:"qrcode_data,omitempty"`    // 可选; 二维码数据。乘客用于值机的二维码字符串，微信会通过此数据为用户生成值机用的二维码。
	IsCancel      *bool  `json:"is_cancel,omitempty"`      // 可选; 是否取消值机。填写true或false。true代表取消，如填写true上述字段（如calss等）均不做判断，机票返回未值机状态，乘客可重新值机。默认填写false。
}

// 更新飞机票信息接口
func Checkin(clt *core.Client, para *CheckinParameters) (err error) {
	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/card/boardingpass/checkin?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
