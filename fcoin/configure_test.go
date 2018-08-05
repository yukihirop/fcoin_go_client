package fcoin

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configure", func() {
	Describe("setConfig", func() {
		c := &Configure{}
		expected := &Configure{
			adapter:   "testAdapter",
			endpoint:  "testEndpoint",
			userAgent: "testUserAgent",
			proxy:     "testProxy",
			apiKey:    "testApiKey",
			secretKey: "testSecretKey",
		}
		subject := c.setConfig(expected)
		It("should be expected", func() {
			Expect(subject).To(Equal(expected))
		})
	})

	Describe("setDefault", func() {
		c := &Configure{}
		subject := c.setDefault()
		expected := &Configure{
			adapter:   "adapter",
			endpoint:  "https://api.fcoin.com/v2/",
			userAgent: "Fcoin Go Client",
			proxy:     "",
			apiKey:    "fcoin public api key",
			secretKey: "fcoin secret api key",
		}

		It("should be expected", func() {
			Expect(subject).To(Equal(expected))
		})
	})
})
