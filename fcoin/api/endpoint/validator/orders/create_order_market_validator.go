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

func (comp *CreateOrderMarketParams) IsValid() (ret bool) {
	pa := comp.params
	switch {
	case pa.isValidSymbolSettingExist():
		switch {
		case pa.isSell():
			ret = pa.isValidAmount()
		case pa.isBuy():
			ret = pa.isValidTotal()
		}
	case comp.isInvalidSymbolSettingExist():
		ret = false
	default:
		ret = pa.isValidSymbol() && pa.isValidSide()
	}
	return
}

func (comp *CreateOrderMarketParams) Messages() (ret map[string]string) {
	pa := comp.params
	if comp.IsValid() {
		ret = map[string]string{}
	}

	var results []map[string]string
	switch {
	case !pa.isValidSymbol():
		results = append(results, presenceErrorMessage(pa.Symbol, "symbol"))
	case pa.isValidSymbolSettingExist():
		switch {
		case pa.isSell():
			if !pa.isValidAmount() {
				results = append(results, betweenErrorMessage(pa.Amount, "amount", pa.min("amount"), pa.max("amount")))
			}
		case pa.isBuy():
			if !pa.isValidTotal() {
				results = append(results, betweenErrorMessage(pa.Total, "total", pa.min("total"), pa.max("total")))
			}
		}
	case comp.isInvalidSymbolSettingExist():
		results = append(results, invalidCreateOrderMarketErrorMessage(pa.Symbol, "symbol"))
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

func (comp *CreateOrderMarketParams) isInvalidSymbolSettingExist() (ret bool) {
	pa := comp.params
	ret = false
	symbols := pa.invalidSymbols().symbols
	for _, symbol := range symbols {
		if symbol == pa.Symbol {
			ret = true
		}
	}
	return
}
