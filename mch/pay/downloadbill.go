package pay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"unicode"

	"github.com/chanxuehong/util"

	"github.com/chanxuehong/wechat/mch/core"
	wechatutil "github.com/chanxuehong/wechat/util"
)

type DownloadBillRequest struct {
	XMLName struct{} `xml:"xml" json:"-"`

	// 必选参数
	BillDate string `xml:"bill_date"` // 下载对账单的日期，格式：20140603
	BillType string `xml:"bill_type"` // 账单类型

	// 可选参数
	DeviceInfo string `xml:"device_info"` // 微信支付分配的终端设备号
	NonceStr   string `xml:"nonce_str"`   // 随机字符串，不长于32位。推荐随机数生成算法
	SignType   string `xml:"sign_type"`   // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	TarType    string `xml:"tar_type"`    // 压缩账单
}

// 下载对账单到到文件.
func DownloadBill(clt *core.Client, filepath string, req *DownloadBillRequest, httpClient *http.Client) (written int64, err error) {
	if req == nil {
		return 0, errors.New("nil request req")
	}

	file, err := os.Create(filepath)
	if err != nil {
		return 0, err
	}
	defer func() {
		file.Close()
		if err != nil {
			os.Remove(filepath)
		}
	}()
	return downloadBillToWriter(clt, file, req, httpClient)
}

// 下载对账单到 io.Writer.
func DownloadBillToWriter(clt *core.Client, writer io.Writer, req *DownloadBillRequest, httpClient *http.Client) (written int64, err error) {
	if writer == nil {
		return 0, errors.New("nil writer")
	}
	if req == nil {
		return 0, errors.New("nil request req")
	}
	return downloadBillToWriter(clt, writer, req, httpClient)
}

var (
	// <xml><return_code><![CDATA[FAIL]]></return_code>
	// <return_msg><![CDATA[require POST method]]></return_msg>
	// </xml>
	downloadBillErrorRootNodeStartElement       = []byte("<xml>")
	downloadBillErrorReturnCodeNodeStartElement = []byte("<return_code>")
	downloadBillErrorReturnMsgNodeStartElement  = []byte("<return_msg>")
)

// 下载对账单到 io.Writer.
func downloadBillToWriter(clt *core.Client, writer io.Writer, req *DownloadBillRequest, httpClient *http.Client) (written int64, err error) {
	if httpClient == nil {
		httpClient = wechatutil.DefaultMediaHttpClient
	}

	m1 := make(map[string]string, 8)
	m1["appid"] = clt.AppId()
	m1["mch_id"] = clt.MchId()
	if subAppId := clt.SubAppId(); subAppId != "" {
		m1["sub_appid"] = subAppId
	}
	if subMchId := clt.SubMchId(); subMchId != "" {
		m1["sub_mch_id"] = subMchId
	}
	m1["bill_date"] = req.BillDate
	if req.BillType != "" {
		m1["bill_type"] = req.BillType
	}
	if req.DeviceInfo != "" {
		m1["device_info"] = req.DeviceInfo
	}
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = wechatutil.NonceStr()
	}
	if req.TarType != "" {
		m1["tar_type"] = req.TarType
	}

	// 签名
	switch req.SignType {
	case "":
		m1["sign"] = core.Sign2(m1, clt.ApiKey(), md5.New())
	case core.SignType_MD5:
		m1["sign_type"] = core.SignType_MD5
		m1["sign"] = core.Sign2(m1, clt.ApiKey(), md5.New())
	case core.SignType_HMAC_SHA256:
		m1["sign_type"] = core.SignType_HMAC_SHA256
		m1["sign"] = core.Sign2(m1, clt.ApiKey(), hmac.New(sha256.New, []byte(clt.ApiKey())))
	default:
		err = fmt.Errorf("unsupported request sign_type: %s", req.SignType)
		return 0, err
	}

	buffer := make([]byte, 32<<10) // 与 io.copyBuffer 里的默认大小一致

	requestBuffer := bytes.NewBuffer(buffer[:0])
	if err = util.EncodeXMLFromMap(requestBuffer, m1, "xml"); err != nil {
		return 0, err
	}

	httpResp, err := httpClient.Post(core.APIBaseURL()+"/pay/downloadbill", "text/xml; charset=utf-8", requestBuffer)
	if err != nil {
		return 0, err
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return 0, err
	}

	switch n, err := io.ReadFull(httpResp.Body, buffer); err {
	case nil:
		// n == len(buffer) == 32KB, 可以认为返回的是对账单而不是xml格式的错误信息
		written, err = bytes.NewReader(buffer).WriteTo(writer)
		if err != nil {
			return written, err
		}
		var n2 int64
		n2, err = io.CopyBuffer(writer, httpResp.Body, buffer)
		written += n2
		return written, err
	case io.ErrUnexpectedEOF:
		content := buffer[:n]
		if bs := trimLeft(content); bytes.HasPrefix(bs, downloadBillErrorRootNodeStartElement) {
			bs = trimLeft(bs[len(downloadBillErrorRootNodeStartElement):])
			if bytes.HasPrefix(bs, downloadBillErrorReturnCodeNodeStartElement) || bytes.HasPrefix(bs, downloadBillErrorReturnMsgNodeStartElement) {
				// 可以认为是错误信息了, 尝试解析xml
				var result core.Error
				if err = xml.Unmarshal(content, &result); err == nil {
					return 0, &result
				}
			}
		}
		return bytes.NewReader(content).WriteTo(writer)
	case io.EOF: // 返回空的body
		return 0, nil
	default: // 其他的错误
		return 0, err
	}
}

func trimLeft(s []byte) []byte {
	for i := 0; i < len(s); i++ {
		if isSpace(s[i]) {
			continue
		}
		return s[i:]
	}
	return s
}

func isSpace(b byte) bool {
	if b > unicode.MaxASCII {
		return false
	}
	return unicode.IsSpace(rune(b))
}
