// 用户分组管理.
package group

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type Group struct {
	Id        int64  `json:"id"`    // 分组id, 由微信分配
	Name      string `json:"name"`  // 分组名字, UTF8编码
	UserCount int    `json:"count"` // 分组内用户数量
}

// Create 创建分组.
func Create(clt *core.Client, name string) (group *Group, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/groups/create?access_token="

	var request struct {
		Group struct {
			Name string `json:"name"`
		} `json:"group"`
	}
	request.Group.Name = name

	var result struct {
		core.Error
		Group `json:"group"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	result.Group.UserCount = 0
	group = &result.Group
	return
}

// Delete 删除分组.
func Delete(clt *core.Client, groupId int64) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/groups/delete?access_token="

	var request struct {
		Group struct {
			Id int64 `json:"id"`
		} `json:"group"`
	}
	request.Group.Id = groupId

	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// Update 修改分组名.
func Update(clt *core.Client, groupId int64, name string) (err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/groups/update?access_token="

	var request struct {
		Group struct {
			Id   int64  `json:"id"`
			Name string `json:"name"`
		} `json:"group"`
	}
	request.Group.Id = groupId
	request.Group.Name = name

	var result core.Error
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// List 查询所有分组.
func List(clt *core.Client) (groups []Group, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/groups/get?access_token="

	var result struct {
		core.Error
		Groups []Group `json:"groups"`
	}
	if err = clt.GetJSON(incompleteURL, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	groups = result.Groups
	return
}
