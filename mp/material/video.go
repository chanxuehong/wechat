package material

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type VideoInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DownURL     string `json:"down_url"`
}

// 获取视频消息素材.
func GetVideo(clt *core.Client, mediaId string) (info *VideoInfo, err error) {
	var request = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}

	var result struct {
		core.Error
		VideoInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.VideoInfo
	return
}
