package tradeogreapi

import (
	"encoding/json"
	"net/http"
)

type TradeOgre struct {
	baseURL    string
	httpClient *http.Client
}

//New creates new TradeOgre object
func New(client *http.Client) *TradeOgre {
	if client == nil {
		client = http.DefaultClient
	}
	url := "https://tradeogre.com/api/v1"
	t := &TradeOgre{url, client}
	return t
}

func (t *TradeOgre) newRequest(method, path string) (*http.Request, error) {
	url := t.baseURL + path
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	return req, err
}

// ListMarkets gets a listing of all markets and basic information including current price, volume, high, low, bid and ask
func (t *TradeOgre) ListMarkets() (*MarketsMap, error) {
	req, err := t.newRequest("GET", "/markets")
	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var markets MarketsMap
	err = json.NewDecoder(resp.Body).Decode(&markets)
	return &markets, err
}

// GetOrderBook gets the current order book for market
func (t *TradeOgre) GetOrderBook(market string) (*OrderBook, error) {
	req, err := t.newRequest("GET", "/orders/"+market)
	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var orderBook OrderBook
	err = json.NewDecoder(resp.Body).Decode(&orderBook)
	return &orderBook, err
}

// GetTicker gets the ticker for market. Volume, high, and low are in the last 24 hours, initialprice is the price from 24 hours ago.
func (t *TradeOgre) GetTicker(market string) (*Ticker, error) {
	req, err := t.newRequest("GET", "/ticker/"+market)
	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var ticker Ticker
	err = json.NewDecoder(resp.Body).Decode(&ticker)
	return &ticker, err
}

// GetTradeHistory gets the history of the last trades on market limited to 100 of the most recent trades. The date is a Unix UTC timestamp.
func (t *TradeOgre) GetTradeHistory(market string) (*TradeHistory, error) {
	req, err := t.newRequest("GET", "/history/"+market)
	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	var tradeHistory TradeHistory
	err = json.NewDecoder(resp.Body).Decode(&tradeHistory)
	return &tradeHistory, err
}
