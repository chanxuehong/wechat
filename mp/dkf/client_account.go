// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package dkf

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/chanxuehong/wechat/mp"
)

// 客服基本信息
type KfInfo struct {
	Id           string `json:"kf_id,string"` // 客服工号
	Account      string `json:"kf_account"`   // 完整客服账号，格式为：账号前缀@公众号微信号
	Nickname     string `json:"kf_nick"`      // 客服昵称
	HeadImageURL string `json:"kf_headimg"`   // 客服头像
}

var ErrNoHeadImage = errors.New("没有头像")

// 获取客服图像的大小, 如果客服没有图像则返回 ErrNoHeadImage 错误.
func (info *KfInfo) HeadImageSize() (size int, err error) {
	HeadImageURL := info.HeadImageURL
	if HeadImageURL == "" {
		err = ErrNoHeadImage
		return
	}

	lastSlashIndex := strings.LastIndex(HeadImageURL, "/")
	if lastSlashIndex == -1 {
		err = fmt.Errorf("invalid HeadImageURL: %s", HeadImageURL)
		return
	}
	HeadImageIndex := lastSlashIndex + 1
	if HeadImageIndex == len(HeadImageURL) {
		err = fmt.Errorf("invalid HeadImageURL: %s", HeadImageURL)
		return
	}

	sizeStr := HeadImageURL[HeadImageIndex:]

	size64, err := strconv.ParseUint(sizeStr, 10, 64)
	if err != nil {
		err = fmt.Errorf("invalid HeadImageURL: %s", HeadImageURL)
		return
	}

	if size64 == 0 {
		size64 = 640
	}
	size = int(size64)
	return
}

// 获取客服基本信息.
func (clt Client) KfList() (KfList []KfInfo, err error) {
	var result struct {
		mp.Error
		KfList []KfInfo `json:"kf_list"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	KfList = result.KfList
	return
}

const (
	OnlineKfInfoStatusPC          = 1
	OnlineKfInfoStatusMobile      = 2
	OnlineKfInfoStatusPCAndMobile = 3
)

// 在线客服接待信息
type OnlineKfInfo struct {
	Id                  string `json:"kf_id,string"`  // 客服工号
	Account             string `json:"kf_account"`    // 完整客服账号，格式为：账号前缀@公众号微信号
	Status              int    `json:"status"`        // 客服在线状态 1：pc在线，2：手机在线。若pc和手机同时在线则为 1+2=3
	AutoAcceptThreshold int    `json:"auto_accept"`   // 客服设置的最大自动接入数
	AcceptingNumber     int    `json:"accepted_case"` // 客服当前正在接待的会话数
}

// 获取在线客服接待信息.
func (clt Client) OnlineKfList() (KfList []OnlineKfInfo, err error) {
	var result struct {
		mp.Error
		KfList []OnlineKfInfo `json:"kf_online_list"`
	}

	incompleteURL := "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result.Error
		return
	}
	KfList = result.KfList
	return
}

// 添加客服账号.
//  account:    完整客服账号，格式为：账号前缀@公众号微信号，账号前缀最多10个字符，必须是英文或者数字字符。
//  nickname:   客服昵称，最长6个汉字或12个英文字符
//  password:   客服账号登录密码
//  isPwdPlain: 标识 password 是否为明文格式, true 表示是明文密码, false 表示是密文密码.
func (clt Client) AddKfAccount(account, nickname, password string, isPwdPlain bool) (err error) {
	if isPwdPlain {
		md5Sum := md5.Sum([]byte(password))
		password = hex.EncodeToString(md5Sum[:])
	}

	request := struct {
		Account  string `json:"kf_account,omitempty"`
		Nickname string `json:"nickname,omitempty"`
		Password string `json:"password,omitempty"`
	}{
		Account:  account,
		Nickname: nickname,
		Password: password,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/customservice/kfaccount/add?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 设置客服信息
//  account:    完整客服账号，格式为：账号前缀@公众号微信号，账号前缀最多10个字符，必须是英文或者数字字符。
//  nickname:   客服昵称，最长6个汉字或12个英文字符
//  password:   客服账号登录密码
//  isPwdPlain: 标识 password 是否为明文格式, true 表示是明文密码, false 表示是密文密码.
func (clt Client) SetKfAccount(account, nickname, password string, isPwdPlain bool) (err error) {
	if isPwdPlain {
		md5Sum := md5.Sum([]byte(password))
		password = hex.EncodeToString(md5Sum[:])
	}

	request := struct {
		Account  string `json:"kf_account,omitempty"`
		Nickname string `json:"nickname,omitempty"`
		Password string `json:"password,omitempty"`
	}{
		Account:  account,
		Nickname: nickname,
		Password: password,
	}

	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/customservice/kfaccount/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 上传客服头像.
//  开发者可调用本接口来上传图片作为客服人员的头像，头像图片文件必须是jpg格式，推荐使用640*640大小的图片以达到最佳效果。
func (clt Client) UploadKfHeadImage(kfAccount, imagePath string) (err error) {
	if kfAccount == "" {
		return errors.New("empty kfAccount")
	}
	file, err := os.Open(imagePath)
	if err != nil {
		return
	}
	defer file.Close()

	return clt.uploadKfHeadImageFromReader(kfAccount, filepath.Base(imagePath), file)
}

// 上传客服头像.
//  开发者可调用本接口来上传图片作为客服人员的头像，头像图片文件必须是jpg格式，推荐使用640*640大小的图片以达到最佳效果。
//  注意参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func (clt Client) UploadKfHeadImageFromReader(kfAccount, filename string, reader io.Reader) (err error) {
	if kfAccount == "" {
		return errors.New("empty kfAccount")
	}
	if filename == "" {
		return errors.New("empty filename")
	}
	if reader == nil {
		return errors.New("nil reader")
	}

	return clt.uploadKfHeadImageFromReader(kfAccount, filename, reader)
}

// 上传客服头像.
//  注意参数 filename 不是文件路径, 是指定 multipart/form-data 里面文件名称
func (clt Client) uploadKfHeadImageFromReader(kfAccount, filename string, reader io.Reader) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?kf_account=" +
		url.QueryEscape(kfAccount) + "&access_token="
	if err = clt.UploadFromReader(incompleteURL, "media", filename, reader, "", nil, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 删除客服账号
func (clt Client) DeleteKfAccount(kfAccount string) (err error) {
	var result mp.Error

	incompleteURL := "https://api.weixin.qq.com/customservice/kfaccount/del?kf_account=" +
		url.QueryEscape(kfAccount) + "&access_token="
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}

	if result.ErrCode != mp.ErrCodeOK {
		err = &result
		return
	}
	return
}
