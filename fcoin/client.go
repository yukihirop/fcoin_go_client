package fcoin

import (
	"fcoin_go_client/fcoin/realtime"
)

type Client interface {
	API
}

func NewClient(opts ...Option) (c *Configure) {
	// https://stackoverflow.com/questions/44543374/cannot-take-the-address-of-and-cannot-call-pointer-method-on?noredirect=1&lq=1
	return newAPI(opts...)
}

func NewRealTimeClient() (c *fcoin_realtime.Configure) {
	return fcoin_realtime.NewRealTimeClient()
}
