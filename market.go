package tradeogreapi

type MarketsMap []map[string]Market

type Market struct {
	InitialPrice float64 `json:"initialprice,string"`
	Price        float64 `json:"price,string"`
	High         float64 `json:"high,string"`
	Low          float64 `json:"low,string"`
	Volume       float64 `json:"volume,string"`
	Bid          float64 `json:"bid,string"`
	Ask          float64 `json:"ask,string"`
}
