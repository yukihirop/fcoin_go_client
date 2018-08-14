package endpoint

import (
	"fcoin_go_client/fcoin/api"
	"fmt"
)

type EndpointConfigure struct {
	*api.APIConfigure
}

type Endpoint struct {
	Symbol     string
	Level      string
	Resolution string
	Side       string
	Price      string
	Amount     string
	Type       string
}

type EndpointOption func(*Endpoint)
type EndpointsOptions []EndpointOption

func Symbol(s string) EndpointOption {
	return func(e *Endpoint) {
		e.Symbol = s
	}
}

func Level(l string) EndpointOption {
	return func(e *Endpoint) {
		e.Level = l
	}
}

func Resolution(r string) EndpointOption {
	return func(e *Endpoint) {
		e.Resolution = r
	}
}

func Side(s string) EndpointOption {
	return func(e *Endpoint) {
		e.Side = s
	}
}

func Price(p float32) EndpointOption {
	return func(e *Endpoint) {
		e.Price = fmt.Sprint(p)
	}
}

func Amount(a float32) EndpointOption {
	return func(e *Endpoint) {
		e.Amount = fmt.Sprint(a)
	}
}

func Type(t string) EndpointOption {
	return func(e *Endpoint) {
		e.Type = t
	}
}

func setEndpoint(opts EndpointsOptions) (ma *Endpoint) {
	ma = &Endpoint{}
	for _, opt := range opts {
		opt(ma)
	}
	return
}
