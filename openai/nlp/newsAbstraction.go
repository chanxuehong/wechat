package nlp

import (
	mpCore "github.com/bububa/wechat/mp/core"
	"github.com/bububa/wechat/openai/core"
	"github.com/bububa/wechat/openai/model"
)

type NewsAbstractionRequest struct {
	Title          string `json:"title,omitempty"`
	Content        string `json:"content,omitempty"`
	Category       string `json:"category,omitempty"`
	DoNewsClassify bool   `json:"do_news_classify,omitempty"`
}

// NewsAbstraction 本服务目前支持对输入新闻进行是否适合提取摘要的分类， 同时支持对给定新闻进行摘要自动提取
func NewsAbstraction(clt *core.Client, uid string, news *NewsAbstractionRequest) (ret *model.NewsAbstraction, err error) {
	const incompleteURL = "https://openai.weixin.qq.com/openapi/nlp/news-abstraction/"

	query, err := Sign(clt, uid, map[string]interface{}{"title": news.Title, "content": news.Content, "category": news.Category, "do_news_classify": news.DoNewsClassify})
	if err != nil {
		return nil, err
	}
	req := map[string]string{
		"query": query,
	}
	var result struct {
		mpCore.Error
		*model.NewsAbstraction
	}

	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != mpCore.ErrCodeOK {
		err = &result.Error
		return
	}
	ret = result.NewsAbstraction
	return
}
