package mass

type commonWithGroupId struct {
	Filter struct {
		GroupId string `json:"group_id"`
	} `json:"filter"`
	MsgType string `json:"msgtype"`
}

type NewsWithGroupId struct {
	commonWithGroupId

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewNewsWithGroupId(groupId, mediaId string) *NewsWithGroupId {
	var msg NewsWithGroupId
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_WITH_GROUPID_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}

type TextWithGroupId struct {
	commonWithGroupId

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewTextWithGroupId(groupId, content string) *TextWithGroupId {
	var msg TextWithGroupId
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_WITH_GROUPID_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

type VoiceWithGroupId struct {
	commonWithGroupId

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewVoiceWithGroupId(groupId, mediaId string) *VoiceWithGroupId {
	var msg VoiceWithGroupId
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_WITH_GROUPID_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

type ImageWithGroupId struct {
	commonWithGroupId

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImageWithGroupId(groupId, mediaId string) *ImageWithGroupId {
	var msg ImageWithGroupId
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_WITH_GROUPID_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

type VideoWithGroupId struct {
	commonWithGroupId

	Video struct {
		MediaId string `json:"media_id"`
	} `json:"mpvideo"`
}

func NewVideoWithGroupId(groupId, mediaId string) *VideoWithGroupId {
	var msg VideoWithGroupId
	msg.Filter.GroupId = groupId
	msg.MsgType = MSG_WITH_GROUPID_TYPE_VIDEO
	msg.Video.MediaId = mediaId

	return &msg
}

type commonWithOpenIdList struct {
	ToUser  []string `json:"touser"`
	MsgType string   `json:"msgtype"`
}

type NewsWithOpenIdList struct {
	commonWithOpenIdList

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewNewsWithOpenIdList(touser []string, mediaId string) *NewsWithOpenIdList {
	if touser == nil {
		touser = make([]string, 0, 16)
	}
	var msg NewsWithOpenIdList
	msg.ToUser = touser
	msg.MsgType = MSG_WITH_OPENIDLIST_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}

type TextWithOpenIdList struct {
	commonWithOpenIdList

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewTextWithOpenIdList(touser []string, content string) *TextWithOpenIdList {
	if touser == nil {
		touser = make([]string, 0, 16)
	}
	var msg TextWithOpenIdList
	msg.ToUser = touser
	msg.MsgType = MSG_WITH_OPENIDLIST_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

type VoiceWithOpenIdList struct {
	commonWithOpenIdList

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewVoiceWithOpenIdList(touser []string, mediaId string) *VoiceWithOpenIdList {
	if touser == nil {
		touser = make([]string, 0, 16)
	}
	var msg VoiceWithOpenIdList
	msg.ToUser = touser
	msg.MsgType = MSG_WITH_OPENIDLIST_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

type ImageWithOpenIdList struct {
	commonWithOpenIdList

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImageWithOpenIdList(touser []string, mediaId string) *ImageWithOpenIdList {
	if touser == nil {
		touser = make([]string, 0, 16)
	}
	var msg ImageWithOpenIdList
	msg.ToUser = touser
	msg.MsgType = MSG_WITH_OPENIDLIST_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

type VideoWithOpenIdList struct {
	commonWithOpenIdList

	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"video"`
}

func NewVideoWithOpenIdList(touser []string, mediaId, title, description string) *VideoWithOpenIdList {
	if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg VideoWithOpenIdList
	msg.ToUser = touser
	msg.MsgType = MSG_WITH_OPENIDLIST_TYPE_VIDEO
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}
