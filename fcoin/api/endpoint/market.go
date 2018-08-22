package endpoint

import (
	"fmt"
)

//http://horie1024.hatenablog.com/entry/2014/08/25/012123
func MarketTicker(c *EndpointConfigure, opts ...ParamsOption) (ret string, err error) {
	ma := setParams(opts)
	url := c.getUrl("market", "MarketTicker") + "/" + ma.Symbol
	ret, err = apiConfig(c).Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketDepth(c *EndpointConfigure, opts ...ParamsOption) (ret string, err error) {
	ma := setParams(opts)
	url := c.getUrl("market", "MarketDepth") + "/" + ma.Symbol + "/" + ma.Level
	ret, err = apiConfig(c).Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketTrades(c *EndpointConfigure, opts ...ParamsOption) (ret string, err error) {
	ma := setParams(opts)
	url := c.getUrl("market", "MarketTrades") + "/" + ma.Symbol
	ret, err = apiConfig(c).Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func MarketCandles(c *EndpointConfigure, opts ...ParamsOption) (ret string, err error) {
	ma := setParams(opts)
	url := c.getUrl("market", "MarketCandles") + "/" + ma.Resolution + "/" + ma.Symbol
	ret, err = apiConfig(c).Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
