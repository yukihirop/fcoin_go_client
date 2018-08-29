package orders

import (
	"errors"
	"fmt"
)

type CreateOrderMarketValidator interface {
	IsValid() bool
	Messages() []error
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

func (comp *CreateOrderMarketParams) Messages() (ret []error) {
	pa := comp.params
	ret = []error{}
	if comp.IsValid() {
		return
	}
	switch {
	case !pa.isValidSymbol():
		ret = append(ret, presenceErrorMessage(pa.Symbol, "symbol"))
	case pa.isValidSymbolSettingExist():
		switch {
		case pa.isSell():
			if !pa.isValidAmount() {
				ret = append(ret, betweenErrorMessage(pa.Amount, "amount", pa.min("amount"), pa.max("amount")))
			}
		case pa.isBuy():
			if !pa.isValidTotal() {
				ret = append(ret, betweenErrorMessage(pa.Total, "total", pa.min("total"), pa.max("total")))
			}
		}
	case comp.isInvalidSymbolSettingExist():
		ret = append(ret, invalidCreateOrderMarketErrorMessage(pa.Symbol, "symbol"))
	}
	return
}

func invalidCreateOrderMarketErrorMessage(target interface{}, targetType string) (ret error) {
	var targetValue string
	if target != nil {
		targetValue = fmt.Sprint(target)
	} else {
		targetValue = "nil"
	}
	errorMessageValue := fmt.Sprintf("%s is %s. %s board is not adapted on-going order.", targetType, targetValue, targetType)
	errorMessage := fmt.Sprintf("{%s: %s}", targetType, errorMessageValue)
	ret = errors.New(errorMessage)
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
