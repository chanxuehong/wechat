package cs

// 获取客服聊天记录 请求消息结构
type RecordRequest struct {
	StartTime int64  `json:"starttime"`
	EndTime   int64  `json:"endtime"`
	OpenId    string `json:"openid"`
	PageSize  int    `json:"pagesize"`
	PageIndex int    `json:"pageindex"`
}

// 一条聊天记录
type RecordItem struct {
	Worker   string `json:"worker"`
	OpenId   string `json:"openid"`
	OperCode int    `json:"opercode"`
	Time     int64  `json:"time"`
	Text     string `json:"text"`
}

// 获取客服聊天记录返回的数据结构
type RecordResponse struct {
	RecordList []RecordItem `json:"recordlist"`
}

// 聊天记录遍历器
type RecordIterator interface {
	HasNext() bool
	NextPage() ([]RecordItem, error)
}
