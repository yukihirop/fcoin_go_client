package endpoint

import (
	"bytes"

	"github.com/spf13/viper"
)

func (c *EndpointConfigure) getUrl(endPoint string, methodName string) (ret string) {
	path := getPath(endPoint, methodName)
	ret = c.Endpoint + path
	return
}

func getPath(endPoint string, methodName string) (ret string) {
	viper.SetConfigType("yaml")
	var rooting = []byte(`
    fcoin:
      endpoints:
        public:
          PublicServerTime: "public/server-time"
          PublicCurrencies: "public/currencies"
          PublicSymbols: "public/symbols"
        market:
          MarketTicker: "market/ticker"
          MarketDepth: "market/depth"
          MarketTrades: "market/trades"
          MarketCandles: "market/candles"
        orders:
          CreateOrderLimit: "orders"
          OrderList: "orders"
          ReferenceOrder: "orders"
          CancelOrder: "orders"
          OrderMatchResults: "orders"
        accounts:
          AccountsBalance: "accounts/balance"
    `)
	viper.ReadConfig(bytes.NewBuffer(rooting))
	path := "fcoin.endpoints" + "." + endPoint + "." + methodName
	ret, _ = viper.Get(path).(string)
	return
}
