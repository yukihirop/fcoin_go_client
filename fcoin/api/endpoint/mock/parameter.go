package mock

import "fmt"

type MockParams struct {
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

type MockParamsOption func(*MockParams)
type MockParamssOptions []MockParamsOption

func Symbol(s string) MockParamsOption {
	return func(e *MockParams) {
		e.Symbol = s
	}
}

func Level(l string) MockParamsOption {
	return func(e *MockParams) {
		e.Level = l
	}
}

func Resolution(r string) MockParamsOption {
	return func(e *MockParams) {
		e.Resolution = r
	}
}

func Side(s string) MockParamsOption {
	return func(e *MockParams) {
		e.Side = s
	}
}

func Price(p float32) MockParamsOption {
	return func(e *MockParams) {
		e.Price = fmt.Sprint(p)
	}
}

func Amount(a float32) MockParamsOption {
	return func(e *MockParams) {
		e.Amount = fmt.Sprint(a)
	}
}

func Type(t string) MockParamsOption {
	return func(e *MockParams) {
		e.Type = t
	}
}

func States(s string) MockParamsOption {
	return func(e *MockParams) {
		e.States = s
	}
}

func PageBefore(pb int) MockParamsOption {
	return func(e *MockParams) {
		e.PageBefore = fmt.Sprint(pb)
	}
}

func PageAfter(pa int) MockParamsOption {
	return func(e *MockParams) {
		e.PageAfter = fmt.Sprint(pa)
	}
}

func PerPage(pp int) MockParamsOption {
	return func(e *MockParams) {
		e.PerPage = fmt.Sprint(pp)
	}
}

func OrderId(o string) MockParamsOption {
	return func(e *MockParams) {
		e.OrderId = fmt.Sprint(o)
	}
}

func setMockParams(opts MockParamssOptions) (mo *MockParams) {
	mo = &MockParams{}
	for _, opt := range opts {
		opt(mo)
	}
	return
}
