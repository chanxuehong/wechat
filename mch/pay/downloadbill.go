package pay

import (
	"bytes"
	"crypto/md5"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/chanxuehong/rand"
	"github.com/chanxuehong/util"

	"github.com/chanxuehong/wechat.v2/mch/core"
	wechatutil "github.com/chanxuehong/wechat.v2/util"
)

type DownloadBillRequest struct {
	AppId    string `xml:"appid"`     // 公众账号ID
	MchId    string `xml:"mch_id"`    // 微信支付分配的商户号
	ApiKey   string `xml:"api_key"`   // 签名密钥
	NonceStr string `xml:"nonce_str"` // 随机字符串，不长于32位。推荐随机数生成算法
	SignType string `xml:"sign_type"` // 签名类型，目前支持HMAC-SHA256和MD5，默认为MD5
	BillDate string `xml:"bill_date"` // 下载对账单的日期，格式：20140603
	BillType string `xml:"bill_type"` // 账单类型
	TarType  string `xml:"tar_type"`  // 压缩账单
}

// 下载对账单到到文件.
func DownloadBill(filepath string, req *DownloadBillRequest, httpClient *http.Client) (written int64, err error) {
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
	return downloadBillToWriter(file, req, httpClient)
}

// 下载对账单到 io.Writer.
func DownloadBillToWriter(writer io.Writer, req *DownloadBillRequest, httpClient *http.Client) (written int64, err error) {
	if writer == nil {
		return 0, errors.New("nil writer")
	}
	if req == nil {
		return 0, errors.New("nil request req")
	}
	return downloadBillToWriter(writer, req, httpClient)
}

var (
	// <xml><return_code><![CDATA[FAIL]]></return_code>
	// <return_msg><![CDATA[require POST method]]></return_msg>
	// </xml>
	downloadBillErrorRootNodeStartElement       = []byte("<xml>")
	downloadBillErrorReturnCodeNodeStartElement = []byte("<return_code>")
)

// 下载对账单到 io.Writer.
func downloadBillToWriter(writer io.Writer, req *DownloadBillRequest, httpClient *http.Client) (written int64, err error) {
	if httpClient == nil {
		httpClient = wechatutil.DefaultHttpClient
	}

	m1 := make(map[string]string, 8)
	m1["appid"] = req.AppId
	m1["mch_id"] = req.MchId
	if req.NonceStr != "" {
		m1["nonce_str"] = req.NonceStr
	} else {
		m1["nonce_str"] = string(rand.NewHex())
	}
	m1["bill_date"] = req.BillDate
	m1["bill_type"] = req.BillType
	if req.TarType != "" {
		m1["tar_type"] = req.TarType
	}
	if req.SignType != "" {
		// m1["sign_type"] = req.SignType
		m1["sign_type"] = "MD5" // TODO(chanxuehong): 目前只支持 MD5, 后期修改
	}
	m1["sign"] = core.Sign(m1, req.ApiKey, md5.New) // TODO(chanxuehong): 目前只支持 MD5, 后期修改

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

	n, err := io.ReadFull(httpResp.Body, buffer)
	switch {
	case err == nil:
		// n == len(buf), 可以认为返回的是对账单而不是xml格式的错误信息
		written, err = bytes.NewReader(buffer).WriteTo(writer)
		if err != nil {
			return written, err
		}
		var n2 int64
		n2, err = io.CopyBuffer(writer, httpResp.Body, buffer)
		written += n2
		return written, err
	case err == io.ErrUnexpectedEOF:
		content := buffer[:n]
		if index := bytes.Index(content, downloadBillErrorRootNodeStartElement); index != -1 {
			if bytes.Contains(content[index+len(downloadBillErrorRootNodeStartElement):], downloadBillErrorReturnCodeNodeStartElement) {
				// 可以认为是错误信息了, 尝试解析xml
				var result core.Error
				if err = xml.Unmarshal(content, &result); err == nil {
					return 0, &result
				}
				// err != nil 执行默认的动作, 写入 writer
			}
		}
		return bytes.NewReader(content).WriteTo(writer)
	case err == io.EOF: // 返回空的body
		return 0, nil
	default: // 其他的错误
		return 0, err
	}
}
