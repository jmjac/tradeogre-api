# tradeogre-api

Simple wrapper for TradeOgre public API

Usage Example:
```go
t, _ := tradeogreapi.New(http.DefaultClient)
ticker, _ := t.GetTicker("BTC-XMR")    
fmt.Println(ticker.Price)    
```
