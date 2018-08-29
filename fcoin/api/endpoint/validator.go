package endpoint

import (
	"fcoin_go_client/fcoin/api/endpoint/validator"
	"fcoin_go_client/fcoin/config"
)

type Validator interface {
	IsValid() bool
	Messages() []error
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

func (vp *ValidatorParams) IsValid() (ret bool) {
	pa := vp.params
	switch {
	case vp.isMarket():
		ret = marketValidator(pa).IsValid()
	case vp.isOrders():
		ret = ordersValidator(pa).IsValid()
	}
	return
}

func (vp *ValidatorParams) Messages() (ret []error) {
	pa := vp.params
	ret = []error{}
	if vp.IsValid() {
		return
	}
	switch {
	case vp.isMarket():
		ret = marketValidator(pa).Messages()
	case vp.isOrders():
		ret = ordersValidator(pa).Messages()
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

func marketValidator(pa *Params) (ret *validator.MarketParams) {
	vp := validatorParams(pa)
	symbol := validator.Symbol(vp.Symbol)
	level := validator.Level(vp.Level)
	resolution := validator.Resolution(vp.Resolution)
	methodName := validator.MethodName(vp.MethodName)
	vsetting := vsetting(vp)
	ret = validator.NewMarketValidator(symbol, level, resolution, methodName, vsetting)
	return
}

func ordersValidator(pa *Params) (ret *validator.OrderParams) {
	vp := validatorParams(pa)
	symbol := validator.Symbol(vp.Symbol)
	side := validator.Side(vp.Side)
	price := validator.Price(vp.Price)
	amount := validator.Amount(vp.Amount)
	total := validator.Total(vp.Total)
	states := validator.States(vp.States)
	methodName := validator.MethodName(vp.MethodName)
	vsetting := vsetting(vp)
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

func vsetting(vp *validator.Params) validator.ParamsOption {
	fixedViper := delegateVSetting(vp).FixedViper
	customViper := delegateVSetting(vp).CustomViper
	customSettingPath := delegateVSetting(vp).CustomSettingPath
	return validator.VSetting(fixedViper, customViper, customSettingPath)
}

func delegateVSetting(vp *validator.Params) (ret *config.ValidationSetting) {
	ret = vp.VSetting
	return
}
