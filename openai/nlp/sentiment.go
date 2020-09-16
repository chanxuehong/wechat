package nlp

import (
	mpCore "github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/openai/core"
)

type SentimentMode = string

const (
	CLASS3 SentimentMode = "3class"
	CLASS6 SentimentMode = "6class"
)

// Ner 词法分析接口(只签名不加密)
func Sentiment(clt *core.Client, uid string, q string, mode SentimentMode) (ret map[string]float64, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/nlp/sentiment/"

	query, err := Sign(clt, uid, map[string]interface{}{"q": q, "mode": mode})
	if err != nil {
		return nil, err
	}
	req := map[string]string{
		"query": query,
	}
	var result struct {
		mpCore.Error
		List [][]interface{} `json:"result"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = make(map[string]float64)
	for _, r := range result.List {
		ret[r[0].(string)] = r[1].(float64)
	}
	return
}
