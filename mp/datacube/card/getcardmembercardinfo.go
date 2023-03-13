package card

import (
	"github.com/bububa/wechat/mp/core"
)

// 会员卡数据
type MemberCardData struct {
	RefDate          string `json:"ref_date"`           // 日期信息, YYYY-MM-DD
	ViewCount        int    `json:"view_cnt"`           // 浏览次数
	ViewUser         int    `json:"view_user"`          // 浏览人数
	ReceiveCount     int    `json:"receive_cnt"`        // 领取次数
	ReceiveUser      int    `json:"receive_user"`       // 领取人数
	VerifyCount      int    `json:"verify_cnt"`         // 使用次数
	VerifyUser       int    `json:"verify_user"`        // 使用人数
	ActiveUser       int    `json:"active_user"`        // 激活人数
	TotalUser        int    `json:"total_user"`         // 有效会员总人数
	TotalReceiveUser int    `json:"total_receive_user"` // 历史领取会员卡总人数
}

// 拉取会员卡数据接口
func GetMemberCardInfo(clt *core.Client, req *Request) (list []MemberCardData, err error) {
	var result struct {
		core.Error
		List []MemberCardData `json:"list"`
	}

	incompleteURL := "https://api.weixin.qq.com/datacube/getcardmembercardinfo?access_token="
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
