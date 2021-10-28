package model

type Repo struct {
	Repo   string      `json:"repo"`
	Branch string      `json:"branch"`
	Data   []*Document `json:"data"`
}
