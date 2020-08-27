package store

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

// Get 获取基本信息
func Get(clt *core.Client) (store *model.Store, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/store/get_info?access_token="

	var result struct {
		core.Error
		Data *model.Store `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, nil, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	store = result.Data
	return
}
