package endpoint

import (
	"fcoin_go_client/fcoin/api/endpoint/validator"
)

type Validator interface {
	IsValid() bool
	Messages() map[string]string
}

type ValidatorParams struct {
	params *Params
}

func NewValidator(opts ...ParamsOption) (ret *ValidatorParams) {
	pa := setParams(opts)
	ret = new(ValidatorParams)
	ret.params = pa
	return
}

func (pa *ValidatorParams) IsValid() (ret bool) {
	switch {
	case pa.isMarket():
		ret = pa.marketValidator().IsValid()
	case pa.isOrders():
		ret = pa.ordersValidator().IsValid()
	}
	return
}

func (pa *ValidatorParams) Messages() (ret map[string]string) {
	if pa.IsValid() {
		ret = map[string]string{}
	}
	switch {
	case pa.isMarket():
		ret = pa.marketValidator().Messages()
	case pa.isOrders():
		ret = pa.ordersValidator().Messages()
	}
	return
}

func (pa *ValidatorParams) isMarket() (ret bool) {
	ret = contains([]string{"MarketDepth", "MarketCandles"}, pa.params.MethodName)
	return
}

func (pa *ValidatorParams) isOrders() (ret bool) {
	ret = contains([]string{"CreateOrderLimit", "CreateOrderMarket", "OrderList"}, pa.params.MethodName)
	return
}

func (pa *ValidatorParams) marketValidator() (ret *validator.MarketParams) {
	symbol := validator.Symbol(pa.params.Symbol)
	level := validator.Level(pa.params.Level)
	resolution := validator.Level(pa.params.Resolution)
	methodName := validator.MethodName(pa.params.MethodName)
	fixedViper := pa.params.VSetting.FixedViper
	customViper := pa.params.VSetting.CustomViper
	customSettingPath := pa.params.VSetting.CustomSettingPath
	vsetting := validator.VSetting(fixedViper, customViper, customSettingPath)
	ret = validator.NewMarketValidator(symbol, level, resolution, methodName, vsetting)
	return
}

func (pa *ValidatorParams) ordersValidator() (ret *validator.OrderParams) {
	symbol := validator.Symbol(pa.params.Symbol)
	side := validator.Side(pa.params.Side)
	price := validator.Price(pa.params.Price)
	amount := validator.Amount(pa.params.Amount)
	total := validator.Total(pa.params.Total)
	states := validator.States(pa.params.States)
	methodName := validator.MethodName(pa.params.MethodName)
	fixedViper := pa.params.VSetting.FixedViper
	customViper := pa.params.VSetting.CustomViper
	customSettingPath := pa.params.VSetting.CustomSettingPath
	vsetting := validator.VSetting(fixedViper, customViper, customSettingPath)
	ret = validator.NewOrdersValidator(symbol, side, price, amount, total, states, methodName, vsetting)
	return
}

func contains(s []string, e string) (ret bool) {
	ret = false
	for _, v := range s {
		if e == v {
			ret = true
		}
	}
	return
}
