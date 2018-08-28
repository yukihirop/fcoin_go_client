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

func (mp *MarketParams) Messages() (ret map[string]string) {
	pa := mp.params
	if mp.IsValid() {
		ret = map[string]string{}
	}
	var results []map[string]string

	if !mp.isValidSymbol() {
		results = append(results, presenceErrorMessage(pa.Symbol, "symbol"))
	}

	switch pa.MethodName {
	case "MarketDepth":

		if !mp.isValidLevel() {
			results = append(results, includesErrorMessage(pa.Level, "level", mp.validLevels()))
		}
	case "MarketCandles":
		if !mp.isValidResolution() {
			results = append(results, includesErrorMessage(pa.Resolution, "resolution", mp.validResolutions()))
		}
	}
	ret = slice2map(results)
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
