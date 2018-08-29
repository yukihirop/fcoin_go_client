package orders

type CreateOrderLimitValidator interface {
	IsValid() bool
	Messages() []error
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

func (colp *CreateOrderLimitParams) Messages() (ret []error) {
	pa := colp.params
	ret = []error{}
	if colp.IsValid() {
		return
	}
	switch {
	case !pa.isValidSymbol():
		ret = append(ret, presenceErrorMessage(pa.Symbol, "symbol"))
	case !pa.isValidSide():
		ret = append(ret, includesErrorMessage(pa.Side, "side", pa.validSides()))
	case pa.isValidSymbolSettingExist():
		if !pa.isValidPrice() {
			ret = append(ret, betweenErrorMessage(pa.Price, "price", pa.min("price"), pa.max("price")))
		}
		if !pa.isValidAmount() {
			ret = append(ret, betweenErrorMessage(pa.Amount, "amount", pa.min("amount"), pa.max("amount")))
		}
	}
	return
}
