package orders

import "fmt"

type ValidSymbolSetting struct {
	symbol string
	amount map[string]string
	price  map[string]string
	total  map[string]string
}

func (pa *Params) validSymbols() (ret map[string]*ValidSymbolSetting) {
	switch {
	case pa.isLimit():
		ret = pa.validSymbolsWhenLimit()
	case pa.isMarket():
		if pa.isSell() {
			ret = pa.validSymbolsWhenrMarketSell()
		} else if pa.isBuy() {
			ret = pa.validSymbolsWhenrMarketBuy()
		}
	}
	return
}

func (pa *Params) validSymbolsWhenLimit() (ret map[string]*ValidSymbolSetting) {
	mainboardA, _ := pa.VSetting.CustomViper.Get("fcoin.validation.limit." + pa.Side + ".mainboard_A").([]interface{})
	mainboardB, _ := pa.VSetting.CustomViper.Get("fcoin.validation.limit." + pa.Side + ".mainboard_B").([]interface{})
	gpm, _ := pa.VSetting.CustomViper.Get("fcoin.validation.limit." + pa.Side + ".gpm").([]interface{})
	symbolSettings := append(mainboardA, mainboardB...)
	symbolSettings = append(symbolSettings, gpm...)

	ret = make(map[string]*ValidSymbolSetting)
	for _, setting := range symbolSettings {
		// e.x.) map[symbol:btcusdt amount:map[min:1 max:10000] price:map[min:1 max:10000]]
		symbol := fmt.Sprint(setting.(map[interface{}]interface{})["symbol"])
		amount := setting.(map[interface{}]interface{})["amount"]
		price := setting.(map[interface{}]interface{})["price"]
		amountMin := fmt.Sprint(amount.(map[interface{}]interface{})["min"])
		amountMax := fmt.Sprint(amount.(map[interface{}]interface{})["max"])
		priceMin := fmt.Sprint(price.(map[interface{}]interface{})["min"])
		priceMax := fmt.Sprint(price.(map[interface{}]interface{})["max"])

		ret[symbol] = &ValidSymbolSetting{
			symbol: symbol,
			amount: map[string]string{
				"min": amountMin,
				"max": amountMax,
			},
			price: map[string]string{
				"min": priceMin,
				"max": priceMax,
			},
		}
	}
	return
}

func (pa *Params) validSymbolsWhenrMarketSell() (ret map[string]*ValidSymbolSetting) {
	mainboardA, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.valid.sell.mainboard_A").([]interface{})
	mainboardB, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.valid.sell.mainboard_B").([]interface{})
	gpm, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.valid.sell.gpm").([]interface{})
	symbolSettings := append(mainboardA, mainboardB...)
	symbolSettings = append(symbolSettings, gpm...)

	ret = make(map[string]*ValidSymbolSetting)
	for _, setting := range symbolSettings {
		symbol := fmt.Sprint(setting.(map[interface{}]interface{})["symbol"])
		amount := setting.(map[interface{}]interface{})["amount"]
		amountMin := fmt.Sprint(amount.(map[interface{}]interface{})["min"])
		amountMax := fmt.Sprint(amount.(map[interface{}]interface{})["max"])

		ret[symbol] = &ValidSymbolSetting{
			symbol: symbol,
			amount: map[string]string{
				"min": amountMin,
				"max": amountMax,
			},
		}
	}
	return
}

func (pa *Params) validSymbolsWhenrMarketBuy() (ret map[string]*ValidSymbolSetting) {
	mainboardA, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.valid.buy.mainboard_A").([]interface{})
	mainboardB, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.valid.buy.mainboard_B").([]interface{})
	gpm, _ := pa.VSetting.CustomViper.Get("fcoin.validation.market.valid.sell.gpm").([]interface{})
	symbolSettings := append(mainboardA, mainboardB...)
	symbolSettings = append(symbolSettings, gpm...)

	ret = make(map[string]*ValidSymbolSetting)
	for _, setting := range symbolSettings {
		symbol := fmt.Sprint(setting.(map[interface{}]interface{})["symbol"])
		total := setting.(map[interface{}]interface{})["total"]
		totalMin := fmt.Sprint(total.(map[interface{}]interface{})["min"])
		totalMax := fmt.Sprint(total.(map[interface{}]interface{})["max"])

		ret[symbol] = &ValidSymbolSetting{
			symbol: symbol,
			total: map[string]string{
				"min": totalMin,
				"max": totalMax,
			},
		}
	}
	return
}
