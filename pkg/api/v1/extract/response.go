package extract

type Response struct {
	Repo    string         `json:"repo"`
	Branch  string         `json:"branch"`
	Data    []*ResponeData `json:"data"`
	Warning []string       `json:"warning"`
	Error   []string       `json:"error"`
}

type ResponeData struct {
	Path  string `json:"path"`
	Title string `json:"title"`
}
