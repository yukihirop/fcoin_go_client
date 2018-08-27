package orders

import "fmt"

type CreateOrderMarketValidator interface {
	IsValid() bool
	Message() map[string]string
}

type CreateOrderMarketParams struct {
	params *Params
}

func NewCreateOrderMarketValidator(opts ...ParamsOption) (ret *CreateOrderMarketParams) {
	pa := setParams(opts)
	pa.Type = "market"
	ret = &CreateOrderMarketParams{}
	ret.params = pa
	ret.params.ValidSymbols = pa.validSymbols()
	ret.params.InvalidSymbols = pa.invalidSymbols()
	return
}

func (pa *CreateOrderMarketParams) IsValid() (ret bool) {
	switch {
	case pa.params.isValidSymbolSettingExist():
		switch {
		case pa.params.isSell():
			ret = pa.params.isValidAmount()
		case pa.params.isBuy():
			ret = pa.params.isValidTotal()
		}
	case pa.isInvalidSymbolSettingExist():
		ret = false
	default:
		ret = pa.params.isValidSymbol() && pa.params.isValidSide()
	}
	return
}

func (pa *CreateOrderMarketParams) Messages() (ret map[string]string) {
	if pa.IsValid() {
		ret = map[string]string{}
	}

	var results []map[string]string
	switch {
	case !pa.params.isValidSymbol():
		results = append(results, presenceErrorMessage(pa.params.Symbol, "symbol"))
	case pa.params.isValidSymbolSettingExist():
		switch {
		case pa.params.isSell():
			if !pa.params.isValidAmount() {
				results = append(results, betweenErrorMessage(pa.params.Amount, "amount", pa.params.min("amount"), pa.params.max("amount")))
			}
		case pa.params.isBuy():
			if !pa.params.isValidTotal() {
				results = append(results, betweenErrorMessage(pa.params.Total, "total", pa.params.min("total"), pa.params.max("total")))
			}
		}
	case pa.isInvalidSymbolSettingExist():
		results = append(results, invalidCreateOrderMarketErrorMessage(pa.params.Symbol, "symbol"))
	}
	ret = slice2map(results)
	return
}

func invalidCreateOrderMarketErrorMessage(target interface{}, targetType string) (ret map[string]string) {
	var targetValue string
	if target != nil {
		targetValue = fmt.Sprint(target)
	} else {
		targetValue = "nil"
	}
	errorMessage := fmt.Sprintf("%s is %s. %s board is not adapted on-going order.", targetType, targetValue, targetType)
	ret = map[string]string{targetType: errorMessage}
	return
}

func (pa *CreateOrderMarketParams) isInvalidSymbolSettingExist() (ret bool) {
	ret = false
	symbols := pa.params.invalidSymbols().symbols
	for _, symbol := range symbols {
		if symbol == pa.params.Symbol {
			ret = true
		}
	}
	return
}
