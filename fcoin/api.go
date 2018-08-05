package fcoin

import (
	"fcoin_go_client/fcoin/endpoint"
	"fmt"
)

type API interface {
	// getter
	Adapter() string
	EndPoint() string
	UserAgent() string
	Proxy() string
	ApiKey() string
	SecretKey() string

	// public
	PublicServerTime() string
	PublicCurrencies() string
	PublicSymbols() string
}

func NewAPI(opts ...Option) API {
	c := (&Configure{}).setDefault()
	for _, opt := range opts {
		opt(c)
	}
	fmt.Println(c)
	return c
}

// getter
func (c *Configure) Adapter() string {
	return c.adapter
}

func (c *Configure) EndPoint() string {
	return c.endpoint
}

func (c *Configure) UserAgent() string {
	return c.userAgent
}

func (c *Configure) Proxy() string {
	return c.proxy
}

func (c *Configure) ApiKey() string {
	return c.apiKey
}

func (c *Configure) SecretKey() string {
	return c.secretKey
}

// //http://horie1024.hatenablog.com/entry/2014/08/25/012123
func (c *Configure) PublicServerTime() string {
	return endpoint.PublicServerTime()
}

func (c *Configure) PublicCurrencies() string {
	return endpoint.PublicCurrencies()
}

func (c *Configure) PublicSymbols() string {
	return endpoint.PublicSymbols()
}
