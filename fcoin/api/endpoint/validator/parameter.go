package validator

import (
	"fcoin_go_client/fcoin/api/endpoint/validator/orders"
	"fcoin_go_client/fcoin/config"

	"github.com/spf13/viper"
)

type Params struct {
	Symbol     string
	Level      string
	Resolution string
	Side       string
	Price      float64
	Amount     float64
	Total      float64
	Type       string
	States     string
	MethodName string

	VSetting *config.ValidationSetting
}

type ParamsOption func(*Params)
type ParamsOptions []ParamsOption

func Symbol(s string) ParamsOption {
	return func(pa *Params) {
		pa.Symbol = s
	}
}

func Level(l string) ParamsOption {
	return func(pa *Params) {
		pa.Level = l
	}
}

func Resolution(r string) ParamsOption {
	return func(pa *Params) {
		pa.Resolution = r
	}
}

func Side(s string) ParamsOption {
	return func(pa *Params) {
		pa.Side = s
	}
}

func Price(p float64) ParamsOption {
	return func(pa *Params) {
		pa.Price = p
	}
}

func Amount(a float64) ParamsOption {
	return func(pa *Params) {
		pa.Amount = a
	}
}

func Total(t float64) ParamsOption {
	return func(pa *Params) {
		pa.Total = t
	}
}

func Type(t string) ParamsOption {
	return func(pa *Params) {
		pa.Type = t
	}
}

func States(s string) ParamsOption {
	return func(pa *Params) {
		pa.States = s
	}
}

func MethodName(m string) ParamsOption {
	return func(pa *Params) {
		pa.MethodName = m
	}
}

func VSetting(fixedViper, customViper *viper.Viper, customSettingPath interface{}) ParamsOption {
	return func(pa *Params) {
		pa.VSetting = new(config.ValidationSetting)
		pa.VSetting.FixedViper = fixedViper
		pa.VSetting.CustomViper = customViper
		pa.VSetting.CustomSettingPath = customSettingPath
	}
}

func ordersParams(pa *Params) *orders.Params {
	var params orders.Params
	params.Symbol = pa.Symbol
	params.Level = pa.Level
	params.Resolution = pa.Resolution
	params.Side = pa.Side
	params.Amount = pa.Amount
	params.Total = pa.Total
	params.Type = pa.Type
	params.States = pa.States
	params.VSetting = pa.VSetting
	return &params
}

func setParams(opts ParamsOptions) (pa *Params) {
	pa = &Params{}
	for _, opt := range opts {
		opt(pa)
	}
	return
}
