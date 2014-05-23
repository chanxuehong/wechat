package wechat

type AccessToken struct {
	Value     string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
}

func GetAccessToken(appid, secret string) (*AccessToken, error) {
	type response struct {
		AccessToken
		Error
	}
	return nil, nil
}
