package groupmass

type commonMsgHead struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	MsgType string `json:"msgtype"`
}

// news ========================================================================

type NewsMsg struct {
	commonMsgHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewNewsMsg(groupId, mediaId string) *NewsMsg {
	var msg NewsMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}

// text ========================================================================

type TextMsg struct {
	commonMsgHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewTextMsg(groupId, content string) *TextMsg {
	var msg TextMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// voice =======================================================================

type VoiceMsg struct {
	commonMsgHead

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewVoiceMsg(groupId, mediaId string) *VoiceMsg {
	var msg VoiceMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// image =======================================================================

type ImageMsg struct {
	commonMsgHead

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImageMsg(groupId, mediaId string) *ImageMsg {
	var msg ImageMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// video =======================================================================

type VideoMsg struct {
	commonMsgHead

	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

func NewVideoMsg(groupId, mediaId string) *VideoMsg {
	var msg VideoMsg
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId

	return &msg
}
