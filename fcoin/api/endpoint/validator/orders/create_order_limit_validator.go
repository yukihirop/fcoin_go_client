package orders

type CreateOrderLimitValidator interface {
	IsValid() bool
	Messages() map[string]string
}

type CreateOrderLimitParams struct {
	params *Params
}

func NewCreateOrderLimitValidator(opts ...ParamsOption) (ret *CreateOrderLimitParams) {
	pa := setParams(opts)
	pa.Type = "limit"
	ret = &CreateOrderLimitParams{}
	ret.params = pa
	ret.params.ValidSymbols = pa.validSymbols()
	return
}

func (pa *CreateOrderLimitParams) IsValid() (ret bool) {
	if pa.params.isValidSymbolSettingExist() {
		ret = pa.params.isValidPrice() && pa.params.isValidAmount()
	} else {
		ret = pa.params.isValidSymbol() && pa.params.isValidSide()
	}
	return
}

func (pa *CreateOrderLimitParams) Messages() (ret map[string]string) {
	if pa.IsValid() {
		ret = map[string]string{}
	}

	var results []map[string]string
	switch {
	case !pa.params.isValidSymbol():
		results = append(results, presenceErrorMessage(pa.params.Symbol, "symbol"))
	case !pa.params.isValidSide():
		results = append(results, includesErrorMessage(pa.params.Side, "side", pa.params.validSides()))
	case pa.params.isValidSymbolSettingExist():
		if !pa.params.isValidPrice() {
			results = append(results, betweenErrorMessage(pa.params.Price, "price", pa.params.min("price"), pa.params.max("price")))
		}
		if !pa.params.isValidAmount() {
			results = append(results, betweenErrorMessage(pa.params.Amount, "amount", pa.params.min("amount"), pa.params.max("amount")))
		}
	}
	ret = slice2map(results)
	return
}
