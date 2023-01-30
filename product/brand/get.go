package brand

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

// Get 获取品牌列表
func Get(clt *core.Client) (brands []model.Brand, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/brand/get?access_token="

	var result struct {
		core.Error
		List []model.Brand `json:"brands"`
	}
	if err = clt.PostJSON(incompleteURL, nil, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	brands = result.List
	return
}
