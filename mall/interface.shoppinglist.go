package mall

type Good struct {
	BizUin     uint64 `json:"biz_uin"`
	ItemCode   string `json:"item_code"`
	SkuId      string `json:"sku_id"`
	Title      string `json:"title"`
	Quantity   uint   `json:"quantity"`
	CreateTime int64  `json:"create_time"`
	UpdateTime int64  `json:"update_time"`
	Source     int    `json:"source"`
	Status     int    `json:"status"`
	FromScene  int    `json:"from_scene"`
}
