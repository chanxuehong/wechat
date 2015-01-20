// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package datacube

import (
	"github.com/chanxuehong/wechat/mp/datacube"
	"time"
)

func (c *Client) DataCubeGetArticleSummary(begin time.Time, end time.Time) (data []datacube.ArticleSummaryData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.ArticleSummaryData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetArticleSummaryUrl(token)

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

func (c *Client) DataCubeGetArticleTotal(begin time.Time, end time.Time) (data []datacube.ArticleTotalData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.ArticleTotalData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetArticleTotalUrl(token)

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

func (c *Client) DataCubeGetUserRead(begin time.Time, end time.Time) (data []datacube.UserReadData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UserReadData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUserReadUrl(token)

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

func (c *Client) DataCubeGetUserReadHour(begin time.Time, end time.Time) (data []datacube.UserReadHourData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UserReadHourData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUserReadHourUrl(token)

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

func (c *Client) DataCubeGetUserShare(begin time.Time, end time.Time) (data []datacube.UserShareData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UserShareData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUserShareUrl(token)

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

func (c *Client) DataCubeGetUserShareHour(begin time.Time, end time.Time) (data []datacube.UserShareHourData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UserShareHourData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUserShareHourUrl(token)

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
