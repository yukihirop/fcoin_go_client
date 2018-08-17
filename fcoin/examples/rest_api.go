package main

import (
	"fcoin_go_client/fcoin"
	"fmt"
	"os"
)

func main() {
	apiKey := os.Getenv("FCOIN_API_KEY")
	secretKey := os.Getenv("FCOIN_SECRET_KEY")
	client := fcoin.NewAPI(fcoin.ApiKey(apiKey), fcoin.SecretKey(secretKey))

	// public
	publicServerTime, _ := client.PublicServerTime()
	publicCurrencies, _ := client.PublicCurrencies()
	publicSymbols, _ := client.PublicSymbols()
	fmt.Println(publicServerTime)
	fmt.Println(publicCurrencies)
	fmt.Println(publicSymbols)

	// market
	marketTicker, _ := client.MarketTicker("ethusdt")
	marketDepth, _ := client.MarketDepth("ethusdt", "L20")
	marketTrades, _ := client.MarketTrades("ethusdt")
	marketCandles, _ := client.MarketCandles("ethusdt", "MN")
	fmt.Println(marketTicker)
	fmt.Println(marketDepth)
	fmt.Println(marketTrades)
	fmt.Println(marketCandles)

	// orders
	//createOrderLimit, _ := client.CreateOrderLimit("ethusdt", "buy", 1, 0.001)
	orderList, _ := client.OrderList("ethusdt", "canceled", nil, nil, 20)
	referenceOrder, _ := client.ReferenceOrder("pUrzd_YxFQNBrQuo_XkQe1mdZY7IyzJNSyWxgkn3sL0=")
	cancelOrder, _ := client.CancelOrder("pUrzd_YxFQNBrQuo_XkQe1mdZY7IyzJNSyWxgkn3sL0=")
	orderMatchResults, _ := client.OrderMatchResults("pUrzd_YxFQNBrQuo_XkQe1mdZY7IyzJNSyWxgkn3sL0=")
	//fmt.Println(createOrderLimit)
	fmt.Println(orderList)
	fmt.Println(referenceOrder)
	fmt.Println(cancelOrder)
	fmt.Println(orderMatchResults)

	// accounts
	accountsBalance, _ := client.AccountsBalance()
	fmt.Println(accountsBalance)
}
