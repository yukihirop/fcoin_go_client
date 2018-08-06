package mock

func (m *Mock) PublicServerTime(cassetName string) (ret string, err error) {
	m.cassetName = cassetName
	url := m.url("public", "PublicServerTime")
	ret, err = m.Get(url)
	return
}

func (m *Mock) PublicCurrencies(cassetName string) (ret string, err error) {
	m.cassetName = cassetName
	url := m.url("public", "PublicCurrencies")
	ret, err = m.Get(url)
	return
}

func (m *Mock) PublicSymbols(cassetName string) (ret string, err error) {
	m.cassetName = cassetName
	url := m.url("public", "PublicSymbols")
	ret, err = m.Get(url)
	return
}
