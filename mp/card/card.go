package card

import (
	"github.com/chanxuehong/wechat/mp/core"
)

// 创建卡券.
func Create(clt *core.Client, card *Card) (cardId string, err error) {
	request := struct {
		Card *Card `json:"card,omitempty"`
	}{
		Card: card,
	}

	var result struct {
		core.Error
		CardId string `json:"card_id"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/create?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	cardId = result.CardId
	return
}

// 查看卡券详情.
func Get(clt *core.Client, cardId string) (card *Card, err error) {
	request := struct {
		CardId string `json:"card_id"`
	}{
		CardId: cardId,
	}

	var result struct {
		core.Error
		Card `json:"card"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/get?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	card = &result.Card
	return
}

type BatchGetQuery struct {
	Offset     int      `json:"offset"`                // 查询卡列表的起始偏移量，从0开始，即offset: 5是指从从列表里的第六个开始读取。
	Count      int      `json:"count"`                 // 需要查询的卡片的数量（数量最大50）。
	StatusList []string `json:"status_list,omitempty"` // 支持开发者拉出指定状态的卡券列表，例：仅拉出通过审核的卡券。
}

type BatchGetResult struct {
	TotalNum   int      `json:"total_num"`
	ItemNum    int      `json:"item_num"`
	CardIdList []string `json:"card_id_list"`
}

// 批量查询卡列表.
func BatchGet(clt *core.Client, query *BatchGetQuery) (rslt *BatchGetResult, err error) {
	var result struct {
		core.Error
		BatchGetResult
	}

	incompleteURL := "https://api.weixin.qq.com/card/batchget?access_token="
	if err = clt.PostJSON(incompleteURL, query, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	result.BatchGetResult.ItemNum = len(result.BatchGetResult.CardIdList)
	rslt = &result.BatchGetResult
	return
}

// 更改卡券信息接口.
//  sendCheck: 是否提交审核，false为修改后不会重新提审，true为修改字段后重新提审，该卡券的状态变为审核中。
func Update(clt *core.Client, cardId string, card *Card) (sendCheck bool, err error) {
	request := struct {
		CardId string `json:"card_id"`
		*Card
	}{
		CardId: cardId,
		Card:   card,
	}

	var result struct {
		core.Error
		SendCheck bool `json:"send_check"`
	}

	incompleteURL := "https://api.weixin.qq.com/card/update?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	sendCheck = result.SendCheck
	return
}

// 库存修改接口.
// cardId:      卡券ID
// increaseNum: 增加库存数量, 可以为负数
func ModifyStock(clt *core.Client, cardId string, increaseNum int) (err error) {
	request := struct {
		CardId             string `json:"card_id"`
		IncreaseStockValue int    `json:"increase_stock_value,omitempty"`
		ReduceStockValue   int    `json:"reduce_stock_value,omitempty"`
	}{
		CardId: cardId,
	}
	switch {
	case increaseNum > 0:
		request.IncreaseStockValue = increaseNum
	case increaseNum < 0:
		request.ReduceStockValue = -increaseNum
	default: // increaseNum == 0
		return
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/card/modifystock?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 删除卡券
func Delete(clt *core.Client, cardId string) (err error) {
	request := struct {
		CardId string `json:"card_id"`
	}{
		CardId: cardId,
	}

	var result core.Error

	incompleteURL := "https://api.weixin.qq.com/card/delete?access_token="
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}

	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
