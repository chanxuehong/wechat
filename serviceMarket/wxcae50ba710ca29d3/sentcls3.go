package wxcae50ba710ca29d3

import (
	"encoding/json"

	"github.com/bububa/wechat/miniapp/wxa/serviceMarket"
	"github.com/bububa/wechat/mp/core"
)

type Sentcls3Result struct {
	Label string  `json:"label,omitempty"`
	Score float64 `json:"score,omitempty"`
}

func Sentcls3(clt *core.Client, q string) ([]Sentcls3Result, error) {
	data := map[string]string{"q": q}
	req := &serviceMarket.InvokeServiceRequest{
		Service: SERVICE,
		Api:     SENTCLS3_API,
		Data:    data,
	}
	resData, err := serviceMarket.InvokeService(clt, req)
	if err != nil {
		return nil, err
	}
	var tmp struct {
		Result [][]interface{} `json:"result"`
	}
	err = json.Unmarshal([]byte(resData), &tmp)
	if err != nil {
		return nil, err
	}
	var resp []Sentcls3Result
	for _, i := range tmp.Result {
		resp = append(resp, Sentcls3Result{
			Label: i[0].(string),
			Score: i[1].(float64),
		})
	}
	return resp, nil
}
