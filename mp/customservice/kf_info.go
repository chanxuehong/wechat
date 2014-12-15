// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package customservice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// 客服基本信息
type KFInfo struct {
	Id           string `json:"kf_id"`      // 客服工号
	Account      string `json:"kf_account"` // 客服账号@微信别名; 微信别名如有修改，旧账号返回旧的微信别名，新增的账号返回新的微信别名
	Nickname     string `json:"kf_nick"`    // 客服昵称
	HeadImageURL string `json:"kf_headimg"`
}

var ErrNoHeadImage = errors.New("没有图像")

// 获取用户图像的大小, 如果用户没有图像则返回 ErrNoHeadImage 错误.
func (this *KFInfo) HeadImageSize() (size int, err error) {
	HeadImageURL := this.HeadImageURL
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

	size64, err := strconv.ParseUint(sizeStr, 10, 8)
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

// 在线客服接待信息
type OnlineKFInfo struct {
	Id                  string `json:"kf_id"`         // 客服工号
	Account             string `json:"kf_account"`    // 客服账号@微信别名; 微信别名如有修改，旧账号返回旧的微信别名，新增的账号返回新的微信别名
	Status              int    `json:"status"`        // 客服在线状态 1：pc在线，2：手机在线, 若pc和手机同时在线则为 1+2=3
	AutoAcceptThreshold int    `json:"auto_accept"`   // 客服设置的最大自动接入数
	AcceptingNumber     int    `json:"accepted_case"` // 客服当前正在接待的会话数
}
