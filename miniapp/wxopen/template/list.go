package template

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 获取帐号下已存在的模板列表
func List(clt *core.Client, offset uint, count uint) (list []Template, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/list?access_token="
	var result struct {
		core.Error
		List []Template `json:"list"`
	}
	req := map[string]uint{"offset": offset, "count": count}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.List, nil
}
