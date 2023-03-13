package wxa

import (
	"strconv"

	"github.com/bububa/wechat/mp/core"
	"github.com/bububa/wechat/util"
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
	incompleteURL := util.StringsJoin("https://api.weixin.qq.com/wxa/getwxamplinkforshow?page=", strconv.Itoa(int(page)), "&num=", strconv.Itoa(int(num)), "&access_token=")
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
