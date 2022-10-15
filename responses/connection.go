package responses

type Connection struct {
	Event  string `json:"event"`
	Status string `json:"status"`
	Pair   string `json:"pair"`
}
