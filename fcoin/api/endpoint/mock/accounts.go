package mock

import "fmt"

func (m *Mock) AccountsBalance() (ret string, err error) {
	url := "https://api.fcoin.com/v2/accounts/balance"
	ret, err = m.Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
