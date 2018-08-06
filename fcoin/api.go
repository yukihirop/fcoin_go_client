package fcoin

import (
	"fcoin_go_client/fcoin/endpoint"
)

type API interface {
	// getter
	GetAdapter() string
	GetEndPoint() string
	GetUserAgent() string
	GetProxy() string
	GetApiKey() string
	GetSecretKey() string

	// public
	PublicServerTime() (string, error)
	PublicCurrencies() (string, error)
	PublicSymbols() (string, error)
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
func (c *Configure) GetAdapter() string {
	return c.Adapter
}

func (c *Configure) GetEndPoint() string {
	return c.Endpoint
}

func (c *Configure) GetUserAgent() string {
	return c.UserAgent
}

func (c *Configure) GetProxy() string {
	return c.Proxy
}

func (c *Configure) GetApiKey() string {
	return c.ApiKey
}

func (c *Configure) GetSecretKey() string {
	return c.SecretKey
}

func endpointConfig(c *Configure) *(endpoint.Configure) {
	var config endpoint.Configure
	config.Adapter = c.Adapter
	config.Endpoint = c.Endpoint
	config.UserAgent = c.UserAgent
	config.Proxy = c.Proxy
	config.ApiKey = c.ApiKey
	config.SecretKey = c.SecretKey
	return &config
}

// //http://horie1024.hatenablog.com/entry/2014/08/25/012123
// http://text.baldanders.info/golang/interface/
func (c *Configure) PublicServerTime() (ret string, err error) {
	ret, err = endpoint.PublicServerTime(endpointConfig(c))
	return
}

func (c *Configure) PublicCurrencies() (ret string, err error) {
	ret, err = endpoint.PublicCurrencies(endpointConfig(c))
	return
}

func (c *Configure) PublicSymbols() (ret string, err error) {
	ret, err = endpoint.PublicSymbols(endpointConfig(c))
	return
}
