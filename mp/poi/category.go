package poi

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 查询门店信息.
func GetWxCategory(clt *core.Client) (categoryList []string, err error) {
	var result struct {
		core.Error
		CategoryList []string `json:"category_list"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/api_getwxcategory?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	categoryList = result.CategoryList
	return
}
