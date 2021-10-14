package tradeogreapi

type OrderBook struct {
	Success bool              `json:"success,string"`
	Buy     map[string]string `json:"buy"`
	Sell    map[string]string `json:"sell"`
}
