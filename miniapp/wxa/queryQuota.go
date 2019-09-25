package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 查询服务商的当月提审限额（quota）和加急次数
func QueryQuota(clt *core.Client) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/queryquota?access_token="
	var result struct {
		core.Error
	}
	req := map[string]interface{}{}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return nil
}
