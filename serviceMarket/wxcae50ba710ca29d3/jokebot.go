package wxcae50ba710ca29d3

import (
	"encoding/json"

	"github.com/chanxuehong/wechat/miniapp/wxa/serviceMarket"
	"github.com/chanxuehong/wechat/mp/core"
)

type JokeBotMode = int

const (
	COLD_JOKE     JokeBotMode = 1 // 冷笑话
	NORMAL_JOKE   JokeBotMode = 2 // 普通笑话
	ROMANTIC_JOKE JokeBotMode = 3 // 浪漫情话
	CHEESY_JOKE   JokeBotMode = 4 // 土味情话
	CHEEP_JOKE    JokeBotMode = 5 // 心灵鸡汤
)

type JokeBotResponse struct {
	ErrCode  int        `json:"err_code"`
	ErrMsg   string     `json:"err_msg"`
	DataList []JokeData `json:"data_list,omitempty"`
}

func (this *JokeBotResponse) IsError() bool {
	return this.ErrCode == -1
}

func (this *JokeBotResponse) Error() string {
	return this.ErrMsg
}

type JokeData struct {
	Result string `json:"result"`
}

func JokeBot(clt *core.Client, mode JokeBotMode) ([]string, error) {
	data := map[string]JokeBotMode{"mode": mode}
	req := &serviceMarket.InvokeServiceRequest{
		Service: SERVICE,
		Api:     JOKE_BOT_API,
		Data:    data,
	}
	resData, err := serviceMarket.InvokeService(clt, req)
	if err != nil {
		return nil, err
	}
	var resp JokeBotResponse
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
