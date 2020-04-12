package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 获取展示的公众号信息
func GetShowWxaItem(clt *core.Client) (info BizInfo, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/getshowwxaitem?access_token="
	var result struct {
		core.Error
		CanOpen  uint   `json:"can_open,omitempty"`
		IsOpen   uint   `json:"is_open,omimtempty"`
		NickName string `json:"nickname"`
		AppId    string `json:"appid"`
		HeadImg  string `json:"headimg"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return BizInfo{
		CanOpen:  result.CanOpen,
		IsOpen:   result.IsOpen,
		NickName: result.NickName,
		AppId:    result.AppId,
		HeadImg:  result.HeadImg,
	}, nil
}
