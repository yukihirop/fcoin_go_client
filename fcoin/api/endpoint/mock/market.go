package mock

import "fmt"

func (m *Mock) MarketTicker(opts ...MockEndpointOption) (ret string, err error) {
	ma := setMockEndpoint(opts)
	url := "https://api.fcoin.com/v2/market/ticker/" + ma.Symbol
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketDepth(opts ...MockEndpointOption) (ret string, err error) {
	ma := setMockEndpoint(opts)
	url := "https://api.fcoin.com/v2/market/depth/" + ma.Level + "/" + ma.Symbol
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketTrades(opts ...MockEndpointOption) (ret string, err error) {
	ma := setMockEndpoint(opts)
	url := "https://api.fcoin.com/v2/market/trades" + ma.Symbol
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketCandles(opts ...MockEndpointOption) (ret string, err error) {
	ma := setMockEndpoint(opts)
	url := "https://api.fcoin.com/v2/market/candles" + ma.Resolution + "/" + ma.Symbol
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
