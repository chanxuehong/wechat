package template

import (
    "github.com/chanxuehong/wechat/mp/core"
)

// 组合模板并添加至帐号下的个人模板库
func Add(clt *core.Client, templateId string, keywordIds []uint) (tplId string, err error) {
    const incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxopen/template/library/get?access_token="
    var result struct {
        core.Error
        Id       string    `json:"template_id"`
    }
    req := map[string]interface{"id": templateId, "keyword_id_list": keywordIds}
    if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
        return
    }
    if result.ErrCode != core.ErrCodeOK {
        err = &result.Error
        return
    }
    return result.Id, nil
}
