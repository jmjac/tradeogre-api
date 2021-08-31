package main

type TradeHistory []TradeHistoryValue

type TradeHistoryValue struct {
	Date     int64  `json:"date"`
	Type     string `json:"type"`
	Price    string `json:"price"`
	Quantity string `json:"quantity"`
}
