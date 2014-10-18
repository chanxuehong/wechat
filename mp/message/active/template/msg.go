// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package template

// 模板消息数据结构, 包括 纯文本模板消息 和 带链接的模板消息
type Msg struct {
	ToUser string `json:"touser"` // 接受者OpenID

	TemplateId string `json:"template_id"` // 模版ID

	// 模版调用时需要的参数赋值内容，具体有哪些参数视模版内容而定;
	// Data 一般是一个 struct 或者 map, 要求被 encoding/json 格式化后满足特定的模板需求.
	Data interface{} `json:"data"`

	// 下面两个字段是 带链接模版消息 才有的
	URL      string `json:"url,omitempty"`      // 用户点击后跳转的URL，该URL必须处于开发者在公众平台网站中设置的域中
	TopColor string `json:"topcolor,omitempty"` // 整个消息的颜色, 可以不设置
}

// 新建 纯文本模板消息
//
//  {
//      "touser": "OPENID",
//      "template_id": "aygtGTLdrjHJP7Bu4EdkptNfYaeFKi98ygn2kitCJ6fAfdmN88naVvX6V5uIV5x0",
//      "data": {
//          "Goods": "苹果",
//          "Unit_price": "RMB 20.13",
//          "Quantity": "5",
//          "Total": "RMB 100.65",
//          "Source": {
//              "Shop": "Jas屌丝商店",
//              "Recommend": "5颗星"
//          }
//      }
//  }
//
func NewMsg(toUser, templateId string, data interface{}) *Msg {
	return &Msg{
		ToUser:     toUser,
		TemplateId: templateId,
		Data:       data,
	}
}

// 新建 带链接的模板消息
//  NOTE:
//  url 置空，则在发送后，点击模板消息会进入一个空白页面（ios），或无法点击（android）。
//  topColor 可以为空
//
//  {
//      "touser": "OPENID",
//      "template_id": "ngqIpbwh8bUfcSsECmogfXcV14J0tQlEpBO27izEYtY",
//      "url": "http://weixin.qq.com/download",
//      "topcolor": "#FF0000",
//      "data": {
//          "User": {
//              "value": "黄先生",
//              "color": "#173177"
//          },
//          "Date": {
//              "value": "06月07日 19时24分",
//              "color": "#173177"
//          },
//          "CardNumber": {
//              "value": "0426",
//              "color": "#173177"
//          },
//          "Type": {
//              "value": "消费",
//              "color": "#173177"
//          },
//          "Money": {
//              "value": "人民币260.00元",
//              "color": "#173177"
//          },
//          "DeadTime": {
//              "value": "06月07日19时24分",
//              "color": "#173177"
//          },
//          "Left": {
//              "value": "6504.09",
//              "color": "#173177"
//          }
//      }
//  }
//
func NewMsgWithLink(toUser, templateId string, data interface{}, url, topColor string) *Msg {
	return &Msg{
		ToUser:     toUser,
		TemplateId: templateId,
		Data:       data,
		URL:        url,
		TopColor:   topColor,
	}
}
