package nlp

import (
	mpCore "github.com/bububa/wechat/mp/core"
	"github.com/bububa/wechat/openai/core"
)

type SensitiveMode = string

const (
	CNN  SensitiveMode = "cnn"
	BERT SensitiveMode = "bert"
)

// Sensitive 文本敏感内容审核，自动审核是否包含违规内容，例如涉政、色情、辱骂

func Sensitive(clt *core.Client, uid string, q string, mode SensitiveMode) (ret map[string]float64, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/nlp/sensitive/"

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
