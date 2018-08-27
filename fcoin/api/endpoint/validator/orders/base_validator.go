package orders

import (
	"strconv"
)

func (pa *Params) isValidSymbol() (ret bool) {
	ret = false
	if pa.Symbol != "" {
		ret = true
	}
	return
}

func (pa *Params) isValidSide() (ret bool) {
	ret = contains(pa.validSides(), pa.Side)
	return
}

func (pa *Params) isValidType() (ret bool) {
	ret = contains(pa.validTypes(), pa.Type)
	return
}

func (pa *Params) isValidStates() (ret bool) {
	ret = contains(pa.validStates(), pa.States)
	return
}

func (pa *Params) isValidPrice() (ret bool) {
	ret = false
	price := pa.Price
	if pa.min("price") <= price && price <= pa.max("price") {
		ret = true
	}
	return
}

func (pa *Params) isValidTotal() (ret bool) {
	ret = false
	total := pa.Total
	if pa.min("total") <= total && total <= pa.max("total") {
		ret = true
	}
	return
}

func (pa *Params) isValidAmount() (ret bool) {
	ret = false
	amount := pa.Amount
	if pa.min("amount") <= amount && amount <= pa.max("amount") {
		ret = true
	}
	return
}

func (pa *Params) min(typeValue string) (ret float64) {
	if pa.isValidSymbolSettingExist() {
		var min string
		for k, v := range pa.ValidSymbols {
			if k != pa.Symbol {
				continue
			}
			if typeValue == "amount" {
				min = v.amount["min"]
			} else if typeValue == "price" {
				min = v.price["min"]
			} else if typeValue == "total" {
				min = v.total["min"]
			}
		}
		ret, _ = strconv.ParseFloat(min, 64)
	} else {
		ret = 0
	}
	return
}

func (pa *Params) max(typeValue string) (ret float64) {
	if pa.isValidSymbolSettingExist() {
		var max string
		for k, v := range pa.ValidSymbols {
			if k != pa.Symbol {
				continue
			}
			if typeValue == "amount" {
				max = v.amount["max"]
			} else if typeValue == "price" {
				max = v.price["max"]
			} else if typeValue == "total" {
				max = v.total["max"]
			}
		}
		ret, _ = strconv.ParseFloat(max, 64)
	} else {
		ret = 0
	}
	return
}

func (pa *Params) isValidSymbolSettingExist() (ret bool) {
	ret = false
	if pa.ValidSymbols[pa.Symbol] != nil {
		ret = true
	}
	return
}

func (pa *Params) validSides() (ret []string) {
	sides, _ := pa.VSetting.FixedViper.Get("fcoin.validation.params.side").([]interface{})
	for _, value := range sides {
		side, _ := value.(string)
		ret = append(ret, side)
	}
	return
}

func (pa *Params) validTypes() (ret []string) {
	types, _ := pa.VSetting.FixedViper.Get("fcoin.validation.params.type").([]interface{})
	for _, value := range types {
		typeValue, _ := value.(string)
		ret = append(ret, typeValue)
	}
	return
}

func (pa *Params) validStates() (ret []string) {
	states, _ := pa.VSetting.FixedViper.Get("fcoin.validation.params.states").([]interface{})
	for _, value := range states {
		state, _ := value.(string)
		ret = append(ret, state)
	}
	return
}
