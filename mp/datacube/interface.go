package datacube

import (
	"errors"

	"github.com/chanxuehong/wechat/mp/core"
)

// 接口分析数据
type InterfaceSummaryData struct {
	RefDate       string `json:"ref_date"`        // 数据的日期, YYYY-MM-DD 格式
	CallbackCount int    `json:"callback_count"`  // 通过服务器配置地址获得消息后, 被动回复用户消息的次数
	FailCount     int    `json:"fail_count"`      // 上述动作的失败次数
	TotalTimeCost int64  `json:"total_time_cost"` // 总耗时, 除以callback_count即为平均耗时
	MaxTimeCost   int64  `json:"max_time_cost"`   // 最大耗时
}

// 获取接口分析数据.
func GetInterfaceSummary(clt *core.Client, req *Request) (list []InterfaceSummaryData, err error) {
	if req == nil {
		err = errors.New("nil Request")
		return
	}

	var result struct {
		core.Error
		List []InterfaceSummaryData `json:"list"`
	}

	incompleteURL := "https://api.weixin.qq.com/datacube/getinterfacesummary?access_token="
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List
	return
}

type InterfaceSummaryHourData struct {
	RefHour int `json:"ref_hour"` // 数据的小时, 包括从000到2300, 分别代表的是[000,100)到[2300,2400), 即每日的第1小时和最后1小时
	InterfaceSummaryData
}

// 获取接口分析分时数据.
func GetInterfaceSummaryHour(clt *core.Client, req *Request) (list []InterfaceSummaryHourData, err error) {
	if req == nil {
		err = errors.New("nil Request")
		return
	}

	var result struct {
		core.Error
		List []InterfaceSummaryHourData `json:"list"`
	}

	incompleteURL := "https://api.weixin.qq.com/datacube/getinterfacesummaryhour?access_token="
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List
	return
}
