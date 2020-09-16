package nlp

import (
	mpCore "github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/openai/core"
	"github.com/chanxuehong/wechat/openai/model"
)

// GetSimilarQuery 获取系统推荐的相似问题
func GetSimilarQuery(clt *core.Client, uid string, q string) (ret []model.SimilarQuery, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/nlp/get_similar_query/"

	query, err := Sign(clt, uid, map[string]interface{}{"query": q})
	if err != nil {
		return nil, err
	}
	req := map[string]string{
		"query": query,
	}
	var result struct {
		mpCore.Error
		List []model.SimilarQuery `json:"data"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.List
	return
}
