package message

import (
	"github.com/chanxuehong/wechat/work/core"
)

type SendRequest struct {
	ToUser                 string    `json:"touser,omitempty"`
	ToParty                string    `json:"toparty,omitempty"`
	ToTag                  string    `json:"totag,omitempty"`
	MsgType                string    `json:"msgtype,omitempty"`
	AgentId                uint64    `json:"agentid,omitempty"`
	Content                string    `json:"content,omitempty"`
	Image                  *Message  `json:"image,omitempty"`
	Voice                  *Message  `json:"voice,omitempty"`
	Video                  *Message  `json:"video,omitempty"`
	File                   *Message  `json:"file,omitempty"`
	TextCard               *Message  `json:"textcard,omitempty"`
	News                   *News     `json:"news,omitempty"`
	MpNews                 *News     `json:"mpnews,omitempty"`
	Markdown               *Markdown `json:"markdown,omitempty"`
	MiniProgramNotice      *Message  `json:"miniprogram_notice,omitempty"`
	Safe                   int       `json:"safe"`
	EnableIdTrans          int       `json:"enable_id_trans,omitempty"`
	EnableDuplicateCheck   int       `json:"enable_duplicate_check,omitempty"`
	DuplicateCheckInterval int64     `json:"duplicate_check_interval,omitempty"`
}

type News struct {
	Articles []Message `json:"articles,omitempty"`
}

type Markdown struct {
	Content string `json:"content,omitempty"`
}

type Message struct {
	Title             string        `json:"title,omitempty"`
	Description       string        `json:"description,omitempty"`
	MediaId           string        `json:"media_id,omitempty"`
	Url               string        `json:"url,omitempty"`
	BtnText           string        `json:"btntext,omitempty"`
	PicUrl            string        `json:"picurl,omitempty"`
	ThumbMediaId      string        `json:"thumb_media_id,omitempty"`
	Author            string        `json:"author,omitempty"`
	ContentSourceUrl  string        `json:"content_source_url,omitempty"`
	Digest            string        `json:"digest,omitempty"`
	AppId             string        `json:"appid,omitempty"`
	Page              string        `json:"page,omitempty"`
	EmphasisFirstItem bool          `json:"emphasis_first_item,omitempty"`
	ContentItem       []ContentItem `json:"content_item,omitempty"`
	TaskId            string        `json:"task_id,omitempty"`
	Btns              []Btn         `json:"btn,omitempty"`
}

type ContentItem struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type Btn struct {
	Key         string `json:"key,omitempty"`
	Name        string `json:"name,omitempty"`
	ReplaceName string `json:"replace_name,omitempty"`
	Color       string `json:"color"`
	IsBold      bool   `json:"is_bold"`
}

type SendResponse struct {
	core.Error
	InvalidUser  string `json:"invaliduser,omitempty"`
	InvalidParty string `json:"invalidparty,omitempty"`
	InvalidTag   string `json:"invalidtag,omitempty"`
}

func Send(clt *core.Client, req *SendRequest) (err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/message/send?access_token="

	var result SendResponse

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}
