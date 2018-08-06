package mock

func PublicServerTime(cassetName string) (ret string, err error) {
	m := &Mock{cassetName: cassetName}
	ret, err = m.Get("https://api.fcoin.com/v2/public/server-time")
	return
}

func PublicCurrencies(cassetName string) (ret string, err error) {
	m := &Mock{cassetName: cassetName}
	ret, err = m.Get("https://api.fcoin.com/v2/public/currencies")
	return
}

func PublicSymbols(cassetName string) (ret string, err error) {
	m := &Mock{cassetName: cassetName}
	ret, err = m.Get("https://api.fcoin.com/v2/public/symbols")
	return
}
