package externalcontact

import (
	"github.com/chanxuehong/wechat/work/core"
)

type RemarkRequest struct {
	UserId           string   `json:"userid"`
	ExternalUserId   string   `json:"external_userid"`
	Remark           string   `json:"remark,omitempty"`
	Description      string   `json:"description,omitempty"`
	RemarkCompany    string   `json:"remark_company,omitempty"`
	RemarkMobiles    []string `json:"remark_mobiles,omitempty"`
	RemarkPicMediaId string   `json:"remark_pic_mediaid,omitempty"`
}

// Remark 修改客户备注信息。
// userid: 企业成员的userid
// external_userid: 群主过滤。如果不填，表示获取全部群主的数据
// remark: 此用户对外部联系人的备注;
// description: 此用户对外部联系人的描述;
// remark_company: 此用户对外部联系人备注的所属公司名称;
// remark_mobiles: 此用户对外部联系人备注的手机号;
// remark_pic_mediaid: 备注图片的mediaid;
func Remark(clt *core.Client, req *RemarkRequest) (err error) {
	const incompleteURL = "https://qyapi.weixin.qq.com/cgi-bin/externalcontact/remark?access_token="
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}
