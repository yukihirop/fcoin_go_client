package orders

type InvalidSymbolSetting struct {
	symbols []string
}

func (pa *Params) invalidSymbols() (ret *InvalidSymbolSetting) {
	mainboardA, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.invalid." + pa.Side + ".mainboard_A").([]interface{})
	mainboardB, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.invalid." + pa.Side + ".mainboard_B").([]interface{})
	gpm, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.invalid." + pa.Side + ".gpm").([]interface{})
	symbolSettings := append(mainboardA, mainboardB...)
	symbolSettings = append(symbolSettings, gpm...)

	ret = new(InvalidSymbolSetting)
	var symbols []string
	for _, symbol := range symbolSettings {
		asymbol := symbol.(string)
		symbols = append(symbols, asymbol)
	}
	ret = &InvalidSymbolSetting{
		symbols: symbols,
	}
	return
}
