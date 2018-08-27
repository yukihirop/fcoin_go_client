package orders

type OrderListValidator interface {
	NewOrderListValidator(...ParamsOption) *OrderListParams
	IsValid() bool
	Messages() map[string]string
}

type OrderListParams struct {
	params *Params
}

func NewOrderListValidator(opts ...ParamsOption) (ret *OrderListParams) {
	pa := setParams(opts)
	ret = &OrderListParams{}
	ret.params = pa
	return
}

func (pa *OrderListParams) IsValid() (ret bool) {
	ret = pa.params.isValidSymbol() && pa.params.isValidStates()
	return
}

func (pa *OrderListParams) Messages() (ret map[string]string) {
	if pa.IsValid() {
		ret = map[string]string{}
	}
	var results []map[string]string
	switch {
	case !pa.params.isValidSymbol():
		results = append(results, presenceErrorMessage(pa.params.Symbol, "symbol"))
	case !pa.params.isValidStates():
		results = append(results, includesErrorMessage(pa.params.States, "states", pa.params.validStates()))
	}
	ret = slice2map(results)
	return
}
