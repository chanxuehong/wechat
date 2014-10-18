// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package product

const (
	// 商品状态
	PRODUCT_STATUS_ALL      = 0 // 所有商品
	PRODUCT_STATUS_ONSHELF  = 1 // 上架商品
	PRODUCT_STATUS_OFFSHELF = 2 // 下架商品
)

const (
	// 快递ID列表
	EXPRESS_ID_PINGYOU = 10000027 // 平邮
	EXPRESS_ID_KUAIDI  = 10000028 // 快递
	EXPRESS_ID_EMS     = 10000029 // EMS
)
