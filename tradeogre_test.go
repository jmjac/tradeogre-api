package tradeogreapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTest() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/markets/", handleMarkets)
	mux.HandleFunc("/orders/BTC-XMR", handleOrderBook)
	mux.HandleFunc("/ticker/BTC-XMR", handleTicker)
	mux.HandleFunc("/history/BTC-XMR", handleTradeHistory)
	server := httptest.NewServer(mux)
	return server
}

func TestListMarkets(t *testing.T) {
	server := setupTest()
	tradeogre := New(server.Client())
	tradeogre.baseURL = server.URL

	markets, err := tradeogre.ListMarkets()
	if err != nil {
		t.Fatal(err)
	}

	btcAeon, ok := (*markets)[0]["BTC-AEON"]
	if ok == false {
		t.Fatal("Expectd BTC-AEON to be present in markets")
	}
	if btcAeon.InitialPrice != 0.00022004 {
		t.Errorf("Got %v, expected 0.00022004 for InitialPrice", btcAeon.InitialPrice)
	}
	if btcAeon.Ask != 0.00025993 {
		t.Errorf("Got %v, expected 0.00025993 for Ask", btcAeon.Ask)
	}

	btcBtcp, ok := (*markets)[1]["BTC-BTCP"]
	if ok == false {
		t.Fatal("Expected BTC-BTCP to be present in markets")
	}
	if btcBtcp.High != 0.00379 {
		t.Errorf("Got %v, expected 0.00325 for High", btcBtcp.High)
	}
	if btcBtcp.Low != 0.0030001 {
		t.Errorf("Got %v, expected 0.00300010 for Low", btcBtcp.Low)
	}

}

func TestGetOrderBook(t *testing.T) {
	server := setupTest()
	defer server.Close()
	tradeogre := New(server.Client())
	tradeogre.baseURL = server.URL

	orderBook, err := tradeogre.GetOrderBook("BTC-XMR")
	if err != nil {
		t.Fatal(err)
	}

	if orderBook.Success == false {
		t.Fatal("Couldn't get orderbook for BTC-XMR")
	}
}

func TestGetTicker(t *testing.T) {
	server := setupTest()
	defer server.Close()
	tradeogre := New(server.Client())
	tradeogre.baseURL = server.URL

	ticker, err := tradeogre.GetTicker("BTC-XMR")
	if err != nil {
		t.Fatal(err)
	}

	if ticker.InitialPrice != 0.02502002 {
		t.Errorf("Got %v, expected 0.02502002", ticker.InitialPrice)
	}

}

func TestGetTradeHistory(t *testing.T) {
	server := setupTest()
	defer server.Close()
	tradeogre := New(server.Client())
	tradeogre.baseURL = server.URL

	tradeHistory, err := tradeogre.GetTradeHistory("BTC-XMR")
	if err != nil {
		t.Fatal(err)
	}

	if len(*tradeHistory) != 5 {
		t.Errorf("Got %v, expected 5", len(*tradeHistory))
	}

}
