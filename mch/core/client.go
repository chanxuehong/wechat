package core

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/chanxuehong/util"

	"github.com/chanxuehong/wechat.v2/internal/debug/mch/api"
	wechatutil "github.com/chanxuehong/wechat.v2/util"
)

type Client struct {
	appId  string
	mchId  string
	apiKey string

	httpClient *http.Client
}

func (clt *Client) AppId() string {
	return clt.appId
}
func (clt *Client) MchId() string {
	return clt.mchId
}
func (clt *Client) ApiKey() string {
	return clt.apiKey
}

// NewClient 创建一个新的 Client.
//  如果 httpClient == nil 则默认用 util.DefaultHttpClient.
func NewClient(appId, mchId, apiKey string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = wechatutil.DefaultHttpClient
	}
	return &Client{
		appId:      appId,
		mchId:      mchId,
		apiKey:     apiKey,
		httpClient: httpClient,
	}
}

// PostXML 是微信支付通用请求方法.
//  err == nil 表示 (return_code == "SUCCESS" && result_code == "SUCCESS").
func (clt *Client) PostXML(url string, req map[string]string) (resp map[string]string, err error) {
	buffer := textBufferPool.Get().(*bytes.Buffer)
	buffer.Reset()
	defer textBufferPool.Put(buffer)

	if err = util.EncodeXMLFromMap(buffer, req, "xml"); err != nil {
		return
	}
	api.DebugPrintPostXMLRequest(url, buffer.Bytes())

	httpResp, err := clt.httpClient.Post(url, "text/xml; charset=utf-8", buffer)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	if resp, err = api.DecodeXMLHttpResponse(httpResp.Body); err != nil {
		return
	}

	// 判断协议状态
	returnCode, ok := resp["return_code"]
	if !ok {
		err = ErrNotFoundReturnCode
		return
	}
	if returnCode != ReturnCodeSuccess {
		err = &Error{
			ReturnCode: returnCode,
			ReturnMsg:  resp["return_msg"],
		}
		return
	}

	// 验证 appid 和 mch_id
	appId, ok := resp["appid"]
	if ok && appId != clt.appId {
		err = fmt.Errorf("appid mismatch, have: %s, want: %s", appId, clt.appId)
		return
	}
	mchId, ok := resp["mch_id"]
	if ok && mchId != clt.mchId {
		err = fmt.Errorf("mch_id mismatch, have: %s, want: %s", mchId, clt.mchId)
		return
	}

	// 验证签名
	signatureHave, ok := resp["sign"]
	if !ok {
		// TODO(chanxuehong): 在适当的时候更新下面的 case
		switch url {
		default:
			err = ErrNotFoundSign
			return
		case "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers":
			// do nothing
		case "https://api2.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers":
			// do nothing
		}
	} else {
		signatureWant := Sign(resp, clt.apiKey, nil)
		if signatureHave != signatureWant {
			err = fmt.Errorf("sign mismatch,\nhave: %s,\nwant: %s", signatureHave, signatureWant)
			return
		}
	}

	resultCode, ok := resp["result_code"]
	if ok && resultCode != ResultCodeSuccess {
		err = &BizError{
			ResultCode:  resultCode,
			ErrCode:     resp["err_code"],
			ErrCodeDesc: resp["err_code_des"],
		}
		return
	}
	return
}
