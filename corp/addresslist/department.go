// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package addresslist

// 创建部门的参数
type DepartmentCreateParameters struct {
	Name     string `json:"name"`            // 部门名称。长度限制为1~64个字符
	ParentId int64  `json:"parentid"`        // 父亲部门id。根部门id为1
	Order    int    `json:"order,omitempty"` // 在父部门中的次序。从1开始，数字越大排序越靠后。
}

// 部门, 获取部门列表返回的单元结构
type Department struct {
	Id       int64  `json:"id"`
	Name     string `json:"name"`
	ParentId int64  `json:"parentid"`
}
