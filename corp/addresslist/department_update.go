// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package addresslist

// 更新部门的参数
//  NOTE: 如果非必须的字段未指定，则不更新该字段之前的设置值
type DepartmentUpdateParameters map[string]interface{}

func NewDepartmentUpdateParameters() DepartmentUpdateParameters {
	return make(map[string]interface{})
}

// 必须; 部门id

func (para DepartmentUpdateParameters) SetId(id int64) {
	para["id"] = id
}
func (para DepartmentUpdateParameters) DelId() {
	delete(para, "id")
}

// 非必须; 更新的部门名称。长度限制为0~64个字符。修改部门名称时指定该参数

func (para DepartmentUpdateParameters) SetName(name string) {
	para["name"] = name
}
func (para DepartmentUpdateParameters) DelName() {
	delete(para, "name")
}

// 非必须; 父亲部门id。根部门id为1

func (para DepartmentUpdateParameters) SetParentId(parentid int64) {
	para["parentid"] = parentid
}
func (para DepartmentUpdateParameters) DelParentId() {
	delete(para, "parentid")
}

// 非必须; 在父部门中的次序。从1开始，数字越大排序越靠后

func (para DepartmentUpdateParameters) SetOrder(order int) {
	if order > 0 {
		para["order"] = order
	}
}
func (para DepartmentUpdateParameters) DelOrder() {
	delete(para, "order")
}
