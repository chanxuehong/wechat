// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package express

const (
	// 支付方式(0-买家承担运费, 1-卖家承担运费)
	ASSUMER_BUYER  = 0
	ASSUMER_SELLER = 1
)

const (
	// 计费单位(0-按件计费, 1-按重量计费, 2-按体积计费，目前只支持按件计费，默认为0)
	VALUATION_BY_ITEM   = 0
	VALUATION_BY_WEIGHT = 1
	VALUATION_BY_VOLUME = 2
)
