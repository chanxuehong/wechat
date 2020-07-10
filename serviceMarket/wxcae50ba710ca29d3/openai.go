package wxcae50ba710ca29d3

import (
	"encoding/json"

	"github.com/chanxuehong/wechat/miniapp/wxa/serviceMarket"
	"github.com/chanxuehong/wechat/mp/core"
)

type OpenAiRequest struct {
	AppId  string `json:"appid"`
	Query  string `json:"query"`
	Name   string `json:"name,omitempty"`
	Avatar string `json:"avatar,omitempty"`
}

type OpenAiResponse struct {
	AnsNodeName string          `json:"ans_node_name,omitempty"` // 技能名称
	Answer      string          `json:"answer,omitempty"`        // 机器人回答
	ListOptions bool            `json:"list_options,omitempty"`  // 是否有相似问题推荐
	Opening     string          `json:"opening,omitempty"`       // 有相似问题推荐时的话术
	Status      string          `json:"status,omitempty"`        // 为NOMATCH代表未命中
	Options     []OpenAiOption  `json:"options,omitempty"`       // 推荐的相似问题
	Title       string          `json:"title,omitempty"`         // 意图名称
	MoreInfo    *OpenAiMoreInfo `json:"more_info,omitempty"`
}

type OpenAiOption struct {
	AnsNodeId   uint64  `json:"ans_node_id,omitempty"`
	AnsNodeName string  `json:"ans_node_name,omitempty"`
	Confidence  float64 `json:"confidence,omitempty"`
	Title       string  `json:"title,omitempty"`
}

type OpenAiMoreInfo struct {
	Music string `json:"music_ans_detail,omitempty"`
	Fm    string `json:"fm_ans_detail,omitempty"`
}

func OpenAi(clt *core.Client, data *OpenAiRequest) (*OpenAiResponse, error) {
	req := &serviceMarket.InvokeServiceRequest{
		Service: SERVICE,
		Api:     OPENAI_API,
		Data:    data,
	}
	resData, err := serviceMarket.InvokeService(clt, req)
	if err != nil {
		return nil, err
	}
	var resp OpenAiResponse
	err = json.Unmarshal([]byte(resData), &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
