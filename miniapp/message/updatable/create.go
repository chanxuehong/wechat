package updatable

import (
	"github.com/bububa/wechat/mp/core"
)

type Activity struct {
	core.Error
	ActivityId     string `json:"activity_id"`
	ExpirationTime int64  `json:"expiration_time"`
}

// 发送模板消息, msg 是经过 encoding/json.Marshal 得到的结果符合微信消息格式的任何数据结构, 一般为 *TemplateMessage 类型.
func Create(clt *core.Client) (activity Activity, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/activityid/create?access_token="
	if err = clt.GetJSON(incompleteURL, &activity); err != nil {
		return
	}
	if activity.ErrCode != core.ErrCodeOK {
		err = &activity.Error
		return
	}
	return activity, nil
}
