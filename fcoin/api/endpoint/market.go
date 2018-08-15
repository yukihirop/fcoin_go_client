package endpoint

import (
	"fmt"
)

//http://horie1024.hatenablog.com/entry/2014/08/25/012123
func MarketTicker(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	ma := setEndpoint(opts)
	url := GetPath(c.Endpoint, "MarketTicker") + "/" + ma.Symbol
	ret, err = c.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketDepth(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	ma := setEndpoint(opts)
	url := GetPath(c.Endpoint, "MarketDepth") + "/" + ma.Symbol + "/" + ma.Level
	ret, err = c.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketTrades(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	ma := setEndpoint(opts)
	url := GetPath(c.Endpoint, "MarketTrades") + "/" + ma.Symbol
	ret, err = c.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketCandles(c *EndpointConfigure, opts ...EndpointOption) (ret string, err error) {
	ma := setEndpoint(opts)
	url := GetPath(c.Endpoint, "MarketCandles") + "/" + ma.Resolution + "/" + ma.Symbol
	ret, err = c.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
