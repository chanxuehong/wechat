package openai

type FmAnsDetail struct {
	Code         int           `json:"err_code,omitempty"`
	Msg          string        `json:"msg,omitempty"`
	AppInfo      *AppInfo      `json:"app_info,omitempty"`
	Feed         *FmFeed       `json:"feed,omitempty"`
	NodeName     string        `json:"node_name,omitempty"`
	SpeakCommand *SpeakCommand `json:"speak_command,omitempty"`
	PlayCommand  *PlayCommand  `json:"audio_play_command,omitempty"`
}

type AppInfo struct {
	AppId string `json:"app_id,omitempty"`
	Name  string `json:"name,omitempty"`
}

type FmFeed struct {
	Cover  string `json:"cover,omitempty"`
	ExtBuf string `json:"ext_buf,omitempty"`
}
