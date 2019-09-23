package wxa

import (
	"fmt"
	"github.com/chanxuehong/wechat/mp/core"
)

// 查询某个指定版本的审核状态（仅供第三方代小程序调用）
func GetAuditStatus(clt *core.Client, auditId uint64) (status uint, reason string, screenshot string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/get_auditstatus?access_token="
	var result struct {
		core.Error
		Status     uint   `json:"status"`     // 审核状态，其中0为审核成功，1为审核失败，2为审核中，3已撤回
		Reason     string `json:"reason"`     // 当status=1，审核被拒绝时，返回的拒绝原因
		Screenshot string `json:"screenshot"` // 当status=1，审核被拒绝时，会返回审核失败的小程序截图示例。 xxx丨yyy丨zzz是media_id可通过获取永久素材接口 拉取截图内容）
	}
	req := map[string]uint64{"auditid", auditId}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.Status, result.Reason, result.Screenshot, nil
}
