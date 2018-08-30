package config

import (
	"github.com/spf13/viper"
)

type ValidationSetting struct {
	FixedViper        *viper.Viper
	CustomViper       *viper.Viper
	CustomSettingPath interface{}
}

func (s *ValidationSetting) SetValidationSetting() {
	s.SetDefaultValidationFixedSetting()
	s.SetDefaultValidationCustomSetting()
	return
}
