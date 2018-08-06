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
				Expect(api.GetAdapter()).To(Equal("testAdapter"))
				Expect(api.GetEndPoint()).To(Equal("testEndPoint"))
				Expect(api.GetUserAgent()).To(Equal("testUserAgent"))
				Expect(api.GetProxy()).To(Equal("testProxy"))
				Expect(api.GetApiKey()).To(Equal("testApiKey"))
				Expect(api.GetSecretKey()).To(Equal("testSecretKey"))
			})
		})

		Context("when args is not given", func() {
			api := fcoin.NewAPI()
			It("should be default", func() {
				Expect(api.GetAdapter()).To(Equal("adapter"))
				Expect(api.GetEndPoint()).To(Equal("https://api.fcoin.com/v2/"))
				Expect(api.GetUserAgent()).To(Equal("Fcoin Go Client"))
				Expect(api.GetProxy()).To(Equal(""))
				Expect(api.GetApiKey()).To(Equal("fcoin public api key"))
				Expect(api.GetSecretKey()).To(Equal("fcoin secret api key"))
			})
		})
	})
})
