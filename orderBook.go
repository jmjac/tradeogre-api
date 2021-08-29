package main

//TODO: convert buy and sell to float64
type OrderBook struct {
	Success string `json:"success"`
	Buy     string `json:"buy"`
	Sell    string `json:"sell"`
}
