// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package poi

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp"
)

type PoiBrief struct {
	BaseInfo struct {
		PoiId          string `json:"poi_id,omitempty"`        // Poi 的id, 只有审核通过后才有
		AvailableState int    `json:"available_state"`         // 门店是否可用状态。1 表示系统错误、2 表示审核中、3 审核通过、4 审核驳回。当该字段为1、2、4 状态时，poi_id 为空
		Sid            string `json:"sid,omitempty"`           // 商户自己的id，用于后续审核通过收到poi_id 的通知时，做对应关系。请商户自己保证唯一识别性
		BusinessName   string `json:"business_name,omitempty"` // 门店名称（仅为商户名，如：国美、麦当劳，不应包含地区、店号等信息，错误示例：北京国美）
		BranchName     string `json:"branch_name,omitempty"`   // 分店名称（不应包含地区信息、不应与门店名重复，错误示例：北京王府井店）
		Address        string `json:"address,omitempty"`       // 门店所在的详细街道地址（不要填写省市信息）
	} `json:"base_info"`
}

// 查询门店列表.
//  begin: 开始位置，0 即为从第一条开始查询
//  limit: 返回数据条数，最大允许50，默认为20
func (clt Client) PoiList(begin, limit int) (list []PoiBrief, totalCount int, err error) {
	if begin < 0 {
		err = fmt.Errorf("invalid begin: %d", begin)
		return
	}
	if limit < 0 {
		err = fmt.Errorf("invalid limit: %d", limit)
		return
	}

	var request = struct {
		Begin int `json:"begin"`
		Limit int `json:"limit,omitempty"`
	}{
		Begin: begin,
		Limit: limit,
	}

	var result struct {
		mp.Error
		PoiList    []PoiBrief `json:"business_list"`
		TotalCount int        `json:"total_count"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/poi/getpoilist?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.PoiList
	totalCount = result.TotalCount
	return
}
