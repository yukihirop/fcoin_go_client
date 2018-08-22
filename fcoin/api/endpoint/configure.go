package endpoint

import (
	"fcoin_go_client/fcoin/api"
)

type Configure struct {
	*api.Configure
	//  The adapter that will be used to connect if none is set
	Adapter   string
	Endpoint  string
	UserAgent string
	Proxy     string
	ApiKey    string
	SecretKey string
}

func apiConfig(c *Configure) *(api.Configure) {
	var config api.Configure
	config.Adapter = c.Adapter
	config.Endpoint = c.Endpoint
	config.UserAgent = c.UserAgent
	config.Proxy = c.Proxy
	config.ApiKey = c.ApiKey
	config.SecretKey = c.SecretKey
	return &config
}
