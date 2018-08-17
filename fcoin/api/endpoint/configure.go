package endpoint

import (
	"fcoin_go_client/fcoin/api"
	"fmt"
)

type EndpointConfigure struct {
	*api.APIConfigure
	//  The adapter that will be used to connect if none is set
	Adapter   string
	Endpoint  string
	UserAgent string
	Proxy     string
	ApiKey    string
	SecretKey string
}

type Endpoint struct {
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

func States(s string) EndpointOption {
	return func(e *Endpoint) {
		e.States = s
	}
}

func PageBefore(pb int) EndpointOption {
	return func(e *Endpoint) {
		e.PageBefore = fmt.Sprint(pb)
	}
}

func PageAfter(pa int) EndpointOption {
	return func(e *Endpoint) {
		e.PageAfter = fmt.Sprint(pa)
	}
}

func PerPage(pp int) EndpointOption {
	return func(e *Endpoint) {
		e.PerPage = fmt.Sprint(pp)
	}
}

func OrderId(o string) EndpointOption {
	return func(e *Endpoint) {
		e.OrderId = o
	}
}

func setEndpoint(opts EndpointsOptions) (ma *Endpoint) {
	ma = &Endpoint{}
	for _, opt := range opts {
		opt(ma)
	}
	return
}

func apiConfig(c *EndpointConfigure) *(api.APIConfigure) {
	var config api.APIConfigure
	config.Adapter = c.Adapter
	config.Endpoint = c.Endpoint
	config.UserAgent = c.UserAgent
	config.Proxy = c.Proxy
	config.ApiKey = c.ApiKey
	config.SecretKey = c.SecretKey
	return &config
}
