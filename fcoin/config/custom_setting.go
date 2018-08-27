package config

import (
	"bytes"
	"fmt"
)

func (s *ValidationSetting) SetDefaultValidationCustomSetting() (err error) {
	s.CustomViper.SetConfigType("yaml")
	var customSetting = []byte(`
fcoin:
  validation:
    limit:
      sell:
        mainboard_A:
          - { symbol: 'btcusdt', amount: { min: 1,     max: 10000 }, price: { min: 1, max: 10000 } }
          - { symbol: 'ethusdt', amount: { min: 0.001, max: 10000 }, price: { min: 1, max: 10000 } }
          - { symbol: 'bchusdt', amount: { min: 0.001, max: 5000 },  price: { min: 1, max: 10000 } }
          - { symbol: 'ltcusdt', amount: { min: 0.001, max: 40000 }, price: { min: 1, max: 10000 } }
          - { symbol: 'etcusdt', amount: { min: 0.001, max: 400000 },price: { min: 1, max: 10000 } }
          - { symbol: 'xrpusdt', amount: { min: 1,     max: 10000 }, price: { min: 1, max: 10000 } }
        mainboard_B:
          - { symbol: 'ftusdt', amount: { min: 1, max: 10000 },  price: { min: 1,         max: 10000 } }
          - { symbol: 'ftbtc',  amount: { min: 1, max: 10000 },  price: { min: 0.0000001, max: 10000 } }
          - { symbol: 'fteth',  amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'zipeth', amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'omgeth', amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'btmusdt',amount: { min: 1, max: 10000 },  price: { min: 1,         max: 10000 } }
          - { symbol: 'zrxeth', amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'bnbusdt',amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'zipusdt',amount: { min: 1, max: 10000 },  price: { min: 1,         max: 10000 } }
          - { symbol: 'fieth',  amount: { min: 1, max: 10000 },  price: { min: 0.0000011, max: 10000 } }
          - { symbol: 'fiusdt', amount: { min: 1, max: 10000 },  price: { min: 1,         max: 10000 } }
        gpm:
          - { symbol: 'fcandyusdt', amount: { min: 0.001, max: 10000 }, price: { min: 1, max: 10000 } }
          - { symbol: 'icxeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'zileth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'aeeth',      amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: '777eth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'guseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'cccxeth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'bancaeth',   amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'praeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'dcceth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ssseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'mdteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'tsteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'pmdeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'rteeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'xpseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'tcteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'dwseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ngoteth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ateth',      amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'soceth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'blzeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ocneth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'datxeth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'gtceth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'leteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'dageth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'yeeeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'aaaeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'nceth',      amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'arpeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'grameth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ifoodeth',   amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'hpceth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'sgcceth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: '3dbeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'xmxeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'rcteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'citeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'eeseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'faireth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'brmeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'sdaeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'cbreth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
      buy:
        mainboard_A:
          - { symbol: 'btcusdt', amount: { min: 1,     max: 10000 }, price: { min: 1, max: 10000 } }
          - { symbol: 'ethusdt', amount: { min: 0.001, max: 10000 }, price: { min: 1, max: 10000 } }
          - { symbol: 'bchusdt', amount: { min: 0.001, max: 5000 },  price: { min: 1, max: 10000 } }
          - { symbol: 'ltcusdt', amount: { min: 0.001, max: 40000 }, price: { min: 1, max: 10000 } }
          - { symbol: 'etcusdt', amount: { min: 0.001, max: 400000 },price: { min: 1, max: 10000 } }
          - { symbol: 'xrpusdt', amount: { min: 1,     max: 10000 }, price: { min: 1, max: 10000 } }
        mainboard_B:
          - { symbol: 'ftusdt', amount: { min: 1, max: 10000 },  price: { min: 1,         max: 10000 } }
          - { symbol: 'ftbtc',  amount: { min: 1, max: 10000 },  price: { min: 0.0000001, max: 10000 } }
          - { symbol: 'fteth',  amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'zipeth', amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'omgeth', amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'btmusdt',amount: { min: 1, max: 10000 },  price: { min: 1,         max: 10000 } }
          - { symbol: 'zrxeth', amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'bnbusdt',amount: { min: 1, max: 10000 },  price: { min: 0.000001,  max: 10000 } }
          - { symbol: 'zipusdt',amount: { min: 1, max: 10000 },  price: { min: 1,         max: 10000 } }
          - { symbol: 'fieth',  amount: { min: 1, max: 10000 },  price: { min: 0.0000011, max: 10000 } }
          - { symbol: 'fiusdt', amount: { min: 1, max: 10000 },  price: { min: 1,         max: 10000 } }
        gpm:
          - { symbol: 'fcandyusdt', amount: { min: 0.001, max: 10000 }, price: { min: 1, max: 10000 } }
          - { symbol: 'icxeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'zileth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'aeeth',      amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: '777eth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'guseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'cccxeth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'bancaeth',   amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'praeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'dcceth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ssseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'mdteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'tsteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'pmdeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'rteeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'xpseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'tcteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'dwseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ngoteth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ateth',      amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'soceth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'blzeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ocneth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'datxeth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'gtceth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'leteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'dageth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'yeeeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'aaaeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'nceth',      amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'arpeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'grameth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'ifoodeth',   amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'hpceth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'sgcceth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: '3dbeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'xmxeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'rcteth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'citeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'eeseth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'faireth',    amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'brmeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'sdaeth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
          - { symbol: 'cbreth',     amount: { min: 0.01, max: 10000 },  price: { min: 0.000001, max: 10000 } }
    market:
      valid:
        sell:
          mainboard_A:
            - { symbol: 'btcusdt', amount: { min: 1,     max: 10000 } }
            - { symbol: 'ethusdt', amount: { min: 0.001, max: 10000 } }
            - { symbol: 'bchusdt', amount: { min: 0.001, max: 5000 } }
            - { symbol: 'ltcusdt', amount: { min: 0.001, max: 40000 } }
            - { symbol: 'etcusdt', amount: { min: 0.001, max: 400000 } }
            - { symbol: 'xrpusdt', amount: { min: 1,     max: 10000 } }
          mainboard_B:
          gpm:
        buy:
          mainboard_A:
            - { symbol: 'btcusdt', total: { min: 1,     max: 10000 } }
            - { symbol: 'ethusdt', total: { min: 0.001, max: 10000 } }
            - { symbol: 'bchusdt', total: { min: 0.001, max: 5000 } }
            - { symbol: 'ltcusdt', total: { min: 0.001, max: 40000 } }
            - { symbol: 'etcusdt', total: { min: 0.001, max: 400000 } }
            - { symbol: 'xrpusdt', total: { min: 1,     max: 10000 } }
          mainboard_B:
          gpm:
      invalid:
        sell:
          mainboard_A:
          mainboard_B:
            - fiusdt
            - fieth
          gpm:
            - fcandyusdt
        buy:
          mainboard_A:
          mainboard_B:
            - fiusdt
            - fieth
          gpm:
            - fcandyusdt
`)

	switch ap := s.CustomSettingPath.(type) {
	case string:
		s.CustomViper.AddConfigPath(ap)
		err = s.CustomViper.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
			return
		}
	case nil:
		err = s.CustomViper.ReadConfig(bytes.NewBuffer(customSetting))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	return
}
