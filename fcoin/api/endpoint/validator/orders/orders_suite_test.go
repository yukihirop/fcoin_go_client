package orders_test

import (
	"fcoin_go_client/fcoin/config"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/spf13/viper"
)

var fixedViper = viper.New()
var customViper = viper.New()
var VSetting = new(config.ValidationSetting)

func SetValidationSetting() {
	VSetting.FixedViper = fixedViper
	VSetting.CustomViper = customViper
	VSetting.CustomSettingPath = nil
	VSetting.SetValidationSetting()
}

func TestOrders(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Orders Suite")
}
