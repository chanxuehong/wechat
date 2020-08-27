package category

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

// Get 获取类目详情
// catId: 父类目ID，可先填0获取根部类目
func Get(clt *core.Client, catId uint64) (categories []model.Category, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/category/get?access_token="

	req := map[string]uint64{
		"f_cat_id": catId,
	}

	var result struct {
		core.Error
		List []model.Category `json:"cat_list"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	categories = result.List
	return
}
