package endpoint_test

import (
	. "fcoin_go_client/fcoin/api/endpoint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Path", func() {
	Describe("GetPath", func() {
		Context("when path exist", func() {
			//public
			publicServerTime := GetPath("public", "PublicServerTime")
			publicCurrencies := GetPath("public", "PublicCurrencies")
			publicSymbols := GetPath("public", "PublicSymbols")
			//market
			marketTicker := GetPath("market", "MarketTicker")
			marketDepth := GetPath("market", "MarketDepth")
			marketTrades := GetPath("market", "MarketTrades")
			marketCandles := GetPath("market", "MarketCandles")

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
			subject := GetPath("do_not_exist", "")
			It("should be blank", func() {
				Expect(subject).To(Equal(""))
			})
		})
	})
})
