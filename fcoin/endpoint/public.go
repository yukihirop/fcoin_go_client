package endpoint

import "fmt"

type Configure struct {
	//  The adapter that will be used to connect if none is set
	Adapter   string
	Endpoint  string
	UserAgent string
	Proxy     string
	ApiKey    string
	SecretKey string
}

//http://horie1024.hatenablog.com/entry/2014/08/25/012123
func PublicServerTime(c *Configure) (ret int, err error) {
	url := Public(c) + "server-time"
	ret, err = Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicCurrencies(c *Configure) (ret int, err error) {
	url := Public(c) + "currencies"
	ret, err = Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func PublicSymbols(c *Configure) (ret int, err error) {
	url := Public(c) + "symbols"
	ret, err = Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func Public(c *Configure) string {
	return c.Endpoint + "public/"
}
