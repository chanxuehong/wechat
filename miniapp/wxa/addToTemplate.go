package wxa

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 将草稿箱的草稿选为小程序代码模版
func AddToTemplate(clt *core.Client, draftId uint64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/addtotemplate?access_token="
	var result struct {
		core.Error
	}
	req := map[string]uint64{"draft_id": draftId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
