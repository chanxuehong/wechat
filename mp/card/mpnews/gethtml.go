<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mpnews

import (
	"github.com/chanxuehong/wechat/mp"
=======
package mpnews

import (
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

// 获取卡券嵌入图文消息的标准格式代码.
//  将返回代码填入上传图文素材接口中content字段，即可获取嵌入卡券的图文消息素材。
<<<<<<< HEAD
func GetHTML(clt *mp.Client, cardId string) (content string, err error) {
=======
func GetHTML(clt *core.Client, cardId string) (content string, err error) {
>>>>>>> github/v2
	request := struct {
		CardId string `json:"card_id"`
	}{
		CardId: cardId,
	}

	var result struct {
<<<<<<< HEAD
		mp.Error
=======
		core.Error
>>>>>>> github/v2
		Content string `json:"content"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/mpnews/gethtml?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result.Error
		return
	}
	content = result.Content
	return
}
