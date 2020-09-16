package model

import (
	"encoding/json"
)

type AnswerType = string

const (
	TEXT  AnswerType = "text"
	MUSIC AnswerType = "music"
	NEWS  AnswerType = "news"
)

type Status = string

const (
	FAQ         Status = "FAQ"
	NOMATCH     Status = "NOMATCH"
	CONTEXT_FAQ Status = "CONTEXT_FAQ"
	GENERAL_FAQ Status = "GENERAL_FAQ"
)

type NLUResult struct {
	RequestID           uint64              `json:"request_id,omitempty"`
	Ret                 int                 `json:"ret,omitempty"`
	AnsNodeID           uint64              `json:"ans_node_id,omitempty"`
	AnsNodeName         string              `json:"ans_node_name,omitempty"`
	SessionID           string              `json:"session_id,omitempty"`
	SceneStatus         string              `json:"scene_status,omitempty"`
	Title               string              `json:"title,omitempty"`
	Answer              string              `json:"answer,omitempty"`
	AnswerOpen          int                 `json:"answer_open,omitempty"`
	AnswerType          AnswerType          `json:"answer_type,omitempty"`
	Article             string              `json:"article,omitempty"`
	Confidence          int                 `json:"confidence,omitempty"`
	CreateTime          json.Number         `json:"create_time,omitempty"`
	DialogSessionStatus string              `json:"dialog_session_status,omitempty"`
	DialogStatus        string              `json:"dialog_status,omitempty"`
	Event               string              `json:"event,omitempty"`
	IntentConfirmStatus string              `json:"intent_confirm_status,omitempty"`
	ListOptions         bool                `json:"list_options,omitempty"`
	MsgID               string              `json:"msg_id,omitempty"`
	Opening             string              `json:"opening,omitempty"`
	Msg                 []Message           `json:"msg,omitempty"`
	FromUserName        string              `json:"from_user_name,omitempty"`
	ToUserName          string              `json:"to_user_name,omitempty"`
	Status              Status              `json:"status,omitempty"`
	SlotInfo            []map[string]string `json:"slot_info,omitempty"`
	Slots               []Slot              `json:"slots_info,omitempty"`
	TakeOptionsOnly     bool                `json:"take_options_only,omitempty"`
	Query               string              `json:"query,omitempty"`
	BidStat             *BidStat            `json:"bid_stat,omitempty"`
}

type Message struct {
	RequestID       uint64 `json:"request_id,omitempty"`
	AnsNodeID       uint64 `json:"ans_node_id,omitempty"`
	AnsNodeName     string `json:"ans_node_name,omitempty"`
	Article         string `json:"article,omitempty"`
	Confidence      int    `json:"confidence,omitempty"`
	Content         string `json:"content,omitempty"`
	DebugInfo       string `json:"debug_info,omitempty"`
	Event           string `json:"event,omitempty"`
	ListOptions     bool   `json:"list_options,omitempty"`
	Opening         string `json:"opening,omitempty"`
	MsgType         string `json:"msg_type,omitempty"`
	RespTitle       string `json:"resp_title,omitempty"`
	SessionID       string `json:"session_id,omitempty"`
	SceneStatus     string `json:"scene_status,omitempty"`
	Status          Status `json:"status,omitempty"`
	TakeOptionsOnly bool   `json:"take_options_only,omitempty"`
}

type Slot struct {
	AllNerTypes   string `json:"all_ner_types,omitempty"`
	ConfirmStatus string `json:"confirm_status,omitempty"`
	Start         int    `json:"start,omitempty"`
	End           int    `json:"end,omitempty"`
	EntityType    string `json:"entity_type,omitempty"`
	Norm          string `json:"norm,omitempty"`
	NormDetail    string `json:"norm_detail,omitempty"`
	SlotName      string `json:"slot_name,omitempty"`
	SlotValue     string `json:"slot_value,omitempty"`
}

type BidStat struct {
	ErrMsg    string `json:"err_msg,omitempty"`
	CurrTime  string `json:"curr_time,omitempty"`
	LastTime  string `json:"last_time,omitempty"`
	LastValid bool   `json:"last_valid,omitempty"`
	UpRet     int    `json:"up_ret,omitempty"`
}
