package open

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"gopkg.in/chanxuehong/wechat.v2/util"
)

//GetComponentAccessToken return component access token
func GetComponentAccessToken(componentAppID, componentAppSecret, componentVerifyTicket string, token IDurationToken) (string, error) {
	if !token.Expired() {
		return token.Value()
	}

	httpClient := util.DefaultHttpClient
	postData := struct {
		ComponentAppID        string `json:"component_appid"`
		ComponentAppSecret    string `json:"component_appsecret"`
		ComponentVerifyTicket string `json:"component_verify_ticket"`
	}{
		componentAppID,
		componentAppSecret,
		componentVerifyTicket,
	}

	postDataJSON, _ := json.Marshal(postData)

	resp, err := httpClient.Post("https://api.weixin.qq.com/cgi-bin/component/api_component_token", "application/json", strings.NewReader(string(postDataJSON)))
	if err != nil {
		return "", err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bytes))

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", resp.Status)
		return "", err
	}

	accToken := struct {
		ComponentAccessToken string `json:"component_access_token"`
		ExpiresIn            int    `json:"expires_in"`
	}{}

	if e := json.Unmarshal(bytes, &accToken); e != nil {
		return "", e
	}
	if e := token.Put(accToken.ComponentAccessToken, time.Second*60); e != nil {
		return "", e
	}

	return accToken.ComponentAccessToken, nil
}
