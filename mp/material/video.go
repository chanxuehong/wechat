package material

import (
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

type Video struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DownloadURL string `json:"down_url"`
}

// 获取视频消息素材信息.
func GetVideo(clt *core.Client, mediaId string) (info *Video, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token="

	var request = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}
	var result struct {
		core.Error
		Video
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.Video
	return
}
