package fcoin_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFcoin(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Fcoin Suite")
}
