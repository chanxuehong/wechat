// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package addresslist

// 更新成员的参数
//  NOTE: 如果非必须的字段未指定，则不更新该字段之前的设置值
type UserUpdateParameters map[string]interface{}

func NewUserUpdateParameters() UserUpdateParameters {
	return make(map[string]interface{})
}

// 必须; 员工UserID。对应管理端的帐号，企业内必须唯一

func (para UserUpdateParameters) SetUserId(userid string) {
	para["userid"] = userid
}
func (para UserUpdateParameters) DelUserId() {
	delete(para, "userid")
}

// 非必须; 成员名称。长度为0~64个字符

func (para UserUpdateParameters) SetName(name string) {
	para["name"] = name
}
func (para UserUpdateParameters) DelName() {
	delete(para, "name")
}

// 非必须; 成员所属部门id列表。注意，每个部门的直属员工上限为1000个

func (para UserUpdateParameters) SetDepartment(department []int64) {
	if department == nil {
		department = make([]int64, 0)
	}
	para["department"] = department
}
func (para UserUpdateParameters) DelDepartment() {
	delete(para, "department")
}

// 非必须; 职位信息。长度为0~64个字符

func (para UserUpdateParameters) SetPosition(position string) {
	para["position"] = position
}
func (para UserUpdateParameters) DelPosition() {
	delete(para, "position")
}

// 非必须; 手机号码。企业内必须唯一，更新后的成员mobile/weixinid/email三者不能同时为空

func (para UserUpdateParameters) SetMobile(mobile string) {
	para["mobile"] = mobile
}
func (para UserUpdateParameters) DelMobile() {
	delete(para, "mobile")
}

// 非必须; 性别。gender=0表示男，=1表示女。默认gender=0

func (para UserUpdateParameters) SetGender(gender int) {
	para["gender"] = gender
}
func (para UserUpdateParameters) DelGender() {
	delete(para, "gender")
}

// 非必须; 办公电话。长度为0~64个字符。必须企业内唯一

func (para UserUpdateParameters) SetTel(tel string) {
	para["tel"] = tel
}
func (para UserUpdateParameters) DelTel() {
	delete(para, "tel")
}

// 非必须; 邮箱。长度为0~64个字符。企业内必须唯一

func (para UserUpdateParameters) SetEmail(email string) {
	para["email"] = email
}
func (para UserUpdateParameters) DelEmail() {
	delete(para, "email")
}

// 非必须; 微信号。企业内必须唯一

func (para UserUpdateParameters) SetWeixinId(weixinid string) {
	para["weixinid"] = weixinid
}
func (para UserUpdateParameters) DelWeixinId() {
	delete(para, "weixinid")
}

// 非必须; 启用/禁用成员

func (para UserUpdateParameters) SetEnable(enable bool) {
	if enable {
		para["enable"] = 1
	} else {
		para["enable"] = 0
	}
}
func (para UserUpdateParameters) DelEnable() {
	delete(para, "enable")
}
