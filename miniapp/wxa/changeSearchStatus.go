package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type SearchStatus = uint

const (
	ENABLE_SEARCHSTATUS  = 0
	DISABLE_SEARCHSTATUS = 1
)

// 设置小程序隐私设置（是否可被搜索）
func ChangeSearchStatus(clt *core.Client, status SearchStatus) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/changewxasearchstatus?access_token="
	var result struct {
		core.Error
	}
	req := map[string]SearchStatus{"status": status}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
