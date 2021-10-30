package model

type Repo struct {
	Repo    string      `json:"repo"`
	Branch  string      `json:"branch"`
	Data    []*Document `json:"data"`
	Warning []string    `json:"warning"`
	Error   []string    `json:"error"`
}
