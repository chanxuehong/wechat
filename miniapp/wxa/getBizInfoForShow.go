package wxa

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
)

type BizInfo struct {
	CanOpen  uint   `json:"can_open,omitempty"`
	IsOpen   uint   `json:"is_open,omimtempty"`
	NickName string `json:"nickname"`
	AppId    string `json:"appid"`
	HeadImg  string `json:"headimg"`
}

// 获取可以用来设置的公众号列表
func GetBizInfoForShow(clt *core.Client, page uint, num uint) (total uint, list []BizInfo, err error) {
	incompleteURL := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxamplinkforshow?page=%d&num=%d&access_token=", page, num)
	var result struct {
		core.Error
		Total       uint      `json:"total_num"`
		BizInfoList []BizInfo `json:"biz_info_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.Total, result.BizInfoList, nil
}
