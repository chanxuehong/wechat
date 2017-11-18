// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://gopkg.in/chanxuehong/wechat.v1 for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/v1/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package material

import (
	"gopkg.in/chanxuehong/wechat.v1/mp"
)

type VideoInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DownURL     string `json:"down_url"`
}

// 获取视频消息素材.
func (clt *Client) GetVideo(mediaId string) (info *VideoInfo, err error) {
	var request = struct {
		MediaId string `json:"media_id"`
	}{
		MediaId: mediaId,
	}

	var result struct {
		mp.Error
		VideoInfo
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/material/get_material?access_token="
	if err = ((*mp.Client)(clt)).PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.VideoInfo
	return
}
