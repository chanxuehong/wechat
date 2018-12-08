package template

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 设置所属行业.
func SetIndustry(clt *core.Client, industryId1, industryId2 int64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/template/api_set_industry?access_token="

	var request = struct {
		IndustryId1 int64 `json:"industry_id1"`
		IndustryId2 int64 `json:"industry_id2"`
	}{
		IndustryId1: industryId1,
		IndustryId2: industryId2,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

type Industry struct {
	FirstClass  string `json:"first_class"`
	SecondClass string `json:"second_class"`
}

// 获取设置的行业信息
func GetIndustry(clt *core.Client) (primaryIndustry, secondaryIndustry Industry, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/template/get_industry?access_token="

	var result struct {
		core.Error
		PrimaryIndustry   Industry `json:"primary_industry"`
		SecondaryIndustry Industry `json:"secondary_industry"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	primaryIndustry = result.PrimaryIndustry
	secondaryIndustry = result.SecondaryIndustry
	return
}

// 从行业模板库选择模板添加到账号后台, 并返回模板id.
//  templateIdShort: 模板库中模板的编号, 有"TM**"和"OPENTMTM**"等形式.
func AddPrivateTemplate(clt *core.Client, templateIdShort string) (templateId string, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/template/api_add_template?access_token="

	var request = struct {
		TemplateIdShort string `json:"template_id_short"`
	}{
		TemplateIdShort: templateIdShort,
	}
	var result struct {
		core.Error
		TemplateId string `json:"template_id"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	templateId = result.TemplateId
	return
}

// 模板数据结构
//  {
//      "template_id": "iPk5sOIt5X_flOVKn5GrTFpncEYTojx6ddbt8WYoV5s",
//      "title": "领取奖金提醒",
//      "primary_industry": "IT科技",
//      "deputy_industry": "互联网|电子商务",
//      "content": "{ {result.DATA} }\n\n领奖金额:{ {withdrawMoney.DATA} }\n领奖  时间:{ {withdrawTime.DATA} }\n银行信息:{ {cardInfo.DATA} }\n到账时间:  { {arrivedTime.DATA} }\n{ {remark.DATA} }",
//      "example": "您已提交领奖申请\n\n领奖金额：xxxx元\n领奖时间：2013-10-10 12:22:22\n银行信息：xx银行(尾号xxxx)\n到账时间：预计xxxxxxx\n\n预计将于xxxx到达您的银行卡"
//  }
type Template struct {
	TemplateId      string `json:"template_id"`
	Title           string `json:"title"`
	PrimaryIndustry string `json:"primary_industry"`
	DeputyIndustry  string `json:"deputy_industry"`
	Content         string `json:"content"`
	Example         string `json:"example"`
}

// 获取模板列表
func GetAllPrivateTemplate(clt *core.Client) (templateList []Template, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token="

	var result struct {
		core.Error
		TemplateList []Template `json:"template_list"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	templateList = result.TemplateList
	return
}

// 删除模板.
func DeletePrivateTemplate(clt *core.Client, templateId string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/template/del_private_template?access_token="

	var request = struct {
		TemplateId string `json:"template_id"`
	}{
		TemplateId: templateId,
	}
	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
