package custom

type CommonHead struct {
	ToUser  string `json:"touser"`  // 接收方帐号(OpenID)
	MsgType string `json:"msgtype"` // text, image, voice, video, music, news
}

// text ========================================================================

type Text struct {
	CommonHead

	Text struct {
		Content string `json:"content"` // 回复的消息内容(换行：在content中能够换行, 微信客户端就支持换行显示)
	} `json:"text"`
}

func NewText(to, content string) *Text {
	msg := Text{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_TEXT,
		},
	}
	msg.Text.Content = content

	return &msg
}

// image =======================================================================

type Image struct {
	CommonHead

	Image struct {
		MediaId string `json:"media_id"` // 通过上传多媒体文件, 得到的id
	} `json:"image"`
}

func NewImage(to, mediaId string) *Image {
	msg := Image{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_IMAGE,
		},
	}
	msg.Image.MediaId = mediaId

	return &msg
}

// voice =======================================================================

type Voice struct {
	CommonHead

	Voice struct {
		MediaId string `json:"media_id"` // 通过上传多媒体文件, 得到的id
	} `json:"voice"`
}

func NewVoice(to, mediaId string) *Voice {
	msg := Voice{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_VOICE,
		},
	}
	msg.Voice.MediaId = mediaId

	return &msg
}

// video =======================================================================

type Video struct {
	CommonHead

	Video struct {
		MediaId     string `json:"media_id"`              // 通过上传多媒体文件, 得到的id
		Title       string `json:"title,omitempty"`       // 视频消息的标题
		Description string `json:"description,omitempty"` // 视频消息的描述
	} `json:"video"`
}

// title, description 可以为 ""
func NewVideo(to, mediaId, title, description string) *Video {
	msg := Video{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_VIDEO,
		},
	}
	msg.Video.MediaId = mediaId
	msg.Video.Title = title
	msg.Video.Description = description

	return &msg
}

// music =======================================================================

type Music struct {
	CommonHead

	Music struct {
		Title        string `json:"title,omitempty"`       // 音乐标题
		Description  string `json:"description,omitempty"` // 音乐描述
		MusicURL     string `json:"musicurl"`              // 音乐链接
		HQMusicURL   string `json:"hqmusicurl"`            // 高质量音乐链接, WIFI环境优先使用该链接播放音乐
		ThumbMediaId string `json:"thumb_media_id"`        // 缩略图的媒体id, 通过上传多媒体文件, 得到的id
	} `json:"music"`
}

// title, description 可以为 ""
func NewMusic(to, thumbMediaId, musicURL, HQMusicURL, title, description string) *Music {
	msg := Music{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_MUSIC,
		},
	}
	msg.Music.ThumbMediaId = thumbMediaId
	msg.Music.MusicURL = musicURL
	msg.Music.HQMusicURL = HQMusicURL
	msg.Music.Title = title
	msg.Music.Description = description

	return &msg
}

// news ========================================================================

// 图文消息里的 item
type NewsArticle struct {
	Title       string `json:"title,omitempty"`       // 图文消息标题
	Description string `json:"description,omitempty"` // 图文消息描述
	URL         string `json:"url,omitempty"`         // 点击图文消息跳转链接
	PicURL      string `json:"picurl,omitempty"`      // 图文消息的图片链接，支持JPG、PNG格式，较好的效果为大图640*320，小图80*80
}

// 图文消息
type News struct {
	CommonHead

	News struct {
		Articles []*NewsArticle `json:"articles"` // 多条图文消息信息, 默认第一个item为大图, 注意, 如果图文数超过10, 则将会无响应
	} `json:"news"`
}

// NOTE: 如果图文消息数量大于微信的限制, 则把多余的截除.
func NewNews(to string, article ...*NewsArticle) *News {
	if len(article) > NewsArticleCountLimit {
		article = article[:NewsArticleCountLimit]
	} else if article == nil {
		article = make([]*NewsArticle, 0, NewsArticleCountLimit)
	}

	msg := News{
		CommonHead: CommonHead{
			ToUser:  to,
			MsgType: MSG_TYPE_NEWS,
		},
	}
	msg.News.Articles = article

	return &msg
}

// NOTE: 如果总的按钮数超过限制, 则截除多余的.
func (msg *News) AppendArticle(article ...*NewsArticle) {
	if len(article) <= 0 {
		return
	}

	switch n := NewsArticleCountLimit - len(msg.News.Articles); {
	case n > 0:
		if len(article) > n {
			article = article[:n]
		}
		msg.News.Articles = append(msg.News.Articles, article...)
	case n == 0:
	default: // n < 0
		msg.News.Articles = msg.News.Articles[:NewsArticleCountLimit]
	}
}
