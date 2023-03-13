package delivery

import (
	"github.com/bububa/wechat/product/core"
	"github.com/bububa/wechat/product/model"
)

// GetFreightTemplate 获取运费模板
func GetFreightTemplate(clt *core.Client) (templates []model.FreightTemplate, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/delivery/get_freight_template?access_token="

	var result struct {
		core.Error
		List []model.FreightTemplate `json:"template_list"`
	}
	if err = clt.PostJSON(incompleteURL, nil, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	templates = result.List
	return
}
