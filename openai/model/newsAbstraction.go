package model

type NewsAbstraction struct {
	Abstraction    string  `json:"abstraction,omitempty"`
	Classification bool    `json:"classification,omitempty"`
	Prob           float64 `json:"prob,omitempty"`
}
