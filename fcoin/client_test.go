package fcoin_test

import (
	"fcoin_go_client/fcoin"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Client", func() {
	Describe("NewClient", func() {
		Context("when args given", func() {
			api := fcoin.NewClient(
				fcoin.EndPoint("testEndPoint"),
			)
			It("should be overrided", func() {
				Expect(api.GetEndPoint()).To(Equal("testEndPoint"))
			})
		})

		Context("when args is not given", func() {
			api := fcoin.NewClient()
			It("should be default", func() {
				Expect(api.GetEndPoint()).To(Equal("https://api.fcoin.com/v2/"))
			})
		})
	})
})
