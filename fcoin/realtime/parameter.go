package fcoin_realtime

import "fmt"

type Params struct {
	Symbol     string
	Level      string
	Resolution string
	Limit      string
}

type ParamsOption func(*Params)
type PramsOptions []ParamsOption

func Symbol(s string) ParamsOption {
	return func(p *Params) {
		p.Symbol = s
	}
}

func Level(l string) ParamsOption {
	return func(p *Params) {
		p.Level = l
	}
}

func Resolution(r string) ParamsOption {
	return func(p *Params) {
		p.Resolution = r
	}
}

func Limit(l int) ParamsOption {
	return func(p *Params) {
		p.Limit = fmt.Sprint(l)
	}
}

func setParams(opts PramsOptions) (p *Params) {
	p = &Params{}
	for _, opt := range opts {
		opt(p)
	}
	return
}
