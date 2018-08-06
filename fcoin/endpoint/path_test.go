package endpoint_test

import (
	"fcoin_go_client/fcoin/endpoint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Path", func() {
	Describe("GetPath", func() {
		Context("when path exist", func() {
			subject := endpoint.GetPath("public", "PublicServerTime")
			It("should return public/server-time", func() {
				Expect(subject).To(Equal("public/server-time"))
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
