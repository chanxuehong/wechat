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
type GroupProductModifyRequest struct {
	GroupId  int64                    `json:"group_id"`
	Products []groupProductModifyUnit `json:"product,omitempty"`
}

type groupProductModifyUnit struct {
	ProductId    string `json:"product_id"`
	ModifyAction int    `json:"mod_action"`
}

func NewGroupProductModifyRequest(groupId int64,
	addProducts []string, delProducts []string) *GroupProductModifyRequest {

	r := GroupProductModifyRequest{
		GroupId: groupId,
	}
	r.Products = make([]groupProductModifyUnit, len(addProducts)+len(delProducts))

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

func (r *GroupProductModifyRequest) AddProduct(productId ...string) {
	switch {
	case len(productId) == 1:
		r.Products = append(r.Products,
			groupProductModifyUnit{
				ProductId:    productId[0],
				ModifyAction: GROUP_PRODUCT_MODIFY_ACTION_ADD,
			},
		)

	case len(productId) > 1:
		products := make([]groupProductModifyUnit, len(productId))
		for i := 0; i < len(productId); i++ {
			products[i].ProductId = productId[i]
			products[i].ModifyAction = GROUP_PRODUCT_MODIFY_ACTION_ADD
		}

		r.Products = append(r.Products, products...)
	}
}

func (r *GroupProductModifyRequest) DeleteProduct(productId ...string) {
	switch {
	case len(productId) == 1:
		r.Products = append(r.Products,
			groupProductModifyUnit{
				ProductId:    productId[0],
				ModifyAction: GROUP_PRODUCT_MODIFY_ACTION_DEL,
			},
		)

	case len(productId) > 1:
		products := make([]groupProductModifyUnit, len(productId))
		for i := 0; i < len(productId); i++ {
			products[i].ProductId = productId[i]
			products[i].ModifyAction = GROUP_PRODUCT_MODIFY_ACTION_DEL
		}

		r.Products = append(r.Products, products...)
	}
}
