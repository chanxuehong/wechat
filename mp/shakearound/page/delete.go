package page

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 删除页面
func Delete(clt *core.Client, pageIds []int64) (err error) {
	request := struct {
		PageIds []int64 `json:"page_ids,omitempty"`
	}{
		PageIds: pageIds,
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/shakearound/page/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
