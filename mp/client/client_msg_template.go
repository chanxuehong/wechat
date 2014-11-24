// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/chanxuehong/wechat/mp/message/active/template"
)

// 发送模版消息
func (c *Client) MsgTemplateSend(msg *template.Msg) (msgid int64, err error) {
	if msg == nil {
		err = errors.New("msg == nil")
		return
	}

	var result struct {
		Error
		MsgId int64 `json:"msgid"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := messageTemplateSendURL(token)

	if err = c.postJSON(url_, msg, &result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		msgid = result.MsgId
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

// 发送模版消息.
//  对于某些用户, template.Msg 不能满足其需求, 所以提供了这个方法供其调用, 由用户自己封装 json格式 消息体!
func (c *Client) MsgTemplateSendRaw(msg []byte) (msgid int64, err error) {
	if len(msg) == 0 {
		err = errors.New("msg is empty")
		return
	}

	var result struct {
		Error
		MsgId int64 `json:"msgid"`
	}

	token, err := c.Token()
	if err != nil {
		return
	}

	hasRetry := false
RETRY:
	url_ := messageTemplateSendURL(token)

	httpResp, err := c.httpClient.Post(url_, "application/json; charset=utf-8", bytes.NewReader(msg))
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	if err = json.NewDecoder(httpResp.Body).Decode(&result); err != nil {
		return
	}

	switch result.ErrCode {
	case errCodeOK:
		msgid = result.MsgId
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
