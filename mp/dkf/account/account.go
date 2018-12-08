// 客户账号管理
package account

import (
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/chanxuehong/wechat/mp/core"
)

// Add 添加客服账号.
//  account:         完整客服账号，格式为：账号前缀@公众号微信号，账号前缀最多10个字符，必须是英文或者数字字符。
//  nickname:        客服昵称，最长6个汉字或12个英文字符
//  password:        客服账号登录密码
//  isPasswordPlain: 标识 password 是否为明文格式, true 表示是明文密码, false 表示是密文密码.
func Add(clt *core.Client, account, nickname, password string, isPasswordPlain bool) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/customservice/kfaccount/add?access_token="

	if password == "" {
		return errors.New("empty password")
	}
	if isPasswordPlain {
		md5Sum := md5.Sum([]byte(password))
		password = hex.EncodeToString(md5Sum[:])
	}

	request := struct {
		Account  string `json:"kf_account"`
		Nickname string `json:"nickname"`
		Password string `json:"password"`
	}{
		Account:  account,
		Nickname: nickname,
		Password: password,
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

// Update 设置客服信息(增量更新, 不更新的可以留空).
//  account:         完整客服账号，格式为：账号前缀@公众号微信号
//  nickname:        客服昵称，最长6个汉字或12个英文字符
//  password:        客服账号登录密码
//  isPasswordPlain: 标识 password 是否为明文格式, true 表示是明文密码, false 表示是密文密码.
func Update(clt *core.Client, account, nickname, password string, isPasswordPlain bool) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/customservice/kfaccount/update?access_token="

	if isPasswordPlain && password != "" {
		md5Sum := md5.Sum([]byte(password))
		password = hex.EncodeToString(md5Sum[:])
	}

	request := struct {
		Account  string `json:"kf_account"`
		Nickname string `json:"nickname,omitempty"`
		Password string `json:"password,omitempty"`
	}{
		Account:  account,
		Nickname: nickname,
		Password: password,
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

// Delete 删除客服账号
func Delete(clt *core.Client, kfAccount string) (err error) {
	// TODO
	//	incompleteURL := "https://api.weixin.qq.com/customservice/kfaccount/del?kf_account=" +
	//		url.QueryEscape(kfAccount) + "&access_token="
	incompleteURL := "https://api.weixin.qq.com/customservice/kfaccount/del?kf_account=" +
		kfAccount + "&access_token="

	var result core.Error
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
