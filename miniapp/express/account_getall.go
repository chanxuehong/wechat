package express

import (
	"github.com/chanxuehong/wechat/component/core"
)

// Account 账号
type Account struct {
	// BizID 快递公司客户编码
	BizID string `json:"biz_id,omitempty"`
	// DeliveryID 快递公司ID
	DeliveryID string `json:"delivery_id,omitempty"`
	// CreateTime 账号绑定时间
	CreateTime int64 `json:"create_time,omitempty"`
	// UpdateTime 账号更新时间
	UpdateTime int64 `json:"update_time,omitempty"`
	// StatusCode 绑定状态
	StatusCode int `json:"status_code,omitempty"`
	// Alias 账号别名
	Alias string `json:"alias,omitempty"`
	// RemarkWrongMsg 账号绑定失败的错误信息（EMS审核结果）
	RemarkWrongMsg string `json:"remark_wrong_msg,omitempty"`
	// RemarkContent 账号绑定时的备注内容（提交 EMS 审核需要）
	RemarkContent string `json:"remark_content,omitempty"`
	// QuotaNum 电子面单余额
	QuotaNum int `json:"quota_num,omitempty"`
	// QuotaUpdateTime 电子面单余额更新时间
	QuotaUpdateTime int64 `json:"quota_update_time,omitempty"`
	// ServiceType 该绑定帐号支持的服务类型
	ServiceType []ServiceType `json:"service_type,omitempty"`
}

// AccountGetAll 获取所有绑定的物流账号
func AccountGetAll(clt *core.Client) (list []Account, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/express/business/account/getall?access_token="
	var result struct {
		core.Error
		// List 账号列表
		List []Account `json:"list,omitempty"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List
	return
}
