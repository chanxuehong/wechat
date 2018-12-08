package mpnews

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 获取卡券嵌入图文消息的标准格式代码.
//  将返回代码填入上传图文素材接口中content字段，即可获取嵌入卡券的图文消息素材。
func GetHTML(clt *core.Client, cardId string) (content string, err error) {
	request := struct {
		CardId string `json:"card_id"`
	}{
		CardId: cardId,
	}

	var result struct {
		core.Error
		Content string `json:"content"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/mpnews/gethtml?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	content = result.Content
	return
}
