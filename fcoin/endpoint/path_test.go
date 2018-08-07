package endpoint_test

import (
	"fcoin_go_client/fcoin/endpoint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Path", func() {
	Describe("GetPath", func() {
		Context("when path exist", func() {
			//public
			publicServerTime := endpoint.GetPath("public", "PublicServerTime")
			publicCurrencies := endpoint.GetPath("public", "PublicCurrencies")
			publicSymbols := endpoint.GetPath("public", "PublicSymbols")
			//market
			marketTicker := endpoint.GetPath("market", "MarketTicker")
			marketDepth := endpoint.GetPath("market", "MarketDepth")
			marketTrades := endpoint.GetPath("market", "MarketTrades")
			marketCandles := endpoint.GetPath("market", "MarketCandles")

			It("should return public/server-time", func() {
				//public
				Expect(publicServerTime).To(Equal("public/server-time"))
				Expect(publicCurrencies).To(Equal("public/currencies"))
				Expect(publicSymbols).To(Equal("public/symbols"))
				//market
				Expect(marketTicker).To(Equal("market/ticker"))
				Expect(marketDepth).To(Equal("market/depth"))
				Expect(marketTrades).To(Equal("market/trades"))
				Expect(marketCandles).To(Equal("market/candles"))
			})
		})

		Context("when path do not exist", func() {
			subject := endpoint.GetPath("do_not_exist", "")
			It("should be blank", func() {
				Expect(subject).To(Equal(""))
			})
		})
	})
})
