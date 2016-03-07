package qrcode

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 获取物料二维码
//  shopId: 门店ID
//  imgId:  物料样式编号：
//          0-二维码，可用于自由设计宣传材料；
//          1-桌贴（二维码），100mm×100mm(宽×高)，可直接张贴
func Get(clt *core.Client, shopId int64, imgId int) (qrcodeURL string, err error) {
	request := struct {
		ShopId int64 `json:"shop_id"`
		ImgId  int   `json:"img_id"`
	}{
		ShopId: shopId,
		ImgId:  imgId,
	}

	var result struct {
		core.Error
		Data struct {
			QrcodeURL string `json:"qrcode_url"`
		} `json:"data"`
	}

	incompleteURL := "https://api.weixin.qq.com/bizwifi/qrcode/get?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}

	qrcodeURL = result.Data.QrcodeURL
	return
}
