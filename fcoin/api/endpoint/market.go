package endpoint

import (
	"fcoin_go_client/fcoin/api"
	"fmt"
)

type Market struct {
	Symbol     string
	Level      string
	Resolution string
}

type MarketOption func(*Market)
type MarketOptions []MarketOption

func Symbol(s string) MarketOption {
	return func(m *Market) {
		m.Symbol = s
	}
}

func Level(l string) MarketOption {
	return func(m *Market) {
		m.Level = l
	}
}

func Resolution(r string) MarketOption {
	return func(m *Market) {
		m.Resolution = r
	}
}

//http://horie1024.hatenablog.com/entry/2014/08/25/012123
func MarketTicker(c *Configure, opts ...MarketOption) (ret string, err error) {
	ma := setMarket(opts)
	url := GetPath(c.Endpoint, "MarketTicker") + "/" + ma.Symbol
	ret, err = api.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketDepth(c *Configure, opts ...MarketOption) (ret string, err error) {
	ma := setMarket(opts)
	url := GetPath(c.Endpoint, "MarketDepth") + "/" + ma.Symbol + "/" + ma.Level
	ret, err = api.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketTrades(c *Configure, opts ...MarketOption) (ret string, err error) {
	ma := setMarket(opts)
	url := GetPath(c.Endpoint, "MarketTrades") + "/" + ma.Symbol
	ret, err = api.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketCandles(c *Configure, opts ...MarketOption) (ret string, err error) {
	ma := setMarket(opts)
	url := GetPath(c.Endpoint, "MarketCandles") + "/" + ma.Resolution + "/" + ma.Symbol
	ret, err = api.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func setMarket(opts MarketOptions) (ma *Market) {
	ma = &Market{}
	for _, opt := range opts {
		opt(ma)
	}
	return
}
