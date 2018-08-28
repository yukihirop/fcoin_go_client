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

func (olp *OrderListParams) IsValid() (ret bool) {
	pa := olp.params
	ret = pa.isValidSymbol() && pa.isValidStates()
	return
}

func (olp *OrderListParams) Messages() (ret map[string]string) {
	pa := olp.params
	if olp.IsValid() {
		ret = map[string]string{}
	}
	var results []map[string]string
	switch {
	case !pa.isValidSymbol():
		results = append(results, presenceErrorMessage(pa.Symbol, "symbol"))
	case !pa.isValidStates():
		results = append(results, includesErrorMessage(pa.States, "states", pa.validStates()))
	}
	ret = slice2map(results)
	return
}
