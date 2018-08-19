package mock

import "fmt"

func (m *Mock) MarketTicker(opts ...MockParamsOption) (ret string, err error) {
	ma := setMockParams(opts)
	url := "https://api.fcoin.com/v2/market/ticker/" + ma.Symbol
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketDepth(opts ...MockParamsOption) (ret string, err error) {
	ma := setMockParams(opts)
	url := "https://api.fcoin.com/v2/market/depth/" + ma.Level + "/" + ma.Symbol
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketTrades(opts ...MockParamsOption) (ret string, err error) {
	ma := setMockParams(opts)
	url := "https://api.fcoin.com/v2/market/trades" + ma.Symbol
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketCandles(opts ...MockParamsOption) (ret string, err error) {
	ma := setMockParams(opts)
	url := "https://api.fcoin.com/v2/market/candles" + ma.Resolution + "/" + ma.Symbol
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
