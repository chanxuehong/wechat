package wxcae50ba710ca29d3

import (
	"encoding/json"

	"github.com/chanxuehong/wechat/miniapp/wxa/serviceMarket"
	"github.com/chanxuehong/wechat/mp/core"
)

type ThumbupBotResponse struct {
	ErrCode  int              `json:"err_code"`
	ErrMsg   string           `json:"err_msg"`
	DataList []ThumbupBotData `json:"data_list,omitempty"`
}

func (this *ThumbupBotResponse) IsError() bool {
	return this.ErrCode == -1
}

func (this *ThumbupBotResponse) Error() string {
	return this.ErrMsg
}

type ThumbupBotData struct {
	Result string `json:"result"`
}

func ThumbupBot(clt *core.Client, q string) ([]string, error) {
	data := map[string]string{"q": q}
	req := &serviceMarket.InvokeServiceRequest{
		Service: SERVICE,
		Api:     THUMBUP_BOT_API,
		Data:    data,
	}
	resData, err := serviceMarket.InvokeService(clt, req)
	if err != nil {
		return nil, err
	}
	var resp ThumbupBotResponse
	err = json.Unmarshal([]byte(resData), &resp)
	if err != nil {
		return nil, err
	}
	if resp.IsError() {
		return nil, &resp
	}
	var lines []string
	for _, i := range resp.DataList {
		if i.Result != "" {
			lines = append(lines, i.Result)
		}
	}
	return lines, nil
}
