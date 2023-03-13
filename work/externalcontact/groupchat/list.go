package groupchat

import (
	"fmt"

	"github.com/bububa/wechat/work/core"
)

type OwnerFilter struct {
	UserIds  []string `json:"userid_list,omitempty"`  // 用户ID列表。最多100个
	PartyIds []uint64 `json:"partyid_list,omitempty"` // 部门ID列表。最多100个
}

type GroupChatList []GroupChat

// List 该接口用于获取配置过客户群管理的客户群列表。
// status_filter: 群状态过滤。0 - 所有列表; 1 - 离职待继承; 2 - 离职继承中; 3 - 离职继承完成; 默认为0
// owner_filter: 群主过滤。如果不填，表示获取全部群主的数据
// offset: 分页，偏移量;
// limit: 分页，预期请求的数据量，取值范围 1 ~ 1000;
func List(clt *core.Client, offset int, limit int, statusFilter uint, ownerFilter *OwnerFilter) (rslt GroupChatList, err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/groupchat/list?access_token="

	if offset < 0 {
		err = fmt.Errorf("invalid offset: %d", offset)
		return
	}
	if limit <= 0 {
		err = fmt.Errorf("invalid limit: %d", limit)
		return
	}

	var request = struct {
		StatusFilter uint         `json:"status_filter,omitempty"`
		OwnerFilter  *OwnerFilter `json:"owner_filter,omitempty"`
		Offset       int          `json:"offset"`
		Limit        int          `json:"limit"`
	}{
		StatusFilter: statusFilter,
		OwnerFilter:  ownerFilter,
		Offset:       offset,
		Limit:        limit,
	}
	var result struct {
		core.Error
		GroupChatList `json:"group_chat_list"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = result.GroupChatList
	return
}
