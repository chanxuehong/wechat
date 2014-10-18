// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package category

// 获取指定子分类的所有SKU 成功时返回结果的数据结构
//
//  {
//      "id": "1075741873",
//      "name": "颜色",
//      "value_list": [
//          {
//              "id": "1079742375",
//              "name": "撞色"
//          },
//          {
//              "id": "1079742376",
//              "name": "桔色"
//          }
//      ]
//  }
type SKU struct {
	Id     string `json:"id"` // int64?
	Name   string `json:"name"`
	Values []struct {
		Id   string `json:"id"` // int64?
		Name string `json:"name"`
	} `json:"value_list,omitempty"`
}
