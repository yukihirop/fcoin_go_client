package fcoin

import (
	"fcoin_go_client/fcoin/config"

	"github.com/spf13/viper"
)

type ValidationSetting struct {
	FixedViper        *viper.Viper
	CustomViper       *viper.Viper
	CustomSettingPath interface{}
}

func ConfigValidationSetting(vs *ValidationSetting) *(config.ValidationSetting) {
	var setting config.ValidationSetting
	setting.FixedViper = vs.FixedViper
	setting.CustomViper = vs.CustomViper
	setting.CustomSettingPath = vs.CustomSettingPath
	return &setting
}

var fixedViper = viper.New()
var customViper = viper.New()
var VSetting = new(ValidationSetting)

func main() {
	VSetting.FixedViper = fixedViper
	VSetting.CustomViper = customViper
	VSetting.CustomSettingPath = nil
	ConfigValidationSetting(VSetting).SetValidationSetting()
}
