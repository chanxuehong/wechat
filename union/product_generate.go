package union

import "github.com/chanxuehong/wechat/mp/core"

// ProductGenerateRequest 获取商品推广素材
type ProductGenerateRequest struct {
	// Pid 推广位PID
	Pid string `json:"pid,omitempty"`
	// ProductList 商品列表
	ProductList []Product `json:"productList,omitempty"`
}

// ShareProduct 商品
type ShareProduct struct {
	// ProductID 商品SPU ID
	ProductID string `json:"productId,omitempty"`
	// AppID 商品所在小商店的AppID
	AppID string `json:"appId,omitempty"`
	// ProductInfo 商品具体信息
	ProductInfo *ProductInfo `json:"info,omitempty"`
	// ShareInfo 推广相关信息
	ShareInfo *ShareInfo `json:"shareInfo,omitempty"`
}

// 拉取会员信息（积分查询）接口
// ProductGenerate 获取商品推广素材
// 通过该接口获取商品的推广素材，包括店铺appID、商品详情页Path、推广文案及推广短链、商品图片等
func ProductGenerate(clt *core.Client, req *ProductGenerateRequest) (list []ShareProduct, err error) {
	var result struct {
		core.Error
		List []ShareProduct `json:"list"`
	}

	incompleteURL := "https://api.weixin.qq.com/union/promoter/product/generate?access_token="
	if err = clt.PostJSON(incompleteURL, req, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.List
	return
}
