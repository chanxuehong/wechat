package component

import (
	"github.com/chanxuehong/wechat/component/core"
)

// 该API用于在授权方令牌（authorizer_access_token）失效时，可用刷新令牌（authorizer_refresh_token）获取新的令牌。请注意，此处token是2小时刷新一次，开发者需要自行进行token的缓存，避免token的获取次数达到每日的限定额度。缓存方法可以参考：http://mp.weixin.qq.com/wiki/2/88b2bf1265a707c031e51f26ca5e6512.html
func RefreshAuthorizerToken(clt *core.Client, appId string, authorizerAppId string, authorizerRefreshToken string) (accessToken string, refreshToken string, expiresIn int64, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/component/api_authorizer_token?component_access_token="
	var result struct {
		core.Error
		AccessToken  string `json:"authorizer_access_token"`
		ExpiresIn    int64  `json:"expires_in"`
		RefreshToken string `json:"authorizer_refresh_token"`
	}
	req := map[string]string{"component_appid": appId, "authorizer_appid": authorizerAppId, "authorizer_refresh_token": authorizerRefreshToken}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.AccessToken, result.RefreshToken, result.ExpiresIn, nil
}
