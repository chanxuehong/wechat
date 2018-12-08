// 客服聊天记录接口
package record

import (
	"fmt"

	"github.com/chanxuehong/wechat/mp/core"
)

type Record struct {
	Worker    string `json:"worker"`   // 客服账号
	OpenId    string `json:"openid"`   // 用户的标识, 对当前公众号唯一
	OperCode  int    `json:"opercode"` // 操作ID(会话状态)
	Timestamp int64  `json:"time"`     // 操作时间, UNIX时间戳
	Text      string `json:"text"`     // 聊天记录
}

type GetRequest struct {
	StartTime int64  `json:"starttime"`        // 查询开始时间, UNIX时间戳
	EndTime   int64  `json:"endtime"`          // 查询结束时间, UNIX时间戳, 每次查询不能跨日查询
	PageIndex int    `json:"pageindex"`        // 查询第几页, 从1开始
	PageSize  int    `json:"pagesize"`         // 每页大小, 每页最多拉取50条
	OpenId    string `json:"openid,omitempty"` // 普通用户的标识, 对当前公众号唯一
}

// Get 获取客服聊天记录
func Get(clt *core.Client, request *GetRequest) (list []Record, err error) {
	const incompleteURL = "https://api.weixin.qq.com/customservice/msgrecord/getrecord?access_token="

	if request.PageIndex < 1 {
		err = fmt.Errorf("Incorrect request.PageIndex: %d", request.PageIndex)
		return
	}
	if request.PageSize <= 0 {
		err = fmt.Errorf("Incorrect request.PageSize: %d", request.PageSize)
		return
	}

	var result struct {
		core.Error
		RecordList []Record `json:"recordlist"`
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	list = result.RecordList
	return
}
