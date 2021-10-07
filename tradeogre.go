package tradeogreapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TradeOgre struct {
	baseURL    *url.URL
	httpClient *http.Client
}

func New() (*TradeOgre, error) {
	url, err := url.Parse("https://tradeogre.com/api/v1")
	t := &TradeOgre{url, http.DefaultClient}
	return t, err
}

func (t *TradeOgre) newRequest(method, path string) (*http.Request, error) {
	url := t.baseURL.String() + path
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	return req, err
}

// ListMarkets returns tickers and prices for all tickers on tradeogre
func (t *TradeOgre) ListMarkets() (market, error) {
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
	var markets market
	err = json.Unmarshal(body, &markets)

	return markets, err
}

func (t *TradeOgre) GetOrderBook(market string) (orderBook, error) {
	req, err := t.newRequest("GET", "/orders/"+market)
	if err != nil {
		return orderBook{}, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return orderBook{}, err
	}

	// Unmarshal the buy and sell to rawMessage that is later converted to string for orderBook
	var tempBook struct {
		Success string          `json:"success"`
		Buy     json.RawMessage `json:"buy"`
		Sell    json.RawMessage `json:"sell"`
		this    json.Number     `json:"this"`
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &tempBook)

	var ob orderBook
	ob.Buy = string(tempBook.Buy)
	ob.Sell = string(tempBook.Sell)
	ob.Success = tempBook.Success

	return ob, err
}

func (t *TradeOgre) GetTicker(market string) (ticker, error) {
	req, err := t.newRequest("GET", "/ticker/"+market)
	if err != nil {
		return ticker{}, err
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return ticker{}, err
	}

	defer resp.Body.Close()
	var tick ticker
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("body = %+v\n", string(body))
	err = json.Unmarshal(body, &tick)
	return tick, err
}

func (t *TradeOgre) GetTradeHistory(market string) (tradeHistory, error) {
	req, err := t.newRequest("GET", "/history/"+market)
	if err != nil {
		return tradeHistory{}, nil
	}

	resp, err := t.httpClient.Do(req)
	if err != nil {
		return tradeHistory{}, nil
	}

	defer resp.Body.Close()
	var th tradeHistory
	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &th)
	return th, err
}
