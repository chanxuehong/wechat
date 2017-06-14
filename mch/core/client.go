package core

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
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

	subAppId string
	subMchId string

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

func (clt *Client) SubAppId() string {
	return clt.subAppId
}
func (clt *Client) SubMchId() string {
	return clt.subMchId
}

// NewClient 创建一个新的 Client.
//  appId:      必选; 公众号的 appid
//  mchId:      必选; 商户号 mch_id
//  apiKey:     必选; 商户的签名 key
//  httpClient: 可选; 默认使用 util.DefaultHttpClient
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

// NewSubMchClient 创建一个新的 Client.
//  appId:      必选; 公众号的 appid
//  mchId:      必选; 商户号 mch_id
//  apiKey:     必选; 商户的签名 key
//  subAppId:   可选; 公众号的 sub_appid
//  subMchId:   必选; 商户号 sub_mch_id
//  httpClient: 可选; 默认使用 util.DefaultHttpClient
func NewSubMchClient(appId, mchId, apiKey string, subAppId, subMchId string, httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = wechatutil.DefaultHttpClient
	}
	return &Client{
		appId:      appId,
		mchId:      mchId,
		apiKey:     apiKey,
		subAppId:   subAppId,
		subMchId:   subMchId,
		httpClient: httpClient,
	}
}

// PostXML 是微信支付通用请求方法.
//  err == nil 表示 (return_code == "SUCCESS" && result_code == "SUCCESS").
func (clt *Client) PostXML(url string, req map[string]string) (resp map[string]string, err error) {
	if req["appid"] == "" {
		req["appid"] = clt.appId
	}
	if req["mch_id"] == "" {
		req["mch_id"] = clt.mchId
	}
	if clt.subAppId != "" && req["sub_appid"] == "" {
		req["sub_appid"] = clt.subAppId
	}
	if clt.subMchId != "" && req["sub_mch_id"] == "" {
		req["sub_mch_id"] = clt.subMchId
	}

	// 获取请求参数的 sign_type 并检查其有效性
	var reqSignType string
	switch signType := req["sign_type"]; signType {
	case "", SignType_MD5:
		reqSignType = SignType_MD5
	case SignType_HMAC_SHA256:
		reqSignType = SignType_HMAC_SHA256
	default:
		err = fmt.Errorf("unsupported request sign_type: %s", signType)
		return nil, err
	}

	// 如果没有签名参数补全签名
	if req["sign"] == "" {
		switch reqSignType {
		case SignType_MD5:
			req["sign"] = Sign2(req, clt.ApiKey(), md5.New())
		case SignType_HMAC_SHA256:
			req["sign"] = Sign2(req, clt.ApiKey(), hmac.New(sha256.New, []byte(clt.ApiKey())))
		}
	}

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
	returnCode := resp["return_code"]
	if returnCode == "" {
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
	appId := resp["appid"]
	if appId != "" && appId != clt.appId {
		err = fmt.Errorf("appid mismatch, have: %s, want: %s", appId, clt.appId)
		return
	}
	mchId := resp["mch_id"]
	if mchId != "" && mchId != clt.mchId {
		err = fmt.Errorf("mch_id mismatch, have: %s, want: %s", mchId, clt.mchId)
		return
	}

	// 验证 sub_appid 和 sub_mch_id
	if clt.subAppId != "" {
		subAppId := resp["sub_appid"]
		if subAppId != "" && subAppId != clt.subAppId {
			err = fmt.Errorf("sub_appid mismatch, have: %s, want: %s", subAppId, clt.subAppId)
			return
		}
	}
	if clt.subMchId != "" {
		subMchId := resp["sub_mch_id"]
		if subMchId != "" && subMchId != clt.subMchId {
			err = fmt.Errorf("sub_mch_id mismatch, have: %s, want: %s", subMchId, clt.subMchId)
			return
		}
	}

	// 验证签名
	signatureHave := resp["sign"]
	if signatureHave == "" {
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
		// 获取返回参数的 sign_type 并检查其有效性
		var respSignType string
		switch signType := resp["sign_type"]; signType {
		case "":
			respSignType = reqSignType // 默认使用请求参数里的算法, 至少目前是这样
		case SignType_MD5:
			respSignType = SignType_MD5
		case SignType_HMAC_SHA256:
			respSignType = SignType_HMAC_SHA256
		default:
			err = fmt.Errorf("unsupported response sign_type: %s", signType)
			return
		}

		// 校验签名
		var signatureWant string
		switch respSignType {
		case SignType_MD5:
			signatureWant = Sign2(resp, clt.apiKey, md5.New())
		case SignType_HMAC_SHA256:
			signatureWant = Sign2(resp, clt.apiKey, hmac.New(sha256.New, []byte(clt.apiKey)))
		}
		if signatureHave != signatureWant {
			err = fmt.Errorf("sign mismatch,\nhave: %s,\nwant: %s", signatureHave, signatureWant)
			return
		}
	}

	resultCode := resp["result_code"]
	if resultCode != "" && resultCode != ResultCodeSuccess {
		err = &BizError{
			ResultCode:  resultCode,
			ErrCode:     resp["err_code"],
			ErrCodeDesc: resp["err_code_des"],
		}
		return
	}
	return
}
