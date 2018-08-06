package fcoin

const DEFAULT_ADAPTER = "adapter"
const DEFAULT_ENDPOINT = "https://api.fcoin.com/v2/"
const DEFAULT_USER_AGENT = "Fcoin Go Client"
const DEFAULT_PROXY = ""
const DEFAULT_API_KEY = "fcoin public api key"
const DEFAULT_SERECT_KEY = "fcoin secret api key"

type Configure struct {
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
type Option func(*Configure)

func Adapter(a string) Option {
	return func(c *Configure) {
		c.Adapter = a
	}
}

func EndPoint(e string) Option {
	return func(c *Configure) {
		c.Endpoint = e
	}
}

func UserAgent(u string) Option {
	return func(c *Configure) {
		c.UserAgent = u
	}
}

func Proxy(p string) Option {
	return func(c *Configure) {
		c.Proxy = p
	}
}

func ApiKey(a string) Option {
	return func(c *Configure) {
		c.ApiKey = a
	}
}

func SecretKey(s string) Option {
	return func(c *Configure) {
		c.SecretKey = s
	}
}

func (c *Configure) setConfig(config *Configure) *Configure {
	return config
}

func (c *Configure) setDefault() *Configure {
	c = &Configure{
		Adapter:   DEFAULT_ADAPTER,
		Endpoint:  DEFAULT_ENDPOINT,
		UserAgent: DEFAULT_USER_AGENT,
		Proxy:     DEFAULT_PROXY,
		ApiKey:    DEFAULT_API_KEY,
		SecretKey: DEFAULT_SERECT_KEY,
	}
	return c
}
