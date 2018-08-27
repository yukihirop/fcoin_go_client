package validator

type MarketParams struct {
	params *Params
}

func NewMarketValidator(opts ...ParamsOption) (ret *MarketParams) {
	pa := setParams(opts)
	ret = &MarketParams{}
	ret.params = pa
	return
}

func (pa *MarketParams) IsValid() (ret bool) {
	if !pa.isValidSymbol() {
		ret = false
		return
	}
	switch pa.params.MethodName {
	case "MarketDepth":
		ret = pa.isValidLevel()
	case "MarketCandles":
		ret = pa.isValidResolution()
	}
	return
}

func (pa *MarketParams) Messages() (ret map[string]string) {
	if pa.IsValid() {
		ret = map[string]string{}
	}
	var results []map[string]string

	if !pa.isValidSymbol() {
		results = append(results, presenceErrorMessage(pa.params.Symbol, "symbol"))
	}

	switch pa.params.MethodName {
	case "MarketDepth":

		if !pa.isValidLevel() {
			results = append(results, includesErrorMessage(pa.params.Level, "level", pa.validLevels()))
		}
	case "MarketCandles":
		if !pa.isValidResolution() {
			results = append(results, includesErrorMessage(pa.params.Resolution, "resolution", pa.validResolutions()))
		}
	}
	ret = slice2map(results)
	return
}

func (pa *MarketParams) isValidSymbol() (ret bool) {
	ret = false
	if pa.params.Symbol != "" {
		ret = true
	}
	return
}

func (pa *MarketParams) isValidLevel() (ret bool) {
	ret = contains(pa.validLevels(), pa.params.Level)
	return
}

func (pa *MarketParams) isValidResolution() (ret bool) {
	ret = contains(pa.validResolutions(), pa.params.Resolution)
	return
}

func (pa *MarketParams) validLevels() (ret []string) {
	levels := pa.params.VSetting.FixedViper.Get("fcoin.validation.params.level").([]interface{})
	for _, level := range levels {
		alevel, _ := level.(string)
		ret = append(ret, alevel)
	}
	return
}

func (pa *MarketParams) validResolutions() (ret []string) {
	resolutions := pa.params.VSetting.FixedViper.Get("fcoin.validation.params.resolution").([]interface{})
	for _, resolution := range resolutions {
		aresolution, _ := resolution.(string)
		ret = append(ret, aresolution)
	}
	return
}
