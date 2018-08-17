package endpoint

import (
	"fmt"
)

//http://horie1024.hatenablog.com/entry/2014/08/25/012123
func PublicServerTime(c *EndpointConfigure) (ret string, err error) {
	url := c.getUrl("public", "PublicServerTime")
	fmt.Println(url)
	ret, err = apiConfig(c).Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicCurrencies(c *EndpointConfigure) (ret string, err error) {
	url := c.getUrl("public", "PublicCurrencies")
	ret, err = apiConfig(c).Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicSymbols(c *EndpointConfigure) (ret string, err error) {
	url := c.getUrl("public", "PublicSymbols")
	ret, err = apiConfig(c).Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
