package bot

// 消息类型，此时固定为text
type Text struct {
	Content             string   `json:"content"`                        // 文本内容，最长不超过2048个字节，必须是utf8编码
	MentionedList       []string `json:"mentioned_list,omitempty"`       // userid的列表，提醒群中的指定成员(@某个成员)，@all表示提醒所有人，如果开发者获取不到userid，可以使用mentioned_mobile_list
	MentionedMobileList []string `json:"metioned_mobile_list,omitempty"` // 手机号列表，提醒手机号对应的群成员(@某个成员)，@all表示提醒所有人
}

func NewText(content string, mentionedList []string, mentionedMobileList []string) *Message {
	return &Mesage{
		Type: TEXT,
		Text: &Text{
			Content:             content,
			MentionedList:       mentionedList,
			MentionedMobileList: mentionedMobileList,
		},
	}
}
