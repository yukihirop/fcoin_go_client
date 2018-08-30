package validator

type MarketParams struct {
	params *Params
}

type MarketValidator interface {
	IsValid() bool
	Messages() []error
}

func NewMarketValidator(opts ...ParamsOption) (ret *MarketParams) {
	pa := setParams(opts)
	ret = &MarketParams{}
	ret.params = pa
	return
}

func (mp *MarketParams) IsValid() (ret bool) {
	pa := mp.params
	if !mp.isValidSymbol() {
		ret = false
		return
	}
	switch pa.MethodName {
	case "MarketDepth":
		ret = mp.isValidLevel()
	case "MarketCandles":
		ret = mp.isValidResolution()
	}
	return
}

func (mp *MarketParams) Messages() (ret []error) {
	pa := mp.params
	ret = []error{}
	if mp.IsValid() {
		return
	}

	if !mp.isValidSymbol() {
		ret = append(ret, presenceErrorMessage(pa.Symbol, "symbol"))
	}

	switch pa.MethodName {
	case "MarketDepth":

		if !mp.isValidLevel() {
			ret = append(ret, includesErrorMessage(pa.Level, "level", mp.validLevels()))
		}
	case "MarketCandles":
		if !mp.isValidResolution() {
			ret = append(ret, includesErrorMessage(pa.Resolution, "resolution", mp.validResolutions()))
		}
	}
	return
}

func (mp *MarketParams) isValidSymbol() (ret bool) {
	pa := mp.params
	ret = false
	if pa.Symbol != "" {
		ret = true
	}
	return
}

func (mp *MarketParams) isValidLevel() (ret bool) {
	pa := mp.params
	ret = contains(mp.validLevels(), pa.Level)
	return
}

func (mp *MarketParams) isValidResolution() (ret bool) {
	pa := mp.params
	ret = contains(mp.validResolutions(), pa.Resolution)
	return
}

func (mp *MarketParams) validLevels() (ret []string) {
	pa := mp.params
	levels := pa.VSetting.FixedViper.Get("fcoin.validation.params.level").([]interface{})
	for _, level := range levels {
		alevel, _ := level.(string)
		ret = append(ret, alevel)
	}
	return
}

func (mp *MarketParams) validResolutions() (ret []string) {
	pa := mp.params
	resolutions := pa.VSetting.FixedViper.Get("fcoin.validation.params.resolution").([]interface{})
	for _, resolution := range resolutions {
		aresolution, _ := resolution.(string)
		ret = append(ret, aresolution)
	}
	return
}
