package cs

// 获取客服聊天记录 请求消息结构
type RecordGetRequest struct {
	StartTime int64  `json:"starttime"` // 查询开始时间，UNIX时间戳
	EndTime   int64  `json:"endtime"`   // 查询结束时间，UNIX时间戳，每次查询不能跨日查询
	OpenId    string `json:"openid"`    // 普通用户的标识，对当前公众号唯一
	PageSize  int    `json:"pagesize"`  // 每页大小，每页最多拉取1000条
	PageIndex int    `json:"pageindex"` // 查询第几页，从1开始
}

// 一条聊天记录
type Record struct {
	Worker   string `json:"worker"`   // 客服账号
	OpenId   string `json:"openid"`   // 用户的标识，对当前公众号唯一
	OperCode int    `json:"opercode"` // 操作ID（会话状态），具体说明见下文
	Time     int64  `json:"time"`     // 操作时间，UNIX时间戳
	Text     string `json:"text"`     // 聊天记录
}

/*
聊天记录遍历器

	iter, err := Client.CSRecordIterator(request)
	if err != nil {
		...
	}

	for iter.HasNext() {
		records, err := iter.NextPage()
		if err != nil {
			...
		}
		...
	}
*/
type RecordIterator interface {
	HasNext() bool
	NextPage() ([]Record, error)
}
