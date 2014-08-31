// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package template

// 模板消息数据结构, 分为 纯文本模板消息 和 带链接的模板消息
//
//  纯文本模板消息:
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
//
//  带链接的模板消息:
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
type Msg struct {
	ToUser     string `json:"touser"`      // 接受者OpenID
	TemplateId string `json:"template_id"` // 模版ID

	// 模版调用时需要的参数赋值内容，具体有哪些参数视模版内容而定
	// 一般是一个 struct 或者 map, 要求这个 Data 被 encoding/json 格式化后满足上面注释的格式.
	Data interface{} `json:"data"`

	// 下面两个字段是 带链接模版消息 才有的
	URL      string `json:"url,omitempty"`      // 用户点击后跳转的URL，该URL必须处于开发者在公众平台网站中设置的域中
	TopColor string `json:"topcolor,omitempty"` // 整个消息的颜色, 可以不设置
}

// 新建 纯文本模板消息
func NewMsg(touser, templateId string, data interface{}) *Msg {
	return &Msg{
		ToUser:     touser,
		TemplateId: templateId,
		Data:       data,
	}
}

// 新建 带链接的模板消息
func NewMsgWithLink(touser, templateId string, data interface{}, url string, topcolor string) *Msg {
	return &Msg{
		ToUser:     touser,
		TemplateId: templateId,
		Data:       data,
		URL:        url,
		TopColor:   topcolor,
	}
}
