package fcoin_test

import (
	"fcoin_go_client/fcoin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Api", func() {
	Describe("NewAPI", func() {
		Context("when args given", func() {
			api := fcoin.NewAPI(
				fcoin.Adapter("testAdapter"),
				fcoin.EndPoint("testEndPoint"),
				fcoin.UserAgent("testUserAgent"),
				fcoin.Proxy("testProxy"),
				fcoin.ApiKey("testApiKey"),
				fcoin.SecretKey("testSecretKey"),
			)
			It("should be overrided", func() {
				Expect(api.Adapter()).To(Equal("testAdapter"))
				Expect(api.EndPoint()).To(Equal("testEndPoint"))
				Expect(api.UserAgent()).To(Equal("testUserAgent"))
				Expect(api.Proxy()).To(Equal("testProxy"))
				Expect(api.ApiKey()).To(Equal("testApiKey"))
				Expect(api.SecretKey()).To(Equal("testSecretKey"))
			})
		})

		Context("when args is not given", func() {
			api := fcoin.NewAPI()
			It("should be default", func() {
				Expect(api.Adapter()).To(Equal("adapter"))
				Expect(api.EndPoint()).To(Equal("https://api.fcoin.com/v2/"))
				Expect(api.UserAgent()).To(Equal("Fcoin Go Client"))
				Expect(api.Proxy()).To(Equal(""))
				Expect(api.ApiKey()).To(Equal("fcoin public api key"))
				Expect(api.SecretKey()).To(Equal("fcoin secret api key"))
			})
		})
	})
	Describe("PublicServerTime", func() {
		api := fcoin.NewAPI()
		subject := api.PublicServerTime()
		expected := "PublicServerTime"

		It("should be expected", func() {
			Expect(subject).To(Equal(expected))
		})
	})

	Describe("PublicCurrencies", func() {
		api := fcoin.NewAPI()
		subject := api.PublicCurrencies()
		expected := "PublicCurrencies"

		It("should be expected", func() {
			Expect(subject).To(Equal(expected))

		})
	})

	Describe("PublicSymbols", func() {
		api := fcoin.NewAPI()
		subject := api.PublicSymbols()
		expected := "PublicSymbols"

		It("should be expected", func() {
			Expect(subject).To(Equal(expected))
		})
	})
})
