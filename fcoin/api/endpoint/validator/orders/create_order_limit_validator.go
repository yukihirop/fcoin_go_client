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

func (colp *CreateOrderLimitParams) IsValid() (ret bool) {
	pa := colp.params
	if pa.isValidSymbolSettingExist() {
		ret = pa.isValidPrice() && pa.isValidAmount()
	} else {
		ret = pa.isValidSymbol() && pa.isValidSide()
	}
	return
}

func (colp *CreateOrderLimitParams) Messages() (ret map[string]string) {
	pa := colp.params
	if colp.IsValid() {
		ret = map[string]string{}
	}

	var results []map[string]string
	switch {
	case !pa.isValidSymbol():
		results = append(results, presenceErrorMessage(pa.Symbol, "symbol"))
	case !pa.isValidSide():
		results = append(results, includesErrorMessage(pa.Side, "side", pa.validSides()))
	case pa.isValidSymbolSettingExist():
		if !pa.isValidPrice() {
			results = append(results, betweenErrorMessage(pa.Price, "price", pa.min("price"), pa.max("price")))
		}
		if !pa.isValidAmount() {
			results = append(results, betweenErrorMessage(pa.Amount, "amount", pa.min("amount"), pa.max("amount")))
		}
	}
	ret = slice2map(results)
	return
}
