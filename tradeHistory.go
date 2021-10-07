package tradeogreapi

type tradeHistory []tradeHistoryValue

type tradeHistoryValue struct {
	Date     int64  `json:"date"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}
