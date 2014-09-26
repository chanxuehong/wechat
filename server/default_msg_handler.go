// @description wechat 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/chanxuehong/wechat for the canonical source repository
// @license     https://github.com/chanxuehong/wechat/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package server

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/chanxuehong/wechat/message/passive/request"
	"net/http"
	"sort"
)

var _ MsgHandler = DefaultMsgHandler{}

type DefaultMsgHandler struct {
	Token string
}

func (handler DefaultMsgHandler) Signature(timestamp, nonce string) (signature string) {
	strArray := sort.StringSlice{timestamp, nonce, handler.Token}
	strArray.Sort()

	n := len(timestamp) + len(nonce) + len(handler.Token)
	buf := make([]byte, 0, n)

	buf = append(buf, strArray[0]...)
	buf = append(buf, strArray[1]...)
	buf = append(buf, strArray[2]...)

	hashSumArray := sha1.Sum(buf)
	return hex.EncodeToString(hashSumArray[:])
}

func (handler DefaultMsgHandler) InvalidRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
}
func (handler DefaultMsgHandler) UnknownMsgHandler(w http.ResponseWriter, r *http.Request, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) TextMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Text, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) ImageMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Image, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) VoiceMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Voice, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) VideoMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Video, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) LocationMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Location, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) LinkMsgHandler(w http.ResponseWriter, r *http.Request, msg *request.Link, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) SubscribeEventHandler(w http.ResponseWriter, r *http.Request, msg *request.SubscribeEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) UnsubscribeEventHandler(w http.ResponseWriter, r *http.Request, msg *request.UnsubscribeEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) SubscribeByScanEventHandler(w http.ResponseWriter, r *http.Request, msg *request.SubscribeByScanEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) ScanEventHandler(w http.ResponseWriter, r *http.Request, msg *request.ScanEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) LocationEventHandler(w http.ResponseWriter, r *http.Request, msg *request.LocationEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MenuClickEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuClickEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MenuViewEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuViewEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MenuScanCodePushEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodePushEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MenuScanCodeWaitMsgEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuScanCodeWaitMsgEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MenuPicSysPhotoEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuPicSysPhotoEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MenuPicPhotoOrAlbumEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuPicPhotoOrAlbumEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MenuPicWeixinEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuPicWeixinEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MenuLocationSelectEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MenuLocationSelectEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MassSendJobFinishEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MassSendJobFinishEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) TemplateSendJobFinishEventHandler(w http.ResponseWriter, r *http.Request, msg *request.TemplateSendJobFinishEvent, rawXMLMsg []byte, timestamp int64) {
}
func (handler DefaultMsgHandler) MerchantOrderEventHandler(w http.ResponseWriter, r *http.Request, msg *request.MerchantOrderEvent, rawXMLMsg []byte, timestamp int64) {
}
