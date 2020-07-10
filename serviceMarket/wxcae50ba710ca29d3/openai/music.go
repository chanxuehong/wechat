package openai

type MusicAnsDetail struct {
	Code         int           `json:"code,omitempty"`
	Msg          string        `json:"msg,omitempty"`
	Feed         *MusicFeed    `json:"feed,omitempty"`
	PlayCommand  *PlayCommand  `json:"play_command,omitempty"`
	SpeakCommand *SpeakCommand `json:"speak_command,omitempty"`
}

type MusicFeed struct {
	ExtBuf string `json:"ext_buf,omitempty"`
}

type PlayCommand struct {
	CanPreGetFront    bool       `json:"can_pre_get_front,omitempty"`
	NeedGetHistory    bool       `json:"need_get_history,omitempty"`
	NeedPreGet        bool       `json:"need_pre_get,omitempty"`
	Recoverable       bool       `json:"recoverable,omitempty"`
	RepeatModeSupport int        `json:"repeat_mode_support,omitempty"`
	Type              int        `json:"type,omitempty"`
	Playlist          []PlayInfo `json:"play_list,omitempty"`
}

type PlayInfo struct {
	Id          string `json:"id,omitempty"`
	AlbumId     string `json:"album_id,omitempty"`
	AlbumName   string `json:"album_name,omitempty"`
	AlbumPicUrl string `json:"album_pic_url,omitempty"`
	Author      string `json:"author,omitempty"`
	Mid         string `json:"mid,omitempty"`
	Name        string `json:"name,omitempty"`
	Url         string `json:"url,omitempty"`
}

type SpeakCommand struct {
	Type int    `json:"speak_type,omitempty"`
	Text string `json:"text,omitempty"`
}
