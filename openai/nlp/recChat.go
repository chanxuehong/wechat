package nlp

import (
	mpCore "github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/openai/core"
)

// RecChat 给定文章内容，输出文章类别
func RecChat(clt *core.Client, uid string) (ret []string, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/nlp/rec_chat/"

	query, err := Sign(clt, uid, map[string]interface{}{})
	if err != nil {
		return nil, err
	}
	req := map[string]string{
		"query": query,
	}
	var result struct {
		mpCore.Error
		Quetions []string `json:"rec_questions"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.Quetions
	return
}
