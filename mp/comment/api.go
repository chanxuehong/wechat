package comment

import (
	"errors"
	"gopkg.in/chanxuehong/wechat.v2/mp/core"
)

const (
	openCommentUri         = "https://api.weixin.qq.com/cgi-bin/comment/open?access_token="
	closeCommentUri        = "https://api.weixin.qq.com/cgi-bin/comment/close?access_token="
	articleListCommentUri  = "https://api.weixin.qq.com/cgi-bin/comment/list?access_token="
	markSelectCommentUri   = "https://api.weixin.qq.com/cgi-bin/comment/markelect?access_token="
	unMarkSelectCommentUri = "https://api.weixin.qq.com/cgi-bin/comment/unmarkelect?access_token="
	deleteCommentUri       = "https://api.weixin.qq.com/cgi-bin/comment/delete?access_token="
	replyCommentUri        = "https://api.weixin.qq.com/cgi-bin/comment/reply/add?access_token="
	deleteReplyCommentUri  = "https://api.weixin.qq.com/cgi-bin/comment/reply/delete?access_token="
)

//打开已群发文章评论
func Open(clt *core.Client, msg_data_id, index int64) (err error) {
	var request = struct {
		MsgDataID int64 `json:"msg_data_id"`
		Index     int64 `json:"index"`
	}{
		MsgDataID: msg_data_id,
		Index:     index,
	}
	var result core.Error
	if err = clt.PostJSON(openCommentUri, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

//关闭已群发文章评论
func Close(clt *core.Client, msg_data_id, index int64) (err error) {
	var request = struct {
		MsgDataID int64 `json:"msg_data_id,omitempty"`
		Index     int64 `json:"index,omitempty"`
	}{
		MsgDataID: msg_data_id,
		Index:     index,
	}
	var result core.Error
	if err = clt.PostJSON(openCommentUri, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 查看指定文章的评论数据
func ArticleList(clt *core.Client, comment *Comment) (list []CommentListResult, count int32, err error) {
	var resp ArticleCommentListResp
	if err = clt.PostJSON(articleListCommentUri, comment, &resp); err != nil {
		return
	}
	if resp.ErrCode != core.ErrCodeOK {
		err = errors.New(resp.ErrMsg)
		return
	}
	list = resp.Comment
	count = resp.Total
	return
}

//将评论标记精选
func MarkSelect(clt *core.Client, msg_data_id, user_comment_id, index int64) (err error) {
	var request = struct {
		MsgDataID     int64 `json:"msg_data_id"`
		UserCommentID int64 `json:"user_comment_id"`
		Index         int64 `json:"index"`
	}{
		MsgDataID:     msg_data_id,
		UserCommentID: user_comment_id,
		Index:         index,
	}
	var result core.Error
	if err = clt.PostJSON(markSelectCommentUri, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

//将评论取消精选
func UnMarkSelect(clt *core.Client, msg_data_id, user_comment_id, index int64) (err error) {
	var request = struct {
		MsgDataID     int64 `json:"msg_data_id"`
		UserCommentID int64 `json:"user_comment_id"`
		Index         int64 `json:"index"`
	}{
		MsgDataID:     msg_data_id,
		UserCommentID: user_comment_id,
		Index:         index,
	}
	var result core.Error
	if err = clt.PostJSON(unMarkSelectCommentUri, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

//删除评论
func Delete(clt *core.Client, msg_data_id, user_comment_id, index int64) (err error) {
	var request = struct {
		MsgDataID     int64 `json:"msg_data_id"`
		UserCommentID int64 `json:"user_comment_id"`
		Index         int64 `json:"index"`
	}{
		MsgDataID:     msg_data_id,
		UserCommentID: user_comment_id,
		Index:         index,
	}
	var result core.Error
	if err = clt.PostJSON(deleteCommentUri, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

//回复评论
func Reply(clt *core.Client, msg_data_id, user_comment_id, index int64, content string) (err error) {
	var request = struct {
		MsgDataID     int64  `json:"msg_data_id"`
		UserCommentID int64  `json:"user_comment_id"`
		Index         int64  `json:"index"`
		Content       string `json:"content"`
	}{
		MsgDataID:     msg_data_id,
		UserCommentID: user_comment_id,
		Index:         index,
		Content:       content,
	}
	var result core.Error
	if err = clt.PostJSON(replyCommentUri, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}

// 删除回复
func DeleteReply(clt *core.Client,msg_data_id, user_comment_id, index int64) (err error) {
	var request = struct {
		MsgDataID     int64  `json:"msg_data_id"`
		UserCommentID int64  `json:"user_comment_id"`
		Index         int64  `json:"index"`
	}{
		MsgDataID:     msg_data_id,
		UserCommentID: user_comment_id,
		Index:         index,
	}
	var result core.Error
	if err = clt.PostJSON(deleteReplyCommentUri, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result
		return
	}
	return
}
