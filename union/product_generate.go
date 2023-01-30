package union

import "github.com/chanxuehong/wechat/mp/core"

// ProductGenerateRequest 获取商品推广素材
type ProductGenerateRequest struct {
	// Pid 推广位PID
	Pid string `json:"pid,omitempty"`
	// ProductList 商品列表
	ProductList []Product `json:"productList,omitempty"`
}

// 拉取会员信息（积分查询）接口
// ProductGenerate 获取商品推广素材
// 通过该接口获取商品的推广素材，包括店铺appID、商品详情页Path、推广文案及推广短链、商品图片等
func ProductGenerate(clt *core.Client, req *ProductGenerateRequest) (products []Product, err error) {
	var result struct {
		core.Error
		List []Product `json:"list"`
	}

	incompleteURL := "https://api.weixin.qq.com/union/promoter/product/generate?access_token="
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	products = result.List
	return
}
