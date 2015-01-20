// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     magicshui(shuiyuzhe@gmail.com)

package datacube

import (
	"github.com/chanxuehong/wechat/mp/datacube"
	"time"
)

func (c *Client) DataCubeGetUpstreamMsg(begin time.Time, end time.Time) (data []datacube.UpstreamMsgData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UpstreamMsgData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUpstreamMsgUrl(token)

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

func (c *Client) DataCubeGetUpstreamMsgHour(begin time.Time, end time.Time) (data []datacube.UpstreamMsgHourData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UpstreamMsgHourData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUpstreamMsgHourUrl(token)

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

func (c *Client) DataCubeGetUpstreamMsgWeek(begin time.Time, end time.Time) (data []datacube.UpstreamMsgWeekData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UpstreamMsgWeekData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUpstreamMsgWeekUrl(token)

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

func (c *Client) DataCubeGetUpstreamMsgMonth(begin time.Time, end time.Time) (data []datacube.UpstreamMsgMonthData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UpstreamMsgMonthData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUpstreamMsgMonthUrl(token)

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

func (c *Client) DataCubeGetUpstreamMsgDist(begin time.Time, end time.Time) (data []datacube.UpstreamMsgDistData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UpstreamMsgDistData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUpstreamMsgDistUrl(token)

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
func (c *Client) DataCubeGetUpstreamMsgDistWeek(begin time.Time, end time.Time) (data []datacube.UpstreamMsgDistWeekData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UpstreamMsgDistWeekData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUpstreamMsgDistWeekUrl(token)

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

func (c *Client) DataCubeGetUpstreamMsgDistMonth(begin time.Time, end time.Time) (data []datacube.UpstreamMsgDistMonthData, err error) {
	var request = struct {
		BeginDate string `json:"begin_date"`
		EndDate   string `json:"end_date"`
	}{
		BeginDate: begin.Format("2006-01-02"),
		EndDate:   end.Format("2006-01-02"),
	}
	var result struct {
		Error
		List []datacube.UpstreamMsgDistMonthData `json:"list"`
	}
	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url := dataCubeGetUpstreamMsgDistMonthUrl(token)

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
