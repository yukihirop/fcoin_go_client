package mock

import "fmt"

func (m *Mock) AccountsBalance() (ret string, err error) {
	url := m.url("accounts", "AccountsBalance")
	ret, err = m.Get(url, nil, nil, true)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
