package group

type Group struct {
	Id   int    `json:"group_id,omitempty"`
	Name string `json:"group_name"`
}

type GroupExt struct {
	Group
	Product []string `json:"product_list"`
}

type GroupModifyRequest struct {
	GroupId int                  `json:"group_id"`
	Product []GroupModifyProduct `json:"product"`
}

type GroupModifyProduct struct {
	ProductId    string `json:"product_id"`
	ModifyAction int    `json:"mod_action"`
}

func (r *GroupModifyRequest) DeleteProduct(productId string) {
	r.Product = append(r.Product,
		GroupModifyProduct{
			ProductId:    productId,
			ModifyAction: GROUPMODIFY_ACTION_DEL,
		},
	)
}
func (r *GroupModifyRequest) AddProduct(productId string) {
	r.Product = append(r.Product,
		GroupModifyProduct{
			ProductId:    productId,
			ModifyAction: GROUPMODIFY_ACTION_ADD,
		},
	)
}
