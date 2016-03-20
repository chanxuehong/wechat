<<<<<<< HEAD
// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package qrcode

import (
	"github.com/chanxuehong/wechat/mp"
=======
package qrcode

import (
	"github.com/chanxuehong/wechat/mp/core"
>>>>>>> github/v2
)

// 获取物料二维码
//  shopId: 门店ID
//  imgId:  物料样式编号：
//          0-二维码，可用于自由设计宣传材料；
//          1-桌贴（二维码），100mm×100mm(宽×高)，可直接张贴
<<<<<<< HEAD
func Get(clt *mp.Client, shopId int64, imgId int) (qrcodeURL string, err error) {
=======
func Get(clt *core.Client, shopId int64, imgId int) (qrcodeURL string, err error) {
>>>>>>> github/v2
	request := struct {
		ShopId int64 `json:"shop_id"`
		ImgId  int   `json:"img_id"`
	}{
		ShopId: shopId,
		ImgId:  imgId,
	}

	var result struct {
<<<<<<< HEAD
		mp.Error
		Data struct {
			QRCodeURL string `json:"qrcode_url"`
=======
		core.Error
		Data struct {
			QrcodeURL string `json:"qrcode_url"`
>>>>>>> github/v2
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/bizwifi/qrcode/get?access_token="
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

<<<<<<< HEAD
	qrcodeURL = result.Data.QRCodeURL
=======
	qrcodeURL = result.Data.QrcodeURL
>>>>>>> github/v2
	return
}
