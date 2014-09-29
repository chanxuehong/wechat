// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"fmt"
	"strconv"
	"strings"
)

type UserInfo struct {
	OpenId   string `json:"openid"`   // 用户的唯一标识
	Nickname string `json:"nickname"` // 用户昵称
	Sex      int    `json:"sex"`      // 用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	City     string `json:"city"`     // 普通用户个人资料填写的城市
	Province string `json:"province"` // 用户个人资料填写的省份
	Country  string `json:"country"`  // 国家，如中国为CN

	// 用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），
	// 用户没有头像时该项为空
	HeadImageURL string `json:"headimgurl,omitempty"`

	// 用户特权信息，json 数组，如微信沃卡用户为（chinaunicom）
	Privilege []string `json:"privilege"`

	// 用户统一标识。针对一个微信开放平台帐号下的应用，同一用户的unionid是唯一的。
	UnionId string `json:"unionid"`
}

// 获取用户图像的大小
//  @headImageURL: 用户头像URL，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像）
//  NOTE: 请确保 headImageURL 不为空
func HeadImageSize(headImageURL string) (size int, err error) {
	index := strings.LastIndex(headImageURL, "/")
	if index == -1 {
		err = fmt.Errorf("invalid headImageURL: %s", headImageURL)
		return
	}
	if index+1 == len(headImageURL) { // "/" 在最后面
		err = fmt.Errorf("invalid headImageURL: %s", headImageURL)
		return
	}

	sizeStr := headImageURL[index+1:]

	size64, err := strconv.ParseUint(sizeStr, 10, 8)
	if err != nil {
		err = fmt.Errorf("invalid headImageURL: %s", headImageURL)
		return
	}

	if size64 == 0 {
		size64 = 640
	}

	size = int(size64)
	return
}
