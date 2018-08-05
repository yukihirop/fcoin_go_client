package fcoin

const DEFAULT_ADAPTER = "adapter"
const DEFAULT_ENDPOINT = "https://api.fcoin.com/v2/"
const DEFAULT_USER_AGENT = "Fcoin Go Client"
const DEFAULT_PROXY = ""
const DEFAULT_API_KEY = "fcoin public api key"
const DEFAULT_SERECT_KEY = "fcoin secret api key"

type Configure struct {
	//  The adapter that will be used to connect if none is set
	adapter   string
	endpoint  string
	userAgent string
	proxy     string
	apiKey    string
	secretKey string
}

//Option is function value of functional options
// http://text.baldanders.info/golang/functional-options-pattern/
type Option func(*Configure)

func Adapter(a string) Option {
	return func(c *Configure) {
		c.adapter = a
	}
}

func EndPoint(e string) Option {
	return func(c *Configure) {
		c.endpoint = e
	}
}

func UserAgent(u string) Option {
	return func(c *Configure) {
		c.userAgent = u
	}
}

func Proxy(p string) Option {
	return func(c *Configure) {
		c.proxy = p
	}
}

func ApiKey(a string) Option {
	return func(c *Configure) {
		c.apiKey = a
	}
}

func SecretKey(s string) Option {
	return func(c *Configure) {
		c.secretKey = s
	}
}

func (c *Configure) setConfig(config *Configure) *Configure {
	return config
}

func (c *Configure) setDefault() *Configure {
	c = &Configure{
		adapter:   DEFAULT_ADAPTER,
		endpoint:  DEFAULT_ENDPOINT,
		userAgent: DEFAULT_USER_AGENT,
		proxy:     DEFAULT_PROXY,
		apiKey:    DEFAULT_API_KEY,
		secretKey: DEFAULT_SERECT_KEY,
	}
	return c
}
