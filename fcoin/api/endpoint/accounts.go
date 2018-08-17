package endpoint

import "fmt"

func AccountsBalance(c *EndpointConfigure) (ret string, err error) {
	url := c.getUrl("accounts", "AccountsBalance")
	ret, err = apiConfig(c).Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
