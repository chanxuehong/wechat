package util

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/chanxuehong/wechat/internal/api"
)

func Download(url, filepath string, httpClient *http.Client) (written int64, err error) {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

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

	return downloadToWriter(url, file, httpClient)
}

func DownloadToWriter(url string, w io.Writer, httpClient *http.Client) (written int64, err error) {
	if w == nil {
		err = errors.New("nil w io.Writer")
		return
	}
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return downloadToWriter(url, w, httpClient)
}

func downloadToWriter(url string, w io.Writer, httpClient *http.Client) (written int64, err error) {
	api.DebugPrintGetRequest(url)
	httpResp, err := httpClient.Get(url)
	if err != nil {
		return
	}
	defer httpResp.Body.Close()

	if httpResp.StatusCode != http.StatusOK {
		err = fmt.Errorf("http.Status: %s", httpResp.Status)
		return
	}

	return io.Copy(w, httpResp.Body)
}
