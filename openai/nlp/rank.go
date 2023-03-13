package nlp

import (
	mpCore "github.com/bububa/wechat/mp/core"
	"github.com/bububa/wechat/openai/core"
	"github.com/bububa/wechat/openai/model"
)

// Rank 句子相似度计算以及排序
func Rank(clt *core.Client, uid string, q string, candidates []string) (ret *model.RankResult, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/nlp/rank/"

	var candidateList []map[string]string
	for _, v := range candidates {
		candidateList = append(candidateList, map[string]string{
			"text": v,
		})
	}
	query, err := Sign(clt, uid, map[string]interface{}{"query": q, "candidates": candidateList})
	if err != nil {
		return nil, err
	}
	req := map[string]string{
		"query": query,
	}
	var result struct {
		mpCore.Error
		*model.RankResult
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.RankResult
	return
}
