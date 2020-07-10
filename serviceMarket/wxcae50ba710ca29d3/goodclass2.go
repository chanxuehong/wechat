package wxcae50ba710ca29d3

import (
	"encoding/json"

	"github.com/chanxuehong/wechat/miniapp/wxa/serviceMarket"
	"github.com/chanxuehong/wechat/mp/core"
)

func GoodClass2(clt *core.Client, content string) (isGood bool, err error) {
	data := map[string]string{"content": content}
	req := &serviceMarket.InvokeServiceRequest{
		Service: SERVICE,
		Api:     GOOD_CLASS2_API,
		Data:    data,
	}
	resData, err := serviceMarket.InvokeService(clt, req)
	if err != nil {
		return false, err
	}
	// 需要关注的字段是model_result(模型预测结果)和es_result(elastic search搜索结果)：如果es_result == -1，那么最终结果依据model_result；如果es_result != -1, 那么最终结果依据es_result
	var tmp struct {
		Model string `json:"model_result"` // 服务分类结果 （0为非商品，1为商品）
		Es    string `json:"es_result"`    // 服务分类结果（0为非商品，1为商品）
	}
	err = json.Unmarshal([]byte(resData), &tmp)
	if err != nil {
		return false, err
	}
	if tmp.Es == "-1" {
		return tmp.Model == "1", nil
	}
	return tmp.Es == "1", nil
}
