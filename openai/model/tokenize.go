package model

type TokenizeResult struct {
	Words       []string `json:"words,omitempty"`
	POSs        []int    `json:"POSs,omitempty"`
	WordsMix    []string `json:"words_mix,omitempty"`
	POSsMix     []int    `json:"POSs_mix,omitempty"`
	Entities    []string `json:"entities,omitempty"`
	EntityTypes []int    `json:"entity_types,omitempty"`
}
