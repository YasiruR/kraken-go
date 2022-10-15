package responses

type GetAssetInfo struct {
	Error  []interface{}    `json:"error"`
	Result map[string]Asset `json:"result"`
}

type Asset struct {
	AClass          string `json:"aclass"`
	AltName         string `json:"altname"`
	Decimals        int    `json:"decimals"`
	DisplayDecimals int    `json:"display_decimals"`
}
