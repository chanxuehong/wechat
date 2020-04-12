package wxa

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type WxAppCategory struct {
	FirstClass  string `json:"first_class"`  // 一级类目名称，可通过“获取授权小程序帐号的可选类目”接口获得
	SecondClass string `json:"second_class"` // 二级类目
	ThirdClass  string `json:"third_class"`  // 三级类目
	FirstId     uint   `json:"first_id"`     // 一级类目的ID，可通过“获取授权小程序帐号的可选类目”接口获得
	SecondId    uint   `json:"second_id"`    // 二级类目的ID
	ThirdId     uint   `json:"third_id"`     // 三级类目的ID
}

// 获取授权小程序帐号已设置的类目, 注意：该接口可获取已设置的二级类目及用于代码审核的可选三级类目。
func GetCategory(clt *core.Client) (categories []WxAppCategory, err error) {
	const incompleteURL = "https://api.weixin.qq.com/wxa/get_category?access_token="
	var result struct {
		core.Error
		CategoryList []WxAppCategory `json:"category_list"` // 可填选的类目列表
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	return result.CategoryList, nil
}
