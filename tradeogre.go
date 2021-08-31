package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	BaseURL    *url.URL
	UserAgent  string
	httpClient *http.Client
}

func (c *Client) getMarkets() (Market, error) {
	url := c.BaseURL.String() + "/markets"
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var markets Market
	err = json.Unmarshal(body, &markets)

	return markets, err
}

func (c *Client) getOrderBook(market string) (OrderBook, error) {
	url := c.BaseURL.String() + "/orders/" + market
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return OrderBook{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return OrderBook{}, err
	}

	// Unmarshal the buy and sell to rawMssage that is later converted to string for orderBook
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

func (c *Client) getTicker(market string) (Ticker, error) {
	url := c.BaseURL.String() + "/ticker/" + market
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return Ticker{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("USer-Agent", c.UserAgent)
	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Ticker{}, err
	}

	defer resp.Body.Close()
	var ticker Ticker
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &ticker)
	return ticker, err
}
