package endpoint

import "fmt"

//http://horie1024.hatenablog.com/entry/2014/08/25/012123
func PublicServerTime(c *Configure) (ret string, err error) {
	url := GetPath(c.Endpoint, "PublicServerTime")
	ret, err = Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicCurrencies(c *Configure) (ret string, err error) {
	url := GetPath(c.Endpoint, "PublicCurrencies")
	ret, err = Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicSymbols(c *Configure) (ret string, err error) {
	url := GetPath(c.Endpoint, "PublicSymbols")
	ret, err = Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
