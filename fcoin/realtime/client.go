package fcoin_realtime

type RealTimeClient interface {
	NewRealTimeClient() *Configure
}

func NewRealTimeClient() (c *Configure) {
	return newRealTimeAPI()
}
