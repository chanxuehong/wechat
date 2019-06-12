package comment

import "gopkg.in/chanxuehong/wechat.v2/mp/core"

type Comment struct {
	base
	//查看指定文章的评论数据
	Begin int  `json:"begin,omitempty"` //必填  起始位置
	Count int  `json:"count,omitempty"` //必填  获取数目（>=50会被拒绝）
	Type  *int `json:"type,omitempty"`  //必填  type=0 普通评论&精选评论 type=1 普通评论 type=2 精选评论
	// 论标记精选  评论取消精选  删除评论  删除回复
	UserCommentID int `json:"user_comment_id,omitempty"` //必填  用户评论id
	//回复评论
	Content string `json:"content,omitempty"` // 必填,   回复内容
}

type base struct {
	// 基本参数
	MsgDataID int64 `json:"msg_data_id,omitempty"` //必填   群发返回的msg_data_id
	Index     int `json:"index,omitempty"`       //多图文时，用来指定第几篇图文，从0开始，不带默认操作该msg_data_id的第一篇图文
}

// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++

// 文章回复
type ArticleCommentListResp struct {
	core.Error
	Total   int32               `json:"total,omitempty"`
	Comment []CommentListResult `json:"comment,omitempty"`
}

type CommentListResult struct {
	UserCommentID int    `json:"user_comment_id,omitempty"`
	OpenID        string `json:"open_id,omitempty"`
	CreateTime    int64  `json:"create_time,omitempty"`
	Content       string `json:"content,omitempty"`
	CommentType   int    `json:"comment_type,omitempty"`
	Reply         *ReplyComment `json:"reply,omitempty"`
}

type ReplyComment struct {
	Content    string `json:"content,omitempty"`
	CreateTime int64  `json:"create_time,omitempty"`
}
