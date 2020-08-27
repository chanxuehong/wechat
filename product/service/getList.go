package service

import (
	"github.com/chanxuehong/wechat/product/core"
)

// GetList 获取用户购买的在有效期内的服务列表
func GetList(clt *core.Client) (services []Service, err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/service/get_list?access_token="

	var result struct {
		core.Error
		List []Service `json:"service_list"`
	}

	if err = clt.PostJSON(incompleteURL, nil, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	services = result.List
	return
}
