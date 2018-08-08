package mock

import "fmt"

type Market struct {
	Symbol     string
	Level      string
	Resolution string
}

type Option func(*Market)
type Options []Option

func Symbol(s string) Option {
	return func(m *Market) {
		m.Symbol = s
	}
}

func Level(l string) Option {
	return func(m *Market) {
		m.Level = l
	}
}

func Resolution(r string) Option {
	return func(m *Market) {
		m.Resolution = r
	}
}

func (m *Mock) MarketTicker(opts ...Option) (ret string, err error) {
	ma := setMarket(opts)
	url := m.url("market", "MarketTicker") + "/" + ma.Symbol
	ret, err = m.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketDepth(opts ...Option) (ret string, err error) {
	ma := setMarket(opts)
	url := m.url("market", "MarketDepth") + "/" + ma.Level + "/" + ma.Symbol
	ret, err = m.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketTrades(opts ...Option) (ret string, err error) {
	ma := setMarket(opts)
	url := m.url("market", "MarketTrades") + "/" + ma.Symbol
	ret, err = m.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) MarketCandles(opts ...Option) (ret string, err error) {
	ma := setMarket(opts)
	url := m.url("market", "MarketCandles") + "/" + ma.Resolution + "/" + ma.Symbol
	ret, err = m.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func setMarket(opts Options) (m *Market) {
	m = &Market{}
	for _, opt := range opts {
		opt(m)
	}
	return
}
