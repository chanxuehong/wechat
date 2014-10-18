// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"strconv"
)

// https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=ACCESS_TOKEN
func _DepartmentCreateURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/department/create?access_token=" +
		accesstoken
}

// https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=ACCESS_TOKEN
func _DepartmentUpdateURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/department/update?access_token=" +
		accesstoken
}

// https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=ACCESS_TOKEN&id=2
func _DepartmentDeleteURL(accesstoken string, id int64) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/department/delete?access_token=" +
		accesstoken + "&id=" + strconv.FormatInt(id, 10)
}

// https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=ACCESS_TOKEN
func _DepartmentListURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/department/list?access_token=" +
		accesstoken
}

// https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=ACCESS_TOKEN
func _UserCreateURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/user/create?access_token=" +
		accesstoken
}

// https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=ACCESS_TOKEN
func _UserUpdateURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/user/update?access_token=" +
		accesstoken
}

// https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=ACCESS_TOKEN&userid=lisi
func _UserDeleteURL(accesstoken string, userid string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/user/delete?access_token=" +
		accesstoken + "&userid=" + userid
}

// https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=ACCESS_TOKEN&userid=lisi
func _UserGetURL(accesstoken string, userid string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=" +
		accesstoken + "&userid=" + userid
}

// https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=ACCESS_TOKEN&department_id=1&fetch_child=0&status=0
func _UserSimpleListURL(accesstoken string, departmentId int64, fetchChild bool, status int) string {
	var fetchChildStr string
	if fetchChild {
		fetchChildStr = "&fetch_child=1&status="
	} else {
		fetchChildStr = "&fetch_child=0&status="
	}

	return "https://qyapi.weixin.qq.com/cgi-bin/user/simplelist?access_token=" + accesstoken +
		"&department_id=" + strconv.FormatInt(departmentId, 10) +
		fetchChildStr + strconv.FormatInt(int64(status), 10)
}

// https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=ACCESS_TOKEN
func _TagCreateURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/tag/create?access_token=" +
		accesstoken
}

// https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=ACCESS_TOKEN
func _TagUpdateURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/tag/update?access_token=" +
		accesstoken
}

// https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=ACCESS_TOKEN&tagid=1
func _TagDeleteURL(accesstoken string, tagid int64) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/tag/delete?access_token=" +
		accesstoken + "&tagid=" + strconv.FormatInt(tagid, 10)
}

// https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=ACCESS_TOKEN&tagid=1
func _TagUserListURL(accesstoken string, tagid int64) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/tag/get?access_token=" +
		accesstoken + "&tagid=" + strconv.FormatInt(tagid, 10)
}

// https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=ACCESS_TOKEN
func _TagUserAddURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/tag/addtagusers?access_token=" +
		accesstoken
}

// https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=ACCESS_TOKEN
func _TagUserDeleteURL(accesstoken string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/tag/deltagusers?access_token=" +
		accesstoken
}
