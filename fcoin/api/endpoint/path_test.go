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
			//orders
			createOrderLimit := GetPath("orders", "CreateOrderLimit")
			orderList := GetPath("orders", "OrderList")
			referenceOrder := GetPath("orders", "ReferenceOrder")
			cancelOrder := GetPath("orders", "CancelOrder")
			orderMatchResults := GetPath("orders", "OrderMatchResults")
			//accounts
			accountsBalance := GetPath("accounts", "AccountsBalance")

			It("should return path", func() {
				//public
				Expect(publicServerTime).To(Equal("public/server-time"))
				Expect(publicCurrencies).To(Equal("public/currencies"))
				Expect(publicSymbols).To(Equal("public/symbols"))
				//market
				Expect(marketTicker).To(Equal("market/ticker"))
				Expect(marketDepth).To(Equal("market/depth"))
				Expect(marketTrades).To(Equal("market/trades"))
				Expect(marketCandles).To(Equal("market/candles"))
				//orders
				Expect(createOrderLimit).To(Equal("orders"))
				Expect(orderList).To(Equal("orders"))
				Expect(referenceOrder).To(Equal("orders"))
				Expect(cancelOrder).To(Equal("orders"))
				Expect(orderMatchResults).To(Equal("orders"))
				//accounts
				Expect(accountsBalance).To(Equal("accounts/balance"))
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
