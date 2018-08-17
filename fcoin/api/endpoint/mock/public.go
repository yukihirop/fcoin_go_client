package mock

import "fmt"

func (m *Mock) PublicServerTime() (ret string, err error) {
	url := m.url("public", "PublicServerTime")
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) PublicCurrencies() (ret string, err error) {
	url := m.url("public", "PublicCurrencies")
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) PublicSymbols() (ret string, err error) {
	url := m.url("public", "PublicSymbols")
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
