package main

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	url, err := url.Parse("https://tradeogre.com/api/v1/")
	if err != nil {
		fmt.Printf("err = %+v\n", err)
	}
	fmt.Printf("url = %+v\n", url)

	httpClient := http.DefaultClient

	c := Client{url, "test-api", httpClient}
	orderBook, err := c.getOrderBook("BTC-XMR")
	if err != nil {
		fmt.Printf("err = %+v\n", err)
	}
	fmt.Printf("orderBook = %+v\n", orderBook)
}
