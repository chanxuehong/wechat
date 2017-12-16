package mp

import "github.com/neugls/wechat/mp/core"

type defaultOpenAccessTokenServer struct {
	token string
}

func (as defaultOpenAccessTokenServer) Token() (token string, err error) {
	return as.token, nil
}

func (as defaultOpenAccessTokenServer) RefreshToken(currentToken string) (token string, err error) {
	return as.token, nil
}

func (as defaultOpenAccessTokenServer) IID01332E16DF5011E5A9D5A4DB30FED8E1() {}

//GetDefaultAccessTokenServer return the default access token of mp
func GetDefaultAccessTokenServer(mpAccessToken string) core.AccessTokenServer {
	return defaultOpenAccessTokenServer{mpAccessToken}
}
