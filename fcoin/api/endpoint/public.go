package endpoint

import (
	"fmt"
)

//http://horie1024.hatenablog.com/entry/2014/08/25/012123
func PublicServerTime(c *EndpointConfigure) (ret string, err error) {
	url := GetPath("public", "PublicServerTime")
	ret, err = c.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicCurrencies(c *EndpointConfigure) (ret string, err error) {
	url := GetPath("public", "PublicCurrencies")
	ret, err = c.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicSymbols(c *EndpointConfigure) (ret string, err error) {
	url := GetPath("public", "PublicSymbols")
	ret, err = c.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
