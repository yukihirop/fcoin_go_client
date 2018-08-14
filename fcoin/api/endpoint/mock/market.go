package mock

import "fmt"

func (m *Mock) MarketTicker(opts ...MockEndpointOption) (ret string, err error) {
	ma := setMockEndpoint(opts)
	url := m.url("market", "MarketTicker") + "/" + ma.Symbol
	ret, err = m.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketDepth(opts ...MockEndpointOption) (ret string, err error) {
	ma := setMockEndpoint(opts)
	url := m.url("market", "MarketDepth") + "/" + ma.Level + "/" + ma.Symbol
	ret, err = m.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketTrades(opts ...MockEndpointOption) (ret string, err error) {
	ma := setMockEndpoint(opts)
	url := m.url("market", "MarketTrades") + "/" + ma.Symbol
	ret, err = m.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketCandles(opts ...MockEndpointOption) (ret string, err error) {
	ma := setMockEndpoint(opts)
	url := m.url("market", "MarketCandles") + "/" + ma.Resolution + "/" + ma.Symbol
	ret, err = m.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
