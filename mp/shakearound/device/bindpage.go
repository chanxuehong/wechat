<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com), magicshui(shuiyuzhe@gmail.com), Harry Rong(harrykobe@gmail.com)

package device

import (
	"github.com/chanxuehong/wechat/mp"
=======
package device

import (
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

type BindPageParameters struct {
	DeviceIdentifier *DeviceIdentifier `json:"device_identifier,omitempty"` // 必须, 设备标识
	PageIds          []int64           `json:"page_ids,omitempty"`          // 必须, 待关联的页面列表
	Bind             int               `json:"bind"`                        // 必须, 关联操作标志位， 0为解除关联关系，1为建立关联关系
	Append           int               `json:"append"`                      // 必须, 新增操作标志位， 0为覆盖，1为新增
}

<<<<<<< HEAD
func BindPage(clt *mp.Client, para *BindPageParameters) (err error) {
	var result mp.Error
=======
func BindPage(clt *core.Client, para *BindPageParameters) (err error) {
	var result core.Error
>>>>>>> github/v2

	incompleteURL := "https://api.weixin.qq.com/shakearound/device/bindpage?access_token="
	if err = clt.PostJSON(incompleteURL, para, &result); err != nil {
		return
	}

<<<<<<< HEAD
	if result.ErrCode != mp.ErrCodeOK {
=======
	if result.ErrCode != core.ErrCodeOK {
>>>>>>> github/v2
		err = &result
		return
	}
	return
}
