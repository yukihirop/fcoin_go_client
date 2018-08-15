package endpoint

import "fmt"

func AccountsBalance(c *EndpointConfigure) (ret string, err error) {
	url := GetPath("accounts", "AccountsBalance")
	ret, err = c.Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
