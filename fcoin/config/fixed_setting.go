package config

import (
	"bytes"
	"fmt"
)

func (s *ValidationSetting) SetDefaultValidationFixedSetting() (err error) {
	s.FixedViper.SetConfigType("yaml")
	var fixedSetting = []byte(`
fcoin:
  validation:
    params:
      side: [buy, sell]
      type: [limit, market]
      resolution: [M1, M3, M5, M15, M30, H1, H4, H6, D1, W1, MN]
      level: [L20, L100, full]
      states: [submitted, partial_filled, canceled, partial_canceled, filled, pending_cancel]`)
  
	err = s.FixedViper.ReadConfig(bytes.NewBuffer(fixedSetting))
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}
