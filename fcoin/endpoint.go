package fcoin

import "fcoin_go_client/fcoin/api/endpoint"

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
}

func endpointConfig(c *Configure) *(endpoint.Configure) {
	var config endpoint.Configure
	config.Adapter = c.Adapter
	config.Endpoint = c.Endpoint
	config.UserAgent = c.UserAgent
	config.Proxy = c.Proxy
	config.ApiKey = c.ApiKey
	config.SecretKey = c.SecretKey
	return &config
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
