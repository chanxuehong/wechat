// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package group

type Group struct {
	Id   int64  `json:"group_id,omitempty"`
	Name string `json:"group_name"`
}

type GroupEx struct {
	Group
	ProductIds []string `json:"product_list,omitempty"` // 商品ID集合
}

// 修改分组商品的请求数据结构
type GroupModifyProductRequest struct {
	GroupId  int64                    `json:"group_id"`
	Products []groupModifyProductUnit `json:"product,omitempty"`
}

type groupModifyProductUnit struct {
	ProductId    string `json:"product_id"`
	ModifyAction int    `json:"mod_action"`
}

func NewGroupModifyProductRequest(groupId int64,
	addProducts []string, delProducts []string) *GroupModifyProductRequest {

	r := GroupModifyProductRequest{
		GroupId: groupId,
	}
	r.Products = make([]groupModifyProductUnit, len(addProducts)+len(delProducts))

	for i := 0; i < len(addProducts); i++ {
		r.Products[i].ProductId = addProducts[i]
		r.Products[i].ModifyAction = GROUP_PRODUCT_MODIFY_ACTION_ADD
	}

	for i, j := len(addProducts), 0; j < len(delProducts); i, j = i+1, j+1 {
		r.Products[i].ProductId = delProducts[j]
		r.Products[i].ModifyAction = GROUP_PRODUCT_MODIFY_ACTION_DEL
	}

	return &r
}

func (r *GroupModifyProductRequest) AddProduct(productId ...string) {
	switch {
	case len(productId) == 1:
		r.Products = append(r.Products,
			groupModifyProductUnit{
				ProductId:    productId[0],
				ModifyAction: GROUP_PRODUCT_MODIFY_ACTION_ADD,
			},
		)

	case len(productId) > 1:
		products := make([]groupModifyProductUnit, len(productId))
		for i := 0; i < len(productId); i++ {
			products[i].ProductId = productId[i]
			products[i].ModifyAction = GROUP_PRODUCT_MODIFY_ACTION_ADD
		}

		r.Products = append(r.Products, products...)
	}
}

func (r *GroupModifyProductRequest) DeleteProduct(productId ...string) {
	switch {
	case len(productId) == 1:
		r.Products = append(r.Products,
			groupModifyProductUnit{
				ProductId:    productId[0],
				ModifyAction: GROUP_PRODUCT_MODIFY_ACTION_DEL,
			},
		)

	case len(productId) > 1:
		products := make([]groupModifyProductUnit, len(productId))
		for i := 0; i < len(productId); i++ {
			products[i].ProductId = productId[i]
			products[i].ModifyAction = GROUP_PRODUCT_MODIFY_ACTION_DEL
		}

		r.Products = append(r.Products, products...)
	}
}
