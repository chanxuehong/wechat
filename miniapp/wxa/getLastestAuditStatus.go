package wxa

import (
	"strconv"

	"github.com/chanxuehong/wechat/mp/core"
)

// 查询最新一次提交的审核状态
func GetLatestAuditStatus(clt *core.Client) (auditId uint64, status uint, reason string, screenshot string, err error) {
	incompleteURL := "https://api.weixin.qq.com/wxa/get_auditstatus?access_token="
	var result struct {
		core.Error
		AuditId    string `json:"auditid"`    // 最新的审核ID
		Status     uint   `json:"status"`     // 审核状态，其中0为审核成功，1为审核失败，2为审核中，3已撤回
		Reason     string `json:"reason"`     // 当status=1，审核被拒绝时，返回的拒绝原因
		Screenshot string `json:"ScreenShot"` // 当status=1，审核被拒绝时，会返回审核失败的小程序截图示例。 xxx丨yyy丨zzz是media_id可通过获取永久素材接口 拉取截图内容）
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	auditId, _ = strconv.ParseUint(result.AuditId, 10, 64)
	return auditId, result.Status, result.Reason, result.Screenshot, nil
}
