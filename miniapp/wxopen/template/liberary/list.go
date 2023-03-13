package liberary

import (
	"github.com/bububa/wechat/mp/core"
)

// 获取小程序模板库标题列表
func List(clt *core.Client, offset uint, count uint) (total uint, list []Template, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/library/list?access_token="
	var result struct {
		core.Error
		List  []Template `json:"list"`
		Total uint       `json:"total_count"`
	}
	req := map[string]uint{"offset": offset, "count": count}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.Total, result.List, nil
}
