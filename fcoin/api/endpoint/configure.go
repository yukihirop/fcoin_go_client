package endpoint

import (
	"fcoin_go_client/fcoin/api"
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
