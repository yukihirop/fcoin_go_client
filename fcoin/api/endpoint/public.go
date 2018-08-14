package endpoint

import (
	"fmt"
)

//http://horie1024.hatenablog.com/entry/2014/08/25/012123
func PublicServerTime(c *EndpointConfigure) (ret string, err error) {
	url := GetPath(c.Endpoint, "PublicServerTime")
	ret, err = c.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicCurrencies(c *EndpointConfigure) (ret string, err error) {
	url := GetPath(c.Endpoint, "PublicCurrencies")
	ret, err = c.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicSymbols(c *EndpointConfigure) (ret string, err error) {
	url := GetPath(c.Endpoint, "PublicSymbols")
	ret, err = c.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
