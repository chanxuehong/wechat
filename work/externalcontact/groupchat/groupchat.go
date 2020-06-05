package groupchat

type GroupChat struct {
	ChatId     string   `json:"chat_id,omitempty"`     // 客户群ID
	Name       string   `json:"name,omitempty"`        // 群名
	Owner      string   `json:"owner,omitempty"`       // 群主ID
	CreateTime int64    `json:"create_time,omitempty"` // 群的创建时间
	Notice     string   `json:"notice,omitempty"`      // 群公告
	Members    []Member `json:"member_list,omitempty"` // 群成员列表
	Status     int      `json:"status,omitempty"`      // 客户群状态。0 - 正常; 1 - 跟进人离职; 2 - 离职继承中; 3 - 离职继承完成
}

type Member struct {
	UserId    string `json:"userid,omitempty"`     // 群成员id
	Type      uint   `json:"type,omitempty"`       // 成员类型。1 - 企业成员; 2 - 外部联系人
	JoinTime  int64  `json:"join_time,omitempty"`  // 入群时间
	JoinScene uint   `json:"join_scene,omitempty"` // 入群方式。1 - 由成员邀请入群（直接邀请入群）;2 - 由成员邀请入群（通过邀请链接入群 ; 3 - 通过扫描群二维码入群
}
