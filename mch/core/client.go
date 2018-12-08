package core

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/chanxuehong/util"

	"github.com/chanxuehong/wechat/internal/debug/mch/api"
	wechatutil "github.com/chanxuehong/wechat/util"
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
	switch url {
	case "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers", "https://api2.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers", // 企业付款
		"https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", "https://api2.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", // 发放普通红包
		"https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack", "https://api2.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack": // 发放裂变红包
		// TODO(chanxuehong): 这几个接口没有标准的 appid 和 mch_id 字段，需要用户在 req 里填写全部参数
		// TODO(chanxuehong): 通读整个支付文档, 可以的话重新考虑逻辑
	default:
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
	}

	// 获取请求参数的 sign_type 并检查其有效性
	var reqSignType string
	switch signType := req["sign_type"]; signType {
	case "", SignType_MD5:
		reqSignType = SignType_MD5
	case SignType_HMAC_SHA256:
		reqSignType = SignType_HMAC_SHA256
	default:
		return nil, fmt.Errorf("unsupported request sign_type: %s", signType)
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
		return nil, err
	}
	body := buffer.Bytes()

	hasRetried := false
RETRY:
	resp, needRetry, err := clt.postXML(url, body, reqSignType)
	if err != nil {
		if needRetry && !hasRetried {
			// TODO(chanxuehong): 打印错误日志
			hasRetried = true
			url = switchRequestURL(url)
			goto RETRY
		}
		return nil, err
	}
	return resp, nil
}

func (clt *Client) postXML(url string, body []byte, reqSignType string) (resp map[string]string, needRetry bool, err error) {
	api.DebugPrintPostXMLRequest(url, body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, false, err
	}
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	req = req.WithContext(ctx)
	httpResp, err := clt.httpClient.Do(req)
	if err != nil {
		return nil, true, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		return nil, true, fmt.Errorf("http.Status: %s", httpResp.Status)
	}

	resp, err = api.DecodeXMLHttpResponse(httpResp.Body)
	if err != nil {
		return nil, false, err
	}

	// 判断协议状态
	returnCode := resp["return_code"]
	if returnCode == "" {
		return nil, false, ErrNotFoundReturnCode
	}
	if returnCode != ReturnCodeSuccess {
		return nil, false, &Error{
			ReturnCode: returnCode,
			ReturnMsg:  resp["return_msg"],
		}
	}

	// 验证 appid 和 mch_id
	appId := resp["appid"]
	if appId != "" && appId != clt.appId {
		return nil, false, fmt.Errorf("appid mismatch, have: %s, want: %s", appId, clt.appId)
	}
	mchId := resp["mch_id"]
	if mchId != "" && mchId != clt.mchId {
		return nil, false, fmt.Errorf("mch_id mismatch, have: %s, want: %s", mchId, clt.mchId)
	}

	// 验证 sub_appid 和 sub_mch_id
	if clt.subAppId != "" {
		subAppId := resp["sub_appid"]
		if subAppId != "" && subAppId != clt.subAppId {
			return nil, false, fmt.Errorf("sub_appid mismatch, have: %s, want: %s", subAppId, clt.subAppId)
		}
	}
	if clt.subMchId != "" {
		subMchId := resp["sub_mch_id"]
		if subMchId != "" && subMchId != clt.subMchId {
			return nil, false, fmt.Errorf("sub_mch_id mismatch, have: %s, want: %s", subMchId, clt.subMchId)
		}
	}

	// 验证签名
	signatureHave := resp["sign"]
	if signatureHave == "" {
		// TODO(chanxuehong): 在适当的时候更新下面的 case
		switch url {
		default:
			return nil, false, ErrNotFoundSign
		case "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers", "https://api2.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers":
			// do nothing
		case "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo", "https://api2.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo":
		// do nothing
		case "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendredpack", "https://api2.mch.weixin.qq.com/mmpaymkttransfers/sendredpack":
			// do nothing
		case "https://api.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack", "https://api2.mch.weixin.qq.com/mmpaymkttransfers/sendgroupredpack":
			// do nothing
		case "https://api.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo", "https://api2.mch.weixin.qq.com/mmpaymkttransfers/gethbinfo":
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
			return nil, false, err
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
			return nil, false, fmt.Errorf("sign mismatch,\nhave: %s,\nwant: %s", signatureHave, signatureWant)
		}
	}

	resultCode := resp["result_code"]
	if resultCode != "" && resultCode != ResultCodeSuccess {
		errCode := resp["err_code"]
		if errCode == "SYSTEMERROR" {
			return nil, true, &BizError{
				ResultCode:  resultCode,
				ErrCode:     errCode,
				ErrCodeDesc: resp["err_code_des"],
			}
		}
		return nil, false, &BizError{
			ResultCode:  resultCode,
			ErrCode:     errCode,
			ErrCodeDesc: resp["err_code_des"],
		}
	}
	return resp, false, nil
}

func switchRequestURL(url string) string {
	switch {
	case strings.HasPrefix(url, "https://api.mch.weixin.qq.com/"):
		return "https://api2.mch.weixin.qq.com/" + url[len("https://api.mch.weixin.qq.com/"):]
	case strings.HasPrefix(url, "https://api2.mch.weixin.qq.com/"):
		return "https://api.mch.weixin.qq.com/" + url[len("https://api2.mch.weixin.qq.com/"):]
	default:
		return url
	}
}
