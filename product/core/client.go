package core

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/chanxuehong/wechat/mp/core"
)

type Client struct {
	*core.Client
	ServiceId  uint64
	SpecificId string
}

// NewClient 创建一个新的 Client.
//  如果 clt == nil 则默认用 util.DefaultHttpClient
func NewClient(srv core.AccessTokenServer, clt *http.Client) *Client {
	return &Client{
		Client: core.NewClient(srv, clt),
	}
}

func NewClientWithToken(token string, clt *http.Client) *Client {
	return &Client{
		Client: core.NewClientWithToken(token, clt),
	}
}

func (clt *Client) SetService(serviceId uint64, specificId string) {
	clt.ServiceId = serviceId
	clt.SpecificId = specificId
}

func (clt *Client) unifyURL(incompleteURL string) (string, error) {
	parsedUrl, err := url.Parse(incompleteURL)
	if err != nil {
		return "", err
	}
	query := parsedUrl.Query()
	if clt.ServiceId > 0 {
		query.Set("service_id", strconv.FormatUint(clt.ServiceId, 10))
	}
	if clt.SpecificId != "" {
		query.Set("specific_id", clt.SpecificId)
	}
	var needAccessToken bool
	if _, found := query["access_token"]; found {
		needAccessToken = true
		query.Del("access_token")
	}
	queryCount := len(query)
	parsedUrl.RawQuery = query.Encode()
	parsedUrl.Fragment = ""
	requestUrl := parsedUrl.String()
	if needAccessToken {
		if queryCount > 0 {
			requestUrl = fmt.Sprintf("%s&access_token=", requestUrl)
		} else {
			requestUrl = fmt.Sprintf("%s?access_token=", requestUrl)
		}
	}
	return requestUrl, nil
}

func (clt *Client) GetJSON(incompleteURL string, response interface{}) (err error) {
	requestURL, err := clt.unifyURL(incompleteURL)
	if err != nil {
		return err
	}
	return clt.Client.GetJSON(requestURL, response)
}

func (clt *Client) PostJSON(incompleteURL string, request interface{}, response interface{}) (err error) {
	requestURL, err := clt.unifyURL(incompleteURL)
	if err != nil {
		return err
	}
	return clt.Client.PostJSON(requestURL, request, response)
}
