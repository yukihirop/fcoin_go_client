package mock

import "fmt"

type MockEndpointConfigure struct {
	//  The adapter that will be used to connect if none is set
	Adapter   string
	Endpoint  string
	UserAgent string
	Proxy     string
	ApiKey    string
	SecretKey string
}

type MockEndpoint struct {
	Symbol     string
	Level      string
	Resolution string
	Side       string
	Price      string
	Amount     string
	Type       string
	States     string
	PageBefore string
	PageAfter  string
	PerPage    string
	OrderId    string
}

type MockEndpointOption func(*MockEndpoint)
type MockEndpointsOptions []MockEndpointOption

func Symbol(s string) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.Symbol = s
	}
}

func Level(l string) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.Level = l
	}
}

func Resolution(r string) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.Resolution = r
	}
}

func Side(s string) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.Side = s
	}
}

func Price(p float32) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.Price = fmt.Sprint(p)
	}
}

func Amount(a float32) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.Amount = fmt.Sprint(a)
	}
}

func Type(t string) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.Type = t
	}
}

func States(s string) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.States = s
	}
}

func PageBefore(pb int) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.PageBefore = fmt.Sprint(pb)
	}
}

func PageAfter(pa int) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.PageAfter = fmt.Sprint(pa)
	}
}

func PerPage(pp int) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.PerPage = fmt.Sprint(pp)
	}
}

func OrderId(o string) MockEndpointOption {
	return func(e *MockEndpoint) {
		e.OrderId = fmt.Sprint(o)
	}
}

func setMockEndpoint(opts MockEndpointsOptions) (mo *MockEndpoint) {
	mo = &MockEndpoint{}
	for _, opt := range opts {
		opt(mo)
	}
	return
}
