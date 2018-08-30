package validator

import (
	"fcoin_go_client/fcoin/api/endpoint/validator/orders"
	"fcoin_go_client/fcoin/config"
)

type OrderValidator interface {
	IsValid() bool
	Messages() []error
}

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
	switch pa.MethodName {
	case "CreateOrderLimit":
		ret = createOrderLimitValidator(op).IsValid()
	case "CreateOrderMarket":
		ret = createOrderMarketValidator(op).IsValid()
	case "OrderList":
		ret = orderListValidator(op).IsValid()
	}
	return
}

func (op *OrderParams) Messages() (ret []error) {
	pa := op.params
	ret = []error{}
	if op.IsValid() {
		return
	}
	switch pa.MethodName {
	case "CreateOrderLimit":
		ret = createOrderLimitValidator(op).Messages()
	case "CreateOrderMarket":
		ret = createOrderMarketValidator(op).Messages()
	case "OrderList":
		ret = orderListValidator(op).Messages()
	}
	return

}

func createOrderLimitValidator(op *OrderParams) (ret *orders.CreateOrderLimitParams) {
	ordersParams := ordersParams(op.params)
	symbol := orders.Symbol(ordersParams.Symbol)
	side := orders.Side(ordersParams.Side)
	price := orders.Price(ordersParams.Price)
	amount := orders.Amount(ordersParams.Amount)
	vsetting := vsetting(op)
	ret = orders.NewCreateOrderLimitValidator(symbol, side, price, amount, vsetting)
	return
}

func createOrderMarketValidator(op *OrderParams) (ret *orders.CreateOrderMarketParams) {
	ordersParams := ordersParams(op.params)
	symbol := orders.Symbol(ordersParams.Symbol)
	side := orders.Side(ordersParams.Side)
	price := orders.Price(ordersParams.Price)
	amount := orders.Amount(ordersParams.Amount)
	total := orders.Total(ordersParams.Total)
	vsetting := vsetting(op)
	ret = orders.NewCreateOrderMarketValidator(symbol, side, price, amount, total, vsetting)
	return
}

func orderListValidator(op *OrderParams) (ret *orders.OrderListParams) {
	ordersParams := ordersParams(op.params)
	symbol := orders.Symbol(ordersParams.Symbol)
	states := orders.States(ordersParams.States)
	vsetting := vsetting(op)
	ret = orders.NewOrderListValidator(symbol, states, vsetting)
	return
}

func vsetting(op *OrderParams) orders.ParamsOption {
	ordersParams := ordersParams(op.params)
	fixedViper := delegateVSetting(ordersParams).FixedViper
	customViper := delegateVSetting(ordersParams).CustomViper
	customSettingPath := delegateVSetting(ordersParams).CustomSettingPath
	return orders.VSetting(fixedViper, customViper, customSettingPath)
}

func delegateVSetting(op *orders.Params) *config.ValidationSetting {
	return op.VSetting
}
