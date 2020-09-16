package nlp

import (
	"fmt"

	mpCore "github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/openai/core"
)

type TranslateTo = string

const (
	TO_EN TranslateTo = "cn2en"
	TO_CN TranslateTo = "en2cn"
)

// Translate 双向中英互译，输入一段中文翻译成英文，或输入一段英文翻译成中文
func Translate(clt *core.Client, uid string, q string, to TranslateTo) (ret string, err error) {
	incompleteURL := fmt.Sprintf("https://openai.weixin.qq.com/openapi/nlp/translate_%s/", to)

	query, err := Sign(clt, uid, map[string]interface{}{"q": q})
	if err != nil {
		return "", err
	}
	req := map[string]string{
		"query": query,
	}
	var result struct {
		mpCore.Error
		Result string `json:"result"`
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.Result
	return
}
