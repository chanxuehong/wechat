package user

import "gopkg.in/chanxuehong/wechat.v2/mp/core"

// 获取黑名单列表
func GetBlacklist(clt *core.Client, begin_openid string) (rslt *ListResult, err error){
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token="

	var request = struct {
		BeginOpenid string `json:"begin_openid"`
	}{
		BeginOpenid: begin_openid,
	}
	var result struct {
		core.Error
		ListResult
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	rslt = &result.ListResult
	return
}

// 批量将用户增加到黑名单中
func BatchAddBlacklist(clt *core.Client, openId_list []string) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token="

	var request = struct {
		OpenidList []string `json:"openid_list"`
	}{
		OpenidList: openId_list,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}

// 批量移出黑名单
func BatchMoveBlacklist(clt *core.Client, openId_list []string) (err error)  {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token="

	var request = struct {
		OpenidList []string `json:"openid_list"`
	}{
		OpenidList: openId_list,
	}
	var result struct {
		core.Error
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return
}
