package fcoin

type API interface {
	// getter
	GetAdapter() string
	GetEndPoint() string
	GetUserAgent() string
	GetProxy() string
	GetApiKey() string
	GetSecretKey() string

	// interface
	Endpoint
}

func NewAPI(opts ...Option) (c *Configure) {
	// https://stackoverflow.com/questions/44543374/cannot-take-the-address-of-and-cannot-call-pointer-method-on?noredirect=1&lq=1
	c = (&Configure{}).setDefault()
	for _, opt := range opts {
		opt(c)
	}
	return
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
