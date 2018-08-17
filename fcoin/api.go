package fcoin

type API interface {
	// getter
	GetEndPoint() string

	// interface
	Endpoint
}

func newAPI(opts ...Option) (c *Configure) {
	// https://stackoverflow.com/questions/44543374/cannot-take-the-address-of-and-cannot-call-pointer-method-on?noredirect=1&lq=1
	c = (&Configure{}).setDefault()
	for _, opt := range opts {
		opt(c)
	}
	return
}

// getter
func (c *Configure) GetEndPoint() string {
	return c.Endpoint
}
