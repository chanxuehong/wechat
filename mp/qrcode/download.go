package qrcode

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/chanxuehong/wechat/internal/debug/api"
	"github.com/chanxuehong/wechat/util"
)

// 二维码图片的URL, 可以通过此URL下载二维码 或者 在线显示此二维码.
func QrcodePicURL(ticket string) string {
	return "https://mp.weixin.qq.com/cgi-bin/showqrcode?ticket=" + url.QueryEscape(ticket)
}

// Download 通过ticket换取二维码, 写入到 filepath 路径的文件.
//  如果 clt == nil 则默认用 util.DefaultHttpClient.
func Download(ticket, filepath string, clt *http.Client) (written int64, err error) {
	file, err := os.Create(filepath)
	if err != nil {
		return
	}
	defer func() {
		file.Close()
		if err != nil {
			os.Remove(filepath)
		}
	}()

	return DownloadToWriter(ticket, file, clt)
}

// DownloadToWriter 通过ticket换取二维码, 写入到 writer.
//  如果 clt == nil 则默认用 util.DefaultHttpClient.
func DownloadToWriter(ticket string, writer io.Writer, clt *http.Client) (written int64, err error) {
	if clt == nil {
		clt = util.DefaultHttpClient
	}

	url := QrcodePicURL(ticket)
	api.DebugPrintGetRequest(url)
	httpResp, err := clt.Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	return io.Copy(writer, httpResp.Body)
}
