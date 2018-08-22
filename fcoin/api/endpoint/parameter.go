package endpoint

import "fmt"

type Params struct {
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

type ParamsOption func(*Params)
type ParamsOptions []ParamsOption

func Symbol(s string) ParamsOption {
	return func(pa *Params) {
		pa.Symbol = s
	}
}

func Level(l string) ParamsOption {
	return func(pa *Params) {
		pa.Level = l
	}
}

func Resolution(r string) ParamsOption {
	return func(pa *Params) {
		pa.Resolution = r
	}
}

func Side(s string) ParamsOption {
	return func(pa *Params) {
		pa.Side = s
	}
}

func Price(p float32) ParamsOption {
	return func(pa *Params) {
		pa.Price = fmt.Sprint(p)
	}
}

func Amount(a float32) ParamsOption {
	return func(pa *Params) {
		pa.Amount = fmt.Sprint(a)
	}
}

func Type(t string) ParamsOption {
	return func(pa *Params) {
		pa.Type = t
	}
}

func States(s string) ParamsOption {
	return func(pa *Params) {
		pa.States = s
	}
}

func PageBefore(pb int) ParamsOption {
	return func(pa *Params) {
		pa.PageBefore = fmt.Sprint(pb)
	}
}

func PageAfter(paa int) ParamsOption {
	return func(pa *Params) {
		pa.PageAfter = fmt.Sprint(paa)
	}
}

func PerPage(pp int) ParamsOption {
	return func(pa *Params) {
		pa.PerPage = fmt.Sprint(pp)
	}
}

func OrderId(o string) ParamsOption {
	return func(pa *Params) {
		pa.OrderId = o
	}
}

func setParams(opts ParamsOptions) (pa *Params) {
	pa = &Params{}
	for _, opt := range opts {
		opt(pa)
	}
	return
}
