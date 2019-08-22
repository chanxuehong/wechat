package mall

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type QueryProductRequest struct {
	KeyList []QueryKey `json:"key_list"`
}

type QueryKey struct {
	ItemCode string `json:"item_code"`
}

type QueryProductResponse struct {
	core.Error
	ProductList []Product `json:"product_list"`
}

// 查询物品信息
// 开发者可以对导入到好物圈中的物品信息进行查询，接口说明如下
func QueryProduct(clt *core.Client, req *QueryProductRequest) (resp QueryProductResponse, err error) {
	const incompleteURL = "https://api.weixin.qq.com/mall/queryproduct?type=batchquery&access_token="
	if err = clt.PostJSON(incompleteURL, req, &resp); err != nil {
		return
	}
	if resp.ErrCode != core.ErrCodeOK {
		err = &resp.Error
		return
	}
	return resp, nil
}
