package model

type SimilarQuery struct {
	Question string  `json:"question,omitempty"`
	Score    float64 `json:"score,omitempty"`
	Source   string  `json:"source,omitempty"`
}
