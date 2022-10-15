package responses

type Ticker struct {
	ChannelID   int    `json:"channelID"`
	ChannelName string `json:"channelName"`
	Pair        string `json:"pair"`
	priceVal
}

type priceVal struct {
	A []interface{} `json:"a"` // best ask price, whole lot volume, lot volume
	B []interface{} `json:"b"` // best bid price, whole lot volume, lot volume
	C []string      `json:"c"` // close price, close volume
	H []string      `json:"h"` // high price: value today, value over last 24 hours
	L []string      `json:"l"` // low price: value today, value over last 24 hours
	O []string      `json:"o"` // open price: value today, value over last 24 hours
	P []string      `json:"p"` // volume weighted avg price: value today, value over last 24 hours
	T []int64       `json:"t"` // num of trades: value today, value over last 24 hours
	V []string      `json:"v"` // volume: value today, value over last 24 hours
}
