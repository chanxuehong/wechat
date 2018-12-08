package poi

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// CategoryList 获取门店类目表.
func CategoryList(clt *core.Client) (list []string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/api_getwxcategory?access_token="

	var result struct {
		core.Error
		CategoryList []string `json:"category_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.CategoryList
	return
}
