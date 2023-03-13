package wxa

import (
	"github.com/bububa/wechat/mp/core"
)

// 获取体验者列表
func GetExperiencer(clt *core.Client) (list []string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/memberauth?access_token="
	var result struct {
		core.Error
		Members []*struct {
			Userstr string `json:"userstr"`
		} `json:"members"`
	}
	req := map[string]string{"action": "get_experiencer"}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	for _, m := range result.Members {
		list = append(list, m.Userstr)
	}
	return list, nil
}
