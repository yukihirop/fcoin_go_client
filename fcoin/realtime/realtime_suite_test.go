package fcoin_realtime_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRealtime(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Realtime Suite")
}
