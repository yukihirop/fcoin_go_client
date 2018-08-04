package helloworld_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHelloworld(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Helloworld Suite")
}
