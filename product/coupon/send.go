package coupon

import (
	"github.com/bububa/wechat/product/core"
)

// Send 发放优惠券
func Send(clt *core.Client, openId string, couponId uint64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/product/coupon/send?access_token="

	req := struct {
		Id     uint64 `json:"coupon_id"`
		OpenId string `json:"openid"`
	}{
		Id:     couponId,
		OpenId: openId,
	}

	var result core.Error
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
