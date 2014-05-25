package mass

// 群发成功时返回的结果
// {
//     "errcode":0,
//     "errmsg":"send job submission success",
//     "msg_id":34182
// }
type MassResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	MsgId   int    `json:"msg_id"`
}

// 删除群发
type DeleteMassRequest struct {
	MsgId int `json:"msgid"`
}
