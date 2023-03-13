package wxcae50ba710ca29d3

import (
	"encoding/json"
	"strconv"

	"github.com/bububa/wechat/miniapp/wxa/serviceMarket"
	"github.com/bububa/wechat/mp/core"
)

type GoodInfoResult struct {
	ProcessedText string     `json:"preprocessed_text,omitempty"` // 输入文本预处理后的结果
	Entities      []GoodInfo `json:"entities,omitempty"`          // 商品属性抽取结果列表
}

type GoodInfo struct {
	Txt   string `json:"txt"`
	Attr  string `json:"attr"`
	Start int    `json:"start"`
	End   int    `json:"end"`
}

func GetGoodInfo(clt *core.Client, q string) (ret *GoodInfoResult, err error) {
	data := map[string]string{"q": q}
	req := &serviceMarket.InvokeServiceRequest{
		Service: SERVICE,
		Api:     GOOD_INFO_API,
		Data:    data,
	}
	resData, err := serviceMarket.InvokeService(clt, req)
	if err != nil {
		return nil, err
	}
	var tmp struct {
		ProcessedText string                     `json:"preprocessed_text,omitempty"` // 输入文本预处理后的结果
		Entities      map[string][][]interface{} `json:"entities,omitempty"`          // 商品属性抽取结果列表
	}
	err = json.Unmarshal([]byte(resData), &tmp)
	if err != nil {
		return nil, err
	}
	resp := &GoodInfoResult{
		ProcessedText: tmp.ProcessedText,
	}
	if infos, found := tmp.Entities["product"]; found {
		for _, info := range infos {
			goodInfo := GoodInfo{
				Txt:   info[0].(string),
				Attr:  info[1].(string),
				Start: toInt(info[2]),
				End:   toInt(info[3]),
			}
			resp.Entities = append(resp.Entities, goodInfo)
		}
	}
	return resp, nil
}

func toInt(v interface{}) int {
	switch v.(type) {
	case float64:
		return int(v.(float64))
	case string:
		tmp, _ := strconv.ParseInt(v.(string), 10, 64)
		return int(tmp)
	default:
		return 0
	}
}
