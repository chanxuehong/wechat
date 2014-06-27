// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package massbyopenids

type CommonHead struct {
	// ToUser 的个数不能超过 ToUserCountLimit
	ToUser  []string `json:"touser,omitempty"`
	MsgType string   `json:"msgtype"`
}

// 如果总的按钮数超过限制, 则截除多余的.
func (msg *CommonHead) AppendUser(touser ...string) {
	if len(touser) <= 0 {
		return
	}

	switch n := ToUserCountLimit - len(msg.ToUser); {
	case n > 0:
		if len(touser) > n {
			touser = touser[:n]
		}
		msg.ToUser = append(msg.ToUser, touser...)
	case n == 0:
	default: // n < 0
		msg.ToUser = msg.ToUser[:ToUserCountLimit]
	}
}

// text ========================================================================

type Text struct {
	CommonHead

	Text struct {
		Content string `json:"content"`
	} `json:"text"`
}

func NewText(touser []string, content string) *Text {
	if len(touser) > ToUserCountLimit {
		touser = touser[:ToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg Text
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_TEXT
	msg.Text.Content = content

	return &msg
}

// image =======================================================================

type Image struct {
	CommonHead

	Image struct {
		MediaId string `json:"media_id"`
	} `json:"image"`
}

func NewImage(touser []string, mediaId string) *Image {
	if len(touser) > ToUserCountLimit {
		touser = touser[:ToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg Image
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_IMAGE
	msg.Image.MediaId = mediaId

	return &msg
}

// voice =======================================================================

type Voice struct {
	CommonHead

	Voice struct {
		MediaId string `json:"media_id"`
	} `json:"voice"`
}

func NewVoice(touser []string, mediaId string) *Voice {
	if len(touser) > ToUserCountLimit {
		touser = touser[:ToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg Voice
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_VOICE
	msg.Voice.MediaId = mediaId

	return &msg
}

// video =======================================================================

type Video struct {
	CommonHead

	Video struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	} `json:"video"`
}

func NewVideo(touser []string, mediaId, title, description string) *Video {
	if len(touser) > ToUserCountLimit {
		touser = touser[:ToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg Video
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_VIDEO
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}

// news ========================================================================

type News struct {
	CommonHead

	News struct {
		MediaId string `json:"media_id"`
	} `json:"mpnews"`
}

func NewNews(touser []string, mediaId string) *News {
	if len(touser) > ToUserCountLimit {
		touser = touser[:ToUserCountLimit]
	} else if touser == nil {
		touser = make([]string, 0, 16)
	}

	var msg News
	msg.ToUser = touser
	msg.MsgType = MSG_TYPE_NEWS
	msg.News.MediaId = mediaId

	return &msg
}
