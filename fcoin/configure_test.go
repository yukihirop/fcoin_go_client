package fcoin

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Configure", func() {
	Describe("setConfig", func() {
		c := &Configure{}
		expected := &Configure{
			Adapter:   "testAdapter",
			Endpoint:  "testEndpoint",
			UserAgent: "testUserAgent",
			Proxy:     "testProxy",
			ApiKey:    "testApiKey",
			SecretKey: "testSecretKey",
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
			Adapter:   "adapter",
			Endpoint:  "https://api.fcoin.com/v2/",
			UserAgent: "Fcoin Go Client",
			Proxy:     "",
			ApiKey:    "fcoin public api key",
			SecretKey: "fcoin secret api key",
		}

		It("should be expected", func() {
			Expect(subject).To(Equal(expected))
		})
	})
})
