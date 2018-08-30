package orders

type OrderListValidator interface {
	IsValid() bool
	Messages() []error
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

func (olp *OrderListParams) Messages() (ret []error) {
	pa := olp.params
	ret = []error{}
	if olp.IsValid() {
		return
	}
	switch {
	case !pa.isValidSymbol():
		ret = append(ret, presenceErrorMessage(pa.Symbol, "symbol"))
	case !pa.isValidStates():
		ret = append(ret, includesErrorMessage(pa.States, "states", pa.validStates()))
	}
	return
}
