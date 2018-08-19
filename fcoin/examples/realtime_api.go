package main

import (
	"fcoin_go_client/fcoin"
	"fmt"
)

func main() {
	client := fcoin.NewRealTimeClient()

	client.OnTicker("ethusdt", func(data map[string]interface{}) {
		fmt.Println(data)
	})

	client.OnDepth("ethusdt", "L20", func(data map[string]interface{}) {
		fmt.Println(data)
	})

	client.OnTrade("ethusdt", 5, func(data map[string]interface{}) {
		fmt.Println(data)
	})

	client.OnCandle("ethusdt", "MN", nil, func(data map[string]interface{}) {
		fmt.Println(data)
	})

	client.OnTopics(func(data map[string]interface{}) {
		fmt.Println(data)
	})

	client.OnHello(func(data map[string]interface{}) {
		fmt.Println(data)
	})
}
