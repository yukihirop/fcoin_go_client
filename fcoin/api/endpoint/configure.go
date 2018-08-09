package endpoint

import (
	"fcoin_go_client/fcoin/api"
	"strconv"
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

func Price(p int) EndpointOption {
	return func(e *Endpoint) {
		e.Price = strconv.Itoa(p)
	}
}

func setEndpoint(opts EndpointsOptions) (ma *Endpoint) {
	ma = &Endpoint{}
	for _, opt := range opts {
		opt(ma)
	}
	return
}
