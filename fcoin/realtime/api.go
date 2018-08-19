package fcoin_realtime

type RealTimeAPI interface {
	OnTicker(string, HandleFunc)
	OnDepth(string, string, HandleFunc)
	OnTrade(string, interface{}, HandleFunc)
	OnCandle(string, string, interface{}, HandleFunc)
	OnTopics(HandleFunc)
	OnHello(HandleFunc)
}

func newRealTimeAPI() (c *Configure) {
	// https://stackoverflow.com/questions/44543374/cannot-take-the-address-of-and-cannot-call-pointer-method-on?noredirect=1&lq=1
	c = (&Configure{}).setDefault()
	return
}

func (c *Configure) OnTicker(symbol string, handler HandleFunc) {
	onTicker(c, handler, Symbol(symbol))
}

func (c *Configure) OnDepth(symbol, level string, handler HandleFunc) {
	onDepth(c, handler, Symbol(symbol), Level(level))
}

func (c *Configure) OnTrade(symbol string, limit interface{}, handler HandleFunc) {
	switch al := limit.(type) {
	case int:
		onTrade(c, handler, Symbol(symbol), Limit(al))
	case nil:
		onTrade(c, handler, Symbol(symbol))
	}
}

func (c *Configure) OnCandle(symbol, resolution string, limit interface{}, handler HandleFunc) {
	switch al := limit.(type) {
	case int:
		onCandle(c, handler, Symbol(symbol), Resolution(resolution), Limit(al))
	case nil:
		onCandle(c, handler, Symbol(symbol), Resolution(resolution))
	}
}

func (c *Configure) OnTopics(handler HandleFunc) {
	onTopics(c, handler)
}

func (c *Configure) OnHello(handler HandleFunc) {
	onHello(c, handler)
}
