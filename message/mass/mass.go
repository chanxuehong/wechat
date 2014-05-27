package mass

// 删除群发
type DeleteMassRequest struct {
	MsgId int `json:"msgid"`
}

func NewDeleteMassRequest(msgid int) *DeleteMassRequest {
	return &DeleteMassRequest{
		MsgId: msgid,
	}
}
