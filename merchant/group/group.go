package group

type Group struct {
	Id   int64  `json:"group_id,omitempty"`
	Name string `json:"group_name"`
}
type GroupExt struct {
	Group
	Products []string `json:"product_list"` // 商品ID集合
}

// 修改分组商品的请求数据结构
type ProductModifyRequest struct {
	GroupId  int64               `json:"group_id"`
	Products []productModifyUnit `json:"product"`
}

type productModifyUnit struct {
	ProductId    string `json:"product_id"`
	ModifyAction int    `json:"mod_action"`
}

func NewProductModifyRequest(groupId int64,
	addProducts []string, delProducts []string) *ProductModifyRequest {

	r := ProductModifyRequest{
		GroupId: groupId,
	}
	r.Products = make([]productModifyUnit, len(addProducts)+len(delProducts))

	for i := 0; i < len(addProducts); i++ {
		r.Products[i].ProductId = addProducts[i]
		r.Products[i].ModifyAction = PRODUCT_MODIFY_ACTION_ADD
	}

	for i, j := len(addProducts), 0; j < len(delProducts); i, j = i+1, j+1 {
		r.Products[i].ProductId = delProducts[j]
		r.Products[i].ModifyAction = PRODUCT_MODIFY_ACTION_DEL
	}

	return &r
}

func (r *ProductModifyRequest) AddProduct(productId ...string) {
	switch {
	case len(productId) == 1:
		r.Products = append(r.Products,
			productModifyUnit{
				ProductId:    productId[0],
				ModifyAction: PRODUCT_MODIFY_ACTION_ADD,
			},
		)

	case len(productId) > 1:
		products := make([]productModifyUnit, len(productId))
		for i := 0; i < len(productId); i++ {
			products[i].ProductId = productId[i]
			products[i].ModifyAction = PRODUCT_MODIFY_ACTION_ADD
		}

		r.Products = append(r.Products, products...)
	}
}

func (r *ProductModifyRequest) DeleteProduct(productId ...string) {
	switch {
	case len(productId) == 1:
		r.Products = append(r.Products,
			productModifyUnit{
				ProductId:    productId[0],
				ModifyAction: PRODUCT_MODIFY_ACTION_DEL,
			},
		)

	case len(productId) > 1:
		products := make([]productModifyUnit, len(productId))
		for i := 0; i < len(productId); i++ {
			products[i].ProductId = productId[i]
			products[i].ModifyAction = PRODUCT_MODIFY_ACTION_DEL
		}

		r.Products = append(r.Products, products...)
	}
}
