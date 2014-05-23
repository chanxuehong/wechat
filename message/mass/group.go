package mass

type commonGroupMsgHead struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	MsgType string `json:"msgtype"`
}

// news ========================================================================

type GroupNews struct {
	commonGroupMsgHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewGroupNews(groupId, mediaId string) *GroupNews {
	var msg GroupNews
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}

// text ========================================================================

type GroupText struct {
	commonGroupMsgHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewGroupText(groupId, content string) *GroupText {
	var msg GroupText
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// voice =======================================================================

type GroupVoice struct {
	commonGroupMsgHead

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewGroupVoice(groupId, mediaId string) *GroupVoice {
	var msg GroupVoice
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// image =======================================================================

type GroupImage struct {
	commonGroupMsgHead

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewGroupImage(groupId, mediaId string) *GroupImage {
	var msg GroupImage
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// video =======================================================================

type GroupVideo struct {
	commonGroupMsgHead

	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

func NewGroupVideo(groupId, mediaId string) *GroupVideo {
	var msg GroupVideo
	msg.Filter.GroupId = groupId
	msg.MsgType = GROUP_MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId

	return &msg
}
