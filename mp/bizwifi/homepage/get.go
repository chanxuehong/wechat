package homepage

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type Homepage struct {
	ShopId     int64  `json:"shop_id"`     // 门店ID
	TemplateId int64  `json:"template_id"` // 模板类型
	URL        string `json:"url"`         // 商家主页链接
}

func Get(clt *core.Client, shopId int64) (homepage *Homepage, err error) {
	request := struct {
		ShopId int64 `json:"shop_id"`
	}{
		ShopId: shopId,
	}

	var result struct {
		core.Error
		Homepage `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/bizwifi/homepage/get?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	homepage = &result.Homepage
	return
}
