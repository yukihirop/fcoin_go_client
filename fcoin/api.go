package fcoin

import (
	"fcoin_go_client/fcoin/endpoint"
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
	PublicServerTime() int
	PublicCurrencies() int
	PublicSymbols() int
}

func NewAPI(opts ...Option) API {
	// https://stackoverflow.com/questions/44543374/cannot-take-the-address-of-and-cannot-call-pointer-method-on?noredirect=1&lq=1
	c := (&Configure{}).setDefault()
	for _, opt := range opts {
		opt(c)
	}
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
// http://text.baldanders.info/golang/interface/
func (c *Configure) PublicServerTime() (ret int) {
	ret, _ = endpoint.PublicServerTime(endpointConfig(c))
	return
}

func (c *Configure) PublicCurrencies() (ret int) {
	ret, _ = endpoint.PublicCurrencies(endpointConfig(c))
	return
}

func (c *Configure) PublicSymbols() (ret int) {
	ret, _ = endpoint.PublicSymbols(endpointConfig(c))
	return
}

func endpointConfig(c *Configure) *(endpoint.Configure) {
	var config endpoint.Configure
	config.Adapter = c.adapter
	config.Endpoint = c.endpoint
	config.UserAgent = c.userAgent
	config.Proxy = c.proxy
	config.ApiKey = c.apiKey
	config.SecretKey = c.secretKey
	return &config
}
