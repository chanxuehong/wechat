package liberary

type Template struct {
	Id       string    `json:"id"`
	Title    string    `json:"title"`
	Keywords []Keyword `json:"keyword_list,omitempty"`
}

type Keyword struct {
	Id      uint64 `json:"keyword_id"`
	Name    string `json:"name"`
	Example string `json:"example"`
}
