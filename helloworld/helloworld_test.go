package helloworld

import (
    . "github.com/onsi/ginkgo"
    . "github.com/onsi/gomega"
)

var _ = Describe("CreateMessage", func() {
    Context("when no prefix", func() {
        prefix := ""

        subject := CreateMessage(prefix)
        expected := "NO PREFIX: " + HELLOWORLD
        It("should be expected", func() {
            Expect(subject).To(Equal(expected))
        })
    })

    Context("when prefix given", func() {
        prefix := "test: "

        subject := CreateMessage(prefix)
        expected := prefix + HELLOWORLD
        It("should be expected", func() {
            Expect(subject).To(Equal(expected))
        })
    })
})