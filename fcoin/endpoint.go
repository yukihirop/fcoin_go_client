package fcoin

import (
	"fcoin_go_client/fcoin/api/endpoint"
)

type Endpoint interface {
	// public
	PublicServerTime() (string, error)
	PublicCurrencies() (string, error)
	PublicSymbols() (string, error)

	// market
	MarketTicker(string) (string, error)
	MarketDepth(string, string) (string, error)
	MarketTrades(string) (string, error)
	MarketCandles(string, string) (string, error)

	// orders
	CreateOrderLimit(string, string, float32, float32) (string error)
	OrderList(string, string, interface{}, interface{}, interface{}) (string, error)
	ReferenceOrder(string) (string, error)
	CancelOrder(string) (string, error)
	OrderMatchResults(string) (string, error)

	// accounts
	AccountsBalance() (string, error)
}

// //http://horie1024.hatenablog.com/entry/2014/08/25/012123
// http://text.baldanders.info/golang/interface/
func (c *Configure) PublicServerTime() (ret string, err error) {
	ret, err = endpoint.PublicServerTime(endpointConfig(c))
	return
}

func (c *Configure) PublicCurrencies() (ret string, err error) {
	ret, err = endpoint.PublicCurrencies(endpointConfig(c))
	return
}

func (c *Configure) PublicSymbols() (ret string, err error) {
	ret, err = endpoint.PublicSymbols(endpointConfig(c))
	return
}

func (c *Configure) MarketTicker(symbol string) (ret string, err error) {
	ret, err = endpoint.MarketTicker(endpointConfig(c), endpoint.Symbol(symbol))
	return
}

func (c *Configure) MarketDepth(symbol string, level string) (ret string, err error) {
	ret, err = endpoint.MarketDepth(endpointConfig(c), endpoint.Symbol(symbol), endpoint.Level(level))
	return
}

func (c *Configure) MarketTrades(symbol string) (ret string, err error) {
	ret, err = endpoint.MarketTrades(endpointConfig(c), endpoint.Symbol(symbol))
	return
}

func (c *Configure) MarketCandles(symbol string, resolution string) (ret string, err error) {
	ret, err = endpoint.MarketCandles(endpointConfig(c), endpoint.Symbol(symbol), endpoint.Resolution(resolution))
	return
}

func (c *Configure) CreateOrderLimit(symbol string, side string, price float32, amount float32) (ret string, err error) {
	ret, err = endpoint.CreateOrderLimit(endpointConfig(c), endpoint.Symbol(symbol), endpoint.Side(side), endpoint.Price(price), endpoint.Amount(amount))
	return
}

func (c *Configure) OrderList(symbol string, states string, pageBefore interface{}, pageAfter interface{}, perPage interface{}) (ret string, err error) {
	pb, ok_pb := pageBefore.(int)
	pa, ok_pa := pageAfter.(int)
	pp, ok_pp := perPage.(int)
	if ok_pb && ok_pa && ok_pp {
		ret, err = endpoint.OrderList(endpointConfig(c), endpoint.Symbol(symbol), endpoint.States(states), endpoint.PageBefore(pb), endpoint.PageAfter(pa), endpoint.PerPage(pp))
	} else if !ok_pb && !ok_pa && ok_pp {
		ret, err = endpoint.OrderList(endpointConfig(c), endpoint.Symbol(symbol), endpoint.States(states), endpoint.PerPage(pp))
	} else if !ok_pb && !ok_pa && !ok_pp {
		ret, err = endpoint.OrderList(endpointConfig(c), endpoint.Symbol(symbol), endpoint.States(states))
	}

	return
}

func (c *Configure) ReferenceOrder(orderId string) (ret string, err error) {
	ret, err = endpoint.ReferenceOrder(endpointConfig(c), endpoint.OrderId(orderId))
	return
}

func (c *Configure) CancelOrder(orderId string) (ret string, err error) {
	ret, err = endpoint.CancelOrder(endpointConfig(c), endpoint.OrderId(orderId))
	return
}

func (c *Configure) OrderMatchResults(orderId string) (ret string, err error) {
	ret, err = endpoint.OrderMatchResults(endpointConfig(c), endpoint.OrderId(orderId))
	return
}

func (c *Configure) AccountsBalance() (ret string, err error) {
	ret, err = endpoint.AccountsBalance(endpointConfig(c))
	return
}
