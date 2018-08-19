package fcoin_realtime

import (
	uuid "github.com/satori/go.uuid"
)

func onTicker(c *Configure, handler HandleFunc, params ...ParamsOption) {
	pa := setParams(params)
	topic := "ticker" + "." + pa.Symbol
	c.on(topic, nil, handler)
}

func onDepth(c *Configure, handler HandleFunc, params ...ParamsOption) {
	pa := setParams(params)
	topic := "depth" + "." + pa.Level + "." + pa.Symbol
	c.on(topic, nil, handler)
}

func onTrade(c *Configure, handler HandleFunc, params ...ParamsOption) {
	pa := setParams(params)
	topic := "trade" + "." + pa.Symbol
	c.on(topic, pa.Limit, handler)
}

func onCandle(c *Configure, handler HandleFunc, params ...ParamsOption) {
	pa := setParams(params)
	topic := "candle" + "." + pa.Resolution + "." + pa.Symbol
	c.on(topic, pa.Limit, handler)
}

func onTopics(c *Configure, handler HandleFunc) {
	c.on("topics", nil, handler)
}

func onHello(c *Configure, handler HandleFunc) {
	c.on("hello", nil, handler)
}

func (c *Configure) on(topic string, limit interface{}, handler HandleFunc) {
	client := c.createWebSocketClient(topic, handler)
	client.write(validSendMessage(topic, limit))
	go client.read()
	for {
		client.callCallback(topic)
	}
}

func validSendMessage(topic string, limit interface{}) (ret map[string]interface{}) {
	u2 := uuid.NewV4()
	al, ok := limit.(string)
	if ok {
		ret = map[string]interface{}{"cmd": "sub", "args": []interface{}{topic, al}, "id": u2.String()}
	} else {
		ret = map[string]interface{}{"cmd": "sub", "args": []interface{}{topic}, "id": u2.String()}
	}
	return
}
