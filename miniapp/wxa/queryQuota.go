package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type Quota struct {
	Rest         uint `json:"rest"`          // quota剩余值
	Limit        uint `json:"limit"`         // 当月分配quota
	SpeedupRest  uint `json:"speedup_rest"`  // 剩余加急次数
	SpeedupLimit uint `json:"speedup_limit"` // 当月分配加急次数
}

// 查询服务商的当月提审限额（quota）和加急次数
func QueryQuota(clt *core.Client) (*quota Quota, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/queryquota?access_token="
	var result struct {
		core.Error
        Quota
	}
	req := map[string]interface{}{}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return nil, err
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return nil, err
	}
	return &Quota {
        Rest: result.Rest,
        Limit: result.Limit,
        SpeedupRest: result.SpeedupRest,
        SpeedupLimit: result.SpeedupLimit,
    }, nil
}
