package api

type APIConfigure struct {
	//  The adapter that will be used to connect if none is set
	Adapter   string
	Endpoint  string
	UserAgent string
	Proxy     string
	ApiKey    string
	SecretKey string
}

//Option is function value of functional options
// http://text.baldanders.info/golang/functional-options-pattern/
type ConfigureOption func(*APIConfigure)

func Adapter(a string) ConfigureOption {
	return func(c *APIConfigure) {
		c.Adapter = a
	}
}

func EndPoint(e string) ConfigureOption {
	return func(c *APIConfigure) {
		c.Endpoint = e
	}
}

func UserAgent(u string) ConfigureOption {
	return func(c *APIConfigure) {
		c.UserAgent = u
	}
}

func Proxy(p string) ConfigureOption {
	return func(c *APIConfigure) {
		c.Proxy = p
	}
}

func ApiKey(a string) ConfigureOption {
	return func(c *APIConfigure) {
		c.ApiKey = a
	}
}

func SecretKey(s string) ConfigureOption {
	return func(c *APIConfigure) {
		c.SecretKey = s
	}
}
