package validator

import (
	"fcoin_go_client/fcoin/api/endpoint/validator/orders"
	"fcoin_go_client/fcoin/config"
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

func (op *OrderParams) IsValid() (ret bool) {
	pa := op.params
	ordersParams := ordersParams(pa)
	switch pa.MethodName {
	case "CreateOrderLimit":
		ret = createOrderLimitValidator(ordersParams).IsValid()
	case "CreateOrderMarket":
		ret = createOrderMarketValidator(ordersParams).IsValid()
	case "OrderList":
		ret = orderListValidator(ordersParams).IsValid()
	}
	return
}

func (op *OrderParams) Messages() (ret map[string]string) {
	pa := op.params
	ordersParams := ordersParams(pa)
	if op.IsValid() {
		ret = map[string]string{}
	}
	switch pa.MethodName {
	case "CreateOrderLimit":
		ret = createOrderLimitValidator(ordersParams).Messages()
	case "CreateOrderMarket":
		ret = createOrderMarketValidator(ordersParams).Messages()
	case "OrderList":
		ret = orderListValidator(ordersParams).Messages()
	}
	return

}

func createOrderLimitValidator(op *orders.Params) (ret *orders.CreateOrderLimitParams) {
	symbol := orders.Symbol(op.Symbol)
	side := orders.Side(op.Side)
	price := orders.Price(op.Price)
	amount := orders.Amount(op.Amount)
	vsetting := vsetting(op)
	ret = orders.NewCreateOrderLimitValidator(symbol, side, price, amount, vsetting)
	return
}

func createOrderMarketValidator(op *orders.Params) (ret *orders.CreateOrderMarketParams) {
	symbol := orders.Symbol(op.Symbol)
	side := orders.Side(op.Side)
	price := orders.Price(op.Price)
	amount := orders.Amount(op.Amount)
	total := orders.Total(op.Total)
	vsetting := vsetting(op)
	ret = orders.NewCreateOrderMarketValidator(symbol, side, price, amount, total, vsetting)
	return
}

func orderListValidator(op *orders.Params) (ret *orders.OrderListParams) {
	symbol := orders.Symbol(op.Symbol)
	states := orders.States(op.States)
	vsetting := vsetting(op)
	ret = orders.NewOrderListValidator(symbol, states, vsetting)
	return
}

func vsetting(op *orders.Params) orders.ParamsOption {
	fixedViper := delegateVSetting(op).FixedViper
	customViper := delegateVSetting(op).CustomViper
	customSettingPath := delegateVSetting(op).CustomSettingPath
	return orders.VSetting(fixedViper, customViper, customSettingPath)
}

func delegateVSetting(op *orders.Params) *config.ValidationSetting {
	return op.VSetting
}
