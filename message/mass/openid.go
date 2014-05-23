package mass

type commonOpenIdMsgHead struct {
	ToUser  []string `json:"touser"`
	MsgType string   `json:"msgtype"`
}

func (msg *commonOpenIdMsgHead) AppendUser(touser ...string) {
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

type OpenIdNews struct {
	commonOpenIdMsgHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewOpenIdNews(touser []string, mediaId string) *OpenIdNews {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdNews
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}

// text ========================================================================

type OpenIdText struct {
	commonOpenIdMsgHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewOpenIdText(touser []string, content string) *OpenIdText {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdText
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// voice =======================================================================

type OpenIdVoice struct {
	commonOpenIdMsgHead

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewOpenIdVoice(touser []string, mediaId string) *OpenIdVoice {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdVoice
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// image =======================================================================

type OpenIdImage struct {
	commonOpenIdMsgHead

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewOpenIdImage(touser []string, mediaId string) *OpenIdImage {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdImage
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// video =======================================================================

type OpenIdVideo struct {
	commonOpenIdMsgHead

	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"video"`
}

func NewOpenIdVideo(touser []string, mediaId, title, description string) *OpenIdVideo {
	if len(touser) > OpenIdMsgToUserCountLimit {
		touser = touser[:OpenIdMsgToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg OpenIdVideo
	msg.ToUser = touser
	msg.MsgType = OPENID_MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}
