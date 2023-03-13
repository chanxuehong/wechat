package nlp

import (
	mpCore "github.com/bububa/wechat/mp/core"
	"github.com/bububa/wechat/openai/core"
	"github.com/bububa/wechat/openai/model"
)

type DocumentClassifyRequest struct {
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
}

// DocumentClassify 给定文章内容，输出文章类别
func DocumentClassify(clt *core.Client, uid string, doc *DocumentClassifyRequest) (ret *model.DocumentClassify, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/nlp/document_classify/"

	query, err := Sign(clt, uid, map[string]interface{}{"title": doc.Title, "content": doc.Content})
	if err != nil {
		return nil, err
	}
	req := map[string]string{
		"query": query,
	}
	var result struct {
		mpCore.Error
		*model.DocumentClassify
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.DocumentClassify
	return
}
