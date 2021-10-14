package tradeogreapi

type TradeHistory []tradeHistoryValue

type tradeHistoryValue struct {
	Date     int64   `json:"date"`
	Type     string  `json:"type"`
	Price    float64 `json:"price,string"`
	Quantity float64 `json:"quantity,string"`
}
