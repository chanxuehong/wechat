package media

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// UploadVideo2 创建视频素材, 返回的素材一般用于群发消息.
//  mediaId:     通过 UploadVideo 上传视频文件得到
//  title:       标题, 可以为空
//  description: 描述, 可以为空
func UploadVideo2(clt *core.Client, mediaId, title, description string) (info *MediaInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/media/uploadvideo?access_token="

	var request = struct {
		MediaId     string `json:"media_id"`
		Title       string `json:"title,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		MediaId:     mediaId,
		Title:       title,
		Description: description,
	}
	var result struct {
		core.Error
		MediaInfo
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	info = &result.MediaInfo
	return
}
