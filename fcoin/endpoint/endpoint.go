package endpoint

type Configure struct {
	//  The adapter that will be used to connect if none is set
	Adapter   string
	Endpoint  string
	UserAgent string
	Proxy     string
	ApiKey    string
	SecretKey string
}
