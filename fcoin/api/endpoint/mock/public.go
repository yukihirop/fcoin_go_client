package mock

import "fmt"

func (m *Mock) PublicServerTime() (ret string, err error) {
	url := "https://api.fcoin.com/v2/public/server-time"
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) PublicCurrencies() (ret string, err error) {
	url := "https://api.fcoin.com/v2/public/currencies"
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func (m *Mock) PublicSymbols() (ret string, err error) {
	url := "https://api.fcoin.com/v2/public/symbols"
	ret, err = m.Get(url, nil, nil, false)
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
