package groupchat

import (
	"github.com/bububa/wechat/work/core"
)

// Get 通过客户群ID，获取详情。包括群名、群成员列表、群成员入群时间、入群方式。（客户群是由具有客户群使用权限的成员创建的外部群）
// chat_id: 客户群ID
func Get(clt *core.Client, chatId string) (rslt *GroupChat, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/get?access_token="

	var result struct {
		core.Error
		*GroupChat `json:"group_chat"`
	}
	if err = clt.PostJSON(incompleteURL, map[string]string{"chat_id": chatId}, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = result.GroupChat
	return
}
