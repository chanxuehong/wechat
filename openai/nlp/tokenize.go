package nlp

import (
	mpCore "github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/openai/core"
	"github.com/chanxuehong/wechat/openai/model"
)

// Tokenize 词法分析接口(只签名不加密)
func Tokenize(clt *core.Client, uid string, q string) (ret *model.TokenizeResult, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/nlp/tokenize/"

	query, err := Sign(clt, uid, map[string]interface{}{"q": q})
	if err != nil {
		return nil, err
	}
	req := map[string]string{
		"query": query,
	}
	var result struct {
		mpCore.Error
		*model.TokenizeResult
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.TokenizeResult
	return
}
