// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package datacube

type UpstreamMsgData struct {
	RefDate    string `json:"ref_date"`
	UserSource int    `json:"user_source"`
	MsgType    int    `json:"msg_type"`
	MsgUser    int    `json:"msg_user"`
	MsgCount   int    `json:"msg_count"`
}

type UpstreamMsgHourData struct {
	UpstreamMsgData
	RefHour int `json:"ref_hour"`
}

type UpstreamMsgWeekData struct {
	UpstreamMsgData
}

type UpstreamMsgMonthData struct {
	UpstreamMsgData
}

type UpstreamMsgDistData struct {
	RefDate       string `json:"ref_date"`
	UserSource    int    `json:"user_source"`
	CountInterval int    `json:"count_interval"`
	MsgUser       int    `json:"msg_user"`
}

type UpstreamMsgDistWeekData struct {
	UpstreamMsgDistData
}
type UpstreamMsgDistMonthData struct {
	UpstreamMsgDistData
}
