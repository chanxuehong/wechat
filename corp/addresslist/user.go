// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package addresslist

// 创建成员的参数
type UserCreateParameters struct {
	UserId     string  `json:"userid"`               // 必须;  员工UserID。对应管理端的帐号，企业内必须唯一
	Name       string  `json:"name"`                 // 必须;  成员名称。长度为1~64个字符
	Department []int64 `json:"department,omitempty"` // 非必须; 成员所属部门id列表。注意，每个部门的直属员工上限为1000个
	Position   string  `json:"position,omitempty"`   // 非必须; 职位信息。长度为0~64个字符
	Mobile     string  `json:"mobile,omitempty"`     // 非必须; 手机号码。企业内必须唯一，mobile/weixinid/email三者不能同时为空
	Gender     int     `json:"gender"`               // 非必须; 性别。gender=0表示男，=1表示女。默认gender=0
	Tel        string  `json:"tel,omitempty"`        // 非必须; 办公电话。长度为0~64个字符
	Email      string  `json:"email,omitempty"`      // 非必须; 邮箱。长度为0~64个字符。企业内必须唯一
	WeixinId   string  `json:"weixinid,omitempty"`   // 非必须; 微信号。企业内必须唯一
}

// 获取成员得到的信息
type UserInfo struct {
	Id         string  `json:"userid"`               // 员工UserID
	Name       string  `json:"name"`                 // 成员名称
	Department []int64 `json:"department,omitempty"` // 成员所属部门id列表
	Position   string  `json:"position"`             // 职位信息
	Mobile     string  `json:"mobile"`               // 手机号码
	Gender     int     `json:"gender"`               // 性别。gender=0表示男，=1表示女
	Tel        string  `json:"tel"`                  // 办公电话
	Email      string  `json:"email"`                // 邮箱
	WeixinId   string  `json:"weixinid"`             // 微信号
	Avatar     string  `json:"avatar"`               // 头像url。注：如果要获取小图将url最后的"/0"改成"/64"即可
	Status     int     `json:"status"`               // 关注状态: 1=已关注，2=已冻结，4=未关注
}

type UserInfoBase struct {
	Id   string `json:"userid"` // 员工UserID
	Name string `json:"name"`   // 成员名称
}
