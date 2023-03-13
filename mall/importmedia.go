package mall

import (
	"github.com/bububa/wechat/mp/core"
)

type ImportMediaRequest struct {
	MediaList []Media `json:"media_list"`
}

// 更新或导入媒体信息
// 开发者可以对好物圈收藏/搜索场景下媒体(音频、视频)信息进行导入或更新
func ImportMedia(clt *core.Client, req *ImportMediaRequest) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/mall/importmedia?access_token="
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
