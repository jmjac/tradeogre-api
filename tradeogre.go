package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TradeOgre struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

func (t *TradeOgre) newRequest(method, path string) (*http.Request, error) {
	url := t.BaseURL.String() + path
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", t.UserAgent)
	return req, err
}

func (t *TradeOgre) getMarkets() (Market, error) {
	req, err := t.newRequest("GET", "/markets")
	if err != nil {
		return nil, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var markets Market
	err = json.Unmarshal(body, &markets)

	return markets, err
}

func (t *TradeOgre) getOrderBook(market string) (OrderBook, error) {
	req, err := t.newRequest("GET", "/orders/"+market)
	if err != nil {
		return OrderBook{}, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return OrderBook{}, err
	}

	// Unmarshal the buy and sell to rawMessage that is later converted to string for orderBook
	var tempBook struct {
		Success string          `json:"success"`
		Buy     json.RawMessage `json:"buy"`
		Sell    json.RawMessage `json:"sell"`
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &tempBook)

	var orderBook OrderBook
	orderBook.Buy = string(tempBook.Buy)
	orderBook.Sell = string(tempBook.Sell)
	orderBook.Success = tempBook.Success

	return orderBook, err
}

func (t *TradeOgre) getTicker(market string) (Ticker, error) {
	req, err := t.newRequest("GET", "/ticker/"+market)
	if err != nil {
		return Ticker{}, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return Ticker{}, err
	}

	defer resp.Body.Close()
	var ticker Ticker
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("body = %+v\n", string(body))
	err = json.Unmarshal(body, &ticker)
	return ticker, err
}

func (t *TradeOgre) getTradeHistory(market string) (TradeHistory, error) {
	req, err := t.newRequest("GET", "/history/"+market)
	if err != nil {
		return TradeHistory{}, nil
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return TradeHistory{}, nil
	}

	defer resp.Body.Close()
	var tradeHistory TradeHistory
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &tradeHistory)
	return tradeHistory, err
}
