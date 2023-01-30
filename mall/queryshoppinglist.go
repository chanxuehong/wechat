package mall

import (
	"github.com/chanxuehong/wechat/mp/core"
	"github.com/chanxuehong/wechat/util"
)

type QueryType = string

const (
	BatchQuery QueryType = "batchquery"
	GetByPage  QueryType = "getbypage"
)

type QueryShoppingListRequest struct {
	OpenId  string     `json:"user_open_id"`       // 用户的openid
	KeyList []QueryKey `json:"key_list,omitempty"` // batchquery模式下必填, 单次请求物品数量不可超过20个
	Offset  int        `json:"offset,omitempty"`   // 按页查询时起始位置偏移，默认0
	Count   int        `json:"count,omitempty"`    // 按页查询时单次最大返回数量，默认20
}

type QueryShoppingListResponse struct {
	core.Error
	GoodsList []Good `json:"goods_list"`
}

// 查询用户收藏信息
// 开发者可以查询用户在好物圈中指定商家的收藏物品
func QueryShoppingList(clt *core.Client, req *QueryShoppingListRequest, queryType QueryType) (resp QueryShoppingListResponse, err error) {
	incompleteURL := util.StringsJoin("https://api.weixin.qq.com/mall/queryshoppinglist?type=", queryType, "&access_token=")
	if err = clt.PostJSON(incompleteURL, req, &resp); err != nil {
		return
	}
	if resp.ErrCode != core.ErrCodeOK {
		err = &resp.Error
		return
	}
	return resp, nil
}
