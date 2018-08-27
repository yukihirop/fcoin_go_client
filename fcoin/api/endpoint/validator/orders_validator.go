package validator

import (
	"fcoin_go_client/fcoin/api/endpoint/validator/orders"
)

type OrderParams struct {
	params *Params
}

func NewOrdersValidator(opts ...ParamsOption) (ret *OrderParams) {
	pa := setParams(opts)
	ret = &OrderParams{}
	ret.params = pa
	return
}

func (pa *OrderParams) IsValid() (ret bool) {
	switch pa.params.MethodName {
	case "CreateOrderLimit":
		ret = pa.createOrderLimitValidator().IsValid()
	case "CreateOrderMarket":
		ret = pa.createOrderMarketValidator().IsValid()
	case "OrderList":
		ret = pa.orderListValidator().IsValid()
	}
	return
}

func (pa *OrderParams) Messages() (ret map[string]string) {
	if pa.IsValid() {
		ret = map[string]string{}
	}
	switch pa.params.MethodName {
	case "CreateOrderLimit":
		ret = pa.createOrderLimitValidator().Messages()
	case "CreateOrderMarket":
		ret = pa.createOrderMarketValidator().Messages()
	case "OrderList":
		ret = pa.orderListValidator().Messages()
	}
	return

}

func (pa *OrderParams) createOrderLimitValidator() (ret *orders.CreateOrderLimitParams) {
	symbol := orders.Symbol(pa.params.Symbol)
	side := orders.Side(pa.params.Side)
	price := orders.Price(pa.params.Price)
	amount := orders.Amount(pa.params.Amount)
	fixedViper := pa.params.VSetting.FixedViper
	customViper := pa.params.VSetting.CustomViper
	customSettingPath := pa.params.VSetting.CustomSettingPath
	vsetting := orders.VSetting(fixedViper, customViper, customSettingPath)
	ret = orders.NewCreateOrderLimitValidator(symbol, side, price, amount, vsetting)
	return
}

func (pa *OrderParams) createOrderMarketValidator() (ret *orders.CreateOrderMarketParams) {
	symbol := orders.Symbol(pa.params.Symbol)
	side := orders.Side(pa.params.Side)
	price := orders.Price(pa.params.Price)
	amount := orders.Amount(pa.params.Amount)
	total := orders.Total(pa.params.Total)
	fixedViper := pa.params.VSetting.FixedViper
	customViper := pa.params.VSetting.CustomViper
	customSettingPath := pa.params.VSetting.CustomSettingPath
	vsetting := orders.VSetting(fixedViper, customViper, customSettingPath)
	ret = orders.NewCreateOrderMarketValidator(symbol, side, price, amount, total, vsetting)
	return
}

func (pa *OrderParams) orderListValidator() (ret *orders.OrderListParams) {
	symbol := orders.Symbol(pa.params.Symbol)
	states := orders.States(pa.params.States)
	fixedViper := pa.params.VSetting.FixedViper
	customViper := pa.params.VSetting.CustomViper
	customSettingPath := pa.params.VSetting.CustomSettingPath
	vsetting := orders.VSetting(fixedViper, customViper, customSettingPath)
	ret = orders.NewOrderListValidator(symbol, states, vsetting)
	return
}
