[![Tests](https://github.com/jmjac/tradeogre-api/actions/workflows/tests.yml/badge.svg)](https://github.com/jmjac/tradeogre-api/actions/workflows/tests.yml)
# tradeogre-api

Simple wrapper for TradeOgre public API [WIP]

Usage Example:
```go
t, _ := tradeogreapi.New(http.DefaultClient)
ticker, _ := t.GetTicker("BTC-XMR")    
fmt.Println(ticker.Price)  

markets, _ := t.ListMarkets()
for _, market := range *markets {
  fmt.Println(market)
}    
```
