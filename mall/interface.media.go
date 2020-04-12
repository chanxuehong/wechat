package mall

type Media struct {
	ItemCode  string    `json:"item_code"`            // 媒体ID，需要保证唯一性
	Title     string    `json:"title"`                // 媒体名称，如歌曲名、视频标题
	Desc      string    `json:"desc,omitempty"`       // 媒体描述，如歌手名、视频描述
	ImageList []string  `json:"image_list"`           // 物品图片链接列表，图片宽度必须大于750px，宽高比建议4:3 - 1:1之间
	AppPath   string    `json:"src_wxapp_path"`       // 媒体来源小程序路径
	Info      MediaInfo `json:"media_info,omitempty"` // 媒体的播放信息
}

type MediaInfo struct {
	Type    int    `json:"type"`     // 区分视频还是音乐, 1代表音频类， 2代表视频类
	PlayUrl string `json:"play_url"` // 媒体的播放url
}
