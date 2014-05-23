package mass

type commonOpenIdMsgHead struct {
	ToUser  []string `json:"touser"`
	MsgType string   `json:"msgtype"`
}

func (msg *commonOpenIdMsgHead) AppendToUser(touser ...string) {
	if len(touser) <= 0 {
		return
	}
	if len(msg.ToUser) >= OpenIdMsgToUserCountLimit {
		return
	}

	if n := OpenIdMsgToUserCountLimit - len(msg.ToUser); len(touser) > n {
		touser = touser[:n]
	}

	msg.ToUser = append(msg.ToUser, touser...)
}

// news ========================================================================

type OpenIdNewsMsg struct {
	commonOpenIdMsgHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewOpenIdNewsMsg(touser []string, mediaId string) *OpenIdNewsMsg {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdNewsMsg
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}

// text ========================================================================

type OpenIdTextMsg struct {
	commonOpenIdMsgHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewOpenIdTextMsg(touser []string, content string) *OpenIdTextMsg {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdTextMsg
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// voice =======================================================================

type OpenIdVoiceMsg struct {
	commonOpenIdMsgHead

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewOpenIdVoiceMsg(touser []string, mediaId string) *OpenIdVoiceMsg {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdVoiceMsg
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// image =======================================================================

type OpenIdImageMsg struct {
	commonOpenIdMsgHead

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewOpenIdImageMsg(touser []string, mediaId string) *OpenIdImageMsg {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdImageMsg
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// video =======================================================================

type OpenIdVideoMsg struct {
	commonOpenIdMsgHead

	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"video"`
}

func NewOpenIdVideoMsg(touser []string, mediaId, title, description string) *OpenIdVideoMsg {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdVideoMsg
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}
