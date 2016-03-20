<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package userinfo

import (
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/card/code"
=======
package userinfo

import (
	"github.com/chanxuehong/wechat/mp/card/code"
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

type CustomField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type UserInfo struct {
	openid          string        `json:"openid"`
	nickname        string        `json:"nickname"`
	sex             string        `json:"sex"`
	CustomFieldList []CustomField `json:"custom_field_list"`
}

// 拉取会员信息（积分查询）接口
<<<<<<< HEAD
func Get(clt *mp.Client, id *code.CardItemIdentifier) (info *UserInfo, err error) {
	var result struct {
		mp.Error
=======
func Get(clt *core.Client, id *code.CardItemIdentifier) (info *UserInfo, err error) {
	var result struct {
		core.Error
>>>>>>> github/v2
		UserInfo
	}

	incompleteURL := "https://api.weixin.qq.com/card/membercard/userinfo/get?access_token="
	if err = clt.PostJSON(incompleteURL, id, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result.Error
		return
	}
	info = &result.UserInfo
	return
}
