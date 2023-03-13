package openai

import (
	mpCore "github.com/bububa/wechat/mp/core"
	"github.com/bububa/wechat/openai/core"
	"github.com/bububa/wechat/openai/model"
)

// AiBot 智能对话接口
func AiBot(clt *core.Client, sign string, query string) (ret *model.NLUResult, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/aibot/"

	req := map[string]string{
		"signature": sign,
		"query":     query,
	}
	var result struct {
		mpCore.Error
		*model.NLUResult
	}
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.NLUResult
	return
}
