package delivery

import (
	"github.com/chanxuehong/wechat/product/core"
	"github.com/chanxuehong/wechat/product/model"
)

// GetCompanyList 获取快递公司列表
func GetCompanyList(clt *core.Client) (companies []model.DeliveryCompany, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/delivery/get_company_list?access_token="

	var result struct {
		core.Error
		Data struct {
			List []model.DeliveryCompany `json:"company_list"`
		} `json:"data"`
	}
	if err = clt.PostJSON(incompleteURL, nil, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	companies = result.Data.List
	return
}
