package device

import (
	"github.com/bububa/wechat/mp/core"
)

type ApplyStatus struct {
	ApplyTime    int64  `json:"apply_time"`    // 提交申请的时间戳
	AuditStatus  int    `json:"audit_status"`  // 审核状态。0：审核未通过、1：审核中、2：审核已通过；审核会在三个工作日内完成
	AuditComment string `json:"audit_comment"` // 审核备注，包括审核不通过的原因
	AuditTime    int64  `json:"audit_time"`    // 确定审核结果的时间戳，若状态为审核中，则该时间值为0
}

// 查询设备ID申请审核状态
func GetApplyStatus(clt *core.Client, applyId int64) (status *ApplyStatus, err error) {
	request := struct {
		ApplyId int64 `json:"apply_id"`
	}{
		ApplyId: applyId,
	}

	var result struct {
		core.Error
		ApplyStatus `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/applystatus?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	status = &result.ApplyStatus
	return
}
