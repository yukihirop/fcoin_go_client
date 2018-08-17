package fcoin

type Client interface {
	API
}

func NewClient(opts ...Option) (c *Configure) {
	// https://stackoverflow.com/questions/44543374/cannot-take-the-address-of-and-cannot-call-pointer-method-on?noredirect=1&lq=1
	return newAPI(opts...)
}
