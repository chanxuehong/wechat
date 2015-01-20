// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package datacube

type InterfaceBaseData struct {
	RefDate       string `json:"ref_date"`        // 数据的日期
	CallbackCount int    `json:"callback_count"`  // 通过服务器配置地址获得消息后，被动回复用户消息的次数
	FailCount     int    `json:"fail_count"`      // 上述动作的失败次数
	TotalTimeCost int64  `json:"total_time_cost"` // 总耗时，除以callback_count即为平均耗时
	MaxTimeCost   int    `json:"max_time_cost"`   // 最大耗时
}
type InterfaceSummaryData struct {
	InterfaceBaseData
}
type InterfaceSummaryHourData struct {
	InterfaceBaseData
	RefHour int `json:"ref_hour,omitempty"` // 数据的小时
}
