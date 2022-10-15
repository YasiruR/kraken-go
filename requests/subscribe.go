package requests

type Subscribe struct {
	Event        string       `json:"event"`
	Pair         []string     `json:"pair"`
	Subscription subscription `json:"subscription"`
}

type subscription struct {
	Name string `json:"name"`
	//Depth int    `json:"depth"`
	//Interval    int    `json:"interval"`
	//RateCounter bool   `json:"ratecounter"`
	//Snapshot    bool   `json:"snapshot"`
	//Token       string `json:"token"`
}
