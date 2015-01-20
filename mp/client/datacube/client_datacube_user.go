// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package datacube

import (
	"github.com/chanxuehong/wechat/mp/datacube"
	"time"
)

func (c *Client) DataCubeGetUserSummary(begin time.Time, end time.Time) (data []datacube.UserSummaryData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UserSummaryData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUserSummaryUrl(token)

	if err = c.postJSON(url, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		data = result.List
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return
	}
}

func (c *Client) DataCubeGetUserCulumate(begin time.Time, end time.Time) (data []datacube.UserCumulateData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UserCumulateData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUserCumulateUrl(token)
	if err = c.postJSON(url, &request, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		data = result.List
		return
	case errCodeInvalidCredential, errCodeTimeout:
		if !hasRetry {
			hasRetry = true

			if token, err = getNewToken(c.tokenService, token); err != nil {
				return
			}
			goto RETRY
		}
		fallthrough
	default:
		err = &result.Error
		return
	}
}
