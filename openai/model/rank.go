package model

type RankResult struct {
	List       []Rank `json:"results,omitempty"`
	ExactMatch bool   `json:"exact_match,omitempty"`
}

type Rank struct {
	Question string  `json:"question,omitempty"`
	Score    float64 `json:"score,omitempty"`
}
