package qrcode

import (
	"github.com/chanxuehong/wechat/mp/core"
)

type TempQrcode struct {
	ExpireSeconds int `json:"expire_seconds,omitempty"`
	PermQrcode
}

type PermQrcode struct {
	Ticket string `json:"ticket"`
	URL    string `json:"url"`
}

// CreateTempQrcode 创建临时二维码.
//  sceneId:       场景值ID, 为32位非0整型
//  expireSeconds: 二维码有效时间, 以秒为单位
func CreateTempQrcode(clt *core.Client, sceneId int32, expireSeconds int) (qrcode *TempQrcode, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="

	var request struct {
		ExpireSeconds int    `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Scene struct {
				SceneId int32 `json:"scene_id"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	request.ExpireSeconds = expireSeconds
	request.ActionName = "QR_SCENE"
	request.ActionInfo.Scene.SceneId = sceneId

	var result struct {
		core.Error
		TempQrcode
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	qrcode = &result.TempQrcode
	return
}

// CreateStrSceneTempQrcode 创建临时二维码.
//  sceneStr:      场景值ID(字符串形式的ID), 字符串类型, 长度限制为1到64
//  expireSeconds: 二维码有效时间, 以秒为单位
func CreateStrSceneTempQrcode(clt *core.Client, sceneStr string, expireSeconds int) (qrcode *TempQrcode, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="

	var request struct {
		ExpireSeconds int    `json:"expire_seconds"`
		ActionName    string `json:"action_name"`
		ActionInfo    struct {
			Scene struct {
				SceneStr string `json:"scene_str"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	request.ExpireSeconds = expireSeconds
	request.ActionName = "QR_STR_SCENE"
	request.ActionInfo.Scene.SceneStr = sceneStr

	var result struct {
		core.Error
		TempQrcode
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	qrcode = &result.TempQrcode
	return
}

// CreatePermQrcode 创建永久二维码
//  sceneId: 场景值ID
func CreatePermQrcode(clt *core.Client, sceneId int32) (qrcode *PermQrcode, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="

	var request struct {
		ActionName string `json:"action_name"`
		ActionInfo struct {
			Scene struct {
				SceneId int32 `json:"scene_id"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	request.ActionName = "QR_LIMIT_SCENE"
	request.ActionInfo.Scene.SceneId = sceneId

	var result struct {
		core.Error
		PermQrcode
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	qrcode = &result.PermQrcode
	return
}

// CreateStrScenePermQrcode 创建永久二维码
//  sceneStr: 场景值ID(字符串形式的ID), 字符串类型, 长度限制为1到64
func CreateStrScenePermQrcode(clt *core.Client, sceneStr string) (qrcode *PermQrcode, err error) {
	const incompleteURL = "https://api.weixin.qq.com/cgi-bin/qrcode/create?access_token="

	var request struct {
		ActionName string `json:"action_name"`
		ActionInfo struct {
			Scene struct {
				SceneStr string `json:"scene_str"`
			} `json:"scene"`
		} `json:"action_info"`
	}
	request.ActionName = "QR_LIMIT_STR_SCENE"
	request.ActionInfo.Scene.SceneStr = sceneStr

	var result struct {
		core.Error
		PermQrcode
	}
	if err = clt.PostJSON(incompleteURL, &request, &result); err != nil {
		return
	}
	if result.ErrCode != core.ErrCodeOK {
		err = &result.Error
		return
	}
	qrcode = &result.PermQrcode
	return
}
