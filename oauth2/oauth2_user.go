// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package oauth2

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

// 获取用户信息(需scope为 snsapi_userinfo).
//  lang 可能的取值是 zh_CN, zh_TW, en; 如果留空 "" 则默认为 zh_CN.
func (c *Client) UserInfo(lang string) (info *UserInfo, err error) {
	switch lang {
	case "":
		lang = Language_zh_CN
	case Language_zh_CN, Language_zh_TW, Language_en:
	default:
		err = fmt.Errorf("lang 必须是 \"\",%s,%s,%s 之一",
			Language_zh_CN, Language_zh_TW, Language_en)
		return
	}

	tok := c.OAuth2Token
	if tok == nil {
		if c.TokenCache == nil {
			err = errors.New("OAuth2Token 和 TokenCache 都没有提供")
			return
		}
		if tok, err = c.TokenCache.Token(); err != nil {
			return
		}
	}
	if tok == nil {
		err = errors.New("没有有效的 OAuth2Token")
		return
	}
	if len(tok.AccessToken) == 0 {
		err = errors.New("没有有效的 AccessToken")
		return
	}
	if len(tok.OpenId) == 0 {
		err = errors.New("没有有效的 OpenId")
		return
	}

	c.OAuth2Token = tok // tok 有可能是从缓存读取的

	// Refresh the OAuth2Token if it has expired.
	if c.accessTokenExpired() {
		if _, err = c.TokenRefresh(); err != nil {
			return
		}
	}

	resp, err := c.httpClient().Get(userInfoURL(c.AccessToken, c.OpenId, lang))
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", resp.Status)
		return
	}

	var result struct {
		UserInfo
		Error
	}

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return
	}

	if result.ErrCode != 0 {
		err = &result.Error
		return
	}

	info = &result.UserInfo
	return
}
