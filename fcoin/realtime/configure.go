package fcoin_realtime

const (
	DEFAULT_WSS_ENDPOINT = "wss://api.fcoin.com/v2/ws"
)

type Configure struct {
	WssEndpoint string
}

type HandleFunc func(data map[string]interface{})

func (c *Configure) setDefault() *Configure {
	c = &Configure{
		WssEndpoint: DEFAULT_WSS_ENDPOINT,
	}
	return c
}
