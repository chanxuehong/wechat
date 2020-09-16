package model

import (
	"encoding/json"
)

type NerType = string

const (
	NUMBER_NER            NerType = "number"
	DATETIME_POINT_NER    NerType = "datetime_point"
	DATETIME_INTERVAL_NER NerType = "datetime_interval"
	DATETIME_DURATION_NER NerType = "datetime_duration"
	DATETIME_REPEAT_NER   NerType = "datetime_repeat"
)

type Ner struct {
	Type NerType         `json:"type,omitempty"`
	Span []int           `json:"span,omitempty"`
	Text string          `json:"text,omitempty"`
	Norm json.RawMessage `json:"norm,omitempty"`
}

type NerNorm struct {
	Year   int    `json:"year,omitempty"`
	Month  int    `json:"month,omitempty"`
	Day    int    `json:"day,omitempty"`
	Hour   int    `json:"hour,omitempty"`
	Minute int    `json:"minute,omitempty"`
	Second int    `json:"second,omitempty"`
	Repeat string `json:"repeat,omitempty"`
}
