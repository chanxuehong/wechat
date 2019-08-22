package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type SubmitAuditRequest struct {
	ItemList []struct {
		Address     string `json:"address"`      // 小程序的页面，可通过“获取小程序的第三方提交代码的页面配置”接口获得
		Tag         string `json:"tag"`          // 小程序的标签，多个标签用空格分隔，标签不能多于10个，标签长度不超过20
		FirstClass  string `json:"first_class"`  // 一级类目名称，可通过“获取授权小程序帐号的可选类目”接口获得
		SecondClass string `json:"second_class"` // 二级类目
		ThirdClass  string `json:"third_class"`  // 三级类目
		FirstId     uint   `json:"first_id"`     // 一级类目的ID，可通过“获取授权小程序帐号的可选类目”接口获得
		SecondId    uint   `json:"second_id"`    // 二级类目的ID
		ThirdId     uint   `json:"third_id"`     // 三级类目的ID
		Title       string `json:"title"`        // 小程序页面的标题,标题长度不超过32
	} `json:"item_list"` // 提交审核项的一个列表（至少填写1项，至多填写5项）
	FeedbackInfo  string `json:"feedback_info,omitempty"`  // 反馈内容，不超过200字
	FeedbackStuff string `json:"feedback_stuff,omitempty"` // 图片media_id列表，中间用“丨”分割，xx丨yy丨zz，不超过5张图片, 其中 media_id 可以通过新增临时素材接口上传而得到
}

// 将第三方提交的代码包提交审核（仅供第三方开发者代小程序调用）
func SubmitAudit(clt *core.Client, req *SubmitAuditRequest) (auditid uint64, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/submit_audit?access_token="
	var result struct {
		core.Error
		AuditId uint64 `json:"auditid"`
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.AuditId, nil
}
