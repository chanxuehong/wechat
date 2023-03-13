package updatable

import (
	"github.com/bububa/wechat/mp/core"
)

type SendRequest struct {
	ActivityId   string       `json:"activity_id"`   // 动态消息的 ID，通过 updatableMessage.createActivityId 接口获取
	TargetState  uint         `json:"target_state"`  // 动态消息修改后的状态 0:未开始, 1:已开始
	TemplateInfo TemplateInfo `json:"template_info"` // 动态消息对应的模板信息
}

type TemplateInfo struct {
	List []Param `json:"parameter_list"` // 模板中需要修改的参数
}

type Param struct {
	Name  string `json:"name"`  // 要修改的参数名, 合法值 member_count, room_limit, path, version_type
	Value string `json:"value"` // 修改后的参数值
}

// 发送模板消息, msg 是经过 encoding/json.Marshal 得到的结果符合微信消息格式的任何数据结构, 一般为 *TemplateMessage 类型.
func Send(clt *core.Client, req *SendRequest) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/updatablemsg/send?access_token="
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
